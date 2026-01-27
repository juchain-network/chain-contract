package tests

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"juchain.org/chain/tools/ci/internal/utils"
)

func TestZ_UpgradesAndInitGuards(t *testing.T) {
	if ctx == nil {
		t.Skip("Context not initialized")
	}
	if len(ctx.GenesisValidators) == 0 {
		t.Skip("No genesis validators configured")
	}

	t.Run("InitGuards", func(t *testing.T) {
		dummy := common.HexToAddress("0x000000000000000000000000000000000000dEaD")
		opts, _ := ctx.GetTransactor(ctx.GenesisValidators[0])

		// Proposal.initialize
		_, err := ctx.Proposal.Initialize(opts, []common.Address{dummy}, dummy, big.NewInt(1))
		utils.AssertTrue(t, err != nil, "Proposal.initialize should fail when already initialized")

		// Validators.initialize
		_, err = ctx.Validators.Initialize(opts, []common.Address{dummy}, dummy, dummy, dummy)
		utils.AssertTrue(t, err != nil, "Validators.initialize should fail when already initialized")

		// Punish.initialize
		_, err = ctx.Punish.Initialize(opts, dummy, dummy, dummy)
		utils.AssertTrue(t, err != nil, "Punish.initialize should fail when already initialized")

		// Staking.initialize
		_, err = ctx.Staking.Initialize(opts, dummy, dummy, dummy)
		utils.AssertTrue(t, err != nil, "Staking.initialize should fail when already initialized")

		// Staking.initializeWithValidators
		_, err = ctx.Staking.InitializeWithValidators(opts, dummy, dummy, dummy, []common.Address{dummy}, big.NewInt(1))
		utils.AssertTrue(t, err != nil, "Staking.initializeWithValidators should fail when already initialized")
	})

	t.Run("ReinitializeV2", func(t *testing.T) {
		minerKey, _ := minerKeyOrSkip(t)
		opts, _ := ctx.GetTransactor(minerKey)

		reinit := func(name string, call func() error) {
			if err := call(); err != nil {
				if strings.Contains(err.Error(), "Miner only") {
					t.Skip("caller is not current miner")
				}
				if strings.Contains(err.Error(), "Already reinitialized") {
					t.Logf("%s already reinitialized", name)
					return
				}
				t.Fatalf("%s reinitialize failed: %v", name, err)
			}
		}

		reinit("Proposal", func() error {
			tx, err := ctx.Proposal.ReinitializeV2(opts)
			if err != nil {
				return err
			}
			return ctx.WaitMined(tx.Hash())
		})
		reinit("Validators", func() error {
			tx, err := ctx.Validators.ReinitializeV2(opts)
			if err != nil {
				return err
			}
			return ctx.WaitMined(tx.Hash())
		})
		reinit("Staking", func() error {
			tx, err := ctx.Staking.ReinitializeV2(opts)
			if err != nil {
				return err
			}
			return ctx.WaitMined(tx.Hash())
		})
		reinit("Punish", func() error {
			tx, err := ctx.Punish.ReinitializeV2(opts)
			if err != nil {
				return err
			}
			return ctx.WaitMined(tx.Hash())
		})

		// Second call should fail
		_, err := ctx.Proposal.ReinitializeV2(opts)
		if err == nil {
			t.Fatal("Proposal reinitializeV2 should fail on second call")
		}
	})
}
