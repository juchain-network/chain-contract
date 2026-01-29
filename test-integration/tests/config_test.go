package tests

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"juchain.org/chain/tools/ci/internal/utils"
)

// Config IDs from Proposal.sol
const (
	ConfigID_ProposalLastingPeriod    = 0
	ConfigID_PunishThreshold          = 1
	ConfigID_RemoveThreshold          = 2
	ConfigID_DecreaseRate             = 3
	ConfigID_WithdrawProfitPeriod     = 4
	ConfigID_BlockReward              = 5
	ConfigID_UnbondingPeriod          = 6
	ConfigID_ValidatorUnjailPeriod    = 7
	ConfigID_MinValidatorStake        = 8
	ConfigID_MaxValidators            = 9
	ConfigID_MinDelegation            = 10
	ConfigID_MinUndelegation          = 11
	ConfigID_DoubleSignSlashAmount    = 12
	ConfigID_DoubleSignRewardAmount   = 13
	ConfigID_BurnAddress              = 14
	ConfigID_DoubleSignWindow         = 15
	ConfigID_CommissionUpdateCooldown = 16
	ConfigID_BaseRewardRatio          = 17
	ConfigID_MaxCommissionRate        = 18
	ConfigID_ProposalCooldown         = 19
)

func TestA_SystemConfigSetup(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized or no validators")
	}

	// Define target parameters for testing environment
	targets := []struct {
		name string
		cid  int64
		val  *big.Int
	}{
		{"ProposalCooldown", ConfigID_ProposalCooldown, big.NewInt(1)},
		{"UnbondingPeriod", ConfigID_UnbondingPeriod, big.NewInt(100)},
		{"ValidatorUnjailPeriod", ConfigID_ValidatorUnjailPeriod, big.NewInt(50)},
		{"WithdrawProfitPeriod", ConfigID_WithdrawProfitPeriod, big.NewInt(20)},
		{"MinValidatorStake", ConfigID_MinValidatorStake, utils.ToWei(1)},
		{"MinDelegation", ConfigID_MinDelegation, utils.ToWei(1)},
		{"CommissionUpdateCooldown", ConfigID_CommissionUpdateCooldown, big.NewInt(50)},
		{"ProposalLastingPeriod", ConfigID_ProposalLastingPeriod, big.NewInt(200)},
	}

	for _, target := range targets {
		t.Logf("Ensuring %s is %v...", target.name, target.val)
		err := ctx.EnsureConfig(target.cid, target.val, nil)
		utils.AssertNoError(t, err, fmt.Sprintf("failed to ensure %s", target.name))
	}
}

func TestB_ConfigBoundaryChecks(t *testing.T) {
	if ctx == nil {
		t.Skip("Context not initialized")
	}

	// Wait for any previous cooldown to expire
	t.Log("Waiting for potential proposal cooldown...")
	waitBlocks(t, 10) // Wait enough blocks (previous test set it to 1, but safety margin)

	proposerKey := ctx.GenesisValidators[0]

	runRevertTest := func(name string, cid uint256, val *big.Int, expectedErr string) {
		t.Run(name, func(t *testing.T) {
			opts, _ := ctx.GetTransactor(proposerKey)
			// Using a fresh nonce is handled by GetTransactor if we actually submitted previously.
			// However, since we expect revert, the nonce might not increment on chain if we simulate?
			// But CreateUpdateConfigProposal submits a transaction.
			// Go-ethereum simulation usually detects revert before broadcast.
			_, err := ctx.Proposal.CreateUpdateConfigProposal(opts, big.NewInt(int64(cid)), val)

			if err == nil {
				t.Fatalf("expected error containing %q, got nil", expectedErr)
			}
			if !strings.Contains(err.Error(), expectedErr) {
				t.Logf("Got error: %v", err)
				if !strings.Contains(err.Error(), "revert") && !strings.Contains(err.Error(), "execution reverted") {
					t.Errorf("Unexpected error type")
				}
				// Strict check enabled
				t.Errorf("expected error %q, got %q", expectedErr, err.Error())
			}
		})
	}

	// [C-02] General Validation
	runRevertTest("Invalid Config ID", 20, big.NewInt(100), "Invalid config ID")
	runRevertTest("Zero Value", ConfigID_ProposalCooldown, big.NewInt(0), "Config value must be positive")

	// [C-03] Threshold Logic
	// Assuming current values: Punish=24, Remove=48, Decrease=24
	runRevertTest("Punish >= Remove", ConfigID_PunishThreshold, big.NewInt(48), "punishThreshold must be < removeThreshold")
	runRevertTest("Remove <= Punish", ConfigID_RemoveThreshold, big.NewInt(24), "removeThreshold must be > punishThreshold")
	runRevertTest("Decrease > Remove", ConfigID_DecreaseRate, big.NewInt(49), "decreaseRate must be <= removeThreshold")

	// [C-04] Consensus & Safety
	runRevertTest("Max Validators Overflow", ConfigID_MaxValidators, big.NewInt(22), "maxValidators exceeds consensus limit")
	// Generic positive check catches zero address first
	runRevertTest("Zero Burn Address", ConfigID_BurnAddress, big.NewInt(0), "Config value must be positive")
	// Burn address out of uint160 range
	burnTooLarge := new(big.Int).Lsh(big.NewInt(1), 160)
	runRevertTest("Burn Address Overflow", ConfigID_BurnAddress, burnTooLarge, "burnAddress invalid")

	// [C-05] Economic
	// Default DoubleSignSlash=50000 ether, Reward=10000 ether
	// Slash < Reward -> Set Slash to 1 wei
	runRevertTest("Slash < Reward", ConfigID_DoubleSignSlashAmount, big.NewInt(1), "doubleSignSlashAmount must be >= doubleSignRewardAmount")

	// Reward > Slash -> Set Reward to 60000 ether
	rewardTooHigh := utils.ToWei(60000)
	runRevertTest("Reward > Slash", ConfigID_DoubleSignRewardAmount, rewardTooHigh, "doubleSignRewardAmount must be <= doubleSignSlashAmount")

	// Invalid Base Ratio (> 10000)
	runRevertTest("Invalid Base Ratio", ConfigID_BaseRewardRatio, big.NewInt(10001), "baseRewardRatio must be <= 10000")

	// Invalid Max Commission (> 10000)
	runRevertTest("Invalid Max Commission", ConfigID_MaxCommissionRate, big.NewInt(10001), "maxCommissionRate must be <= 10000")
}

// Helper type for uint256 since I used it in struct definition
type uint256 = uint64
