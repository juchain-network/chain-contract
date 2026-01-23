// SPDX-License-Identifier: MIT

pragma solidity ^0.8.29;

import {Params} from "./Params.sol";
import {IValidators} from "./IValidators.sol";
import {ReentrancyGuard} from "@openzeppelin/contracts/utils/ReentrancyGuard.sol";

contract Proposal is Params, ReentrancyGuard {
    // Default configuration constants
    uint256 private constant DEFAULT_PROPOSAL_LASTING_PERIOD = 604800; // 7 days in blocks
    uint256 private constant DEFAULT_PUNISH_THRESHOLD = 24; // blocks
    uint256 private constant DEFAULT_REMOVE_THRESHOLD = 48; // blocks
    uint256 private constant DEFAULT_DECREASE_RATE = 24; // %
    uint256 private constant DEFAULT_WITHDRAW_PROFIT_PERIOD = 86400; // 1 day in blocks
    uint256 private constant DEFAULT_BLOCK_REWARD = 0.2 ether; // 2 * 10^17 wei
    uint256 private constant DEFAULT_UNBONDING_PERIOD = 604800; // 7 days in blocks
    uint256 private constant DEFAULT_VALIDATOR_UNJAIL_PERIOD = 86400; // 1 day in blocks
    uint256 private constant DEFAULT_MIN_VALIDATOR_STAKE = 100000 ether; // Minimum validator stake
    uint256 private constant DEFAULT_MAX_VALIDATORS = 21; // Maximum active validators
    uint256 private constant DEFAULT_MIN_DELEGATION = 10 ether; // 10 JU
    uint256 private constant DEFAULT_MIN_UNDELEGATION = 1 ether; // 1 JU
    uint256 private constant DEFAULT_DOUBLE_SIGN_SLASH_AMOUNT = 50000 ether;
    uint256 private constant DEFAULT_DOUBLE_SIGN_REWARD_AMOUNT = 10000 ether;
    uint256 private constant DEFAULT_DOUBLE_SIGN_WINDOW = 86400; // 1 day in blocks
    uint256 private constant DEFAULT_COMMISSION_UPDATE_COOLDOWN = 604800; // 7 days in blocks
    uint256 private constant DEFAULT_BASE_REWARD_RATIO = 3000; // 30.00%
    uint256 private constant DEFAULT_MAX_COMMISSION_RATE = 6000; // 60.00%
    uint256 private constant DEFAULT_PROPOSAL_COOLDOWN = 100; // 100 blocks
    address private constant DEFAULT_BURN_ADDRESS = 0x000000000000000000000000000000000000dEaD;

    // How many blocks a proposal will exist
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
    // Minimum delegation amount per delegator
    uint256 public minDelegation;
    // Minimum undelegation amount per delegator
    uint256 public minUndelegation;
    // Double-sign slash amount (absolute, in wei)
    uint256 public doubleSignSlashAmount;
    // Double-sign reporter reward amount (absolute, in wei)
    uint256 public doubleSignRewardAmount;
    // Double-sign evidence window (in blocks)
    uint256 public doubleSignWindow;
    // Burn address for slashed funds after reward
    address public burnAddress;

    // Commission update cooldown (in blocks)
    uint256 public commissionUpdateCooldown;
    // Base reward ratio (0-100)
    uint256 public baseRewardRatio;
    // Max commission rate (0-10000)
    uint256 public maxCommissionRate;
    // Proposal cooldown (in blocks)
    uint256 public proposalCooldown;

    // Validator address => last proposal block number
    mapping(address => uint256) public lastProposalBlock;

    // record
    mapping(address => bool) public pass;
    // Record when proposal passed (for 7-day staking requirement)
    mapping(address => uint256) public proposalPassedHeight;
    // Proposal nonce per proposer to ensure unique proposal IDs
    mapping(address => uint256) public proposerNonces;

    struct ProposalInfo {
        // who propose this proposal
        address proposer;
        // time create proposal
        uint256 createTime;
        // block number when proposal was created
        uint256 createBlock;
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
    uint256 public revision;
    uint256[50] private __gap;

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
        require(validators.isValidatorActive(msg.sender), "Validator only");
    }

    /**
     * @dev Initializes the Proposal contract with validators and default parameters.
     * @param vals Array of initial validator addresses.
     * @param validators_ Address of the Validators contract.
     */
    function initialize(
        address[] calldata vals,
        address validators_,
        uint256 epoch_
    ) external onlyNotInitialized {
        require(validators_ != address(0), "Invalid validators address");

        _initializeEpoch(epoch_);
        validators = IValidators(validators_);
        for (uint256 i = 0; i < vals.length; i++) {
            require(vals[i] != address(0), "Invalid validator address");
            pass[vals[i]] = true;
            // Set proposalPassedHeight for genesis validators (uses proposalLastingPeriod for block-based validation)
            // This ensures consistency with normal proposal flow, even though genesis validators
            // don't need to pass isProposalValidForStaking() check (they use initializeWithValidators)
            proposalPassedHeight[vals[i]] = block.number;
        }

        proposalLastingPeriod = DEFAULT_PROPOSAL_LASTING_PERIOD;
        punishThreshold = DEFAULT_PUNISH_THRESHOLD;
        removeThreshold = DEFAULT_REMOVE_THRESHOLD;
        decreaseRate = DEFAULT_DECREASE_RATE;
        withdrawProfitPeriod = DEFAULT_WITHDRAW_PROFIT_PERIOD;
        blockReward = DEFAULT_BLOCK_REWARD;
        unbondingPeriod = DEFAULT_UNBONDING_PERIOD;
        validatorUnjailPeriod = DEFAULT_VALIDATOR_UNJAIL_PERIOD;
        minValidatorStake = DEFAULT_MIN_VALIDATOR_STAKE;
        maxValidators = DEFAULT_MAX_VALIDATORS;
        minDelegation = DEFAULT_MIN_DELEGATION;
        minUndelegation = DEFAULT_MIN_UNDELEGATION;
        doubleSignSlashAmount = DEFAULT_DOUBLE_SIGN_SLASH_AMOUNT;
        doubleSignRewardAmount = DEFAULT_DOUBLE_SIGN_REWARD_AMOUNT;
        doubleSignWindow = DEFAULT_DOUBLE_SIGN_WINDOW;
        commissionUpdateCooldown = DEFAULT_COMMISSION_UPDATE_COOLDOWN;
        baseRewardRatio = DEFAULT_BASE_REWARD_RATIO;
        maxCommissionRate = DEFAULT_MAX_COMMISSION_RATE;
        proposalCooldown = DEFAULT_PROPOSAL_COOLDOWN;
        burnAddress = DEFAULT_BURN_ADDRESS;
        revision = 1;
        initialized = true;
    }

    /**
     * @dev Reinitialize for upgrades (v2).
     * @notice Only miner can call. Can be called once.
     */
    function reinitializeV2() external onlyInitialized onlyMiner {
        require(revision < 2, "Already reinitialized");
        revision = 2;
    }

    /**
     * @dev Creates a new proposal for validator management.
     * @param dst Address of the validator being proposed for addition or removal.
     * @param flag Boolean indicating whether to add (true) or remove (false) the validator.
     * @param details Description of the proposal.
     * @return bytes32 Unique identifier for the created proposal.
     */
    function createProposal(
        address dst,
        bool flag,
        string calldata details
    ) external onlyValidator returns (bytes32) {
        _checkProposalCooldown();
        // Only add additional checks for add proposals
        if (flag) {
            // Check if validator is already in top validator set
            bool isTop = validators.isTopValidator(dst);

            // Only block add proposals for validators already in top set
            if (isTop) {
                revert("Validator is already in top validator set");
            }

            // If proposal was passed before, check if it's expired
            if (pass[dst]) {
                uint256 passedHeight = proposalPassedHeight[dst];
                // If proposal has expired, clear the pass status to allow resubmission
                if (block.number > passedHeight + proposalLastingPeriod) {
                    pass[dst] = false;
                    proposalPassedHeight[dst] = 0;
                } else {
                    // Proposal is still valid, can't resubmit add proposal
                    revert("Can't add an already passed dst");
                }
            }
        }
        // Simplified requirement: only check for add proposals, remove proposals can be resubmitted freely
        require(
            (!pass[dst] && flag) || !flag,
            "Can't add an already exist dst"
        );

        // Get current nonce for the proposer
        uint256 currentNonce = proposerNonces[msg.sender];

        // generate proposal id using nonce instead of block.timestamp
        // forge-lint: disable-next-line(asm-keccak256)
        bytes32 id = keccak256(abi.encode(msg.sender, dst, flag, details, currentNonce));
        require(bytes(details).length <= 3000, "Details too long");
        require(proposals[id].createTime == 0, "Proposal already exists");

        // Increment nonce for the proposer
        proposerNonces[msg.sender]++;

        ProposalInfo memory proposal = ProposalInfo({proposer: address(0), createTime: 0, createBlock: 0, proposalType: 0, dst: address(0), flag: false, details: "", cid: 0, newValue: 0});
        proposal.proposer = msg.sender;
        proposal.dst = dst;
        proposal.flag = flag;
        proposal.details = details;
        proposal.createTime = block.timestamp;
        proposal.createBlock = block.number;
        proposal.proposalType = 1;

        proposals[id] = proposal;
        emit LogCreateProposal(id, msg.sender, dst, flag, block.timestamp);
        return id;
    }

    /**
     * @dev Creates a proposal to update system configuration parameters.
     * @param cid Configuration parameter ID.
     * @param newValue New value for the configuration parameter.
     * @return bytes32 Unique identifier for the created proposal.
     */
    function createUpdateConfigProposal(uint256 cid, uint256 newValue) external onlyValidator returns (bytes32) {
        _checkProposalCooldown();
        // Validate config parameters before creating proposal
        require(validateConfig(cid, newValue), "Config validation failed");

        // Get current nonce for the proposer
        uint256 currentNonce = proposerNonces[msg.sender];

        // generate proposal id using nonce instead of block.timestamp
        // forge-lint: disable-next-line(asm-keccak256)
        bytes32 id = keccak256(abi.encode(msg.sender, cid, newValue, currentNonce));

        // Increment nonce for the proposer
        proposerNonces[msg.sender]++;

        ProposalInfo memory proposal = ProposalInfo({proposer: address(0), createTime: 0, createBlock: 0, proposalType: 0, dst: address(0), flag: false, details: "", cid: 0, newValue: 0});
        proposal.proposer = msg.sender;
        proposal.cid = cid;
        proposal.newValue = newValue;
        proposal.createTime = block.timestamp;
        proposal.createBlock = block.number;
        proposal.proposalType = 2;

        proposals[id] = proposal;
        emit LogCreateConfigProposal(id, msg.sender, cid, newValue, block.timestamp);
        return id;
    }

    /**
     * @dev Casts a vote on a proposal.
     * @param id Unique identifier of the proposal.
     * @param auth Boolean indicating approval (true) or rejection (false) of the proposal.
     * @return bool Returns true if the vote was successful.
     */
    function voteProposal(bytes32 id, bool auth) external onlyValidator onlyNotEpoch nonReentrant returns (bool) {
        require(proposals[id].createTime != 0, "Proposal does not exist");
        require(votes[msg.sender][id].voteTime == 0, "You can't vote for a proposal twice");
        require(block.number < proposals[id].createBlock + proposalLastingPeriod, "Proposal expired");

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
        uint256 votingCount = validators.getVotingValidatorCount();
        uint256 threshold = votingCount / 2 + 1;
        if (results[id].agree >= threshold) {
            results[id].resultExist = true;

            // Handle different proposal types with if-else statements
            uint256 proposalType = proposals[id].proposalType;
            if (proposalType == 1) {
                if (proposals[id].flag) {
                    // add to validators
                    pass[proposals[id].dst] = true;
                    // Record proposal passed height for 7-day staking requirement
                    proposalPassedHeight[proposals[id].dst] = block.number;
                    // Validator needs to stake first, then will be activated at next epoch
                } else {
                    pass[proposals[id].dst] = false;
                    proposalPassedHeight[proposals[id].dst] = 0; // Clear passed height
                    validators.tryRemoveValidator(proposals[id].dst);
                }
            } else if (proposalType == 2) {
                updateConfig(proposals[id].cid, proposals[id].newValue);
            } else {
                revert("Invalid proposal type");
            }

            emit LogPassProposal(id, block.timestamp);

            return true;
        }

        if (results[id].reject >= threshold) {
            results[id].resultExist = true;
            emit LogRejectProposal(id, block.timestamp);
        }

        return true;
    }

    /**
     * @dev Validate configuration parameters
     * @param cid Configuration ID:
     *   - 0: proposalLastingPeriod (must > 0)
     *   - 1: punishThreshold (must > 0)
     *   - 2: removeThreshold (must > 0)
     *   - 3: decreaseRate (must > 0, prevents division by zero)
     *   - 4: withdrawProfitPeriod (must > 0)
     *   - 5: blockReward (must > 0, in wei)
     *   - 6: unbondingPeriod (must > 0)
     *   - 7: validatorUnjailPeriod (must > 0)
     *   - 8: minValidatorStake (must > 0, in wei)
     *   - 9: maxValidators (must > 0)
     *   - 10: minDelegation (must > 0, in wei)
     *   - 11: minUndelegation (must > 0, in wei)
     *   - 12: doubleSignSlashAmount (must > 0)
     *   - 13: doubleSignRewardAmount (must > 0)
     *   - 14: burnAddress (must be non-zero)
     *   - 15: doubleSignWindow (must > 0)
     *   - 16: commissionUpdateCooldown (must > 0)
     *   - 17: baseRewardRatio (must <= 10000)
     *   - 18: maxCommissionRate (must <= 10000)
     * @param value New configuration value
     */
    function validateConfig(uint256 cid, uint256 value) internal view returns (bool) {
        require(cid <= 19, "Invalid config ID");
        require(value > 0, "Config value must be positive");
        if (cid == 1) {
            require(value < removeThreshold, "punishThreshold must be < removeThreshold");
        } else if (cid == 2) {
            require(punishThreshold < value, "removeThreshold must be > punishThreshold");
            require(decreaseRate <= value, "removeThreshold must be >= decreaseRate");
        } else if (cid == 3) {
            require(value <= removeThreshold, "decreaseRate must be <= removeThreshold");
        } else if (cid == 9) {
            require(value <= CONSENSUS_MAX_VALIDATORS, "maxValidators exceeds consensus limit");
        } else if (cid == 12) {
            require(value >= doubleSignRewardAmount, "doubleSignSlashAmount must be >= doubleSignRewardAmount");
        } else if (cid == 13) {
            require(value <= doubleSignSlashAmount, "doubleSignRewardAmount must be <= doubleSignSlashAmount");
        } else if (cid == 14) {
            require(value <= type(uint160).max, "burnAddress invalid");
            require(address(uint160(value)) != address(0), "burnAddress must be non-zero");
        } else if (cid == 15) {
            require(value > 0, "doubleSignWindow must be positive");
        } else if (cid == 16) {
            require(value > 0, "commissionUpdateCooldown must be positive");
        } else if (cid == 17) {
            require(value <= 10000, "baseRewardRatio must be <= 10000");
        } else if (cid == 18) {
            require(value <= 10000, "maxCommissionRate must be <= 10000");
        } else if (cid == 19) {
            require(value > 0, "proposalCooldown must be positive");
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

        // Use if-else statements for better robustness
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
            blockReward = value;
        } else if (cid == 6) {
            unbondingPeriod = value;
        } else if (cid == 7) {
            validatorUnjailPeriod = value;
        } else if (cid == 8) {
            minValidatorStake = value;
        } else if (cid == 9) {
            maxValidators = value;
        } else if (cid == 10) {
            minDelegation = value;
        } else if (cid == 11) {
            minUndelegation = value;
        } else if (cid == 12) {
            doubleSignSlashAmount = value;
        } else if (cid == 13) {
            doubleSignRewardAmount = value;
        } else if (cid == 14) {
            burnAddress = address(uint160(value));
        } else if (cid == 15) {
            doubleSignWindow = value;
        } else if (cid == 16) {
            commissionUpdateCooldown = value;
        } else if (cid == 17) {
            baseRewardRatio = value;
        } else if (cid == 18) {
            maxCommissionRate = value;
        } else if (cid == 19) {
            proposalCooldown = value;
        } else {
            revert("Unknown config ID"); // Fail fast for new config IDs
        }
    }

    /**
     * @dev Sets a validator as unpassed (ineligible).
     * @param val Address of the validator to mark as unpassed.
     * @return bool Returns true if the operation was successful.
     * @notice This function is called when validator is removed due to punishment.
     * @notice Validator must pass a reproposal to regain validator status.
     */
    function setUnpassed(address val) external onlyValidatorsContract returns (bool) {
        // set validator unpass
        pass[val] = false;
        proposalPassedHeight[val] = 0; // Clear passed height

        emit LogSetUnpassed(val, block.timestamp);
        return true;
    }

    /**
     * @dev Checks if a validator's proposal is valid for staking.
     * @param validator Address of the validator to check.
     * @return bool Returns true if the validator's proposal is valid for staking.
     * @notice This function is ONLY used to check if a NEW validator can register within the specified block period.
     * @notice Once a validator is registered (has selfStake >= MIN_VALIDATOR_STAKE), this check is no longer applied.
     * @notice Validators are removed by setUnpassed() when punished.
     */
    function isProposalValidForStaking(address validator) external view returns (bool) {
        if (!pass[validator]) {
            return false;
        }
        uint256 passedHeight = proposalPassedHeight[validator];
        // Check if within 7 days (604800 blocks) - only applies to NEW registrations
        return block.number <= passedHeight + proposalLastingPeriod;
    }

    /**
     * @dev Internal function to check and update proposal cooldown for a validator.
     */
    function _checkProposalCooldown() internal {
        uint256 lastBlock = lastProposalBlock[msg.sender];
        if (lastBlock > 0) {
            require(block.number >= lastBlock + proposalCooldown, "Proposal creation too frequent");
        }
        lastProposalBlock[msg.sender] = block.number;
    }
}
