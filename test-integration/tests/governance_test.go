package tests

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

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

	// Proposer rotation counter
	proposerIndex := 0

	// Helper to extract proposal ID
	getPropID := func(tx *types.Transaction) [32]byte {
		receipt, err := ctx.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
		if err != nil { return [32]byte{} }
		for _, log := range receipt.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*log); err == nil { return ev.Id }
			if ev, err := ctx.Proposal.ParseLogCreateConfigProposal(*log); err == nil { return ev.Id }
		}
		return [32]byte{}
	}

	// Helper to find an active validator proposer
	getActiveProposer := func() *ecdsa.PrivateKey {
		for i := 0; i < len(ctx.GenesisValidators)*2; i++ {
			k := ctx.GenesisValidators[proposerIndex%len(ctx.GenesisValidators)]
			proposerIndex++
			addr := crypto.PubkeyToAddress(k.PublicKey)
			active, _ := ctx.Validators.IsValidatorActive(nil, addr)
			if active {
				return k
			}
		}
		return ctx.GenesisValidators[0]
	}

	// Setup: Ensure stable config for this test group
	t.Run("Setup_Governance", func(t *testing.T) {
		updateConfig := func(cid uint256, val int64, name string) {
			var tx *types.Transaction
			var err error
			for {
				pk := getActiveProposer()
				opts, _ := ctx.GetTransactor(pk)
				tx, err = ctx.Proposal.CreateUpdateConfigProposal(opts, big.NewInt(int64(cid)), big.NewInt(val))
				if err == nil { break }
				if strings.Contains(err.Error(), "Proposal creation too frequent") {
					time.Sleep(1 * time.Second)
					continue
				}
				t.Fatalf("setup config %s failed: %v", name, err)
			}
			ctx.WaitMined(tx.Hash())
			propID := getPropID(tx)
			for _, vk := range ctx.GenesisValidators {
				vo, _ := ctx.GetTransactor(vk)
				ctx.Proposal.VoteProposal(vo, propID, true)
			}
			ctx.WaitMined(tx.Hash())
			time.Sleep(2 * time.Second)
		}
		updateConfig(0, 1000, "ProposalLastingPeriod")
		updateConfig(19, 1, "ProposalCooldown")
		
		t.Log("Waiting for fresh epoch (20 blocks)...")
		waitBlocks(t, 20)
	})

	// Helper to create and pass a proposal
	createAndPassProposal := func(dst common.Address, flag bool, desc string) error {
		proposerKey := getActiveProposer()
		var tx *types.Transaction
		var err error
		
		for {
			proposerOpts, _ := ctx.GetTransactor(proposerKey)
			proposerOpts.Value = nil
			tx, err = ctx.Proposal.CreateProposal(proposerOpts, dst, flag, desc)
			if err == nil { break }
			if strings.Contains(err.Error(), "Proposal creation too frequent") {
				time.Sleep(2 * time.Second)
				continue
			}
			if strings.Contains(err.Error(), "Validator only") {
				proposerKey = getActiveProposer()
				continue
			}
			return fmt.Errorf("createProposal failed: %w", err)
		}
		ctx.WaitMined(tx.Hash())
		proposalID := getPropID(tx)

		agreeCount := 0
		for _, voterKey := range ctx.GenesisValidators {
			voterAddr := crypto.PubkeyToAddress(voterKey.PublicKey)
			active, _ := ctx.Validators.IsValidatorActive(nil, voterAddr)
			if !active { continue }

			voterOpts, _ := ctx.GetTransactor(voterKey)
			var txVote *types.Transaction
			for retry := 0; retry < 5; retry++ {
				txVote, err = ctx.Proposal.VoteProposal(voterOpts, proposalID, true)
				if err == nil { break }
				if strings.Contains(err.Error(), "Epoch block forbidden") {
					t.Logf("Validator %s vote hit epoch block, waiting 1s...", voterAddr.Hex())
					time.Sleep(1 * time.Second)
					continue
				}
				break
			}
			if err == nil {
				ctx.WaitMined(txVote.Hash())
				agreeCount++
			}
		}
		
		votingCount, _ := ctx.Validators.GetVotingValidatorCount(nil)
		threshold := votingCount.Uint64()/2 + 1
		t.Logf("Proposal %x status: %d votes received, threshold required: %d", proposalID, agreeCount, threshold)
		
		time.Sleep(1 * time.Second)
		pass, _ := ctx.Proposal.Pass(nil, dst)
		if flag && !pass { return fmt.Errorf("proposal should be passed") }
		if !flag && pass { return fmt.Errorf("proposal should be removed") }
		return nil
	}

	t.Run("G-01_AddValidator", func(t *testing.T) {
		_, addr, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		err := createAndPassProposal(addr, true, "G-01 Add")
		utils.AssertNoError(t, err, "add validator proposal failed")
	})

	t.Run("G-02_RemoveValidator", func(t *testing.T) {
		_, addr, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		createAndPassProposal(addr, true, "G-02 Prep") // Pass it first
		err := createAndPassProposal(addr, false, "G-02 Remove")
		utils.AssertNoError(t, err, "remove validator proposal failed")
	})

	t.Run("G-03_ReOnboard", func(t *testing.T) {
		_, addr, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		err := createAndPassProposal(addr, true, "G-03 Add")
		utils.AssertNoError(t, err, "revive proposal failed")
	})

	t.Run("G-13_FlipFlop", func(t *testing.T) {
		_, addr, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		err := createAndPassProposal(addr, true, "G-13 Add")
		utils.AssertNoError(t, err, "G-13 add failed")
		err = createAndPassProposal(addr, false, "G-13 Remove")
		utils.AssertNoError(t, err, "G-13 remove failed")
	})

	t.Run("G-11_GhostRemoval", func(t *testing.T) {
		randomAddr := common.HexToAddress("0x1234567890123456789012345678901234567890")
		err := createAndPassProposal(randomAddr, false, "G-11 Ghost")
		utils.AssertNoError(t, err, "ghost removal failed")
	})

	t.Run("G-06_DuplicateProposal", func(t *testing.T) {
		_, addr, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		createAndPassProposal(addr, true, "G-06 Prep")
		
		proposerKey := getActiveProposer()
		opts, _ := ctx.GetTransactor(proposerKey)
		_, err := ctx.Proposal.CreateProposal(opts, addr, true, "G-06 Should Fail")
		if err == nil {
			t.Fatal("Expected failure for already passed dst, got success")
		}
		t.Log("Duplicate add rejected correctly:", err)
	})

	t.Run("G-05_Cooldown", func(t *testing.T) {
		adminKey := getActiveProposer()
		adminOpts, _ := ctx.GetTransactor(adminKey)
		txC, _ := ctx.Proposal.CreateUpdateConfigProposal(adminOpts, big.NewInt(19), big.NewInt(5))
		ctx.WaitMined(txC.Hash())
		propID := getPropID(txC)
		for _, k := range ctx.GenesisValidators {
			voterAddr := crypto.PubkeyToAddress(k.PublicKey)
			active, _ := ctx.Validators.IsValidatorActive(nil, voterAddr)
			if !active { continue }
			vo, _ := ctx.GetTransactor(k)
			ctx.Proposal.VoteProposal(vo, propID, true)
		}
		time.Sleep(2 * time.Second)

		proposerKey := getActiveProposer()
		opts, _ := ctx.GetTransactor(proposerKey)
		tx, err := ctx.Proposal.CreateProposal(opts, common.HexToAddress("0x9999"), false, "G-05 1")
		utils.AssertNoError(t, err, "first proposal failed")
		ctx.WaitMined(tx.Hash())
		
		_, err2 := ctx.Proposal.CreateProposal(opts, common.HexToAddress("0x8888"), false, "G-05 2")
		if err2 == nil { t.Fatal("Expected cooldown error") }
		t.Log("Cooldown triggered correctly:", err2)

		t.Log("Waiting for cooldown to expire...")
		waitBlocks(t, 6)
		
		txR, _ := ctx.Proposal.CreateUpdateConfigProposal(adminOpts, big.NewInt(19), big.NewInt(1))
		ctx.WaitMined(txR.Hash())
		propIDR := getPropID(txR)
		for _, k := range ctx.GenesisValidators {
			voterAddr := crypto.PubkeyToAddress(k.PublicKey)
			active, _ := ctx.Validators.IsValidatorActive(nil, voterAddr)
			if !active { continue }
			vo, _ := ctx.GetTransactor(k)
			ctx.Proposal.VoteProposal(vo, propIDR, true)
		}
	})
	
	t.Run("G-07_FrontRunning", func(t *testing.T) {
		fakeValKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(100005))
		regOpts, _ := ctx.GetTransactor(fakeValKey)
		regOpts.Value = utils.ToWei(100000)
		_, err := ctx.Staking.RegisterValidator(regOpts, big.NewInt(1000))
		if err == nil { t.Fatal("Expected register failure (no proposal)") }
		t.Log("Registration without proposal correctly blocked:", err)
	})

	t.Run("V-02_DescriptionBoundary", func(t *testing.T) {
		proposerKey := getActiveProposer()
		opts, _ := ctx.GetTransactor(proposerKey)
		valAddr := crypto.PubkeyToAddress(proposerKey.PublicKey)
		longMoniker := strings.Repeat("a", 71)
		_, err := ctx.Validators.CreateOrEditValidator(opts, valAddr, longMoniker, "", "", "", "")
		if err == nil { t.Fatal("Should fail with moniker > 70 bytes") }
		t.Logf("Caught expected error: %v", err)
	})

	t.Run("G-16_SmoothExpansion", func(t *testing.T) {
		currentSet, _ := ctx.Validators.GetActiveValidators(nil)
		initialCount := len(currentSet)
		
		v1Key, v1Addr, _ := ctx.CreateAndFundAccount(utils.ToWei(100005))
		createAndPassProposal(v1Addr, true, "G-16 V1")
		
		v1Opts, _ := ctx.GetTransactor(v1Key)
		v1Opts.Value = utils.ToWei(100000)
		tx1, err := ctx.Staking.RegisterValidator(v1Opts, big.NewInt(1000))
		utils.AssertNoError(t, err, "v1 register failed")
		ctx.WaitMined(tx1.Hash())
		
		v2Key, v2Addr, _ := ctx.CreateAndFundAccount(utils.ToWei(100005))
		createAndPassProposal(v2Addr, true, "G-16 V2")
		
		v2Opts, _ := ctx.GetTransactor(v2Key)
		v2Opts.Value = utils.ToWei(100000)
		_, err = ctx.Staking.RegisterValidator(v2Opts, big.NewInt(1000))
		if err != nil {
			t.Log("V2 registration correctly blocked in same epoch:", err)
		} else {
			t.Log("Warning: V2 register succeeded unexpectedly.")
		}

		topValidators, _ := ctx.Validators.GetTopValidators(nil)
		t.Logf("Smooth expansion check: initial=%d current=%d", initialCount, len(topValidators))
	})
}

func waitNextBlock() {
	waitBlocks(nil, 1)
}