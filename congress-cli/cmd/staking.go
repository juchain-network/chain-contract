package cmd

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
)

// Staking contract ABI (simplified for CLI usage)
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
		"inputs": [{"internalType": "uint256", "name": "limit", "type": "uint256"}],
		"name": "getTopValidators",
		"outputs": [{"internalType": "address[]", "name": "", "type": "address[]"}],
		"stateMutability": "view",
		"type": "function"
	}
]`

// Staking transaction file names
const (
	RegisterValidatorFile = "registerValidator.json"
	DelegateFile          = "delegate.json"
	UndelegateFile        = "undelegate.json"
	ClaimRewardsFile      = "claimRewards.json"
)

// StakingCmd creates the main staking command
func StakingCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "staking",
		Short: "Staking operations for JPoSA consensus",
		Long:  "Manage validator staking, delegation, and rewards in JuChain PoSA consensus",
	}

	cmd.AddCommand(
		registerValidatorCmd(),
		delegateCmd(),
		undelegateCmd(),
		claimRewardsCmd(),
		queryValidatorCmd(),
		queryDelegationCmd(),
		listTopValidatorsCmd(),
	)

	return cmd
}

// registerValidatorCmd creates command for validator registration
func registerValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "register-validator",
		Short: "Create validator registration transaction",
		Long:  "Create a transaction to register as a validator by staking JU tokens and setting commission rate",
		Run:   createRegisterValidatorTx,
	}

	cmd.Flags().String("proposer", "", "Proposer address (required)")
	cmd.Flags().String("stake-amount", "", "Amount of JU to stake (required, minimum 10000)")
	cmd.Flags().String("commission-rate", "", "Commission rate in basis points (0-10000, e.g., 500 = 5%)")

	cmd.MarkFlagRequired("proposer")
	cmd.MarkFlagRequired("stake-amount")
	cmd.MarkFlagRequired("commission-rate")

	return cmd
}

// delegateCmd creates command for delegation
func delegateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delegate",
		Short: "Create delegation transaction",
		Long:  "Create a transaction to delegate JU tokens to a validator",
		Run:   createDelegateTx,
	}

	cmd.Flags().String("delegator", "", "Delegator address (required)")
	cmd.Flags().String("validator", "", "Validator address to delegate to (required)")
	cmd.Flags().String("amount", "", "Amount of JU to delegate (required, minimum 1)")

	cmd.MarkFlagRequired("delegator")
	cmd.MarkFlagRequired("validator")
	cmd.MarkFlagRequired("amount")

	return cmd
}

// undelegateCmd creates command for undelegation
func undelegateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "undelegate",
		Short: "Create undelegation transaction",
		Long:  "Create a transaction to start the 7-day unbonding process for delegated tokens",
		Run:   createUndelegateTx,
	}

	cmd.Flags().String("delegator", "", "Delegator address (required)")
	cmd.Flags().String("validator", "", "Validator address to undelegate from (required)")
	cmd.Flags().String("amount", "", "Amount of JU to undelegate (required)")

	cmd.MarkFlagRequired("delegator")
	cmd.MarkFlagRequired("validator")
	cmd.MarkFlagRequired("amount")

	return cmd
}

// claimRewardsCmd creates command for claiming rewards
func claimRewardsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "claim-rewards",
		Short: "Create claim rewards transaction",
		Long:  "Create a transaction to claim accumulated staking rewards",
		Run:   createClaimRewardsTx,
	}

	cmd.Flags().String("claimer", "", "Claimer address (required)")
	cmd.Flags().String("validator", "", "Validator address to claim rewards from (required)")

	cmd.MarkFlagRequired("claimer")
	cmd.MarkFlagRequired("validator")

	return cmd
}

// queryValidatorCmd creates command for querying validator info
func queryValidatorCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-validator",
		Short: "Query validator information",
		Long:  "Get detailed information about a validator including stake and status",
		Run:   queryValidatorInfo,
	}

	cmd.Flags().String("address", "", "Validator address to query (required)")
	cmd.MarkFlagRequired("address")

	return cmd
}

// queryDelegationCmd creates command for querying delegation info
func queryDelegationCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "query-delegation",
		Short: "Query delegation information",
		Long:  "Get delegation details including amount and pending rewards",
		Run:   queryDelegationInfo,
	}

	cmd.Flags().String("delegator", "", "Delegator address (required)")
	cmd.Flags().String("validator", "", "Validator address (required)")

	cmd.MarkFlagRequired("delegator")
	cmd.MarkFlagRequired("validator")

	return cmd
}

// listTopValidatorsCmd creates command for listing top validators
func listTopValidatorsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-top-validators",
		Short: "List top validators by stake",
		Long:  "Get the list of top validators ranked by total stake",
		Run:   queryTopValidators,
	}

	cmd.Flags().Int("limit", 21, "Maximum number of validators to return")

	return cmd
}

// Implementation functions

func createRegisterValidatorTx(cmd *cobra.Command, args []string) {
	rpc := GetRPCEndpoint(cmd)
	proposer, _ := cmd.Flags().GetString("proposer")
	stakeAmountStr, _ := cmd.Flags().GetString("stake-amount")
	commissionRateStr, _ := cmd.Flags().GetString("commission-rate")

	// Validate inputs
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateAddress(proposer); err != nil {
		PrintValidationError(err)
		return
	}

	// Parse amounts
	stakeAmount, ok := new(big.Int).SetString(stakeAmountStr, 10)
	if !ok {
		PrintValidationError(fmt.Errorf("invalid stake amount: %s", stakeAmountStr))
		return
	}

	// Convert to wei (multiply by 10^18)
	weiAmount := new(big.Int).Mul(stakeAmount, big.NewInt(1e18))

	commissionRate, ok := new(big.Int).SetString(commissionRateStr, 10)
	if !ok {
		PrintValidationError(fmt.Errorf("invalid commission rate: %s", commissionRateStr))
		return
	}

	// Validate commission rate (0-10000)
	if commissionRate.Cmp(big.NewInt(10000)) > 0 {
		PrintValidationError(fmt.Errorf("commission rate must be between 0 and 10000 (100%%)"))
		return
	}

	PrintInfo("Creating validator registration transaction")

	if err := innerCreateRegisterValidatorTx(proposer, weiAmount, commissionRate, rpc); err != nil {
		PrintError("Failed to create register validator transaction", err)
		return
	}
}

func innerCreateRegisterValidatorTx(proposer string, stakeAmount, commissionRate *big.Int, rpc string) error {
	stakingAbi, err := abi.JSON(strings.NewReader(stakingABI))
	if err != nil {
		return fmt.Errorf("failed to parse staking ABI: %w", err)
	}

	abiData, err := stakingAbi.Pack("registerValidator", commissionRate)
	if err != nil {
		return fmt.Errorf("failed to pack registerValidator data: %w", err)
	}

	err = CreateRawTx(common.HexToAddress(proposer), common.HexToAddress(StakingContractAddr), stakeAmount, abiData, rpc, RegisterValidatorFile)
	if err != nil {
		return fmt.Errorf("failed to create raw transaction: %w", err)
	}

	PrintSuccess("Validator registration transaction created successfully!")
	PrintInfo(fmt.Sprintf("Transaction file: %s", RegisterValidatorFile))
	PrintInfo(fmt.Sprintf("Stake amount: %s wei", stakeAmount.String()))
	PrintInfo(fmt.Sprintf("Commission rate: %s basis points", commissionRate.String()))
	return nil
}

func createDelegateTx(cmd *cobra.Command, args []string) {
	rpc := GetRPCEndpoint(cmd)
	delegator, _ := cmd.Flags().GetString("delegator")
	validatorStr, _ := cmd.Flags().GetString("validator")
	amountStr, _ := cmd.Flags().GetString("amount")

	// Validate inputs
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateAddresses(delegator, validatorStr); err != nil {
		PrintValidationError(err)
		return
	}

	// Parse amount
	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		PrintValidationError(fmt.Errorf("invalid amount: %s", amountStr))
		return
	}

	// Convert to wei
	weiAmount := new(big.Int).Mul(amount, big.NewInt(1e18))

	PrintInfo("Creating delegation transaction")

	if err := innerCreateDelegateTx(delegator, validatorStr, weiAmount, rpc); err != nil {
		PrintError("Failed to create delegate transaction", err)
		return
	}
}

func innerCreateDelegateTx(delegator, validator string, amount *big.Int, rpc string) error {
	stakingAbi, err := abi.JSON(strings.NewReader(stakingABI))
	if err != nil {
		return fmt.Errorf("failed to parse staking ABI: %w", err)
	}

	abiData, err := stakingAbi.Pack("delegate", common.HexToAddress(validator))
	if err != nil {
		return fmt.Errorf("failed to pack delegate data: %w", err)
	}

	err = CreateRawTx(common.HexToAddress(delegator), common.HexToAddress(StakingContractAddr), amount, abiData, rpc, DelegateFile)
	if err != nil {
		return fmt.Errorf("failed to create raw transaction: %w", err)
	}

	PrintSuccess("Delegation transaction created successfully!")
	PrintInfo(fmt.Sprintf("Transaction file: %s", DelegateFile))
	PrintInfo(fmt.Sprintf("Amount: %s wei", amount.String()))
	PrintInfo(fmt.Sprintf("Validator: %s", validator))
	return nil
}

func createUndelegateTx(cmd *cobra.Command, args []string) {
	rpc := GetRPCEndpoint(cmd)
	delegator, _ := cmd.Flags().GetString("delegator")
	validatorStr, _ := cmd.Flags().GetString("validator")
	amountStr, _ := cmd.Flags().GetString("amount")

	// Validate inputs
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateAddresses(delegator, validatorStr); err != nil {
		PrintValidationError(err)
		return
	}

	// Parse amount
	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		PrintValidationError(fmt.Errorf("invalid amount: %s", amountStr))
		return
	}

	// Convert to wei
	weiAmount := new(big.Int).Mul(amount, big.NewInt(1e18))

	PrintInfo("Creating undelegation transaction")

	if err := innerCreateUndelegateTx(delegator, validatorStr, weiAmount, rpc); err != nil {
		PrintError("Failed to create undelegate transaction", err)
		return
	}
}

func innerCreateUndelegateTx(delegator, validator string, amount *big.Int, rpc string) error {
	stakingAbi, err := abi.JSON(strings.NewReader(stakingABI))
	if err != nil {
		return fmt.Errorf("failed to parse staking ABI: %w", err)
	}

	abiData, err := stakingAbi.Pack("undelegate", common.HexToAddress(validator), amount)
	if err != nil {
		return fmt.Errorf("failed to pack undelegate data: %w", err)
	}

	err = CreateRawTx(common.HexToAddress(delegator), common.HexToAddress(StakingContractAddr), big.NewInt(0), abiData, rpc, UndelegateFile)
	if err != nil {
		return fmt.Errorf("failed to create raw transaction: %w", err)
	}

	PrintSuccess("Undelegation transaction created successfully!")
	PrintInfo(fmt.Sprintf("Transaction file: %s", UndelegateFile))
	PrintInfo(fmt.Sprintf("Amount: %s wei", amount.String()))
	PrintInfo(fmt.Sprintf("Validator: %s", validator))
	PrintInfo("Unbonding period: 7 days")
	return nil
}

func createClaimRewardsTx(cmd *cobra.Command, args []string) {
	rpc := GetRPCEndpoint(cmd)
	claimer, _ := cmd.Flags().GetString("claimer")
	validatorStr, _ := cmd.Flags().GetString("validator")

	// Validate inputs
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateAddresses(claimer, validatorStr); err != nil {
		PrintValidationError(err)
		return
	}

	PrintInfo("Creating claim rewards transaction")

	if err := innerCreateClaimRewardsTx(claimer, validatorStr, rpc); err != nil {
		PrintError("Failed to create claim rewards transaction", err)
		return
	}
}

func innerCreateClaimRewardsTx(claimer, validator string, rpc string) error {
	stakingAbi, err := abi.JSON(strings.NewReader(stakingABI))
	if err != nil {
		return fmt.Errorf("failed to parse staking ABI: %w", err)
	}

	abiData, err := stakingAbi.Pack("claimRewards", common.HexToAddress(validator))
	if err != nil {
		return fmt.Errorf("failed to pack claimRewards data: %w", err)
	}

	err = CreateRawTx(common.HexToAddress(claimer), common.HexToAddress(StakingContractAddr), big.NewInt(0), abiData, rpc, ClaimRewardsFile)
	if err != nil {
		return fmt.Errorf("failed to create raw transaction: %w", err)
	}

	PrintSuccess("Claim rewards transaction created successfully!")
	PrintInfo(fmt.Sprintf("Transaction file: %s", ClaimRewardsFile))
	PrintInfo(fmt.Sprintf("Validator: %s", validator))
	return nil
}

func queryValidatorInfo(cmd *cobra.Command, args []string) {
	rpc := GetRPCEndpoint(cmd)
	addressStr, _ := cmd.Flags().GetString("address")

	// Validate inputs
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateAddress(addressStr); err != nil {
		PrintValidationError(err)
		return
	}

	validator := common.HexToAddress(addressStr)

	// Connect to blockchain
	client, err := ethclient.Dial(rpc)
	if err != nil {
		PrintError("Failed to connect to RPC", err)
		return
	}
	defer client.Close()

	// Create call
	parsedABI, err := abi.JSON(strings.NewReader(stakingABI))
	if err != nil {
		PrintError("Failed to parse ABI", err)
		return
	}

	data, err := parsedABI.Pack("getValidatorInfo", validator)
	if err != nil {
		PrintError("Failed to pack call data", err)
		return
	}

	// Make call
	msg := ethereum.CallMsg{
		To:   &common.Address{}, // Will be set below
		Data: data,
	}
	stakingAddr := common.HexToAddress(StakingContractAddr)
	msg.To = &stakingAddr

	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		PrintError("Failed to call contract", err)
		return
	}

	// Unpack result
	unpacked, err := parsedABI.Unpack("getValidatorInfo", result)
	if err != nil {
		PrintError("Failed to unpack result", err)
		return
	}

	if len(unpacked) != 5 {
		PrintError("Unexpected result length", fmt.Errorf("got %d, expected 5", len(unpacked)))
		return
	}

	// Convert wei to JU
	selfStake := unpacked[0].(*big.Int)
	totalDelegated := unpacked[1].(*big.Int)
	commissionRate := unpacked[2].(*big.Int)
	isJailed := unpacked[3].(bool)
	jailUntilBlock := unpacked[4].(*big.Int)

	selfStakeJU := new(big.Int).Div(selfStake, big.NewInt(1e18))
	totalDelegatedJU := new(big.Int).Div(totalDelegated, big.NewInt(1e18))
	totalStakeJU := new(big.Int).Add(selfStakeJU, totalDelegatedJU)

	PrintSuccess("Validator Information")
	fmt.Printf("Address: %s\n", addressStr)
	fmt.Printf("Self Stake: %s JU\n", selfStakeJU.String())
	fmt.Printf("Total Delegated: %s JU\n", totalDelegatedJU.String())
	fmt.Printf("Total Stake: %s JU\n", totalStakeJU.String())
	fmt.Printf("Commission Rate: %s basis points\n", commissionRate.String())
	fmt.Printf("Is Jailed: %t\n", isJailed)
	fmt.Printf("Jail Until Block: %s\n", jailUntilBlock.String())
}

func queryDelegationInfo(cmd *cobra.Command, args []string) {
	rpc := GetRPCEndpoint(cmd)
	delegatorStr, _ := cmd.Flags().GetString("delegator")
	validatorStr, _ := cmd.Flags().GetString("validator")

	// Validate inputs
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateAddresses(delegatorStr, validatorStr); err != nil {
		PrintValidationError(err)
		return
	}

	delegator := common.HexToAddress(delegatorStr)
	validator := common.HexToAddress(validatorStr)

	// Connect to blockchain
	client, err := ethclient.Dial(rpc)
	if err != nil {
		PrintError("Failed to connect to RPC", err)
		return
	}
	defer client.Close()

	// Create call
	parsedABI, err := abi.JSON(strings.NewReader(stakingABI))
	if err != nil {
		PrintError("Failed to parse ABI", err)
		return
	}

	data, err := parsedABI.Pack("getDelegationInfo", delegator, validator)
	if err != nil {
		PrintError("Failed to pack call data", err)
		return
	}

	// Make call
	msg := ethereum.CallMsg{
		To:   &common.Address{},
		Data: data,
	}
	stakingAddr := common.HexToAddress(StakingContractAddr)
	msg.To = &stakingAddr

	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		PrintError("Failed to call contract", err)
		return
	}

	// Unpack result
	unpacked, err := parsedABI.Unpack("getDelegationInfo", result)
	if err != nil {
		PrintError("Failed to unpack result", err)
		return
	}

	if len(unpacked) != 4 {
		PrintError("Unexpected result length", fmt.Errorf("got %d, expected 4", len(unpacked)))
		return
	}

	// Convert wei to JU
	amount := unpacked[0].(*big.Int)
	pendingRewards := unpacked[1].(*big.Int)
	unbondingAmount := unpacked[2].(*big.Int)
	unbondingBlock := unpacked[3].(*big.Int)

	amountJU := new(big.Int).Div(amount, big.NewInt(1e18))
	pendingRewardsJU := new(big.Int).Div(pendingRewards, big.NewInt(1e18))
	unbondingAmountJU := new(big.Int).Div(unbondingAmount, big.NewInt(1e18))

	PrintSuccess("Delegation Information")
	fmt.Printf("Delegator: %s\n", delegatorStr)
	fmt.Printf("Validator: %s\n", validatorStr)
	fmt.Printf("Delegated Amount: %s JU\n", amountJU.String())
	fmt.Printf("Pending Rewards: %s JU\n", pendingRewardsJU.String())
	fmt.Printf("Unbonding Amount: %s JU\n", unbondingAmountJU.String())
	fmt.Printf("Unbonding Block: %s\n", unbondingBlock.String())
}

func queryTopValidators(cmd *cobra.Command, args []string) {
	rpc := GetRPCEndpoint(cmd)
	limit, _ := cmd.Flags().GetInt("limit")

	// Validate inputs
	if err := ValidateRPCURL(rpc); err != nil {
		PrintValidationError(err)
		return
	}

	// Connect to blockchain
	client, err := ethclient.Dial(rpc)
	if err != nil {
		PrintError("Failed to connect to RPC", err)
		return
	}
	defer client.Close()

	// Create call
	parsedABI, err := abi.JSON(strings.NewReader(stakingABI))
	if err != nil {
		PrintError("Failed to parse ABI", err)
		return
	}

	data, err := parsedABI.Pack("getTopValidators", big.NewInt(int64(limit)))
	if err != nil {
		PrintError("Failed to pack call data", err)
		return
	}

	// Make call
	msg := ethereum.CallMsg{
		To:   &common.Address{},
		Data: data,
	}
	stakingAddr := common.HexToAddress(StakingContractAddr)
	msg.To = &stakingAddr

	result, err := client.CallContract(context.Background(), msg, nil)
	if err != nil {
		PrintError("Failed to call contract", err)
		return
	}

	// Unpack result
	unpacked, err := parsedABI.Unpack("getTopValidators", result)
	if err != nil {
		PrintError("Failed to unpack result", err)
		return
	}

	if len(unpacked) != 1 {
		PrintError("Unexpected result length", fmt.Errorf("got %d, expected 1", len(unpacked)))
		return
	}

	validators := unpacked[0].([]common.Address)

	PrintSuccess("Top Validators")
	fmt.Printf("Count: %d\n", len(validators))
	for i, validator := range validators {
		fmt.Printf("%d. %s\n", i+1, validator.Hex())
	}
}
