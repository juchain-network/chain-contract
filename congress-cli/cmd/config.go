package cmd

// 智能合约地址常量
const (
	ValidatorContractAddr = "0x000000000000000000000000000000000000f000"
	PunishContractAddr    = "0x000000000000000000000000000000000000f001"
	ProposalContractAddr  = "0x000000000000000000000000000000000000f002"
)

// 默认配置
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

// 操作类型
const (
	OperationAdd    = "add"
	OperationRemove = "remove"
)

// 投票类型
const (
	VoteApprove = true
	VoteReject  = false
)

// 文件名模板
const (
	CreateProposalFile       = "createProposal.json"
	CreateConfigProposalFile = "createUpdateConfigProposal.json"
	VoteProposalFile         = "voteProposal.json"
	WithdrawProfitsFile      = "withdrawProfits.json"
)

// 状态码
const (
	ValidatorStatusActive   = 1
	ValidatorStatusInactive = 0
)
