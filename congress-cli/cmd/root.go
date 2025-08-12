package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "congress-cli",
	Short: "Juchain blockchain governance command line tool",
	Long: `Congress CLI is a command line tool for Juchain blockchain governance.
It provides comprehensive functionality for validator management and proposal voting.

Features:
- Create and vote on validator addition/removal proposals
- Create and vote on configuration update proposals  
- Query validator information and manage rewards
- Sign and broadcast transactions to the blockchain network

Use "congress-cli [command] --help" for more information about a command.`,
	PersistentPreRun: validateGlobalFlags,
}

func validateGlobalFlags(cmd *cobra.Command, args []string) {
	// Skip validation for help and version commands
	if cmd.Name() == "help" || cmd.Name() == "version" {
		return
	}

	// Check if command requires RPC connection
	requiresRPC := []string{"miners", "miner", "create_proposal", "create_config_proposal", "vote_proposal", "withdraw_profits", "send"}
	cmdName := cmd.Name()

	for _, name := range requiresRPC {
		if cmdName == name {
			rpc, _ := cmd.Flags().GetString("rpc_laddr")
			if rpc == "" {
				PrintValidationError(fmt.Errorf("RPC endpoint is required for command '%s'", cmdName))
				os.Exit(1)
			}
			if err := ValidateRPCURL(rpc); err != nil {
				PrintValidationError(err)
				os.Exit(1)
			}
			break
		}
	}

	// Check if command requires chain ID
	requiresChainID := []string{"create_proposal", "create_config_proposal", "vote_proposal", "withdraw_profits", "sign"}

	for _, name := range requiresChainID {
		if cmdName == name {
			chainID, _ := cmd.Flags().GetInt64("chainId")
			if chainID == 0 {
				PrintValidationError(fmt.Errorf("chain ID is required for command '%s'", cmdName))
				os.Exit(1)
			}
			if err := ValidateChainID(chainID); err != nil {
				PrintValidationError(err)
				os.Exit(1)
			}
			break
		}
	}
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		PrintError("Command execution failed", err)
		os.Exit(1)
	}
}

func init() {
	// Global flags
	rootCmd.PersistentFlags().StringP("rpc_laddr", "l", "", "RPC endpoint URL (e.g., http://localhost:8545)")
	rootCmd.PersistentFlags().Int64P("chainId", "c", 0, "Chain ID (e.g., 202599)")

	// Add subcommands
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
