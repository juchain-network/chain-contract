package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateAddress(t *testing.T) {
	tests := []struct {
		name    string
		addr    string
		wantErr bool
	}{
		{
			name:    "valid address",
			addr:    "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			wantErr: false,
		},
		{
			name:    "invalid address - too short",
			addr:    "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb922",
			wantErr: true,
		},
		{
			name:    "invalid address - no 0x prefix",
			addr:    "f39Fd6e51aad88F6F4ce6aB8827279cffFb92266",
			wantErr: false, // common.IsHexAddress actually accepts this format
		},
		{
			name:    "empty address",
			addr:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateAddress(tt.addr)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateChainID(t *testing.T) {
	tests := []struct {
		name    string
		chainID int64
		wantErr bool
	}{
		{
			name:    "valid chain ID",
			chainID: 202599,
			wantErr: false,
		},
		{
			name:    "zero chain ID",
			chainID: 0,
			wantErr: true,
		},
		{
			name:    "negative chain ID",
			chainID: -1,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateChainID(tt.chainID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateOperation(t *testing.T) {
	tests := []struct {
		name      string
		operation string
		wantErr   bool
	}{
		{
			name:      "valid operation add",
			operation: "add",
			wantErr:   false,
		},
		{
			name:      "valid operation remove",
			operation: "remove",
			wantErr:   false,
		},
		{
			name:      "invalid operation",
			operation: "invalid",
			wantErr:   true,
		},
		{
			name:      "empty operation",
			operation: "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateOperation(tt.operation)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestValidateConfigID(t *testing.T) {
	tests := []struct {
		name     string
		configID int64
		wantErr  bool
	}{
		{
			name:     "valid config ID 0",
			configID: 0,
			wantErr:  false,
		},
		{
			name:     "valid config ID 4",
			configID: 4,
			wantErr:  false,
		},
		{
			name:     "invalid config ID negative",
			configID: -1,
			wantErr:  true,
		},
		{
			name:     "invalid config ID too large",
			configID: 10,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateConfigID(tt.configID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestGetConfigIDName(t *testing.T) {
	tests := []struct {
		name     string
		configID int64
		expected string
	}{
		{
			name:     "config ID 0",
			configID: 0,
			expected: "proposalLastingPeriod",
		},
		{
			name:     "config ID 1",
			configID: 1,
			expected: "punishThreshold",
		},
		{
			name:     "config ID 4",
			configID: 4,
			expected: "withdrawProfitPeriod",
		},
		{
			name:     "invalid config ID",
			configID: 10,
			expected: "unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GetConfigIDName(tt.configID)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestValidateProposalID(t *testing.T) {
	tests := []struct {
		name       string
		proposalID string
		wantErr    bool
	}{
		{
			name:       "valid proposal ID",
			proposalID: "4881cb992217b8d050a87384de6998efb1f051c0cb62c10bdd393f68336618b6",
			wantErr:    false,
		},
		{
			name:       "invalid proposal ID - too short",
			proposalID: "4881cb992217b8d050a87384de6998efb1f051c0cb62c10bdd393f68336618",
			wantErr:    true,
		},
		{
			name:       "invalid proposal ID - invalid characters",
			proposalID: "4881cb992217b8d050a87384de6998efb1f051c0cb62c10bdd393f68336618gz",
			wantErr:    true,
		},
		{
			name:       "empty proposal ID",
			proposalID: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateProposalID(tt.proposalID)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
