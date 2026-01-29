package tests

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
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

		checkReinit := func(name string, call func() (*types.Transaction, error)) {
			tx, err := call()
			if err == nil {
				// If simulation didn't catch it, WaitMined must
				err = ctx.WaitMined(tx.Hash())
			}
			
			if err == nil {
				t.Errorf("%s.initialize should have failed", name)
			} else {
				t.Logf("%s.initialize failed as expected: %v", name, err)
			}
		}

		// Proposal.initialize
		checkReinit("Proposal", func() (*types.Transaction, error) {
			return ctx.Proposal.Initialize(opts, []common.Address{dummy}, dummy, big.NewInt(1))
		})

		// Validators.initialize
		checkReinit("Validators", func() (*types.Transaction, error) {
			return ctx.Validators.Initialize(opts, []common.Address{dummy}, dummy, dummy, dummy)
		})

		// Punish.initialize
		checkReinit("Punish", func() (*types.Transaction, error) {
			return ctx.Punish.Initialize(opts, dummy, dummy, dummy)
		})

		// Staking.initialize
		checkReinit("Staking", func() (*types.Transaction, error) {
			return ctx.Staking.Initialize(opts, dummy, dummy, dummy)
		})

		// Staking.initializeWithValidators
		checkReinit("StakingValidators", func() (*types.Transaction, error) {
			return ctx.Staking.InitializeWithValidators(opts, dummy, dummy, dummy, []common.Address{dummy}, big.NewInt(1))
		})
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
		tx, err := ctx.Proposal.ReinitializeV2(opts)
		if err == nil {
			err = ctx.WaitMined(tx.Hash())
		}
		if err == nil {
			t.Fatal("Proposal reinitializeV2 should fail on second call")
		}
	})
}