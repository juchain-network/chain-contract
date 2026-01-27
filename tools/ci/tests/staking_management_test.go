package tests

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestD_StakingManagement(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized or no validators")
	}

	// [S-01] Add Stake
	t.Run("S-01_AddStake", func(t *testing.T) {
		valKey, valAddr, err := createAndRegisterValidator(t, "S-01 Validator")
		utils.AssertNoError(t, err, "failed to create validator")
		if valKey == nil {
			t.Fatal("valKey is nil but err is nil")
		}
		
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
	t.Run("S-02_DecreaseStake", func(t *testing.T) {
		// Need to create with more funds so we can increase then decrease
		valKey, valAddr, _ := createAndRegisterValidator(t, "S-02 Validator")
		
		// 1. Add Stake first (Current is 100k, add 10k to make it 110k)
		addOpts, _ := ctx.GetTransactor(valKey)
		addOpts.Value = utils.ToWei(10000)
		txAdd, _ := ctx.Staking.AddValidatorStake(addOpts)
		ctx.WaitMined(txAdd.Hash())

		// 2. Decrease Stake (Decrease 5k, leaving 105k which is > 100k min)
		decAmount := utils.ToWei(5000)
		opts, _ := ctx.GetTransactor(valKey)
		
		tx, err := ctx.Staking.DecreaseValidatorStake(opts, decAmount)
		utils.AssertNoError(t, err, "decrease stake failed")
		ctx.WaitMined(tx.Hash())
		
		info, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		expected := utils.ToWei(105000)
		utils.AssertBigIntEq(t, info.SelfStake, expected, "stake not decreased correctly")
	})

	// [S-03] Edit Info
	t.Run("S-03_EditInfo", func(t *testing.T) {
		valKey, valAddr, _ := createAndRegisterValidator(t, "S-03 Validator")
		newFeeAddr := common.HexToAddress("0xFEeb")
		opts, _ := ctx.GetTransactor(valKey)
		tx, err := ctx.Validators.CreateOrEditValidator(opts, newFeeAddr, "NewMoniker", "ident", "site", "email", "details")
		utils.AssertNoError(t, err, "edit validator failed")
		ctx.WaitMined(tx.Hash())
		feeAddr, _, _, _, _, _ := ctx.Validators.GetValidatorInfo(nil, valAddr)
		utils.AssertTrue(t, feeAddr == newFeeAddr, "fee address not updated")
	})

	// [S-04] Update Commission
	t.Run("S-04_UpdateCommission", func(t *testing.T) {
		valKey, valAddr, _ := createAndRegisterValidator(t, "S-04 Validator")
		newRate := big.NewInt(2000)
		opts, _ := ctx.GetTransactor(valKey)
		tx, err := ctx.Staking.UpdateCommissionRate(opts, newRate)
		utils.AssertNoError(t, err, "update commission failed")
		ctx.WaitMined(tx.Hash())
		info, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		utils.AssertBigIntEq(t, info.CommissionRate, newRate, "commission rate not updated")
	})

	t.Run("S-07_DecreaseBelowMin", func(t *testing.T) {
		valKey, _, err := createAndRegisterValidator(t, "S-07 Validator")
		utils.AssertNoError(t, err, "failed to create validator")

		decAmount := big.NewInt(1)
		opts, _ := ctx.GetTransactor(valKey)
		_, err = ctx.Staking.DecreaseValidatorStake(opts, decAmount)
		utils.AssertTrue(t, err != nil, "should fail decreasing below min")
	})

	t.Run("S-09_FrequentCommissionUpdate", func(t *testing.T) {
		valKey, _, err := createAndRegisterValidator(t, "S-09 Validator")
		utils.AssertNoError(t, err, "failed to create validator")

		opts, _ := ctx.GetTransactor(valKey)
		tx, _ := ctx.Staking.UpdateCommissionRate(opts, big.NewInt(1500))
		ctx.WaitMined(tx.Hash())
		_, err = ctx.Staking.UpdateCommissionRate(opts, big.NewInt(1600))
		utils.AssertTrue(t, err != nil, "should fail frequent update")
	})

	t.Run("S-11_DoubleRegister", func(t *testing.T) {
		valKey, _, err := createAndRegisterValidator(t, "S-11 Validator")
		utils.AssertNoError(t, err, "failed to create validator")

		opts, _ := ctx.GetTransactor(valKey)
		opts.Value = utils.ToWei(100000)
		
		_, err = ctx.Staking.RegisterValidator(opts, big.NewInt(1000))
		if err == nil {
			t.Fatal("Expected error 'Already registered', got nil")
		}
		t.Log("Double registration blocked as expected")
	})

	t.Run("S-05_Reincarnation", func(t *testing.T) {
		valKey, valAddr, err := createAndRegisterValidator(t, "S-05 Validator")
		utils.AssertNoError(t, err, "failed to create validator")

		opts, _ := ctx.GetTransactor(valKey)
		ctx.Staking.ResignValidator(opts)
		t.Logf("Validator %s resigned", valAddr.Hex())
	})
}

func createAndRegisterValidator(t *testing.T, name string) (*ecdsa.PrivateKey, common.Address, error) {
	// INCREASE FUNDING: Give 250,000 ETH to allow for AddStake tests
	key, addr, err := ctx.CreateAndFundAccount(utils.ToWei(250000))
	if err != nil { return nil, common.Address{}, err }

	proposerKey := ctx.GenesisValidators[0]
	var tx *types.Transaction
	for {
		opts, _ := ctx.GetTransactor(proposerKey)
		opts.Value = nil
		tx, err = ctx.Proposal.CreateProposal(opts, addr, true, name)
		if err == nil { break }
		waitNextBlock()
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
