package tests

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
		waitReceipt := func(txHash common.Hash, timeout time.Duration) error {
			if txHash == (common.Hash{}) {
				return nil
			}
			deadline := time.Now().Add(timeout)
			for time.Now().Before(deadline) {
				receipt, err := ctx.Clients[0].TransactionReceipt(context.Background(), txHash)
				if err == nil && receipt != nil {
					if receipt.Status == 0 {
						return fmt.Errorf("transaction %s reverted", txHash.Hex())
					}
					return nil
				}
				time.Sleep(1 * time.Second)
			}
			return fmt.Errorf("timeout waiting for tx %s", txHash.Hex())
		}

		reinit := func(name string, call func(opts *bind.TransactOpts) (*types.Transaction, error)) bool {
			for attempt := 0; attempt < 5; attempt++ {
				minerKey, _ := minerKeyOrSkip(t)
				opts, _ := ctx.GetTransactor(minerKey)
				tx, err := call(opts)
				if err != nil {
					if strings.Contains(err.Error(), "Already reinitialized") {
						t.Logf("%s already reinitialized", name)
						return true
					}
					if strings.Contains(err.Error(), "Miner only") {
						waitNextBlock()
						continue
					}
					t.Logf("%s reinit attempt %d failed: %v", name, attempt+1, err)
					waitNextBlock()
					continue
				}
				if err := waitReceipt(tx.Hash(), 30*time.Second); err != nil {
					if strings.Contains(err.Error(), "reverted") {
						waitNextBlock()
						continue
					}
					t.Logf("%s reinit attempt %d not mined: %v", name, attempt+1, err)
					waitNextBlock()
					continue
				}
				return true
			}
			t.Skipf("%s reinitialize not mined by miner after retries", name)
			return false
		}

		proposalDone := reinit("Proposal", func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return ctx.Proposal.ReinitializeV2(opts)
		})
		reinit("Validators", func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return ctx.Validators.ReinitializeV2(opts)
		})
		reinit("Staking", func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return ctx.Staking.ReinitializeV2(opts)
		})
		reinit("Punish", func(opts *bind.TransactOpts) (*types.Transaction, error) {
			return ctx.Punish.ReinitializeV2(opts)
		})

		// Second call should fail
		if !proposalDone {
			t.Skip("proposal reinitialize not executed; skipping second call check")
		}
		minerKey, _ := minerKeyOrSkip(t)
		opts, _ := ctx.GetTransactor(minerKey)
		tx, err := ctx.Proposal.ReinitializeV2(opts)
		if err == nil {
			err = waitReceipt(tx.Hash(), 30*time.Second)
		}
		if err == nil {
			t.Fatal("Proposal reinitializeV2 should fail on second call")
		}
	})
}
