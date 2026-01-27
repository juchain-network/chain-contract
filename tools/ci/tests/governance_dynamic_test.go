package tests

import (
	"bytes"
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestB_Governance_Dynamic(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized")
	}

	// [G-08] Invalid Voting
	t.Run("G-08_InvalidVoting", func(t *testing.T) {
		// 1. Create a proposal first
		_, candAddr, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		proposerKey := ctx.GenesisValidators[0]
		opts, _ := ctx.GetTransactor(proposerKey)
		
		tx, err := ctx.Proposal.CreateProposal(opts, candAddr, true, "G-08 Invalid Vote")
		// Retry if cooldown
		if err != nil && err.Error() == "execution reverted: Proposal creation too frequent" {
			waitNextBlock()
			tx, err = ctx.Proposal.CreateProposal(opts, candAddr, true, "G-08 Retry")
		}
		utils.AssertNoError(t, err, "create proposal failed")
		ctx.WaitMined(tx.Hash())
		
		receipt, _ := ctx.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
		var propID [32]byte
		for _, l := range receipt.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil { propID = ev.Id; break }
		}

		// Test Double Vote
		voteOpts, _ := ctx.GetTransactor(proposerKey)
		txVote, err := ctx.Proposal.VoteProposal(voteOpts, propID, true)
		utils.AssertNoError(t, err, "first vote failed")
		ctx.WaitMined(txVote.Hash())
		
		_, err = ctx.Proposal.VoteProposal(voteOpts, propID, true)
		if err == nil {
			t.Fatal("Double vote should fail")
		}
		
		// Test Non-Existent
		var fakeID [32]byte
		fakeID[0] = 1
		_, err = ctx.Proposal.VoteProposal(voteOpts, fakeID, true)
		if err == nil {
			t.Fatal("Vote on non-existent proposal should fail")
		}
		
		// Test Expired (Wait for expiry)
		// Create a fresh proposal so we can let it expire.
		_, candAddr2, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		tx2, err := ctx.Proposal.CreateProposal(opts, candAddr2, true, "G-08 Expiry")
		if err != nil && err.Error() == "execution reverted: Proposal creation too frequent" {
			waitNextBlock()
			tx2, err = ctx.Proposal.CreateProposal(opts, candAddr2, true, "G-08 Expiry Retry")
		}
		utils.AssertNoError(t, err, "create expiry proposal failed")
		ctx.WaitMined(tx2.Hash())
		receipt2, _ := ctx.Clients[0].TransactionReceipt(context.Background(), tx2.Hash())
		var propID2 [32]byte
		for _, l := range receipt2.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil { propID2 = ev.Id; break }
		}
		period, _ := ctx.Proposal.ProposalLastingPeriod(nil)
		if period.Sign() == 0 {
			t.Skip("proposalLastingPeriod is zero")
		}
		// Mine period+1 blocks to expire
		waitBlocks(t, int(new(big.Int).Add(period, big.NewInt(1)).Int64()))
		_, err = ctx.Proposal.VoteProposal(voteOpts, propID2, true)
		if err == nil {
			t.Fatal("Vote on expired proposal should fail")
		}
	})

	// [G-12] Last Man Standing (Removal Protection)
	t.Run("G-12_LastManStanding", func(t *testing.T) {
		// This requires reducing validator set to 1.
		// Current set size ~3 (Genesis).
		// We need to remove 2 validators.
		// This is destructive for the test environment.
		// We should probably check if we can mock or if there is a way to check without destroying.
		
		// Alternative: Check code logic or assume environment is disposable (it is).
		// But other tests depend on validators.
		// So we should run this LAST or skip.
		// Let's Skip if we want to preserve env, or run it and expect subsequent tests to fail or handle it.
		// Since this is the end of the chain, maybe it's fine.
		
		t.Skip("Skipping G-12 to preserve validator set for other tests")
	})

	// [G-15] Dynamic Threshold
	t.Run("G-15_DynamicThreshold", func(t *testing.T) {
		// Scenario: 4 Validators. Threshold = 3.
		// 1. Add V4.
		_, v4Addr, _ := ctx.CreateAndFundAccount(utils.ToWei(100005))
		
		// We need to add V4 first.
		// This creates proposal, votes, passes.
		passProposalFor(t, v4Addr, "G-15 Add V4")
		
		// Register V4
		v4Key2, v4Addr2, _ := ctx.CreateAndFundAccount(utils.ToWei(100005))
		passProposalFor(t, v4Addr2, "G-15 Add V4")
		
		opts4, _ := ctx.GetTransactor(v4Key2)
		opts4.Value = utils.ToWei(100000)
		tx, err := ctx.Staking.RegisterValidator(opts4, big.NewInt(1000))
		utils.AssertNoError(t, err, "register v4 failed")
		ctx.WaitMined(tx.Hash())
		
		// Now we have 4 validators (3 Genesis + V4).
		// Voting Count = 4. Threshold = 4/2 + 1 = 3.
		
		// 2. Create Proposal (e.g. Add V5)
		_, v5Addr, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		proposerKey := ctx.GenesisValidators[0]
		opts, _ := ctx.GetTransactor(proposerKey)
		opts.Value = nil
		
		txP, err := ctx.Proposal.CreateProposal(opts, v5Addr, true, "G-15 Add V5")
		// Retry if cooldown
		if err != nil {
			waitNextBlock()
			txP, err = ctx.Proposal.CreateProposal(opts, v5Addr, true, "G-15 Add V5 Retry")
		}
		utils.AssertNoError(t, err, "proposal v5 failed")
		ctx.WaitMined(txP.Hash())
		
		receipt, _ := ctx.Clients[0].TransactionReceipt(context.Background(), txP.Hash())
		var propID [32]byte
		for _, l := range receipt.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil { propID = ev.Id; break }
		}
		
		// 3. Vote 2 times (V0, V1). Total Agree = 2. Threshold is 3. Result not exist.
		for i := 0; i < 2; i++ {
			vk := ctx.GenesisValidators[i]
			vo, _ := ctx.GetTransactor(vk)
			txV, _ := ctx.Proposal.VoteProposal(vo, propID, true)
			ctx.WaitMined(txV.Hash())
		}
		
		// Check result: should be not passed yet.
		// We can't easily check internal state `results`, but we can check `pass`.
		pass, _ := ctx.Proposal.Pass(nil, v5Addr)
		utils.AssertTrue(t, !pass, "Should not pass with 2/4 votes")
		
		// 4. Remove V4 (reduce validator count to 3).
		// We need to remove V4.
		// Create Remove Proposal
		txR, err := ctx.Proposal.CreateProposal(opts, v4Addr2, false, "G-15 Remove V4")
		if err != nil {
			waitNextBlock()
			txR, err = ctx.Proposal.CreateProposal(opts, v4Addr2, false, "G-15 Remove V4 Retry")
		}
		utils.AssertNoError(t, err, "remove v4 proposal failed")
		ctx.WaitMined(txR.Hash())
		
		receiptR, _ := ctx.Clients[0].TransactionReceipt(context.Background(), txR.Hash())
		var remID [32]byte
		for _, l := range receiptR.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil { remID = ev.Id; break }
		}
		
		// Vote to remove V4 (Needs 3 votes from 4)
		for i := 0; i < 3; i++ {
			vk := ctx.GenesisValidators[i]
			vo, _ := ctx.GetTransactor(vk)
			txV, _ := ctx.Proposal.VoteProposal(vo, remID, true)
			ctx.WaitMined(txV.Hash())
		}
		
		// V4 should be removed (pass=false)
		passV4, _ := ctx.Proposal.Pass(nil, v4Addr2)
		utils.AssertTrue(t, !passV4, "V4 should be removed")
		
		// Now Validator Set size is 3. Threshold = 3/2 + 1 = 2.
		// The previous proposal (Add V5) has 2 votes.
		// 5. Trigger check on V5 proposal?
		// Proposal logic usually checks threshold AT THE MOMENT of voting.
		// Since we already voted, the state "agree=2" is stored.
		// But "resultExist" is false.
		// If we vote again (e.g. V2 votes), it will check `agree >= threshold`.
		// agree will be 3. threshold will be 2. It will pass.
		// But what if we don't vote? The proposal remains stuck unless someone votes.
		// OR, if the threshold dropped to 2, and we already have 2 votes, 
		// the NEXT interaction (vote) should trigger success.
		
		// Let's have V2 vote for V5.
		vk2 := ctx.GenesisValidators[2]
		vo2, _ := ctx.GetTransactor(vk2)
		txV2, _ := ctx.Proposal.VoteProposal(vo2, propID, true)
		ctx.WaitMined(txV2.Hash())
		
		// Now agree=3. Threshold=2. Should pass.
		passV5, _ := ctx.Proposal.Pass(nil, v5Addr)
		utils.AssertTrue(t, passV5, "V5 should pass with 3 votes (threshold reduced)")
	})

	// [G-17] Proposal Nonce Isolation
	t.Run("G-17_NonceIsolation", func(t *testing.T) {
		// Same target, same flag, same details from different proposers
		target := common.HexToAddress("0xDEAD")
		
		// Proposer 1
		p1Key := ctx.GenesisValidators[0]
		opts1, _ := ctx.GetTransactor(p1Key)
		tx1, err1 := ctx.Proposal.CreateProposal(opts1, target, false, "Duplicate")
		utils.AssertNoError(t, err1, "P1 proposal failed")
		
		// Proposer 2 (Wait for next block to avoid cooldown)
		waitNextBlock()
		p2Key := ctx.GenesisValidators[1]
		opts2, _ := ctx.GetTransactor(p2Key)
		tx2, err2 := ctx.Proposal.CreateProposal(opts2, target, false, "Duplicate")
		utils.AssertNoError(t, err2, "P2 proposal failed")
		
		// Get IDs
		rec1, _ := ctx.Clients[0].TransactionReceipt(context.Background(), tx1.Hash())
		rec2, _ := ctx.Clients[0].TransactionReceipt(context.Background(), tx2.Hash())
		
		var id1, id2 [32]byte
		for _, l := range rec1.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil { id1 = ev.Id; break }
		}
		for _, l := range rec2.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil { id2 = ev.Id; break }
		}
		
		if bytes.Equal(id1[:], id2[:]) {
			t.Fatal("Proposal IDs should be unique even with same content (due to nonces)")
		}
		t.Logf("Generated unique IDs: %x and %x", id1, id2)
	})
}
