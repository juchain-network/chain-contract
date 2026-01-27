package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"juchain.org/chain/tools/ci/internal/utils"
)

// TestZ_LastManStanding ensures the last validator cannot be removed (protection path).
// This is destructive; keep it at the end of the suite.
func TestZ_LastManStanding(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) < 3 {
		t.Skip("Need at least 3 genesis validators")
	}
	highest, _ := ctx.Validators.GetHighestValidators(nil)
	if len(highest) <= 1 {
		t.Skip("validator set already reduced to 1")
	}

	removeByProposal := func(target common.Address) error {
		proposerKey := ctx.GenesisValidators[0]
		opts, _ := ctx.GetTransactor(proposerKey)
		opts.Value = nil

		tx, err := ctx.Proposal.CreateProposal(opts, target, false, fmt.Sprintf("G-12 Remove %s", target.Hex()))
		if err != nil && err.Error() == "execution reverted: Proposal creation too frequent" {
			waitNextBlock()
			tx, err = ctx.Proposal.CreateProposal(opts, target, false, fmt.Sprintf("G-12 Remove %s Retry", target.Hex()))
		}
		if err != nil {
			return err
		}
		ctx.WaitMined(tx.Hash())

		receipt, _ := ctx.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
		var propID [32]byte
		for _, l := range receipt.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil {
				propID = ev.Id
				break
			}
		}
		for _, vk := range ctx.GenesisValidators {
			vo, _ := ctx.GetTransactor(vk)
			ctx.Proposal.VoteProposal(vo, propID, true)
		}
		waitNextBlock()

		pass, _ := ctx.Proposal.Pass(nil, target)
		if pass {
			return fmt.Errorf("expected pass=false for removed validator")
		}
		return nil
	}

	// Remove two validators to leave only one.
	v1 := common.HexToAddress(ctx.Config.Validators[1].Address)
	v2 := common.HexToAddress(ctx.Config.Validators[2].Address)
	utils.AssertNoError(t, removeByProposal(v1), "remove v1 failed")
	utils.AssertNoError(t, removeByProposal(v2), "remove v2 failed")

	// Now attempt to remove the last remaining validator.
	last := common.HexToAddress(ctx.Config.Validators[0].Address)
	proposerKey := ctx.GenesisValidators[0]
	opts, _ := ctx.GetTransactor(proposerKey)
	tx, err := ctx.Proposal.CreateProposal(opts, last, false, "G-12 Last Man")
	if err != nil && err.Error() == "execution reverted: Proposal creation too frequent" {
		waitNextBlock()
		tx, err = ctx.Proposal.CreateProposal(opts, last, false, "G-12 Last Man Retry")
	}
	utils.AssertNoError(t, err, "last man proposal failed")
	ctx.WaitMined(tx.Hash())
	receipt, _ := ctx.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
	var propID [32]byte
	for _, l := range receipt.Logs {
		if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil {
			propID = ev.Id
			break
		}
	}
	for _, vk := range ctx.GenesisValidators {
		vo, _ := ctx.GetTransactor(vk)
		ctx.Proposal.VoteProposal(vo, propID, true)
	}
	waitNextBlock()

	// Protection: last validator should remain in highest set and pass status unchanged (still true).
	pass, _ := ctx.Proposal.Pass(nil, last)
	utils.AssertTrue(t, pass, "last validator should remain passed")
	highest, _ = ctx.Validators.GetHighestValidators(nil)
	if len(highest) != 1 {
		t.Fatalf("expected highestValidatorsSet length = 1, got %d", len(highest))
	}
}
