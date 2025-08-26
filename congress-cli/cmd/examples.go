package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func helpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "examples",
		Short: "Show usage examples",
		Long:  "Show common usage examples for congress-cli commands",
		Run:   showExamples,
	}
	return cmd
}

func showExamples(cmd *cobra.Command, _ []string) {
	fmt.Println("Congress CLI Usage Examples")
	fmt.Println("===========================")
	fmt.Println()

	fmt.Println("🔍 QUERY COMMANDS")
	fmt.Println()

	fmt.Println("1. Query all validators:")
	fmt.Println("   # Test network")
	fmt.Println("   congress-cli miners -c 2025 -l https://testnet-rpc.juchain.org")
	fmt.Println("   # Main network")
	fmt.Println("   congress-cli miners -c 210000 -l https://rpc.juchain.org")
	fmt.Println()

	fmt.Println("2. Query specific validator:")
	fmt.Println("   # Test network")
	fmt.Println("   congress-cli miner -c 2025 -l https://testnet-rpc.juchain.org \\")
	fmt.Println("     -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b")
	fmt.Println("   # Main network")
	fmt.Println("   congress-cli miner -c 210000 -l https://rpc.juchain.org \\")
	fmt.Println("     -a 0x311B37f01c04B84d1f94645BfBd58D82fc03F709")
	fmt.Println()

	fmt.Println("📝 PROPOSAL COMMANDS")
	fmt.Println()

	fmt.Println("3. Create validator addition proposal:")
	fmt.Println("   congress-cli create_proposal -c 2025 -l https://testnet-rpc.juchain.org \\")
	fmt.Println("     -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \\")
	fmt.Println("     -t 0x029DAB47e268575D4AC167De64052FB228B5fA41 \\")
	fmt.Println("     -o add")
	fmt.Println()

	fmt.Println("4. Create validator removal proposal:")
	fmt.Println("   congress-cli create_proposal -c 210000 -l https://rpc.juchain.org \\")
	fmt.Println("     -p 0xccafa71c31bc11ba24d526fd27ba57d743152807 \\")
	fmt.Println("     -t 0x4d432df142823ca25b21bc3f9744ed21a275bdea \\")
	fmt.Println("     -o remove")
	fmt.Println()

	fmt.Println("5. Create configuration update proposal:")
	fmt.Println("   congress-cli create_config_proposal -c 2025 -l https://testnet-rpc.juchain.org \\")
	fmt.Println("     -p 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \\")
	fmt.Println("     -i 0 -v 86400")
	fmt.Println()

	fmt.Println("🗳️  VOTING COMMANDS")
	fmt.Println()

	fmt.Println("6. Vote APPROVE on proposal:")
	fmt.Println("   congress-cli vote_proposal -c 2025 -l https://testnet-rpc.juchain.org \\")
	fmt.Println("     -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \\")
	fmt.Println("     -i PROPOSAL_ID \\")
	fmt.Println("     -a")
	fmt.Println()

	fmt.Println("7. Vote REJECT on proposal:")
	fmt.Println("   congress-cli vote_proposal -c 2025 -l https://testnet-rpc.juchain.org \\")
	fmt.Println("     -s 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b \\")
	fmt.Println("     -i PROPOSAL_ID")
	fmt.Println("   # Note: Omit -a flag for reject vote")
	fmt.Println()

	fmt.Println("✍️  TRANSACTION COMMANDS")
	fmt.Println()

	fmt.Println("8. Sign transaction:")
	fmt.Println("   # Test network")
	fmt.Println("   congress-cli sign -f createProposal.json \\")
	fmt.Println("     -k miner1.key -p password.file -c 2025")
	fmt.Println("   # Main network")
	fmt.Println("   congress-cli sign -f createProposal.json \\")
	fmt.Println("     -k /path/to/keystore/UTC--xxx -p /path/to/password.txt -c 210000")
	fmt.Println()

	fmt.Println("9. Broadcast signed transaction:")
	fmt.Println("   # Test network")
	fmt.Println("   congress-cli send -f createProposal_signed.json \\")
	fmt.Println("     -l https://testnet-rpc.juchain.org")
	fmt.Println("   # Main network")
	fmt.Println("   congress-cli send -f createProposal_signed.json \\")
	fmt.Println("     -l https://rpc.juchain.org")
	fmt.Println()

	fmt.Println("💰 REWARD COMMANDS")
	fmt.Println()

	fmt.Println("10. Withdraw validator profits:")
	fmt.Println("    # Test network")
	fmt.Println("    congress-cli withdraw_profits -c 2025 -l https://testnet-rpc.juchain.org \\")
	fmt.Println("      -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b")
	fmt.Println("    # Main network")
	fmt.Println("    congress-cli withdraw_profits -c 210000 -l https://rpc.juchain.org \\")
	fmt.Println("      -a 0xccafa71c31bc11ba24d526fd27ba57d743152807")
	fmt.Println()

	fmt.Println("📋 REFERENCE")
	fmt.Println()
	fmt.Println("Configuration IDs (-i parameter):")
	fmt.Println("  0: proposalLastingPeriod (proposal duration)")
	fmt.Println("  1: punishThreshold (punishment threshold)")
	fmt.Println("  2: removeThreshold (removal threshold)")
	fmt.Println("  3: decreaseRate (decrease rate)")
	fmt.Println("  4: withdrawProfitPeriod (withdrawal period)")
	fmt.Println()

	fmt.Println("Network Information:")
	fmt.Println("  Test Network: Chain ID 2025, RPC https://testnet-rpc.juchain.org")
	fmt.Println("  Main Network: Chain ID 210000, RPC https://rpc.juchain.org")
	fmt.Println()

	fmt.Println("💡 TIPS")
	fmt.Println("- Always record the Proposal ID from create_proposal output")
	fmt.Println("- Use -a flag for APPROVE vote, omit for REJECT vote")
	fmt.Println("- Ensure validator node is synced before voting")
	fmt.Println("- Check validator status before creating proposals")
	fmt.Println()
}

func init() {
	rootCmd.AddCommand(helpCmd())
}
