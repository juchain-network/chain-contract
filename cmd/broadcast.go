package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func SendSignedTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "send",
		Short: "send signed tx from file to remote rpc endpoint",
		Run:   sendSignedTx,
	}
	sendSignedTxCmdFlags(cmd)
	return cmd
}

func sendSignedTxCmdFlags(cmd *cobra.Command) {
	cmd.Flags().StringP("proposer", "p", "", "proposer addr (must be valid validator)")
	_ = cmd.MarkFlagRequired("proposer")
	cmd.Flags().StringP("file", "f", "", "signed tx file")
	_ = cmd.MarkFlagRequired("file")
}

func sendSignedTx(cmd *cobra.Command, _ []string) {
	file, _ := cmd.Flags().GetString("file")
	rpc, _ := cmd.Flags().GetString("rpc_laddr")
	proposer, _ := cmd.Flags().GetString("proposer")

	innerSendSignedTx(proposer, file, rpc)
}

func innerSendSignedTx(proposer, file, rpc string) {
	_, err := SendSignedTx(proposer, rpc, file)
	if err != nil {
		fmt.Println("send tx Err:", err)
		return
	}
	fmt.Println("send tx success!")
}
