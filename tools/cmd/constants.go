package cmd

// Smart contract address constants - JuChain system contract addresses
const (
	// Main contract addresses
	ValidatorContractAddr = "0x000000000000000000000000000000000000f010"
	PunishContractAddr    = "0x000000000000000000000000000000000000f011"
	ProposalContractAddr  = "0x000000000000000000000000000000000000f012"
	StakingContractAddr   = "0x000000000000000000000000000000000000f013"
)

// Default configuration constants
const (
	DefaultGasMultiplier = 120 // Gas estimation multiplier (percentage)
	DefaultTimeout       = 30  // Transaction confirmation timeout (seconds)
	DefaultCheckInterval = 5   // Transaction status check interval (seconds)
)

// Configuration ID mapping
var ConfigIDNames = map[int64]string{
	0: "proposalLastingPeriod", // Proposal duration
	1: "punishThreshold",       // Punishment threshold
	2: "removeThreshold",       // Removal threshold
	3: "decreaseRate",          // Decrease rate
	4: "withdrawProfitPeriod",  // Profit withdrawal period
}

// Operation type constants
const (
	OperationAdd    = "add"
	OperationRemove = "remove"
)

// Vote type constants
const (
	VoteApprove = true
	VoteReject  = false
)

// Filename template constants - Generic operations
const (
	CreateProposalFile       = "createProposal.json"
	CreateConfigProposalFile = "createUpdateConfigProposal.json"
	VoteProposalFile         = "voteProposal.json"
	WithdrawProfitsFile      = "withdrawProfits.json"
)

// Filename template constants - Staking operations
const (
	RegisterValidatorFile = "registerValidator.json"
	EditValidatorFile     = "editValidator.json"
	DelegateFile          = "delegate.json"
	UndelegateFile        = "undelegate.json"
	ClaimRewardsFile      = "claimRewards.json"
)

// Status code constants
const (
	ValidatorStatusActive   = 1
	ValidatorStatusInactive = 0
)

// Staking ABI constants
const stakingABI = `[
	{
		"inputs": [{"internalType": "uint256", "name": "commissionRate", "type": "uint256"}],
		"name": "registerValidator",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [{"internalType": "address", "name": "validator", "type": "address"}],
		"name": "delegate",
		"outputs": [],
		"stateMutability": "payable",
		"type": "function"
	},
	{
		"inputs": [
			{"internalType": "address", "name": "validator", "type": "address"},
			{"internalType": "uint256", "name": "amount", "type": "uint256"}
		],
		"name": "undelegate",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [{"internalType": "address", "name": "validator", "type": "address"}],
		"name": "claimRewards",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [{"internalType": "address", "name": "validator", "type": "address"}],
		"name": "getValidatorInfo",
		"outputs": [
			{"internalType": "uint256", "name": "selfStake", "type": "uint256"},
			{"internalType": "uint256", "name": "totalDelegated", "type": "uint256"},
			{"internalType": "uint256", "name": "commissionRate", "type": "uint256"},
			{"internalType": "bool", "name": "isJailed", "type": "bool"},
			{"internalType": "uint256", "name": "jailUntilBlock", "type": "uint256"}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{"internalType": "address", "name": "delegator", "type": "address"},
			{"internalType": "address", "name": "validator", "type": "address"}
		],
		"name": "getDelegationInfo",
		"outputs": [
			{"internalType": "uint256", "name": "amount", "type": "uint256"},
			{"internalType": "uint256", "name": "pendingRewards", "type": "uint256"},
			{"internalType": "uint256", "name": "unbondingAmount", "type": "uint256"},
			{"internalType": "uint256", "name": "unbondingBlock", "type": "uint256"}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [{"internalType": "address[]", "name": "validators", "type": "address[]"}],
		"name": "getTopValidators",
		"outputs": [{"internalType": "address[]", "name": "", "type": "address[]"}],
		"stateMutability": "view",
		"type": "function"
	}
]`

// Validators contract ABI for getTopValidators
const validatorsABI = `[
	{
		"inputs": [],
		"name": "getTopValidators",
		"outputs": [{"internalType": "address[]", "name": "", "type": "address[]"}],
		"stateMutability": "view",
		"type": "function"
	}
]`
