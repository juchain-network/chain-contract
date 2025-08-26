package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	configFile string
)

// Config represents the application configuration
type Config struct {
	RPCEndpoint string `json:"rpc_endpoint"`
	ChainID     int64  `json:"chain_id"`
}

// Default configuration values
var defaultConfig = Config{
	RPCEndpoint: "http://localhost:8545",
	ChainID:     2025,
}

// getConfigPath returns the configuration file path
func getConfigPath() (string, error) {
	if configFile != "" {
		return configFile, nil
	}

	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	configDir := filepath.Join(home, ".congress-cli")
	if err := os.MkdirAll(configDir, 0755); err != nil {
		return "", err
	}

	return filepath.Join(configDir, "config.json"), nil
}

// loadConfig loads configuration from file
func loadConfig() (*Config, error) {
	configPath, err := getConfigPath()
	if err != nil {
		return &defaultConfig, err
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			// Config file doesn't exist, return default config
			return &defaultConfig, nil
		}
		return &defaultConfig, err
	}

	var config Config
	if err := json.Unmarshal(data, &config); err != nil {
		return &defaultConfig, err
	}

	return &config, nil
}

// saveConfig saves configuration to file
func saveConfig(config *Config) error {
	configPath, err := getConfigPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, data, 0644)
}

// ConfigCmd creates the main config command
func ConfigCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "config",
		Short: "Manage congress-cli configuration",
		Long:  "Set and get configuration values for congress-cli including RPC endpoint and chain ID",
	}

	cmd.AddCommand(
		configSetCmd(),
		configGetCmd(),
		configListCmd(),
	)

	return cmd
}

// configSetCmd creates command for setting configuration values
func configSetCmd() *cobra.Command {
	var rpcEndpoint string
	var chainID int64

	cmd := &cobra.Command{
		Use:   "set",
		Short: "Set configuration values",
		Long:  "Set configuration values like RPC endpoint and chain ID",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := loadConfig()
			if err != nil {
				fmt.Printf("Error loading config: %v\n", err)
				return
			}

			changed := false
			if cmd.Flags().Changed("rpc") {
				config.RPCEndpoint = rpcEndpoint
				fmt.Printf("✅ RPC endpoint set to: %s\n", rpcEndpoint)
				changed = true
			}
			if cmd.Flags().Changed("chain-id") {
				config.ChainID = chainID
				fmt.Printf("✅ Chain ID set to: %d\n", chainID)
				changed = true
			}

			if !changed {
				fmt.Println("No configuration values provided. Use --rpc or --chain-id flags.")
				return
			}

			if err := saveConfig(config); err != nil {
				fmt.Printf("Error saving config: %v\n", err)
				return
			}
		},
	}

	cmd.Flags().StringVar(&rpcEndpoint, "rpc", "", "RPC endpoint URL")
	cmd.Flags().Int64Var(&chainID, "chain-id", 0, "Chain ID")

	return cmd
}

// configGetCmd creates command for getting configuration values
func configGetCmd() *cobra.Command {
	var getRpc bool
	var getChainID bool

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get configuration values",
		Long:  "Get current configuration values",
		Run: func(cmd *cobra.Command, args []string) {
			config, err := loadConfig()
			if err != nil {
				fmt.Printf("Error loading config: %v\n", err)
				return
			}

			if getRpc {
				fmt.Printf("RPC endpoint: %s\n", config.RPCEndpoint)
			}
			if getChainID {
				fmt.Printf("Chain ID: %d\n", config.ChainID)
			}
			if !getRpc && !getChainID {
				// Show all config if no specific flags
				fmt.Printf("RPC endpoint: %s\n", config.RPCEndpoint)
				fmt.Printf("Chain ID: %d\n", config.ChainID)
			}
		},
	}

	cmd.Flags().BoolVar(&getRpc, "rpc", false, "Get RPC endpoint")
	cmd.Flags().BoolVar(&getChainID, "chain-id", false, "Get chain ID")

	return cmd
}

// configListCmd creates command for listing all configuration
func configListCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List all configuration values",
		Long:  "Display all current configuration values",
		Run:   configList,
	}

	return cmd
}

func configList(cmd *cobra.Command, args []string) {
	config, err := loadConfig()
	if err != nil {
		PrintWarning(fmt.Sprintf("Could not load config: %v", err))
		config = &defaultConfig
	}

	PrintSuccess("Current Configuration")
	fmt.Printf("RPC endpoint: %s\n", config.RPCEndpoint)
	fmt.Printf("Chain ID: %d\n", config.ChainID)

	// Show config file location
	configPath, err := getConfigPath()
	if err == nil {
		fmt.Printf("Config file: %s\n", configPath)
	}
}

// GetRPCEndpoint gets RPC endpoint with fallback logic
func GetRPCEndpoint(cmd *cobra.Command) string {
	// First check command line flags
	if cmd.Flags().Changed("rpc_laddr") {
		value, _ := cmd.Flags().GetString("rpc_laddr")
		return value
	}

	// Load from config file
	config, err := loadConfig()
	if err != nil {
		return defaultConfig.RPCEndpoint
	}

	return config.RPCEndpoint
}

// GetChainID gets chain ID with fallback logic
func GetChainID(cmd *cobra.Command) int64 {
	// First check command line flags
	if cmd.Flags().Changed("chainId") {
		value, _ := cmd.Flags().GetInt64("chainId")
		return value
	}

	// Load from config file
	config, err := loadConfig()
	if err != nil {
		return defaultConfig.ChainID
	}

	return config.ChainID
}
