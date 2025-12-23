package cmd

import (
	"fmt"
	"math/big"
	"os"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"juchain.org/chain/tools/contracts"
)

// ValidateAddress validates Ethereum address format
func ValidateAddress(addr string) error {
	if addr == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if !common.IsHexAddress(addr) {
		return fmt.Errorf("invalid address format: %s", addr)
	}
	return nil
}

// ValidateAddresses validates multiple addresses
func ValidateAddresses(addrs ...string) error {
	for _, addr := range addrs {
		if err := ValidateAddress(addr); err != nil {
			return err
		}
	}
	return nil
}

// ValidateChainID validates chain ID
func ValidateChainID(chainID int64) error {
	if chainID <= 0 {
		return fmt.Errorf("invalid chain ID: %d, must be positive", chainID)
	}
	return nil
}

// ValidateRPCURL validates RPC URL format
func ValidateRPCURL(url string) error {
	if url == "" {
		return fmt.Errorf("RPC URL cannot be empty")
	}
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "ws://") && !strings.HasPrefix(url, "wss://") {
		return fmt.Errorf("invalid RPC URL format: %s", url)
	}
	return nil
}

// ValidateFile validates file existence
func ValidateFile(filepath string) error {
	if filepath == "" {
		return fmt.Errorf("file path cannot be empty")
	}
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filepath)
	}
	return nil
}

// ValidateOperation validates operation type
func ValidateOperation(operation string) error {
	if operation != "add" && operation != "remove" {
		return fmt.Errorf("invalid operation: %s, must be 'add' or 'remove'", operation)
	}
	return nil
}

// ValidateConfigID validates configuration ID
func ValidateConfigID(cid int64) error {
	if cid < 0 || cid > 9 {
		return fmt.Errorf("invalid config ID: %d, must be 0-9", cid)
	}
	return nil
}

// ValidateProposalID validates proposal ID format
func ValidateProposalID(proposalID string) error {
	if proposalID == "" {
		return fmt.Errorf("proposal ID cannot be empty")
	}

	// Allow 0x prefix
	cleanID := strings.TrimPrefix(proposalID, "0x")

	// Validate if it's a valid hexadecimal string
	if !regexp.MustCompile(`^[0-9a-fA-F]+$`).MatchString(cleanID) {
		return fmt.Errorf("invalid proposal ID format: %s", proposalID)
	}
	if len(cleanID) != 64 {
		return fmt.Errorf("proposal ID must be 64 characters long: %s", proposalID)
	}
	return nil
}

// GetConfigIDName gets the name of the configuration ID
func GetConfigIDName(cid int64) string {
	switch cid {
	case 0:
		return "proposalLastingPeriod"
	case 1:
		return "punishThreshold"
	case 2:
		return "removeThreshold"
	case 3:
		return "decreaseRate"
	case 4:
		return "withdrawProfitPeriod"
	case 5:
		return "blockReward"
	case 6:
		return "unbondingPeriod"
	case 7:
		return "validatorUnjailPeriod"
	default:
		return "unknown"
	}
}

// PrintValidationError prints validation error message
func PrintValidationError(err error) {
	fmt.Printf("❌ Validation Error: %v\n", err)
}

// PrintSuccess prints success message
func PrintSuccess(message string) {
	fmt.Printf("✅ %s\n", message)
}

// PrintInfo prints information message
func PrintInfo(message string) {
	fmt.Printf("ℹ️  %s\n", message)
}

// PrintWarning prints warning message
func PrintWarning(message string) {
	fmt.Printf("⚠️  %s\n", message)
}

// PrintError prints error message
func PrintError(message string, err error) {
	fmt.Printf("❌ %s: %v\n", message, err)
}

func EtherToWei(ether string) *big.Int {
	parts := strings.Split(ether, ".")
	if len(parts) == 1 {
		result := new(big.Int)
		result.SetString(parts[0], 10)
		return result.Mul(result, big.NewInt(1e18))
	}

	integerPart := new(big.Int)
	integerPart.SetString(parts[0], 10)
	integerPart.Mul(integerPart, big.NewInt(1e18))

	decimalPart := new(big.Int)
	decimalPart.SetString(parts[1], 10)

	zeros := 18 - len(parts[1])
	if zeros > 0 {
		decimalPart.Mul(decimalPart, new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(zeros)), nil))
	}

	return integerPart.Add(integerPart, decimalPart)
}

func WeiToEther(wei *big.Int) string {
	eth := new(big.Float).Quo(new(big.Float).SetInt(wei), big.NewFloat(1e18))
	return eth.Text('f', 18)
}

func GetContractInstance(rpc string) (*contracts.Validators, *contracts.Staking, *contracts.Proposal, error) {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		PrintError("Failed to connect to RPC endpoint", err)
		return nil, nil, nil, err
	}
	defer client.Close()

	// Get validator information
	validatorInstance, err := contracts.NewValidators(common.HexToAddress(ValidatorContractAddr), client)
	if err != nil {
		PrintError("Failed to instantiate validator contract", err)
		return nil, nil, nil, err
	}
	// Connect to staking contract
	stakingInstance, err := contracts.NewStaking(common.HexToAddress(StakingContractAddr), client)
	if err != nil {
		PrintError("Failed to instantiate staking contract", err)
		return nil, nil, nil, err
	}
	// Connect to proposal contract
	proposalInstance, err := contracts.NewProposal(common.HexToAddress(ProposalContractAddr), client)
	if err != nil {
		PrintError("Failed to instantiate proposal contract", err)
		return nil, nil, nil, err
	}
	return validatorInstance, stakingInstance, proposalInstance, nil
}
