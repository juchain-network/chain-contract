package tests

import (
	"strings"
	"testing"

	"juchain.org/chain/tools/ci/internal/utils"
)

func TestY_UpdateActiveValidatorSet(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized")
	}

	minerKey, _ := minerKeyOrSkip(t)
	epoch, _ := ctx.Proposal.Epoch(nil)
	if epoch.Sign() == 0 {
		t.Skip("epoch not set")
	}

	// Non-epoch call should fail
	t.Run("V-07_UpdateSetNonEpoch", func(t *testing.T) {
		opts, _ := ctx.GetTransactor(minerKey)
		current, _ := ctx.Validators.GetActiveValidators(nil)
		_, err := ctx.Validators.UpdateActiveValidatorSet(opts, current, epoch)
		if err == nil {
			t.Log("UpdateActiveValidatorSet succeeded (likely already at epoch), skipping non-epoch check")
			return
		}
		if strings.Contains(err.Error(), "Miner only") || strings.Contains(err.Error(), "forbidden system transaction") {
			t.Skip("caller is not current miner or system blocked")
		}
		if !strings.Contains(err.Error(), "Block epoch only") {
			t.Fatalf("unexpected error: %v", err)
		}
	})

	// Normal epoch update path
	t.Run("V-08_UpdateSetEpoch", func(t *testing.T) {
		waitForNextEpochBlock(t)

		highest, err := ctx.Validators.GetHighestValidators(nil)
		utils.AssertNoError(t, err, "getHighestValidators failed")
		expected, err := ctx.Staking.GetTopValidators(nil, highest)
		utils.AssertNoError(t, err, "getTopValidators failed")
		if len(expected) == 0 {
			t.Skip("expected validator set empty")
		}

		opts, _ := ctx.GetTransactor(minerKey)
		tx, err := ctx.Validators.UpdateActiveValidatorSet(opts, expected, epoch)
		if err != nil {
			if strings.Contains(err.Error(), "Miner only") || strings.Contains(err.Error(), "forbidden system transaction") {
				t.Skip("caller is not current miner or system blocked")
			}
			t.Fatalf("updateActiveValidatorSet failed: %v", err)
		}
		// Mine the epoch block (if needed) and wait
		waitBlocks(t, 1)
		ctx.WaitMined(tx.Hash())

		newSet, _ := ctx.Validators.GetActiveValidators(nil)
		if len(newSet) != len(expected) {
			t.Fatalf("validator set length mismatch: expected %d, got %d", len(expected), len(newSet))
		}
	})
}