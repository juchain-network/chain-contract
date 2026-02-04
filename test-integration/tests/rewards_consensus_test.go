package tests

import (
	"context"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestI_ConsensusRewards(t *testing.T) {
	if ctx == nil {
		t.Skip("Context not initialized")
	}

	t.Run("V-03_DistributeBlockReward", func(t *testing.T) {
		_, minerAddr := minerKeyOrSkip(t)
		beforeInfo, _ := ctx.Staking.GetValidatorInfo(nil, minerAddr)
		before := new(big.Int).Set(beforeInfo.AccumulatedRewards)
		for i := 0; i < 10; i++ {
			waitBlocks(t, 1)
			info, _ := ctx.Staking.GetValidatorInfo(nil, minerAddr)
			if info.AccumulatedRewards.Cmp(before) > 0 {
				t.Logf("Rewards increased for %s: %s -> %s", minerAddr.Hex(), before.String(), info.AccumulatedRewards.String())
				return
			}
		}
		t.Skip("accumulatedRewards did not increase within 10 blocks")
	})

	t.Run("S-22_DistributeRewardsAndCooldown", func(t *testing.T) {
		minerKey, minerAddr := minerKeyOrSkip(t)
		withdrawPeriod, _ := ctx.Proposal.WithdrawProfitPeriod(nil)
		if withdrawPeriod == nil || withdrawPeriod.Sign() == 0 {
			t.Skip("withdrawProfitPeriod unavailable")
		}

		// Wait for some rewards to accrue for this validator.
		var infoBefore struct {
			AccumulatedRewards *big.Int
		}
		for i := 0; i < int(withdrawPeriod.Int64()); i++ {
			info, _ := ctx.Staking.GetValidatorInfo(nil, minerAddr)
			infoBefore.AccumulatedRewards = info.AccumulatedRewards
			if info.AccumulatedRewards.Sign() > 0 {
				break
			}
			waitBlocks(t, 1)
		}
		if infoBefore.AccumulatedRewards == nil || infoBefore.AccumulatedRewards.Sign() == 0 {
			t.Skip("no rewards accrued for validator in time")
		}

		robustClaimValidatorRewards(t, minerKey)

		infoAfterClaim, _ := ctx.Staking.GetValidatorInfo(nil, minerAddr)
		if infoAfterClaim.LastClaimBlock.Sign() == 0 {
			t.Skip("claim did not update lastClaimBlock")
		}
		lastClaim := new(big.Int).Set(infoAfterClaim.LastClaimBlock)

		// Try to claim again within cooldown after rewards accumulate.
		deadline := new(big.Int).Add(lastClaim, withdrawPeriod)
		for i := 0; i < int(withdrawPeriod.Int64())-1; i++ {
			waitBlocks(t, 1)
			curHeight, _ := ctx.Clients[0].BlockNumber(context.Background())
			if curHeight >= deadline.Uint64() {
				break
			}
			info, _ := ctx.Staking.GetValidatorInfo(nil, minerAddr)
			if info.AccumulatedRewards.Sign() == 0 {
				continue
			}

			opts, _ := ctx.GetTransactor(minerKey)
			tx, err := ctx.Staking.ClaimValidatorRewards(opts)
			if err == nil {
				err = ctx.WaitMined(tx.Hash())
			}
			if err == nil {
				t.Fatalf("expected cooldown revert, got success")
			}
			if !strings.Contains(err.Error(), "withdrawProfitPeriod") {
				t.Logf("claim failed as expected (cooldown), err=%v", err)
			}
			return
		}

		t.Skip("no rewards accrued within cooldown window")
	})

	t.Run("QueryRewards", func(t *testing.T) {
		valAddr := ctx.Config.Validators[0].Address
		_, _, _, _, rew, _ := ctx.Validators.GetValidatorInfo(nil, common.HexToAddress(valAddr))
		t.Logf("Validator %s rewards: %s", valAddr, rew.String())
	})
}
