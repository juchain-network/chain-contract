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
    // Record when proposal passed (for 7-day staking requirement)
    mapping(address => uint256) public proposalPassedTime;
    // Period for validator to stake after proposal passed (7 days)
    uint256 public constant STAKING_DEADLINE_PERIOD = 7 days;
    
    // Maximum violations allowed for auto unjail without reproposal
    uint256 public constant MAX_VIOLATIONS_FOR_AUTO_UNJAIL = 3;
    
    // Violation count for graded punishment mechanism
    // violationCount[validator] = 0: No violations
    // violationCount[validator] = 1-3: Can auto unjail and restore pass
    // violationCount[validator] >= 4: Must repropose before unjail
    mapping(address => uint256) public violationCount;

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
    event LogViolationCountIncreased(address indexed val, uint256 newCount, uint256 time);
    event LogPassAutoRestored(address indexed val, uint256 time);
    event LogViolationCountReset(address indexed val, uint256 time);

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
            // Set proposalPassedTime for genesis validators (no 7-day limit for genesis)
            // This ensures consistency with normal proposal flow, even though genesis validators
            // don't need to pass isProposalValidForStaking() check (they use initializeWithValidators)
            proposalPassedTime[vals[i]] = block.timestamp;
        }

        proposalLastingPeriod = 7 days;
        punishThreshold = 24;
        removeThreshold = 48;
        decreaseRate = 24;
        withdrawProfitPeriod = 86400;
        initialized = true;
        increasePeriod = 60*60*24*365; //增发周期 1分钟 * 60 * 24*365
        receiverAddr = 0x9014B4DB9D30CeD67DB9d6B096f5DCDbA28cE639;
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

        // generate proposal id
        bytes32 id = keccak256(abi.encodePacked(msg.sender, dst, flag, details, block.timestamp));
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
        bytes32 id = keccak256(abi.encodePacked(msg.sender, cid, newValue, block.timestamp));

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

        // Check threshold using only votes from currently active validators
        // This ensures that votes from validators who were removed after voting are not counted
        uint256 activeAgree = getActiveVoteCount(id, true);
        uint256 activeReject = getActiveVoteCount(id, false);
        uint256 activeValidatorCount = validators.getActiveValidatorCount();
        uint256 threshold = activeValidatorCount / 2 + 1;

        if (activeAgree >= threshold) {
            results[id].resultExist = true;

            if (proposals[id].proposalType == 1) {
                if (proposals[id].flag) {
                    // add to validators
                    pass[proposals[id].dst] = true;
                    // Record proposal passed time for 7-day staking requirement
                    proposalPassedTime[proposals[id].dst] = block.timestamp;
                    // Reset violation count after proposal passed
                    if (violationCount[proposals[id].dst] > 0) {
                        violationCount[proposals[id].dst] = 0;
                        emit LogViolationCountReset(proposals[id].dst, block.timestamp);
                    }
                    // Validator needs to stake first, then will be activated at next epoch
                } else {
                    pass[proposals[id].dst] = false;
                    proposalPassedTime[proposals[id].dst] = 0; // Clear passed time
                    validators.tryRemoveValidator(proposals[id].dst);
                }
            } else if (proposals[id].proposalType == 2) {
                updateConfig(proposals[id].cid, proposals[id].newValue);
            }

            emit LogPassProposal(id, block.timestamp);

            return true;
        }

        if (activeReject >= threshold) {
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
            // casting to 'uint160' is safe because:
            // 1. Ethereum addresses are 20 bytes (160 bits), which matches uint160 exactly
            // 2. The value is expected to be a valid address passed through the proposal system
            // 3. Addresses are typically converted from address to uint256(uint160(addr)) when creating proposals
            // 4. This conversion is reversible and safe for valid Ethereum addresses
            // forge-lint: disable-next-line(unsafe-typecast)
            receiverAddr = address(uint160(value));
        }
    }

    /**
     * @dev Set validator as unpassed and increase violation count 
     * @param val Validator address
     * @notice This function is called when validator is removed due to punishment
     * @notice violationCount is incremented to track repeated violations
     */
    function setUnpassed(address val) external onlyValidatorsContract returns (bool) {
        // set validator unpass
        pass[val] = false;
        proposalPassedTime[val] = 0; // Clear passed time
        
        // Increment violation count for graded punishment mechanism
        violationCount[val] = violationCount[val] + 1;
        
        emit LogSetUnpassed(val, block.timestamp);
        emit LogViolationCountIncreased(val, violationCount[val], block.timestamp);
        return true;
    }
    
    /**
     * @dev Auto restore pass status for violators 
     * @param validator Validator address
     * @notice This function is called when validator unjails
     * @notice Only restores pass if violationCount <= 3 (允许 3 次以下违规自动恢复)
     * @return Whether pass status was restored
     */
    function autoRestorePass(address validator) external onlyValidatorsContract returns (bool) {
        // Auto restore for violations <= MAX_VIOLATIONS_FOR_AUTO_UNJAIL
        if (violationCount[validator] <= MAX_VIOLATIONS_FOR_AUTO_UNJAIL && violationCount[validator] > 0) {
            pass[validator] = true;
            // Set proposal passed time to current time (no 7-day wait for auto restore)
            proposalPassedTime[validator] = block.timestamp;
            emit LogPassAutoRestored(validator, block.timestamp);
            return true;
        }
        return false;
    }
    
    /**
     * @dev Reset violation count after validator runs normally for N epochs 
     * @param validator Validator address
     * @notice This function is called when validator has been running normally
     * @notice Allows validator to get a fresh start after demonstrating good behavior
     */
    function resetViolationCount(address validator) external onlyValidatorsContract returns (bool) {
        if (violationCount[validator] > 0) {
            violationCount[validator] = 0;
            emit LogViolationCountReset(validator, block.timestamp);
            return true;
        }
        return false;
    }
    
    /**
     * @dev Get violation count for a validator
     * @param validator Validator address
     * @return Current violation count
     */
    function getViolationCount(address validator) external view returns (uint256) {
        return violationCount[validator];
    }

    /**
     * @dev Get count of votes from currently active validators only
     * @param id Proposal ID
     * @param isAgree Whether to count agree votes (true) or reject votes (false)
     * @return Count of votes from currently active validators
     */
    function getActiveVoteCount(bytes32 id, bool isAgree) internal view returns (uint256) {
        address[] memory activeValidators = validators.getActiveValidators();
        uint256 count = 0;
        
        for (uint256 i = 0; i < activeValidators.length; i++) {
            address voter = activeValidators[i];
            VoteInfo memory vote = votes[voter][id];
            // Check if this validator voted and the vote matches the requested type
            if (vote.voteTime != 0 && vote.auth == isAgree) {
                count++;
            }
        }
        
        return count;
    }

    /**
     * @dev Check if proposal passed time is still valid (within 7 days) for NEW registration
     * @notice This function is ONLY used to check if a NEW validator can register within 7 days
     * @notice Once a validator is registered (has selfStake >= MIN_VALIDATOR_STAKE), 
     *         this check is no longer applied. Validators are removed by setUnpassed() when punished.
     * @param validator Validator address
     * @return Whether the proposal is still valid for new staking registration
     */
    function isProposalValidForStaking(address validator) external view returns (bool) {
        if (!pass[validator]) {
            return false;
        }
        uint256 passedTime = proposalPassedTime[validator];
        if (passedTime == 0) {
            return false; // No proposal passed time recorded
        }
        // Check if within 7 days - only applies to NEW registrations
        return block.timestamp <= passedTime + STAKING_DEADLINE_PERIOD;
    }
}
