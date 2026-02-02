package tests

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func keyForAddress(addr common.Address) *ecdsa.PrivateKey {
	if ctx == nil {
		return nil
	}
	for _, k := range ctx.GenesisValidators {
		if crypto.PubkeyToAddress(k.PublicKey) == addr {
			return k
		}
	}
	return nil
}

func minerKeyOrSkip(t *testing.T) (*ecdsa.PrivateKey, common.Address) {
	if ctx == nil || len(ctx.Clients) == 0 {
		t.Skip("Context not initialized")
	}
	header, err := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
	if err != nil {
		t.Skipf("failed to read header: %v", err)
	}
	coinbase := header.Coinbase
	key := keyForAddress(coinbase)

	if key != nil {
		return key, coinbase
	}

	// Try all validators if coinbase doesn't match (might be a system node)
	for _, k := range ctx.GenesisValidators {
		return k, crypto.PubkeyToAddress(k.PublicKey)
	}

	t.Skip("no validator keys available")
	return nil, common.Address{}
}

// waitForNextEpochBlock mines or waits until an epoch block has passed and state has settled.
// It also handles triggering updateActiveValidatorSet as the miner if we are at an epoch block.
func waitForNextEpochBlock(t *testing.T) uint64 {
	if ctx == nil {
		t.Skip("Context not initialized")
	}
	epochBig, err := ctx.Proposal.Epoch(nil)
	if err != nil || epochBig.Sign() == 0 {
		t.Skip("epoch not available")
	}
	epoch := epochBig.Uint64()
	header, err := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
	if err != nil {
		t.Skipf("failed to read header: %v", err)
	}
	cur := header.Number.Uint64()

	// Target is the next multiple of epoch
	nextEpochBlock := ((cur / epoch) + 1) * epoch
	blocksToWait := nextEpochBlock - cur

	t.Logf("Current block %d, next epoch block %d, waiting %d blocks...", cur, nextEpochBlock, blocksToWait)

	// Wait until we are at nextEpochBlock - 1
	if blocksToWait > 1 {
		waitBlocks(t, int(blocksToWait-1))
	}

	// Now wait block by block until we hit exactly nextEpochBlock
	for {
		header, _ := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
		if header.Number.Uint64() >= nextEpochBlock {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}

	// Trigger the validator set update.
	highest, _ := ctx.Validators.GetHighestValidators(nil)
	top, _ := ctx.Staking.GetTopValidators(nil, highest)

	success := false
	// Try aggressively to trigger the update
	for retry := 0; retry < 10; retry++ {
		curHeader, _ := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
		curHeight := curHeader.Number.Uint64()

		// If we passed the window, we might have missed it or it happened
		if curHeight > nextEpochBlock+5 {
			break
		}

		for _, vk := range ctx.GenesisValidators {
			addr := crypto.PubkeyToAddress(vk.PublicKey)

			// Refresh nonce to be safe; allow epoch-block tx for validator-set update
			ctx.RefreshNonce(addr)
			opts, _ := ctx.GetTransactorNoEpochWait(vk, true)

			// Try to update for the current block OR the next one
			// We try both because we might be at the end of the block
			targets := []uint64{curHeight, curHeight + 1}

			for _, target := range targets {
				if target%epoch != 0 {
					continue
				}
				tx, err := ctx.Validators.UpdateActiveValidatorSet(opts, top, big.NewInt(int64(target)))
				if err == nil {
					t.Logf("Triggered UpdateActiveValidatorSet successfully at block %d from %s", target, addr.Hex())
					ctx.WaitMined(tx.Hash())
					success = true
					goto Done
				}
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
Done:

	if !success {
		t.Log("UpdateActiveValidatorSet was not triggered (might be already done or missed)")
	}

	newHeight, _ := ctx.Clients[0].BlockNumber(context.Background())
	fmt.Printf("Epoch wait complete. New height: %d\n", newHeight)
	ctx.SyncNonces()
	return newHeight
}
