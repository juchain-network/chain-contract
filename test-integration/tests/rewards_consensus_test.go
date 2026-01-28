package tests

import (
	"strings"
	"testing"

	"juchain.org/chain/tools/ci/internal/utils"
)

func TestI_ConsensusRewards(t *testing.T) {
	if ctx == nil {
		t.Skip("Context not initialized")
	}

	minerKey, minerAddr := minerKeyOrSkip(t)

	// Validators.distributeBlockReward (tx fee reward)
	t.Run("V-03_DistributeBlockReward", func(t *testing.T) {
		opts, _ := ctx.GetTransactor(minerKey)
		opts.Value = utils.ToWei(1)

		_, _, incomingBefore, _, _, _ := ctx.Validators.GetValidatorInfo(nil, minerAddr)
		tx, err := ctx.Validators.DistributeBlockReward(opts)
		if err != nil {
			if strings.Contains(err.Error(), "Miner only") {
				t.Skip("caller is not current miner")
			}
			if strings.Contains(err.Error(), "Block is already rewarded") {
				t.Skip("block reward already distributed in this block")
			}
			t.Fatalf("distributeBlockReward failed: %v", err)
		}
		ctx.WaitMined(tx.Hash())

		_, _, incomingAfter, _, _, _ := ctx.Validators.GetValidatorInfo(nil, minerAddr)
		jailed, _ := ctx.Staking.IsValidatorJailed(nil, minerAddr)
		if jailed {
			utils.AssertBigIntEq(t, incomingAfter, incomingBefore, "jailed miner should not gain fee income")
		} else {
			utils.AssertTrue(t, incomingAfter.Cmp(incomingBefore) > 0, "fee income should increase")
		}
	})

	// Staking.distributeRewards + claim cooldown
	t.Run("S-22_DistributeRewardsAndCooldown", func(t *testing.T) {
		opts, _ := ctx.GetTransactor(minerKey)
		opts.Value = utils.ToWei(1)

		infoBefore, _ := ctx.Staking.GetValidatorInfo(nil, minerAddr)
		tx, err := ctx.Staking.DistributeRewards(opts)
		if err != nil {
			if strings.Contains(err.Error(), "Miner only") {
				t.Skip("caller is not current miner")
			}
			t.Fatalf("distributeRewards failed: %v", err)
		}
		ctx.WaitMined(tx.Hash())

		infoAfter, _ := ctx.Staking.GetValidatorInfo(nil, minerAddr)
		if infoAfter.AccumulatedRewards.Cmp(infoBefore.AccumulatedRewards) <= 0 {
			t.Skip("rewards not accumulated (possibly already distributed this block)")
		}

		// Claim once (should succeed)
		claimOpts, _ := ctx.GetTransactor(minerKey)
		txC, err := ctx.Staking.ClaimValidatorRewards(claimOpts)
		if err != nil {
			t.Fatalf("claimValidatorRewards failed: %v", err)
		}
		ctx.WaitMined(txC.Hash())

		// Distribute again and attempt immediate claim to trigger cooldown
		opts2, _ := ctx.GetTransactor(minerKey)
		opts2.Value = utils.ToWei(1)
		tx2, err := ctx.Staking.DistributeRewards(opts2)
		if err == nil {
			ctx.WaitMined(tx2.Hash())
		}

		claimOpts2, _ := ctx.GetTransactor(minerKey)
		_, err = ctx.Staking.ClaimValidatorRewards(claimOpts2)
		if err == nil {
			t.Fatal("should fail claimValidatorRewards due to withdrawProfitPeriod cooldown")
		}
	})
}
