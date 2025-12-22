package cmd

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
	"juchain.org/chain/tools/contracts"
)

func CreateRawTx(
	caller common.Address,
	contract common.Address,
	value *big.Int,
	data []byte,
	rpc string,
	output string,
) error {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		return fmt.Errorf("failed to connect to RPC: %v", err)
	}
	defer client.Close()

	nonce, err := client.PendingNonceAt(context.Background(), caller)
	if err != nil {
		return fmt.Errorf("failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("SuggestGasPrice Err:", err)
		return nil
	}

	msg := ethereum.CallMsg{
		From:  caller,
		To:    &contract,
		Data:  data,
		Value: value,
	}
	gasLimit, err := client.EstimateGas(context.Background(), msg)
	if err != nil {
		return fmt.Errorf("gas estimation failed: %v", err)
	}
	gasLimit = gasLimit * DefaultGasMultiplier / 100

	tx := types.NewTransaction(
		nonce,
		contract,
		value,
		gasLimit,
		gasPrice,
		data,
	)

	rawTx := map[string]interface{}{
		"nonce":    tx.Nonce(),
		"gasPrice": tx.GasPrice(),
		"gasLimit": tx.Gas(),
		"to":       tx.To().Hex(),
		"value":    tx.Value(),
		"data":     common.Bytes2Hex(tx.Data()),
	}

	file, err := os.Create(output)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(rawTx)
}

func SignRawTx(
	inputFile string,
	privateKey *ecdsa.PrivateKey,
	chainID *big.Int,
	outputFile string,
) error {
	var rawTx map[string]interface{}
	file, err := os.Open(inputFile)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()
	if err := json.NewDecoder(file).Decode(&rawTx); err != nil {
		return fmt.Errorf("invalid JSON: %v", err)
	}

	// Handle value conversion safely for large numbers
	var value *big.Int
	switch v := rawTx["value"].(type) {
	case float64:
		// For large numbers, we need to handle them carefully
		valueStr := fmt.Sprintf("%.0f", v)
		value = new(big.Int)
		value.SetString(valueStr, 10)
	case string:
		value = new(big.Int)
		value.SetString(v, 10)
	default:
		return fmt.Errorf("invalid value type: %T", v)
	}

	tx := types.NewTransaction(
		uint64(rawTx["nonce"].(float64)),
		common.HexToAddress(rawTx["to"].(string)),
		value,
		uint64(rawTx["gasLimit"].(float64)),
		big.NewInt(int64(rawTx["gasPrice"].(float64))),
		common.Hex2Bytes(rawTx["data"].(string)),
	)

	signer := types.NewEIP155Signer(chainID)
	signedTx, err := types.SignTx(tx, signer, privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign: %v", err)
	}

	signedData, err := signedTx.MarshalJSON()
	if err != nil {
		return fmt.Errorf("failed to encode signed tx: %v", err)
	}
	return os.WriteFile(outputFile, signedData, 0644)
}

func SendSignedTx(rpcURL string, signedTxFile string) (common.Hash, error) {
	client, err := ethclient.Dial(rpcURL)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to connect to RPC: %v", err)
	}
	defer client.Close()

	// Read signed tx
	data, err := os.ReadFile(signedTxFile)
	if err != nil {
		return common.Hash{}, fmt.Errorf("failed to read file: %v", err)
	}

	var tx types.Transaction
	if err := tx.UnmarshalJSON(data); err != nil {
		return common.Hash{}, fmt.Errorf("invalid signed tx: %v", err)
	}

	tx.ChainId()

	sender, err := types.Sender(types.NewEIP155Signer(tx.ChainId()), &tx)
	if err != nil {
		return common.Hash{}, fmt.Errorf("invalid signed tx: %v", err)
	}
	// Broadcast
	err = client.SendTransaction(context.Background(), &tx)
	if err != nil {
		fmt.Printf("send tx error %v\n", err)
		return common.Hash{}, err
	}

	err, blockHeight := waitEthTxFinished(client, tx.Hash())
	if err != nil {
		return tx.Hash(), err
	}
	time.Sleep(3 * time.Second)
	proposer := sender.Hex()
	fmt.Printf("read sender from signed tx is %s\n", proposer)
	err, _ = QueryProposalId(blockHeight.Uint64(), proposer, client)
	return tx.Hash(), err
}

func waitEthTxFinished(client *ethclient.Client, txhash common.Hash) (error, *big.Int) {
	PrintInfo(fmt.Sprintf("Waiting for transaction confirmation: %s", txhash.String()))
	timeout := time.NewTimer(DefaultTimeout * time.Second)
	ticker := time.NewTicker(DefaultCheckInterval * time.Second)
	defer timeout.Stop()
	defer ticker.Stop()

	for {
		select {
		case <-timeout.C:
			return errors.New("transaction confirmation timeout"), nil
		case <-ticker.C:
			receipt, err := client.TransactionReceipt(context.Background(), txhash)
			if err == ethereum.NotFound {
				continue
			} else if err != nil {
				return err, nil
			}
			PrintSuccess(fmt.Sprintf("Transaction confirmed in block %v", receipt.BlockNumber))
			return nil, receipt.BlockNumber
		}
	}
}

// Build proposal ID
// flag true to add candidate, false to remove candidate
func BuildProposalId(from, dest string, flag bool, details string, blockNum uint64, client *ethclient.Client) (string, error) {
	sender := common.HexToAddress(from) // Proposer
	dst := common.HexToAddress(dest)    // Candidate
	flagValue := big.NewInt(0)
	if flag {
		flagValue.SetInt64(1)
	}
	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
	if err != nil {
		return "", err
	}

	// // Calculate proposal ID (equivalent to Solidity keccak256(abi.encodePacked(...)))
	// id := crypto.Keccak256(
	// 	sender.Bytes(),
	// 	dst.Bytes(),
	// 	common.LeftPadBytes(flagValue.Bytes(), 32), // Pad flag to 32 bytes
	// 	[]byte(nil),
	// 	common.LeftPadBytes(big.NewInt(int64(block.Header().Time)).Bytes(), 32), // timestamp (uint256)
	// )
	// return common.BytesToHash(id).Hex(), nil

	return buildId(sender, dst, flag, "", int64(block.Header().Time)), nil

}

func buildId(
	sender common.Address,
	dst common.Address,
	flag bool,
	details string,
	timestamp int64,
) string {
	// Pack arguments in the same order as Solidity
	data := append(
		sender.Bytes(),
		dst.Bytes()...,
	)
	if flag {
		data = append(data, byte(1))
	} else {
		data = append(data, byte(0))
	}
	data = append(data, []byte(details)...)
	data = append(data, big.NewInt(timestamp).Bytes()...)

	// Compute Keccak-256 hash (Ethereum's custom SHA-3 variant)
	hash := sha3.NewLegacyKeccak256()
	hash.Write(data)
	id := hash.Sum(nil)

	return hex.EncodeToString(id)

	// // Compute Keccak-256 hash (not SHA-3!)
	// hash := crypto.Keccak256(data)
	// return hex.EncodeToString(hash) // Returns bytes32 as hex string
}

// Query contracts proposal ID
func QueryProposalId(blockHeight uint64, proposer string, client *ethclient.Client) (error, string) {
	instance, err := contracts.NewProposal(common.HexToAddress(ProposalContractAddr), client)
	if err != nil {
		fmt.Printf("Failed to instantiate contract: %v", err)
		return err, ""
	}

	// Define query filter
	filterOpts := &bind.FilterOpts{
		Start:   blockHeight,  // Starting block number
		End:     &blockHeight, // End block number (nil means latest block)
		Context: context.Background(),
	}
	// Query event logs
	logs, err := instance.FilterLogCreateProposal(filterOpts, nil, []common.Address{common.HexToAddress(proposer)}, nil)
	if err != nil {
		fmt.Printf("Failed to filter LogCreateProposal events: %v", err)
		return err, ""
	}
	// Iterate through logs
	var proposalId string
	for logs.Next() {
		event := logs.Event
		proposalId = hex.EncodeToString(event.Id[:])
		fmt.Println("--------CreateProposal----------")
		fmt.Printf("Proposal ID: %s\n", proposalId)
		fmt.Printf("Proposer: %s\n", event.Proposer.Hex())
		fmt.Printf("Destination: %s\n", event.Dst.Hex())
		fmt.Printf("Flag: %v\n", event.Flag)
		fmt.Printf("Time: %d\n", event.Time)
		fmt.Printf("Block: %d\n", event.Raw.BlockNumber)
		fmt.Println("-----")
	}
	if logs.Error() != nil {
		fmt.Printf("Error reading logs: %v", logs.Error())
		return logs.Error(), ""
	}

	// Query event logs
	logs1, err := instance.FilterLogCreateConfigProposal(filterOpts, nil, []common.Address{common.HexToAddress(proposer)})
	if err != nil {
		fmt.Printf("Failed to filter LogCreateConfigProposal events: %v", err)
		return err, ""
	}
	// Iterate through logs
	proposalId = "0x"
	for logs1.Next() {
		event := logs1.Event
		proposalId = hex.EncodeToString(event.Id[:])
		fmt.Println("--------CreateConfigProposal----------")
		fmt.Printf("Proposal ID: %s\n", proposalId)
		fmt.Printf("Proposer: %s\n", event.Proposer.Hex())
		fmt.Printf("CID: %s\n", event.Cid)
		fmt.Printf("NewValue: %v\n", event.NewValue)
		fmt.Printf("Time: %d\n", event.Time)
		fmt.Printf("Block: %d\n", event.Raw.BlockNumber)
		fmt.Println("-----")
	}
	if logs.Error() != nil {
		fmt.Printf("Error reading logs: %v", logs.Error())
		return logs.Error(), ""
	}

	return nil, proposalId
}

func ReadKeystoreFile(filepath, password string) (*ecdsa.PrivateKey, error) {
	// 1. Read the keystore file
	keyjson, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read keystore: %v", err)
	}

	// 2. Decrypt with password
	key, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt key: %v (wrong password?)", err)
	}

	return key.PrivateKey, nil
}

func WriteJsonFile(data map[string]interface{}, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func ReadFileToString(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}
	return string(data), nil
}
