// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PunishMetaData contains all meta data concerning the Punish contract.
var PunishMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"PROPOSAL_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUNISH_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cleanPunishRecord\",\"inputs\":[{\"name\":\"val\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreaseMissedBlocksCounter\",\"inputs\":[{\"name\":\"epoch\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getPunishRecord\",\"inputs\":[{\"name\":\"val\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getPunishValidatorsLen\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"validators_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposal_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"staking_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"punish\",\"inputs\":[{\"name\":\"val\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"punishValidators\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"LogDecreaseMissedBlocksCounter\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogPunishValidator\",\"inputs\":[{\"name\":\"val\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055610ee0806100405f395ff3fe608060405234801561000f575f5ffd5b50600436106100b1575f3560e01c80639f759dba1161006e5780639f759dba14610152578063c0c53b8b1461015b578063d93d2cb914610170578063e0d8ea5314610183578063ea7221a11461018b578063f62af26c1461019e575f5ffd5b8063158ef93e146100b557806332f3c17f146100d6578063437ccda81461010c5780635d4f0cb61461012d57806363e1d451146101365780638c872d0514610149575b5f5ffd5b5f546100c19060ff1681565b60405190151581526020015b60405180910390f35b6100fe6100e4366004610d75565b6001600160a01b03165f9081526003602052604090205490565b6040519081526020016100cd565b61011561f01281565b6040516001600160a01b0390911681526020016100cd565b61011561f01381565b6100c1610144366004610d75565b6101b1565b61011561f01181565b61011561f01081565b61016e610169366004610d95565b610369565b005b61016e61017e366004610dd5565b610476565b6004546100fe565b61016e610199366004610d75565b610680565b6101156101ac366004610dd5565b610ac1565b5f6101ba610ae9565b6101c2610b2b565b6001600160a01b0382165f90815260036020526040902054156101f8576001600160a01b0382165f908152600360205260408120555b6001600160a01b0382165f9081526003602052604090206002015460ff168015610223575060045415155b156103605760045461023790600190610e00565b6001600160a01b0383165f908152600360205260409020600101541461030357600480545f919061026a90600190610e00565b8154811061027a5761027a610e19565b5f9182526020808320909101546001600160a01b03868116845260039092526040909220600101546004805492909316935083929181106102bd576102bd610e19565b5f91825260208083209190910180546001600160a01b0319166001600160a01b039485161790558583168252600390526040808220600190810154949093168252902001555b600480548061031457610314610e2d565b5f828152602080822083015f1990810180546001600160a01b03191690559092019092556001600160a01b038416825260039052604081206001810191909155600201805460ff191690555b5060015b919050565b610371610b7c565b6001600160a01b0383166103cc5760405162461bcd60e51b815260206004820152601a60248201527f496e76616c69642076616c696461746f7273206164647265737300000000000060448201526064015b60405180910390fd5b6001600160a01b0382166104225760405162461bcd60e51b815260206004820152601860248201527f496e76616c69642070726f706f73616c2061646472657373000000000000000060448201526064016103c3565b5f8054600180546001600160a01b03199081166001600160a01b0396871617825560028054909116948616949094179093556001600160a81b031916610100949093169390930260ff191691909117179055565b61047e610bc4565b610486610c00565b61048e610ae9565b8061049881610c53565b435f908152600660205260409020805460ff191660011790556004541561067c57600154604080516344c1aa9960e01b815290515f926001600160a01b0316916344c1aa999160048083019260209291908290030181865afa158015610500573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105249190610e41565b90505f60015f9054906101000a90046001600160a01b03166001600160a01b0316632897183d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610577573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061059b9190610e41565b90505f6105a88284610e6c565b6004549091505f5b8181101561064e575f600482815481106105cc576105cc610e19565b5f9182526020808320909101546001600160a01b0316808352600390915260409091205490915084101561062c576001600160a01b0381165f9081526003602052604081208054869290610621908490610e00565b909155506106459050565b6001600160a01b0381165f908152600360205260408120555b506001016105b0565b506040517f181d51be54e8e8eaca6eae0eab32d4162099236bd519e7238d015d0870db4641905f90a1505050505b5050565b610688610bc4565b610690610ae9565b610698610c9d565b6106a0610cef565b435f908152600560209081526040808320805460ff191660011790556001600160a01b0384168352600390915290206002015460ff1661074757600480546001600160a01b0383165f81815260036020526040812060018082018590558085019095557f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b90930180546001600160a01b0319168317905552600201805460ff191690911790555b6001600160a01b0381165f90815260036020526040812080549161076a83610e7f565b919050555060015f9054906101000a90046001600160a01b03166001600160a01b03166344c1aa996040518163ffffffff1660e01b8152600401602060405180830381865afa1580156107bf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107e39190610e41565b6001600160a01b0382165f908152600360205260409020546108059190610e97565b5f03610956576001600160a01b038082165f90815260036020908152604080832092909255600254600154835163f945b62360e01b8152935191851694635a4d66c0948794929091169263f945b623926004808401939192918290030181865afa158015610875573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108999190610e41565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044015f604051808303815f87803b1580156108dc575f5ffd5b505af11580156108ee573d5f5f3e3d5ffd5b50505f546040516340a141ff60e01b81526001600160a01b03858116600483015261010090920490911692506340a141ff91506024015f604051808303815f87803b15801561093b575f5ffd5b505af115801561094d573d5f5f3e3d5ffd5b50505050610a52565b60015f9054906101000a90046001600160a01b03166001600160a01b031663cb1ea7256040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109a6573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906109ca9190610e41565b6001600160a01b0382165f908152600360205260409020546109ec9190610e97565b5f03610a52575f546040516305dd095960e41b81526001600160a01b03838116600483015261010090920490911690635dd09590906024015f604051808303815f87803b158015610a3b575f5ffd5b505af1158015610a4d573d5f5f3e3d5ffd5b505050505b806001600160a01b03167f770e0cca42c35d00240986ce8d3ed438be04663c91dac6576b79537d7c180f1e42604051610a8d91815260200190565b60405180910390a2610abe60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b50565b60048181548110610ad0575f80fd5b5f918252602090912001546001600160a01b0316905081565b5f5460ff16610b295760405162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b60448201526064016103c3565b565b3361f01014610b295760405162461bcd60e51b815260206004820152601860248201527f56616c696461746f727320636f6e7472616374206f6e6c79000000000000000060448201526064016103c3565b5f5460ff1615610b295760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b60448201526064016103c3565b334114610b295760405162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b60448201526064016103c3565b435f9081526006602052604090205460ff1615610b295760405162461bcd60e51b8152602060048201526011602482015270105b1c9958591e48191958dc99585cd959607a1b60448201526064016103c3565b610c5d8143610e97565b15610abe5760405162461bcd60e51b815260206004820152601060248201526f426c6f636b2065706f6368206f6e6c7960801b60448201526064016103c3565b435f9081526005602052604090205460ff1615610b295760405162461bcd60e51b815260206004820152601060248201526f105b1c9958591e481c1d5b9a5cda195960821b60448201526064016103c3565b610cf7610d1d565b60027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0054600203610b2957604051633ee5aeb560e01b815260040160405180910390fd5b80356001600160a01b0381168114610364575f5ffd5b5f60208284031215610d85575f5ffd5b610d8e82610d5f565b9392505050565b5f5f5f60608486031215610da7575f5ffd5b610db084610d5f565b9250610dbe60208501610d5f565b9150610dcc60408501610d5f565b90509250925092565b5f60208284031215610de5575f5ffd5b5035919050565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610e1357610e13610dec565b92915050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52603160045260245ffd5b5f60208284031215610e51575f5ffd5b5051919050565b634e487b7160e01b5f52601260045260245ffd5b5f82610e7a57610e7a610e58565b500490565b5f60018201610e9057610e90610dec565b5060010190565b5f82610ea557610ea5610e58565b50069056fea2646970667358221220909517b2e472c467904537e869cf8be44477edfde19c8d164d1ef5dba6c8ad2364736f6c634300081d0033",
}

// PunishABI is the input ABI used to generate the binding from.
// Deprecated: Use PunishMetaData.ABI instead.
var PunishABI = PunishMetaData.ABI

// PunishBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PunishMetaData.Bin instead.
var PunishBin = PunishMetaData.Bin

// DeployPunish deploys a new Ethereum contract, binding an instance of Punish to it.
func DeployPunish(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Punish, error) {
	parsed, err := PunishMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PunishBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Punish{PunishCaller: PunishCaller{contract: contract}, PunishTransactor: PunishTransactor{contract: contract}, PunishFilterer: PunishFilterer{contract: contract}}, nil
}

// Punish is an auto generated Go binding around an Ethereum contract.
type Punish struct {
	PunishCaller     // Read-only binding to the contract
	PunishTransactor // Write-only binding to the contract
	PunishFilterer   // Log filterer for contract events
}

// PunishCaller is an auto generated read-only Go binding around an Ethereum contract.
type PunishCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PunishTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PunishTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PunishFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PunishFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PunishSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PunishSession struct {
	Contract     *Punish           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PunishCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PunishCallerSession struct {
	Contract *PunishCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PunishTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PunishTransactorSession struct {
	Contract     *PunishTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PunishRaw is an auto generated low-level Go binding around an Ethereum contract.
type PunishRaw struct {
	Contract *Punish // Generic contract binding to access the raw methods on
}

// PunishCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PunishCallerRaw struct {
	Contract *PunishCaller // Generic read-only contract binding to access the raw methods on
}

// PunishTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PunishTransactorRaw struct {
	Contract *PunishTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPunish creates a new instance of Punish, bound to a specific deployed contract.
func NewPunish(address common.Address, backend bind.ContractBackend) (*Punish, error) {
	contract, err := bindPunish(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Punish{PunishCaller: PunishCaller{contract: contract}, PunishTransactor: PunishTransactor{contract: contract}, PunishFilterer: PunishFilterer{contract: contract}}, nil
}

// NewPunishCaller creates a new read-only instance of Punish, bound to a specific deployed contract.
func NewPunishCaller(address common.Address, caller bind.ContractCaller) (*PunishCaller, error) {
	contract, err := bindPunish(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PunishCaller{contract: contract}, nil
}

// NewPunishTransactor creates a new write-only instance of Punish, bound to a specific deployed contract.
func NewPunishTransactor(address common.Address, transactor bind.ContractTransactor) (*PunishTransactor, error) {
	contract, err := bindPunish(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PunishTransactor{contract: contract}, nil
}

// NewPunishFilterer creates a new log filterer instance of Punish, bound to a specific deployed contract.
func NewPunishFilterer(address common.Address, filterer bind.ContractFilterer) (*PunishFilterer, error) {
	contract, err := bindPunish(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PunishFilterer{contract: contract}, nil
}

// bindPunish binds a generic wrapper to an already deployed contract.
func bindPunish(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PunishMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Punish *PunishRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Punish.Contract.PunishCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Punish *PunishRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Punish.Contract.PunishTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Punish *PunishRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Punish.Contract.PunishTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Punish *PunishCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Punish.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Punish *PunishTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Punish.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Punish *PunishTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Punish.Contract.contract.Transact(opts, method, params...)
}

// PROPOSALADDR is a free data retrieval call binding the contract method 0x437ccda8.
//
// Solidity: function PROPOSAL_ADDR() view returns(address)
func (_Punish *PunishCaller) PROPOSALADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "PROPOSAL_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROPOSALADDR is a free data retrieval call binding the contract method 0x437ccda8.
//
// Solidity: function PROPOSAL_ADDR() view returns(address)
func (_Punish *PunishSession) PROPOSALADDR() (common.Address, error) {
	return _Punish.Contract.PROPOSALADDR(&_Punish.CallOpts)
}

// PROPOSALADDR is a free data retrieval call binding the contract method 0x437ccda8.
//
// Solidity: function PROPOSAL_ADDR() view returns(address)
func (_Punish *PunishCallerSession) PROPOSALADDR() (common.Address, error) {
	return _Punish.Contract.PROPOSALADDR(&_Punish.CallOpts)
}

// PUNISHADDR is a free data retrieval call binding the contract method 0x8c872d05.
//
// Solidity: function PUNISH_ADDR() view returns(address)
func (_Punish *PunishCaller) PUNISHADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "PUNISH_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PUNISHADDR is a free data retrieval call binding the contract method 0x8c872d05.
//
// Solidity: function PUNISH_ADDR() view returns(address)
func (_Punish *PunishSession) PUNISHADDR() (common.Address, error) {
	return _Punish.Contract.PUNISHADDR(&_Punish.CallOpts)
}

// PUNISHADDR is a free data retrieval call binding the contract method 0x8c872d05.
//
// Solidity: function PUNISH_ADDR() view returns(address)
func (_Punish *PunishCallerSession) PUNISHADDR() (common.Address, error) {
	return _Punish.Contract.PUNISHADDR(&_Punish.CallOpts)
}

// STAKINGADDR is a free data retrieval call binding the contract method 0x5d4f0cb6.
//
// Solidity: function STAKING_ADDR() view returns(address)
func (_Punish *PunishCaller) STAKINGADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "STAKING_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKINGADDR is a free data retrieval call binding the contract method 0x5d4f0cb6.
//
// Solidity: function STAKING_ADDR() view returns(address)
func (_Punish *PunishSession) STAKINGADDR() (common.Address, error) {
	return _Punish.Contract.STAKINGADDR(&_Punish.CallOpts)
}

// STAKINGADDR is a free data retrieval call binding the contract method 0x5d4f0cb6.
//
// Solidity: function STAKING_ADDR() view returns(address)
func (_Punish *PunishCallerSession) STAKINGADDR() (common.Address, error) {
	return _Punish.Contract.STAKINGADDR(&_Punish.CallOpts)
}

// VALIDATORADDR is a free data retrieval call binding the contract method 0x9f759dba.
//
// Solidity: function VALIDATOR_ADDR() view returns(address)
func (_Punish *PunishCaller) VALIDATORADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "VALIDATOR_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORADDR is a free data retrieval call binding the contract method 0x9f759dba.
//
// Solidity: function VALIDATOR_ADDR() view returns(address)
func (_Punish *PunishSession) VALIDATORADDR() (common.Address, error) {
	return _Punish.Contract.VALIDATORADDR(&_Punish.CallOpts)
}

// VALIDATORADDR is a free data retrieval call binding the contract method 0x9f759dba.
//
// Solidity: function VALIDATOR_ADDR() view returns(address)
func (_Punish *PunishCallerSession) VALIDATORADDR() (common.Address, error) {
	return _Punish.Contract.VALIDATORADDR(&_Punish.CallOpts)
}

// GetPunishRecord is a free data retrieval call binding the contract method 0x32f3c17f.
//
// Solidity: function getPunishRecord(address val) view returns(uint256)
func (_Punish *PunishCaller) GetPunishRecord(opts *bind.CallOpts, val common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "getPunishRecord", val)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPunishRecord is a free data retrieval call binding the contract method 0x32f3c17f.
//
// Solidity: function getPunishRecord(address val) view returns(uint256)
func (_Punish *PunishSession) GetPunishRecord(val common.Address) (*big.Int, error) {
	return _Punish.Contract.GetPunishRecord(&_Punish.CallOpts, val)
}

// GetPunishRecord is a free data retrieval call binding the contract method 0x32f3c17f.
//
// Solidity: function getPunishRecord(address val) view returns(uint256)
func (_Punish *PunishCallerSession) GetPunishRecord(val common.Address) (*big.Int, error) {
	return _Punish.Contract.GetPunishRecord(&_Punish.CallOpts, val)
}

// GetPunishValidatorsLen is a free data retrieval call binding the contract method 0xe0d8ea53.
//
// Solidity: function getPunishValidatorsLen() view returns(uint256)
func (_Punish *PunishCaller) GetPunishValidatorsLen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "getPunishValidatorsLen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPunishValidatorsLen is a free data retrieval call binding the contract method 0xe0d8ea53.
//
// Solidity: function getPunishValidatorsLen() view returns(uint256)
func (_Punish *PunishSession) GetPunishValidatorsLen() (*big.Int, error) {
	return _Punish.Contract.GetPunishValidatorsLen(&_Punish.CallOpts)
}

// GetPunishValidatorsLen is a free data retrieval call binding the contract method 0xe0d8ea53.
//
// Solidity: function getPunishValidatorsLen() view returns(uint256)
func (_Punish *PunishCallerSession) GetPunishValidatorsLen() (*big.Int, error) {
	return _Punish.Contract.GetPunishValidatorsLen(&_Punish.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Punish *PunishCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Punish *PunishSession) Initialized() (bool, error) {
	return _Punish.Contract.Initialized(&_Punish.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Punish *PunishCallerSession) Initialized() (bool, error) {
	return _Punish.Contract.Initialized(&_Punish.CallOpts)
}

// PunishValidators is a free data retrieval call binding the contract method 0xf62af26c.
//
// Solidity: function punishValidators(uint256 ) view returns(address)
func (_Punish *PunishCaller) PunishValidators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "punishValidators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunishValidators is a free data retrieval call binding the contract method 0xf62af26c.
//
// Solidity: function punishValidators(uint256 ) view returns(address)
func (_Punish *PunishSession) PunishValidators(arg0 *big.Int) (common.Address, error) {
	return _Punish.Contract.PunishValidators(&_Punish.CallOpts, arg0)
}

// PunishValidators is a free data retrieval call binding the contract method 0xf62af26c.
//
// Solidity: function punishValidators(uint256 ) view returns(address)
func (_Punish *PunishCallerSession) PunishValidators(arg0 *big.Int) (common.Address, error) {
	return _Punish.Contract.PunishValidators(&_Punish.CallOpts, arg0)
}

// CleanPunishRecord is a paid mutator transaction binding the contract method 0x63e1d451.
//
// Solidity: function cleanPunishRecord(address val) returns(bool)
func (_Punish *PunishTransactor) CleanPunishRecord(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _Punish.contract.Transact(opts, "cleanPunishRecord", val)
}

// CleanPunishRecord is a paid mutator transaction binding the contract method 0x63e1d451.
//
// Solidity: function cleanPunishRecord(address val) returns(bool)
func (_Punish *PunishSession) CleanPunishRecord(val common.Address) (*types.Transaction, error) {
	return _Punish.Contract.CleanPunishRecord(&_Punish.TransactOpts, val)
}

// CleanPunishRecord is a paid mutator transaction binding the contract method 0x63e1d451.
//
// Solidity: function cleanPunishRecord(address val) returns(bool)
func (_Punish *PunishTransactorSession) CleanPunishRecord(val common.Address) (*types.Transaction, error) {
	return _Punish.Contract.CleanPunishRecord(&_Punish.TransactOpts, val)
}

// DecreaseMissedBlocksCounter is a paid mutator transaction binding the contract method 0xd93d2cb9.
//
// Solidity: function decreaseMissedBlocksCounter(uint256 epoch) returns()
func (_Punish *PunishTransactor) DecreaseMissedBlocksCounter(opts *bind.TransactOpts, epoch *big.Int) (*types.Transaction, error) {
	return _Punish.contract.Transact(opts, "decreaseMissedBlocksCounter", epoch)
}

// DecreaseMissedBlocksCounter is a paid mutator transaction binding the contract method 0xd93d2cb9.
//
// Solidity: function decreaseMissedBlocksCounter(uint256 epoch) returns()
func (_Punish *PunishSession) DecreaseMissedBlocksCounter(epoch *big.Int) (*types.Transaction, error) {
	return _Punish.Contract.DecreaseMissedBlocksCounter(&_Punish.TransactOpts, epoch)
}

// DecreaseMissedBlocksCounter is a paid mutator transaction binding the contract method 0xd93d2cb9.
//
// Solidity: function decreaseMissedBlocksCounter(uint256 epoch) returns()
func (_Punish *PunishTransactorSession) DecreaseMissedBlocksCounter(epoch *big.Int) (*types.Transaction, error) {
	return _Punish.Contract.DecreaseMissedBlocksCounter(&_Punish.TransactOpts, epoch)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address validators_, address proposal_, address staking_) returns()
func (_Punish *PunishTransactor) Initialize(opts *bind.TransactOpts, validators_ common.Address, proposal_ common.Address, staking_ common.Address) (*types.Transaction, error) {
	return _Punish.contract.Transact(opts, "initialize", validators_, proposal_, staking_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address validators_, address proposal_, address staking_) returns()
func (_Punish *PunishSession) Initialize(validators_ common.Address, proposal_ common.Address, staking_ common.Address) (*types.Transaction, error) {
	return _Punish.Contract.Initialize(&_Punish.TransactOpts, validators_, proposal_, staking_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address validators_, address proposal_, address staking_) returns()
func (_Punish *PunishTransactorSession) Initialize(validators_ common.Address, proposal_ common.Address, staking_ common.Address) (*types.Transaction, error) {
	return _Punish.Contract.Initialize(&_Punish.TransactOpts, validators_, proposal_, staking_)
}

// Punish is a paid mutator transaction binding the contract method 0xea7221a1.
//
// Solidity: function punish(address val) returns()
func (_Punish *PunishTransactor) Punish(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _Punish.contract.Transact(opts, "punish", val)
}

// Punish is a paid mutator transaction binding the contract method 0xea7221a1.
//
// Solidity: function punish(address val) returns()
func (_Punish *PunishSession) Punish(val common.Address) (*types.Transaction, error) {
	return _Punish.Contract.Punish(&_Punish.TransactOpts, val)
}

// Punish is a paid mutator transaction binding the contract method 0xea7221a1.
//
// Solidity: function punish(address val) returns()
func (_Punish *PunishTransactorSession) Punish(val common.Address) (*types.Transaction, error) {
	return _Punish.Contract.Punish(&_Punish.TransactOpts, val)
}

// PunishLogDecreaseMissedBlocksCounterIterator is returned from FilterLogDecreaseMissedBlocksCounter and is used to iterate over the raw logs and unpacked data for LogDecreaseMissedBlocksCounter events raised by the Punish contract.
type PunishLogDecreaseMissedBlocksCounterIterator struct {
	Event *PunishLogDecreaseMissedBlocksCounter // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PunishLogDecreaseMissedBlocksCounterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PunishLogDecreaseMissedBlocksCounter)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PunishLogDecreaseMissedBlocksCounter)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PunishLogDecreaseMissedBlocksCounterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PunishLogDecreaseMissedBlocksCounterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PunishLogDecreaseMissedBlocksCounter represents a LogDecreaseMissedBlocksCounter event raised by the Punish contract.
type PunishLogDecreaseMissedBlocksCounter struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogDecreaseMissedBlocksCounter is a free log retrieval operation binding the contract event 0x181d51be54e8e8eaca6eae0eab32d4162099236bd519e7238d015d0870db4641.
//
// Solidity: event LogDecreaseMissedBlocksCounter()
func (_Punish *PunishFilterer) FilterLogDecreaseMissedBlocksCounter(opts *bind.FilterOpts) (*PunishLogDecreaseMissedBlocksCounterIterator, error) {

	logs, sub, err := _Punish.contract.FilterLogs(opts, "LogDecreaseMissedBlocksCounter")
	if err != nil {
		return nil, err
	}
	return &PunishLogDecreaseMissedBlocksCounterIterator{contract: _Punish.contract, event: "LogDecreaseMissedBlocksCounter", logs: logs, sub: sub}, nil
}

// WatchLogDecreaseMissedBlocksCounter is a free log subscription operation binding the contract event 0x181d51be54e8e8eaca6eae0eab32d4162099236bd519e7238d015d0870db4641.
//
// Solidity: event LogDecreaseMissedBlocksCounter()
func (_Punish *PunishFilterer) WatchLogDecreaseMissedBlocksCounter(opts *bind.WatchOpts, sink chan<- *PunishLogDecreaseMissedBlocksCounter) (event.Subscription, error) {

	logs, sub, err := _Punish.contract.WatchLogs(opts, "LogDecreaseMissedBlocksCounter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PunishLogDecreaseMissedBlocksCounter)
				if err := _Punish.contract.UnpackLog(event, "LogDecreaseMissedBlocksCounter", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogDecreaseMissedBlocksCounter is a log parse operation binding the contract event 0x181d51be54e8e8eaca6eae0eab32d4162099236bd519e7238d015d0870db4641.
//
// Solidity: event LogDecreaseMissedBlocksCounter()
func (_Punish *PunishFilterer) ParseLogDecreaseMissedBlocksCounter(log types.Log) (*PunishLogDecreaseMissedBlocksCounter, error) {
	event := new(PunishLogDecreaseMissedBlocksCounter)
	if err := _Punish.contract.UnpackLog(event, "LogDecreaseMissedBlocksCounter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PunishLogPunishValidatorIterator is returned from FilterLogPunishValidator and is used to iterate over the raw logs and unpacked data for LogPunishValidator events raised by the Punish contract.
type PunishLogPunishValidatorIterator struct {
	Event *PunishLogPunishValidator // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PunishLogPunishValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PunishLogPunishValidator)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PunishLogPunishValidator)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PunishLogPunishValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PunishLogPunishValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PunishLogPunishValidator represents a LogPunishValidator event raised by the Punish contract.
type PunishLogPunishValidator struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogPunishValidator is a free log retrieval operation binding the contract event 0x770e0cca42c35d00240986ce8d3ed438be04663c91dac6576b79537d7c180f1e.
//
// Solidity: event LogPunishValidator(address indexed val, uint256 time)
func (_Punish *PunishFilterer) FilterLogPunishValidator(opts *bind.FilterOpts, val []common.Address) (*PunishLogPunishValidatorIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Punish.contract.FilterLogs(opts, "LogPunishValidator", valRule)
	if err != nil {
		return nil, err
	}
	return &PunishLogPunishValidatorIterator{contract: _Punish.contract, event: "LogPunishValidator", logs: logs, sub: sub}, nil
}

// WatchLogPunishValidator is a free log subscription operation binding the contract event 0x770e0cca42c35d00240986ce8d3ed438be04663c91dac6576b79537d7c180f1e.
//
// Solidity: event LogPunishValidator(address indexed val, uint256 time)
func (_Punish *PunishFilterer) WatchLogPunishValidator(opts *bind.WatchOpts, sink chan<- *PunishLogPunishValidator, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Punish.contract.WatchLogs(opts, "LogPunishValidator", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PunishLogPunishValidator)
				if err := _Punish.contract.UnpackLog(event, "LogPunishValidator", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogPunishValidator is a log parse operation binding the contract event 0x770e0cca42c35d00240986ce8d3ed438be04663c91dac6576b79537d7c180f1e.
//
// Solidity: event LogPunishValidator(address indexed val, uint256 time)
func (_Punish *PunishFilterer) ParseLogPunishValidator(log types.Log) (*PunishLogPunishValidator, error) {
	event := new(PunishLogPunishValidator)
	if err := _Punish.contract.UnpackLog(event, "LogPunishValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
