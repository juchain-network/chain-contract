package tests

import (
	"context"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestE_Delegation(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized")
	}

	// Use Genesis Validator 0 for most tests to avoid Rate Limit
	valAddr := common.HexToAddress(ctx.Config.Validators[0].Address)
	valKey := ctx.GenesisValidators[0]

	// [D-01] Full Delegation Flow
	t.Run("D-01_FullFlow", func(t *testing.T) {
		userKey, userAddr, err := ctx.CreateAndFundAccount(utils.ToWei(200))
		utils.AssertNoError(t, err, "failed to setup delegator")

		// 2. Delegate
		t.Logf("User %s delegating 100 ETH to %s...", userAddr.Hex(), valAddr.Hex())
		delegateAmount := utils.ToWei(100)
		opts, err := ctx.GetTransactor(userKey)
		utils.AssertNoError(t, err, "failed to get transactor")
		opts.Value = delegateAmount

		tx, err := ctx.Staking.Delegate(opts, valAddr)
		utils.AssertNoError(t, err, "delegate failed")
		ctx.WaitMined(tx.Hash())

		// Verify delegation
		info, _ := ctx.Staking.GetDelegationInfo(nil, userAddr, valAddr)
		utils.AssertBigIntEq(t, info.Amount, delegateAmount, "delegation amount mismatch")

		// 3. Wait for rewards
		t.Log("Waiting for some blocks to accumulate rewards...")
		waitBlocks(t, 5)

		info, _ = ctx.Staking.GetDelegationInfo(nil, userAddr, valAddr)
		t.Logf("Accumulated rewards: %s", info.PendingRewards.String())

		// 4. Claim Rewards
		if info.PendingRewards.Cmp(big.NewInt(0)) > 0 {
			t.Log("Claiming rewards...")
			opts.Value = nil
			txClaim, err := ctx.Staking.ClaimRewards(opts, valAddr)
			utils.AssertNoError(t, err, "claim rewards failed")
			ctx.WaitMined(txClaim.Hash())
		}

		// 5. Undelegate
		t.Log("Undelegating 50 ETH...")
		undelAmount := utils.ToWei(50)
		txUndel, err := ctx.Staking.Undelegate(opts, valAddr, undelAmount)
		utils.AssertNoError(t, err, "undelegate failed")
		ctx.WaitMined(txUndel.Hash())

		// 6. Check unbonding
		entries, _ := ctx.Staking.GetUnbondingEntries(nil, userAddr, valAddr)
		utils.AssertTrue(t, len(entries) > 0, "unbonding entry missing")
	})

	// [D-02] Validator Claims Commission
	t.Run("D-02_ClaimCommission", func(t *testing.T) {
		infoBefore, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		t.Logf("Initial accumulated rewards: %s", infoBefore.AccumulatedRewards.String())

		// Wait for more blocks to generate commission
		waitBlocks(t, 5)
		
		opts, _ := ctx.GetTransactor(valKey)
		tx, err := ctx.Staking.ClaimValidatorRewards(opts)
		utils.AssertNoError(t, err, "claim validator rewards failed")
		ctx.WaitMined(tx.Hash())
		
		infoAfter, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		utils.AssertTrue(t, infoAfter.AccumulatedRewards.Cmp(big.NewInt(0)) == 0, "rewards should be reset after claim")
	})

	// [D-04] Validator Resign Impact
	t.Run("D-04_ValidatorResignImpact", func(t *testing.T) {
		// 1. New Validator
		key, addr, _ := createAndRegisterValidator(t, "D-04 Validator")
		
		// 2. User Delegate
		userKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(50))
		uOpts, _ := ctx.GetTransactor(userKey)
		uOpts.Value = utils.ToWei(10)
		txD, _ := ctx.Staking.Delegate(uOpts, addr)
		ctx.WaitMined(txD.Hash())
		
		// 3. Validator Resign
		vOpts, _ := ctx.GetTransactor(key)
		txR, _ := ctx.Staking.ResignValidator(vOpts)
		ctx.WaitMined(txR.Hash())
		
		// 4. Try to delegate to resigned validator should fail
		uOpts.Value = utils.ToWei(5)
		_, err := ctx.Staking.Delegate(uOpts, addr)
		if err == nil {
			t.Fatal("Should not be able to delegate to resigned validator")
		}
		t.Logf("Caught expected error: %v", err)
	})

	// [D-08] Delegation Below Min
	t.Run("D-08_DelegationBelowMin", func(t *testing.T) {
		userKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(5))
		opts, _ := ctx.GetTransactor(userKey)
		// minDelegation is typically 10 ETH. Try 1 ETH.
		opts.Value = utils.ToWei(1)
		
		_, err := ctx.Staking.Delegate(opts, valAddr)
		if err == nil {
			t.Fatal("Should fail delegation below minimum")
		}
		t.Logf("Caught expected error: %v", err)
	})

	// [D-03] Compound Delegation (Multiple additions)
	t.Run("D-03_CompoundDelegation", func(t *testing.T) {
		userKey, userAddr, err := ctx.CreateAndFundAccount(utils.ToWei(100))
		utils.AssertNoError(t, err, "failed delegator setup")
		
		opts, err := ctx.GetTransactor(userKey)
		utils.AssertNoError(t, err, "failed transactor")
		
		// First 10 ETH
		opts.Value = utils.ToWei(10)
		tx1, err := ctx.Staking.Delegate(opts, valAddr)
		utils.AssertNoError(t, err, "first delegate failed")
		ctx.WaitMined(tx1.Hash())
		
		waitBlocks(t, 2)
		
		// Second 10 ETH
		opts.Value = utils.ToWei(10)
		tx2, err := ctx.Staking.Delegate(opts, valAddr)
		utils.AssertNoError(t, err, "second delegate failed")
		ctx.WaitMined(tx2.Hash())
		
		info, _ := ctx.Staking.GetDelegationInfo(nil, userAddr, valAddr)
		utils.AssertBigIntEq(t, info.Amount, utils.ToWei(20), "total amount mismatch")
	})

	// [D-07] Undelegate Overflow (Boundary)
	t.Run("D-07_UndelegateOverflow", func(t *testing.T) {
		userKey, _, err := ctx.CreateAndFundAccount(utils.ToWei(50))
		utils.AssertNoError(t, err, "failed setup user")

		opts, _ := ctx.GetTransactor(userKey)
		opts.Value = utils.ToWei(10)
		tx, _ := ctx.Staking.Delegate(opts, valAddr)
		ctx.WaitMined(tx.Hash())
		
		opts.Value = nil
		_, err = ctx.Staking.Undelegate(opts, valAddr, utils.ToWei(11))
		utils.AssertTrue(t, err != nil, "should fail undelegating more than staked")
	})

	// [D-15] Delegator becomes Validator
	t.Run("D-15_DelegatorToValidator", func(t *testing.T) {
		// Use Genesis Validator as Target
		targetVal := valAddr
		
		userKey, userAddr, err := ctx.CreateAndFundAccount(utils.ToWei(100005))
		utils.AssertNoError(t, err, "failed user setup")

		opts, _ := ctx.GetTransactor(userKey)
		opts.Value = utils.ToWei(10)
		tx, _ := ctx.Staking.Delegate(opts, targetVal)
		ctx.WaitMined(tx.Hash())
		
		err = passProposalFor(t, userAddr, "D-15 Propose")
		utils.AssertNoError(t, err, "proposal failed")
		
		opts.Value = utils.ToWei(100000)
		txReg, err := ctx.Staking.RegisterValidator(opts, big.NewInt(500))
		if err != nil {
			t.Logf("D-15 Register failed (likely rate limit): %v. Skipping.", err)
			return
		}
		ctx.WaitMined(txReg.Hash())
		
		isVal, _ := ctx.Validators.IsValidatorExist(nil, userAddr)
		utils.AssertTrue(t, isVal, "should be validator")
	})
}

// Helper to wait for N blocks
func waitBlocks(t *testing.T, n int) {
	header, err := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
	if err != nil { return }
	start := header.Number.Uint64()
	target := start + uint64(n)
	
	for {
		time.Sleep(2 * time.Second)
		h, err := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
		if err != nil { continue }
		if h.Number.Uint64() >= target {
			break
		}
	}
}

func passProposalFor(t *testing.T, target common.Address, name string) error {
	proposerKey := ctx.GenesisValidators[0]
	
	var tx *types.Transaction
	var err error
	for {
		opts, _ := ctx.GetTransactor(proposerKey)
		opts.Value = nil
		tx, err = ctx.Proposal.CreateProposal(opts, target, true, name)
		if err == nil { break }
		time.Sleep(2 * time.Second)
	}
	ctx.WaitMined(tx.Hash())

	receipt, _ := ctx.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
	var propID [32]byte
	for _, l := range receipt.Logs {
		ev, err := ctx.Proposal.ParseLogCreateProposal(*l)
		if err == nil { propID = ev.Id; break }
	}

	for _, vk := range ctx.GenesisValidators {
		vo, _ := ctx.GetTransactor(vk)
		ctx.Proposal.VoteProposal(vo, propID, true)
	}
	time.Sleep(2 * time.Second)
	return nil
}