// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import {Params} from './Params.sol';
import {IValidators} from './IValidators.sol';

contract Proposal is Params {
    // How long a proposal will exist
    uint256 public proposalLastingPeriod;
    uint256 public punishThreshold;
    uint256 public removeThreshold;
    uint256 public decreaseRate;
    // Validator have to wait withdrawProfitPeriod blocks to withdraw his profits
    uint256 public withdrawProfitPeriod;
    // Block reward per block (in wei)
    uint256 public blockReward;
    // Unbonding period in blocks (time before delegators can withdraw undelegated funds)
    uint256 public unbondingPeriod;
    // Validator unjail period in blocks (time before jailed validator can unjail)
    uint256 public validatorUnjailPeriod;
    // Minimum staking amount to become a validator
    uint256 public minValidatorStake;
    // Maximum validators in active set
    uint256 public maxValidators;

    // record
    mapping(address => bool) public pass;
    // Record when proposal passed (for 7-day staking requirement)
    mapping(address => uint256) public proposalPassedTime;
    // Period for validator to stake after proposal passed (7 days)
    uint256 public constant STAKING_DEADLINE_PERIOD = 7 days;

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

    IValidators validators;

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
        _onlyValidator();
        _;
    }

    function _onlyValidator() internal view {
        require(validators.isActiveValidator(msg.sender), 'Validator only');
    }

    function initialize(
        address[] calldata vals,
        address _validators
    ) external onlyNotInitialized {
        require(_validators != address(0), "Invalid validators address");
        
        validators = IValidators(_validators);
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
        // Default block reward: 0.2 ether per block (17,280 JU/day ÷ 86,400 blocks/day)
        blockReward = 200_000_000_000_000_000; // 0.2 ether = 2 * 10^17 wei
        // Default unbonding period: 7 days in blocks (604800 blocks = 7 days * 24 hours * 3600 seconds / 1 second per block)
        unbondingPeriod = 604800;
        // Default validator unjail period: 24 hours in blocks (86400 blocks = 24 hours * 3600 seconds / 1 second per block)
        validatorUnjailPeriod = 86400;
        // Default minimum staking amount to become a validator: 100000 ether
        minValidatorStake = 100000 ether;
        // Default maximum number of validators in active set: 21
        maxValidators = 21;
        initialized = true;
    }

    function createProposal(
        address dst,
        bool flag,
        string calldata details
    ) external onlyValidator returns (bool) {
        // can't add an already exist dst or remove a not exist dst
        require(
            (!pass[dst] && flag) || (pass[dst] && !flag),
            "Can't add an already exist dst or Can't remove a not passed dst"
        );

        // generate proposal id
        // forge-lint: disable-next-line(asm-keccak256)
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

    function createUpdateConfigProposal(uint256 cid, uint256 newValue) external onlyValidator returns (bool) {
        // Validate config parameters before creating proposal
        validateConfig(cid, newValue);
        
        // forge-lint: disable-next-line(asm-keccak256)
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
        if (results[id].agree >= validators.getActiveValidatorCount() / 2 + 1) {
            results[id].resultExist = true;

            if (proposals[id].proposalType == 1) {
                if (proposals[id].flag) {
                    // add to validators
                    pass[proposals[id].dst] = true;
                    // Record proposal passed time for 7-day staking requirement
                    proposalPassedTime[proposals[id].dst] = block.timestamp;
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

        if (results[id].reject >= validators.getActiveValidatorCount() / 2 + 1) {
            results[id].resultExist = true;
            emit LogRejectProposal(id, block.timestamp);
        }
        
        return true;
    }

    /**
     * @dev Validate configuration parameters
     * @param cid Configuration ID:
     *   - 0: proposalLastingPeriod (1 hour - 30 days)
     *   - 1: punishThreshold (must > 0)
     *   - 2: removeThreshold (must > 0)
     *   - 3: decreaseRate (must > 0, prevents division by zero)
     *   - 4: withdrawProfitPeriod (must > 0)
     *   - 5: blockReward (must > 0, in wei)
     *   - 6: unbondingPeriod (must > 0)
     *   - 7: validatorUnjailPeriod (must > 0)
     *   - 8: minValidatorStake (must > 0, in wei)
     *   - 9: maxValidators (must > 0)
     * @param value New configuration value
     */
    function validateConfig(uint256 cid, uint256 value) internal pure returns (bool) {
        require(cid >= 0 && cid <= 9, "Invalid config ID");
        
        // Check specific rules
        if (cid == 0) {
            require(value >= 1 hours && value <= 30 days, "Invalid proposal period");
        } else {
            // All other configs require positive values
            require(value > 0, "Config value must be positive");
        }
        return true;
    }

    /**
     * @dev Update system configuration
     * @param cid Configuration ID
     * @param value New configuration value
     */
    function updateConfig(uint256 cid, uint256 value) private {
        validateConfig(cid, value);
        
        // Since validateConfig already checks cid is between 0-9, no need for else checks
        if (cid == 0) proposalLastingPeriod = value;
        if (cid == 1) punishThreshold = value;
        if (cid == 2) removeThreshold = value;
        if (cid == 3) decreaseRate = value;
        if (cid == 4) withdrawProfitPeriod = value;
        if (cid == 5) blockReward = value;
        if (cid == 6) unbondingPeriod = value;
        if (cid == 7) validatorUnjailPeriod = value;
        if (cid == 8) minValidatorStake = value;
        if (cid == 9) maxValidators = value;
    }

    /**
     * @dev Set validator as unpassed
     * @param val Validator address
     * @notice This function is called when validator is removed due to punishment
     * @notice Validator must pass a reproposal to regain validator status
     */
    function setUnpassed(address val) external onlyValidatorsContract returns (bool) {
        // set validator unpass
        pass[val] = false;
        proposalPassedTime[val] = 0; // Clear passed time
        
        emit LogSetUnpassed(val, block.timestamp);
        return true;
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
        // Check if within 7 days - only applies to NEW registrations
        return block.timestamp <= passedTime + STAKING_DEADLINE_PERIOD;
    }
}
