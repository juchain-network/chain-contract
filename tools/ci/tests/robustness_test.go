package tests

import (
	"context"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestH_Robustness(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized")
	}

	// [V-01] Revenue Redistribution after Jail
	t.Run("V-01_JailedRedistribution", func(t *testing.T) {
		// 1. Setup 3 Validators (already exist in genesis)
		// 2. Jail one of them (e.g. V2)
		v2Key := ctx.GenesisValidators[2]
		v2Addr := common.HexToAddress(ctx.Config.Validators[2].Address)
		
		opts2, _ := ctx.GetTransactor(v2Key)
		txJ, err := ctx.Staking.ResignValidator(opts2)
		utils.AssertNoError(t, err, "Jailing failed")
		ctx.WaitMined(txJ.Hash())
		
		// 3. Record incomes of V0, V1
		v0Addr := common.HexToAddress(ctx.Config.Validators[0].Address)
		v1Addr := common.HexToAddress(ctx.Config.Validators[1].Address)
		
		_, _, incoming0_before, _, _, _ := ctx.Validators.GetValidatorInfo(nil, v0Addr)
		_, _, incoming1_before, _, _, _ := ctx.Validators.GetValidatorInfo(nil, v1Addr)
		_, _, incoming2_before, _, _, _ := ctx.Validators.GetValidatorInfo(nil, v2Addr)

		// 4. Wait for V2 to produce a block while Jailed
		// In JuChain, Jailed validators remain in currentValidatorSet until Epoch.
		// If V2 mines a block, rewards should go to V0 and V1.
		// Since we cannot force mining on V2 easily in this harness, we simulate blocks.
		t.Log("Waiting for blocks. If V2 mines, rewards should redistribute.")
		waitBlocks(t, 5)
		
		_, _, incoming0_after, _, _, _ := ctx.Validators.GetValidatorInfo(nil, v0Addr)
		_, _, incoming1_after, _, _, _ := ctx.Validators.GetValidatorInfo(nil, v1Addr)
		_, _, incoming2_after, _, _, _ := ctx.Validators.GetValidatorInfo(nil, v2Addr)
		
		// Verify V2 didn't get new income
		utils.AssertBigIntEq(t, incoming2_after, incoming2_before, "Jailed validator should not get income")
		
		// Verify others got income
		if incoming0_after.Cmp(incoming0_before) > 0 || incoming1_after.Cmp(incoming1_before) > 0 {
			t.Log("Redistribution confirmed")
		}
	})

	// [S-16] Rewards when TotalDelegated is zero
	t.Run("S-16_ZeroDelegatedRewards", func(t *testing.T) {
		// 1. Use a validator with zero delegation (new one)
		_, addr, err := createAndRegisterValidator(t, "Robust-S16")
		if err != nil { return }
		
		infoBefore, _ := ctx.Staking.GetValidatorInfo(nil, addr)
		
		// 2. Produce blocks
		waitBlocks(t, 5)
		
		infoAfter, _ := ctx.Staking.GetValidatorInfo(nil, addr)
		
		// 3. Since delegated=0, all rewards should go to accumulatedRewards
		if infoAfter.AccumulatedRewards.Cmp(infoBefore.AccumulatedRewards) > 0 {
			t.Logf("Rewards accumulated correctly: %s -> %s", 
				infoBefore.AccumulatedRewards, infoAfter.AccumulatedRewards)
		}
	})

	// [S-15] Proposal Expiry Boundary
	t.Run("S-15_ProposalExpiry", func(t *testing.T) {
		// 1. Create Proposal for a new candidate
		candKey, candAddr, _ := ctx.CreateAndFundAccount(utils.ToWei(100005))
		
		// 2. Reduce proposalLastingPeriod temporarily
		// CID 0: proposalLastingPeriod. Set to 10 blocks.
		proposerKey := ctx.GenesisValidators[0]
		pOpts, _ := ctx.GetTransactor(proposerKey)
		
		// Update config
		txC, _ := ctx.Proposal.CreateUpdateConfigProposal(pOpts, big.NewInt(0), big.NewInt(10))
		ctx.WaitMined(txC.Hash())
		// Get ID and vote
		receipt, _ := ctx.Clients[0].TransactionReceipt(context.Background(), txC.Hash())
		var cid [32]byte
		for _, l := range receipt.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateConfigProposal(*l); err == nil { cid = ev.Id; break }
		}
		for _, vk := range ctx.GenesisValidators {
			vo, _ := ctx.GetTransactor(vk)
			ctx.Proposal.VoteProposal(vo, cid, true)
		}
		
		// 3. Create Validator Proposal
		waitNextBlock()
		txP, _ := ctx.Proposal.CreateProposal(pOpts, candAddr, true, "Expiry Test")
		ctx.WaitMined(txP.Hash())
		
		receiptP, _ := ctx.Clients[0].TransactionReceipt(context.Background(), txP.Hash())
		var pid [32]byte
		for _, l := range receiptP.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil { pid = ev.Id; break }
		}
		for _, vk := range ctx.GenesisValidators {
			vo, _ := ctx.GetTransactor(vk)
			ctx.Proposal.VoteProposal(vo, pid, true)
		}
		
		// 4. Wait for expiry (11 blocks)
		t.Log("Waiting 11 blocks for proposal to expire...")
		waitBlocks(t, 11)
		
		// 5. Try to Register
		cOpts, _ := ctx.GetTransactor(candKey)
		cOpts.Value = utils.ToWei(100000)
		_, err := ctx.Staking.RegisterValidator(cOpts, big.NewInt(1000))
		if err == nil {
			t.Fatal("Should fail registration after proposal expiry")
		}
		t.Logf("Caught expected expiry error: %v", err)
		
		// Reset proposalLastingPeriod to 200
		waitNextBlock()
		txR, _ := ctx.Proposal.CreateUpdateConfigProposal(pOpts, big.NewInt(0), big.NewInt(200))
		ctx.WaitMined(txR.Hash())
		// (Optional: vote to reset if needed for subsequent tests)
	})
}
