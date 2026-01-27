package tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"juchain.org/chain/tools/ci/internal/utils"
)

// Config IDs from Proposal.sol
const (
	ConfigID_ProposalLastingPeriod   = 0
	ConfigID_WithdrawProfitPeriod    = 4
	ConfigID_UnbondingPeriod         = 6
	ConfigID_ValidatorUnjailPeriod   = 7
	ConfigID_CommissionUpdateCooldown = 16
	ConfigID_ProposalCooldown        = 19
)

func TestA_SystemConfigSetup(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized or no validators")
	}

	// Define target parameters for testing environment
	targets := []struct {
		name string
		cid  uint256
		val  uint256
		getter func() (*big.Int, error)
	}{
		// Prioritize ProposalCooldown to speed up subsequent tests
		{
			name: "ProposalCooldown",
			cid:  ConfigID_ProposalCooldown,
			val:  10, // 10 blocks
			getter: func() (*big.Int, error) { return ctx.Proposal.ProposalCooldown(nil) },
		},
		{
			name: "UnbondingPeriod",
			cid:  ConfigID_UnbondingPeriod,
			val:  100, // 100 blocks
			getter: func() (*big.Int, error) { return ctx.Proposal.UnbondingPeriod(nil) },
		},
		{
			name: "ValidatorUnjailPeriod",
			cid:  ConfigID_ValidatorUnjailPeriod,
			val:  50, // 50 blocks
			getter: func() (*big.Int, error) { return ctx.Proposal.ValidatorUnjailPeriod(nil) },
		},
		{
			name: "WithdrawProfitPeriod",
			cid:  ConfigID_WithdrawProfitPeriod,
			val:  20, // 20 blocks
			getter: func() (*big.Int, error) { return ctx.Proposal.WithdrawProfitPeriod(nil) },
		},
		{
			name: "CommissionUpdateCooldown",
			cid:  ConfigID_CommissionUpdateCooldown,
			val:  50, // 50 blocks
			getter: func() (*big.Int, error) { return ctx.Proposal.CommissionUpdateCooldown(nil) },
		},
		// We set ProposalLastingPeriod last to avoid expiring previous proposals too fast during this setup
		{
			name: "ProposalLastingPeriod",
			cid:  ConfigID_ProposalLastingPeriod,
			val:  200, // 200 blocks
			getter: func() (*big.Int, error) { return ctx.Proposal.ProposalLastingPeriod(nil) },
		},
	}

	proposerKey := ctx.GenesisValidators[0]

	for _, target := range targets {
		t.Logf("Updating %s to %d...", target.name, target.val)

		// 1. Check current value
		current, err := target.getter()
		utils.AssertNoError(t, err, "failed to get current value")
		
		targetVal := big.NewInt(int64(target.val))
		if current.Cmp(targetVal) == 0 {
			t.Logf("%s is already %s, skipping", target.name, current)
			continue
		}

		// 2. Create Proposal with Retry
		// Rotate proposers first
		proposerIndex := int(target.cid) % len(ctx.GenesisValidators)
		proposerKey = ctx.GenesisValidators[proposerIndex]
		
		var tx *types.Transaction
		
		// Retry loop
		for {
			proposerOpts, _ := ctx.GetTransactor(proposerKey)
			proposerOpts.Value = nil
			
			tx, err = ctx.Proposal.CreateUpdateConfigProposal(proposerOpts, big.NewInt(int64(target.cid)), targetVal)
			if err == nil {
				break // Success
			}
			
			// If failed, likely cooldown. Wait for a block.
			t.Logf("Proposal failed (%v), waiting for next block...", err)
			
			// Get current block
			header, _ := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
			currentHeight := header.Number.Uint64()
			
			// Wait until height increases
			for {
				time.Sleep(1 * time.Second)
				newHeader, _ := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
				if newHeader.Number.Uint64() > currentHeight {
					break
				}
			}
		}
		
		err = ctx.WaitMined(tx.Hash())
		utils.AssertNoError(t, err, "proposal tx failed")

		// 3. Get ID
		receipt, err := ctx.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
		utils.AssertNoError(t, err, "failed to get receipt")
		
		var proposalID [32]byte
		found := false
		for _, log := range receipt.Logs {
			event, err := ctx.Proposal.ParseLogCreateConfigProposal(*log)
			if err == nil {
				proposalID = event.Id
				found = true
				break
			}
		}
		utils.AssertTrue(t, found, "LogCreateConfigProposal not found")

		// 4. Vote
		for _, voterKey := range ctx.GenesisValidators {
			voterOpts, _ := ctx.GetTransactor(voterKey)
			txVote, err := ctx.Proposal.VoteProposal(voterOpts, proposalID, true)
			if err == nil {
				ctx.WaitMined(txVote.Hash())
			}
		}

		// 5. Verify Update
		// Wait a block for state update? Usually immediate after vote.
		time.Sleep(1 * time.Second)
		
		newVal, err := target.getter()
		utils.AssertNoError(t, err, "failed to get new value")
		utils.AssertBigIntEq(t, newVal, targetVal, "config update failed")
		t.Logf("%s updated successfully to %s", target.name, newVal)
	}
}

// Helper type for uint256 since I used it in struct definition
type uint256 = uint64
