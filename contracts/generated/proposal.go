// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package generated

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
	Bin: "0x6080604052348015600e575f5ffd5b5060cd80601a5f395ff3fe6080604052348015600e575f5ffd5b50600436106044575f3560e01c8063158ef93e1460485780631b5e358c1460685780633a061bd31460875780636233be5d14608f575b5f5ffd5b5f5460539060ff1681565b60405190151581526020015b60405180910390f35b607061f00181565b6040516001600160a01b039091168152602001605f565b607061f00081565b607061f0028156fea2646970667358221220d8088d76208a680a0ffdcd825505fcf71115d2637088f0fc24ba1911a224838764736f6c634300081e0033",
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
	Bin: "0x6080604052348015600e575f5ffd5b506117e48061001c5f395ff3fe608060405234801561000f575f5ffd5b5060043610610127575f3560e01c806382c4b3b2116100a9578063cb1ea7251161006e578063cb1ea7251461030a578063d24806eb14610313578063d51cade814610326578063d63e6ce714610339578063e823c81414610342575f5ffd5b806382c4b3b2146102a457806394522b6d146102c6578063a224cee7146102cf578063a3dcb4d2146102e4578063a4c4d922146102f7575f5ffd5b806332ed5b12116100ef57806332ed5b12146102045780633a061bd31461022b57806344c1aa99146102345780634c6b25b11461023d5780636233be5d1461029b575f5ffd5b8063158ef93e1461012b57806315ea27811461014c5780631b5e358c1461015f5780631db5ade8146101805780632897183d146101ed575b5f5ffd5b5f546101379060ff1681565b60405190151581526020015b60405180910390f35b61013761015a36600461126a565b61034b565b61016861f00181565b6040516001600160a01b039091168152602001610143565b6101c661018e36600461128c565b600b60209081525f92835260408084209091529082529020805460018201546002909201546001600160a01b03909116919060ff1683565b604080516001600160a01b0390941684526020840192909252151590820152606001610143565b6101f660045481565b604051908152602001610143565b6102176102123660046112b6565b610404565b6040516101439897969594939291906112cd565b61016861f00081565b6101f660035481565b61027861024b3660046112b6565b600a6020525f908152604090205461ffff8082169162010000810490911690640100000000900460ff1683565b6040805161ffff9485168152939092166020840152151590820152606001610143565b61016861f00281565b6101376102b236600461126a565b60086020525f908152604090205460ff1681565b6101f660055481565b6102e26102dd36600461134a565b6104e0565b005b600754610168906001600160a01b031681565b6101376103053660046113c8565b61067d565b6101f660025481565b6101376103213660046113f6565b610d33565b610137610334366004611416565b610e91565b6101f660065481565b6101f660015481565b5f3361f000146103a25760405162461bcd60e51b815260206004820152601860248201527f56616c696461746f727320636f6e7472616374206f6e6c79000000000000000060448201526064015b60405180910390fd5b6001600160a01b0382165f8181526008602052604090819020805460ff19169055517f4e0b191f7f5c32b1b5e3704b68874b1a3980147cae00be8ece271bfb5b92c07a906103f39042815260200190565b60405180910390a25060015b919050565b60096020525f9081526040902080546001820154600283015460038401546004850180546001600160a01b03958616969495939493831693600160a01b90930460ff16929190610453906114a4565b80601f016020809104026020016040519081016040528092919081815260200182805461047f906114a4565b80156104ca5780601f106104a1576101008083540402835291602001916104ca565b820191905f5260205f20905b8154815290600101906020018083116104ad57829003601f168201915b5050505050908060050154908060060154905088565b5f5460ff16156105285760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b6044820152606401610399565b600c80546001600160a01b03191661f0001790555f5b8181101561061f575f838383818110610559576105596114dc565b905060200201602081019061056e919061126a565b6001600160a01b0316036105c45760405162461bcd60e51b815260206004820152601960248201527f496e76616c69642076616c696461746f722061646472657373000000000000006044820152606401610399565b600160085f8585858181106105db576105db6114dc565b90506020020160208101906105f0919061126a565b6001600160a01b0316815260208101919091526040015f20805460ff191691151591909117905560010161053e565b505062093a806001908155601860028190556030600355600455620151806005555f805460ff19169091179055506301e13380600655600780546001600160a01b031916739014b4db9d30ced67db9d6b096f5dcdba28ce639179055565b600c54604051631015428760e21b81523360048201525f916001600160a01b0316906340550a1c90602401602060405180830381865afa1580156106c3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106e791906114f0565b6107245760405162461bcd60e51b815260206004820152600e60248201526d56616c696461746f72206f6e6c7960901b6044820152606401610399565b5f8381526009602052604081206001015490036107785760405162461bcd60e51b8152602060048201526012602482015271141c9bdc1bdcd85b081b9bdd08195e1a5cdd60721b6044820152606401610399565b335f908152600b60209081526040808320868452909152902060010154156107ee5760405162461bcd60e51b815260206004820152602360248201527f596f752063616e277420766f746520666f7220612070726f706f73616c20747760448201526269636560e81b6064820152608401610399565b600180545f8581526009602052604090209091015461080d919061151f565b421061084e5760405162461bcd60e51b815260206004820152601060248201526f141c9bdc1bdcd85b08195e1c1a5c995960821b6044820152606401610399565b335f818152600b60209081526040808320878452825291829020426001820181905581546001600160a01b031916851782556002909101805460ff191687151590811790915583519081529182015285917f6c59bda68cac318717c60c7c9635a78a0f0613f9887cc18a7157f5745a86d14e910160405180910390a38115610913575f838152600a60205260409020546108ed9061ffff166001611532565b5f848152600a60205260409020805461ffff191661ffff92909216919091179055610961565b5f838152600a60205260409020546109369062010000900461ffff166001611532565b5f848152600a60205260409020805461ffff92909216620100000263ffff0000199092169190911790555b5f838152600a6020526040902054640100000000900460ff161561098757506001610d2d565b600c54604080516313bce04b60e31b815290516002926001600160a01b031691639de70258916004808301925f9291908290030181865afa1580156109ce573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526109f5919081019061156b565b51610a009190611636565b610a0b90600161151f565b5f848152600a602052604090205461ffff1610610c2b575f838152600a60209081526040808320805464ff0000000019166401000000001790556009909152902060020154600103610bae575f83815260096020526040902060030154600160a01b900460ff1615610b1e575f83815260096020818152604080842060030180546001600160a01b03908116865260088452828620805460ff19166001179055600c5495899052939092529054905163503cc43160e11b8152908216600482015291169063a0798862906024016020604051808303815f875af1158015610af4573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b1891906114f0565b50610be9565b5f83815260096020818152604080842060030180546001600160a01b03908116865260088452828620805460ff19169055600c5495899052939092529054905163a1ff465560e01b8152908216600482015291169063a1ff4655906024015f604051808303815f87803b158015610b93575f5ffd5b505af1158015610ba5573d5f5f3e3d5ffd5b50505050610be9565b5f8381526009602052604090206002908101549003610be9575f8381526009602052604090206005810154600690910154610be99190611186565b827f90d2e923947d9356c1c04391cb9e2e9c5d4ad6c165a849787b0c7569bbe99e2442604051610c1b91815260200190565b60405180910390a2506001610d2d565b600c54604080516313bce04b60e31b815290516002926001600160a01b031691639de70258916004808301925f9291908290030181865afa158015610c72573d5f5f3e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052610c99919081019061156b565b51610ca49190611636565b610caf90600161151f565b5f848152600a602052604090205462010000900461ffff1610610d29575f838152600a602052604090819020805464ff0000000019166401000000001790555183907f36bdb56d707cdf53eadffe319a71ddf97736be67b8caab47b7720201a6b65ca090610d209042815260200190565b60405180910390a25b5060015b92915050565b6040516bffffffffffffffffffffffff193360601b16602082015260348101839052605481018290524260748201525f908190609401604051602081830303815290604052805190602001209050610d89611201565b33815260c0810185905260e08101849052426020808301918252600260408085018281525f87815260099094529220845181546001600160a01b0319166001600160a01b039182161782559351600182015591519082015560608301516003820180546080860151929094166001600160a81b031990941693909317600160a01b911515919091021790915560a08201518291906004820190610e2c90826116a1565b5060c0820151600582015560e090910151600690910155604080518681526020810186905242818301529051339184917f8bfc061277ae1778974ada10db7f9664ab1d67c455c025c025b438c52c69d1819181900360600190a3506001949350505050565b6001600160a01b0384165f9081526008602052604081205460ff16158015610eb65750835b80610ee157506001600160a01b0385165f9081526008602052604090205460ff168015610ee1575083155b610f55576040805162461bcd60e51b81526020600482015260248101919091527f43616e74227420616464206120616c726561647920657869737420647374206f60448201527f722043616e7422742072656d6f76652061206e6f7420706173736564206473746064820152608401610399565b5f338686868642604051602001610f719695949392919061175c565b60408051601f1981840301815291905280516020909101209050610bb8831115610fd05760405162461bcd60e51b815260206004820152601060248201526f44657461696c7320746f6f206c6f6e6760801b6044820152606401610399565b5f818152600960205260409020600101541561102e5760405162461bcd60e51b815260206004820152601760248201527f50726f706f73616c20616c7265616479206578697374730000000000000000006044820152606401610399565b611036611201565b3381526001600160a01b03871660608201528515156080820152604080516020601f87018190048102820181019092528581529086908690819084018382808284375f92018290525060a0860194855242602080880191825260016040808a018281528b86526009909352909320885181546001600160a01b0319166001600160a01b03918216178255925193810193909355516002830155606087015160038301805460808a0151929093166001600160a81b031990931692909217600160a01b911515919091021790559351859493506004840192506111199150826116a1565b5060c0820151600582015560e0909101516006909101556040805187151581524260208201526001600160a01b03891691339185917f1af05d46b8c1ec021d82b7128cff40e91a1c2337deffc010df48eeddef8da56c910160405180910390a45060019695505050505050565b815f036111935760015550565b816001036111a15760025550565b816002036111af5760035550565b816003036111bd5760045550565b816004036111cb5760055550565b816005036111d95760065550565b816006036111fd57600780546001600160a01b0319166001600160a01b0383161790555b5050565b6040518061010001604052805f6001600160a01b031681526020015f81526020015f81526020015f6001600160a01b031681526020015f15158152602001606081526020015f81526020015f81525090565b6001600160a01b0381168114611267575f5ffd5b50565b5f6020828403121561127a575f5ffd5b813561128581611253565b9392505050565b5f5f6040838503121561129d575f5ffd5b82356112a881611253565b946020939093013593505050565b5f602082840312156112c6575f5ffd5b5035919050565b60018060a01b038916815287602082015286604082015260018060a01b0386166060820152841515608082015261010060a08201525f845180610100840152806020870161012085015e5f6101208285010152610120601f19601f8301168401019150508360c08301528260e08301529998505050505050505050565b5f5f6020838503121561135b575f5ffd5b823567ffffffffffffffff811115611371575f5ffd5b8301601f81018513611381575f5ffd5b803567ffffffffffffffff811115611397575f5ffd5b8560208260051b84010111156113ab575f5ffd5b6020919091019590945092505050565b8015158114611267575f5ffd5b5f5f604083850312156113d9575f5ffd5b8235915060208301356113eb816113bb565b809150509250929050565b5f5f60408385031215611407575f5ffd5b50508035926020909101359150565b5f5f5f5f60608587031215611429575f5ffd5b843561143481611253565b93506020850135611444816113bb565b9250604085013567ffffffffffffffff81111561145f575f5ffd5b8501601f8101871361146f575f5ffd5b803567ffffffffffffffff811115611485575f5ffd5b876020828401011115611496575f5ffd5b949793965060200194505050565b600181811c908216806114b857607f821691505b6020821081036114d657634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215611500575f5ffd5b8151611285816113bb565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610d2d57610d2d61150b565b61ffff8181168382160190811115610d2d57610d2d61150b565b634e487b7160e01b5f52604160045260245ffd5b80516103ff81611253565b5f6020828403121561157b575f5ffd5b815167ffffffffffffffff811115611591575f5ffd5b8201601f810184136115a1575f5ffd5b805167ffffffffffffffff8111156115bb576115bb61154c565b8060051b604051601f19603f830116810181811067ffffffffffffffff821117156115e8576115e861154c565b604052918252602081840181019290810187841115611605575f5ffd5b6020850194505b8385101561162b5761161d85611560565b81526020948501940161160c565b509695505050505050565b5f8261165057634e487b7160e01b5f52601260045260245ffd5b500490565b601f82111561169c57805f5260205f20601f840160051c8101602085101561167a5750805b601f840160051c820191505b81811015611699575f8155600101611686565b50505b505050565b815167ffffffffffffffff8111156116bb576116bb61154c565b6116cf816116c984546114a4565b84611655565b6020601f821160018114611701575f83156116ea5750848201515b5f19600385901b1c1916600184901b178455611699565b5f84815260208120601f198516915b828110156117305787850151825560209485019460019092019101611710565b508482101561174d57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6bffffffffffffffffffffffff198760601b1681526bffffffffffffffffffffffff198660601b16601482015284151560f81b602882015282846029830137602992019182015260490194935050505056fea26469706673582212205114b82925895ded73f79d5114182d17de9c6a2e682018f89413c526009f393c64736f6c634300081e0033",
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
	Bin: "0x6080604052348015600e575f5ffd5b50610dca8061001c5f395ff3fe608060405234801561000f575f5ffd5b50600436106100a6575f3560e01c806363e1d4511161006e57806363e1d451146101345780638129fc1c14610147578063d93d2cb914610151578063e0d8ea5314610164578063ea7221a11461016c578063f62af26c1461017f575f5ffd5b8063158ef93e146100aa5780631b5e358c146100cb57806332f3c17f146100ec5780633a061bd3146101225780636233be5d1461012b575b5f5ffd5b5f546100b69060ff1681565b60405190151581526020015b60405180910390f35b6100d461f00181565b6040516001600160a01b0390911681526020016100c2565b6101146100fa366004610c6c565b6001600160a01b03165f9081526002602052604090205490565b6040519081526020016100c2565b6100d461f00081565b6100d461f00281565b6100b6610142366004610c6c565b610192565b61014f6103b5565b005b61014f61015f366004610c99565b610429565b600354610114565b61014f61017a366004610c6c565b61086a565b6100d461018d366004610c99565b610c44565b5f805460ff166101bd5760405162461bcd60e51b81526004016101b490610cb0565b60405180910390fd5b3361f0001461020e5760405162461bcd60e51b815260206004820152601860248201527f56616c696461746f727320636f6e7472616374206f6e6c79000000000000000060448201526064016101b4565b6001600160a01b0382165f9081526002602052604090205415610244576001600160a01b0382165f908152600260205260408120555b6001600160a01b0382165f908152600260208190526040909120015460ff168015610270575060035415155b156103ad5760035461028490600190610cea565b6001600160a01b0383165f908152600260205260409020600101541461035057600380545f91906102b790600190610cea565b815481106102c7576102c7610d03565b5f9182526020808320909101546001600160a01b038681168452600290925260409092206001015460038054929093169350839291811061030a5761030a610d03565b5f91825260208083209190910180546001600160a01b0319166001600160a01b039485161790558583168252600290526040808220600190810154949093168252902001555b600380548061036157610361610d17565b5f828152602080822083015f1990810180546001600160a01b03191690559092019092556001600160a01b038416825260029081905260408220600181019290925501805460ff191690555b506001919050565b5f5460ff16156103fd5760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b60448201526064016101b4565b5f8054600180546001600160a01b03191661f00217905562f000016001600160a81b0319909116179055565b3341146104655760405162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b60448201526064016101b4565b435f9081526005602052604090205460ff16156104b85760405162461bcd60e51b8152602060048201526011602482015270105b1c9958591e48191958dc99585cd959607a1b60448201526064016101b4565b5f5460ff166104d95760405162461bcd60e51b81526004016101b490610cb0565b806104e48143610d3f565b156105245760405162461bcd60e51b815260206004820152601060248201526f426c6f636b2065706f6368206f6e6c7960801b60448201526064016101b4565b435f908152600560205260409020805460ff1916600117905560035415610866575f5b60035481101561083c5760015f9054906101000a90046001600160a01b03166001600160a01b0316632897183d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105a1573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105c59190610d52565b60015f9054906101000a90046001600160a01b03166001600160a01b03166344c1aa996040518163ffffffff1660e01b8152600401602060405180830381865afa158015610615573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106399190610d52565b6106439190610d69565b60025f6003848154811061065957610659610d03565b5f9182526020808320909101546001600160a01b0316835282019290925260400190205411156107f85760015f9054906101000a90046001600160a01b03166001600160a01b0316632897183d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106d3573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106f79190610d52565b60015f9054906101000a90046001600160a01b03166001600160a01b03166344c1aa996040518163ffffffff1660e01b8152600401602060405180830381865afa158015610747573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061076b9190610d52565b6107759190610d69565b60025f6003848154811061078b5761078b610d03565b5f9182526020808320909101546001600160a01b031683528201929092526040019020546107b99190610cea565b60025f600384815481106107cf576107cf610d03565b5f9182526020808320909101546001600160a01b03168352820192909252604001902055610834565b5f60025f6003848154811061080f5761080f610d03565b5f9182526020808320909101546001600160a01b031683528201929092526040019020555b600101610547565b506040517f181d51be54e8e8eaca6eae0eab32d4162099236bd519e7238d015d0870db4641905f90a15b5050565b3341146108a65760405162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b60448201526064016101b4565b5f5460ff166108c75760405162461bcd60e51b81526004016101b490610cb0565b435f9081526004602052604090205460ff16156109195760405162461bcd60e51b815260206004820152601060248201526f105b1c9958591e481c1d5b9a5cda195960821b60448201526064016101b4565b435f908152600460209081526040808320805460ff191660011790556001600160a01b0384168352600291829052909120015460ff166109c257600380546001600160a01b0383165f818152600260208190526040822060018082018690558086019096557fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b90940180546001600160a01b0319168417905591905201805460ff191690911790555b6001600160a01b0381165f9081526002602052604081208054916109e583610d7c565b919050555060015f9054906101000a90046001600160a01b03166001600160a01b03166344c1aa996040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a3a573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a5e9190610d52565b6001600160a01b0382165f90815260026020526040902054610a809190610d3f565b5f03610b02575f546040516340a141ff60e01b81526001600160a01b038381166004830152610100909204909116906340a141ff906024015f604051808303815f87803b158015610acf575f5ffd5b505af1158015610ae1573d5f5f3e3d5ffd5b5050506001600160a01b0382165f9081526002602052604081205550610bfe565b60015f9054906101000a90046001600160a01b03166001600160a01b031663cb1ea7256040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b52573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b769190610d52565b6001600160a01b0382165f90815260026020526040902054610b989190610d3f565b5f03610bfe575f546040516305dd095960e41b81526001600160a01b03838116600483015261010090920490911690635dd09590906024015f604051808303815f87803b158015610be7575f5ffd5b505af1158015610bf9573d5f5f3e3d5ffd5b505050505b806001600160a01b03167f770e0cca42c35d00240986ce8d3ed438be04663c91dac6576b79537d7c180f1e42604051610c3991815260200190565b60405180910390a250565b60038181548110610c53575f80fd5b5f918252602090912001546001600160a01b0316905081565b5f60208284031215610c7c575f5ffd5b81356001600160a01b0381168114610c92575f5ffd5b9392505050565b5f60208284031215610ca9575f5ffd5b5035919050565b6020808252600c908201526b139bdd081a5b9a5d081e595d60a21b604082015260600190565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610cfd57610cfd610cd6565b92915050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52603160045260245ffd5b634e487b7160e01b5f52601260045260245ffd5b5f82610d4d57610d4d610d2b565b500690565b5f60208284031215610d62575f5ffd5b5051919050565b5f82610d7757610d77610d2b565b500490565b5f60018201610d8d57610d8d610cd6565b506001019056fea2646970667358221220b46eeec28237596a4262b5b3bd90e42a4da1b317f90dd1e359472012a5a053c264736f6c634300081e0033",
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
	Bin: "0x60556032600b8282823980515f1a607314602657634e487b7160e01b5f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040525f5ffdfea2646970667358221220b26b9e1ecd34bff9421bcdd3cc77e0bd7170ae1887e46ec6235830cc01a80ac364736f6c634300081e0033",
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
	Bin: "0x6080604052348015600e575f5ffd5b5061313a8061001c5f395ff3fe60806040526004361061013b575f3560e01c80636969a25c116100b3578063a224cee71161006d578063a224cee71461036f578063a406fcb71461038e578063a43569b3146103ad578063afeea115146103dd578063b6c88519146103f1578063d6c0edad14610410575f5ffd5b80636969a25c146102a25780638a11d7c9146102c157806398e3b626146102f15780639de7025814610310578063a079886214610331578063a1ff465514610350575f5ffd5b806340550a1c1161010457806340550a1c146101f057806340a141ff1461020f5780634b3d500b146102305780635dd095901461024f5780636233be5d1461026e5780636846992a14610283575f5ffd5b8062362a771461013f5780631303f7cf14610173578063158ef93e146101965780631b5e358c146101ae5780633a061bd3146101db575b5f5ffd5b34801561014a575f5ffd5b5061015e610159366004612916565b610418565b60405190151581526020015b60405180910390f35b34801561017e575f5ffd5b5061018860045481565b60405190815260200161016a565b3480156101a1575f5ffd5b505f5461015e9060ff1681565b3480156101b9575f5ffd5b506101c361f00181565b6040516001600160a01b03909116815260200161016a565b3480156101e6575f5ffd5b506101c361f00081565b3480156101fb575f5ffd5b5061015e61020a366004612916565b61075d565b34801561021a575f5ffd5b5061022e610229366004612916565b6107b9565b005b34801561023b575f5ffd5b506101c361024a366004612931565b61080d565b34801561025a575f5ffd5b5061022e610269366004612916565b610835565b348015610279575f5ffd5b506101c361f00281565b34801561028e575f5ffd5b5061022e61029d36600461298c565b610886565b3480156102ad575f5ffd5b506101c36102bc366004612931565b610a54565b3480156102cc575f5ffd5b506102e06102db366004612916565b610a63565b60405161016a959493929190612a5a565b3480156102fc575f5ffd5b5061015e61030b366004612916565b610dfe565b34801561031b575f5ffd5b50610324610e52565b60405161016a9190612aa8565b34801561033c575f5ffd5b5061015e61034b366004612916565b610eb2565b34801561035b575f5ffd5b5061022e61036a366004612916565b6110b0565b34801561037a575f5ffd5b5061022e610389366004612af3565b6110fa565b348015610399575f5ffd5b5061015e6103a8366004612bb1565b6114b4565b3480156103b8575f5ffd5b506103cc6103c7366004612916565b611982565b60405161016a959493929190612cea565b3480156103e8575f5ffd5b50610324611d25565b3480156103fc575f5ffd5b5061015e61040b366004612dc1565b611d83565b61022e611f0e565b5f33816001600160a01b0384165f90815260016020526040902054600160a01b900460ff16600281111561044e5761044e612a46565b036104965760405162461bcd60e51b815260206004820152601360248201527215985b1a59185d1bdc881b9bdd08195e1a5cdd606a1b60448201526064015b60405180910390fd5b6001600160a01b038381165f908152600160205260409020548116908216146105185760405162461bcd60e51b815260206004820152602e60248201527f596f7520617265206e6f742074686520666565207265636569766572206f662060448201526d3a3434b9903b30b634b230ba37b960911b606482015260840161048d565b600554604080516394522b6d60e01b8152905143926001600160a01b0316916394522b6d9160048083019260209291908290030181865afa15801561055f573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105839190612e9f565b6001600160a01b0385165f908152600160205260409020600801546105a89190612eca565b11156106425760405162461bcd60e51b815260206004820152605c60248201527f596f75206d757374207761697420656e6f75676820626c6f636b7320746f207760448201527f6974686472617720796f75722070726f66697473206166746572206c6174657360648201527f74207769746864726177206f6620746869732076616c696461746f7200000000608482015260a40161048d565b6001600160a01b0383165f90815260016020526040902060060154806106aa5760405162461bcd60e51b815260206004820152601a60248201527f596f7520646f6e2774206861766520616e792070726f66697473000000000000604482015260640161048d565b6001600160a01b0384165f908152600160205260408120600681019190915543600890910155801561070b576040516001600160a01b0383169082156108fc029083905f818181858888f19350505050158015610709573d5f5f3e3d5ffd5b505b604080518281524260208201526001600160a01b0380851692908716917f51a69b4502f660774c9339825c7b5adbf0b8622289134647e29728ec5d9b3bb9910160405180910390a35060019392505050565b5f805b6002548110156107b157826001600160a01b03166002828154811061078757610787612edd565b5f918252602090912001546001600160a01b0316036107a95750600192915050565b600101610760565b505f92915050565b3361f001146108015760405162461bcd60e51b815260206004820152601460248201527350756e69736820636f6e7472616374206f6e6c7960601b604482015260640161048d565b61080a8161207f565b50565b6003818154811061081c575f80fd5b5f918252602090912001546001600160a01b0316905081565b3361f0011461087d5760405162461bcd60e51b815260206004820152601460248201527350756e69736820636f6e7472616374206f6e6c7960601b604482015260640161048d565b61080a8161217c565b3341146108c25760405162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015260640161048d565b435f9081526007602090815260408083206001845290915290205460ff161561092d5760405162461bcd60e51b815260206004820152601a60248201527f56616c696461746f727320616c72656164792075706461746564000000000000604482015260640161048d565b5f5460ff1661094e5760405162461bcd60e51b815260040161048d90612ef1565b806109598143612f2b565b156109995760405162461bcd60e51b815260206004820152601060248201526f426c6f636b2065706f6368206f6e6c7960801b604482015260640161048d565b435f90815260076020908152604080832060018085529252909120805460ff191690911790558251610a045760405162461bcd60e51b815260206004820152601460248201527356616c696461746f722073657420656d7074792160601b604482015260640161048d565b8251610a1790600290602086019061288b565b507feacea8f3c22f06c0b18306bdb04d0a967255129e8ce0094debb0a0ff89d006b583604051610a479190612aa8565b60405180910390a1505050565b6002818154811061081c575f80fd5b6001600160a01b038181165f908152600160209081526040808320815160c08101909252805494851682529293849384938493849384939192830190600160a01b900460ff166002811115610aba57610aba612a46565b6002811115610acb57610acb612a46565b8152602001600182016040518060a00160405290815f82018054610aee90612f3e565b80601f0160208091040260200160405190810160405280929190818152602001828054610b1a90612f3e565b8015610b655780601f10610b3c57610100808354040283529160200191610b65565b820191905f5260205f20905b815481529060010190602001808311610b4857829003601f168201915b50505050508152602001600182018054610b7e90612f3e565b80601f0160208091040260200160405190810160405280929190818152602001828054610baa90612f3e565b8015610bf55780601f10610bcc57610100808354040283529160200191610bf5565b820191905f5260205f20905b815481529060010190602001808311610bd857829003601f168201915b50505050508152602001600282018054610c0e90612f3e565b80601f0160208091040260200160405190810160405280929190818152602001828054610c3a90612f3e565b8015610c855780601f10610c5c57610100808354040283529160200191610c85565b820191905f5260205f20905b815481529060010190602001808311610c6857829003601f168201915b50505050508152602001600382018054610c9e90612f3e565b80601f0160208091040260200160405190810160405280929190818152602001828054610cca90612f3e565b8015610d155780601f10610cec57610100808354040283529160200191610d15565b820191905f5260205f20905b815481529060010190602001808311610cf857829003601f168201915b50505050508152602001600482018054610d2e90612f3e565b80601f0160208091040260200160405190810160405280929190818152602001828054610d5a90612f3e565b8015610da55780601f10610d7c57610100808354040283529160200191610da5565b820191905f5260205f20905b815481529060010190602001808311610d8857829003601f168201915b505050505081525050815260200160068201548152602001600782015481526020016008820154815250509050805f01518160200151826060015183608001518460a00151955095509550955095505091939590929450565b5f805b6003548110156107b157826001600160a01b031660038281548110610e2857610e28612edd565b5f918252602090912001546001600160a01b031603610e4a5750600192915050565b600101610e01565b60606002805480602002602001604051908101604052809291908181526020018280548015610ea857602002820191905f5260205f20905b81546001600160a01b03168152600190910190602001808311610e8a575b5050505050905090565b5f3361f00214610efd5760405162461bcd60e51b815260206004820152601660248201527550726f706f73616c20636f6e7472616374206f6e6c7960501b604482015260640161048d565b5f5460ff16610f1e5760405162461bcd60e51b815260040161048d90612ef1565b60016001600160a01b0383165f90815260016020526040902054600160a01b900460ff166002811115610f5357610f53612a46565b03610f6057506001919050565b610f698261228d565b60026001600160a01b0383165f90815260016020526040902054600160a01b900460ff166002811115610f9e57610f9e612a46565b0361104b576006546040516363e1d45160e01b81526001600160a01b038481166004830152909116906363e1d451906024016020604051808303815f875af1158015610fec573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110109190612f76565b61104b5760405162461bcd60e51b815260206004820152600c60248201526b18db19585b8819985a5b195960a21b604482015260640161048d565b6001600160a01b0382165f81815260016020908152604091829020805460ff60a01b1916600160a01b17905590514281527f8bef9a500ef702fa4b7c82318f7b750176b75d33c8897ad10a35e5e5e4161362910160405180910390a25060015b919050565b3361f002146108015760405162461bcd60e51b815260206004820152601660248201527550726f706f73616c20636f6e7472616374206f6e6c7960501b604482015260640161048d565b5f5460ff16156111425760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604482015260640161048d565b600580546001600160a01b031990811661f002179091556006805490911661f0011790555f5b818110156114a3575f83838381811061118357611183612edd565b90506020020160208101906111989190612916565b6001600160a01b0316036111ee5760405162461bcd60e51b815260206004820152601960248201527f496e76616c69642076616c696461746f72206164647265737300000000000000604482015260640161048d565b61121883838381811061120357611203612edd565b905060200201602081019061020a9190612916565b61127757600283838381811061123057611230612edd565b90506020020160208101906112459190612916565b81546001810183555f928352602090922090910180546001600160a01b0319166001600160a01b039092169190911790555b6112a183838381811061128c5761128c612edd565b905060200201602081019061030b9190612916565b6113005760038383838181106112b9576112b9612edd565b90506020020160208101906112ce9190612916565b81546001810183555f928352602090922090910180546001600160a01b0319166001600160a01b039092169190911790555b5f60018185858581811061131657611316612edd565b905060200201602081019061132b9190612916565b6001600160a01b03908116825260208201929092526040015f205416036113ce5782828281811061135e5761135e612edd565b90506020020160208101906113739190612916565b60015f85858581811061138857611388612edd565b905060200201602081019061139d9190612916565b6001600160a01b03908116825260208201929092526040015f2080546001600160a01b031916929091169190911790555b5f60015f8585858181106113e4576113e4612edd565b90506020020160208101906113f99190612916565b6001600160a01b0316815260208101919091526040015f2054600160a01b900460ff16600281111561142d5761142d612a46565b0361149b576001805f85858581811061144857611448612edd565b905060200201602081019061145d9190612916565b6001600160a01b0316815260208101919091526040015f20805460ff60a01b1916600160a01b83600281111561149557611495612a46565b02179055505b600101611168565b50505f805460ff1916600117905550565b5f805460ff166114d65760405162461bcd60e51b815260040161048d90612ef1565b6001600160a01b038c166115225760405162461bcd60e51b8152602060048201526013602482015272496e76616c696420666565206164647265737360681b604482015260640161048d565b61162c8b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f9201919091525050604080516020601f8f018190048102820181019092528d815292508d91508c90819084018382808284375f9201919091525050604080516020601f8e018190048102820181019092528c815292508c91508b90819084018382808284375f9201919091525050604080516020601f8d018190048102820181019092528b815292508b91508a90819084018382808284375f9201919091525050604080516020601f8c018190048102820181019092528a815292508a91508990819084018382808284375f92019190915250611d8392505050565b61166e5760405162461bcd60e51b815260206004820152601360248201527224b73b30b634b2103232b9b1b934b83a34b7b760691b604482015260640161048d565b60055460405163416259d960e11b81523360048201819052916001600160a01b0316906382c4b3b290602401602060405180830381865afa1580156116b5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116d99190612f76565b6117255760405162461bcd60e51b815260206004820152601c60248201527f596f75206d75737420626520617574686f72697a656420666972737400000000604482015260640161048d565b6001600160a01b038181165f908152600160205260409020548116908e1614611776576001600160a01b038181165f90815260016020526040902080546001600160a01b031916918f169190911790555b6040518060a001604052808d8d8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250505090825250604080516020601f8e018190048102820181019092528c815291810191908d908d90819084018382808284375f92019190915250505090825250604080516020601f8c018190048102820181019092528a815291810191908b908b90819084018382808284375f92019190915250505090825250604080516020601f8a01819004810282018101909252888152918101919089908990819084018382808284375f92019190915250505090825250604080516020601f8801819004810282018101909252868152918101919087908790819084018382808284375f9201829052509390945250506001600160a01b038416815260016020819052604090912083519101915081906118ca9082612fe1565b50602082015160018201906118df9082612fe1565b50604082015160028201906118f49082612fe1565b50606082015160038201906119099082612fe1565b506080820151600482019061191e9082612fe1565b509050508c6001600160a01b0316816001600160a01b03167fb8421f65501371f54d58de1937ff1e1ccdb76423ef6f84acea1814a0f6362ca04260405161196791815260200190565b60405180910390a35060019c9b505050505050505050505050565b6001600160a01b038181165f908152600160209081526040808320815160c081019092528054948516825260609485948594859485949293909291830190600160a01b900460ff1660028111156119db576119db612a46565b60028111156119ec576119ec612a46565b8152602001600182016040518060a00160405290815f82018054611a0f90612f3e565b80601f0160208091040260200160405190810160405280929190818152602001828054611a3b90612f3e565b8015611a865780601f10611a5d57610100808354040283529160200191611a86565b820191905f5260205f20905b815481529060010190602001808311611a6957829003601f168201915b50505050508152602001600182018054611a9f90612f3e565b80601f0160208091040260200160405190810160405280929190818152602001828054611acb90612f3e565b8015611b165780601f10611aed57610100808354040283529160200191611b16565b820191905f5260205f20905b815481529060010190602001808311611af957829003601f168201915b50505050508152602001600282018054611b2f90612f3e565b80601f0160208091040260200160405190810160405280929190818152602001828054611b5b90612f3e565b8015611ba65780601f10611b7d57610100808354040283529160200191611ba6565b820191905f5260205f20905b815481529060010190602001808311611b8957829003601f168201915b50505050508152602001600382018054611bbf90612f3e565b80601f0160208091040260200160405190810160405280929190818152602001828054611beb90612f3e565b8015611c365780601f10611c0d57610100808354040283529160200191611c36565b820191905f5260205f20905b815481529060010190602001808311611c1957829003601f168201915b50505050508152602001600482018054611c4f90612f3e565b80601f0160208091040260200160405190810160405280929190818152602001828054611c7b90612f3e565b8015611cc65780601f10611c9d57610100808354040283529160200191611cc6565b820191905f5260205f20905b815481529060010190602001808311611ca957829003601f168201915b50505091909252505050815260068201546020808301919091526007830154604080840191909152600890930154606092830152928201518051938101519281015191810151608090910151939b929a50909850965090945092505050565b60606003805480602002602001604051908101604052809291908181526020018280548015610ea857602002820191905f5260205f209081546001600160a01b03168152600190910190602001808311610e8a575050505050905090565b5f604686511115611dcf5760405162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840dadedcd2d6cae440d8cadccee8d60531b604482015260640161048d565b610bb885511115611e225760405162461bcd60e51b815260206004820152601760248201527f496e76616c6964206964656e74697479206c656e677468000000000000000000604482015260640161048d565b608c84511115611e6d5760405162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840eecac4e6d2e8ca40d8cadccee8d60531b604482015260640161048d565b608c83511115611eb65760405162461bcd60e51b8152602060048201526014602482015273092dcecc2d8d2c840cadac2d2d840d8cadccee8d60631b604482015260640161048d565b61011882511115611f025760405162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840c8cae8c2d2d8e640d8cadccee8d60531b604482015260640161048d565b50600195945050505050565b334114611f4a5760405162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015260640161048d565b435f90815260076020908152604080832083805290915290205460ff1615611fb45760405162461bcd60e51b815260206004820152601960248201527f426c6f636b20697320616c726561647920726577617264656400000000000000604482015260640161048d565b5f5460ff16611fd55760405162461bcd60e51b815260040161048d90612ef1565b435f9081526007602090815260408083208380528252808320805460ff1916600190811790915533808552925282205490913491600160a01b900460ff16600281111561202457612024612a46565b0361202d575050565b612037815f61235f565b604080518281524260208201526001600160a01b038416917f7dc4e5df59513708dca355b8706273a5df7b810a4cec8019f2a4b9bb166a1a0491015b60405180910390a25050565b6001600160a01b0381165f9081526001602052604090206006810154815460ff60a01b1916600160a11b179091556120b68261217c565b60035460011015612178576120ca826124ea565b6005546040516315ea278160e01b81526001600160a01b038481166004830152909116906315ea2781906024016020604051808303815f875af1158015612113573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906121379190612f76565b50604080518281524260208201526001600160a01b038416917fa26de7ab324eac08c596549f421e5c8741213d237d2e9a2c9c0ebde0a7a849fe9101612073565b5050565b5f6001600160a01b0382165f90815260016020526040902054600160a01b900460ff1660028111156121b0576121b0612a46565b14806121bf5750600254600110155b156121c75750565b6001600160a01b0381165f90815260016020526040902060060154801561224d576121f2818361235f565b6004546121ff9082612641565b6004556001600160a01b0382165f908152600160205260409020600701546122279082612641565b6001600160a01b0383165f90815260016020526040812060078101929092556006909101555b604080518281524260208201526001600160a01b038416917fe294e9d73f8eee23e21b2e1567960625a6b5d339cb127b55d0d09473a99512359101612073565b5f5b6003548110156122db57816001600160a01b0316600382815481106122b6576122b6612edd565b5f918252602090912001546001600160a01b0316036122d3575050565b60010161228f565b50600380546001810182555f919091527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0180546001600160a01b0319166001600160a01b0383169081179091556040514281527f1e3310ad6891b30e03874ec3d1422a6386c5da63d9faf595f5d99eeaf443b99a9060200160405180910390a250565b815f0361236a575050565b5f612374826126a8565b9050805f0361238257505050565b5f808061238f8685612758565b90506123a561239e82866127b2565b8790612830565b92505f5b600254811015612484575f600282815481106123c7576123c7612edd565b5f918252602090912001546001600160a01b0316905060026001600160a01b0382165f90815260016020526040902054600160a01b900460ff16600281111561241257612412612a46565b141580156124325750866001600160a01b0316816001600160a01b031614155b1561247b576001600160a01b0381165f9081526001602052604090206006015461245c9084612641565b6001600160a01b0382165f908152600160205260409020600601559250825b506001016123a9565b505f8311801561249c57506001600160a01b03821615155b156124e2576001600160a01b0382165f908152600160205260409020600601546124c69084612641565b6001600160a01b0383165f908152600160205260409020600601555b505050505050565b5f5b600354811080156124ff57506003546001105b15612178576003818154811061251757612517612edd565b5f918252602090912001546001600160a01b039081169083160361262f576003546125449060019061309b565b81146125c1576003805461255a9060019061309b565b8154811061256a5761256a612edd565b5f91825260209091200154600380546001600160a01b03909216918390811061259557612595612edd565b905f5260205f20015f6101000a8154816001600160a01b0302191690836001600160a01b031602179055505b60038054806125d2576125d26130ae565b5f8281526020902081015f1990810180546001600160a01b03191690550190556040516001600160a01b038316907f7521e44559c870c316e84e60bc4785d9c034a8ab1d6acdce8134ac03f946c6ed906120739042815260200190565b80612639816130c2565b9150506124ec565b5f8061264d8385612eca565b90508381101561269f5760405162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015260640161048d565b90505b92915050565b5f80805b600254811015612751575f600282815481106126ca576126ca612edd565b5f918252602090912001546001600160a01b0316905060026001600160a01b0382165f90815260016020526040902054600160a01b900460ff16600281111561271557612715612a46565b141580156127355750846001600160a01b0316816001600160a01b031614155b156127485782612744816130c2565b9350505b506001016126ac565b5092915050565b5f5f82116127a85760405162461bcd60e51b815260206004820152601a60248201527f536166654d6174683a206469766973696f6e206279207a65726f000000000000604482015260640161048d565b61269f82846130da565b5f825f036127c157505f6126a2565b5f6127cc83856130ed565b9050826127d985836130da565b1461269f5760405162461bcd60e51b815260206004820152602160248201527f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f6044820152607760f81b606482015260840161048d565b5f828211156128815760405162461bcd60e51b815260206004820152601e60248201527f536166654d6174683a207375627472616374696f6e206f766572666c6f770000604482015260640161048d565b61269f828461309b565b828054828255905f5260205f209081019282156128de579160200282015b828111156128de57825182546001600160a01b0319166001600160a01b039091161782556020909201916001909101906128a9565b506128ea9291506128ee565b5090565b5b808211156128ea575f81556001016128ef565b6001600160a01b038116811461080a575f5ffd5b5f60208284031215612926575f5ffd5b813561269f81612902565b5f60208284031215612941575f5ffd5b5035919050565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b038111828210171561298457612984612948565b604052919050565b5f5f6040838503121561299d575f5ffd5b82356001600160401b038111156129b2575f5ffd5b8301601f810185136129c2575f5ffd5b80356001600160401b038111156129db576129db612948565b8060051b6129eb6020820161295c565b91825260208184018101929081019088841115612a06575f5ffd5b6020850194505b83851015612a345784359250612a2283612902565b82825260209485019490910190612a0d565b98602097909701359750505050505050565b634e487b7160e01b5f52602160045260245ffd5b6001600160a01b038616815260a0810160038610612a8657634e487b7160e01b5f52602160045260245ffd5b8560208301528460408301528360608301528260808301529695505050505050565b602080825282518282018190525f918401906040840190835b81811015612ae85783516001600160a01b0316835260209384019390920191600101612ac1565b509095945050505050565b5f5f60208385031215612b04575f5ffd5b82356001600160401b03811115612b19575f5ffd5b8301601f81018513612b29575f5ffd5b80356001600160401b03811115612b3e575f5ffd5b8560208260051b8401011115612b52575f5ffd5b6020919091019590945092505050565b80356110ab81612902565b5f5f83601f840112612b7d575f5ffd5b5081356001600160401b03811115612b93575f5ffd5b602083019150836020828501011115612baa575f5ffd5b9250929050565b5f5f5f5f5f5f5f5f5f5f5f60c08c8e031215612bcb575f5ffd5b612bd48c612b62565b9a5060208c01356001600160401b03811115612bee575f5ffd5b612bfa8e828f01612b6d565b909b5099505060408c01356001600160401b03811115612c18575f5ffd5b612c248e828f01612b6d565b90995097505060608c01356001600160401b03811115612c42575f5ffd5b612c4e8e828f01612b6d565b90975095505060808c01356001600160401b03811115612c6c575f5ffd5b612c788e828f01612b6d565b90955093505060a08c01356001600160401b03811115612c96575f5ffd5b612ca28e828f01612b6d565b915080935050809150509295989b509295989b9093969950565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b60a081525f612cfc60a0830188612cbc565b8281036020840152612d0e8188612cbc565b90508281036040840152612d228187612cbc565b90508281036060840152612d368186612cbc565b90508281036080840152612d4a8185612cbc565b98975050505050505050565b5f82601f830112612d65575f5ffd5b81356001600160401b03811115612d7e57612d7e612948565b612d91601f8201601f191660200161295c565b818152846020838601011115612da5575f5ffd5b816020850160208301375f918101602001919091529392505050565b5f5f5f5f5f60a08688031215612dd5575f5ffd5b85356001600160401b03811115612dea575f5ffd5b612df688828901612d56565b95505060208601356001600160401b03811115612e11575f5ffd5b612e1d88828901612d56565b94505060408601356001600160401b03811115612e38575f5ffd5b612e4488828901612d56565b93505060608601356001600160401b03811115612e5f575f5ffd5b612e6b88828901612d56565b92505060808601356001600160401b03811115612e86575f5ffd5b612e9288828901612d56565b9150509295509295909350565b5f60208284031215612eaf575f5ffd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b808201808211156126a2576126a2612eb6565b634e487b7160e01b5f52603260045260245ffd5b6020808252600c908201526b139bdd081a5b9a5d081e595d60a21b604082015260600190565b634e487b7160e01b5f52601260045260245ffd5b5f82612f3957612f39612f17565b500690565b600181811c90821680612f5257607f821691505b602082108103612f7057634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215612f86575f5ffd5b8151801515811461269f575f5ffd5b601f821115612fdc57805f5260205f20601f840160051c81016020851015612fba5750805b601f840160051c820191505b81811015612fd9575f8155600101612fc6565b50505b505050565b81516001600160401b03811115612ffa57612ffa612948565b61300e816130088454612f3e565b84612f95565b6020601f821160018114613040575f83156130295750848201515b5f19600385901b1c1916600184901b178455612fd9565b5f84815260208120601f198516915b8281101561306f578785015182556020948501946001909201910161304f565b508482101561308c57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b818103818111156126a2576126a2612eb6565b634e487b7160e01b5f52603160045260245ffd5b5f600182016130d3576130d3612eb6565b5060010190565b5f826130e8576130e8612f17565b500490565b80820281158282048414176126a2576126a2612eb656fea264697066735822122060bf2f8db10aff54d5f83e86162e433e7cdba64b7627c38ccf9bf4d1e90d140164736f6c634300081e0033",
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
