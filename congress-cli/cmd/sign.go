package cmd

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

func SignRawTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sign",
		Short: "sign raw tx from file",
		Run:   signRawTx,
	}
	signRawTxCmdFlags(cmd)
	return cmd
}

func signRawTxCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("file", "f", "", "raw tx file")
	_ = cmd.MarkFlagRequired("file")
	cmd.Flags().StringP("key", "k", "", "singer wallet file")
	_ = cmd.MarkFlagRequired("key")
	cmd.Flags().StringP("password", "p", "", "singer  password file ")
	_ = cmd.MarkFlagRequired("password")
}

func signRawTx(cmd *cobra.Command, _ []string) {
	chainId := GetChainID(cmd) // Use config-aware function
	file, _ := cmd.Flags().GetString("file")
	key, _ := cmd.Flags().GetString("key")
	passwordFile, _ := cmd.Flags().GetString("password")

	// 验证输入参数
	if err := ValidateChainID(chainId); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateFile(file); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateFile(key); err != nil {
		PrintValidationError(err)
		return
	}

	if err := ValidateFile(passwordFile); err != nil {
		PrintValidationError(err)
		return
	}

	PrintInfo(fmt.Sprintf("Signing transaction from file: %s", file))

	password, err := fethchKeyFromFile(passwordFile)
	if err != nil {
		PrintError("Failed to read password file", err)
		return
	}

	privateKey, err := ReadKeystoreFile(key, password)
	if err != nil {
		PrintError("Failed to decrypt keystore file", err)
		return
	}

	if err := innerSignRawTx(chainId, file, privateKey); err != nil {
		PrintError("Failed to sign transaction", err)
		return
	}
}

func fethchKeyFromFile(path string) (string, error) {
	text, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	lines := strings.Split(string(text), "\n")
	// Sanitise DOS line endings.
	for i := range lines {
		lines[i] = strings.TrimRight(lines[i], "\r")
	}
	return lines[0], nil
}

func innerSignRawTx(chainId int64, file string, privateKey *ecdsa.PrivateKey) error {
	outputFile := addSuffixToFilename(file, "_signed")

	err := SignRawTx(file, privateKey, big.NewInt(chainId), outputFile)
	if err != nil {
		return fmt.Errorf("failed to sign transaction: %w", err)
	}

	PrintSuccess("Transaction signed successfully!")
	PrintInfo(fmt.Sprintf("Signed transaction saved to: %s", outputFile))
	return nil
}

func addSuffixToFilename(filename, suffix string) string {
	ext := filepath.Ext(filename)
	base := strings.TrimSuffix(filename, ext)
	return base + suffix + ext
}
