// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package validators

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

// ParamsMetaData contains all meta data concerning the Params contract.
var ParamsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ProposalAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PunishContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValidatorContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060e48061001f6000396000f3fe6080604052348015600f57600080fd5b506004361060465760003560e01c8063158ef93e14604b5780631b5e358c1460655780633a061bd31460875780636233be5d14608d575b600080fd5b60516093565b604080519115158252519081900360200190f35b606b609c565b604080516001600160a01b039092168252519081900360200190f35b606b60a2565b606b60a8565b60005460ff1681565b61f00181565b61f00081565b61f0028156fea2646970667358221220000e6e43cca5141127f6ecdfad3d67119b746836dec146c7581f7816b5b34ac164736f6c63430007060033",
}

// ParamsABI is the input ABI used to generate the binding from.
// Deprecated: Use ParamsMetaData.ABI instead.
var ParamsABI = ParamsMetaData.ABI

// ParamsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ParamsMetaData.Bin instead.
var ParamsBin = ParamsMetaData.Bin

// DeployParams deploys a new Ethereum contract, binding an instance of Params to it.
func DeployParams(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Params, error) {
	parsed, err := ParamsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ParamsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Params{ParamsCaller: ParamsCaller{contract: contract}, ParamsTransactor: ParamsTransactor{contract: contract}, ParamsFilterer: ParamsFilterer{contract: contract}}, nil
}

// Params is an auto generated Go binding around an Ethereum contract.
type Params struct {
	ParamsCaller     // Read-only binding to the contract
	ParamsTransactor // Write-only binding to the contract
	ParamsFilterer   // Log filterer for contract events
}

// ParamsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ParamsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParamsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ParamsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParamsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ParamsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ParamsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ParamsSession struct {
	Contract     *Params           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParamsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ParamsCallerSession struct {
	Contract *ParamsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ParamsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ParamsTransactorSession struct {
	Contract     *ParamsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ParamsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ParamsRaw struct {
	Contract *Params // Generic contract binding to access the raw methods on
}

// ParamsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ParamsCallerRaw struct {
	Contract *ParamsCaller // Generic read-only contract binding to access the raw methods on
}

// ParamsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ParamsTransactorRaw struct {
	Contract *ParamsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewParams creates a new instance of Params, bound to a specific deployed contract.
func NewParams(address common.Address, backend bind.ContractBackend) (*Params, error) {
	contract, err := bindParams(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Params{ParamsCaller: ParamsCaller{contract: contract}, ParamsTransactor: ParamsTransactor{contract: contract}, ParamsFilterer: ParamsFilterer{contract: contract}}, nil
}

// NewParamsCaller creates a new read-only instance of Params, bound to a specific deployed contract.
func NewParamsCaller(address common.Address, caller bind.ContractCaller) (*ParamsCaller, error) {
	contract, err := bindParams(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ParamsCaller{contract: contract}, nil
}

// NewParamsTransactor creates a new write-only instance of Params, bound to a specific deployed contract.
func NewParamsTransactor(address common.Address, transactor bind.ContractTransactor) (*ParamsTransactor, error) {
	contract, err := bindParams(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ParamsTransactor{contract: contract}, nil
}

// NewParamsFilterer creates a new log filterer instance of Params, bound to a specific deployed contract.
func NewParamsFilterer(address common.Address, filterer bind.ContractFilterer) (*ParamsFilterer, error) {
	contract, err := bindParams(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ParamsFilterer{contract: contract}, nil
}

// bindParams binds a generic wrapper to an already deployed contract.
func bindParams(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ParamsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Params *ParamsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Params.Contract.ParamsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Params *ParamsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Params.Contract.ParamsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Params *ParamsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Params.Contract.ParamsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Params *ParamsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Params.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Params *ParamsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Params.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Params *ParamsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Params.Contract.contract.Transact(opts, method, params...)
}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Params *ParamsCaller) ProposalAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Params.contract.Call(opts, &out, "ProposalAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Params *ParamsSession) ProposalAddr() (common.Address, error) {
	return _Params.Contract.ProposalAddr(&_Params.CallOpts)
}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Params *ParamsCallerSession) ProposalAddr() (common.Address, error) {
	return _Params.Contract.ProposalAddr(&_Params.CallOpts)
}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Params *ParamsCaller) PunishContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Params.contract.Call(opts, &out, "PunishContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Params *ParamsSession) PunishContractAddr() (common.Address, error) {
	return _Params.Contract.PunishContractAddr(&_Params.CallOpts)
}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Params *ParamsCallerSession) PunishContractAddr() (common.Address, error) {
	return _Params.Contract.PunishContractAddr(&_Params.CallOpts)
}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Params *ParamsCaller) ValidatorContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Params.contract.Call(opts, &out, "ValidatorContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Params *ParamsSession) ValidatorContractAddr() (common.Address, error) {
	return _Params.Contract.ValidatorContractAddr(&_Params.CallOpts)
}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Params *ParamsCallerSession) ValidatorContractAddr() (common.Address, error) {
	return _Params.Contract.ValidatorContractAddr(&_Params.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Params *ParamsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Params.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Params *ParamsSession) Initialized() (bool, error) {
	return _Params.Contract.Initialized(&_Params.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Params *ParamsCallerSession) Initialized() (bool, error) {
	return _Params.Contract.Initialized(&_Params.CallOpts)
}

// ProposalMetaData contains all meta data concerning the Proposal contract.
var ProposalMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogCreateConfigProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogCreateProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogPassProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogRejectProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogSetUnpassed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"auth\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogVote\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ProposalAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PunishContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValidatorContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"createProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"createUpdateConfigProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decreaseRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increasePeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vals\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pass\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalLastingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"createTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"proposalType\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"flag\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"cid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"punishThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiverAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"removeThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"results\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"agree\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reject\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"resultExist\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"setUnpassed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"auth\",\"type\":\"bool\"}],\"name\":\"voteProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"votes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"voteTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auth\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawProfitPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506116ea806100206000396000f3fe608060405234801561001057600080fd5b506004361061012c5760003560e01c806382c4b3b2116100ad578063cb1ea72511610071578063cb1ea72514610408578063d24806eb14610410578063d51cade814610433578063d63e6ce7146104b8578063e823c814146104c05761012c565b806382c4b3b21461033d57806394522b6d14610363578063a224cee71461036b578063a3dcb4d2146103db578063a4c4d922146103e35761012c565b806332ed5b12116100f457806332ed5b12146102075780633a061bd3146102e157806344c1aa99146102e95780634c6b25b1146102f15780636233be5d146103355761012c565b8063158ef93e1461013157806315ea27811461014d5780631b5e358c146101735780631db5ade8146101975780632897183d146101ed575b600080fd5b6101396104c8565b604080519115158252519081900360200190f35b6101396004803603602081101561016357600080fd5b50356001600160a01b03166104d1565b61017b610586565b604080516001600160a01b039092168252519081900360200190f35b6101c3600480360360408110156101ad57600080fd5b506001600160a01b03813516906020013561058c565b604080516001600160a01b0390941684526020840192909252151582820152519081900360600190f35b6101f56105c5565b60408051918252519081900360200190f35b6102246004803603602081101561021d57600080fd5b50356105cb565b60405180896001600160a01b03168152602001888152602001878152602001866001600160a01b03168152602001851515815260200180602001848152602001838152602001828103825285818151815260200191508051906020019080838360005b8381101561029f578181015183820152602001610287565b50505050905090810190601f1680156102cc5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b61017b6106ab565b6101f56106b1565b61030e6004803603602081101561030757600080fd5b50356106b7565b6040805161ffff948516815292909316602083015215158183015290519081900360600190f35b61017b6106e4565b6101396004803603602081101561035357600080fd5b50356001600160a01b03166106ea565b6101f56106ff565b6103d96004803603602081101561038157600080fd5b810190602081018135600160201b81111561039b57600080fd5b8201836020820111156103ad57600080fd5b803590602001918460208302840111600160201b831117156103ce57600080fd5b509092509050610705565b005b61017b610894565b610139600480360360408110156103f957600080fd5b508035906020013515156108a3565b6101f5611055565b6101396004803603604081101561042657600080fd5b508035906020013561105b565b6101396004803603606081101561044957600080fd5b6001600160a01b03823516916020810135151591810190606081016040820135600160201b81111561047a57600080fd5b82018360208201111561048c57600080fd5b803590602001918460018302840111600160201b831117156104ad57600080fd5b5090925090506111b2565b6101f56114b4565b6101f56114ba565b60005460ff1681565b60003361f00014610529576040805162461bcd60e51b815260206004820152601860248201527f56616c696461746f727320636f6e7472616374206f6e6c790000000000000000604482015290519081900360640190fd5b6001600160a01b038216600081815260086020908152604091829020805460ff19169055815142815291517f4e0b191f7f5c32b1b5e3704b68874b1a3980147cae00be8ece271bfb5b92c07a9281900390910190a2506001919050565b61f00181565b600b6020908152600092835260408084209091529082529020805460018201546002909201546001600160a01b03909116919060ff1683565b60045481565b600960209081526000918252604091829020805460018083015460028085015460038601546004870180548a516101009782161597909702600019011693909304601f81018990048902860189019099528885526001600160a01b03958616989397919695811695600160a01b90910460ff169490939092918301828280156106955780601f1061066a57610100808354040283529160200191610695565b820191906000526020600020905b81548152906001019060200180831161067857829003601f168201915b5050505050908060050154908060060154905088565b61f00081565b60035481565b600a6020526000908152604090205461ffff8082169162010000810490911690600160201b900460ff1683565b61f00281565b60086020526000908152604090205460ff1681565b60055481565b60005460ff1615610753576040805162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604482015290519081900360640190fd5b600c80546001600160a01b03191661f00017905560005b8181101561083757600083838381811061078057fe5b905060200201356001600160a01b03166001600160a01b031614156107ec576040805162461bcd60e51b815260206004820152601960248201527f496e76616c69642076616c696461746f72206164647265737300000000000000604482015290519081900360640190fd5b6001600860008585858181106107fe57fe5b602090810292909201356001600160a01b0316835250810191909152604001600020805460ff191691151591909117905560010161076a565b505062093a8060019081556018600281905560306003556004556170806005556000805460ff191690911790555062015180600655600780546001600160a01b03191673f869b51b53f72036d84e3edf3ba09c5dd3d89a66179055565b6007546001600160a01b031681565b600c5460408051631015428760e21b815233600482015290516000926001600160a01b0316916340550a1c916024808301926020929190829003018186803b1580156108ee57600080fd5b505afa158015610902573d6000803e3d6000fd5b505050506040513d602081101561091857600080fd5b505161095c576040805162461bcd60e51b815260206004820152600e60248201526d56616c696461746f72206f6e6c7960901b604482015290519081900360640190fd5b6000838152600960205260409020600101546109b4576040805162461bcd60e51b8152602060048201526012602482015271141c9bdc1bdcd85b081b9bdd08195e1a5cdd60721b604482015290519081900360640190fd5b336000908152600b6020908152604080832086845290915290206001015415610a0e5760405162461bcd60e51b81526004018080602001828103825260238152602001806116526023913960400191505060405180910390fd5b60018054600085815260096020526040902090910154014210610a6b576040805162461bcd60e51b815260206004820152601060248201526f141c9bdc1bdcd85b08195e1c1a5c995960821b604482015290519081900360640190fd5b336000818152600b60209081526040808320878452825291829020426001820181905581546001600160a01b031916851782556002909101805460ff1916871515908117909155835190815291820152815186927f6c59bda68cac318717c60c7c9635a78a0f0613f9887cc18a7157f5745a86d14e928290030190a38115610b16576000838152600a60205260409020805461ffff8082166001011661ffff19909116179055610b47565b6000838152600a602052604090208054600161ffff62010000808404821692909201160263ffff0000199091161790555b6000838152600a6020526040902054600160201b900460ff1615610b6d5750600161104f565b600c54604080516313bce04b60e31b815290516002926001600160a01b031691639de70258916004808301926000929190829003018186803b158015610bb257600080fd5b505afa158015610bc6573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610bef57600080fd5b8101908080516040519392919084600160201b821115610c0e57600080fd5b908301906020820185811115610c2357600080fd5b82518660208202830111600160201b82111715610c3f57600080fd5b82525081516020918201928201910280838360005b83811015610c6c578181015183820152602001610c54565b505050509050016040525050505181610c8157fe5b6000858152600a602052604090205491900460010161ffff90911610610ebb576000838152600a60209081526040808320805464ff000000001916600160201b179055600990915290206002015460011415610e4057600083815260096020526040902060030154600160a01b900460ff1615610da957600083815260096020818152604080842060030180546001600160a01b03908116865260088452828620805460ff19166001179055600c548987529484529054825163503cc43160e11b81529082166004820152915193169363a079886293602480840194939192918390030190829087803b158015610d7757600080fd5b505af1158015610d8b573d6000803e3d6000fd5b505050506040513d6020811015610da157600080fd5b50610e3b9050565b600083815260096020818152604080842060030180546001600160a01b03908116865260088452828620805460ff19169055600c548987529490935254815163a1ff465560e01b815290831660048201529051929091169263a1ff46559260248084019382900301818387803b158015610e2257600080fd5b505af1158015610e36573d6000803e3d6000fd5b505050505b610e7d565b60008381526009602052604090206002908101541415610e7d5760008381526009602052604090206005810154600690910154610e7d91906114c0565b60408051428152905184917f90d2e923947d9356c1c04391cb9e2e9c5d4ad6c165a849787b0c7569bbe99e24919081900360200190a250600161104f565b600c54604080516313bce04b60e31b815290516002926001600160a01b031691639de70258916004808301926000929190829003018186803b158015610f0057600080fd5b505afa158015610f14573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526020811015610f3d57600080fd5b8101908080516040519392919084600160201b821115610f5c57600080fd5b908301906020820185811115610f7157600080fd5b82518660208202830111600160201b82111715610f8d57600080fd5b82525081516020918201928201910280838360005b83811015610fba578181015183820152602001610fa2565b505050509050016040525050505181610fcf57fe5b6000858152600a60205260409020549190046001016201000090910461ffff161061104b576000838152600a6020908152604091829020805464ff000000001916600160201b1790558151428152915185927f36bdb56d707cdf53eadffe319a71ddf97736be67b8caab47b7720201a6b65ca092908290030190a25b5060015b92915050565b60025481565b604080513360601b60208083019190915260348201859052605482018490524260748084019190915283518084039091018152609490920190925280519101206000906110a6611557565b33815260c0810185905260e081018490524260208083019182526002604080850182815260008781526009855291909120855181546001600160a01b039182166001600160a01b031991821617835595516001830155915192810192909255606085015160038301805460808801511515600160a01b0260ff60a01b1993909416961695909517161790925560a0830151805184939261114d9260048501929101906115b0565b5060c0820151600582015560e090910151600690910155604080518681526020810186905242818301529051339184917f8bfc061277ae1778974ada10db7f9664ab1d67c455c025c025b438c52c69d1819181900360600190a3506001949350505050565b6001600160a01b03841660009081526008602052604081205460ff161580156111d85750835b8061120457506001600160a01b03851660009081526008602052604090205460ff168015611204575083155b61123f5760405162461bcd60e51b81526004018080602001828103825260408152602001806116756040913960400191505060405180910390fd5b600033868686864260405160200180876001600160a01b031660601b8152601401866001600160a01b031660601b815260140185151560f81b81526001018484808284379190910192835250506040805180830381526020928301909152805191012095505050610bb886111592506112f5915050576040805162461bcd60e51b815260206004820152601060248201526f44657461696c7320746f6f206c6f6e6760801b604482015290519081900360640190fd5b60008181526009602052604090206001015415611359576040805162461bcd60e51b815260206004820152601760248201527f50726f706f73616c20616c726561647920657869737473000000000000000000604482015290519081900360640190fd5b611361611557565b3381526001600160a01b03871660608201528515156080820152604080516020601f8701819004810282018101909252858152908690869081908401838280828437600092018290525060a0860194855242602080880191825260016040808a018281528b8652600984529420895181546001600160a01b039182166001600160a01b031991821617835594519282019290925593516002850155606089015160038501805460808c01511515600160a01b0260ff60a01b1993909416951694909417161790915594518051879692955061144594506004860193509101906115b0565b5060c0820151600582015560e09091015160069091015560408051871515815242602082015281516001600160a01b038a1692339286927f1af05d46b8c1ec021d82b7128cff40e91a1c2337deffc010df48eeddef8da56c929181900390910190a45060019695505050505050565b60065481565b60015481565b816114cf576001819055611553565b81600114156114e2576002819055611553565b81600214156114f5576003819055611553565b8160031415611508576004819055611553565b816004141561151b576005819055611553565b816005141561152e576006819055611553565b816006141561155357600780546001600160a01b0319166001600160a01b0383161790555b5050565b60405180610100016040528060006001600160a01b03168152602001600081526020016000815260200160006001600160a01b031681526020016000151581526020016060815260200160008152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f0160209004810192826115e6576000855561162c565b82601f106115ff57805160ff191683800117855561162c565b8280016001018555821561162c579182015b8281111561162c578251825591602001919060010190611611565b5061163892915061163c565b5090565b5b80821115611638576000815560010161163d56fe596f752063616e277420766f746520666f7220612070726f706f73616c20747769636543616e74227420616464206120616c726561647920657869737420647374206f722043616e7422742072656d6f76652061206e6f742070617373656420647374a2646970667358221220fdddedc4134e21798aac6760adc903e0fa94ba89bd0ac5c39322b1f8ddd6f8c564736f6c63430007060033",
}

// ProposalABI is the input ABI used to generate the binding from.
// Deprecated: Use ProposalMetaData.ABI instead.
var ProposalABI = ProposalMetaData.ABI

// ProposalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProposalMetaData.Bin instead.
var ProposalBin = ProposalMetaData.Bin

// DeployProposal deploys a new Ethereum contract, binding an instance of Proposal to it.
func DeployProposal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Proposal, error) {
	parsed, err := ProposalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProposalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proposal{ProposalCaller: ProposalCaller{contract: contract}, ProposalTransactor: ProposalTransactor{contract: contract}, ProposalFilterer: ProposalFilterer{contract: contract}}, nil
}

// Proposal is an auto generated Go binding around an Ethereum contract.
type Proposal struct {
	ProposalCaller     // Read-only binding to the contract
	ProposalTransactor // Write-only binding to the contract
	ProposalFilterer   // Log filterer for contract events
}

// ProposalCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProposalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProposalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProposalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProposalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProposalSession struct {
	Contract     *Proposal         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProposalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProposalCallerSession struct {
	Contract *ProposalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ProposalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProposalTransactorSession struct {
	Contract     *ProposalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ProposalRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProposalRaw struct {
	Contract *Proposal // Generic contract binding to access the raw methods on
}

// ProposalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProposalCallerRaw struct {
	Contract *ProposalCaller // Generic read-only contract binding to access the raw methods on
}

// ProposalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProposalTransactorRaw struct {
	Contract *ProposalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProposal creates a new instance of Proposal, bound to a specific deployed contract.
func NewProposal(address common.Address, backend bind.ContractBackend) (*Proposal, error) {
	contract, err := bindProposal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proposal{ProposalCaller: ProposalCaller{contract: contract}, ProposalTransactor: ProposalTransactor{contract: contract}, ProposalFilterer: ProposalFilterer{contract: contract}}, nil
}

// NewProposalCaller creates a new read-only instance of Proposal, bound to a specific deployed contract.
func NewProposalCaller(address common.Address, caller bind.ContractCaller) (*ProposalCaller, error) {
	contract, err := bindProposal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalCaller{contract: contract}, nil
}

// NewProposalTransactor creates a new write-only instance of Proposal, bound to a specific deployed contract.
func NewProposalTransactor(address common.Address, transactor bind.ContractTransactor) (*ProposalTransactor, error) {
	contract, err := bindProposal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProposalTransactor{contract: contract}, nil
}

// NewProposalFilterer creates a new log filterer instance of Proposal, bound to a specific deployed contract.
func NewProposalFilterer(address common.Address, filterer bind.ContractFilterer) (*ProposalFilterer, error) {
	contract, err := bindProposal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProposalFilterer{contract: contract}, nil
}

// bindProposal binds a generic wrapper to an already deployed contract.
func bindProposal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ProposalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proposal *ProposalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proposal.Contract.ProposalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proposal *ProposalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.Contract.ProposalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proposal *ProposalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proposal.Contract.ProposalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proposal *ProposalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proposal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proposal *ProposalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proposal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proposal *ProposalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proposal.Contract.contract.Transact(opts, method, params...)
}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Proposal *ProposalCaller) ProposalAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "ProposalAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Proposal *ProposalSession) ProposalAddr() (common.Address, error) {
	return _Proposal.Contract.ProposalAddr(&_Proposal.CallOpts)
}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Proposal *ProposalCallerSession) ProposalAddr() (common.Address, error) {
	return _Proposal.Contract.ProposalAddr(&_Proposal.CallOpts)
}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Proposal *ProposalCaller) PunishContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "PunishContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Proposal *ProposalSession) PunishContractAddr() (common.Address, error) {
	return _Proposal.Contract.PunishContractAddr(&_Proposal.CallOpts)
}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Proposal *ProposalCallerSession) PunishContractAddr() (common.Address, error) {
	return _Proposal.Contract.PunishContractAddr(&_Proposal.CallOpts)
}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Proposal *ProposalCaller) ValidatorContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "ValidatorContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Proposal *ProposalSession) ValidatorContractAddr() (common.Address, error) {
	return _Proposal.Contract.ValidatorContractAddr(&_Proposal.CallOpts)
}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Proposal *ProposalCallerSession) ValidatorContractAddr() (common.Address, error) {
	return _Proposal.Contract.ValidatorContractAddr(&_Proposal.CallOpts)
}

// DecreaseRate is a free data retrieval call binding the contract method 0x2897183d.
//
// Solidity: function decreaseRate() view returns(uint256)
func (_Proposal *ProposalCaller) DecreaseRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "decreaseRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DecreaseRate is a free data retrieval call binding the contract method 0x2897183d.
//
// Solidity: function decreaseRate() view returns(uint256)
func (_Proposal *ProposalSession) DecreaseRate() (*big.Int, error) {
	return _Proposal.Contract.DecreaseRate(&_Proposal.CallOpts)
}

// DecreaseRate is a free data retrieval call binding the contract method 0x2897183d.
//
// Solidity: function decreaseRate() view returns(uint256)
func (_Proposal *ProposalCallerSession) DecreaseRate() (*big.Int, error) {
	return _Proposal.Contract.DecreaseRate(&_Proposal.CallOpts)
}

// IncreasePeriod is a free data retrieval call binding the contract method 0xd63e6ce7.
//
// Solidity: function increasePeriod() view returns(uint256)
func (_Proposal *ProposalCaller) IncreasePeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "increasePeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IncreasePeriod is a free data retrieval call binding the contract method 0xd63e6ce7.
//
// Solidity: function increasePeriod() view returns(uint256)
func (_Proposal *ProposalSession) IncreasePeriod() (*big.Int, error) {
	return _Proposal.Contract.IncreasePeriod(&_Proposal.CallOpts)
}

// IncreasePeriod is a free data retrieval call binding the contract method 0xd63e6ce7.
//
// Solidity: function increasePeriod() view returns(uint256)
func (_Proposal *ProposalCallerSession) IncreasePeriod() (*big.Int, error) {
	return _Proposal.Contract.IncreasePeriod(&_Proposal.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Proposal *ProposalCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Proposal *ProposalSession) Initialized() (bool, error) {
	return _Proposal.Contract.Initialized(&_Proposal.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Proposal *ProposalCallerSession) Initialized() (bool, error) {
	return _Proposal.Contract.Initialized(&_Proposal.CallOpts)
}

// Pass is a free data retrieval call binding the contract method 0x82c4b3b2.
//
// Solidity: function pass(address ) view returns(bool)
func (_Proposal *ProposalCaller) Pass(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "pass", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Pass is a free data retrieval call binding the contract method 0x82c4b3b2.
//
// Solidity: function pass(address ) view returns(bool)
func (_Proposal *ProposalSession) Pass(arg0 common.Address) (bool, error) {
	return _Proposal.Contract.Pass(&_Proposal.CallOpts, arg0)
}

// Pass is a free data retrieval call binding the contract method 0x82c4b3b2.
//
// Solidity: function pass(address ) view returns(bool)
func (_Proposal *ProposalCallerSession) Pass(arg0 common.Address) (bool, error) {
	return _Proposal.Contract.Pass(&_Proposal.CallOpts, arg0)
}

// ProposalLastingPeriod is a free data retrieval call binding the contract method 0xe823c814.
//
// Solidity: function proposalLastingPeriod() view returns(uint256)
func (_Proposal *ProposalCaller) ProposalLastingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "proposalLastingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalLastingPeriod is a free data retrieval call binding the contract method 0xe823c814.
//
// Solidity: function proposalLastingPeriod() view returns(uint256)
func (_Proposal *ProposalSession) ProposalLastingPeriod() (*big.Int, error) {
	return _Proposal.Contract.ProposalLastingPeriod(&_Proposal.CallOpts)
}

// ProposalLastingPeriod is a free data retrieval call binding the contract method 0xe823c814.
//
// Solidity: function proposalLastingPeriod() view returns(uint256)
func (_Proposal *ProposalCallerSession) ProposalLastingPeriod() (*big.Int, error) {
	return _Proposal.Contract.ProposalLastingPeriod(&_Proposal.CallOpts)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(address proposer, uint256 createTime, uint256 proposalType, address dst, bool flag, string details, uint256 cid, uint256 newValue)
func (_Proposal *ProposalCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Proposer     common.Address
	CreateTime   *big.Int
	ProposalType *big.Int
	Dst          common.Address
	Flag         bool
	Details      string
	Cid          *big.Int
	NewValue     *big.Int
}, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Proposer     common.Address
		CreateTime   *big.Int
		ProposalType *big.Int
		Dst          common.Address
		Flag         bool
		Details      string
		Cid          *big.Int
		NewValue     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proposer = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.CreateTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProposalType = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Dst = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Flag = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Details = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.Cid = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.NewValue = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(address proposer, uint256 createTime, uint256 proposalType, address dst, bool flag, string details, uint256 cid, uint256 newValue)
func (_Proposal *ProposalSession) Proposals(arg0 [32]byte) (struct {
	Proposer     common.Address
	CreateTime   *big.Int
	ProposalType *big.Int
	Dst          common.Address
	Flag         bool
	Details      string
	Cid          *big.Int
	NewValue     *big.Int
}, error) {
	return _Proposal.Contract.Proposals(&_Proposal.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(address proposer, uint256 createTime, uint256 proposalType, address dst, bool flag, string details, uint256 cid, uint256 newValue)
func (_Proposal *ProposalCallerSession) Proposals(arg0 [32]byte) (struct {
	Proposer     common.Address
	CreateTime   *big.Int
	ProposalType *big.Int
	Dst          common.Address
	Flag         bool
	Details      string
	Cid          *big.Int
	NewValue     *big.Int
}, error) {
	return _Proposal.Contract.Proposals(&_Proposal.CallOpts, arg0)
}

// PunishThreshold is a free data retrieval call binding the contract method 0xcb1ea725.
//
// Solidity: function punishThreshold() view returns(uint256)
func (_Proposal *ProposalCaller) PunishThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "punishThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunishThreshold is a free data retrieval call binding the contract method 0xcb1ea725.
//
// Solidity: function punishThreshold() view returns(uint256)
func (_Proposal *ProposalSession) PunishThreshold() (*big.Int, error) {
	return _Proposal.Contract.PunishThreshold(&_Proposal.CallOpts)
}

// PunishThreshold is a free data retrieval call binding the contract method 0xcb1ea725.
//
// Solidity: function punishThreshold() view returns(uint256)
func (_Proposal *ProposalCallerSession) PunishThreshold() (*big.Int, error) {
	return _Proposal.Contract.PunishThreshold(&_Proposal.CallOpts)
}

// ReceiverAddr is a free data retrieval call binding the contract method 0xa3dcb4d2.
//
// Solidity: function receiverAddr() view returns(address)
func (_Proposal *ProposalCaller) ReceiverAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "receiverAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReceiverAddr is a free data retrieval call binding the contract method 0xa3dcb4d2.
//
// Solidity: function receiverAddr() view returns(address)
func (_Proposal *ProposalSession) ReceiverAddr() (common.Address, error) {
	return _Proposal.Contract.ReceiverAddr(&_Proposal.CallOpts)
}

// ReceiverAddr is a free data retrieval call binding the contract method 0xa3dcb4d2.
//
// Solidity: function receiverAddr() view returns(address)
func (_Proposal *ProposalCallerSession) ReceiverAddr() (common.Address, error) {
	return _Proposal.Contract.ReceiverAddr(&_Proposal.CallOpts)
}

// RemoveThreshold is a free data retrieval call binding the contract method 0x44c1aa99.
//
// Solidity: function removeThreshold() view returns(uint256)
func (_Proposal *ProposalCaller) RemoveThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "removeThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemoveThreshold is a free data retrieval call binding the contract method 0x44c1aa99.
//
// Solidity: function removeThreshold() view returns(uint256)
func (_Proposal *ProposalSession) RemoveThreshold() (*big.Int, error) {
	return _Proposal.Contract.RemoveThreshold(&_Proposal.CallOpts)
}

// RemoveThreshold is a free data retrieval call binding the contract method 0x44c1aa99.
//
// Solidity: function removeThreshold() view returns(uint256)
func (_Proposal *ProposalCallerSession) RemoveThreshold() (*big.Int, error) {
	return _Proposal.Contract.RemoveThreshold(&_Proposal.CallOpts)
}

// Results is a free data retrieval call binding the contract method 0x4c6b25b1.
//
// Solidity: function results(bytes32 ) view returns(uint16 agree, uint16 reject, bool resultExist)
func (_Proposal *ProposalCaller) Results(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Agree       uint16
	Reject      uint16
	ResultExist bool
}, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "results", arg0)

	outstruct := new(struct {
		Agree       uint16
		Reject      uint16
		ResultExist bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Agree = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.Reject = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.ResultExist = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Results is a free data retrieval call binding the contract method 0x4c6b25b1.
//
// Solidity: function results(bytes32 ) view returns(uint16 agree, uint16 reject, bool resultExist)
func (_Proposal *ProposalSession) Results(arg0 [32]byte) (struct {
	Agree       uint16
	Reject      uint16
	ResultExist bool
}, error) {
	return _Proposal.Contract.Results(&_Proposal.CallOpts, arg0)
}

// Results is a free data retrieval call binding the contract method 0x4c6b25b1.
//
// Solidity: function results(bytes32 ) view returns(uint16 agree, uint16 reject, bool resultExist)
func (_Proposal *ProposalCallerSession) Results(arg0 [32]byte) (struct {
	Agree       uint16
	Reject      uint16
	ResultExist bool
}, error) {
	return _Proposal.Contract.Results(&_Proposal.CallOpts, arg0)
}

// Votes is a free data retrieval call binding the contract method 0x1db5ade8.
//
// Solidity: function votes(address , bytes32 ) view returns(address voter, uint256 voteTime, bool auth)
func (_Proposal *ProposalCaller) Votes(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	Voter    common.Address
	VoteTime *big.Int
	Auth     bool
}, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "votes", arg0, arg1)

	outstruct := new(struct {
		Voter    common.Address
		VoteTime *big.Int
		Auth     bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Voter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.VoteTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Auth = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// Votes is a free data retrieval call binding the contract method 0x1db5ade8.
//
// Solidity: function votes(address , bytes32 ) view returns(address voter, uint256 voteTime, bool auth)
func (_Proposal *ProposalSession) Votes(arg0 common.Address, arg1 [32]byte) (struct {
	Voter    common.Address
	VoteTime *big.Int
	Auth     bool
}, error) {
	return _Proposal.Contract.Votes(&_Proposal.CallOpts, arg0, arg1)
}

// Votes is a free data retrieval call binding the contract method 0x1db5ade8.
//
// Solidity: function votes(address , bytes32 ) view returns(address voter, uint256 voteTime, bool auth)
func (_Proposal *ProposalCallerSession) Votes(arg0 common.Address, arg1 [32]byte) (struct {
	Voter    common.Address
	VoteTime *big.Int
	Auth     bool
}, error) {
	return _Proposal.Contract.Votes(&_Proposal.CallOpts, arg0, arg1)
}

// WithdrawProfitPeriod is a free data retrieval call binding the contract method 0x94522b6d.
//
// Solidity: function withdrawProfitPeriod() view returns(uint256)
func (_Proposal *ProposalCaller) WithdrawProfitPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "withdrawProfitPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawProfitPeriod is a free data retrieval call binding the contract method 0x94522b6d.
//
// Solidity: function withdrawProfitPeriod() view returns(uint256)
func (_Proposal *ProposalSession) WithdrawProfitPeriod() (*big.Int, error) {
	return _Proposal.Contract.WithdrawProfitPeriod(&_Proposal.CallOpts)
}

// WithdrawProfitPeriod is a free data retrieval call binding the contract method 0x94522b6d.
//
// Solidity: function withdrawProfitPeriod() view returns(uint256)
func (_Proposal *ProposalCallerSession) WithdrawProfitPeriod() (*big.Int, error) {
	return _Proposal.Contract.WithdrawProfitPeriod(&_Proposal.CallOpts)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xd51cade8.
//
// Solidity: function createProposal(address dst, bool flag, string details) returns(bool)
func (_Proposal *ProposalTransactor) CreateProposal(opts *bind.TransactOpts, dst common.Address, flag bool, details string) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "createProposal", dst, flag, details)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xd51cade8.
//
// Solidity: function createProposal(address dst, bool flag, string details) returns(bool)
func (_Proposal *ProposalSession) CreateProposal(dst common.Address, flag bool, details string) (*types.Transaction, error) {
	return _Proposal.Contract.CreateProposal(&_Proposal.TransactOpts, dst, flag, details)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xd51cade8.
//
// Solidity: function createProposal(address dst, bool flag, string details) returns(bool)
func (_Proposal *ProposalTransactorSession) CreateProposal(dst common.Address, flag bool, details string) (*types.Transaction, error) {
	return _Proposal.Contract.CreateProposal(&_Proposal.TransactOpts, dst, flag, details)
}

// CreateUpdateConfigProposal is a paid mutator transaction binding the contract method 0xd24806eb.
//
// Solidity: function createUpdateConfigProposal(uint256 cid, uint256 newValue) returns(bool)
func (_Proposal *ProposalTransactor) CreateUpdateConfigProposal(opts *bind.TransactOpts, cid *big.Int, newValue *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "createUpdateConfigProposal", cid, newValue)
}

// CreateUpdateConfigProposal is a paid mutator transaction binding the contract method 0xd24806eb.
//
// Solidity: function createUpdateConfigProposal(uint256 cid, uint256 newValue) returns(bool)
func (_Proposal *ProposalSession) CreateUpdateConfigProposal(cid *big.Int, newValue *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.CreateUpdateConfigProposal(&_Proposal.TransactOpts, cid, newValue)
}

// CreateUpdateConfigProposal is a paid mutator transaction binding the contract method 0xd24806eb.
//
// Solidity: function createUpdateConfigProposal(uint256 cid, uint256 newValue) returns(bool)
func (_Proposal *ProposalTransactorSession) CreateUpdateConfigProposal(cid *big.Int, newValue *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.CreateUpdateConfigProposal(&_Proposal.TransactOpts, cid, newValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] vals) returns()
func (_Proposal *ProposalTransactor) Initialize(opts *bind.TransactOpts, vals []common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "initialize", vals)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] vals) returns()
func (_Proposal *ProposalSession) Initialize(vals []common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.Initialize(&_Proposal.TransactOpts, vals)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] vals) returns()
func (_Proposal *ProposalTransactorSession) Initialize(vals []common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.Initialize(&_Proposal.TransactOpts, vals)
}

// SetUnpassed is a paid mutator transaction binding the contract method 0x15ea2781.
//
// Solidity: function setUnpassed(address val) returns(bool)
func (_Proposal *ProposalTransactor) SetUnpassed(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "setUnpassed", val)
}

// SetUnpassed is a paid mutator transaction binding the contract method 0x15ea2781.
//
// Solidity: function setUnpassed(address val) returns(bool)
func (_Proposal *ProposalSession) SetUnpassed(val common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.SetUnpassed(&_Proposal.TransactOpts, val)
}

// SetUnpassed is a paid mutator transaction binding the contract method 0x15ea2781.
//
// Solidity: function setUnpassed(address val) returns(bool)
func (_Proposal *ProposalTransactorSession) SetUnpassed(val common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.SetUnpassed(&_Proposal.TransactOpts, val)
}

// VoteProposal is a paid mutator transaction binding the contract method 0xa4c4d922.
//
// Solidity: function voteProposal(bytes32 id, bool auth) returns(bool)
func (_Proposal *ProposalTransactor) VoteProposal(opts *bind.TransactOpts, id [32]byte, auth bool) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "voteProposal", id, auth)
}

// VoteProposal is a paid mutator transaction binding the contract method 0xa4c4d922.
//
// Solidity: function voteProposal(bytes32 id, bool auth) returns(bool)
func (_Proposal *ProposalSession) VoteProposal(id [32]byte, auth bool) (*types.Transaction, error) {
	return _Proposal.Contract.VoteProposal(&_Proposal.TransactOpts, id, auth)
}

// VoteProposal is a paid mutator transaction binding the contract method 0xa4c4d922.
//
// Solidity: function voteProposal(bytes32 id, bool auth) returns(bool)
func (_Proposal *ProposalTransactorSession) VoteProposal(id [32]byte, auth bool) (*types.Transaction, error) {
	return _Proposal.Contract.VoteProposal(&_Proposal.TransactOpts, id, auth)
}

// ProposalLogCreateConfigProposalIterator is returned from FilterLogCreateConfigProposal and is used to iterate over the raw logs and unpacked data for LogCreateConfigProposal events raised by the Proposal contract.
type ProposalLogCreateConfigProposalIterator struct {
	Event *ProposalLogCreateConfigProposal // Event containing the contract specifics and raw log

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
func (it *ProposalLogCreateConfigProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalLogCreateConfigProposal)
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
		it.Event = new(ProposalLogCreateConfigProposal)
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
func (it *ProposalLogCreateConfigProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalLogCreateConfigProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalLogCreateConfigProposal represents a LogCreateConfigProposal event raised by the Proposal contract.
type ProposalLogCreateConfigProposal struct {
	Id       [32]byte
	Proposer common.Address
	Cid      *big.Int
	NewValue *big.Int
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogCreateConfigProposal is a free log retrieval operation binding the contract event 0x8bfc061277ae1778974ada10db7f9664ab1d67c455c025c025b438c52c69d181.
//
// Solidity: event LogCreateConfigProposal(bytes32 indexed id, address indexed proposer, uint256 cid, uint256 newValue, uint256 time)
func (_Proposal *ProposalFilterer) FilterLogCreateConfigProposal(opts *bind.FilterOpts, id [][32]byte, proposer []common.Address) (*ProposalLogCreateConfigProposalIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "LogCreateConfigProposal", idRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return &ProposalLogCreateConfigProposalIterator{contract: _Proposal.contract, event: "LogCreateConfigProposal", logs: logs, sub: sub}, nil
}

// WatchLogCreateConfigProposal is a free log subscription operation binding the contract event 0x8bfc061277ae1778974ada10db7f9664ab1d67c455c025c025b438c52c69d181.
//
// Solidity: event LogCreateConfigProposal(bytes32 indexed id, address indexed proposer, uint256 cid, uint256 newValue, uint256 time)
func (_Proposal *ProposalFilterer) WatchLogCreateConfigProposal(opts *bind.WatchOpts, sink chan<- *ProposalLogCreateConfigProposal, id [][32]byte, proposer []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "LogCreateConfigProposal", idRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalLogCreateConfigProposal)
				if err := _Proposal.contract.UnpackLog(event, "LogCreateConfigProposal", log); err != nil {
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

// ParseLogCreateConfigProposal is a log parse operation binding the contract event 0x8bfc061277ae1778974ada10db7f9664ab1d67c455c025c025b438c52c69d181.
//
// Solidity: event LogCreateConfigProposal(bytes32 indexed id, address indexed proposer, uint256 cid, uint256 newValue, uint256 time)
func (_Proposal *ProposalFilterer) ParseLogCreateConfigProposal(log types.Log) (*ProposalLogCreateConfigProposal, error) {
	event := new(ProposalLogCreateConfigProposal)
	if err := _Proposal.contract.UnpackLog(event, "LogCreateConfigProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalLogCreateProposalIterator is returned from FilterLogCreateProposal and is used to iterate over the raw logs and unpacked data for LogCreateProposal events raised by the Proposal contract.
type ProposalLogCreateProposalIterator struct {
	Event *ProposalLogCreateProposal // Event containing the contract specifics and raw log

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
func (it *ProposalLogCreateProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalLogCreateProposal)
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
		it.Event = new(ProposalLogCreateProposal)
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
func (it *ProposalLogCreateProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalLogCreateProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalLogCreateProposal represents a LogCreateProposal event raised by the Proposal contract.
type ProposalLogCreateProposal struct {
	Id       [32]byte
	Proposer common.Address
	Dst      common.Address
	Flag     bool
	Time     *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogCreateProposal is a free log retrieval operation binding the contract event 0x1af05d46b8c1ec021d82b7128cff40e91a1c2337deffc010df48eeddef8da56c.
//
// Solidity: event LogCreateProposal(bytes32 indexed id, address indexed proposer, address indexed dst, bool flag, uint256 time)
func (_Proposal *ProposalFilterer) FilterLogCreateProposal(opts *bind.FilterOpts, id [][32]byte, proposer []common.Address, dst []common.Address) (*ProposalLogCreateProposalIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "LogCreateProposal", idRule, proposerRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &ProposalLogCreateProposalIterator{contract: _Proposal.contract, event: "LogCreateProposal", logs: logs, sub: sub}, nil
}

// WatchLogCreateProposal is a free log subscription operation binding the contract event 0x1af05d46b8c1ec021d82b7128cff40e91a1c2337deffc010df48eeddef8da56c.
//
// Solidity: event LogCreateProposal(bytes32 indexed id, address indexed proposer, address indexed dst, bool flag, uint256 time)
func (_Proposal *ProposalFilterer) WatchLogCreateProposal(opts *bind.WatchOpts, sink chan<- *ProposalLogCreateProposal, id [][32]byte, proposer []common.Address, dst []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "LogCreateProposal", idRule, proposerRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalLogCreateProposal)
				if err := _Proposal.contract.UnpackLog(event, "LogCreateProposal", log); err != nil {
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

// ParseLogCreateProposal is a log parse operation binding the contract event 0x1af05d46b8c1ec021d82b7128cff40e91a1c2337deffc010df48eeddef8da56c.
//
// Solidity: event LogCreateProposal(bytes32 indexed id, address indexed proposer, address indexed dst, bool flag, uint256 time)
func (_Proposal *ProposalFilterer) ParseLogCreateProposal(log types.Log) (*ProposalLogCreateProposal, error) {
	event := new(ProposalLogCreateProposal)
	if err := _Proposal.contract.UnpackLog(event, "LogCreateProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalLogPassProposalIterator is returned from FilterLogPassProposal and is used to iterate over the raw logs and unpacked data for LogPassProposal events raised by the Proposal contract.
type ProposalLogPassProposalIterator struct {
	Event *ProposalLogPassProposal // Event containing the contract specifics and raw log

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
func (it *ProposalLogPassProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalLogPassProposal)
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
		it.Event = new(ProposalLogPassProposal)
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
func (it *ProposalLogPassProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalLogPassProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalLogPassProposal represents a LogPassProposal event raised by the Proposal contract.
type ProposalLogPassProposal struct {
	Id   [32]byte
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogPassProposal is a free log retrieval operation binding the contract event 0x90d2e923947d9356c1c04391cb9e2e9c5d4ad6c165a849787b0c7569bbe99e24.
//
// Solidity: event LogPassProposal(bytes32 indexed id, uint256 time)
func (_Proposal *ProposalFilterer) FilterLogPassProposal(opts *bind.FilterOpts, id [][32]byte) (*ProposalLogPassProposalIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "LogPassProposal", idRule)
	if err != nil {
		return nil, err
	}
	return &ProposalLogPassProposalIterator{contract: _Proposal.contract, event: "LogPassProposal", logs: logs, sub: sub}, nil
}

// WatchLogPassProposal is a free log subscription operation binding the contract event 0x90d2e923947d9356c1c04391cb9e2e9c5d4ad6c165a849787b0c7569bbe99e24.
//
// Solidity: event LogPassProposal(bytes32 indexed id, uint256 time)
func (_Proposal *ProposalFilterer) WatchLogPassProposal(opts *bind.WatchOpts, sink chan<- *ProposalLogPassProposal, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "LogPassProposal", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalLogPassProposal)
				if err := _Proposal.contract.UnpackLog(event, "LogPassProposal", log); err != nil {
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

// ParseLogPassProposal is a log parse operation binding the contract event 0x90d2e923947d9356c1c04391cb9e2e9c5d4ad6c165a849787b0c7569bbe99e24.
//
// Solidity: event LogPassProposal(bytes32 indexed id, uint256 time)
func (_Proposal *ProposalFilterer) ParseLogPassProposal(log types.Log) (*ProposalLogPassProposal, error) {
	event := new(ProposalLogPassProposal)
	if err := _Proposal.contract.UnpackLog(event, "LogPassProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalLogRejectProposalIterator is returned from FilterLogRejectProposal and is used to iterate over the raw logs and unpacked data for LogRejectProposal events raised by the Proposal contract.
type ProposalLogRejectProposalIterator struct {
	Event *ProposalLogRejectProposal // Event containing the contract specifics and raw log

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
func (it *ProposalLogRejectProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalLogRejectProposal)
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
		it.Event = new(ProposalLogRejectProposal)
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
func (it *ProposalLogRejectProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalLogRejectProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalLogRejectProposal represents a LogRejectProposal event raised by the Proposal contract.
type ProposalLogRejectProposal struct {
	Id   [32]byte
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRejectProposal is a free log retrieval operation binding the contract event 0x36bdb56d707cdf53eadffe319a71ddf97736be67b8caab47b7720201a6b65ca0.
//
// Solidity: event LogRejectProposal(bytes32 indexed id, uint256 time)
func (_Proposal *ProposalFilterer) FilterLogRejectProposal(opts *bind.FilterOpts, id [][32]byte) (*ProposalLogRejectProposalIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "LogRejectProposal", idRule)
	if err != nil {
		return nil, err
	}
	return &ProposalLogRejectProposalIterator{contract: _Proposal.contract, event: "LogRejectProposal", logs: logs, sub: sub}, nil
}

// WatchLogRejectProposal is a free log subscription operation binding the contract event 0x36bdb56d707cdf53eadffe319a71ddf97736be67b8caab47b7720201a6b65ca0.
//
// Solidity: event LogRejectProposal(bytes32 indexed id, uint256 time)
func (_Proposal *ProposalFilterer) WatchLogRejectProposal(opts *bind.WatchOpts, sink chan<- *ProposalLogRejectProposal, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "LogRejectProposal", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalLogRejectProposal)
				if err := _Proposal.contract.UnpackLog(event, "LogRejectProposal", log); err != nil {
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

// ParseLogRejectProposal is a log parse operation binding the contract event 0x36bdb56d707cdf53eadffe319a71ddf97736be67b8caab47b7720201a6b65ca0.
//
// Solidity: event LogRejectProposal(bytes32 indexed id, uint256 time)
func (_Proposal *ProposalFilterer) ParseLogRejectProposal(log types.Log) (*ProposalLogRejectProposal, error) {
	event := new(ProposalLogRejectProposal)
	if err := _Proposal.contract.UnpackLog(event, "LogRejectProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalLogSetUnpassedIterator is returned from FilterLogSetUnpassed and is used to iterate over the raw logs and unpacked data for LogSetUnpassed events raised by the Proposal contract.
type ProposalLogSetUnpassedIterator struct {
	Event *ProposalLogSetUnpassed // Event containing the contract specifics and raw log

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
func (it *ProposalLogSetUnpassedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalLogSetUnpassed)
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
		it.Event = new(ProposalLogSetUnpassed)
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
func (it *ProposalLogSetUnpassedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalLogSetUnpassedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalLogSetUnpassed represents a LogSetUnpassed event raised by the Proposal contract.
type ProposalLogSetUnpassed struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogSetUnpassed is a free log retrieval operation binding the contract event 0x4e0b191f7f5c32b1b5e3704b68874b1a3980147cae00be8ece271bfb5b92c07a.
//
// Solidity: event LogSetUnpassed(address indexed val, uint256 time)
func (_Proposal *ProposalFilterer) FilterLogSetUnpassed(opts *bind.FilterOpts, val []common.Address) (*ProposalLogSetUnpassedIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "LogSetUnpassed", valRule)
	if err != nil {
		return nil, err
	}
	return &ProposalLogSetUnpassedIterator{contract: _Proposal.contract, event: "LogSetUnpassed", logs: logs, sub: sub}, nil
}

// WatchLogSetUnpassed is a free log subscription operation binding the contract event 0x4e0b191f7f5c32b1b5e3704b68874b1a3980147cae00be8ece271bfb5b92c07a.
//
// Solidity: event LogSetUnpassed(address indexed val, uint256 time)
func (_Proposal *ProposalFilterer) WatchLogSetUnpassed(opts *bind.WatchOpts, sink chan<- *ProposalLogSetUnpassed, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "LogSetUnpassed", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalLogSetUnpassed)
				if err := _Proposal.contract.UnpackLog(event, "LogSetUnpassed", log); err != nil {
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

// ParseLogSetUnpassed is a log parse operation binding the contract event 0x4e0b191f7f5c32b1b5e3704b68874b1a3980147cae00be8ece271bfb5b92c07a.
//
// Solidity: event LogSetUnpassed(address indexed val, uint256 time)
func (_Proposal *ProposalFilterer) ParseLogSetUnpassed(log types.Log) (*ProposalLogSetUnpassed, error) {
	event := new(ProposalLogSetUnpassed)
	if err := _Proposal.contract.UnpackLog(event, "LogSetUnpassed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProposalLogVoteIterator is returned from FilterLogVote and is used to iterate over the raw logs and unpacked data for LogVote events raised by the Proposal contract.
type ProposalLogVoteIterator struct {
	Event *ProposalLogVote // Event containing the contract specifics and raw log

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
func (it *ProposalLogVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ProposalLogVote)
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
		it.Event = new(ProposalLogVote)
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
func (it *ProposalLogVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ProposalLogVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ProposalLogVote represents a LogVote event raised by the Proposal contract.
type ProposalLogVote struct {
	Id    [32]byte
	Voter common.Address
	Auth  bool
	Time  *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogVote is a free log retrieval operation binding the contract event 0x6c59bda68cac318717c60c7c9635a78a0f0613f9887cc18a7157f5745a86d14e.
//
// Solidity: event LogVote(bytes32 indexed id, address indexed voter, bool auth, uint256 time)
func (_Proposal *ProposalFilterer) FilterLogVote(opts *bind.FilterOpts, id [][32]byte, voter []common.Address) (*ProposalLogVoteIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Proposal.contract.FilterLogs(opts, "LogVote", idRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &ProposalLogVoteIterator{contract: _Proposal.contract, event: "LogVote", logs: logs, sub: sub}, nil
}

// WatchLogVote is a free log subscription operation binding the contract event 0x6c59bda68cac318717c60c7c9635a78a0f0613f9887cc18a7157f5745a86d14e.
//
// Solidity: event LogVote(bytes32 indexed id, address indexed voter, bool auth, uint256 time)
func (_Proposal *ProposalFilterer) WatchLogVote(opts *bind.WatchOpts, sink chan<- *ProposalLogVote, id [][32]byte, voter []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _Proposal.contract.WatchLogs(opts, "LogVote", idRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ProposalLogVote)
				if err := _Proposal.contract.UnpackLog(event, "LogVote", log); err != nil {
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

// ParseLogVote is a log parse operation binding the contract event 0x6c59bda68cac318717c60c7c9635a78a0f0613f9887cc18a7157f5745a86d14e.
//
// Solidity: event LogVote(bytes32 indexed id, address indexed voter, bool auth, uint256 time)
func (_Proposal *ProposalFilterer) ParseLogVote(log types.Log) (*ProposalLogVote, error) {
	event := new(ProposalLogVote)
	if err := _Proposal.contract.UnpackLog(event, "LogVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PunishMetaData contains all meta data concerning the Punish contract.
var PunishMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[],\"name\":\"LogDecreaseMissedBlocksCounter\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogPunishValidator\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ProposalAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PunishContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValidatorContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"cleanPunishRecord\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"decreaseMissedBlocksCounter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getPunishRecord\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPunishValidatorsLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"punish\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"punishValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610d8b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806363e1d4511161007157806363e1d451146101365780638129fc1c1461015c578063d93d2cb914610166578063e0d8ea5314610183578063ea7221a11461018b578063f62af26c146101b1576100a9565b8063158ef93e146100ae5780631b5e358c146100ca57806332f3c17f146100ee5780633a061bd3146101265780636233be5d1461012e575b600080fd5b6100b66101ce565b604080519115158252519081900360200190f35b6100d26101d7565b604080516001600160a01b039092168252519081900360200190f35b6101146004803603602081101561010457600080fd5b50356001600160a01b03166101dd565b60408051918252519081900360200190f35b6100d26101f8565b6100d26101fe565b6100b66004803603602081101561014c57600080fd5b50356001600160a01b0316610204565b610164610431565b005b6101646004803603602081101561017c57600080fd5b50356104b4565b610114610921565b610164600480360360208110156101a157600080fd5b50356001600160a01b0316610927565b6100d2600480360360208110156101c757600080fd5b5035610d2b565b60005460ff1681565b61f00181565b6001600160a01b031660009081526002602052604090205490565b61f00081565b61f00281565b6000805460ff1661024b576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b3361f000146102a1576040805162461bcd60e51b815260206004820152601860248201527f56616c696461746f727320636f6e7472616374206f6e6c790000000000000000604482015290519081900360640190fd5b6001600160a01b038216600090815260026020526040902054156102d9576001600160a01b0382166000908152600260205260408120555b6001600160a01b0382166000908152600260208190526040909120015460ff168015610306575060035415155b15610429576003546001600160a01b038316600090815260026020526040902060010154600019909101146103d0576003805460009190600019810190811061034b57fe5b60009182526020808320909101546001600160a01b038681168452600290925260409092206001015460038054929093169350839291811061038957fe5b600091825260208083209190910180546001600160a01b0319166001600160a01b039485161790558583168252600290526040808220600190810154949093168252902001555b60038054806103db57fe5b60008281526020808220830160001990810180546001600160a01b03191690559092019092556001600160a01b038416825260029081905260408220600181019290925501805460ff191690555b506001919050565b60005460ff161561047f576040805162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604482015290519081900360640190fd5b600080546001805461f0026001600160a01b0319909116178155610100600160a81b031990911662f000001760ff1916179055565b3341146104f5576040805162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015290519081900360640190fd5b4360009081526005602052604090205460ff161561054e576040805162461bcd60e51b8152602060048201526011602482015270105b1c9958591e48191958dc99585cd959607a1b604482015290519081900360640190fd5b60005460ff16610594576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b8080438161059e57fe5b06156105e4576040805162461bcd60e51b815260206004820152601060248201526f426c6f636b2065706f6368206f6e6c7960801b604482015290519081900360640190fd5b436000908152600560205260409020805460ff1916600117905560035461060a5761091d565b60005b6003548110156108f257600160009054906101000a90046001600160a01b03166001600160a01b0316632897183d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561066557600080fd5b505afa158015610679573d6000803e3d6000fd5b505050506040513d602081101561068f57600080fd5b5051600154604080516344c1aa9960e01b815290516001600160a01b03909216916344c1aa9991600480820192602092909190829003018186803b1580156106d657600080fd5b505afa1580156106ea573d6000803e3d6000fd5b505050506040513d602081101561070057600080fd5b50518161070957fe5b04600260006003848154811061071b57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400190205411156108b157600160009054906101000a90046001600160a01b03166001600160a01b0316632897183d6040518163ffffffff1660e01b815260040160206040518083038186803b15801561079457600080fd5b505afa1580156107a8573d6000803e3d6000fd5b505050506040513d60208110156107be57600080fd5b5051600154604080516344c1aa9960e01b815290516001600160a01b03909216916344c1aa9991600480820192602092909190829003018186803b15801561080557600080fd5b505afa158015610819573d6000803e3d6000fd5b505050506040513d602081101561082f57600080fd5b50518161083857fe5b04600260006003848154811061084a57fe5b60009182526020808320909101546001600160a01b0316835282019290925260400181205460038054939091039260029291908590811061088757fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020556108ea565b600060026000600384815481106108c457fe5b60009182526020808320909101546001600160a01b031683528201929092526040019020555b60010161060d565b506040517f181d51be54e8e8eaca6eae0eab32d4162099236bd519e7238d015d0870db464190600090a15b5050565b60035490565b334114610968576040805162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015290519081900360640190fd5b60005460ff166109ae576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b4360009081526004602052604090205460ff1615610a06576040805162461bcd60e51b815260206004820152601060248201526f105b1c9958591e481c1d5b9a5cda195960821b604482015290519081900360640190fd5b436000908152600460209081526040808320805460ff191660011790556001600160a01b0384168352600291829052909120015460ff16610ab157600380546001600160a01b0383166000818152600260208190526040822060018082018690558086019096557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b90940180546001600160a01b0319168417905591905201805460ff191690911790555b6001600160a01b03808216600090815260026020908152604091829020805460019081019091555482516344c1aa9960e01b815292519316926344c1aa99926004808201939291829003018186803b158015610b0c57600080fd5b505afa158015610b20573d6000803e3d6000fd5b505050506040513d6020811015610b3657600080fd5b50516001600160a01b03821660009081526002602052604090205481610b5857fe5b06610be25760008054604080516340a141ff60e01b81526001600160a01b0385811660048301529151610100909304909116926340a141ff9260248084019382900301818387803b158015610bac57600080fd5b505af1158015610bc0573d6000803e3d6000fd5b5050506001600160a01b03821660009081526002602052604081205550610ce9565b600160009054906101000a90046001600160a01b03166001600160a01b031663cb1ea7256040518163ffffffff1660e01b815260040160206040518083038186803b158015610c3057600080fd5b505afa158015610c44573d6000803e3d6000fd5b505050506040513d6020811015610c5a57600080fd5b50516001600160a01b03821660009081526002602052604090205481610c7c57fe5b06610ce95760008054604080516305dd095960e41b81526001600160a01b038581166004830152915161010090930490911692635dd095909260248084019382900301818387803b158015610cd057600080fd5b505af1158015610ce4573d6000803e3d6000fd5b505050505b6040805142815290516001600160a01b038316917f770e0cca42c35d00240986ce8d3ed438be04663c91dac6576b79537d7c180f1e919081900360200190a250565b60038181548110610d3b57600080fd5b6000918252602090912001546001600160a01b031690508156fea26469706673582212207bfabdd92b39d320f31e0a25b191be6bf050c893a89f8013b1bdadd09af245ee64736f6c63430007060033",
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

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Punish *PunishCaller) ProposalAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "ProposalAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Punish *PunishSession) ProposalAddr() (common.Address, error) {
	return _Punish.Contract.ProposalAddr(&_Punish.CallOpts)
}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Punish *PunishCallerSession) ProposalAddr() (common.Address, error) {
	return _Punish.Contract.ProposalAddr(&_Punish.CallOpts)
}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Punish *PunishCaller) PunishContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "PunishContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Punish *PunishSession) PunishContractAddr() (common.Address, error) {
	return _Punish.Contract.PunishContractAddr(&_Punish.CallOpts)
}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Punish *PunishCallerSession) PunishContractAddr() (common.Address, error) {
	return _Punish.Contract.PunishContractAddr(&_Punish.CallOpts)
}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Punish *PunishCaller) ValidatorContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Punish.contract.Call(opts, &out, "ValidatorContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Punish *PunishSession) ValidatorContractAddr() (common.Address, error) {
	return _Punish.Contract.ValidatorContractAddr(&_Punish.CallOpts)
}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Punish *PunishCallerSession) ValidatorContractAddr() (common.Address, error) {
	return _Punish.Contract.ValidatorContractAddr(&_Punish.CallOpts)
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

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Punish *PunishTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Punish.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Punish *PunishSession) Initialize() (*types.Transaction, error) {
	return _Punish.Contract.Initialize(&_Punish.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Punish *PunishTransactorSession) Initialize() (*types.Transaction, error) {
	return _Punish.Contract.Initialize(&_Punish.TransactOpts)
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

// SafeMathMetaData contains all meta data concerning the SafeMath contract.
var SafeMathMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207674d636ad5bb1dab55026686963488ff46d753661c80f4b603713678415b58c64736f6c63430007060033",
}

// SafeMathABI is the input ABI used to generate the binding from.
// Deprecated: Use SafeMathMetaData.ABI instead.
var SafeMathABI = SafeMathMetaData.ABI

// SafeMathBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SafeMathMetaData.Bin instead.
var SafeMathBin = SafeMathMetaData.Bin

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SafeMathMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// ValidatorsMetaData contains all meta data concerning the Validators contract.
var ValidatorsMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogActive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogAddToTopValidators\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockReward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogDistributeBlockReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogEditValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogRemoveFromTopValidators\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hb\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogRemoveValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hb\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogRemoveValidatorIncoming\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"newSet\",\"type\":\"address[]\"}],\"name\":\"LogUpdateValidator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fee\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"hb\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"time\",\"type\":\"uint256\"}],\"name\":\"LogWithdrawProfits\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ProposalAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PunishContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ValidatorContractAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"feeAddr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"createOrEditValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidatorSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeBlockReward\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTopValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getValidatorDescription\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumValidators.Status\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"highestValidatorsSet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vals\",\"type\":\"address[]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"isActiveValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"isTopValidator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"removeValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"removeValidatorIncoming\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalJailedHB\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"tryActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"val\",\"type\":\"address\"}],\"name\":\"tryRemoveValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"newSet\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"updateActiveValidatorSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"moniker\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"website\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"email\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"details\",\"type\":\"string\"}],\"name\":\"validateDescription\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"withdrawProfits\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506133cd806100206000396000f3fe60806040526004361061013f5760003560e01c80636969a25c116100b6578063a224cee71161006f578063a224cee714610538578063a406fcb7146105b3578063a43569b31461077e578063afeea115146109c1578063b6c88519146109d6578063d6c0edad14610c9b5761013f565b80636969a25c146103975780638a11d7c9146103c157806398e3b6261461043a5780639de702581461046d578063a0798862146104d2578063a1ff4655146105055761013f565b806340550a1c1161010857806340550a1c1461020d57806340a141ff146102405780634b3d500b146102755780635dd095901461029f5780636233be5d146102d25780636846992a146102e75761013f565b8062362a77146101445780631303f7cf1461018b578063158ef93e146101b25780631b5e358c146101c75780633a061bd3146101f8575b600080fd5b34801561015057600080fd5b506101776004803603602081101561016757600080fd5b50356001600160a01b0316610ca3565b604080519115158252519081900360200190f35b34801561019757600080fd5b506101a0610f75565b60408051918252519081900360200190f35b3480156101be57600080fd5b50610177610f7b565b3480156101d357600080fd5b506101dc610f84565b604080516001600160a01b039092168252519081900360200190f35b34801561020457600080fd5b506101dc610f8a565b34801561021957600080fd5b506101776004803603602081101561023057600080fd5b50356001600160a01b0316610f90565b34801561024c57600080fd5b506102736004803603602081101561026357600080fd5b50356001600160a01b0316610feb565b005b34801561028157600080fd5b506101dc6004803603602081101561029857600080fd5b5035611044565b3480156102ab57600080fd5b50610273600480360360208110156102c257600080fd5b50356001600160a01b031661106e565b3480156102de57600080fd5b506101dc6110c4565b3480156102f357600080fd5b506102736004803603604081101561030a57600080fd5b810190602081018135600160201b81111561032457600080fd5b82018360208201111561033657600080fd5b803590602001918460208302840111600160201b8311171561035757600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955050913592506110ca915050565b3480156103a357600080fd5b506101dc600480360360208110156103ba57600080fd5b5035611311565b3480156103cd57600080fd5b506103f4600480360360208110156103e457600080fd5b50356001600160a01b0316611321565b60405180866001600160a01b0316815260200185600281111561041357fe5b81526020018481526020018381526020018281526020019550505050505060405180910390f35b34801561044657600080fd5b506101776004803603602081101561045d57600080fd5b50356001600160a01b03166116e1565b34801561047957600080fd5b50610482611733565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156104be5781810151838201526020016104a6565b505050509050019250505060405180910390f35b3480156104de57600080fd5b50610177600480360360208110156104f557600080fd5b50356001600160a01b0316611795565b34801561051157600080fd5b506102736004803603602081101561052857600080fd5b50356001600160a01b03166119ca565b34801561054457600080fd5b506102736004803603602081101561055b57600080fd5b810190602081018135600160201b81111561057557600080fd5b82018360208201111561058757600080fd5b803590602001918460208302840111600160201b831117156105a857600080fd5b509092509050611a19565b3480156105bf57600080fd5b50610177600480360360c08110156105d657600080fd5b6001600160a01b038235169190810190604081016020820135600160201b81111561060057600080fd5b82018360208201111561061257600080fd5b803590602001918460018302840111600160201b8311171561063357600080fd5b919390929091602081019035600160201b81111561065057600080fd5b82018360208201111561066257600080fd5b803590602001918460018302840111600160201b8311171561068357600080fd5b919390929091602081019035600160201b8111156106a057600080fd5b8201836020820111156106b257600080fd5b803590602001918460018302840111600160201b831117156106d357600080fd5b919390929091602081019035600160201b8111156106f057600080fd5b82018360208201111561070257600080fd5b803590602001918460018302840111600160201b8311171561072357600080fd5b919390929091602081019035600160201b81111561074057600080fd5b82018360208201111561075257600080fd5b803590602001918460018302840111600160201b8311171561077357600080fd5b509092509050611d79565b34801561078a57600080fd5b506107b1600480360360208110156107a157600080fd5b50356001600160a01b03166122b1565b60405180806020018060200180602001806020018060200186810386528b818151815260200191508051906020019080838360005b838110156107fe5781810151838201526020016107e6565b50505050905090810190601f16801561082b5780820380516001836020036101000a031916815260200191505b5086810385528a5181528a516020918201918c019080838360005b8381101561085e578181015183820152602001610846565b50505050905090810190601f16801561088b5780820380516001836020036101000a031916815260200191505b5086810384528951815289516020918201918b019080838360005b838110156108be5781810151838201526020016108a6565b50505050905090810190601f1680156108eb5780820380516001836020036101000a031916815260200191505b5086810383528851815288516020918201918a019080838360005b8381101561091e578181015183820152602001610906565b50505050905090810190601f16801561094b5780820380516001836020036101000a031916815260200191505b50868103825287518152875160209182019189019080838360005b8381101561097e578181015183820152602001610966565b50505050905090810190601f1680156109ab5780820380516001836020036101000a031916815260200191505b509a505050505050505050505060405180910390f35b3480156109cd57600080fd5b50610482612678565b3480156109e257600080fd5b50610177600480360360a08110156109f957600080fd5b810190602081018135600160201b811115610a1357600080fd5b820183602082011115610a2557600080fd5b803590602001918460018302840111600160201b83111715610a4657600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610a9857600080fd5b820183602082011115610aaa57600080fd5b803590602001918460018302840111600160201b83111715610acb57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610b1d57600080fd5b820183602082011115610b2f57600080fd5b803590602001918460018302840111600160201b83111715610b5057600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610ba257600080fd5b820183602082011115610bb457600080fd5b803590602001918460018302840111600160201b83111715610bd557600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295949360208101935035915050600160201b811115610c2757600080fd5b820183602082011115610c3957600080fd5b803590602001918460018302840111600160201b83111715610c5a57600080fd5b91908080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152509295506126d8945050505050565b61027361287d565b600033816001600160a01b038416600090815260016020526040902054600160a01b900460ff166002811115610cd557fe5b1415610d1e576040805162461bcd60e51b815260206004820152601360248201527215985b1a59185d1bdc881b9bdd08195e1a5cdd606a1b604482015290519081900360640190fd5b6001600160a01b03838116600090815260016020526040902054811690821614610d795760405162461bcd60e51b815260040180806020018281038252602e81526020018061336a602e913960400191505060405180910390fd5b600554604080516394522b6d60e01b8152905143926001600160a01b0316916394522b6d916004808301926020929190829003018186803b158015610dbd57600080fd5b505afa158015610dd1573d6000803e3d6000fd5b505050506040513d6020811015610de757600080fd5b50516001600160a01b038516600090815260016020526040902060080154011115610e435760405162461bcd60e51b815260040180806020018281038252605c8152602001806132ed605c913960600191505060405180910390fd5b6001600160a01b03831660009081526001602052604090206006015480610eb1576040805162461bcd60e51b815260206004820152601a60248201527f596f7520646f6e2774206861766520616e792070726f66697473000000000000604482015290519081900360640190fd5b6001600160a01b03841660009081526001602052604081206006810191909155436008909101558015610f16576040516001600160a01b0383169082156108fc029083906000818181858888f19350505050158015610f14573d6000803e3d6000fd5b505b816001600160a01b0316846001600160a01b03167f51a69b4502f660774c9339825c7b5adbf0b8622289134647e29728ec5d9b3bb98342604051808381526020018281526020019250505060405180910390a36001925050505b919050565b60045481565b60005460ff1681565b61f00181565b61f00081565b6000805b600254811015610fe257826001600160a01b031660028281548110610fb557fe5b6000918252602090912001546001600160a01b03161415610fda576001915050610f70565b600101610f94565b50600092915050565b3361f00114611038576040805162461bcd60e51b815260206004820152601460248201527350756e69736820636f6e7472616374206f6e6c7960601b604482015290519081900360640190fd5b61104181612a1d565b50565b6003818154811061105457600080fd5b6000918252602090912001546001600160a01b0316905081565b3361f001146110bb576040805162461bcd60e51b815260206004820152601460248201527350756e69736820636f6e7472616374206f6e6c7960601b604482015290519081900360640190fd5b61104181612b2b565b61f00281565b33411461110b576040805162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015290519081900360640190fd5b4360009081526007602090815260408083206001845290915290205460ff161561117c576040805162461bcd60e51b815260206004820152601a60248201527f56616c696461746f727320616c72656164792075706461746564000000000000604482015290519081900360640190fd5b60005460ff166111c2576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b808043816111cc57fe5b0615611212576040805162461bcd60e51b815260206004820152601060248201526f426c6f636b2065706f6368206f6e6c7960801b604482015290519081900360640190fd5b43600090815260076020908152604080832060018085529252909120805460ff191690911790558251611283576040805162461bcd60e51b815260206004820152601460248201527356616c696461746f722073657420656d7074792160601b604482015290519081900360640190fd5b82516112969060029060208601906131f6565b507feacea8f3c22f06c0b18306bdb04d0a967255129e8ce0094debb0a0ff89d006b5836040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156112f95781810151838201526020016112e1565b505050509050019250505060405180910390a1505050565b6002818154811061105457600080fd5b6001600160a01b038181166000908152600160209081526040808320815160c08101909252805494851682529293849384938493849384939192830190600160a01b900460ff16600281111561137357fe5b600281111561137e57fe5b8152602001600182016040518060a0016040529081600082018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561142b5780601f106114005761010080835404028352916020019161142b565b820191906000526020600020905b81548152906001019060200180831161140e57829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156114cd5780601f106114a2576101008083540402835291602001916114cd565b820191906000526020600020905b8154815290600101906020018083116114b057829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f8101839004830285018301909152808452938101939083018282801561155f5780601f106115345761010080835404028352916020019161155f565b820191906000526020600020905b81548152906001019060200180831161154257829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156115f35780601f106115c8576101008083540402835291602001916115f3565b820191906000526020600020905b8154815290600101906020018083116115d657829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156116875780601f1061165c57610100808354040283529160200191611687565b820191906000526020600020905b81548152906001019060200180831161166a57829003601f168201915b50505050508152505081526020016006820154815260200160078201548152602001600882015481525050905080600001518160200151826060015183608001518460a00151955095509550955095505091939590929450565b6000805b600354811015610fe257826001600160a01b03166003828154811061170657fe5b6000918252602090912001546001600160a01b0316141561172b576001915050610f70565b6001016116e5565b6060600280548060200260200160405190810160405280929190818152602001828054801561178b57602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161176d575b5050505050905090565b60003361f002146117e6576040805162461bcd60e51b815260206004820152601660248201527550726f706f73616c20636f6e7472616374206f6e6c7960501b604482015290519081900360640190fd5b60005460ff1661182c576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b60016001600160a01b038316600090815260016020526040902054600160a01b900460ff16600281111561185c57fe5b141561186a57506001610f70565b61187382612c43565b60026001600160a01b038316600090815260016020526040902054600160a01b900460ff1660028111156118a357fe5b141561196457600654604080516363e1d45160e01b81526001600160a01b038581166004830152915191909216916363e1d4519160248083019260209291908290030181600087803b1580156118f857600080fd5b505af115801561190c573d6000803e3d6000fd5b505050506040513d602081101561192257600080fd5b5051611964576040805162461bcd60e51b815260206004820152600c60248201526b18db19585b8819985a5b195960a21b604482015290519081900360640190fd5b6001600160a01b038216600081815260016020908152604091829020805460ff60a01b1916600160a01b179055815142815291517f8bef9a500ef702fa4b7c82318f7b750176b75d33c8897ad10a35e5e5e41613629281900390910190a2506001919050565b3361f00214611038576040805162461bcd60e51b815260206004820152601660248201527550726f706f73616c20636f6e7472616374206f6e6c7960501b604482015290519081900360640190fd5b60005460ff1615611a67576040805162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604482015290519081900360640190fd5b600580546001600160a01b031990811661f002179091556006805490911661f00117905560005b81811015611d67576000838383818110611aa457fe5b905060200201356001600160a01b03166001600160a01b03161415611b10576040805162461bcd60e51b815260206004820152601960248201527f496e76616c69642076616c696461746f72206164647265737300000000000000604482015290519081900360640190fd5b611b34838383818110611b1f57fe5b905060200201356001600160a01b0316610f90565b611b83576002838383818110611b4657fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b611ba7838383818110611b9257fe5b905060200201356001600160a01b03166116e1565b611bf6576003838383818110611bb957fe5b835460018101855560009485526020948590200180546001600160a01b0319166001600160a01b0395909202939093013593909316929092179055505b6000600181858585818110611c0757fe5b6001600160a01b0360209182029390930135831684528301939093526040909101600020541691909114159050611cbc57828282818110611c4457fe5b905060200201356001600160a01b031660016000858585818110611c6457fe5b905060200201356001600160a01b03166001600160a01b03166001600160a01b0316815260200190815260200160002060000160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b600060016000858585818110611cce57fe5b602090810292909201356001600160a01b031683525081019190915260400160002054600160a01b900460ff166002811115611d0657fe5b1415611d5f576001806000858585818110611d1d57fe5b602090810292909201356001600160a01b0316835250810191909152604001600020805460ff60a01b1916600160a01b836002811115611d5957fe5b02179055505b600101611a8e565b50506000805460ff1916600117905550565b6000805460ff16611dc0576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b6001600160a01b038c16611e11576040805162461bcd60e51b8152602060048201526013602482015272496e76616c696420666565206164647265737360681b604482015290519081900360640190fd5b611f208b8b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f8f018190048102820181019092528d815292508d91508c908190840183828082843760009201919091525050604080516020601f8e018190048102820181019092528c815292508c91508b908190840183828082843760009201919091525050604080516020601f8d018190048102820181019092528b815292508b91508a908190840183828082843760009201919091525050604080516020601f8c018190048102820181019092528a815292508a91508990819084018382808284376000920191909152506126d892505050565b611f67576040805162461bcd60e51b815260206004820152601360248201527224b73b30b634b2103232b9b1b934b83a34b7b760691b604482015290519081900360640190fd5b6005546040805163416259d960e11b81523360048201819052915191926001600160a01b0316916382c4b3b291602480820192602092909190829003018186803b158015611fb457600080fd5b505afa158015611fc8573d6000803e3d6000fd5b505050506040513d6020811015611fde57600080fd5b5051612031576040805162461bcd60e51b815260206004820152601c60248201527f596f75206d75737420626520617574686f72697a656420666972737400000000604482015290519081900360640190fd5b6001600160a01b038181166000908152600160205260409020548116908e1614612084576001600160a01b03818116600090815260016020526040902080546001600160a01b031916918f169190911790555b6040518060a001604052808d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f8e018190048102820181019092528c815291810191908d908d9081908401838280828437600092019190915250505090825250604080516020601f8c018190048102820181019092528a815291810191908b908b9081908401838280828437600092019190915250505090825250604080516020601f8a0181900481028201810190925288815291810191908990899081908401838280828437600092019190915250505090825250604080516020601f88018190048102820181019092528681529181019190879087908190840183828082843760009201829052509390945250506001600160a01b0384168152600160208181526040909220845180519190920193506121e2928492019061325b565b5060208281015180516121fb926001850192019061325b565b506040820151805161221791600284019160209091019061325b565b506060820151805161223391600384019160209091019061325b565b506080820151805161224f91600484019160209091019061325b565b509050508c6001600160a01b0316816001600160a01b03167fb8421f65501371f54d58de1937ff1e1ccdb76423ef6f84acea1814a0f6362ca0426040518082815260200191505060405180910390a35060019c9b505050505050505050505050565b6001600160a01b038181166000908152600160209081526040808320815160c081019092528054948516825260609485948594859485949293909291830190600160a01b900460ff16600281111561230557fe5b600281111561231057fe5b8152602001600182016040518060a0016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156123bd5780601f10612392576101008083540402835291602001916123bd565b820191906000526020600020905b8154815290600101906020018083116123a057829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561245f5780601f106124345761010080835404028352916020019161245f565b820191906000526020600020905b81548152906001019060200180831161244257829003601f168201915b5050509183525050600282810180546040805160206001841615610100026000190190931694909404601f810183900483028501830190915280845293810193908301828280156124f15780601f106124c6576101008083540402835291602001916124f1565b820191906000526020600020905b8154815290600101906020018083116124d457829003601f168201915b505050918352505060038201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156125855780601f1061255a57610100808354040283529160200191612585565b820191906000526020600020905b81548152906001019060200180831161256857829003601f168201915b505050918352505060048201805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529382019392918301828280156126195780601f106125ee57610100808354040283529160200191612619565b820191906000526020600020905b8154815290600101906020018083116125fc57829003601f168201915b50505091909252505050815260068201546020808301919091526007830154604080840191909152600890930154606092830152928201518051938101519281015191810151608090910151939b929a50909850965090945092505050565b6060600380548060200260200160405190810160405280929190818152602001828054801561178b576020028201919060005260206000209081546001600160a01b0316815260019091019060200180831161176d575050505050905090565b600060468651111561272a576040805162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840dadedcd2d6cae440d8cadccee8d60531b604482015290519081900360640190fd5b610bb885511115612782576040805162461bcd60e51b815260206004820152601760248201527f496e76616c6964206964656e74697479206c656e677468000000000000000000604482015290519081900360640190fd5b608c845111156127d2576040805162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840eecac4e6d2e8ca40d8cadccee8d60531b604482015290519081900360640190fd5b608c83511115612820576040805162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840cadac2d2d840d8cadccee8d60631b604482015290519081900360640190fd5b61011882511115612871576040805162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840c8cae8c2d2d8e640d8cadccee8d60531b604482015290519081900360640190fd5b50600195945050505050565b3341146128be576040805162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015290519081900360640190fd5b43600090815260076020908152604080832083805290915290205460ff161561292e576040805162461bcd60e51b815260206004820152601960248201527f426c6f636b20697320616c726561647920726577617264656400000000000000604482015290519081900360640190fd5b60005460ff16612974576040805162461bcd60e51b815260206004820152600c60248201526b139bdd081a5b9a5d081e595d60a21b604482015290519081900360640190fd5b4360009081526007602090815260408083208380528252808320805460ff1916600190811790915533808552925282205490913491600160a01b900460ff1660028111156129be57fe5b14156129cb575050612a1b565b6129d6816000612d16565b6040805182815242602082015281516001600160a01b038516927f7dc4e5df59513708dca355b8706273a5df7b810a4cec8019f2a4b9bb166a1a04928290030190a250505b565b6001600160a01b03811660009081526001602052604090206006810154815460ff60a01b1916600160a11b17909155612a5582612b2b565b60035460011015612b2757612a6982612e9e565b600554604080516315ea278160e01b81526001600160a01b038581166004830152915191909216916315ea27819160248083019260209291908290030181600087803b158015612ab857600080fd5b505af1158015612acc573d6000803e3d6000fd5b505050506040513d6020811015612ae257600080fd5b50506040805182815242602082015281516001600160a01b038516927fa26de7ab324eac08c596549f421e5c8741213d237d2e9a2c9c0ebde0a7a849fe928290030190a25b5050565b60006001600160a01b038216600090815260016020526040902054600160a01b900460ff166002811115612b5b57fe5b1480612b6a5750600254600110155b15612b7457611041565b6001600160a01b0381166000908152600160205260409020600601548015612bfd57612ba08183612d16565b600454612bad9082612fd4565b6004556001600160a01b038216600090815260016020526040902060070154612bd69082612fd4565b6001600160a01b038316600090815260016020526040812060078101929092556006909101555b6040805182815242602082015281516001600160a01b038516927fe294e9d73f8eee23e21b2e1567960625a6b5d339cb127b55d0d09473a9951235928290030190a25050565b60005b600354811015612c9057816001600160a01b031660038281548110612c6757fe5b6000918252602090912001546001600160a01b03161415612c885750611041565b600101612c46565b50600380546001810182556000919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0383166001600160a01b031990911681179091556040805142815290517f1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a9181900360200190a250565b81612d2057612b27565b6000612d2b82613037565b905080612d385750612b27565b60008080612d4686856130d9565b9050612d5c612d558286613140565b8790613199565b925060005b600254811015612e3557600060028281548110612d7a57fe5b6000918252602090912001546001600160a01b0316905060026001600160a01b038216600090815260016020526040902054600160a01b900460ff166002811115612dc157fe5b14158015612de15750866001600160a01b0316816001600160a01b031614155b15612e2c576001600160a01b038116600090815260016020526040902060060154612e0c9084612fd4565b6001600160a01b0382166000908152600160205260409020600601559250825b50600101612d61565b50600083118015612e4e57506001600160a01b03821615155b15612e96576001600160a01b038216600090815260016020526040902060060154612e799084612fd4565b6001600160a01b0383166000908152600160205260409020600601555b505050505050565b60005b60035481108015612eb457506003546001105b15612b275760038181548110612ec657fe5b6000918252602090912001546001600160a01b0383811691161415612fcc57600354600019018114612f5957600380546000198101908110612f0457fe5b600091825260209091200154600380546001600160a01b039092169183908110612f2a57fe5b9060005260206000200160006101000a8154816001600160a01b0302191690836001600160a01b031602179055505b6003805480612f6457fe5b6000828152602090819020820160001990810180546001600160a01b03191690559091019091556040805142815290516001600160a01b038516927f7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed928290030190a2612b27565b600101612ea1565b60008282018381101561302e576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b90505b92915050565b60008060005b6002548110156130d25760006002828154811061305657fe5b6000918252602090912001546001600160a01b0316905060026001600160a01b038216600090815260016020526040902054600160a01b900460ff16600281111561309d57fe5b141580156130bd5750846001600160a01b0316816001600160a01b031614155b156130c9576001909201915b5060010161303d565b5092915050565b600080821161312f576040805162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015290519081900360640190fd5b81838161313857fe5b049392505050565b60008261314f57506000613031565b8282028284828161315c57fe5b041461302e5760405162461bcd60e51b81526004018080602001828103825260218152602001806133496021913960400191505060405180910390fd5b6000828211156131f0576040805162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015290519081900360640190fd5b50900390565b82805482825590600052602060002090810192821561324b579160200282015b8281111561324b57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190613216565b506132579291506132d7565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282613291576000855561324b565b82601f106132aa57805160ff191683800117855561324b565b8280016001018555821561324b579182015b8281111561324b5782518255916020019190600101906132bc565b5b8082111561325757600081556001016132d856fe596f75206d757374207761697420656e6f75676820626c6f636b7320746f20776974686472617720796f75722070726f66697473206166746572206c6174657374207769746864726177206f6620746869732076616c696461746f72536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f77596f7520617265206e6f742074686520666565207265636569766572206f6620746869732076616c696461746f72a2646970667358221220e3a67fa3efe316c9a46ed6da64aad2987616d569806632f4e9e21b1e0d68524664736f6c63430007060033",
}

// ValidatorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ValidatorsMetaData.ABI instead.
var ValidatorsABI = ValidatorsMetaData.ABI

// ValidatorsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ValidatorsMetaData.Bin instead.
var ValidatorsBin = ValidatorsMetaData.Bin

// DeployValidators deploys a new Ethereum contract, binding an instance of Validators to it.
func DeployValidators(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Validators, error) {
	parsed, err := ValidatorsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ValidatorsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
}

// Validators is an auto generated Go binding around an Ethereum contract.
type Validators struct {
	ValidatorsCaller     // Read-only binding to the contract
	ValidatorsTransactor // Write-only binding to the contract
	ValidatorsFilterer   // Log filterer for contract events
}

// ValidatorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ValidatorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ValidatorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ValidatorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ValidatorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ValidatorsSession struct {
	Contract     *Validators       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ValidatorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ValidatorsCallerSession struct {
	Contract *ValidatorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ValidatorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ValidatorsTransactorSession struct {
	Contract     *ValidatorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ValidatorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ValidatorsRaw struct {
	Contract *Validators // Generic contract binding to access the raw methods on
}

// ValidatorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ValidatorsCallerRaw struct {
	Contract *ValidatorsCaller // Generic read-only contract binding to access the raw methods on
}

// ValidatorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ValidatorsTransactorRaw struct {
	Contract *ValidatorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewValidators creates a new instance of Validators, bound to a specific deployed contract.
func NewValidators(address common.Address, backend bind.ContractBackend) (*Validators, error) {
	contract, err := bindValidators(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Validators{ValidatorsCaller: ValidatorsCaller{contract: contract}, ValidatorsTransactor: ValidatorsTransactor{contract: contract}, ValidatorsFilterer: ValidatorsFilterer{contract: contract}}, nil
}

// NewValidatorsCaller creates a new read-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsCaller(address common.Address, caller bind.ContractCaller) (*ValidatorsCaller, error) {
	contract, err := bindValidators(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsCaller{contract: contract}, nil
}

// NewValidatorsTransactor creates a new write-only instance of Validators, bound to a specific deployed contract.
func NewValidatorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ValidatorsTransactor, error) {
	contract, err := bindValidators(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ValidatorsTransactor{contract: contract}, nil
}

// NewValidatorsFilterer creates a new log filterer instance of Validators, bound to a specific deployed contract.
func NewValidatorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ValidatorsFilterer, error) {
	contract, err := bindValidators(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ValidatorsFilterer{contract: contract}, nil
}

// bindValidators binds a generic wrapper to an already deployed contract.
func bindValidators(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ValidatorsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.ValidatorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.ValidatorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Validators *ValidatorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Validators.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Validators *ValidatorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Validators *ValidatorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Validators.Contract.contract.Transact(opts, method, params...)
}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Validators *ValidatorsCaller) ProposalAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "ProposalAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Validators *ValidatorsSession) ProposalAddr() (common.Address, error) {
	return _Validators.Contract.ProposalAddr(&_Validators.CallOpts)
}

// ProposalAddr is a free data retrieval call binding the contract method 0x6233be5d.
//
// Solidity: function ProposalAddr() view returns(address)
func (_Validators *ValidatorsCallerSession) ProposalAddr() (common.Address, error) {
	return _Validators.Contract.ProposalAddr(&_Validators.CallOpts)
}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Validators *ValidatorsCaller) PunishContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "PunishContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Validators *ValidatorsSession) PunishContractAddr() (common.Address, error) {
	return _Validators.Contract.PunishContractAddr(&_Validators.CallOpts)
}

// PunishContractAddr is a free data retrieval call binding the contract method 0x1b5e358c.
//
// Solidity: function PunishContractAddr() view returns(address)
func (_Validators *ValidatorsCallerSession) PunishContractAddr() (common.Address, error) {
	return _Validators.Contract.PunishContractAddr(&_Validators.CallOpts)
}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Validators *ValidatorsCaller) ValidatorContractAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "ValidatorContractAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Validators *ValidatorsSession) ValidatorContractAddr() (common.Address, error) {
	return _Validators.Contract.ValidatorContractAddr(&_Validators.CallOpts)
}

// ValidatorContractAddr is a free data retrieval call binding the contract method 0x3a061bd3.
//
// Solidity: function ValidatorContractAddr() view returns(address)
func (_Validators *ValidatorsCallerSession) ValidatorContractAddr() (common.Address, error) {
	return _Validators.Contract.ValidatorContractAddr(&_Validators.CallOpts)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address)
func (_Validators *ValidatorsCaller) CurrentValidatorSet(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "currentValidatorSet", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address)
func (_Validators *ValidatorsSession) CurrentValidatorSet(arg0 *big.Int) (common.Address, error) {
	return _Validators.Contract.CurrentValidatorSet(&_Validators.CallOpts, arg0)
}

// CurrentValidatorSet is a free data retrieval call binding the contract method 0x6969a25c.
//
// Solidity: function currentValidatorSet(uint256 ) view returns(address)
func (_Validators *ValidatorsCallerSession) CurrentValidatorSet(arg0 *big.Int) (common.Address, error) {
	return _Validators.Contract.CurrentValidatorSet(&_Validators.CallOpts, arg0)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_Validators *ValidatorsCaller) GetActiveValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getActiveValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_Validators *ValidatorsSession) GetActiveValidators() ([]common.Address, error) {
	return _Validators.Contract.GetActiveValidators(&_Validators.CallOpts)
}

// GetActiveValidators is a free data retrieval call binding the contract method 0x9de70258.
//
// Solidity: function getActiveValidators() view returns(address[])
func (_Validators *ValidatorsCallerSession) GetActiveValidators() ([]common.Address, error) {
	return _Validators.Contract.GetActiveValidators(&_Validators.CallOpts)
}

// GetTopValidators is a free data retrieval call binding the contract method 0xafeea115.
//
// Solidity: function getTopValidators() view returns(address[])
func (_Validators *ValidatorsCaller) GetTopValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getTopValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetTopValidators is a free data retrieval call binding the contract method 0xafeea115.
//
// Solidity: function getTopValidators() view returns(address[])
func (_Validators *ValidatorsSession) GetTopValidators() ([]common.Address, error) {
	return _Validators.Contract.GetTopValidators(&_Validators.CallOpts)
}

// GetTopValidators is a free data retrieval call binding the contract method 0xafeea115.
//
// Solidity: function getTopValidators() view returns(address[])
func (_Validators *ValidatorsCallerSession) GetTopValidators() ([]common.Address, error) {
	return _Validators.Contract.GetTopValidators(&_Validators.CallOpts)
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address val) view returns(string, string, string, string, string)
func (_Validators *ValidatorsCaller) GetValidatorDescription(opts *bind.CallOpts, val common.Address) (string, string, string, string, string, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getValidatorDescription", val)

	if err != nil {
		return *new(string), *new(string), *new(string), *new(string), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(string)).(*string)

	return out0, out1, out2, out3, out4, err

}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address val) view returns(string, string, string, string, string)
func (_Validators *ValidatorsSession) GetValidatorDescription(val common.Address) (string, string, string, string, string, error) {
	return _Validators.Contract.GetValidatorDescription(&_Validators.CallOpts, val)
}

// GetValidatorDescription is a free data retrieval call binding the contract method 0xa43569b3.
//
// Solidity: function getValidatorDescription(address val) view returns(string, string, string, string, string)
func (_Validators *ValidatorsCallerSession) GetValidatorDescription(val common.Address) (string, string, string, string, string, error) {
	return _Validators.Contract.GetValidatorDescription(&_Validators.CallOpts, val)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address val) view returns(address, uint8, uint256, uint256, uint256)
func (_Validators *ValidatorsCaller) GetValidatorInfo(opts *bind.CallOpts, val common.Address) (common.Address, uint8, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "getValidatorInfo", val)

	if err != nil {
		return *new(common.Address), *new(uint8), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(uint8)).(*uint8)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address val) view returns(address, uint8, uint256, uint256, uint256)
func (_Validators *ValidatorsSession) GetValidatorInfo(val common.Address) (common.Address, uint8, *big.Int, *big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorInfo(&_Validators.CallOpts, val)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address val) view returns(address, uint8, uint256, uint256, uint256)
func (_Validators *ValidatorsCallerSession) GetValidatorInfo(val common.Address) (common.Address, uint8, *big.Int, *big.Int, *big.Int, error) {
	return _Validators.Contract.GetValidatorInfo(&_Validators.CallOpts, val)
}

// HighestValidatorsSet is a free data retrieval call binding the contract method 0x4b3d500b.
//
// Solidity: function highestValidatorsSet(uint256 ) view returns(address)
func (_Validators *ValidatorsCaller) HighestValidatorsSet(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "highestValidatorsSet", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestValidatorsSet is a free data retrieval call binding the contract method 0x4b3d500b.
//
// Solidity: function highestValidatorsSet(uint256 ) view returns(address)
func (_Validators *ValidatorsSession) HighestValidatorsSet(arg0 *big.Int) (common.Address, error) {
	return _Validators.Contract.HighestValidatorsSet(&_Validators.CallOpts, arg0)
}

// HighestValidatorsSet is a free data retrieval call binding the contract method 0x4b3d500b.
//
// Solidity: function highestValidatorsSet(uint256 ) view returns(address)
func (_Validators *ValidatorsCallerSession) HighestValidatorsSet(arg0 *big.Int) (common.Address, error) {
	return _Validators.Contract.HighestValidatorsSet(&_Validators.CallOpts, arg0)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Validators *ValidatorsCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Validators *ValidatorsSession) Initialized() (bool, error) {
	return _Validators.Contract.Initialized(&_Validators.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Validators *ValidatorsCallerSession) Initialized() (bool, error) {
	return _Validators.Contract.Initialized(&_Validators.CallOpts)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address who) view returns(bool)
func (_Validators *ValidatorsCaller) IsActiveValidator(opts *bind.CallOpts, who common.Address) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "isActiveValidator", who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address who) view returns(bool)
func (_Validators *ValidatorsSession) IsActiveValidator(who common.Address) (bool, error) {
	return _Validators.Contract.IsActiveValidator(&_Validators.CallOpts, who)
}

// IsActiveValidator is a free data retrieval call binding the contract method 0x40550a1c.
//
// Solidity: function isActiveValidator(address who) view returns(bool)
func (_Validators *ValidatorsCallerSession) IsActiveValidator(who common.Address) (bool, error) {
	return _Validators.Contract.IsActiveValidator(&_Validators.CallOpts, who)
}

// IsTopValidator is a free data retrieval call binding the contract method 0x98e3b626.
//
// Solidity: function isTopValidator(address who) view returns(bool)
func (_Validators *ValidatorsCaller) IsTopValidator(opts *bind.CallOpts, who common.Address) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "isTopValidator", who)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTopValidator is a free data retrieval call binding the contract method 0x98e3b626.
//
// Solidity: function isTopValidator(address who) view returns(bool)
func (_Validators *ValidatorsSession) IsTopValidator(who common.Address) (bool, error) {
	return _Validators.Contract.IsTopValidator(&_Validators.CallOpts, who)
}

// IsTopValidator is a free data retrieval call binding the contract method 0x98e3b626.
//
// Solidity: function isTopValidator(address who) view returns(bool)
func (_Validators *ValidatorsCallerSession) IsTopValidator(who common.Address) (bool, error) {
	return _Validators.Contract.IsTopValidator(&_Validators.CallOpts, who)
}

// TotalJailedHB is a free data retrieval call binding the contract method 0x1303f7cf.
//
// Solidity: function totalJailedHB() view returns(uint256)
func (_Validators *ValidatorsCaller) TotalJailedHB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "totalJailedHB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalJailedHB is a free data retrieval call binding the contract method 0x1303f7cf.
//
// Solidity: function totalJailedHB() view returns(uint256)
func (_Validators *ValidatorsSession) TotalJailedHB() (*big.Int, error) {
	return _Validators.Contract.TotalJailedHB(&_Validators.CallOpts)
}

// TotalJailedHB is a free data retrieval call binding the contract method 0x1303f7cf.
//
// Solidity: function totalJailedHB() view returns(uint256)
func (_Validators *ValidatorsCallerSession) TotalJailedHB() (*big.Int, error) {
	return _Validators.Contract.TotalJailedHB(&_Validators.CallOpts)
}

// ValidateDescription is a free data retrieval call binding the contract method 0xb6c88519.
//
// Solidity: function validateDescription(string moniker, string identity, string website, string email, string details) pure returns(bool)
func (_Validators *ValidatorsCaller) ValidateDescription(opts *bind.CallOpts, moniker string, identity string, website string, email string, details string) (bool, error) {
	var out []interface{}
	err := _Validators.contract.Call(opts, &out, "validateDescription", moniker, identity, website, email, details)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ValidateDescription is a free data retrieval call binding the contract method 0xb6c88519.
//
// Solidity: function validateDescription(string moniker, string identity, string website, string email, string details) pure returns(bool)
func (_Validators *ValidatorsSession) ValidateDescription(moniker string, identity string, website string, email string, details string) (bool, error) {
	return _Validators.Contract.ValidateDescription(&_Validators.CallOpts, moniker, identity, website, email, details)
}

// ValidateDescription is a free data retrieval call binding the contract method 0xb6c88519.
//
// Solidity: function validateDescription(string moniker, string identity, string website, string email, string details) pure returns(bool)
func (_Validators *ValidatorsCallerSession) ValidateDescription(moniker string, identity string, website string, email string, details string) (bool, error) {
	return _Validators.Contract.ValidateDescription(&_Validators.CallOpts, moniker, identity, website, email, details)
}

// CreateOrEditValidator is a paid mutator transaction binding the contract method 0xa406fcb7.
//
// Solidity: function createOrEditValidator(address feeAddr, string moniker, string identity, string website, string email, string details) returns(bool)
func (_Validators *ValidatorsTransactor) CreateOrEditValidator(opts *bind.TransactOpts, feeAddr common.Address, moniker string, identity string, website string, email string, details string) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "createOrEditValidator", feeAddr, moniker, identity, website, email, details)
}

// CreateOrEditValidator is a paid mutator transaction binding the contract method 0xa406fcb7.
//
// Solidity: function createOrEditValidator(address feeAddr, string moniker, string identity, string website, string email, string details) returns(bool)
func (_Validators *ValidatorsSession) CreateOrEditValidator(feeAddr common.Address, moniker string, identity string, website string, email string, details string) (*types.Transaction, error) {
	return _Validators.Contract.CreateOrEditValidator(&_Validators.TransactOpts, feeAddr, moniker, identity, website, email, details)
}

// CreateOrEditValidator is a paid mutator transaction binding the contract method 0xa406fcb7.
//
// Solidity: function createOrEditValidator(address feeAddr, string moniker, string identity, string website, string email, string details) returns(bool)
func (_Validators *ValidatorsTransactorSession) CreateOrEditValidator(feeAddr common.Address, moniker string, identity string, website string, email string, details string) (*types.Transaction, error) {
	return _Validators.Contract.CreateOrEditValidator(&_Validators.TransactOpts, feeAddr, moniker, identity, website, email, details)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0xd6c0edad.
//
// Solidity: function distributeBlockReward() payable returns()
func (_Validators *ValidatorsTransactor) DistributeBlockReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "distributeBlockReward")
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0xd6c0edad.
//
// Solidity: function distributeBlockReward() payable returns()
func (_Validators *ValidatorsSession) DistributeBlockReward() (*types.Transaction, error) {
	return _Validators.Contract.DistributeBlockReward(&_Validators.TransactOpts)
}

// DistributeBlockReward is a paid mutator transaction binding the contract method 0xd6c0edad.
//
// Solidity: function distributeBlockReward() payable returns()
func (_Validators *ValidatorsTransactorSession) DistributeBlockReward() (*types.Transaction, error) {
	return _Validators.Contract.DistributeBlockReward(&_Validators.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] vals) returns()
func (_Validators *ValidatorsTransactor) Initialize(opts *bind.TransactOpts, vals []common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "initialize", vals)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] vals) returns()
func (_Validators *ValidatorsSession) Initialize(vals []common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Initialize(&_Validators.TransactOpts, vals)
}

// Initialize is a paid mutator transaction binding the contract method 0xa224cee7.
//
// Solidity: function initialize(address[] vals) returns()
func (_Validators *ValidatorsTransactorSession) Initialize(vals []common.Address) (*types.Transaction, error) {
	return _Validators.Contract.Initialize(&_Validators.TransactOpts, vals)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address val) returns()
func (_Validators *ValidatorsTransactor) RemoveValidator(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "removeValidator", val)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address val) returns()
func (_Validators *ValidatorsSession) RemoveValidator(val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RemoveValidator(&_Validators.TransactOpts, val)
}

// RemoveValidator is a paid mutator transaction binding the contract method 0x40a141ff.
//
// Solidity: function removeValidator(address val) returns()
func (_Validators *ValidatorsTransactorSession) RemoveValidator(val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RemoveValidator(&_Validators.TransactOpts, val)
}

// RemoveValidatorIncoming is a paid mutator transaction binding the contract method 0x5dd09590.
//
// Solidity: function removeValidatorIncoming(address val) returns()
func (_Validators *ValidatorsTransactor) RemoveValidatorIncoming(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "removeValidatorIncoming", val)
}

// RemoveValidatorIncoming is a paid mutator transaction binding the contract method 0x5dd09590.
//
// Solidity: function removeValidatorIncoming(address val) returns()
func (_Validators *ValidatorsSession) RemoveValidatorIncoming(val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RemoveValidatorIncoming(&_Validators.TransactOpts, val)
}

// RemoveValidatorIncoming is a paid mutator transaction binding the contract method 0x5dd09590.
//
// Solidity: function removeValidatorIncoming(address val) returns()
func (_Validators *ValidatorsTransactorSession) RemoveValidatorIncoming(val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.RemoveValidatorIncoming(&_Validators.TransactOpts, val)
}

// TryActive is a paid mutator transaction binding the contract method 0xa0798862.
//
// Solidity: function tryActive(address validator) returns(bool)
func (_Validators *ValidatorsTransactor) TryActive(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "tryActive", validator)
}

// TryActive is a paid mutator transaction binding the contract method 0xa0798862.
//
// Solidity: function tryActive(address validator) returns(bool)
func (_Validators *ValidatorsSession) TryActive(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.TryActive(&_Validators.TransactOpts, validator)
}

// TryActive is a paid mutator transaction binding the contract method 0xa0798862.
//
// Solidity: function tryActive(address validator) returns(bool)
func (_Validators *ValidatorsTransactorSession) TryActive(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.TryActive(&_Validators.TransactOpts, validator)
}

// TryRemoveValidator is a paid mutator transaction binding the contract method 0xa1ff4655.
//
// Solidity: function tryRemoveValidator(address val) returns()
func (_Validators *ValidatorsTransactor) TryRemoveValidator(opts *bind.TransactOpts, val common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "tryRemoveValidator", val)
}

// TryRemoveValidator is a paid mutator transaction binding the contract method 0xa1ff4655.
//
// Solidity: function tryRemoveValidator(address val) returns()
func (_Validators *ValidatorsSession) TryRemoveValidator(val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.TryRemoveValidator(&_Validators.TransactOpts, val)
}

// TryRemoveValidator is a paid mutator transaction binding the contract method 0xa1ff4655.
//
// Solidity: function tryRemoveValidator(address val) returns()
func (_Validators *ValidatorsTransactorSession) TryRemoveValidator(val common.Address) (*types.Transaction, error) {
	return _Validators.Contract.TryRemoveValidator(&_Validators.TransactOpts, val)
}

// UpdateActiveValidatorSet is a paid mutator transaction binding the contract method 0x6846992a.
//
// Solidity: function updateActiveValidatorSet(address[] newSet, uint256 epoch) returns()
func (_Validators *ValidatorsTransactor) UpdateActiveValidatorSet(opts *bind.TransactOpts, newSet []common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "updateActiveValidatorSet", newSet, epoch)
}

// UpdateActiveValidatorSet is a paid mutator transaction binding the contract method 0x6846992a.
//
// Solidity: function updateActiveValidatorSet(address[] newSet, uint256 epoch) returns()
func (_Validators *ValidatorsSession) UpdateActiveValidatorSet(newSet []common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.UpdateActiveValidatorSet(&_Validators.TransactOpts, newSet, epoch)
}

// UpdateActiveValidatorSet is a paid mutator transaction binding the contract method 0x6846992a.
//
// Solidity: function updateActiveValidatorSet(address[] newSet, uint256 epoch) returns()
func (_Validators *ValidatorsTransactorSession) UpdateActiveValidatorSet(newSet []common.Address, epoch *big.Int) (*types.Transaction, error) {
	return _Validators.Contract.UpdateActiveValidatorSet(&_Validators.TransactOpts, newSet, epoch)
}

// WithdrawProfits is a paid mutator transaction binding the contract method 0x00362a77.
//
// Solidity: function withdrawProfits(address validator) returns(bool)
func (_Validators *ValidatorsTransactor) WithdrawProfits(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Validators.contract.Transact(opts, "withdrawProfits", validator)
}

// WithdrawProfits is a paid mutator transaction binding the contract method 0x00362a77.
//
// Solidity: function withdrawProfits(address validator) returns(bool)
func (_Validators *ValidatorsSession) WithdrawProfits(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.WithdrawProfits(&_Validators.TransactOpts, validator)
}

// WithdrawProfits is a paid mutator transaction binding the contract method 0x00362a77.
//
// Solidity: function withdrawProfits(address validator) returns(bool)
func (_Validators *ValidatorsTransactorSession) WithdrawProfits(validator common.Address) (*types.Transaction, error) {
	return _Validators.Contract.WithdrawProfits(&_Validators.TransactOpts, validator)
}

// ValidatorsLogActiveIterator is returned from FilterLogActive and is used to iterate over the raw logs and unpacked data for LogActive events raised by the Validators contract.
type ValidatorsLogActiveIterator struct {
	Event *ValidatorsLogActive // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogActiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogActive)
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
		it.Event = new(ValidatorsLogActive)
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
func (it *ValidatorsLogActiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogActiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogActive represents a LogActive event raised by the Validators contract.
type ValidatorsLogActive struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogActive is a free log retrieval operation binding the contract event 0x8bef9a500ef702fa4b7c82318f7b750176b75d33c8897ad10a35e5e5e4161362.
//
// Solidity: event LogActive(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogActive(opts *bind.FilterOpts, val []common.Address) (*ValidatorsLogActiveIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogActive", valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogActiveIterator{contract: _Validators.contract, event: "LogActive", logs: logs, sub: sub}, nil
}

// WatchLogActive is a free log subscription operation binding the contract event 0x8bef9a500ef702fa4b7c82318f7b750176b75d33c8897ad10a35e5e5e4161362.
//
// Solidity: event LogActive(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogActive(opts *bind.WatchOpts, sink chan<- *ValidatorsLogActive, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogActive", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogActive)
				if err := _Validators.contract.UnpackLog(event, "LogActive", log); err != nil {
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

// ParseLogActive is a log parse operation binding the contract event 0x8bef9a500ef702fa4b7c82318f7b750176b75d33c8897ad10a35e5e5e4161362.
//
// Solidity: event LogActive(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogActive(log types.Log) (*ValidatorsLogActive, error) {
	event := new(ValidatorsLogActive)
	if err := _Validators.contract.UnpackLog(event, "LogActive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogAddToTopValidatorsIterator is returned from FilterLogAddToTopValidators and is used to iterate over the raw logs and unpacked data for LogAddToTopValidators events raised by the Validators contract.
type ValidatorsLogAddToTopValidatorsIterator struct {
	Event *ValidatorsLogAddToTopValidators // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogAddToTopValidatorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogAddToTopValidators)
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
		it.Event = new(ValidatorsLogAddToTopValidators)
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
func (it *ValidatorsLogAddToTopValidatorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogAddToTopValidatorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogAddToTopValidators represents a LogAddToTopValidators event raised by the Validators contract.
type ValidatorsLogAddToTopValidators struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddToTopValidators is a free log retrieval operation binding the contract event 0x1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a.
//
// Solidity: event LogAddToTopValidators(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogAddToTopValidators(opts *bind.FilterOpts, val []common.Address) (*ValidatorsLogAddToTopValidatorsIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogAddToTopValidators", valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogAddToTopValidatorsIterator{contract: _Validators.contract, event: "LogAddToTopValidators", logs: logs, sub: sub}, nil
}

// WatchLogAddToTopValidators is a free log subscription operation binding the contract event 0x1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a.
//
// Solidity: event LogAddToTopValidators(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogAddToTopValidators(opts *bind.WatchOpts, sink chan<- *ValidatorsLogAddToTopValidators, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogAddToTopValidators", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogAddToTopValidators)
				if err := _Validators.contract.UnpackLog(event, "LogAddToTopValidators", log); err != nil {
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

// ParseLogAddToTopValidators is a log parse operation binding the contract event 0x1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a.
//
// Solidity: event LogAddToTopValidators(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogAddToTopValidators(log types.Log) (*ValidatorsLogAddToTopValidators, error) {
	event := new(ValidatorsLogAddToTopValidators)
	if err := _Validators.contract.UnpackLog(event, "LogAddToTopValidators", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogDistributeBlockRewardIterator is returned from FilterLogDistributeBlockReward and is used to iterate over the raw logs and unpacked data for LogDistributeBlockReward events raised by the Validators contract.
type ValidatorsLogDistributeBlockRewardIterator struct {
	Event *ValidatorsLogDistributeBlockReward // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogDistributeBlockRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogDistributeBlockReward)
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
		it.Event = new(ValidatorsLogDistributeBlockReward)
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
func (it *ValidatorsLogDistributeBlockRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogDistributeBlockRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogDistributeBlockReward represents a LogDistributeBlockReward event raised by the Validators contract.
type ValidatorsLogDistributeBlockReward struct {
	Coinbase    common.Address
	BlockReward *big.Int
	Time        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogDistributeBlockReward is a free log retrieval operation binding the contract event 0x7dc4e5df59513708dca355b8706273a5df7b810a4cec8019f2a4b9bb166a1a04.
//
// Solidity: event LogDistributeBlockReward(address indexed coinbase, uint256 blockReward, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogDistributeBlockReward(opts *bind.FilterOpts, coinbase []common.Address) (*ValidatorsLogDistributeBlockRewardIterator, error) {

	var coinbaseRule []interface{}
	for _, coinbaseItem := range coinbase {
		coinbaseRule = append(coinbaseRule, coinbaseItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogDistributeBlockReward", coinbaseRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogDistributeBlockRewardIterator{contract: _Validators.contract, event: "LogDistributeBlockReward", logs: logs, sub: sub}, nil
}

// WatchLogDistributeBlockReward is a free log subscription operation binding the contract event 0x7dc4e5df59513708dca355b8706273a5df7b810a4cec8019f2a4b9bb166a1a04.
//
// Solidity: event LogDistributeBlockReward(address indexed coinbase, uint256 blockReward, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogDistributeBlockReward(opts *bind.WatchOpts, sink chan<- *ValidatorsLogDistributeBlockReward, coinbase []common.Address) (event.Subscription, error) {

	var coinbaseRule []interface{}
	for _, coinbaseItem := range coinbase {
		coinbaseRule = append(coinbaseRule, coinbaseItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogDistributeBlockReward", coinbaseRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogDistributeBlockReward)
				if err := _Validators.contract.UnpackLog(event, "LogDistributeBlockReward", log); err != nil {
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

// ParseLogDistributeBlockReward is a log parse operation binding the contract event 0x7dc4e5df59513708dca355b8706273a5df7b810a4cec8019f2a4b9bb166a1a04.
//
// Solidity: event LogDistributeBlockReward(address indexed coinbase, uint256 blockReward, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogDistributeBlockReward(log types.Log) (*ValidatorsLogDistributeBlockReward, error) {
	event := new(ValidatorsLogDistributeBlockReward)
	if err := _Validators.contract.UnpackLog(event, "LogDistributeBlockReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogEditValidatorIterator is returned from FilterLogEditValidator and is used to iterate over the raw logs and unpacked data for LogEditValidator events raised by the Validators contract.
type ValidatorsLogEditValidatorIterator struct {
	Event *ValidatorsLogEditValidator // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogEditValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogEditValidator)
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
		it.Event = new(ValidatorsLogEditValidator)
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
func (it *ValidatorsLogEditValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogEditValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogEditValidator represents a LogEditValidator event raised by the Validators contract.
type ValidatorsLogEditValidator struct {
	Val  common.Address
	Fee  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogEditValidator is a free log retrieval operation binding the contract event 0xb8421f65501371f54d58de1937ff1e1ccdb76423ef6f84acea1814a0f6362ca0.
//
// Solidity: event LogEditValidator(address indexed val, address indexed fee, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogEditValidator(opts *bind.FilterOpts, val []common.Address, fee []common.Address) (*ValidatorsLogEditValidatorIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogEditValidator", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogEditValidatorIterator{contract: _Validators.contract, event: "LogEditValidator", logs: logs, sub: sub}, nil
}

// WatchLogEditValidator is a free log subscription operation binding the contract event 0xb8421f65501371f54d58de1937ff1e1ccdb76423ef6f84acea1814a0f6362ca0.
//
// Solidity: event LogEditValidator(address indexed val, address indexed fee, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogEditValidator(opts *bind.WatchOpts, sink chan<- *ValidatorsLogEditValidator, val []common.Address, fee []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogEditValidator", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogEditValidator)
				if err := _Validators.contract.UnpackLog(event, "LogEditValidator", log); err != nil {
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

// ParseLogEditValidator is a log parse operation binding the contract event 0xb8421f65501371f54d58de1937ff1e1ccdb76423ef6f84acea1814a0f6362ca0.
//
// Solidity: event LogEditValidator(address indexed val, address indexed fee, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogEditValidator(log types.Log) (*ValidatorsLogEditValidator, error) {
	event := new(ValidatorsLogEditValidator)
	if err := _Validators.contract.UnpackLog(event, "LogEditValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogRemoveFromTopValidatorsIterator is returned from FilterLogRemoveFromTopValidators and is used to iterate over the raw logs and unpacked data for LogRemoveFromTopValidators events raised by the Validators contract.
type ValidatorsLogRemoveFromTopValidatorsIterator struct {
	Event *ValidatorsLogRemoveFromTopValidators // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogRemoveFromTopValidatorsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogRemoveFromTopValidators)
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
		it.Event = new(ValidatorsLogRemoveFromTopValidators)
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
func (it *ValidatorsLogRemoveFromTopValidatorsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogRemoveFromTopValidatorsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogRemoveFromTopValidators represents a LogRemoveFromTopValidators event raised by the Validators contract.
type ValidatorsLogRemoveFromTopValidators struct {
	Val  common.Address
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRemoveFromTopValidators is a free log retrieval operation binding the contract event 0x7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed.
//
// Solidity: event LogRemoveFromTopValidators(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogRemoveFromTopValidators(opts *bind.FilterOpts, val []common.Address) (*ValidatorsLogRemoveFromTopValidatorsIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogRemoveFromTopValidators", valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogRemoveFromTopValidatorsIterator{contract: _Validators.contract, event: "LogRemoveFromTopValidators", logs: logs, sub: sub}, nil
}

// WatchLogRemoveFromTopValidators is a free log subscription operation binding the contract event 0x7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed.
//
// Solidity: event LogRemoveFromTopValidators(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogRemoveFromTopValidators(opts *bind.WatchOpts, sink chan<- *ValidatorsLogRemoveFromTopValidators, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogRemoveFromTopValidators", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogRemoveFromTopValidators)
				if err := _Validators.contract.UnpackLog(event, "LogRemoveFromTopValidators", log); err != nil {
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

// ParseLogRemoveFromTopValidators is a log parse operation binding the contract event 0x7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed.
//
// Solidity: event LogRemoveFromTopValidators(address indexed val, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogRemoveFromTopValidators(log types.Log) (*ValidatorsLogRemoveFromTopValidators, error) {
	event := new(ValidatorsLogRemoveFromTopValidators)
	if err := _Validators.contract.UnpackLog(event, "LogRemoveFromTopValidators", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogRemoveValidatorIterator is returned from FilterLogRemoveValidator and is used to iterate over the raw logs and unpacked data for LogRemoveValidator events raised by the Validators contract.
type ValidatorsLogRemoveValidatorIterator struct {
	Event *ValidatorsLogRemoveValidator // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogRemoveValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogRemoveValidator)
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
		it.Event = new(ValidatorsLogRemoveValidator)
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
func (it *ValidatorsLogRemoveValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogRemoveValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogRemoveValidator represents a LogRemoveValidator event raised by the Validators contract.
type ValidatorsLogRemoveValidator struct {
	Val  common.Address
	Hb   *big.Int
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRemoveValidator is a free log retrieval operation binding the contract event 0xa26de7ab324eac08c596549f421e5c8741213d237d2e9a2c9c0ebde0a7a849fe.
//
// Solidity: event LogRemoveValidator(address indexed val, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogRemoveValidator(opts *bind.FilterOpts, val []common.Address) (*ValidatorsLogRemoveValidatorIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogRemoveValidator", valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogRemoveValidatorIterator{contract: _Validators.contract, event: "LogRemoveValidator", logs: logs, sub: sub}, nil
}

// WatchLogRemoveValidator is a free log subscription operation binding the contract event 0xa26de7ab324eac08c596549f421e5c8741213d237d2e9a2c9c0ebde0a7a849fe.
//
// Solidity: event LogRemoveValidator(address indexed val, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogRemoveValidator(opts *bind.WatchOpts, sink chan<- *ValidatorsLogRemoveValidator, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogRemoveValidator", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogRemoveValidator)
				if err := _Validators.contract.UnpackLog(event, "LogRemoveValidator", log); err != nil {
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

// ParseLogRemoveValidator is a log parse operation binding the contract event 0xa26de7ab324eac08c596549f421e5c8741213d237d2e9a2c9c0ebde0a7a849fe.
//
// Solidity: event LogRemoveValidator(address indexed val, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogRemoveValidator(log types.Log) (*ValidatorsLogRemoveValidator, error) {
	event := new(ValidatorsLogRemoveValidator)
	if err := _Validators.contract.UnpackLog(event, "LogRemoveValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogRemoveValidatorIncomingIterator is returned from FilterLogRemoveValidatorIncoming and is used to iterate over the raw logs and unpacked data for LogRemoveValidatorIncoming events raised by the Validators contract.
type ValidatorsLogRemoveValidatorIncomingIterator struct {
	Event *ValidatorsLogRemoveValidatorIncoming // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogRemoveValidatorIncomingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogRemoveValidatorIncoming)
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
		it.Event = new(ValidatorsLogRemoveValidatorIncoming)
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
func (it *ValidatorsLogRemoveValidatorIncomingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogRemoveValidatorIncomingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogRemoveValidatorIncoming represents a LogRemoveValidatorIncoming event raised by the Validators contract.
type ValidatorsLogRemoveValidatorIncoming struct {
	Val  common.Address
	Hb   *big.Int
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogRemoveValidatorIncoming is a free log retrieval operation binding the contract event 0xe294e9d73f8eee23e21b2e1567960625a6b5d339cb127b55d0d09473a9951235.
//
// Solidity: event LogRemoveValidatorIncoming(address indexed val, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogRemoveValidatorIncoming(opts *bind.FilterOpts, val []common.Address) (*ValidatorsLogRemoveValidatorIncomingIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogRemoveValidatorIncoming", valRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogRemoveValidatorIncomingIterator{contract: _Validators.contract, event: "LogRemoveValidatorIncoming", logs: logs, sub: sub}, nil
}

// WatchLogRemoveValidatorIncoming is a free log subscription operation binding the contract event 0xe294e9d73f8eee23e21b2e1567960625a6b5d339cb127b55d0d09473a9951235.
//
// Solidity: event LogRemoveValidatorIncoming(address indexed val, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogRemoveValidatorIncoming(opts *bind.WatchOpts, sink chan<- *ValidatorsLogRemoveValidatorIncoming, val []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogRemoveValidatorIncoming", valRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogRemoveValidatorIncoming)
				if err := _Validators.contract.UnpackLog(event, "LogRemoveValidatorIncoming", log); err != nil {
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

// ParseLogRemoveValidatorIncoming is a log parse operation binding the contract event 0xe294e9d73f8eee23e21b2e1567960625a6b5d339cb127b55d0d09473a9951235.
//
// Solidity: event LogRemoveValidatorIncoming(address indexed val, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogRemoveValidatorIncoming(log types.Log) (*ValidatorsLogRemoveValidatorIncoming, error) {
	event := new(ValidatorsLogRemoveValidatorIncoming)
	if err := _Validators.contract.UnpackLog(event, "LogRemoveValidatorIncoming", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogUpdateValidatorIterator is returned from FilterLogUpdateValidator and is used to iterate over the raw logs and unpacked data for LogUpdateValidator events raised by the Validators contract.
type ValidatorsLogUpdateValidatorIterator struct {
	Event *ValidatorsLogUpdateValidator // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogUpdateValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogUpdateValidator)
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
		it.Event = new(ValidatorsLogUpdateValidator)
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
func (it *ValidatorsLogUpdateValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogUpdateValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogUpdateValidator represents a LogUpdateValidator event raised by the Validators contract.
type ValidatorsLogUpdateValidator struct {
	NewSet []common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogUpdateValidator is a free log retrieval operation binding the contract event 0xeacea8f3c22f06c0b18306bdb04d0a967255129e8ce0094debb0a0ff89d006b5.
//
// Solidity: event LogUpdateValidator(address[] newSet)
func (_Validators *ValidatorsFilterer) FilterLogUpdateValidator(opts *bind.FilterOpts) (*ValidatorsLogUpdateValidatorIterator, error) {

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogUpdateValidator")
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogUpdateValidatorIterator{contract: _Validators.contract, event: "LogUpdateValidator", logs: logs, sub: sub}, nil
}

// WatchLogUpdateValidator is a free log subscription operation binding the contract event 0xeacea8f3c22f06c0b18306bdb04d0a967255129e8ce0094debb0a0ff89d006b5.
//
// Solidity: event LogUpdateValidator(address[] newSet)
func (_Validators *ValidatorsFilterer) WatchLogUpdateValidator(opts *bind.WatchOpts, sink chan<- *ValidatorsLogUpdateValidator) (event.Subscription, error) {

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogUpdateValidator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogUpdateValidator)
				if err := _Validators.contract.UnpackLog(event, "LogUpdateValidator", log); err != nil {
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

// ParseLogUpdateValidator is a log parse operation binding the contract event 0xeacea8f3c22f06c0b18306bdb04d0a967255129e8ce0094debb0a0ff89d006b5.
//
// Solidity: event LogUpdateValidator(address[] newSet)
func (_Validators *ValidatorsFilterer) ParseLogUpdateValidator(log types.Log) (*ValidatorsLogUpdateValidator, error) {
	event := new(ValidatorsLogUpdateValidator)
	if err := _Validators.contract.UnpackLog(event, "LogUpdateValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ValidatorsLogWithdrawProfitsIterator is returned from FilterLogWithdrawProfits and is used to iterate over the raw logs and unpacked data for LogWithdrawProfits events raised by the Validators contract.
type ValidatorsLogWithdrawProfitsIterator struct {
	Event *ValidatorsLogWithdrawProfits // Event containing the contract specifics and raw log

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
func (it *ValidatorsLogWithdrawProfitsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ValidatorsLogWithdrawProfits)
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
		it.Event = new(ValidatorsLogWithdrawProfits)
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
func (it *ValidatorsLogWithdrawProfitsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ValidatorsLogWithdrawProfitsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ValidatorsLogWithdrawProfits represents a LogWithdrawProfits event raised by the Validators contract.
type ValidatorsLogWithdrawProfits struct {
	Val  common.Address
	Fee  common.Address
	Hb   *big.Int
	Time *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogWithdrawProfits is a free log retrieval operation binding the contract event 0x51a69b4502f660774c9339825c7b5adbf0b8622289134647e29728ec5d9b3bb9.
//
// Solidity: event LogWithdrawProfits(address indexed val, address indexed fee, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) FilterLogWithdrawProfits(opts *bind.FilterOpts, val []common.Address, fee []common.Address) (*ValidatorsLogWithdrawProfitsIterator, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Validators.contract.FilterLogs(opts, "LogWithdrawProfits", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return &ValidatorsLogWithdrawProfitsIterator{contract: _Validators.contract, event: "LogWithdrawProfits", logs: logs, sub: sub}, nil
}

// WatchLogWithdrawProfits is a free log subscription operation binding the contract event 0x51a69b4502f660774c9339825c7b5adbf0b8622289134647e29728ec5d9b3bb9.
//
// Solidity: event LogWithdrawProfits(address indexed val, address indexed fee, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) WatchLogWithdrawProfits(opts *bind.WatchOpts, sink chan<- *ValidatorsLogWithdrawProfits, val []common.Address, fee []common.Address) (event.Subscription, error) {

	var valRule []interface{}
	for _, valItem := range val {
		valRule = append(valRule, valItem)
	}
	var feeRule []interface{}
	for _, feeItem := range fee {
		feeRule = append(feeRule, feeItem)
	}

	logs, sub, err := _Validators.contract.WatchLogs(opts, "LogWithdrawProfits", valRule, feeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ValidatorsLogWithdrawProfits)
				if err := _Validators.contract.UnpackLog(event, "LogWithdrawProfits", log); err != nil {
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

// ParseLogWithdrawProfits is a log parse operation binding the contract event 0x51a69b4502f660774c9339825c7b5adbf0b8622289134647e29728ec5d9b3bb9.
//
// Solidity: event LogWithdrawProfits(address indexed val, address indexed fee, uint256 hb, uint256 time)
func (_Validators *ValidatorsFilterer) ParseLogWithdrawProfits(log types.Log) (*ValidatorsLogWithdrawProfits, error) {
	event := new(ValidatorsLogWithdrawProfits)
	if err := _Validators.contract.UnpackLog(event, "LogWithdrawProfits", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
