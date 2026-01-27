package tests

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestI_ValidatorExtras(t *testing.T) {
	if ctx == nil || len(ctx.GenesisValidators) == 0 {
		t.Skip("Context not initialized")
	}

	valKey := ctx.GenesisValidators[0]
	valAddr := common.HexToAddress(ctx.Config.Validators[0].Address)
	pass, _ := ctx.Proposal.Pass(nil, valAddr)
	if !pass {
		t.Skip("validator not authorized for edit tests")
	}

	// Description boundary checks (identity, website, email, details)
	t.Run("V-02b_DescriptionBoundaryFields", func(t *testing.T) {
		opts, _ := ctx.GetTransactor(valKey)

		tooLong := func(n int) string {
			b := make([]byte, n)
			for i := range b {
				b[i] = 'a'
			}
			return string(b)
		}

		// identity > 3000
		_, err := ctx.Validators.CreateOrEditValidator(opts, valAddr, "ok", tooLong(3001), "", "", "")
		if err == nil {
			t.Fatal("identity > 3000 should fail")
		}
		// website > 140
		_, err = ctx.Validators.CreateOrEditValidator(opts, valAddr, "ok", "", tooLong(141), "", "")
		if err == nil {
			t.Fatal("website > 140 should fail")
		}
		// email > 140
		_, err = ctx.Validators.CreateOrEditValidator(opts, valAddr, "ok", "", "", tooLong(141), "")
		if err == nil {
			t.Fatal("email > 140 should fail")
		}
		// details > 280
		_, err = ctx.Validators.CreateOrEditValidator(opts, valAddr, "ok", "", "", "", tooLong(281))
		if err == nil {
			t.Fatal("details > 280 should fail")
		}
	})

	// Withdraw profits exceptions
	t.Run("V-04_WithdrawProfitsExceptions", func(t *testing.T) {
		feeAddr, _, incoming, _, _, _ := ctx.Validators.GetValidatorInfo(nil, valAddr)

		// Non-fee address should fail
		userKey, _, _ := ctx.CreateAndFundAccount(utils.ToWei(1))
		userOpts, _ := ctx.GetTransactor(userKey)
		_, err := ctx.Validators.WithdrawProfits(userOpts, valAddr)
		if err == nil || !strings.Contains(err.Error(), "fee receiver") {
			t.Fatalf("expected non-fee withdrawal to fail, got: %v", err)
		}

		// Zero-profit path (ensure fee key available)
		feeKey := keyForAddress(feeAddr)
		if feeKey == nil {
			t.Skip("fee address key not available for zero-profit check")
		}
		feeOpts, _ := ctx.GetTransactor(feeKey)

		if incoming.Cmp(big.NewInt(0)) > 0 {
			// Withdraw once to clear profits
			tx, err := ctx.Validators.WithdrawProfits(feeOpts, valAddr)
			if err != nil {
				t.Skipf("cannot withdraw to clear profits: %v", err)
			}
			ctx.WaitMined(tx.Hash())
			// Wait cooldown, then try again to hit "no profits"
			period, _ := ctx.Proposal.WithdrawProfitPeriod(nil)
			waitBlocks(t, int(new(big.Int).Add(period, big.NewInt(1)).Int64()))
		}

		_, err = ctx.Validators.WithdrawProfits(feeOpts, valAddr)
		if err == nil {
			t.Fatal("expected withdraw with zero profits to fail")
		}
	})
}
