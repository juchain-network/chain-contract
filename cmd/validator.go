package cmd

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"juchain.org/chain/contract/contracts/generated"
)

func ValidatorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "miners",
		Short: "list all valid miners",
		Run:   listValidators,
	}
	return cmd
}

func listValidators(cmd *cobra.Command, _ []string) {
	rpc, _ := cmd.Flags().GetString("rpc_laddr")

	client, err := ethclient.Dial(rpc)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v", err)
		return
	}
	// Get validator information
	contractAddress := common.HexToAddress(validatorAddr)
	instance, err := generated.NewValidators(contractAddress, client)
	if err != nil {
		fmt.Printf("Failed to instantiate contract: %v", err)
		return
	}
	vals, err := instance.GetTopValidators(&bind.CallOpts{})
	if err != nil {
		fmt.Printf("Failed to call GetValidatorInfo: %v", err)
	} else {
		for _, val := range vals {
			queryOneInfo(val.Hex(), instance)
			// fmt.Printf("miner%v %v\n", i, val)
		}
	}
}

func ValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "miner",
		Short: "query one miner info",
		Run:   queryValidator,
	}
	queryValidatorFlags(cmd)
	return cmd
}

func queryValidatorFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("addr", "a", "", "validator address to query")
	_ = cmd.MarkFlagRequired("addr")
}

func queryValidator(cmd *cobra.Command, _ []string) {
	rpc, _ := cmd.Flags().GetString("rpc_laddr")
	addr, _ := cmd.Flags().GetString("addr")

	client, err := ethclient.Dial(rpc)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v", err)
		return
	}

	contractAddress := common.HexToAddress(validatorAddr)
	instance, err := generated.NewValidators(contractAddress, client)
	if err != nil {
		fmt.Printf("Failed to instantiate contract: %v", err)
	}

	queryOneInfo(addr, instance)
}

func queryOneInfo(addr string, instance *generated.Validators) {
	feeAddr, status, aacIncoming, totalJailedHB, lastWithdrawProfitsBlock, err := instance.GetValidatorInfo(&bind.CallOpts{}, common.HexToAddress(addr))
	if err != nil {
		fmt.Printf("Failed to call GetValidatorInfo address %v: %v", addr, err)
	} else {
		fmt.Println("Validator ", addr, "Fee Addr", feeAddr, "Status", status, "Accumulated Rewards", aacIncoming, "Confiscated Rewards", totalJailedHB, "Last Withdraw Block", lastWithdrawProfitsBlock)
	}
}

func WithdrawProfitsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw_profits",
		Short: "claim miner's reward",
		Run:   validatorClaim,
	}
	validatorClaimFlags(cmd)
	return cmd
}

func validatorClaimFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("addr", "a", "", "validator address to claim")
	_ = cmd.MarkFlagRequired("addr")
}

func validatorClaim(cmd *cobra.Command, _ []string) {
	rpc, _ := cmd.Flags().GetString("rpc_laddr")
	addr, _ := cmd.Flags().GetString("addr")
	innerValidatorClaim(addr, rpc)
}

func innerValidatorClaim(addr string, rpc string) {
	validatorAbi, err := abi.JSON(strings.NewReader(generated.ValidatorsABI))
	if err != nil {
		fmt.Println("JSON NewReader Err:", err)
		return
	}

	abiData, err := validatorAbi.Pack("withdrawProfits", common.HexToAddress(addr))
	if err != nil {
		fmt.Println("validatorAbi.Pack withdrawProfits Err:", err)
		return
	}
	err = CreateRawTx(common.HexToAddress(addr), common.HexToAddress(validatorAddr), nil, abiData, rpc, "withdrawProfits.json")
	if err != nil {
		fmt.Println("create tx Err:", err)
		return
	}
	fmt.Println("crete tx success!")

}
