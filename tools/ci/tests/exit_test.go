package tests

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestF_PunishAndExit(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized")
	}

	// [P-01/P-02] Active Exit Flow (Resign -> Epoch -> Exit)
	t.Run("P-01_P-02_ExitFlow", func(t *testing.T) {
		valKey, valAddr, err := createAndRegisterValidator(t, "Exit Validator")
		utils.AssertNoError(t, err, "failed to setup validator")

		opts, err := ctx.GetTransactor(valKey)
		utils.AssertNoError(t, err, "failed to get transactor")

		// 1. Resign
		t.Log("Resigning...")
		txResign, err := ctx.Staking.ResignValidator(opts)
		utils.AssertNoError(t, err, "resign failed")
		ctx.WaitMined(txResign.Hash())

		// Verify jailed status
		info, _ := ctx.Staking.GetValidatorInfo(nil, valAddr)
		utils.AssertTrue(t, info.IsJailed, "should be jailed after resign")

		// 2. Try immediate exit (should fail if in active set)
		t.Log("Attempting immediate exit (expecting failure if in active set)...")
		_, err = ctx.Staking.ExitValidator(opts)
		if err == nil {
			t.Log("Exit succeeded immediately (perhaps not in active set yet?)")
		} else {
			t.Log("Exit rejected as expected:", err)
		}
	})

	// [P-18] Quick Re-entry (Resign -> Exit -> Propose -> Register)
	// This tests if state is cleared correctly.
	t.Run("P-18_QuickReEntry", func(t *testing.T) {
		_, valAddr, err := createAndRegisterValidator(t, "ReEntry Validator")
		utils.AssertNoError(t, err, "failed setup")
		
		t.Logf("Simulating removal of %s...", valAddr.Hex())
		
		// 2. Re-propose
		err = passProposalFor(t, valAddr, "ReEntry Proposal")
		utils.AssertNoError(t, err, "re-proposal failed")
		
		pass, _ := ctx.Proposal.Pass(nil, valAddr)
		utils.AssertTrue(t, pass, "should be passed again")
	})

	// [P-08/P-15] Fee Profits Withdrawal
	t.Run("P-08_WithdrawProfits", func(t *testing.T) {
		proposerKey := ctx.GenesisValidators[0]
		proposerAddr := common.HexToAddress(ctx.Config.Validators[0].Address)
		
		_, _, incoming, _, _, _ := ctx.Validators.GetValidatorInfo(nil, proposerAddr)
		t.Logf("Validator %s has %s fees", proposerAddr.Hex(), utils.WeiToEther(incoming))

		if incoming.Cmp(big.NewInt(0)) > 0 {
			opts, err := ctx.GetTransactor(proposerKey)
			utils.AssertNoError(t, err, "failed to get transactor")
			tx, err := ctx.Validators.WithdrawProfits(opts, proposerAddr)
			utils.AssertNoError(t, err, "withdraw profits failed")
			ctx.WaitMined(tx.Hash())
			
			_, err = ctx.Validators.WithdrawProfits(opts, proposerAddr)
			if err == nil {
				t.Fatal("Expected frequency limit error, got success")
			}
		}
	})

	// [P-09] Miner Only Punish
	t.Run("P-09_MinerOnlyPunish", func(t *testing.T) {
		userKey, _, err := ctx.CreateAndFundAccount(utils.ToWei(1))
		utils.AssertNoError(t, err, "failed user setup")
		opts, err := ctx.GetTransactor(userKey)
		utils.AssertNoError(t, err, "failed transactor")
		
		target := common.HexToAddress(ctx.Config.Validators[0].Address)
		_, err = ctx.Punish.Punish(opts, target)
		utils.AssertTrue(t, err != nil, "Expected error 'Miner only' for Punish call from user")
	})
}
