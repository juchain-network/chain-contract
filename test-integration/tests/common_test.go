package tests

import (
	"flag"
	"os"
	"testing"

	"github.com/ethereum/go-ethereum/log"
	
	"juchain.org/chain/tools/ci/internal/config"
	"juchain.org/chain/tools/ci/internal/context"
)

var (
	ctx *context.CIContext
	configPath = flag.String("config", "../config.yaml", "Path to test configuration file")
)

func TestMain(m *testing.M) {
	// Parse flags
	flag.Parse()

	// Initialize logger
	log.SetDefault(log.NewLogger(log.NewTerminalHandlerWithLevel(os.Stderr, log.LevelInfo, true)))

	// Load config
	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		// Fallback to example if not found (for CI environment without real secrets)
		// But in real run, we fail.
		log.Error("Failed to load config", "err", err)
		os.Exit(1)
	}

	// Init context
	c, err := context.NewCIContext(cfg)
	if err != nil {
		log.Error("Failed to init context", "err", err)
		os.Exit(1)
	}
	ctx = c

	os.Exit(m.Run())
}
