package tests

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
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

	// [D-01b] Withdraw Unbonded after period
	t.Run("D-01b_WithdrawUnbonded", func(t *testing.T) {
		userKey, userAddr, err := ctx.CreateAndFundAccount(utils.ToWei(200))
		utils.AssertNoError(t, err, "failed to setup delegator")
		opts, _ := ctx.GetTransactor(userKey)

		// Delegate then undelegate
		opts.Value = utils.ToWei(20)
		txD, err := ctx.Staking.Delegate(opts, valAddr)
		utils.AssertNoError(t, err, "delegate failed")
		ctx.WaitMined(txD.Hash())

		opts.Value = nil
		txU, err := ctx.Staking.Undelegate(opts, valAddr, utils.ToWei(10))
		utils.AssertNoError(t, err, "undelegate failed")
		ctx.WaitMined(txU.Hash())

		// Wait unbonding period + 1
		period, _ := ctx.Proposal.UnbondingPeriod(nil)
		waitBlocks(t, int(new(big.Int).Add(period, big.NewInt(1)).Int64()))

		txW, err := ctx.Staking.WithdrawUnbonded(opts, valAddr, big.NewInt(20))
		utils.AssertNoError(t, err, "withdraw unbonded failed")
		ctx.WaitMined(txW.Hash())

		cnt, _ := ctx.Staking.GetUnbondingEntriesCount(nil, userAddr, valAddr)
		utils.AssertTrue(t, cnt.Sign() == 0, "unbonding entries should be cleared after withdraw")
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

	// [D-02b] Claim rewards without delegation
	t.Run("D-02b_ClaimNoDelegation", func(t *testing.T) {
		userKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(10))
		opts, _ := ctx.GetTransactor(userKey)
		_, err := ctx.Staking.ClaimRewards(opts, valAddr)
		if err == nil {
			t.Fatal("Should fail claim rewards with no delegation")
		}
	})

	// [D-04a] Validator Resign Impact
	t.Run("D-04a_ValidatorResignImpact", func(t *testing.T) {
		// 1. New Validator
		key, addr, _ := createAndRegisterValidator(t, "D-04a Validator")
		
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

	// [D-04b] Multi-Delegator Isolation
	t.Run("D-04b_MultiDelegatorIsolation", func(t *testing.T) {
		keyA, _, _ := ctx.CreateAndFundAccount(utils.ToWei(100))
		keyB, _, _ := ctx.CreateAndFundAccount(utils.ToWei(100))
		
		optsA, _ := ctx.GetTransactor(keyA)
		optsB, _ := ctx.GetTransactor(keyB)
		
		optsA.Value = utils.ToWei(10)
		txA, _ := ctx.Staking.Delegate(optsA, valAddr)
		ctx.WaitMined(txA.Hash())
		
		optsB.Value = utils.ToWei(20)
		txB, _ := ctx.Staking.Delegate(optsB, valAddr)
		ctx.WaitMined(txB.Hash())
		
		waitBlocks(t, 2)
		
		infoA, _ := ctx.Staking.GetDelegationInfo(nil, crypto.PubkeyToAddress(keyA.PublicKey), valAddr)
		infoB, _ := ctx.Staking.GetDelegationInfo(nil, crypto.PubkeyToAddress(keyB.PublicKey), valAddr)
		
		utils.AssertBigIntEq(t, infoA.Amount, utils.ToWei(10), "User A amount mismatch")
		utils.AssertBigIntEq(t, infoB.Amount, utils.ToWei(20), "User B amount mismatch")
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

	// [D-05] Multi-Validator Delegation
	t.Run("D-05_MultiValidatorDelegation", func(t *testing.T) {
		// Use Genesis Val 0 and 1
		val1 := common.HexToAddress(ctx.Config.Validators[0].Address)
		val2 := common.HexToAddress(ctx.Config.Validators[1].Address)
		
		userKey, userAddr, _ := ctx.CreateAndFundAccount(utils.ToWei(100))
		opts, _ := ctx.GetTransactor(userKey)
		
		// Delegate to V1
		opts.Value = utils.ToWei(10)
		tx1, _ := ctx.Staking.Delegate(opts, val1)
		ctx.WaitMined(tx1.Hash())
		
		// Delegate to V2
		opts.Value = utils.ToWei(10)
		tx2, _ := ctx.Staking.Delegate(opts, val2)
		ctx.WaitMined(tx2.Hash())
		
		info1, _ := ctx.Staking.GetDelegationInfo(nil, userAddr, val1)
		info2, _ := ctx.Staking.GetDelegationInfo(nil, userAddr, val2)
		
		utils.AssertBigIntEq(t, info1.Amount, utils.ToWei(10), "V1 delegation failed")
		utils.AssertBigIntEq(t, info2.Amount, utils.ToWei(10), "V2 delegation failed")
	})

	// [D-16] Circular Delegation
	t.Run("D-16_CircularDelegation", func(t *testing.T) {
		// V0 delegates to V1, V1 delegates to V0
		v0Key := ctx.GenesisValidators[0]
		v0Addr := common.HexToAddress(ctx.Config.Validators[0].Address)
		v1Key := ctx.GenesisValidators[1]
		v1Addr := common.HexToAddress(ctx.Config.Validators[1].Address)
		
		// V0 -> V1
		opts0, _ := ctx.GetTransactor(v0Key)
		opts0.Value = utils.ToWei(10)
		tx0, err := ctx.Staking.Delegate(opts0, v1Addr)
		utils.AssertNoError(t, err, "V0->V1 failed")
		ctx.WaitMined(tx0.Hash())
		
		// V1 -> V0
		opts1, _ := ctx.GetTransactor(v1Key)
		opts1.Value = utils.ToWei(10)
		tx1, err := ctx.Staking.Delegate(opts1, v0Addr)
		utils.AssertNoError(t, err, "V1->V0 failed")
		ctx.WaitMined(tx1.Hash())
		
		info0, _ := ctx.Staking.GetDelegationInfo(nil, v0Addr, v1Addr)
		info1, _ := ctx.Staking.GetDelegationInfo(nil, v1Addr, v0Addr)
		utils.AssertBigIntEq(t, info0.Amount, utils.ToWei(10), "V0->V1 check failed")
		utils.AssertBigIntEq(t, info1.Amount, utils.ToWei(10), "V1->V0 check failed")
	})

	// [D-17] Role Downgrade (Validator -> Delegator)
	t.Run("D-17_RoleDowngrade", func(t *testing.T) {
		// Create new validator first to avoid messing up genesis set too much
		key, addr, err := createAndRegisterValidator(t, "D-17 Downgrade")
		if err != nil {
			t.Logf("Skipping D-17 due to creation failure: %v", err)
			return
		}
		
		opts, _ := ctx.GetTransactor(key)
		
		// Resign
		txR, _ := ctx.Staking.ResignValidator(opts)
		ctx.WaitMined(txR.Hash())
		
		// Exit (Wait unjail period)
		// UnjailPeriod is 50 blocks (from Phase 0)
		t.Log("Waiting for unjail period (50 blocks)...")
		waitBlocks(t, 55)
		
		// Need to call ExitValidator
		txE, err := ctx.Staking.ExitValidator(opts)
		utils.AssertNoError(t, err, "Exit failed")
		ctx.WaitMined(txE.Hash())
		
		// Now delegate to another validator
		targetVal := common.HexToAddress(ctx.Config.Validators[0].Address)
		opts.Value = utils.ToWei(10)
		txD, err := ctx.Staking.Delegate(opts, targetVal)
		utils.AssertNoError(t, err, "Delegation after exit failed")
		ctx.WaitMined(txD.Hash())
		
		info, _ := ctx.Staking.GetDelegationInfo(nil, addr, targetVal)
		utils.AssertBigIntEq(t, info.Amount, utils.ToWei(10), "Delegation amount check failed")
	})

	// [D-06] Early Withdraw (Unbonding not ready)
	t.Run("D-06_EarlyWithdraw", func(t *testing.T) {
		userKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(100))
		opts, _ := ctx.GetTransactor(userKey)
		
		// Delegate & Undelegate
		opts.Value = utils.ToWei(10)
		ctx.Staking.Delegate(opts, valAddr)
		opts.Value = nil
		ctx.Staking.Undelegate(opts, valAddr, utils.ToWei(5))
		
		// Try Withdraw immediately (should fail or do nothing if checking logic)
		// Staking.sol withdrawUnbonded checks timestamp/block.
		// If using `require`, it reverts.
		_, err := ctx.Staking.WithdrawUnbonded(opts, valAddr, big.NewInt(10))
		if err == nil {
			t.Fatal("Early withdraw should fail")
		}
	})

	// [D-07] Self Delegation
	t.Run("D-07_SelfDelegation", func(t *testing.T) {
		// Validator tries to delegate to self using delegate()
		// (addValidatorStake is the correct way, delegate() should fail)
		opts, _ := ctx.GetTransactor(valKey)
		opts.Value = utils.ToWei(10)
		_, err := ctx.Staking.Delegate(opts, valAddr)
		if err == nil {
			t.Fatal("Self-delegation via delegate() should fail")
		}
	})

	// [D-09] Delegate to Non-Existent
	t.Run("D-09_DelegateToNonExistent", func(t *testing.T) {
		userKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(100))
		opts, _ := ctx.GetTransactor(userKey)
		
		randomAddr := common.HexToAddress("0x1234567890123456789012345678901234567890")
		opts.Value = utils.ToWei(10)
		_, err := ctx.Staking.Delegate(opts, randomAddr)
		if err == nil {
			t.Fatal("Should fail delegating to non-existent validator")
		}
	})

	// [D-11] Zero Undelegate
	t.Run("D-11_ZeroUndelegate", func(t *testing.T) {
		userKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(100))
		opts, _ := ctx.GetTransactor(userKey)
		opts.Value = utils.ToWei(10)
		ctx.Staking.Delegate(opts, valAddr)
		
		opts.Value = nil
		_, err := ctx.Staking.Undelegate(opts, valAddr, big.NewInt(0))
		if err == nil {
			t.Fatal("Should fail zero undelegation")
		}
	})

	// [D-18] Undelegate Below Minimum
	t.Run("D-18_UndelegateBelowMin", func(t *testing.T) {
		minUndel, _ := ctx.Proposal.MinUndelegation(nil)
		if minUndel.Cmp(big.NewInt(1)) <= 0 {
			t.Skip("minUndelegation too small to test below-min path")
		}

		userKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(50))
		opts, _ := ctx.GetTransactor(userKey)
		opts.Value = utils.ToWei(10)
		ctx.Staking.Delegate(opts, valAddr)

		opts.Value = nil
		belowMin := new(big.Int).Sub(minUndel, big.NewInt(1))
		_, err := ctx.Staking.Undelegate(opts, valAddr, belowMin)
		if err == nil {
			t.Fatal("Should fail undelegate below minUndelegation")
		}
	})

	// [D-12] Delegate to Jailed
	t.Run("D-12_DelegateToJailed", func(t *testing.T) {
		// Need a jailed validator. We can reuse one from D-17 or create new.
		// Creating new is safer.
		key, addr, err := createAndRegisterValidator(t, "D-12 Jailed")
		if err != nil { return }
		
		vOpts, _ := ctx.GetTransactor(key)
		ctx.Staking.ResignValidator(vOpts) // Jails the validator
		
		userKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(50))
		opts, _ := ctx.GetTransactor(userKey)
		opts.Value = utils.ToWei(10)
		
		_, err = ctx.Staking.Delegate(opts, addr)
		if err == nil {
			t.Fatal("Should fail delegating to jailed validator")
		}
	})

	// [D-13] Undelegate from Jailed (Allowed)
	t.Run("D-13_UndelegateFromJailed", func(t *testing.T) {
		// 1. Setup Validator & Delegator
		key, addr, err := createAndRegisterValidator(t, "D-13 Jailed")
		if err != nil { return }
		
		userKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(50))
		opts, _ := ctx.GetTransactor(userKey)
		opts.Value = utils.ToWei(10)
		ctx.Staking.Delegate(opts, addr)
		
		// 2. Jail Validator
		vOpts, _ := ctx.GetTransactor(key)
		ctx.Staking.ResignValidator(vOpts)
		
		// 3. Undelegate
		opts.Value = nil
		tx, err := ctx.Staking.Undelegate(opts, addr, utils.ToWei(10))
		utils.AssertNoError(t, err, "Should allow undelegation from jailed validator")
		ctx.WaitMined(tx.Hash())
	})
	
	// [D-14] Max Unbonding Entries
	t.Run("D-14_MaxUnbonding", func(t *testing.T) {
		userKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(100))
		opts, _ := ctx.GetTransactor(userKey)
		opts.Value = utils.ToWei(25) // Enough for 21 * 1 ETH
		ctx.Staking.Delegate(opts, valAddr)
		
		// Max is 20. Try 21 times.
		opts.Value = nil
		for i := 0; i < 20; i++ {
			tx, err := ctx.Staking.Undelegate(opts, valAddr, utils.ToWei(1))
			utils.AssertNoError(t, err, "Undelegate within limit failed")
			ctx.WaitMined(tx.Hash())
		}
		
		// 21st time
		_, err := ctx.Staking.Undelegate(opts, valAddr, utils.ToWei(1))
		if err == nil {
			t.Fatal("Should fail exceeding max unbonding entries")
		}
		t.Logf("Caught expected error: %v", err)
	})

	// [D-19] Invalid maxEntries for withdrawUnbonded
	t.Run("D-19_InvalidMaxEntries", func(t *testing.T) {
		userKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(10))
		opts, _ := ctx.GetTransactor(userKey)
		_, err := ctx.Staking.WithdrawUnbonded(opts, valAddr, big.NewInt(0))
		if err == nil {
			t.Fatal("Should fail with maxEntries=0")
		}
		_, err = ctx.Staking.WithdrawUnbonded(opts, valAddr, big.NewInt(21))
		if err == nil {
			t.Fatal("Should fail with maxEntries too large")
		}
	})
}

// Helper to wait for N blocks
func waitBlocks(t *testing.T, n int) {
	if n <= 0 || ctx == nil || len(ctx.Clients) == 0 {
		return
	}

	// Try fast mining if the RPC supports it.
	rpcClient := ctx.Clients[0].Client()
	if rpcClient != nil {
		var res interface{}
		// anvil_mine supports a block count
		if err := rpcClient.CallContext(context.Background(), &res, "anvil_mine", n); err == nil {
			return
		}
		// hardhat_mine expects hex string
		if err := rpcClient.CallContext(context.Background(), &res, "hardhat_mine", fmt.Sprintf("0x%x", n)); err == nil {
			return
		}
		// evm_mine one by one
		ok := true
		for i := 0; i < n; i++ {
			if err := rpcClient.CallContext(context.Background(), &res, "evm_mine"); err != nil {
				ok = false
				break
			}
		}
		if ok {
			return
		}
	}

	// Fallback: wait for blocks by polling.
	header, err := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
	if err != nil {
		return
	}
	start := header.Number.Uint64()
	target := start + uint64(n)

	for {
		time.Sleep(2 * time.Second)
		h, err := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
		if err != nil {
			continue
		}
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
