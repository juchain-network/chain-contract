package context

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	
	"juchain.org/chain/tools/contracts"
	"juchain.org/chain/tools/ci/internal/config"
)

// Addresses of system contracts (Hardcoded for now, or load from somewhere)
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

	// System Contracts (bound to the first client)
	Validators *contracts.Validators
	Punish     *contracts.Punish
	Proposal   *contracts.Proposal
	Staking    *contracts.Staking

	// Accounts
	FunderKey         *ecdsa.PrivateKey
	GenesisValidators []*ecdsa.PrivateKey
	mu                sync.Mutex
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

	// Use the first client for general queries
	primaryClient := clients[0]
	chainID, err := primaryClient.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain id: %w", err)
	}

	// Init contracts
	val, err := contracts.NewValidators(ValidatorsAddr, primaryClient)
	if err != nil { return nil, err }
	pun, err := contracts.NewPunish(PunishAddr, primaryClient)
	if err != nil { return nil, err }
	prop, err := contracts.NewProposal(ProposalAddr, primaryClient)
	if err != nil { return nil, err }
	stk, err := contracts.NewStaking(StakingAddr, primaryClient)
	if err != nil { return nil, err }

	// Parse funder key
	funderKey, err := crypto.HexToECDSA(cfg.Funder.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("invalid funder private key: %w", err)
	}

	// Parse genesis validators keys
	var genesisValidators []*ecdsa.PrivateKey
	for i, v := range cfg.Validators {
		key, err := crypto.HexToECDSA(v.PrivateKey)
		if err != nil {
			return nil, fmt.Errorf("invalid validator private key at index %d: %w", i, err)
		}
		genesisValidators = append(genesisValidators, key)
	}

	return &CIContext{
		Config:            cfg,
		Clients:           clients,
		ChainID:           chainID,
		Validators:        val,
		Punish:            pun,
		Proposal:          prop,
		Staking:           stk,
		FunderKey:         funderKey,
		GenesisValidators: genesisValidators,
	}, nil
}

// GetTransactor returns a bind.TransactOpts for the given private key
func (c *CIContext) GetTransactor(key *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(key, c.ChainID)
	if err != nil {
		return nil, err
	}
	
	// Simply allow gas estimation to handle it, or set a high limit
	// opts.GasLimit = 5000000 
	return opts, nil
}

// CreateAccount generates a new random account and funds it
func (c *CIContext) CreateAndFundAccount(amount *big.Int) (*ecdsa.PrivateKey, common.Address, error) {
	key, err := crypto.GenerateKey()
	if err != nil {
		return nil, common.Address{}, err
	}
	addr := crypto.PubkeyToAddress(key.PublicKey)

	// Fund it
	funderOpts, err := c.GetTransactor(c.FunderKey)
	if err != nil {
		return nil, common.Address{}, err
	}
	// Important: We need to manage nonce manually if concurrent requests happen, 
	// but for simplicity here we rely on pending state or simple blocking.
	// For robustness in CI, we should probably lock the funder.
	c.mu.Lock()
	defer c.mu.Unlock()

	nonce, err := c.Clients[0].PendingNonceAt(context.Background(), crypto.PubkeyToAddress(c.FunderKey.PublicKey))
	if err != nil {
		return nil, common.Address{}, err
	}
	funderOpts.Nonce = big.NewInt(int64(nonce))
	funderOpts.Value = amount

	// Just a simple transfer? TransactOpts is usually for contract calls.
	// For simple transfer, we use raw types.Transaction, but to keep it simple, 
	// let's assume we might have a helper or just use one of the clients.
	// ACTUALLY: bind.TransactOpts doesn't expose a "SendTransaction" method for ETH transfers.
	// We have to use the client.
	
	// Let's implement transfer using ethclient
	// Gas Limit 21000
	gasLimit := uint64(21000)
	gasPrice, err := c.Clients[0].SuggestGasPrice(context.Background())
	if err != nil {
		return nil, common.Address{}, err
	}

	tx := types.NewTransaction(nonce, addr, amount, gasLimit, gasPrice, nil)
	
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(c.ChainID), c.FunderKey)
	if err != nil {
		return nil, common.Address{}, err
	}

	err = c.Clients[0].SendTransaction(context.Background(), signedTx)
	if err != nil {
		return nil, common.Address{}, err
	}

	log.Info("Funded account", "address", addr.Hex(), "tx", signedTx.Hash().Hex())
	
	// Wait for receipt?
	// In a real CI, we might want to wait.
	if err := c.WaitMined(signedTx.Hash()); err != nil {
        return nil, common.Address{}, fmt.Errorf("funding tx failed: %w", err)
    }

	return key, addr, nil
}

// WaitMined waits for a tx to be mined
func (c *CIContext) WaitMined(txHash common.Hash) error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	queryTicker := time.NewTicker(1 * time.Second)
	defer queryTicker.Stop()

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("timeout waiting for tx %s", txHash.Hex())
		case <-queryTicker.C:
			receipt, err := c.Clients[0].TransactionReceipt(context.Background(), txHash)
			if err == nil && receipt != nil {
				if receipt.Status == 1 {
					return nil
				}
				return fmt.Errorf("transaction failed (status 0)")
			}
		}
	}
}
