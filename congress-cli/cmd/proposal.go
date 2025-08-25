package cmd

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"juchain.org/chain/congress-cli/contracts/generated"
)

const (
	validatorAddr = ValidatorContractAddr
	punishAddr    = PunishContractAddr
	proposalAddr  = ProposalContractAddr
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

	err = CreateRawTx(common.HexToAddress(proposer), common.HexToAddress(proposalAddr), nil, abiData, rpc, CreateProposalFile)
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

	err = CreateRawTx(common.HexToAddress(proposer), common.HexToAddress(proposalAddr), nil, abiData, rpc, CreateConfigProposalFile)
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

	err = CreateRawTx(common.HexToAddress(signer), common.HexToAddress(proposalAddr), nil, abiData, rpc, VoteProposalFile)
	if err != nil {
		return fmt.Errorf("failed to create vote transaction: %w", err)
	}

	PrintSuccess("Vote transaction created successfully!")
	PrintInfo(fmt.Sprintf("Transaction file: %s", VoteProposalFile))
	return nil
}
