package cmd

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"
	"juchain.org/chain/contract/contracts/generated"
)

const (
	validatorAddr = "0x000000000000000000000000000000000000f000"
	punishAddr    = "0x000000000000000000000000000000000000f001"
	proposalAddr  = "0x000000000000000000000000000000000000f002"
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
	rpc, _ := cmd.Flags().GetString("rpc_laddr")
	proposer, _ := cmd.Flags().GetString("proposer")
	target, _ := cmd.Flags().GetString("target")
	operation, _ := cmd.Flags().GetString("operation")
	flag := true
	if operation == "add" {
		flag = true
		fmt.Printf("create add new miner tx\n")
	} else if operation == "remove" {
		flag = false
		fmt.Printf("create remove miner tx\n")
	} else {
		fmt.Printf("Invalid operation %s\n", operation)
		return
	}

	innerCreateProposal(proposer, target, flag, rpc)
}

func innerCreateProposal(proposer, target string, flag bool, rpc string) {
	proposalAbi, err := abi.JSON(strings.NewReader(generated.ProposalABI))
	if err != nil {
		fmt.Println("JSON NewReader Err:", err)
		return
	}

	abiData, err := proposalAbi.Pack("createProposal", common.HexToAddress(target), flag, "")
	if err != nil {
		fmt.Println("proposalAbi.Pack createProposal Err:", err)
		return
	}
	err = CreateRawTx(common.HexToAddress(proposer), common.HexToAddress(proposalAddr), nil, abiData, rpc, "createProposal.json")
	if err != nil {
		fmt.Println("create tx Err:", err)
		return
	}
	fmt.Println("crete tx success!")
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
	rpc, _ := cmd.Flags().GetString("rpc_laddr")
	proposer, _ := cmd.Flags().GetString("proposer")
	cid, _ := cmd.Flags().GetInt64("cid")
	cvalue, _ := cmd.Flags().GetInt64("value")
	innerCreateConfigProposal(proposer, cid, cvalue, rpc)
}

func innerCreateConfigProposal(proposer string, cid, cvalue int64, rpc string) {
	proposalAbi, err := abi.JSON(strings.NewReader(generated.ProposalABI))
	if err != nil {
		fmt.Println("JSON NewReader Err:", err)
		return
	}

	abiData, err := proposalAbi.Pack("createUpdateConfigProposal", big.NewInt(cid), big.NewInt(cvalue))
	if err != nil {
		fmt.Println("proposalAbi.Pack createUpdateConfigProposal Err:", err)
		return
	}
	err = CreateRawTx(common.HexToAddress(proposer), common.HexToAddress(proposalAddr), nil, abiData, rpc, "createUpdateConfigProposal.json")
	if err != nil {
		fmt.Println("create tx Err:", err)
		return
	}
	fmt.Println("crete tx success!")
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
	cmd.Flags().StringP("signer", "s", "", "singer addr (must be valid validator)")
	_ = cmd.MarkFlagRequired("signer")
	cmd.Flags().StringP("proposalId", "i", "", "proposal id")
	_ = cmd.MarkFlagRequired("proposalId")
	cmd.Flags().BoolP("approve", "a", false, "approve this proposal or not")
	_ = cmd.MarkFlagRequired("approve")
}

func voteProposalTx(cmd *cobra.Command, _ []string) {
	rpc, _ := cmd.Flags().GetString("rpc_laddr")

	signer, _ := cmd.Flags().GetString("signer")
	proposalId, _ := cmd.Flags().GetString("proposalId")
	approve, _ := cmd.Flags().GetBool("approve")

	innerVoteProposal(signer, proposalId, approve, rpc)
}

func innerVoteProposal(signer, proposalId string, flag bool, rpc string) {

	proposalAbi, err := abi.JSON(strings.NewReader(generated.ProposalABI))
	if err != nil {
		fmt.Println("JSON NewReader Err:", err)
		return
	}

	var proposalIdBytes [32]byte
	copy(proposalIdBytes[:], common.HexToHash(proposalId).Bytes())

	abiData, err := proposalAbi.Pack("voteProposal", proposalIdBytes, flag)
	if err != nil {
		fmt.Println("proposalAbi.Pack createProposal Err:", err)
		return
	}
	err = CreateRawTx(common.HexToAddress(signer), common.HexToAddress(proposalAddr), nil, abiData, rpc, "voteProposal.json")
	if err != nil {
		fmt.Printf("create tx error:%v", err)
		return
	}
	fmt.Println("create tx success!")
}
