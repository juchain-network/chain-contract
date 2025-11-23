package cmd

// 智能合约地址常量 - JuChain 系统合约地址
const (
	// 主要合约地址
	ValidatorContractAddr = "0x000000000000000000000000000000000000f000"
	PunishContractAddr    = "0x000000000000000000000000000000000000f001"
	ProposalContractAddr  = "0x000000000000000000000000000000000000f002"
	StakingContractAddr   = "0x000000000000000000000000000000000000f003"
)

// 默认配置常量
const (
	DefaultGasMultiplier = 120 // Gas估算乘数（百分比）
	DefaultTimeout       = 30  // 交易确认超时时间（秒）
	DefaultCheckInterval = 5   // 交易状态检查间隔（秒）
)

// 配置ID映射
var ConfigIDNames = map[int64]string{
	0: "proposalLastingPeriod", // 提案持续期
	1: "punishThreshold",       // 惩罚阈值
	2: "removeThreshold",       // 移除阈值
	3: "decreaseRate",          // 减少率
	4: "withdrawProfitPeriod",  // 提取收益周期
}

// 操作类型常量
const (
	OperationAdd    = "add"
	OperationRemove = "remove"
)

// 投票类型常量
const (
	VoteApprove = true
	VoteReject  = false
)

// 文件名模板常量 - 通用操作
const (
	CreateProposalFile       = "createProposal.json"
	CreateConfigProposalFile = "createUpdateConfigProposal.json"
	VoteProposalFile         = "voteProposal.json"
	WithdrawProfitsFile      = "withdrawProfits.json"
)

// 文件名模板常量 - Staking操作
const (
	RegisterValidatorFile = "registerValidator.json"
	EditValidatorFile     = "editValidator.json"
	DelegateFile          = "delegate.json"
	UndelegateFile        = "undelegate.json"
	ClaimRewardsFile      = "claimRewards.json"
)

// 状态码常量
const (
	ValidatorStatusActive   = 1
	ValidatorStatusInactive = 0
)

// Staking ABI 常量
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
		"inputs": [],
		"name": "getTopValidators",
		"outputs": [{"internalType": "address[]", "name": "", "type": "address[]"}],
		"stateMutability": "view",
		"type": "function"
	}
]`
