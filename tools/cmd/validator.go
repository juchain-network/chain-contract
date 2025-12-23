package cmd

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"juchain.org/chain/tools/contracts"
)

func ValidatorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "list all validators",
		Run:   listValidators,
	}
	return cmd
}

func listValidators(cmd *cobra.Command, _ []string) {
	rpc := GetRPCEndpoint(cmd) // Use config-aware function

	// Validate RPC URL
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	client, err := ethclient.Dial(rpc)
	if err != nil {
		PrintError("Failed to connect to RPC endpoint", err)
		return
	}
	defer client.Close()

	// Get validator information
	contractAddress := common.HexToAddress(ValidatorContractAddr)
	instance, err := contracts.NewValidators(contractAddress, client)
	if err != nil {
		PrintError("Failed to instantiate validator contract", err)
		return
	}

	PrintInfo("Fetching validator information...")
	vals, err := instance.GetTopValidators(&bind.CallOpts{})
	if err != nil {
		PrintError("Failed to get validators", err)
		return
	}

	if len(vals) == 0 {
		PrintInfo("No validators found")
		return
	}

	PrintInfo(fmt.Sprintf("Found %d validators:", len(vals)))
	for i, val := range vals {
		fmt.Printf("\n--- Validator %d ---\n", i+1)
		queryOneInfo(val.Hex(), instance)
	}
}

func ValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query",
		Short: "query a validator by address",
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
	rpc := GetRPCEndpoint(cmd) // Use config-aware function
	addr, _ := cmd.Flags().GetString("addr")

	// Validate input parameters
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateAddress(addr); err != nil {
		PrintValidationError(err)
		return
	}

	client, err := ethclient.Dial(rpc)
	if err != nil {
		PrintError("Failed to connect to RPC endpoint", err)
		return
	}
	defer client.Close()

	contractAddress := common.HexToAddress(ValidatorContractAddr)
	instance, err := contracts.NewValidators(contractAddress, client)
	if err != nil {
		PrintError("Failed to instantiate validator contract", err)
		return
	}

	PrintInfo(fmt.Sprintf("Querying validator information for: %s", addr))
	queryOneInfo(addr, instance)
}

func queryOneInfo(addr string, instance *contracts.Validators) {
	feeAddr, status, aacIncoming, totalJailedHB, lastWithdrawProfitsBlock, err := instance.GetValidatorInfo(&bind.CallOpts{}, common.HexToAddress(addr))
	if err != nil {
		PrintError(fmt.Sprintf("Failed to get validator info for %s", addr), err)
		return
	}

	fmt.Printf("Address: %s\n", addr)
	fmt.Printf("Fee Address: %s\n", feeAddr.Hex())
	// Print friendly status label instead of raw number
	fmt.Printf("Status: %s\n", formatValidatorStatus(uint64(status)))
	fmt.Printf("Accumulated Rewards: %s\n", aacIncoming.String())
	// totalJailedHB represents total penalized (forfeited) rewards
	fmt.Printf("Penalized Rewards: %s\n", totalJailedHB.String())
	fmt.Printf("Last Withdraw Block: %s\n", lastWithdrawProfitsBlock.String())
}

// formatValidatorStatus converts numeric status to a concise, friendly label
// 1 -> "Active ✅"; 2 -> "Inactive ❌"; others -> "Unknown"
func formatValidatorStatus(status uint64) string {
	switch status {
	case 1:
		return "Active ✅"
	case 2:
		return "Inactive ❌"
	default:
		return "Unknown"
	}
}

func WithdrawProfitsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "withdraw_profits",
		Short: "claim validator's reward",
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
	rpc := GetRPCEndpoint(cmd) // Use config-aware function
	addr, _ := cmd.Flags().GetString("addr")

	// Validate input parameters
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateAddress(addr); err != nil {
		PrintValidationError(err)
		return
	}

	PrintInfo(fmt.Sprintf("Creating withdraw profits transaction for validator: %s", addr))
	if err := innerValidatorClaim(addr, rpc); err != nil {
		PrintError("Failed to create withdraw transaction", err)
		return
	}
}

func innerValidatorClaim(addr string, rpc string) error {
	validatorAbi, err := abi.JSON(strings.NewReader(contracts.ValidatorsABI))
	if err != nil {
		return fmt.Errorf("failed to parse validator ABI: %w", err)
	}

	abiData, err := validatorAbi.Pack("withdrawProfits", common.HexToAddress(addr))
	if err != nil {
		return fmt.Errorf("failed to pack withdrawProfits data: %w", err)
	}

	err = CreateRawTx(common.HexToAddress(addr), common.HexToAddress(ValidatorContractAddr), nil, abiData, rpc, WithdrawProfitsFile)
	if err != nil {
		return fmt.Errorf("failed to create withdraw transaction: %w", err)
	}

	PrintSuccess("Withdraw profits transaction created successfully!")
	PrintInfo(fmt.Sprintf("Transaction file: %s", WithdrawProfitsFile))
	PrintWarning("Note: Withdrawal has minimum waiting period restrictions")
	return nil
}
