// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {Params} from './Params.sol';
import {Validators} from './Validators.sol';

contract Proposal is Params {
    // How long a proposal will exist
    uint256 public proposalLastingPeriod;
    uint256 public punishThreshold;
    uint256 public removeThreshold;
    uint256 public decreaseRate;
    // Validator have to wait withdrawProfitPeriod blocks to withdraw his profits
    uint256 public withdrawProfitPeriod;
    // period time to increase aac
    uint256 public increasePeriod;
    // address to receive acc
    address public receiverAddr;

    // record
    mapping(address => bool) public pass;

    struct ProposalInfo {
        // who propose this proposal
        address proposer;
        // time create proposal
        uint256 createTime;
        uint256 proposalType;
        // validator proposal
        // propose who to be a validator
        address dst;
        // flag = true means add dst to validators
        // flag = false means remove dst from validators
        bool flag;
        // optional detail info of proposal
        string details;
        // config proposal
        // config id to update
        uint256 cid;
        // new value
        uint256 newValue;
    }

    struct ResultInfo {
        // number agree this proposal
        uint16 agree;
        // number reject this proposal
        uint16 reject;
        // means you can get proposal of current vote.
        bool resultExist;
    }

    struct VoteInfo {
        address voter;
        uint256 voteTime;
        bool auth;
    }

    mapping(bytes32 => ProposalInfo) public proposals;
    mapping(bytes32 => ResultInfo) public results;
    mapping(address => mapping(bytes32 => VoteInfo)) public votes;

    Validators validators;

    event LogCreateProposal(bytes32 indexed id, address indexed proposer, address indexed dst, bool flag, uint256 time);
    event LogCreateConfigProposal(
        bytes32 indexed id,
        address indexed proposer,
        uint256 cid,
        uint256 newValue,
        uint256 time
    );
    event LogVote(bytes32 indexed id, address indexed voter, bool auth, uint256 time);
    event LogPassProposal(bytes32 indexed id, uint256 time);
    event LogRejectProposal(bytes32 indexed id, uint256 time);
    event LogSetUnpassed(address indexed val, uint256 time);

    modifier onlyValidator() {
        require(validators.isActiveValidator(msg.sender), 'Validator only');
        _;
    }

    function initialize(
        address[] calldata vals,
        address _validators
    ) external onlyNotInitialized {
        require(_validators != address(0), "Invalid validators address");
        
        validators = Validators(_validators);
        for (uint256 i = 0; i < vals.length; i++) {
            require(vals[i] != address(0), 'Invalid validator address');
            pass[vals[i]] = true;
        }

        proposalLastingPeriod = 7 days;
        punishThreshold = 24;
        removeThreshold = 48;
        decreaseRate = 24;
        withdrawProfitPeriod = 86400;
        initialized = true;
        increasePeriod = 60*60*24*365; // Issuance period: 1 minute * 60 * 24*365
        receiverAddr = 0x9014B4DB9D30CeD67DB9d6B096f5DCDbA28cE639;
    }

    /**
     * @dev Efficiently compute hash for validator proposal
     * @param proposer Address of the proposer
     * @param dst Target validator address
     * @param flag Add/remove flag
     * @param details Proposal details
     * @param blockTimestamp Block timestamp
     * @return id The computed hash
     */
    function _hashValidatorProposal(
        address proposer,
        address dst,
        bool flag,
        string calldata details,
        uint256 blockTimestamp
    ) private pure returns (bytes32 id) {
        assembly {
            let ptr := mload(0x40)
            let detailsLen := details.length
            
            // Pack data tightly: proposer(20) + dst(20) + flag(1) + details + timestamp(32)
            let totalLen := add(0x49, detailsLen)  // 20 + 20 + 1 + detailsLen + 32
            
            // Store proposer (20 bytes)
            mstore(ptr, shl(96, proposer))
            
            // Store dst (20 bytes) 
            mstore(add(ptr, 0x14), shl(96, dst))
            
            // Store flag (1 byte)
            mstore8(add(ptr, 0x28), flag)
            
            // Copy details from calldata
            calldatacopy(add(ptr, 0x29), details.offset, detailsLen)
            
            // Store timestamp (32 bytes)
            mstore(add(ptr, add(0x29, detailsLen)), blockTimestamp)
            
            // Compute hash
            id := keccak256(ptr, totalLen)
        }
    }

    /**
     * @dev Efficiently compute hash for config proposal
     * @param proposer Address of the proposer
     * @param cid Configuration ID
     * @param newValue New configuration value
     * @param blockTimestamp Block timestamp
     * @return id The computed hash
     */
    function _hashConfigProposal(
        address proposer,
        uint256 cid,
        uint256 newValue,
        uint256 blockTimestamp
    ) private pure returns (bytes32 id) {
        assembly {
            let ptr := mload(0x40)
            
            // Pack data tightly: proposer(20) + cid(32) + newValue(32) + timestamp(32)
            // Total: 116 bytes
            
            // Store proposer (20 bytes)
            mstore(ptr, shl(96, proposer))
            
            // Store cid (32 bytes)
            mstore(add(ptr, 0x14), cid)
            
            // Store newValue (32 bytes)
            mstore(add(ptr, 0x34), newValue)
            
            // Store timestamp (32 bytes)
            mstore(add(ptr, 0x54), blockTimestamp)
            
            // Compute hash
            id := keccak256(ptr, 0x74)  // 116 bytes = 0x74
        }
    }

    function createProposal(
        address dst,
        bool flag,
        string calldata details
    ) external returns (bool) {
        // can't add a already dst or remove a not exist dst
        require(
            (!pass[dst] && flag) || (pass[dst] && !flag),
            'Cant"t add a already exist dst or Cant"t remove a not passed dst'
        );

        // generate proposal id using optimized hash function
        bytes32 id = _hashValidatorProposal(msg.sender, dst, flag, details, block.timestamp);
        require(bytes(details).length <= 3000, 'Details too long');
        require(proposals[id].createTime == 0, 'Proposal already exists');

        ProposalInfo memory proposal;
        proposal.proposer = msg.sender;
        proposal.dst = dst;
        proposal.flag = flag;
        proposal.details = details;
        proposal.createTime = block.timestamp;
        proposal.proposalType = 1;

        proposals[id] = proposal;
        emit LogCreateProposal(id, msg.sender, dst, flag, block.timestamp);
        return true;
    }

    function createUpdateConfigProposal(uint256 cid, uint256 newValue) external returns (bool) {
        // generate proposal id using optimized hash function
        bytes32 id = _hashConfigProposal(msg.sender, cid, newValue, block.timestamp);

        ProposalInfo memory proposal;
        proposal.proposer = msg.sender;
        proposal.cid = cid;
        proposal.newValue = newValue;
        proposal.createTime = block.timestamp;
        proposal.proposalType = 2;

        proposals[id] = proposal;
        emit LogCreateConfigProposal(id, msg.sender, cid, newValue, block.timestamp);
        return true;
    }

    function voteProposal(bytes32 id, bool auth) external onlyValidator returns (bool) {
        require(proposals[id].createTime != 0, 'Proposal not exist');
        require(votes[msg.sender][id].voteTime == 0, "You can't vote for a proposal twice");
        require(block.timestamp < proposals[id].createTime + proposalLastingPeriod, 'Proposal expired');

        votes[msg.sender][id].voteTime = block.timestamp;
        votes[msg.sender][id].voter = msg.sender;
        votes[msg.sender][id].auth = auth;
        emit LogVote(id, msg.sender, auth, block.timestamp);

        // update dst status if proposal is passed
        if (auth) {
            results[id].agree = results[id].agree + 1;
        } else {
            results[id].reject = results[id].reject + 1;
        }

        if (results[id].resultExist) {
            // do nothing if proposal already has result.
            return true;
        }

        if (results[id].agree >= validators.getActiveValidators().length / 2 + 1) {
            results[id].resultExist = true;

            if (proposals[id].proposalType == 1) {
                if (proposals[id].flag) {
                    // add to validators
                    pass[proposals[id].dst] = true;
                    // try to active validator if it isn't the first time
                    validators.tryActive(proposals[id].dst);
                } else {
                    pass[proposals[id].dst] = false;
                    validators.tryRemoveValidator(proposals[id].dst);
                }
            } else if (proposals[id].proposalType == 2) {
                updateConfig(proposals[id].cid, proposals[id].newValue);
            }

            emit LogPassProposal(id, block.timestamp);

            return true;
        }

        if (results[id].reject >= validators.getActiveValidators().length / 2 + 1) {
            results[id].resultExist = true;
            emit LogRejectProposal(id, block.timestamp);
        }

        return true;
    }

    function updateConfig(uint256 cid, uint256 value) private {
        if (cid == 0) {
            proposalLastingPeriod = value;
        } else if (cid == 1) {
            punishThreshold = value;
        } else if (cid == 2) {
            removeThreshold = value;
        } else if (cid == 3) {
            decreaseRate = value;
        } else if (cid == 4) {
            withdrawProfitPeriod = value;
        } else if (cid == 5) {
            increasePeriod = value;
        } else if (cid == 6) {
            receiverAddr = address(uint160(value));
        }
    }

    function setUnpassed(address val) external onlyValidatorsContract returns (bool) {
        // set validator unpass
        pass[val] = false;

        emit LogSetUnpassed(val, block.timestamp);
        return true;
    }
}
