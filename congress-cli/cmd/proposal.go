package cmd

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"juchain.org/chain/congress-cli/contracts/generated"
)

func CreateProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create_proposal",
		Short: "create proposal tx",
		Run:   createProposalTx,
	}
	createProposalFlags(cmd)
	return cmd
}

func createProposalFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("proposer", "p", "", "proposer addr (must be valid validator)")
	_ = cmd.MarkFlagRequired("proposer")
	cmd.Flags().StringP("target", "t", "", "target addr (to be a validator)")
	_ = cmd.MarkFlagRequired("target")
	cmd.Flags().StringP("operation", "o", "", "operation add|remove")
	_ = cmd.MarkFlagRequired("operation")
}

func createProposalTx(cmd *cobra.Command, _ []string) {
	rpc := GetRPCEndpoint(cmd) // Use config-aware function
	proposer, _ := cmd.Flags().GetString("proposer")
	target, _ := cmd.Flags().GetString("target")
	operation, _ := cmd.Flags().GetString("operation")

	// 验证输入参数
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateAddresses(proposer, target); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateOperation(operation); err != nil {
		PrintValidationError(err)
		return
	}

	flag := operation == OperationAdd
	if flag {
		PrintInfo("Creating proposal to add new validator")
	} else {
		PrintInfo("Creating proposal to remove validator")
	}

	if err := innerCreateProposal(proposer, target, flag, rpc); err != nil {
		PrintError("Failed to create proposal", err)
		return
	}
}

func innerCreateProposal(proposer, target string, flag bool, rpc string) error {
	proposalAbi, err := abi.JSON(strings.NewReader(generated.ProposalABI))
	if err != nil {
		return fmt.Errorf("failed to parse proposal ABI: %w", err)
	}

	abiData, err := proposalAbi.Pack("createProposal", common.HexToAddress(target), flag, "")
	if err != nil {
		return fmt.Errorf("failed to pack proposal data: %w", err)
	}

	err = CreateRawTx(common.HexToAddress(proposer), common.HexToAddress(ProposalContractAddr), nil, abiData, rpc, CreateProposalFile)
	if err != nil {
		return fmt.Errorf("failed to create raw transaction: %w", err)
	}

	PrintSuccess("Proposal transaction created successfully!")
	PrintInfo(fmt.Sprintf("Transaction file: %s", CreateProposalFile))
	return nil
}

func CreateConfigProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create_config_proposal",
		Short: "create update config proposal tx",
		Run:   createConfigProposalTx,
	}
	createConfigProposalFlags(cmd)
	return cmd
}

func createConfigProposalFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("proposer", "p", "", "proposer addr (must be valid validator)")
	_ = cmd.MarkFlagRequired("proposer")
	cmd.Flags().Int64P("cid", "i", 0, "config id (0 proposalLastingPeriod, 1 punishThreshold, 2 removeThreshold, 3 decreaseRate, 4 withdrawProfitPeriod)")
	_ = cmd.MarkFlagRequired("cid")
	cmd.Flags().Int64P("value", "v", 0, "new config value")
	_ = cmd.MarkFlagRequired("value")
}

func createConfigProposalTx(cmd *cobra.Command, _ []string) {
	rpc := GetRPCEndpoint(cmd) // Use config-aware function
	proposer, _ := cmd.Flags().GetString("proposer")
	cid, _ := cmd.Flags().GetInt64("cid")
	cvalue, _ := cmd.Flags().GetInt64("value")

	// 验证输入参数
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateAddress(proposer); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateConfigID(cid); err != nil {
		PrintValidationError(err)
		return
	}

	if cvalue < 0 {
		PrintValidationError(fmt.Errorf("config value must be non-negative: %d", cvalue))
		return
	}

	PrintInfo(fmt.Sprintf("Creating config update proposal for %s (ID: %d) with value: %d",
		GetConfigIDName(cid), cid, cvalue))

	if err := innerCreateConfigProposal(proposer, cid, cvalue, rpc); err != nil {
		PrintError("Failed to create config proposal", err)
		return
	}
}

func innerCreateConfigProposal(proposer string, cid, cvalue int64, rpc string) error {
	proposalAbi, err := abi.JSON(strings.NewReader(generated.ProposalABI))
	if err != nil {
		return fmt.Errorf("failed to parse proposal ABI: %w", err)
	}

	abiData, err := proposalAbi.Pack("createUpdateConfigProposal", big.NewInt(cid), big.NewInt(cvalue))
	if err != nil {
		return fmt.Errorf("failed to pack config proposal data: %w", err)
	}

	err = CreateRawTx(common.HexToAddress(proposer), common.HexToAddress(ProposalContractAddr), nil, abiData, rpc, CreateConfigProposalFile)
	if err != nil {
		return fmt.Errorf("failed to create raw transaction: %w", err)
	}

	PrintSuccess("Config proposal transaction created successfully!")
	PrintInfo(fmt.Sprintf("Transaction file: %s", CreateConfigProposalFile))
	return nil
}

func VoteProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "vote_proposal",
		Short: "vote proposal tx",
		Run:   voteProposalTx,
	}
	voteProposalCmdFlags(cmd)
	return cmd
}

func voteProposalCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("signer", "s", "", "signer addr (must be valid validator)")
	_ = cmd.MarkFlagRequired("signer")
	cmd.Flags().StringP("proposalId", "i", "", "proposal id (64-character hex string)")
	_ = cmd.MarkFlagRequired("proposalId")
	cmd.Flags().BoolP("approve", "a", false, "approve this proposal (use -a for YES, omit for NO)")
}

func voteProposalTx(cmd *cobra.Command, _ []string) {
	rpc := GetRPCEndpoint(cmd) // Use config-aware function
	signer, _ := cmd.Flags().GetString("signer")
	proposalId, _ := cmd.Flags().GetString("proposalId")
	approve, _ := cmd.Flags().GetBool("approve")

	// 验证输入参数
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateAddress(signer); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateProposalID(proposalId); err != nil {
		PrintValidationError(err)
		return
	}

	voteType := "REJECT"
	if approve {
		voteType = "APPROVE"
	}
	PrintInfo(fmt.Sprintf("Voting %s on proposal: %s", voteType, proposalId))

	if err := innerVoteProposal(signer, proposalId, approve, rpc); err != nil {
		PrintError("Failed to vote on proposal", err)
		return
	}
}

func innerVoteProposal(signer, proposalId string, flag bool, rpc string) error {
	proposalAbi, err := abi.JSON(strings.NewReader(generated.ProposalABI))
	if err != nil {
		return fmt.Errorf("failed to parse proposal ABI: %w", err)
	}

	var proposalIdBytes [32]byte
	copy(proposalIdBytes[:], common.HexToHash(proposalId).Bytes())

	abiData, err := proposalAbi.Pack("voteProposal", proposalIdBytes, flag)
	if err != nil {
		return fmt.Errorf("failed to pack vote proposal data: %w", err)
	}

	err = CreateRawTx(common.HexToAddress(signer), common.HexToAddress(ProposalContractAddr), nil, abiData, rpc, VoteProposalFile)
	if err != nil {
		return fmt.Errorf("failed to create vote transaction: %w", err)
	}

	PrintSuccess("Vote transaction created successfully!")
	PrintInfo(fmt.Sprintf("Transaction file: %s", VoteProposalFile))
	return nil
}

// QueryProposalCmd creates a command to query a specific proposal
func QueryProposalCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposal",
		Short: "query proposal by ID",
		Run:   queryProposalTx,
	}
	queryProposalFlags(cmd)
	return cmd
}

func queryProposalFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("id", "i", "", "proposal ID (64-character hex string)")
	_ = cmd.MarkFlagRequired("id")
}

func queryProposalTx(cmd *cobra.Command, _ []string) {
	rpc := GetRPCEndpoint(cmd)
	proposalId, _ := cmd.Flags().GetString("id")

	// 验证输入参数
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateProposalID(proposalId); err != nil {
		PrintValidationError(err)
		return
	}

	PrintInfo("Fetching proposal details...")

	if err := innerQueryProposal(proposalId, rpc); err != nil {
		PrintError("Failed to query proposal", err)
		return
	}
}

func innerQueryProposal(proposalId string, rpc string) error {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return fmt.Errorf("failed to connect to RPC: %w", err)
	}
	defer client.Close()

	proposalContract, err := generated.NewProposal(common.HexToAddress(ProposalContractAddr), client)
	if err != nil {
		return fmt.Errorf("failed to instantiate proposal contract: %w", err)
	}

	var proposalIdBytes [32]byte
	copy(proposalIdBytes[:], common.HexToHash(proposalId).Bytes())

	proposal, err := proposalContract.Proposals(nil, proposalIdBytes)
	if err != nil {
		return fmt.Errorf("failed to query proposal: %w", err)
	}

	// 显示提案信息
	fmt.Println("📋 Proposal Details:")
	fmt.Printf("Proposal ID: %s\n", proposalId)
	fmt.Printf("Proposer: %s (验证者地址)\n", proposal.Proposer.Hex())

	// 根据提案类型显示不同信息
	if proposal.ProposalType.Int64() == 1 { // 验证者管理提案
		if proposal.Flag {
			fmt.Printf("Target Address: %s (待添加验证者)\n", proposal.Dst.Hex())
			fmt.Printf("Action: Add New Validator (Flag: true)\n")
		} else {
			fmt.Printf("Target Address: %s (待移除验证者)\n", proposal.Dst.Hex())
			fmt.Printf("Action: Remove Validator (Flag: false)\n")
		}
	} else if proposal.ProposalType.Int64() == 2 { // 配置更新提案
		fmt.Printf("Config ID: %s (%s)\n", proposal.Cid.String(), getConfigIDName(proposal.Cid.Int64()))
		fmt.Printf("New Value: %s\n", proposal.NewValue.String())
		fmt.Printf("Action: Update Configuration\n")
	}

	fmt.Printf("Proposal Type: %s (%s)\n", proposal.ProposalType.String(), getProposalTypeName(proposal.ProposalType.Int64()))
	fmt.Printf("Create Time: %s\n", timeToString(proposal.CreateTime.Int64()))

	if proposal.Details != "" {
		fmt.Printf("Details: %s\n", proposal.Details)
	}

	return nil
}

// 获取提案类型名称
func getProposalTypeName(proposalType int64) string {
	switch proposalType {
	case 1:
		return "Validator Management 验证者管理"
	case 2:
		return "Configuration Update 配置更新"
	default:
		return "Unknown Type 未知类型"
	}
}

// 获取配置项名称
func getConfigIDName(cid int64) string {
	switch cid {
	case 0:
		return "Proposal Lasting Period 提案持续时间"
	case 1:
		return "Punish Threshold 惩罚阈值"
	case 2:
		return "Remove Threshold 移除阈值"
	case 3:
		return "Decrease Rate 减少率"
	case 4:
		return "Withdraw Profit Period 提取收益周期"
	default:
		return "Unknown Config 未知配置"
	}
}

// QueryProposalsCmd creates a command to query all proposals
func QueryProposalsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "proposals",
		Short: "query all proposals",
		Run:   queryProposalsTx,
	}
	return cmd
}

func queryProposalsTx(cmd *cobra.Command, _ []string) {
	rpc := GetRPCEndpoint(cmd)

	// 验证输入参数
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	PrintInfo("Fetching all proposals...")

	if err := innerQueryProposals(rpc); err != nil {
		PrintError("Failed to query proposals", err)
		return
	}
}

func innerQueryProposals(rpc string) error {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return fmt.Errorf("failed to connect to RPC: %w", err)
	}
	defer client.Close()

	proposalContract, err := generated.NewProposal(common.HexToAddress(ProposalContractAddr), client)
	if err != nil {
		return fmt.Errorf("failed to instantiate proposal contract: %w", err)
	}

	// 获取当前区块号
	currentBlock, err := client.BlockNumber(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get current block number: %w", err)
	}

	// 设置查询范围（从创世块到当前块）
	opts := &bind.FilterOpts{
		Start: 0,
		End:   &currentBlock,
	}

	// 收集所有提案 ID
	proposalIDs := make(map[string]bool)

	// 查询验证者管理提案事件
	validatorProposalIter, err := proposalContract.FilterLogCreateProposal(opts, nil, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to filter validator proposals: %w", err)
	}
	defer validatorProposalIter.Close()

	for validatorProposalIter.Next() {
		event := validatorProposalIter.Event
		proposalID := common.BytesToHash(event.Id[:]).Hex()
		proposalIDs[proposalID] = true
	}

	// 查询配置更新提案事件
	configProposalIter, err := proposalContract.FilterLogCreateConfigProposal(opts, nil, nil)
	if err != nil {
		return fmt.Errorf("failed to filter config proposals: %w", err)
	}
	defer configProposalIter.Close()

	for configProposalIter.Next() {
		event := configProposalIter.Event
		proposalID := common.BytesToHash(event.Id[:]).Hex()
		proposalIDs[proposalID] = true
	}

	if len(proposalIDs) == 0 {
		fmt.Println("📋 No proposals found.")
		return nil
	}

	fmt.Printf("ℹ️  Found %d proposal(s):\n\n", len(proposalIDs))

	// 查询每个提案的详细信息
	count := 1
	for proposalID := range proposalIDs {
		fmt.Printf("--- Proposal %d ---\n", count)

		var proposalIdBytes [32]byte
		copy(proposalIdBytes[:], common.HexToHash(proposalID).Bytes())

		proposal, err := proposalContract.Proposals(nil, proposalIdBytes)
		if err != nil {
			fmt.Printf("❌ Failed to query proposal %s: %v\n", proposalID, err)
			continue
		}

		// 显示提案信息
		fmt.Printf("ID: %s\n", proposalID)
		fmt.Printf("Proposer: %s (验证者地址)\n", proposal.Proposer.Hex())

		// 根据提案类型显示不同信息
		if proposal.ProposalType.Int64() == 1 { // 验证者管理提案
			if proposal.Flag {
				fmt.Printf("Target: %s (待添加验证者)\n", proposal.Dst.Hex())
				fmt.Printf("Action: Add New Validator (添加验证者)\n")
			} else {
				fmt.Printf("Target: %s (待移除验证者)\n", proposal.Dst.Hex())
				fmt.Printf("Action: Remove Validator (移除验证者)\n")
			}
		} else if proposal.ProposalType.Int64() == 2 { // 配置更新提案
			fmt.Printf("Config ID: %s (%s)\n", proposal.Cid.String(), getConfigIDName(proposal.Cid.Int64()))
			fmt.Printf("New Value: %s\n", proposal.NewValue.String())
			fmt.Printf("Action: Update Configuration (更新配置)\n")
		}

		fmt.Printf("Type: %s (%s)\n", proposal.ProposalType.String(), getProposalTypeName(proposal.ProposalType.Int64()))
		fmt.Printf("Create Time: %s\n", timeToString(proposal.CreateTime.Int64()))

		// 查询提案投票结果和状态
		result, err := proposalContract.Results(nil, proposalIdBytes)
		if err != nil {
			fmt.Printf("⚠️  Status: Cannot query result (%v)\n", err)
		} else {
			statusIcon, statusText := getProposalStatus(result.Agree, result.Reject, result.ResultExist)
			fmt.Printf("Status: %s %s\n", statusIcon, statusText)
			if result.Agree > 0 || result.Reject > 0 {
				fmt.Printf("Votes: 👍 %d agree, 👎 %d reject\n", result.Agree, result.Reject)
			}
		}

		if proposal.Details != "" {
			fmt.Printf("Details: %s\n", proposal.Details)
		}

		fmt.Println()
		count++
	}

	return nil
}

// 时间转换辅助函数
func timeToString(timestamp int64) string {
	if timestamp == 0 {
		return "N/A"
	}
	return time.Unix(timestamp, 0).UTC().String()
}

// 获取提案状态的辅助函数
func getProposalStatus(agree uint16, reject uint16, resultExist bool) (string, string) {
	if !resultExist {
		if agree == 0 && reject == 0 {
			return "⏳", "Pending (No votes yet / 等待投票)"
		} else {
			return "⏳", "Pending (Voting in progress / 投票进行中)"
		}
	}
	
	// resultExist = true 意味着提案已经有了最终结果
	// 根据 Proposal.sol 逻辑：超过半数同意则通过，超过半数反对则失败
	if agree > reject {
		return "✅", "Passed (提案通过)"
	} else {
		return "❌", "Rejected (提案被拒绝)"
	}
}
