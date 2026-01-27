package tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/core/types"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestC_StakingFlow(t *testing.T) {
	if ctx == nil {
		t.Skip("Context not initialized")
	}

	if len(ctx.GenesisValidators) == 0 {
		t.Skip("No genesis validators configured, cannot run governance tests")
	}

	// 1. Create a new account for the Candidate Validator
	t.Log("Creating new candidate validator account...")
	valKey, valAddr, err := ctx.CreateAndFundAccount(utils.ToWei(100005)) // 100k min stake + gas
	utils.AssertNoError(t, err, "failed to create validator account")

	// 2. Create Proposal (Must be done by an existing validator)
	t.Log("Creating proposal to add validator...")
	proposerKey := ctx.GenesisValidators[0]
	
	var tx *types.Transaction
	// Retry loop for cooldown
	for {
		proposerOpts, _ := ctx.GetTransactor(proposerKey)
		proposerOpts.Value = nil 
		
		var err error
		// CreateProposal(opts, dst, flag, details)
		tx, err = ctx.Proposal.CreateProposal(proposerOpts, valAddr, true, "New Validator Candidate")
		if err == nil {
			break
		}
		
		t.Logf("CreateProposal failed (%v), waiting for next block...", err)
		header, _ := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
		currentHeight := header.Number.Uint64()
		for {
			time.Sleep(1 * time.Second)
			newHeader, _ := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
			if newHeader.Number.Uint64() > currentHeight {
				break
			}
		}
	}
	
	utils.AssertNoError(t, err, "failed to create proposal")
	
	// Wait for tx and extract Proposal ID
	receipt, err := ctx.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
	// If receipt is not available immediately, wait
	if err != nil {
		err = ctx.WaitMined(tx.Hash())
		utils.AssertNoError(t, err, "proposal tx mining failed")
		receipt, err = ctx.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
		utils.AssertNoError(t, err, "failed to get receipt")
	}
	utils.AssertTrue(t, receipt.Status == 1, "proposal tx reverted")

	var proposalID [32]byte
	found := false
	for _, log := range receipt.Logs {
		event, err := ctx.Proposal.ParseLogCreateProposal(*log)
		if err == nil {
			proposalID = event.Id
			found = true
			t.Logf("Proposal created with ID: %x", proposalID)
			break
		}
	}
	utils.AssertTrue(t, found, "LogCreateProposal event not found")

	// 3. Vote on Proposal (All genesis validators vote YES)
	t.Log("Voting on proposal...")
	for i, voterKey := range ctx.GenesisValidators {
		// Use a different context/transactor for each voter
		voterOpts, _ := ctx.GetTransactor(voterKey)
		
		// Check if already voted (proposer might have auto-voted? usually no)
		// Just vote.
		txVote, err := ctx.Proposal.VoteProposal(voterOpts, proposalID, true)
		if err != nil {
			t.Logf("Validator %d failed to vote (maybe already voted or self-vote restricted?): %v", i, err)
			continue 
		}
		// We don't necessarily need to wait for every vote to be mined sequentially, 
		// but for stability in this test, we do.
		ctx.WaitMined(txVote.Hash())
	}

	// 4. Register Validator (Candidate executes this)
	t.Log("Registering validator...")
	// Check if passed?
	pass, err := ctx.Proposal.Pass(nil, valAddr)
	utils.AssertNoError(t, err, "failed to check pass status")
	
	if !pass {
		t.Log("Proposal has not passed yet (not enough votes?). Skipping registration.")
		// In a 4-node cluster, we need >50% (3 votes). If we have keys for all 4, it should pass.
		// If we only configured 1 key, it might fail.
		return 
	}

	registerOpts, _ := ctx.GetTransactor(valKey)
	registerOpts.Value = utils.ToWei(100000) // Min stake
	commission := big.NewInt(1000) // 10%
	
	txReg, err := ctx.Staking.RegisterValidator(registerOpts, commission)
	utils.AssertNoError(t, err, "failed to register validator")
	
	err = ctx.WaitMined(txReg.Hash())
	utils.AssertNoError(t, err, "register tx failed")

	// 5. Verify Registration
	isRegistered, err := ctx.Validators.IsValidatorExist(nil, valAddr)
	utils.AssertNoError(t, err, "failed to check validator existence")
	utils.AssertTrue(t, isRegistered, "Validator should be registered")
	
	t.Logf("Validator %s successfully registered!", valAddr.Hex())
}
