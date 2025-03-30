package cmd

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/assert"
)

func TestCreaetProposalTx(t *testing.T) {
	proposer := "0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b"
	target := "0x029DAB47e268575D4AC167De64052FB228B5fA41"
	rpc := "https://testnet-rpc.juchain.org"
	innerCreateProposal(proposer, target, false, rpc)
}

func TestSignTx(t *testing.T) {
	file := "createProposal.json"
	key := "0xca881281fb10b53a87d00cbfae29f7cf8cfe8ac7c8389b3d20b24fc6bc3f3ff9"
	chainId := 202599
	key = strings.TrimPrefix(key, "0x")
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		fmt.Printf("invalid private key: %v", err)
		return
	}
	innerSignRawTx(int64(chainId), file, privateKey)
}

func TestSendTx(t *testing.T) {
	proposer := "0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b"
	file := "createProposal_signed.json"
	rpc := "https://testnet-rpc.juchain.org"
	innerSendSignedTx(proposer, file, rpc)
}

func TestBuldProposalId(t *testing.T) {
	// Wait for tx to be finished executing with hash 0xb75dc353e433ce38edb359ae9aa9f88fa265ff9436fac164b6afb97f0aa87795
	// tx confirmed in block 776421
	// Proposal ID: 13013b639ad153f5207ec6b0aa168b142168cddba47e077f91df7c40aaba44b8
	// Proposer: 0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b
	// Destination: 0x029DAB47e268575D4AC167De64052FB228B5fA41
	// Flag: true
	// Time: 1743150640
	proposer := "0x016103822e9a3425DfeaFDCd57c9F7fC2bA72a8b"
	target := "0x029DAB47e268575D4AC167De64052FB228B5fA41"
	rpc := "https://testnet-rpc.juchain.org"
	blockNum := 776421

	client, err := ethclient.Dial(rpc)
	if err != nil {
		fmt.Printf("failed to connect to RPC: %v", err)
		return
	}
	defer client.Close()
	proposalId, err := BuildProposalId(proposer, target, true, "", uint64(blockNum), client)
	if err != nil {
		fmt.Printf("failed to build proposal id: %v", err)
		return
	}
	fmt.Printf("build proposal id %s \n", proposalId)
	expected := "13013b639ad153f5207ec6b0aa168b142168cddba47e077f91df7c40aaba44b8"
	assert.Equal(t, expected, proposalId, "Proposal ID mismatch")
}
