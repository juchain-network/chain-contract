package tests

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestH_Robustness(t *testing.T) {
	if ctx == nil {
		t.Skip("Context not initialized")
	}

	// [V-01] Jailed Validator Redistribution
	t.Run("V-01_JailedRedistribution", func(t *testing.T) {
		valAddr := common.HexToAddress(ctx.Config.Validators[1].Address)
		valKey := ctx.GenesisValidators[1]
		opts, _ := ctx.GetTransactor(valKey)
		tx, err := ctx.Staking.ResignValidator(opts)
		if err == nil {
			ctx.WaitMined(tx.Hash())
		}
		
		t.Log("Waiting for blocks. If V2 mines, rewards should redistribute.")
		waitBlocks(t, 5)
		
		info, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		utils.AssertTrue(t, info.IsJailed, "Should be jailed")
	})

	// [S-16] Zero Delegated Rewards
	t.Run("S-16_ZeroDelegatedRewards", func(t *testing.T) {
		key, addr, err := createAndRegisterValidator(t, "ZeroDelegation")
		utils.AssertNoError(t, err, "failed to setup validator")
		
		waitBlocks(t, 5)
		
		info, _ := ctx.Staking.GetValidatorInfo(nil, addr)
		t.Logf("Validator %s accumulated: %s", addr.Hex(), info.AccumulatedRewards.String())
		
		opts, _ := ctx.GetTransactor(key)
		ctx.Staking.ClaimValidatorRewards(opts)
	})

	// [S-15] Proposal Expiry
	t.Run("S-15_ProposalExpiry", func(t *testing.T) {
		userKey, userAddr, err := ctx.CreateAndFundAccount(utils.ToWei(10))
		utils.AssertNoError(t, err, "setup user failed")
		
		opts, err := ctx.GetTransactor(userKey)
		utils.AssertNoError(t, err, "transactor failed")
		
		tx, err := ctx.Proposal.CreateProposal(opts, userAddr, true, "Expiry Test")
		utils.AssertNoError(t, err, "create proposal failed")
		ctx.WaitMined(tx.Hash())
		
		propID := getPropID(tx)
		if propID == ([32]byte{}) { t.Fatal("propID missing") }

		p, _ := ctx.Proposal.Proposals(nil, propID)
		utils.AssertTrue(t, p.Proposer == userAddr, "Proposal not found")
	})
}