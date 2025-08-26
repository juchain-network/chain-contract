package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/ethereum/go-ethereum/common"
)

// ValidateAddress 验证以太坊地址格式
func ValidateAddress(addr string) error {
	if addr == "" {
		return fmt.Errorf("address cannot be empty")
	}
	if !common.IsHexAddress(addr) {
		return fmt.Errorf("invalid address format: %s", addr)
	}
	return nil
}

// ValidateAddresses 批量验证地址
func ValidateAddresses(addrs ...string) error {
	for _, addr := range addrs {
		if err := ValidateAddress(addr); err != nil {
			return err
		}
	}
	return nil
}

// ValidateChainID 验证链ID
func ValidateChainID(chainID int64) error {
	if chainID <= 0 {
		return fmt.Errorf("invalid chain ID: %d, must be positive", chainID)
	}
	return nil
}

// ValidateRPCURL 验证RPC URL格式
func ValidateRPCURL(url string) error {
	if url == "" {
		return fmt.Errorf("RPC URL cannot be empty")
	}
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "ws://") && !strings.HasPrefix(url, "wss://") {
		return fmt.Errorf("invalid RPC URL format: %s", url)
	}
	return nil
}

// ValidateFile 验证文件存在性
func ValidateFile(filepath string) error {
	if filepath == "" {
		return fmt.Errorf("file path cannot be empty")
	}
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", filepath)
	}
	return nil
}

// ValidateOperation 验证操作类型
func ValidateOperation(operation string) error {
	if operation != "add" && operation != "remove" {
		return fmt.Errorf("invalid operation: %s, must be 'add' or 'remove'", operation)
	}
	return nil
}

// ValidateConfigID 验证配置ID
func ValidateConfigID(cid int64) error {
	if cid < 0 || cid > 4 {
		return fmt.Errorf("invalid config ID: %d, must be 0-4", cid)
	}
	return nil
}

// ValidateProposalID 验证提案ID格式
func ValidateProposalID(proposalID string) error {
	if proposalID == "" {
		return fmt.Errorf("proposal ID cannot be empty")
	}

	// 允许 0x 前缀
	cleanID := strings.TrimPrefix(proposalID, "0x")

	// 验证是否为有效的十六进制字符串
	if !regexp.MustCompile(`^[0-9a-fA-F]+$`).MatchString(cleanID) {
		return fmt.Errorf("invalid proposal ID format: %s", proposalID)
	}
	if len(cleanID) != 64 {
		return fmt.Errorf("proposal ID must be 64 characters long: %s", proposalID)
	}
	return nil
}

// GetConfigIDName 获取配置ID的名称
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
	default:
		return "unknown"
	}
}

// PrintValidationError 打印验证错误信息
func PrintValidationError(err error) {
	fmt.Printf("❌ Validation Error: %v\n", err)
}

// PrintSuccess 打印成功信息
func PrintSuccess(message string) {
	fmt.Printf("✅ %s\n", message)
}

// PrintInfo 打印信息
func PrintInfo(message string) {
	fmt.Printf("ℹ️  %s\n", message)
}

// PrintWarning 打印警告信息
func PrintWarning(message string) {
	fmt.Printf("⚠️  %s\n", message)
}

// PrintError 打印错误信息
func PrintError(message string, err error) {
	fmt.Printf("❌ %s: %v\n", message, err)
}
