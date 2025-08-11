package cmd

import (
	"testing"

	"github.com/ethereum/go-ethereum/core"
	"github.com/stretchr/testify/assert"
)

// TestHashGeneration tests hash generation with default mainnet configuration
func TestHashGeneration(t *testing.T) {
	t.Log("Testing hash generation with default mainnet configuration...")

	// Create genesis with MainnetChainConfig
	genesis := core.DefaultGenesisBlock()

	// Generate the genesis block and get its hash
	block := genesis.ToBlock()
	hash := block.Hash()

	t.Logf("Generated hash: %s", hash.Hex())

	// The actual mainnet genesis hash
	expectedHash := "0xd4e56740f876aef8c010b86a40d5f56745a118d0906a34e69aec8c0db1cb8fa3"
	assert.Equal(t, expectedHash, hash.Hex(), "Generated hash should match mainnet genesis hash")

	// Also check the config
	config := genesis.Config
	t.Logf("ChainID: %d", config.ChainID)
	t.Logf("BerlinBlock: %v", config.BerlinBlock)
	t.Logf("LondonBlock: %v", config.LondonBlock)

	if config.ShanghaiTime != nil {
		t.Logf("ShanghaiTime: %d", *config.ShanghaiTime)
	} else {
		t.Logf("ShanghaiTime: nil")
	}

	if config.CancunTime != nil {
		t.Logf("CancunTime: %d", *config.CancunTime)
	} else {
		t.Logf("CancunTime: nil")
	}

	// Verify configuration values
	assert.NotNil(t, config.ChainID, "ChainID should not be nil")
	assert.Equal(t, int64(1), config.ChainID.Int64(), "ChainID should be 1 for mainnet")
}

// TestGenesisBlockStructure tests the structure of the genesis block
func TestGenesisBlockStructure(t *testing.T) {
	genesis := core.DefaultGenesisBlock()
	block := genesis.ToBlock()

	// Test basic block properties
	assert.NotNil(t, block, "Block should not be nil")
	assert.NotNil(t, block.Hash(), "Block hash should not be nil")
	assert.NotNil(t, block.Header(), "Block header should not be nil")

	// Test header properties
	header := block.Header()
	assert.Equal(t, uint64(0), header.Number.Uint64(), "Genesis block number should be 0")
	assert.NotEmpty(t, header.Root.Hex(), "State root should not be empty")

	t.Logf("Block Number: %d", header.Number.Uint64())
	t.Logf("State Root: %s", header.Root.Hex())
	t.Logf("Block Hash: %s", block.Hash().Hex())
}

// TestChainConfigValues tests specific chain configuration values
func TestChainConfigValues(t *testing.T) {
	genesis := core.DefaultGenesisBlock()
	config := genesis.Config

	// Test that BerlinBlock and LondonBlock are properly set in mainnet config
	assert.NotNil(t, config.BerlinBlock, "BerlinBlock should be set in mainnet config")
	assert.NotNil(t, config.LondonBlock, "LondonBlock should be set in mainnet config")

	// Test other important config values
	assert.NotNil(t, config.ChainID, "ChainID should not be nil")
	assert.NotNil(t, config.HomesteadBlock, "HomesteadBlock should not be nil")

	// Verify expected block numbers
	assert.Equal(t, uint64(12244000), config.BerlinBlock.Uint64(), "Berlin block should be 12244000")
	assert.Equal(t, uint64(12965000), config.LondonBlock.Uint64(), "London block should be 12965000")

	t.Logf("Homestead Block: %v", config.HomesteadBlock)
	t.Logf("DAO Fork Block: %v", config.DAOForkBlock)
	t.Logf("EIP150 Block: %v", config.EIP150Block)
	t.Logf("EIP155 Block: %v", config.EIP155Block)
	t.Logf("EIP158 Block: %v", config.EIP158Block)
	t.Logf("Byzantium Block: %v", config.ByzantiumBlock)
	t.Logf("Constantinople Block: %v", config.ConstantinopleBlock)
	t.Logf("Petersburg Block: %v", config.PetersburgBlock)
	t.Logf("Istanbul Block: %v", config.IstanbulBlock)
	t.Logf("Muir Glacier Block: %v", config.MuirGlacierBlock)
	t.Logf("Berlin Block: %v", config.BerlinBlock)
	t.Logf("London Block: %v", config.LondonBlock)
}
