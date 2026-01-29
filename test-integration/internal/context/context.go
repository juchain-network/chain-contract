package context

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	
	"juchain.org/chain/tools/contracts"
	"juchain.org/chain/tools/ci/internal/config"
)

// Addresses of system contracts
var (
	ValidatorsAddr = common.HexToAddress("0x000000000000000000000000000000000000f010")
	PunishAddr     = common.HexToAddress("0x000000000000000000000000000000000000F011")
	ProposalAddr   = common.HexToAddress("0x000000000000000000000000000000000000F012")
	StakingAddr    = common.HexToAddress("0x000000000000000000000000000000000000F013")
)

type CIContext struct {
	Config  *config.Config
	Clients []*ethclient.Client
	ChainID *big.Int

	// System Contracts
	Validators *contracts.Validators
	Punish     *contracts.Punish
	Proposal   *contracts.Proposal
	Staking    *contracts.Staking

	// Accounts
	FunderKey         *ecdsa.PrivateKey
	GenesisValidators []*ecdsa.PrivateKey
	
	mu                sync.Mutex
	nonces            map[common.Address]uint64
	clientIndex       int
	proposerIndex     int
}

func NewCIContext(cfg *config.Config) (*CIContext, error) {
	if len(cfg.RPCs) == 0 {
		return nil, fmt.Errorf("no rpcs provided")
	}

	var clients []*ethclient.Client
	for _, url := range cfg.RPCs {
		client, err := ethclient.Dial(url)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to %s: %w", url, err)
		}
		clients = append(clients, client)
	}

	primaryClient := clients[0]
	chainID, err := primaryClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain id: %w", err)
	}

	val, err := contracts.NewValidators(ValidatorsAddr, primaryClient)
	if err != nil { return nil, err }
	pun, err := contracts.NewPunish(PunishAddr, primaryClient)
	if err != nil { return nil, err }
	prop, err := contracts.NewProposal(ProposalAddr, primaryClient)
	if err != nil { return nil, err }
	stk, err := contracts.NewStaking(StakingAddr, primaryClient)
	if err != nil { return nil, err }

	funderKey, err := crypto.HexToECDSA(cfg.Funder.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid funder private key: %w", err)
	}

	var genesisValidators []*ecdsa.PrivateKey
	for i, v := range cfg.Validators {
		key, err := crypto.HexToECDSA(v.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("invalid validator private key at index %d: %w", i, err)
		}
		genesisValidators = append(genesisValidators, key)
	}

	c := &CIContext{
		Config:            cfg,
		Clients:           clients,
		ChainID:           chainID,
		Validators:        val,
		Punish:            pun,
		Proposal:          prop,
		Staking:           stk,
		FunderKey:         funderKey,
		GenesisValidators: genesisValidators,
		nonces:            make(map[common.Address]uint64),
	}

	// Auto-Initialize if needed
	err = c.autoInitialize()
	if err != nil {
		fmt.Printf("⚠️ autoInitialize failed: %v\n", err)
	}

	return c, nil
}

func (c *CIContext) autoInitialize() error {
	// Robust check: if MinValidatorStake is default (100k JU), we need setup.
	minStake, err := c.Proposal.MinValidatorStake(nil)
	if err == nil && minStake.Cmp(big.NewInt(1000000000000000000)) == 0 {
		fmt.Printf("✅ System already configured (MinValidatorStake = 1 JU).\n")
		return nil
	}

	fmt.Printf("🔧 System unconfigured (MinValidatorStake = %v), performing auto-initialization...\n", minStake)
	
	// Check if we need to call initialize() at all
	initialized, _ := c.Proposal.Initialized(nil)
	if !initialized {
		var valAddrs []common.Address
		for _, vk := range c.GenesisValidators {
			valAddrs = append(valAddrs, crypto.PubkeyToAddress(vk.PublicKey))
		}

		// 1. Initialize Proposal
		opts, _ := c.GetTransactor(c.GenesisValidators[0])
		fmt.Printf("  > Initializing Proposal...\n")
		tx, err := c.Proposal.Initialize(opts, valAddrs, ValidatorsAddr, big.NewInt(20))
		if err == nil {
			c.WaitMined(tx.Hash())
		}

		// 2. Initialize Staking with Validators
		opts, _ = c.GetTransactor(c.GenesisValidators[1])
		fmt.Printf("  > Initializing Staking with Validators...\n")
		tx, err = c.Staking.InitializeWithValidators(opts, ValidatorsAddr, ProposalAddr, PunishAddr, valAddrs, big.NewInt(1000))
		if err == nil {
			c.WaitMined(tx.Hash())
		}

		// 3. Initialize Validators
		opts, _ = c.GetTransactor(c.GenesisValidators[2])
		fmt.Printf("  > Initializing Validators...\n")
		tx, err = c.Validators.Initialize(opts, valAddrs, ProposalAddr, PunishAddr, StakingAddr)
		if err == nil {
			c.WaitMined(tx.Hash())
		}
	}

	// 4. Always ensure test-friendly parameters if they are not set
	fmt.Printf("  > Configuring system parameters...\n")
	_ = c.EnsureConfig(19, big.NewInt(1), nil)                   // ProposalCooldown
	_ = c.EnsureConfig(8, big.NewInt(1000000000000000000), nil)  // MinValidatorStake
	_ = c.EnsureConfig(10, big.NewInt(1000000000000000000), nil) // MinDelegation
	
	fmt.Printf("✅ Auto-initialization complete.\n")
	return nil
}

func (c *CIContext) GetTransactor(key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	if key == nil { return nil, fmt.Errorf("nil private key") }
	
c.mu.Lock()
	defer c.mu.Unlock()

	addr := crypto.PubkeyToAddress(key.PublicKey)
	
	var maxNonce uint64
	for _, client := range c.Clients {
		n, err := client.PendingNonceAt(context.Background(), addr)
		if err == nil && n > maxNonce {
			maxNonce = n
		}
	}

	if cached, ok := c.nonces[addr]; ok && cached >= maxNonce {
		maxNonce = cached
	}
	c.nonces[addr] = maxNonce + 1

	opts, err := bind.NewKeyedTransactorWithChainID(key, c.ChainID)
	if err != nil {
		return nil, err
	}
	
	opts.Nonce = big.NewInt(int64(maxNonce))
	opts.GasLimit = 20000000 
	opts.GasPrice = big.NewInt(1000000000) 
	return opts, nil
}

func (c *CIContext) CreateAndFundAccount(amount *big.Int) (*ecdsa.PrivateKey, common.Address, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, common.Address{}, err
	}
	addr := crypto.PubkeyToAddress(key.PublicKey)

	opts, err := c.GetTransactor(c.FunderKey)
	if err != nil { return nil, common.Address{}, err }

	tx := types.NewTransaction(opts.Nonce.Uint64(), addr, amount, 21000, opts.GasPrice, nil)
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(c.ChainID), c.FunderKey)
	if err != nil { return nil, common.Address{}, err }

	for _, client := range c.Clients {
		_ = client.SendTransaction(context.Background(), signedTx)
	}

	log.Info("Funded account", "address", addr.Hex(), "tx", signedTx.Hash().Hex())
	
	if err := c.WaitMined(signedTx.Hash()); err != nil {
        return nil, common.Address{}, fmt.Errorf("funding tx failed: %w", err)
    }

	time.Sleep(2 * time.Second)

	return key, addr, nil
}

func (c *CIContext) WaitMined(txHash common.Hash) error {
	if txHash == (common.Hash{}) { return nil }
	
	ctx, cancel := context.WithTimeout(context.Background(), 120*time.Second)
	defer cancel()
	
	queryTicker := time.NewTicker(1 * time.Second)
	defer queryTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for tx %s", txHash.Hex())
		case <-queryTicker.C:
			fmt.Print(".")
			
			for i, client := range c.Clients {
				receipt, err := client.TransactionReceipt(context.Background(), txHash)
				if err == nil && receipt != nil {
					fmt.Println()
					if receipt.Status == 1 {
						return nil
					}
					// Revert reason
					tx, _, err := client.TransactionByHash(context.Background(), txHash)
					if err == nil {
						from, _ := types.Sender(types.LatestSignerForChainID(c.ChainID), tx)
						msg := ethereum.CallMsg{
							From:     from,
							To:       tx.To(),
							Gas:      tx.Gas(),
							GasPrice: tx.GasPrice(),
							Value:    tx.Value(),
							Data:     tx.Data(),
						}
						_, errCall := client.CallContract(context.Background(), msg, receipt.BlockNumber)
						if errCall != nil {
							return fmt.Errorf("transaction %s failed on node %d: %v", txHash.Hex(), i, errCall)
						}
					}
					return fmt.Errorf("transaction %s failed on node %d (status 0)", txHash.Hex(), i)
				}
			}
		}
	}
}

func (c *CIContext) EnsureConfig(cid int64, targetVal *big.Int, currentVal *big.Int) error {
	if currentVal != nil && currentVal.Cmp(targetVal) == 0 {
		return nil
	}
	
	log.Info("Updating config", "cid", cid, "target", targetVal, "current", currentVal)
	
	if len(c.GenesisValidators) == 0 { return fmt.Errorf("no genesis validators") }
	
	var err error
	for i := 0; i < len(c.GenesisValidators); i++ {
		c.mu.Lock()
		proposerKey := c.GenesisValidators[c.proposerIndex % len(c.GenesisValidators)]
		c.proposerIndex++
		c.mu.Unlock()

		opts, errG := c.GetTransactor(proposerKey)
		if errG != nil { continue }
		
		tx, errCall := c.Proposal.CreateUpdateConfigProposal(opts, big.NewInt(cid), targetVal)
		if errCall == nil {
			err = c.WaitMined(tx.Hash())
			if err != nil { return err }
			
			receipt, _ := c.Clients[0].TransactionReceipt(context.Background(), tx.Hash())
			var propID [32]byte
			found := false
			for _, l := range receipt.Logs {
				if ev, err := c.Proposal.ParseLogCreateConfigProposal(*l); err == nil {
					propID = ev.Id
					found = true
					break
				}
			}
			if !found { return fmt.Errorf("proposal log not found for tx %s", tx.Hash().Hex()) }

			for _, vk := range c.GenesisValidators {
				voterAddr := crypto.PubkeyToAddress(vk.PublicKey)
				exist, _ := c.Validators.IsValidatorExist(nil, voterAddr)
				if !exist { continue }
				
				vo, _ := c.GetTransactor(vk)
				txV, errV := c.Proposal.VoteProposal(vo, propID, true)
				if errV == nil {
					c.WaitMined(txV.Hash())
				}
			}
			return nil
		}
		
		err = errCall
		if strings.Contains(err.Error(), "Proposal creation too frequent") {
			continue 
		}
		return fmt.Errorf("createUpdateConfigProposal failed: %w", err)
	}
	return fmt.Errorf("all proposers in cooldown: %w", err)
}
