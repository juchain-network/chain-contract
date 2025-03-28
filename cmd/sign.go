package cmd

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
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
	cmd.Flags().StringP("password", "p", "", "singer wallet file password ")
	_ = cmd.MarkFlagRequired("password")
}

func signRawTx(cmd *cobra.Command, _ []string) {
	chainId, _ := cmd.Flags().GetInt64("chainId")
	file, _ := cmd.Flags().GetString("file")
	key, _ := cmd.Flags().GetString("key")
	password, _ := cmd.Flags().GetString("password")

	privateKey, err := ReadKeystoreFile(key, password)
	if err != nil {
		fmt.Printf("read wallet file error: %v", err)
		return
	}

	innerSignRawTx(chainId, file, privateKey)
}

func innerSignRawTx(chainId int64, file string, privateKey *ecdsa.PrivateKey) {
	// key = strings.TrimPrefix(key, "0x")
	// privateKey, err := crypto.HexToECDSA(key)
	// if err != nil {
	// 	fmt.Printf("invalid private key: %v", err)
	// 	return
	// }

	err := SignRawTx(file, privateKey, big.NewInt(chainId), addSuffixToFilename(file, "_signed"))
	if err != nil {
		fmt.Println("sign tx Err:", err)
		return
	}
	fmt.Println("sign tx success!")
}

func addSuffixToFilename(filename, suffix string) string {
	ext := filepath.Ext(filename)
	base := strings.TrimSuffix(filename, ext)
	return base + suffix + ext
}
