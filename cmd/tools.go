package cmd

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
	"juchain.org/chain/contract/contracts/generated/proposal"
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
	gasLimit = gasLimit * 120 / 100

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

	tx := types.NewTransaction(
		uint64(rawTx["nonce"].(float64)),
		common.HexToAddress(rawTx["to"].(string)),
		big.NewInt(int64(rawTx["value"].(float64))),
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

func SendSignedTx(proposer, rpcURL string, signedTxFile string) (common.Hash, error) {
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
	err, _ = QueryProposalId(blockHeight.Uint64(), proposer, client)
	return tx.Hash(), err
}

func waitEthTxFinished(client *ethclient.Client, txhash common.Hash) (error, *big.Int) {
	fmt.Printf("Wait for tx to be finished executing with hash %s\n", txhash.String())
	timeout := time.NewTimer(30 * time.Second)
	oneSecondtimeout := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-timeout.C:
			fmt.Println("tx timeout")
			return errors.New("eth tx timeout"), nil
		case <-oneSecondtimeout.C:
			receipt, err := client.TransactionReceipt(context.Background(), txhash)
			if err == ethereum.NotFound {
				continue
			} else if err != nil {
				return err, nil
			}
			fmt.Printf("tx confirmed in block %v\n", receipt.BlockNumber)
			return nil, receipt.BlockNumber
		}
	}
}

// 创建提案id
// flag true新增候选人，false踢出候选人
func BuildProposalId(from, dest string, flag bool, details string, blockNum uint64, client *ethclient.Client) (string, error) {
	sender := common.HexToAddress(from) // 提案人
	dst := common.HexToAddress(dest)    // 候选人
	flagValue := big.NewInt(0)
	if flag {
		flagValue.SetInt64(1)
	}
	block, err := client.BlockByNumber(context.Background(), big.NewInt(int64(blockNum)))
	if err != nil {
		return "", err
	}

	// // 计算提案ID（与Solidity keccak256(abi.encodePacked(...)) 等效）
	// id := crypto.Keccak256(
	// 	sender.Bytes(),
	// 	dst.Bytes(),
	// 	common.LeftPadBytes(flagValue.Bytes(), 32), // 将flag填充为32字节
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

// 查询生成的提案ID
func QueryProposalId(blockHeight uint64, proposer string, client *ethclient.Client) (error, string) {
	instance, err := proposal.NewProposal(common.HexToAddress(proposalAddr), client)
	if err != nil {
		fmt.Printf("Failed to instantiate contract: %v", err)
		return err, ""
	}

	// 定义查询过滤器
	filterOpts := &bind.FilterOpts{
		Start:   blockHeight,  // 起始区块号
		End:     &blockHeight, // 结束区块号（nil 表示最新区块）
		Context: context.Background(),
	}

	// 查询事件日志
	logs, err := instance.FilterLogCreateProposal(filterOpts, nil, []common.Address{common.HexToAddress(proposer)}, nil)
	if err != nil {
		fmt.Printf("Failed to filter LogCreateProposal events: %v", err)
		return err, ""
	}

	// 遍历日志
	proposalId := "0x"
	for logs.Next() {
		event := logs.Event
		proposalId = hex.EncodeToString(event.Id[:])
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
	return nil, proposalId
}

func ReadKeystoreFile(filepath, password string) (*ecdsa.PrivateKey, error) {
	// 1. Read the keystore file
	keyjson, err := ioutil.ReadFile(filepath)
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
