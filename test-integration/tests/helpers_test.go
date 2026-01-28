package tests

import (
	"context"
	"crypto/ecdsa"
	"testing"

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
	if key == nil {
		t.Skipf("no local key for miner/coinbase %s", coinbase.Hex())
	}
	return key, coinbase
}

// waitForNextEpochBlock mines or waits until the NEXT block is an epoch block.
// Call this, then send the tx, then mine/wait 1 block to include it in an epoch block.
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
	// Blocks to mine so that (cur + blocksToMine + 1) % epoch == 0
	blocksToMine := (epoch - ((cur+1)%epoch)) % epoch
	if blocksToMine > 0 {
		waitBlocks(t, int(blocksToMine))
	}
	return epoch
}
