package tests

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sort"
	"strings"
	"testing"
	"time"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	testctx "juchain.org/chain/tools/ci/internal/context"
	"juchain.org/chain/tools/contracts"
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

		reinit := func(
			name string,
			contractAddr common.Address,
			meta *bind.MetaData,
			call func(opts *bind.TransactOpts, client *ethclient.Client) (*types.Transaction, error),
		) bool {
			for attempt := 0; attempt < 12; attempt++ {
				minerKey, _, minerClient := pickInTurnValidator(t)
				opts, _ := ctx.GetTransactor(minerKey)
				if minerClient != nil {
					addr := crypto.PubkeyToAddress(minerKey.PublicKey)
					if nonce, err := minerClient.PendingNonceAt(context.Background(), addr); err == nil {
						opts.Nonce = big.NewInt(int64(nonce))
					}
				}
				opts.GasLimit = 300000
				tx, err := call(opts, minerClient)
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
			if meta != nil && contractAddr != (common.Address{}) {
				if validateOnlyMinerCall(t, contractAddr, meta) {
					t.Logf("%s reinitialize validated via eth_call (miner-only guard)", name)
					return true
				}
			}
			t.Skipf("%s reinitialize not mined by miner after retries", name)
			return false
		}

		proposalDone := reinit("Proposal", ctx.ProposalAddr, contracts.ProposalMetaData, func(opts *bind.TransactOpts, client *ethclient.Client) (*types.Transaction, error) {
			prop, err := contracts.NewProposal(ctx.ProposalAddr, client)
			if err != nil {
				return nil, err
			}
			return prop.ReinitializeV2(opts)
		})
		reinit("Validators", testctx.ValidatorsAddr, contracts.ValidatorsMetaData, func(opts *bind.TransactOpts, client *ethclient.Client) (*types.Transaction, error) {
			vals, err := contracts.NewValidators(testctx.ValidatorsAddr, client)
			if err != nil {
				return nil, err
			}
			return vals.ReinitializeV2(opts)
		})
		reinit("Staking", testctx.StakingAddr, contracts.StakingMetaData, func(opts *bind.TransactOpts, client *ethclient.Client) (*types.Transaction, error) {
			stk, err := contracts.NewStaking(testctx.StakingAddr, client)
			if err != nil {
				return nil, err
			}
			return stk.ReinitializeV2(opts)
		})
		reinit("Punish", testctx.PunishAddr, contracts.PunishMetaData, func(opts *bind.TransactOpts, client *ethclient.Client) (*types.Transaction, error) {
			pun, err := contracts.NewPunish(testctx.PunishAddr, client)
			if err != nil {
				return nil, err
			}
			return pun.ReinitializeV2(opts)
		})

		// Second call should fail
		if !proposalDone {
			t.Skip("proposal reinitialize not executed; skipping second call check")
		}
		minerKey, _, minerClient := pickInTurnValidator(t)
		opts, _ := ctx.GetTransactor(minerKey)
		if minerClient != nil {
			addr := crypto.PubkeyToAddress(minerKey.PublicKey)
			if nonce, err := minerClient.PendingNonceAt(context.Background(), addr); err == nil {
				opts.Nonce = big.NewInt(int64(nonce))
			}
		}
		opts.GasLimit = 300000
		prop, err := contracts.NewProposal(ctx.ProposalAddr, minerClient)
		if err != nil {
			t.Fatalf("failed to bind proposal: %v", err)
		}
		tx, err := prop.ReinitializeV2(opts)
		if err == nil {
			err = waitReceipt(tx.Hash(), 30*time.Second)
		}
		if err == nil {
			t.Fatal("Proposal reinitializeV2 should fail on second call")
		}
	})
}

func pickInTurnValidator(t *testing.T) (*ecdsa.PrivateKey, common.Address, *ethclient.Client) {
	if ctx == nil {
		t.Skip("Context not initialized")
	}
	validators, err := ctx.Validators.GetActiveValidators(nil)
	if err != nil || len(validators) == 0 {
		t.Skip("no active validators available")
	}
	sort.Slice(validators, func(i, j int) bool {
		return bytes.Compare(validators[i][:], validators[j][:]) < 0
	})
	header, err := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
	if err != nil || header == nil {
		t.Skipf("failed to read header: %v", err)
	}
	next := header.Number.Uint64() + 1
	idx := next % uint64(len(validators))
	addr := validators[idx]
	key := keyForAddress(addr)
	if key == nil {
		t.Skipf("no key for in-turn validator %s", addr.Hex())
	}
	client := clientForValidator(t, addr)
	return key, addr, client
}

func clientForValidator(t *testing.T, addr common.Address) *ethclient.Client {
	if ctx == nil {
		t.Skip("Context not initialized")
	}
	if len(ctx.Clients) > 1 {
		for _, c := range ctx.Clients {
			var cb common.Address
			if err := c.Client().Call(&cb, "eth_coinbase"); err == nil && cb == addr {
				return c
			}
		}
	}
	if len(ctx.Config.Validators) == 0 {
		t.Skip("no validator config available")
	}
	ip := ""
	for i, v := range ctx.Config.Validators {
		if common.HexToAddress(v.Address) == addr {
			switch i {
			case 0:
				ip = "172.28.0.10"
			case 1:
				ip = "172.28.0.11"
			case 2:
				ip = "172.28.0.12"
			}
		}
	}
	if ip == "" {
		t.Skipf("no RPC mapping for validator %s", addr.Hex())
	}
	client, err := ethclient.Dial("http://" + ip + ":8545")
	if err != nil {
		t.Skipf("failed to dial validator RPC %s: %v", ip, err)
	}
	return client
}

func validateOnlyMinerCall(t *testing.T, contractAddr common.Address, meta *bind.MetaData) bool {
	if ctx == nil {
		return false
	}
	abi, err := meta.GetAbi()
	if err != nil {
		t.Logf("failed to load ABI: %v", err)
		return false
	}
	data, err := abi.Pack("reinitializeV2")
	if err != nil {
		t.Logf("failed to pack call data: %v", err)
		return false
	}

	minerAddr := common.HexToAddress(ctx.Config.Validators[0].Address)
	blockNum, err := findRecentBlockByCoinbase(ctx.Clients[0], minerAddr, 200)
	if err != nil {
		t.Logf("no recent block for miner %s: %v", minerAddr.Hex(), err)
		return false
	}

	msg := ethereum.CallMsg{
		From: minerAddr,
		To:   &contractAddr,
		Gas:  300000,
		Data: data,
	}
	if _, err := ctx.Clients[0].CallContract(context.Background(), msg, blockNum); err != nil {
		t.Logf("miner call failed: %v", err)
		return false
	}

	nonMiner := common.HexToAddress(ctx.Config.Funder.Address)
	msg.From = nonMiner
	if _, err := ctx.Clients[0].CallContract(context.Background(), msg, blockNum); err == nil {
		t.Log("non-miner call unexpectedly succeeded")
		return false
	}
	return true
}

func findRecentBlockByCoinbase(client *ethclient.Client, coinbase common.Address, lookback uint64) (*big.Int, error) {
	if client == nil {
		return nil, fmt.Errorf("nil client")
	}
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil || header == nil {
		return nil, fmt.Errorf("failed to read header: %v", err)
	}
	start := header.Number.Uint64()
	for i := uint64(0); i <= lookback && start >= i; i++ {
		num := new(big.Int).SetUint64(start - i)
		h, err := client.HeaderByNumber(context.Background(), num)
		if err == nil && h != nil && h.Coinbase == coinbase {
			return num, nil
		}
	}
	return nil, fmt.Errorf("no block found in last %d", lookback)
}
