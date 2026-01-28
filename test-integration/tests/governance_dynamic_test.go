package tests

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestB_Governance_Dynamic(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized")
	}

	proposerIndex := 0
	getProposer := func() *ecdsa.PrivateKey {
		k := ctx.GenesisValidators[proposerIndex%len(ctx.GenesisValidators)]
		proposerIndex++
		return k
	}

	// [G-08] Invalid Voting
	t.Run("G-08_InvalidVoting", func(t *testing.T) {
		// 1. Create a proposal first
		_, candAddr, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		var tx *types.Transaction
		var err error
		
		for {
			proposerKey := getProposer()
			opts, _ := ctx.GetTransactor(proposerKey)
			tx, err = ctx.Proposal.CreateProposal(opts, candAddr, true, "G-08 Invalid Vote")
			if err == nil { break }
			if strings.Contains(err.Error(), "Proposal creation too frequent") {
				time.Sleep(2 * time.Second)
				continue
			}
			t.Fatalf("create proposal failed: %v", err)
		}
		ctx.WaitMined(tx.Hash())
		
		receipt, _ := ctx.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
		var propID [32]byte
		for _, l := range receipt.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil { propID = ev.Id; break }
		}

		// Test Double Vote
		voteOpts, _ := ctx.GetTransactor(ctx.GenesisValidators[0])
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
		_, candAddr2, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		var tx2 *types.Transaction
		for {
			proposerKey := getProposer()
			opts, _ := ctx.GetTransactor(proposerKey)
			tx2, err = ctx.Proposal.CreateProposal(opts, candAddr2, true, "G-08 Expiry")
			if err == nil { break }
			if strings.Contains(err.Error(), "Proposal creation too frequent") {
				time.Sleep(2 * time.Second)
				continue
			}
			t.Fatalf("create expiry proposal failed: %v", err)
		}
		ctx.WaitMined(tx2.Hash())
		receipt2, _ := ctx.Clients[0].TransactionReceipt(context.Background(), tx2.Hash())
		var propID2 [32]byte
		for _, l := range receipt2.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil { propID2 = ev.Id; break }
		}
		period, _ := ctx.Proposal.ProposalLastingPeriod(nil)
		if period.Sign() > 0 {
			t.Logf("Waiting %s blocks for expiry...", period)
			waitBlocks(t, int(new(big.Int).Add(period, big.NewInt(1)).Int64()))
			_, err = ctx.Proposal.VoteProposal(voteOpts, propID2, true)
			if err == nil {
				t.Fatal("Vote on expired proposal should fail")
			}
		}
	})

	// [G-12] Last Man Standing (Removal Protection)
	t.Run("G-12_LastManStanding", func(t *testing.T) {
		t.Skip("Skipping G-12 to preserve validator set for other tests")
	})

	// [G-15] Dynamic Threshold
	t.Run("G-15_DynamicThreshold", func(t *testing.T) {
		// 1. Add V4 to increase threshold to 3
		_, v4Addr, _ := ctx.CreateAndFundAccount(utils.ToWei(100005))
		err := passProposalFor(t, v4Addr, "G-15 Add V4")
		utils.AssertNoError(t, err, "G-15 add proposal failed")
		
		v4KeyReal, v4AddrReal, _ := ctx.CreateAndFundAccount(utils.ToWei(100005))
		err = passProposalFor(t, v4AddrReal, "G-15 Add V4 Real")
		utils.AssertNoError(t, err, "G-15 real add failed")
		
		opts4, _ := ctx.GetTransactor(v4KeyReal)
		opts4.Value = utils.ToWei(100000)
		tx, err := ctx.Staking.RegisterValidator(opts4, big.NewInt(1000))
		utils.AssertNoError(t, err, "register v4 failed")
		ctx.WaitMined(tx.Hash())
		
		// 2. Create Proposal (e.g. Add V5)
		_, v5Addr, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		var txP *types.Transaction
		for {
			proposerKey := getProposer()
			opts, _ := ctx.GetTransactor(proposerKey)
			txP, err = ctx.Proposal.CreateProposal(opts, v5Addr, true, "G-15 Add V5")
			if err == nil { break }
			if strings.Contains(err.Error(), "Proposal creation too frequent") {
				time.Sleep(2 * time.Second)
				continue
			}
			t.Fatalf("proposal v5 failed: %v", err)
		}
		ctx.WaitMined(txP.Hash())
		
		receipt, _ := ctx.Clients[0].TransactionReceipt(context.Background(), txP.Hash())
		var propID [32]byte
		for _, l := range receipt.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil { propID = ev.Id; break }
		}
		
		// 3. Vote 2 times. Total Agree = 2. Threshold is 3.
		for i := 0; i < 2; i++ {
			vk := ctx.GenesisValidators[i]
			vo, _ := ctx.GetTransactor(vk)
			txV, _ := ctx.Proposal.VoteProposal(vo, propID, true)
			ctx.WaitMined(txV.Hash())
		}
		
		pass, _ := ctx.Proposal.Pass(nil, v5Addr)
		utils.AssertTrue(t, !pass, "Should not pass with 2/4 votes")
		
		// 4. Remove V4 (reduce validator count to 3).
		// Use proposerKey 0 to remove V4
		p0Key := ctx.GenesisValidators[0]
		p0Opts, _ := ctx.GetTransactor(p0Key)
		txR, _ := ctx.Proposal.CreateProposal(p0Opts, v4AddrReal, false, "G-15 Remove V4")
		ctx.WaitMined(txR.Hash())
		recR, _ := ctx.Clients[0].TransactionReceipt(context.Background(), txR.Hash())
		var pidR [32]byte
		for _, l := range recR.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil { pidR = ev.Id; break }
		}
		for _, vk := range ctx.GenesisValidators {
			vo, _ := ctx.GetTransactor(vk)
			ctx.Proposal.VoteProposal(vo, pidR, true)
		}
		time.Sleep(2 * time.Second)
		
		// 5. Trigger check by 3rd vote (V2)
		vk2 := ctx.GenesisValidators[2]
		vo2, _ := ctx.GetTransactor(vk2)
		txV2, _ := ctx.Proposal.VoteProposal(vo2, propID, true)
		ctx.WaitMined(txV2.Hash())
		
		passV5, _ := ctx.Proposal.Pass(nil, v5Addr)
		utils.AssertTrue(t, passV5, "V5 should pass after threshold reduction")
	})

	// [G-17] Proposal Nonce Isolation
	t.Run("G-17_NonceIsolation", func(t *testing.T) {
		target := common.HexToAddress("0xDEAD")
		
		// Proposer 1
		p1Key := getProposer()
		var tx1 *types.Transaction
		for {
			opts1, _ := ctx.GetTransactor(p1Key)
			var err error
			tx1, err = ctx.Proposal.CreateProposal(opts1, target, false, "Duplicate")
			if err == nil { break }
			if strings.Contains(err.Error(), "Proposal creation too frequent") {
				time.Sleep(2 * time.Second)
				continue
			}
			t.Fatalf("P1 proposal failed: %v", err)
		}
		ctx.WaitMined(tx1.Hash())
		
		// Proposer 2
		p2Key := getProposer()
		var tx2 *types.Transaction
		for {
			opts2, _ := ctx.GetTransactor(p2Key)
			var err error
			tx2, err = ctx.Proposal.CreateProposal(opts2, target, false, "Duplicate")
			if err == nil { break }
			if strings.Contains(err.Error(), "Proposal creation too frequent") {
				time.Sleep(2 * time.Second)
				continue
			}
			t.Fatalf("P2 proposal failed: %v", err)
		}
		ctx.WaitMined(tx2.Hash())
		
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
			t.Fatal("Proposal IDs should be unique")
		}
	})
}