package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "juchain",
	Short: "Juchain contract command line",
	Long:  "Juchain contract command line",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	// Global flags or subcommand additions here
	rootCmd.PersistentFlags().StringP("rpc_laddr", "l", "", "remote rpc addr")
	rootCmd.PersistentFlags().Int64P("chainId", "c", 0, "chain id")

	rootCmd.AddCommand(
		CreateProposalCmd(),
		CreateConfigProposalCmd(),
		VoteProposalCmd(),
		SignRawTxCmd(),
		SendSignedTxCmd(),
		WithdrawProfitsCmd(),
		ValidatorsCmd(),
		ValidatorCmd(),
	)
}
