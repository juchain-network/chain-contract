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
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	testctx "juchain.org/chain/tools/ci/internal/context"
	"juchain.org/chain/tools/contracts"
)

func TestZ_UpgradesAndInitGuards(t *testing.T) {
	if ctx == nil {
		t.Fatalf("Context not initialized")
	}
	if len(ctx.GenesisValidators) == 0 {
		t.Fatalf("No genesis validators configured")
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
		waitReceipt := func(client *ethclient.Client, txHash common.Hash, timeout time.Duration) error {
			if txHash == (common.Hash{}) {
				return nil
			}
			if client == nil {
				client = ctx.Clients[0]
			}
			deadline := time.Now().Add(timeout)
			for time.Now().Before(deadline) {
				receipt, err := client.TransactionReceipt(context.Background(), txHash)
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

		makeTxOpts := func(key *ecdsa.PrivateKey, client *ethclient.Client) *bind.TransactOpts {
			if client == nil {
				client = ctx.Clients[0]
			}
			opts, err := bind.NewKeyedTransactorWithChainID(key, ctx.ChainID)
			if err != nil {
				t.Fatalf("failed to create transactor: %v", err)
			}
			addr := crypto.PubkeyToAddress(key.PublicKey)
			nonce, err := client.PendingNonceAt(context.Background(), addr)
			if err != nil {
				t.Fatalf("failed to get nonce for %s: %v", addr.Hex(), err)
			}
			opts.Nonce = big.NewInt(int64(nonce))
			gasPrice, err := client.SuggestGasPrice(context.Background())
			if err != nil {
				gasPrice = big.NewInt(1000000000) // 1 gwei fallback
			}
			opts.GasPrice = gasPrice
			return opts
		}

		sendLocalReinit := func(client *ethclient.Client, from common.Address, contractAddr common.Address, meta *bind.MetaData) (common.Hash, error) {
			if client == nil {
				return common.Hash{}, fmt.Errorf("nil client")
			}
			abi, err := meta.GetAbi()
			if err != nil {
				return common.Hash{}, err
			}
			data, err := abi.Pack("reinitializeV2")
			if err != nil {
				return common.Hash{}, err
			}
			args := map[string]interface{}{
				"from":     from,
				"to":       contractAddr,
				"gas":      hexutil.Uint64(300000),
				"gasPrice": (*hexutil.Big)(big.NewInt(1000000000)), // 1 gwei
				"data":     hexutil.Bytes(data),
			}
			var txHash common.Hash
			if err := client.Client().Call(&txHash, "eth_sendTransaction", args); err != nil {
				return common.Hash{}, err
			}
			return txHash, nil
		}

		reinit := func(
			name string,
			contractAddr common.Address,
			meta *bind.MetaData,
			call func(opts *bind.TransactOpts, client *ethclient.Client) (*types.Transaction, error),
			revision func() (*big.Int, error),
		) bool {
			for attempt := 0; attempt < 24; attempt++ {
				minerKey, minerAddr, minerClient := pickInTurnValidator(t)
				if minerClient != nil {
					var cb common.Address
					if err := minerClient.Client().Call(&cb, "eth_coinbase"); err == nil && cb != minerAddr {
						t.Fatalf("validator RPC coinbase mismatch: expected %s got %s", minerAddr.Hex(), cb.Hex())
					}
				}
				var (
					txHash common.Hash
					err    error
				)
				txHash, err = sendLocalReinit(minerClient, minerAddr, contractAddr, meta)
				if err != nil {
					msg := err.Error()
					if strings.Contains(msg, "authentication needed") || strings.Contains(msg, "does not exist") || strings.Contains(msg, "not available") {
						opts := makeTxOpts(minerKey, minerClient)
						opts.GasLimit = 300000
						var tx *types.Transaction
						tx, err = call(opts, minerClient)
						if err == nil {
							txHash = tx.Hash()
						}
					}
				}
				if err != nil {
					if strings.Contains(err.Error(), "Already reinitialized") {
						t.Logf("%s already reinitialized", name)
						return true
					}
					if strings.Contains(err.Error(), "Miner only") {
						if attempt == 0 {
							t.Logf("%s reinit rejected by miner-only guard at send time", name)
						}
						waitNextBlock()
						continue
					}
					t.Logf("%s reinit attempt %d failed: %v", name, attempt+1, err)
					waitNextBlock()
					continue
				}
				if err := waitReceipt(minerClient, txHash, 30*time.Second); err != nil {
					if strings.Contains(err.Error(), "reverted") {
						if revision != nil {
							if rev, rerr := revision(); rerr == nil && rev.Cmp(big.NewInt(2)) >= 0 {
								t.Logf("%s already at revision %s", name, rev.String())
								return true
							}
						}
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
					t.Logf("%s miner-only guard validated via eth_call", name)
				}
			}
			if revision != nil {
				if rev, err := revision(); err == nil && rev.Cmp(big.NewInt(2)) >= 0 {
					t.Logf("%s already at revision %s", name, rev.String())
					return true
				}
			}
			t.Fatalf("%s reinitialize not mined by miner after retries", name)
			return false // unreachable
		}

		proposalDone := reinit("Proposal", ctx.ProposalAddr, contracts.ProposalMetaData, func(opts *bind.TransactOpts, client *ethclient.Client) (*types.Transaction, error) {
			prop, err := contracts.NewProposal(ctx.ProposalAddr, client)
			if err != nil {
				return nil, err
			}
			return prop.ReinitializeV2(opts)
		}, func() (*big.Int, error) {
			return ctx.Proposal.Revision(nil)
		})
		reinit("Validators", testctx.ValidatorsAddr, contracts.ValidatorsMetaData, func(opts *bind.TransactOpts, client *ethclient.Client) (*types.Transaction, error) {
			vals, err := contracts.NewValidators(testctx.ValidatorsAddr, client)
			if err != nil {
				return nil, err
			}
			return vals.ReinitializeV2(opts)
		}, func() (*big.Int, error) {
			return ctx.Validators.Revision(nil)
		})
		reinit("Staking", testctx.StakingAddr, contracts.StakingMetaData, func(opts *bind.TransactOpts, client *ethclient.Client) (*types.Transaction, error) {
			stk, err := contracts.NewStaking(testctx.StakingAddr, client)
			if err != nil {
				return nil, err
			}
			return stk.ReinitializeV2(opts)
		}, func() (*big.Int, error) {
			return ctx.Staking.Revision(nil)
		})
		reinit("Punish", testctx.PunishAddr, contracts.PunishMetaData, func(opts *bind.TransactOpts, client *ethclient.Client) (*types.Transaction, error) {
			pun, err := contracts.NewPunish(testctx.PunishAddr, client)
			if err != nil {
				return nil, err
			}
			return pun.ReinitializeV2(opts)
		}, func() (*big.Int, error) {
			return ctx.Punish.Revision(nil)
		})

		// Second call should fail
		if !proposalDone {
			t.Fatalf("proposal reinitialize not executed; cannot verify second call guard")
		}
		_, minerAddr, minerClient := pickInTurnValidator(t)
		txHash, err := sendLocalReinit(minerClient, minerAddr, ctx.ProposalAddr, contracts.ProposalMetaData)
		if err == nil {
			err = waitReceipt(minerClient, txHash, 30*time.Second)
		}
		if err == nil {
			t.Fatal("Proposal reinitializeV2 should fail on second call")
		}
	})
}

func pickInTurnValidator(t *testing.T) (*ecdsa.PrivateKey, common.Address, *ethclient.Client) {
	if ctx == nil {
		t.Fatalf("Context not initialized")
	}
	validators, err := ctx.Validators.GetActiveValidators(nil)
	if err != nil || len(validators) == 0 {
		t.Fatalf("no active validators available")
	}
	sort.Slice(validators, func(i, j int) bool {
		return bytes.Compare(validators[i][:], validators[j][:]) < 0
	})
	for attempt := 0; attempt < 12; attempt++ {
		header, err := ctx.Clients[0].HeaderByNumber(context.Background(), nil)
		if err != nil || header == nil {
			waitNextBlock()
			continue
		}
		next := header.Number.Uint64() + 1
		idx := next % uint64(len(validators))
		addr := validators[idx]
		key := keyForAddress(addr)
		if key == nil {
			known := make([]string, 0, len(ctx.GenesisValidators))
			for _, k := range ctx.GenesisValidators {
				known = append(known, crypto.PubkeyToAddress(k.PublicKey).Hex())
			}
			t.Fatalf("no key for in-turn validator %s (known=%s)", addr.Hex(), strings.Join(known, ","))
		}
		client := clientForValidator(t, addr)
		if client != nil {
			h2, err := client.HeaderByNumber(context.Background(), nil)
			if err == nil && h2 != nil {
				next2 := h2.Number.Uint64() + 1
				if validators[next2%uint64(len(validators))] == addr {
					return key, addr, client
				}
			}
		}
		waitNextBlock()
	}
	t.Fatalf("no in-turn validator matched across clients")
	return nil, common.Address{}, nil
}

func clientForValidator(t *testing.T, addr common.Address) *ethclient.Client {
	if ctx == nil {
		t.Fatalf("Context not initialized")
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
		t.Fatalf("no validator config available; update test_config.yaml validators list")
	}
	ip := ""
	rpcURL := ""
	for i, v := range ctx.Config.Validators {
		if common.HexToAddress(v.Address) == addr {
			if i < len(ctx.Config.ValidatorRPCs) && ctx.Config.ValidatorRPCs[i] != "" {
				rpcURL = ctx.Config.ValidatorRPCs[i]
			}
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
	if rpcURL == "" && ip != "" {
		rpcURL = "http://" + ip + ":8545"
	}
	if rpcURL == "" {
		t.Fatalf("no RPC mapping for validator %s; ensure validator_rpcs or validators[0..2] match node0-2", addr.Hex())
	}
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		t.Fatalf("failed to dial validator RPC %s: %v", rpcURL, err)
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
