package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func helpCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "guide",
		Short: "Show usage guide",
		Long:  "Show comprehensive usage guide and examples for congress-cli commands",
		Run:   showUsageGuide,
	}
	return cmd
}

func showUsageGuide(cmd *cobra.Command, _ []string) {
	fmt.Println("Congress CLI Usage Examples")
	fmt.Println("===========================")
	fmt.Println()

	fmt.Println("⚙️  CONFIG COMMANDS")
	fmt.Println()

	fmt.Println("1. View current configuration:")
	fmt.Println("   congress-cli config get")
	fmt.Println("   congress-cli config list")
	fmt.Println()

	fmt.Println("2. Set configuration for local network:")
	fmt.Println("   congress-cli config set --rpc http://localhost:8545 --chain-id 202599")
	fmt.Println()

	fmt.Println("3. Set configuration for test network:")
	fmt.Println("   congress-cli config set --rpc https://testnet-rpc.juchain.org --chain-id 202599")
	fmt.Println()

	fmt.Println("4. Set configuration for main network:")
	fmt.Println("   congress-cli config set --rpc https://rpc.juchain.org --chain-id 210000")
	fmt.Println()

	fmt.Println("5. Get specific config value:")
	fmt.Println("   congress-cli config get --rpc")
	fmt.Println("   congress-cli config get --chain-id")
	fmt.Println()

	fmt.Println("🔍 QUERY COMMANDS")
	fmt.Println()

	fmt.Println("6. Query all validators:")
	fmt.Println("   # Local network (default)")
	fmt.Println("   congress-cli miners")
	fmt.Println("   # With explicit RPC")
	fmt.Println("   congress-cli miners -c 202599 -l https://testnet-rpc.juchain.org")
	fmt.Println("   congress-cli miners -c 210000 -l https://rpc.juchain.org")
	fmt.Println()

	fmt.Println("7. Query specific validator:")
	fmt.Println("   # Local network")
	fmt.Println("   congress-cli miner -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	fmt.Println("   # Test network")
	fmt.Println("   congress-cli miner -c 202599 -l https://testnet-rpc.juchain.org \\")
	fmt.Println("     -a 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b")
	fmt.Println()

	fmt.Println("8. Query all proposals:")
	fmt.Println("   congress-cli proposals")
	fmt.Println()

	fmt.Println("9. Query specific proposal:")
	fmt.Println("   congress-cli proposal -i <PROPOSAL_ID>")
	fmt.Println()

	fmt.Println("🥩 STAKING COMMANDS")
	fmt.Println()

	fmt.Println("10. Query top validators in staking:")
	fmt.Println("    congress-cli staking list-top-validators")
	fmt.Println()

	fmt.Println("11. Query validator staking info:")
	fmt.Println("    congress-cli staking query-validator --address 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	fmt.Println()

	fmt.Println("12. Register validator for staking:")
	fmt.Println("    congress-cli staking register-validator \\")
	fmt.Println("      --proposer 0xYOUR_VALIDATOR_ADDRESS \\")
	fmt.Println("      --stake-amount 10000 \\")
	fmt.Println("      --commission-rate 500")
	fmt.Println()

	fmt.Println("13. Edit validator info:")
	fmt.Println("    congress-cli staking edit-validator \\")
	fmt.Println("      --validator 0xYOUR_VALIDATOR_ADDRESS \\")
	fmt.Println("      --fee-addr 0xYOUR_FEE_ADDRESS \\")
	fmt.Println("      --moniker \"Your Validator Name\" \\")
	fmt.Println("      --details \"Validator description\"")
	fmt.Println()

	fmt.Println("📝 PROPOSAL COMMANDS")
	fmt.Println()

	fmt.Println("14. Create validator addition proposal:")
	fmt.Println("    congress-cli create_proposal \\")
	fmt.Println("      -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \\")
	fmt.Println("      -t 0x50C554aC9c134491818fa6f21d504f2AE5BD9c26 \\")
	fmt.Println("      -o add")
	fmt.Println()

	fmt.Println("15. Create validator removal proposal:")
	fmt.Println("    congress-cli create_proposal \\")
	fmt.Println("      -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \\")
	fmt.Println("      -t 0x50C554aC9c134491818fa6f21d504f2AE5BD9c26 \\")
	fmt.Println("      -o remove")
	fmt.Println()

	fmt.Println("16. Create configuration update proposal:")
	fmt.Println("    congress-cli create_config_proposal \\")
	fmt.Println("      -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \\")
	fmt.Println("      -i 0 -v 86400")
	fmt.Println()

	fmt.Println("🗳️  VOTING COMMANDS")
	fmt.Println()

	fmt.Println("17. Vote APPROVE on proposal:")
	fmt.Println("    congress-cli vote_proposal \\")
	fmt.Println("      -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \\")
	fmt.Println("      -i <PROPOSAL_ID> \\")
	fmt.Println("      -a")
	fmt.Println()

	fmt.Println("18. Vote REJECT on proposal:")
	fmt.Println("    congress-cli vote_proposal \\")
	fmt.Println("      -s 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \\")
	fmt.Println("      -i <PROPOSAL_ID>")
	fmt.Println("    # Note: Omit -a flag for reject vote")
	fmt.Println()

	fmt.Println("✍️  TRANSACTION COMMANDS")
	fmt.Println()

	fmt.Println("19. Sign transaction:")
	fmt.Println("    congress-cli sign -f createProposal.json \\")
	fmt.Println("      -k /path/to/keystore/UTC--xxx \\")
	fmt.Println("      -p /path/to/password.txt")
	fmt.Println()

	fmt.Println("20. Broadcast signed transaction:")
	fmt.Println("    congress-cli send -f createProposal_signed.json")
	fmt.Println()

	fmt.Println("💰 REWARD COMMANDS")
	fmt.Println()

	fmt.Println("21. Withdraw validator profits:")
	fmt.Println("    congress-cli withdraw_profits -a 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	fmt.Println()

	fmt.Println("� UTILITY COMMANDS")
	fmt.Println()

	fmt.Println("22. Check version:")
	fmt.Println("    congress-cli version")
	fmt.Println()

	fmt.Println("23. View help:")
	fmt.Println("    congress-cli --help")
	fmt.Println("    congress-cli [command] --help")
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

	fmt.Println("Current Network Validators:")
	fmt.Println("  Validator1: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266")
	fmt.Println("  Validator2: 0x70997970C51812dc3A010C7d01b50e0d17dc79C8")
	fmt.Println("  Validator3: 0x3C44CdDdB6a900fa2b585dd299e03d12FA4293BC")
	fmt.Println("  Validator4: 0x90F79bf6EB2c4f870365E785982E1f101E93b906")
	fmt.Println("  Validator5: 0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")
	fmt.Println("  Validator6: 0x9965507D1a55bcC2695C58ba16FB37d819B0A4dc")
	fmt.Println("  Validator7: 0x50C554aC9c134491818fa6f21d504f2AE5BD9c26")
	fmt.Println()

	fmt.Println("Network Information:")
	fmt.Println("  Local Network: Chain ID 202599, RPC http://localhost:8545")
	fmt.Println("  Test Network: Chain ID 202599, RPC https://testnet-rpc.juchain.org")
	fmt.Println("  Main Network: Chain ID 210000, RPC https://rpc.juchain.org")
	fmt.Println()

	fmt.Println("💡 TIPS")
	fmt.Println("- Always record the Proposal ID from create_proposal output")
	fmt.Println("- Use -a flag for APPROVE vote, omit for REJECT vote")
	fmt.Println("- Ensure validator node is synced before voting")
	fmt.Println("- Check validator status before creating proposals")
	fmt.Println("- Use 'congress-cli config set' to avoid repeating --rpc and --chain-id flags")
	fmt.Println("- Config file is stored at ~/.congress-cli/config.json")
	fmt.Println("- Command line flags override config file settings")
	fmt.Println("- Local network defaults to localhost:8545 with Chain ID 202599")
	fmt.Println("- Use 'congress-cli guide' to see this help anytime")
	fmt.Println()

	fmt.Println("📖 COMPLETE WORKFLOW EXAMPLE: Adding a New Validator")
	fmt.Println()
	fmt.Println("Step 1: Configure your environment")
	fmt.Println("    congress-cli config set --rpc http://localhost:8545 --chain-id 202599")
	fmt.Println()
	fmt.Println("Step 2: Check current validators")
	fmt.Println("    congress-cli miners")
	fmt.Println()
	fmt.Println("Step 3: Create addition proposal")
	fmt.Println("    congress-cli create_proposal \\")
	fmt.Println("      -p 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266 \\")
	fmt.Println("      -t 0xNEW_VALIDATOR_ADDRESS \\")
	fmt.Println("      -o add")
	fmt.Println()
	fmt.Println("Step 4: Sign the proposal")
	fmt.Println("    congress-cli sign -f createProposal.json \\")
	fmt.Println("      -k /path/to/validator1/keystore/UTC--xxx \\")
	fmt.Println("      -p /path/to/validator1/password.txt")
	fmt.Println()
	fmt.Println("Step 5: Broadcast the proposal")
	fmt.Println("    congress-cli send -f createProposal_signed.json")
	fmt.Println("    # Record the Proposal ID from output")
	fmt.Println()
	fmt.Println("Step 6: All validators vote (repeat for each validator)")
	fmt.Println("    congress-cli vote_proposal \\")
	fmt.Println("      -s 0xVALIDATOR_ADDRESS \\")
	fmt.Println("      -i PROPOSAL_ID \\")
	fmt.Println("      -a")
	fmt.Println("    congress-cli sign -f voteProposal.json \\")
	fmt.Println("      -k /path/to/validator/keystore/UTC--xxx \\")
	fmt.Println("      -p /path/to/validator/password.txt")
	fmt.Println("    congress-cli send -f voteProposal_signed.json")
	fmt.Println()
	fmt.Println("Step 7: Verify the new validator was added")
	fmt.Println("    congress-cli miners")
	fmt.Println("    congress-cli miner -a 0xNEW_VALIDATOR_ADDRESS")
	fmt.Println()
}

func init() {
	rootCmd.AddCommand(helpCmd())
}
