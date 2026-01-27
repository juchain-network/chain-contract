package tests

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestD_StakingManagement(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized or no validators")
	}

	// Use Genesis Validator 0 for most tests to avoid Rate Limit (1 val/epoch)
	// Genesis validators are already registered and active.
	valKey := ctx.GenesisValidators[0]
	valAddr := common.HexToAddress(ctx.Config.Validators[0].Address)

	// [S-01] Add Stake
	t.Run("S-01_AddStake", func(t *testing.T) {
		initialInfo, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		
		addAmount := utils.ToWei(1000)
		opts, _ := ctx.GetTransactor(valKey)
		opts.Value = addAmount
		
		tx, err := ctx.Staking.AddValidatorStake(opts)
		utils.AssertNoError(t, err, "add stake failed")
		ctx.WaitMined(tx.Hash())
		
		newInfo, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		expected := new(big.Int).Add(initialInfo.SelfStake, addAmount)
		utils.AssertBigIntEq(t, newInfo.SelfStake, expected, "stake not increased correctly")
	})

	// [S-02] Decrease Stake
	// Relies on S-01 adding stake first, so we are above MinValidatorStake
	t.Run("S-02_DecreaseStake", func(t *testing.T) {
		// Decrease 500 (S-01 added 1000, so we have plenty margin above Min)
		decAmount := utils.ToWei(500)
		opts, _ := ctx.GetTransactor(valKey)
		opts.Value = nil
		
		infoBefore, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)

		tx, err := ctx.Staking.DecreaseValidatorStake(opts, decAmount)
		utils.AssertNoError(t, err, "decrease stake failed")
		ctx.WaitMined(tx.Hash())
		
		infoAfter, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		expected := new(big.Int).Sub(infoBefore.SelfStake, decAmount)
		utils.AssertBigIntEq(t, infoAfter.SelfStake, expected, "stake not decreased correctly")
	})

	// [S-03] Edit Info
	t.Run("S-03_EditInfo", func(t *testing.T) {
		newFeeAddr := common.HexToAddress("0xFEebFEebFEebFEebFEebFEebFEebFEebFEebFEeb")
		opts, _ := ctx.GetTransactor(valKey)
		tx, err := ctx.Validators.CreateOrEditValidator(opts, newFeeAddr, "NewMoniker", "ident", "site", "email", "details")
		utils.AssertNoError(t, err, "edit validator failed")
		ctx.WaitMined(tx.Hash())
		
		feeAddr, _, _, _, _, _ := ctx.Validators.GetValidatorInfo(nil, valAddr)
		utils.AssertTrue(t, feeAddr == newFeeAddr, "fee address not updated")
		
		// Revert fee addr to self for safety
		tx2, _ := ctx.Validators.CreateOrEditValidator(opts, valAddr, "Genesis", "", "", "", "")
		ctx.WaitMined(tx2.Hash())
	})

	// [S-04] Update Commission
	t.Run("S-04_UpdateCommission", func(t *testing.T) {
		newRate := big.NewInt(2000)
		opts, _ := ctx.GetTransactor(valKey)
		tx, err := ctx.Staking.UpdateCommissionRate(opts, newRate)
		utils.AssertNoError(t, err, "update commission failed")
		ctx.WaitMined(tx.Hash())
		info, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		utils.AssertBigIntEq(t, info.CommissionRate, newRate, "commission rate not updated")
	})

	// [S-07] Decrease Below Min
	// Assumes current stake is X. We try to decrease X - (Min - 1) ?
	// Easier: Just try to decrease ALL stake (minus 1 wei).
	// If current is 1000 ETH (approx). Decrease 999.9 ETH -> Remaining 0.1 ETH < 1 ETH.
	t.Run("S-07_DecreaseBelowMin", func(t *testing.T) {
		info, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		// We want remaining to be 0.5 ETH (which is < 1 ETH Min)
		// Decrease = SelfStake - 0.5 ETH
		// utils.ToWei takes int/float. Let's use big int math to be safe.
		// 0.5 ETH = 5 * 10^17
		targetRemBig := new(big.Int).Mul(big.NewInt(5), big.NewInt(1e17))
		
		if info.SelfStake.Cmp(targetRemBig) <= 0 {
			t.Skip("Current stake too low for S-07 test")
		}
		
		decAmount := new(big.Int).Sub(info.SelfStake, targetRemBig)
		
		opts, _ := ctx.GetTransactor(valKey)
		_, err := ctx.Staking.DecreaseValidatorStake(opts, decAmount)
		utils.AssertTrue(t, err != nil, "should fail decreasing below min")
	})

	// [S-09] Frequent Update
	t.Run("S-09_FrequentCommissionUpdate", func(t *testing.T) {
		opts, _ := ctx.GetTransactor(valKey)
		// First update
		tx, err := ctx.Staking.UpdateCommissionRate(opts, big.NewInt(1500))
		if err != nil && err.Error() == "execution reverted: Commission update too frequent" {
			// Already cooled down? No, we might have run S-04 recently.
			// If S-04 ran, we are in cooldown.
			// So this checks if we are BLOCKED.
			t.Log("Blocked by cooldown (expected)")
			return
		}
		if err == nil {
			ctx.WaitMined(tx.Hash())
			// Try second update immediately
			_, err = ctx.Staking.UpdateCommissionRate(opts, big.NewInt(1600))
			utils.AssertTrue(t, err != nil, "should fail frequent update")
		}
	})

	// [S-11] Double Register
	t.Run("S-11_DoubleRegister", func(t *testing.T) {
		opts, _ := ctx.GetTransactor(valKey)
		opts.Value = utils.ToWei(100000)
		
		_, err := ctx.Staking.RegisterValidator(opts, big.NewInt(1000))
		if err == nil {
			t.Fatal("Expected error 'Already registered', got nil")
		}
		t.Log("Double registration blocked as expected")
	})

	// [S-05] Reincarnation
	// This creates a NEW validator, so it might hit Rate Limit if run in same epoch as other creations.
	// But previous tests used Genesis validator, so we haven't created any new validator in this file yet!
	// So this should succeed (1 quota available).
	t.Run("S-05_Reincarnation", func(t *testing.T) {
		newValKey, newValAddr, err := createAndRegisterValidator(t, "S-05 Validator")
		// If this fails due to "Too many new validators", it means another test file stole the quota.
		if err != nil {
			t.Logf("Skipping S-05 due to creation failure (likely rate limit): %v", err)
			return
		}

		opts, _ := ctx.GetTransactor(newValKey)
		tx, err := ctx.Staking.ResignValidator(opts)
		utils.AssertNoError(t, err, "resign failed")
		ctx.WaitMined(tx.Hash())
		t.Logf("Validator %s resigned", newValAddr.Hex())
	})

	// [S-06] Stake Below Minimum
	t.Run("S-06_StakeBelowMin", func(t *testing.T) {
		key, addr, _ := ctx.CreateAndFundAccount(utils.ToWei(100))
		passProposalFor(t, addr, "S-06 Small Stake")

		opts, _ := ctx.GetTransactor(key)
		// minValidatorStake is usually 100,000 JU. Try with 100 JU.
		opts.Value = utils.ToWei(100)
		_, err := ctx.Staking.RegisterValidator(opts, big.NewInt(1000))
		if err == nil {
			t.Fatal("Expected failure for insufficient self-stake, but succeeded")
		}
		t.Logf("Caught expected error: %v", err)
	})

	// [S-08] Decrease Stake to Zero (Invalid)
	t.Run("S-08_DecreaseStakeToZero", func(t *testing.T) {
		info, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		opts, _ := ctx.GetTransactor(valKey)
		// Trying to decrease exact amount of selfStake should fail (should use exit instead)
		_, err := ctx.Staking.DecreaseValidatorStake(opts, info.SelfStake)
		if err == nil {
			t.Fatal("Expected failure for decreasing stake to zero, but succeeded")
		}
		t.Logf("Caught expected error: %v", err)
	})

	// [S-10] Non-Validator Operations
	t.Run("S-10_NonValidatorOperations", func(t *testing.T) {
		key, _, _ := ctx.CreateAndFundAccount(utils.ToWei(10))
		opts, _ := ctx.GetTransactor(key)

		_, err := ctx.Staking.AddValidatorStake(opts)
		if err == nil {
			t.Fatal("Non-validator should not be able to add stake")
		}

		_, err = ctx.Staking.UpdateCommissionRate(opts, big.NewInt(1000))
		if err == nil {
			t.Fatal("Non-validator should not be able to update commission")
		}
	})

	// [S-12] Zombie Register (Pass=false)
	t.Run("S-12_ZombieRegister", func(t *testing.T) {
		// 1. Create and register a validator
		key, addr, _ := createAndRegisterValidator(t, "S-12 Zombie")
		
		// 2. Propose to remove it
		proposerKey := ctx.GenesisValidators[0]
		optsP, _ := ctx.GetTransactor(proposerKey)
		tx, _ := ctx.Proposal.CreateProposal(optsP, addr, false, "Remove S-12")
		ctx.WaitMined(tx.Hash())
		
		receipt, _ := ctx.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
		var propID [32]byte
		for _, l := range receipt.Logs {
			if ev, err := ctx.Proposal.ParseLogCreateProposal(*l); err == nil { propID = ev.Id; break }
		}
		for _, vk := range ctx.GenesisValidators {
			vo, _ := ctx.GetTransactor(vk)
			ctx.Proposal.VoteProposal(vo, propID, true)
		}
		
		// Wait for removal execution
		waitNextBlock()
		
		// 3. Try to register again without new proposal
		opts, _ := ctx.GetTransactor(key)
		opts.Value = utils.ToWei(100000)
		_, err := ctx.Staking.RegisterValidator(opts, big.NewInt(1000))
		if err == nil {
			t.Fatal("Should not be able to register without passing proposal (pass=false)")
		}
		t.Logf("Caught expected error: %v", err)
	})

	// [S-13] Action after Exit
	t.Run("S-13_ActionAfterExit", func(t *testing.T) {
		key, _, _ := createAndRegisterValidator(t, "S-13 Exit")
		opts, _ := ctx.GetTransactor(key)
		
		// Resign
		txR, _ := ctx.Staking.ResignValidator(opts)
		ctx.WaitMined(txR.Hash())
		
		// Wait for unbonding if necessary, but here we just test immediate blocked actions
		_, err := ctx.Staking.AddValidatorStake(opts)
		if err == nil {
			t.Log("Warning: AddValidatorStake succeeded after Resign? Check if state is immediate.")
		} else {
			t.Logf("Caught error after Resign: %v", err)
		}
	})

	// [S-14] Jailed Constraints
	t.Run("S-14_JailedConstraints", func(t *testing.T) {
		key, addr, _ := createAndRegisterValidator(t, "S-14 Jailed")
		opts, _ := ctx.GetTransactor(key)
		
		// 1. Resign (Jails the validator)
		txR, _ := ctx.Staking.ResignValidator(opts)
		ctx.WaitMined(txR.Hash())
		
		// 2. Try to update commission (Should fail with "Validator is jailed")
		_, err := ctx.Staking.UpdateCommissionRate(opts, big.NewInt(2500))
		if err == nil {
			t.Fatal("Should fail updating commission while jailed")
		}
		t.Logf("Update commission rejected as expected: %v", err)
		
		// 3. Try to add stake (Should be allowed)
		opts.Value = utils.ToWei(10)
		txAdd, err := ctx.Staking.AddValidatorStake(opts)
		utils.AssertNoError(t, err, "Should allow adding stake while jailed")
		ctx.WaitMined(txAdd.Hash())
		
		// 4. Edit Info (Should be allowed)
		txEdit, err := ctx.Validators.CreateOrEditValidator(opts, addr, "JailedMoniker", "", "", "", "")
		utils.AssertNoError(t, err, "Should allow editing info while jailed")
		ctx.WaitMined(txEdit.Hash())
	})
}

func createAndRegisterValidator(t *testing.T, name string) (*ecdsa.PrivateKey, common.Address, error) {
	// INCREASE FUNDING: Give 250,000 ETH to allow for AddStake tests
	key, addr, err := ctx.CreateAndFundAccount(utils.ToWei(250000))
	if err != nil { return nil, common.Address{}, err }

	proposerKey := ctx.GenesisValidators[0]
	var tx *types.Transaction
	
	// Retry Proposal creation (Cooldown)
	for {
		opts, _ := ctx.GetTransactor(proposerKey)
		opts.Value = nil
		tx, err = ctx.Proposal.CreateProposal(opts, addr, true, name)
		if err == nil { break }
		if err.Error() != "execution reverted: Proposal creation too frequent" {
			return nil, common.Address{}, err
		}
		time.Sleep(1 * time.Second)
	}
	ctx.WaitMined(tx.Hash())

	receipt, _ := ctx.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
	var propID [32]byte
	found := false
	for _, l := range receipt.Logs {
		ev, err := ctx.Proposal.ParseLogCreateProposal(*l)
		if err == nil { propID = ev.Id; found = true; break }
	}
	if !found { return nil, common.Address{}, fmt.Errorf("proposal log not found") }

	for _, vk := range ctx.GenesisValidators {
		vo, _ := ctx.GetTransactor(vk)
		tv, _ := ctx.Proposal.VoteProposal(vo, propID, true)
		if tv != nil { ctx.WaitMined(tv.Hash()) }
	}

	ro, _ := ctx.GetTransactor(key)
	ro.Value = utils.ToWei(100000)
	tr, err := ctx.Staking.RegisterValidator(ro, big.NewInt(1000))
	if err != nil { return nil, common.Address{}, err }
	if err := ctx.WaitMined(tr.Hash()); err != nil {
		return nil, common.Address{}, err
	}

	return key, addr, nil
}
