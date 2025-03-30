package cmd

import (
	"fmt"
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
	// 获取validator信息
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
			queryOneInfo(val.Hex(), instance, client)
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

	queryOneInfo(addr, instance, client)
}

func queryOneInfo(addr string, instance *generated.Validators, client *ethclient.Client) {
	feeAddr, status, aacIncoming, totalJailedHB, lastWithdrawProfitsBlock, err := instance.GetValidatorInfo(&bind.CallOpts{}, common.HexToAddress(addr))
	if err != nil {
		fmt.Printf("Failed to call GetValidatorInfo address %v: %v", addr, err)
	} else {
		fmt.Println("矿工 ", addr, "奖励地址", feeAddr, "活动状态", status, "累计奖励", aacIncoming, "罚没奖励", totalJailedHB, "上次提取奖励区块", lastWithdrawProfitsBlock)
	}
}
