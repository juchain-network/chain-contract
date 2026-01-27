package tests

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestG_PunishPaths(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized")
	}

	minerKey, minerAddr := minerKeyOrSkip(t)

	// Pick a target validator different from miner if possible
	target := common.HexToAddress(ctx.Config.Validators[0].Address)
	if target == minerAddr && len(ctx.Config.Validators) > 1 {
		target = common.HexToAddress(ctx.Config.Validators[1].Address)
	}

	// [P-XX] Punish normal path (missed blocks counter++)
	t.Run("P-23_PunishNormal", func(t *testing.T) {
		opts, _ := ctx.GetTransactor(minerKey)
		tx, err := ctx.Punish.Punish(opts, target)
		if err != nil {
			if strings.Contains(err.Error(), "Miner only") {
				t.Skip("caller is not current miner")
			}
			t.Fatalf("punish failed: %v", err)
		}
		ctx.WaitMined(tx.Hash())

		rec, _ := ctx.Punish.GetPunishRecord(nil, target)
		utils.AssertTrue(t, rec.Cmp(big.NewInt(0)) > 0, "missedBlocksCounter should increase")
		length, _ := ctx.Punish.GetPunishValidatorsLen(nil)
		utils.AssertTrue(t, length.Cmp(big.NewInt(0)) > 0, "punishValidators length should be > 0")
	})

	// executePending no-op paths
	t.Run("P-24_ExecutePendingNoop", func(t *testing.T) {
		epoch, _ := ctx.Proposal.Epoch(nil)
		if epoch.Sign() > 0 {
			header, _ := ctx.Clients[0].HeaderByNumber(nil, nil)
			if header.Number.Uint64()%epoch.Uint64() == 0 {
				waitBlocks(t, 1)
			}
		}
		opts, _ := ctx.GetTransactor(ctx.GenesisValidators[0])
		tx, err := ctx.Punish.ExecutePending(opts, big.NewInt(0))
		utils.AssertNoError(t, err, "executePending(0) failed")
		ctx.WaitMined(tx.Hash())

		tx2, err := ctx.Punish.ExecutePending(opts, big.NewInt(1))
		utils.AssertNoError(t, err, "executePending(1) failed")
		ctx.WaitMined(tx2.Hash())
	})

	// decreaseMissedBlocksCounter at epoch
	t.Run("P-25_DecreaseMissedBlocksCounter", func(t *testing.T) {
		epoch, _ := ctx.Proposal.Epoch(nil)
		if epoch.Sign() == 0 {
			t.Skip("epoch not set")
		}

		// Align to epoch so the next mined block is an epoch block
		waitForNextEpochBlock(t)
		opts, _ := ctx.GetTransactor(minerKey)
		tx, err := ctx.Punish.DecreaseMissedBlocksCounter(opts, epoch)
		if err != nil {
			if strings.Contains(err.Error(), "Miner only") {
				t.Skip("caller is not current miner")
			}
			t.Fatalf("decreaseMissedBlocksCounter failed: %v", err)
		}
		// Mine the epoch block (if needed) and wait for receipt
		waitBlocks(t, 1)
		ctx.WaitMined(tx.Hash())
	})
}
