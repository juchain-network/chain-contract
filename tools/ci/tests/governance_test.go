package tests

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestB_Governance(t *testing.T) {
	if ctx == nil {
		t.Skip("Context not initialized")
	}
	if len(ctx.GenesisValidators) == 0 {
		t.Skip("No genesis validators configured")
	}

	// Helper to create proposal with retry
	createProposalWithRetry := func(opts *bind.TransactOpts, dst common.Address, flag bool, desc string) (*types.Transaction, error) {
		var tx *types.Transaction
		var err error
		for {
			tx, err = ctx.Proposal.CreateProposal(opts, dst, flag, desc)
			if err == nil {
				return tx, nil
			}
			if err.Error() == "execution reverted: Proposal creation too frequent" {
				t.Logf("CreateProposal hit cooldown, waiting for next block...")
				waitNextBlock()
				continue
			}
			return nil, err
		}
	}

	// Helper to create and pass a proposal
	createAndPassProposal := func(dst common.Address, flag bool, desc string) error {
		proposerKey := ctx.GenesisValidators[0]
		proposerOpts, _ := ctx.GetTransactor(proposerKey)
		proposerOpts.Value = nil

		tx, err := createProposalWithRetry(proposerOpts, dst, flag, desc)
		if err != nil { return err }
		ctx.WaitMined(tx.Hash())

		receipt, _ := ctx.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
		var proposalID [32]byte
		for _, log := range receipt.Logs {
			event, err := ctx.Proposal.ParseLogCreateProposal(*log)
			if err == nil { proposalID = event.Id; break }
		}

		// Vote (Need majority)
		for i, voterKey := range ctx.GenesisValidators {
			voterOpts, _ := ctx.GetTransactor(voterKey)
			txVote, err := ctx.Proposal.VoteProposal(voterOpts, proposalID, true)
			if err == nil {
				ctx.WaitMined(txVote.Hash())
				t.Logf("Validator %d voted YES", i)
			}
		}
		
		// Wait a bit for state sync
		time.Sleep(2 * time.Second)
		
		pass, err := ctx.Proposal.Pass(nil, dst)
		if err != nil { return err }
		if flag && !pass { return fmt.Errorf("proposal passed but dst status not updated (expected true)") }
		if !flag && pass { return fmt.Errorf("proposal passed but dst status not updated (expected false)") }
		return nil
	}

	_, candidateAddr, err := ctx.CreateAndFundAccount(utils.ToWei(1))
	utils.AssertNoError(t, err, "create candidate failed")

	t.Run("G-01_AddValidator", func(t *testing.T) {
		err := createAndPassProposal(candidateAddr, true, "G-01 Test Add")
		utils.AssertNoError(t, err, "add validator proposal failed")
	})

	t.Run("G-02_RemoveValidator", func(t *testing.T) {
		err := createAndPassProposal(candidateAddr, false, "G-02 Test Remove")
		utils.AssertNoError(t, err, "remove validator proposal failed")
	})

	t.Run("G-03_ReOnboard", func(t *testing.T) {
		err := createAndPassProposal(candidateAddr, true, "G-03 Test Revive")
		utils.AssertNoError(t, err, "revive proposal failed")
	})

	t.Run("G-13_FlipFlop", func(t *testing.T) {
		err := createAndPassProposal(candidateAddr, false, "G-13 Remove")
		utils.AssertNoError(t, err, "G-13 remove failed")
		err = createAndPassProposal(candidateAddr, true, "G-13 Add")
		utils.AssertNoError(t, err, "G-13 add failed")
	})

	t.Run("G-11_GhostRemoval", func(t *testing.T) {
		randomAddr := common.HexToAddress("0x1234567890123456789012345678901234567890")
		err := createAndPassProposal(randomAddr, false, "G-11 Ghost")
		utils.AssertNoError(t, err, "ghost removal failed")
	})

	t.Run("G-06_DuplicateProposal", func(t *testing.T) {
		// Use candidateAddr which IS passed (from G-13)
		proposerKey := ctx.GenesisValidators[0]
		opts, _ := ctx.GetTransactor(proposerKey)
		
		// Should fail because already passed
		_, err := ctx.Proposal.CreateProposal(opts, candidateAddr, true, "G-06 Should Fail")
		if err == nil {
			t.Fatal("Expected failure for already passed dst, got success")
		}
		t.Log("Duplicate/Invalid add rejected correctly:", err)
	})

	t.Run("G-05_Cooldown", func(t *testing.T) {
		if len(ctx.GenesisValidators) < 2 { t.Skip("Need 2 validators") }
		
		proposerKey := ctx.GenesisValidators[1]
		opts, _ := ctx.GetTransactor(proposerKey)
		
		// 1. Send first (don't care if it was already in cooldown, we just need one success or fail)
		tx, err := ctx.Proposal.CreateProposal(opts, common.HexToAddress("0x9999"), false, "G-05 1")
		if err != nil {
			if err.Error() == "execution reverted: Proposal creation too frequent" {
				t.Log("Already in cooldown, test condition met")
				return
			}
			t.Fatalf("Unexpected error: %v", err)
		}
		ctx.WaitMined(tx.Hash())
		
		// 2. Immediate second call should fail
		_, err2 := ctx.Proposal.CreateProposal(opts, common.HexToAddress("0x8888"), false, "G-05 2")
		if err2 == nil {
			t.Fatal("Expected cooldown error, got nil")
		}
		t.Log("Cooldown triggered correctly")
	})
	
	t.Run("G-07_FrontRunning", func(t *testing.T) {
		frKey, frontRunner, _ := ctx.CreateAndFundAccount(utils.ToWei(100005))
		proposerKey := ctx.GenesisValidators[0]
		opts, _ := ctx.GetTransactor(proposerKey)
		opts.Value = nil
		
		tx, err := createProposalWithRetry(opts, frontRunner, true, "G-07 FrontRun")
		utils.AssertNoError(t, err, "front-run proposal failed")
		ctx.WaitMined(tx.Hash())
		
		regOpts, _ := ctx.GetTransactor(frKey)
		regOpts.Value = utils.ToWei(100000)
		
		_, err = ctx.Staking.RegisterValidator(regOpts, big.NewInt(1000))
		if err == nil {
			t.Fatal("Expected register failure (front-running), got success")
		}
	})

	// [V-02] Description Boundary Validation
	t.Run("V-02_DescriptionBoundary", func(t *testing.T) {
		proposerKey := ctx.GenesisValidators[0]
		opts, _ := ctx.GetTransactor(proposerKey)
		
		valAddr := common.HexToAddress(ctx.Config.Validators[0].Address)
		
		// Moniker > 70 bytes
		longMoniker := ""
		for i := 0; i < 71; i++ { longMoniker += "a" }
		
		_, err := ctx.Validators.CreateOrEditValidator(opts, valAddr, longMoniker, "", "", "", "")
		if err == nil {
			t.Fatal("Should fail with moniker > 70 bytes")
		}
		t.Logf("Caught expected error: %v", err)
	})

	// [G-16] Smooth Expansion (Rate Limiting)
	t.Run("G-16_SmoothExpansion", func(t *testing.T) {
		// 1. Get current active count
		currentSet, _ := ctx.Validators.GetActiveValidators(nil)
		initialCount := len(currentSet)
		
		// 2. Register 2 new validators (V1, V2)
		// We need to fully onboard them (Propose -> Vote -> Register)
		// Assuming we can do this fast enough within an epoch (or simulated)
		
		// V1
		_, v1Addr, err := ctx.CreateAndFundAccount(utils.ToWei(100005))
		utils.AssertNoError(t, err, "create v1 failed")
		err = createAndPassProposal(v1Addr, true, "G-16 V1")
		utils.AssertNoError(t, err, "v1 proposal failed")
		
		// Register V1
		// We need key for V1? Ah, CreateAndFundAccount returns key.
		// Wait, CreateAndFundAccount returns (key, addr, err)
		// But in this helper block I ignored key.
		// I need to refactor to get key or just re-fund if needed?
		// Actually, I can just use a helper that returns key.
		
		// Let's copy-paste account creation to get key
		v1Key, _, _ := ctx.CreateAndFundAccount(utils.ToWei(100005)) // Use new account for V1 actually
		// The previous line was just funding 'v1Addr' but discarding key.
		// Let's fix this properly.
		v1Addr = crypto.PubkeyToAddress(v1Key.PublicKey)
		
		// Re-do proposal for correct address
		err = createAndPassProposal(v1Addr, true, "G-16 V1 Real")
		utils.AssertNoError(t, err, "v1 real proposal failed")
		
		v1Opts, _ := ctx.GetTransactor(v1Key)
		v1Opts.Value = utils.ToWei(100000)
		tx1, err := ctx.Staking.RegisterValidator(v1Opts, big.NewInt(1000))
		utils.AssertNoError(t, err, "v1 register failed")
		ctx.WaitMined(tx1.Hash())
		
		// V2 (Should FAIL registration in same epoch)
		v2Key, v2Addr, _ := ctx.CreateAndFundAccount(utils.ToWei(100005))
		err = createAndPassProposal(v2Addr, true, "G-16 V2")
		utils.AssertNoError(t, err, "v2 proposal failed")
		
		v2Opts, _ := ctx.GetTransactor(v2Key)
		v2Opts.Value = utils.ToWei(100000)
		_, err = ctx.Staking.RegisterValidator(v2Opts, big.NewInt(1000))
		if err == nil {
			t.Fatal("Expected v2 register to fail due to epoch limit, but it succeeded")
		}
		t.Log("V2 registration correctly blocked in same epoch:", err)

		// 3. Check getTopValidators
		topValidators, err := ctx.Validators.GetTopValidators(nil)
		utils.AssertNoError(t, err, "getTopValidators failed")
		
		expectedLen := initialCount + 1
		if len(topValidators) != expectedLen {
			t.Fatalf("Smooth expansion failed: expected %d validators, got %d", expectedLen, len(topValidators))
		}
		t.Logf("Smooth expansion verified: %d -> %d (capped by registration limit)", initialCount, len(topValidators))
	})
}

func waitNextBlock() {
	header, _ := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
	current := header.Number.Uint64()
	for {
		time.Sleep(1 * time.Second)
		newH, _ := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
		if newH.Number.Uint64() > current { break }
	}
}
