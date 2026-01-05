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

// ProposalMetaData contains all meta data concerning the Proposal contract.
var ProposalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"PROPOSAL_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUNISH_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blockReward\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createProposal\",\"inputs\":[{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createUpdateConfigProposal\",\"inputs\":[{\"name\":\"cid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreaseRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"vals\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"validators_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isProposalValidForStaking\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minDelegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minUndelegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minValidatorStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pass\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalLastingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalPassedHeight\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposals\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"createBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposalType\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"cid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposerNonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"punishThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"results\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"agree\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"reject\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"resultExist\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUnpassed\",\"inputs\":[{\"name\":\"val\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbondingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorUnjailPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voteProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"auth\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"votes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"voteTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"auth\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawProfitPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"LogCreateConfigProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogCreateProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dst\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogPassProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogRejectProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogSetUnpassed\",\"inputs\":[{\"name\":\"val\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogVote\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"auth\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055611d86806100405f395ff3fe608060405234801561000f575f5ffd5b50600436106101c6575f3560e01c80635d4f0cb6116100fe5780639f759dba1161009e578063d24806eb1161006e578063d24806eb14610430578063d51cade814610443578063e823c81414610456578063f945b6231461045f575f5ffd5b80639f759dba14610402578063a4c4d9221461040b578063b2b2732a1461041e578063cb1ea72514610427575f5ffd5b806382c4b3b2116100d957806382c4b3b2146103af5780638c872d05146103d157806394522b6d146103da5780639e4353cc146103e3575f5ffd5b80635d4f0cb61461037e5780636cf6d675146103875780637668dd2414610390575f5ffd5b80631db5ade811610169578063437ccda811610144578063437ccda8146102e157806344c1aa9914610302578063462d0b2e1461030b5780634c6b25b114610320575f5ffd5b80631db5ade8146102435780632897183d146102b057806332ed5b12146102b9575f5ffd5b80630ac168a1116101a45780630ac168a1146101f8578063107a03a814610201578063158ef93e1461022457806315ea278114610230575f5ffd5b8063017ddd35146101ca57806302985992146101e657806308ac5256146101ef575b5f5ffd5b6101d360095481565b6040519081526020015b60405180910390f35b6101d3600b5481565b6101d3600a5481565b6101d360065481565b61021461020f366004611898565b610468565b60405190151581526020016101dd565b5f546102149060ff1681565b61021461023e366004611898565b6104bd565b6102896102513660046118b8565b601260209081525f92835260408084209091529082529020805460018201546002909201546001600160a01b03909116919060ff1683565b604080516001600160a01b03909416845260208401929092521515908201526060016101dd565b6101d360045481565b6102cc6102c73660046118e0565b61052c565b6040516101dd999897969594939291906118f7565b6102ea61f01281565b6040516001600160a01b0390911681526020016101dd565b6101d360035481565b61031e61031936600461197c565b61060e565b005b61035b61032e3660046118e0565b60116020525f908152604090205461ffff8082169162010000810490911690640100000000900460ff1683565b6040805161ffff94851681529390921660208401521515908201526060016101dd565b6102ea61f01381565b6101d360075481565b6101d361039e366004611898565b600e6020525f908152604090205481565b6102146103bd366004611898565b600d6020525f908152604090205460ff1681565b6102ea61f01181565b6101d360055481565b6101d36103f1366004611898565b600f6020525f908152604090205481565b6102ea61f01081565b610214610419366004611a0b565b610837565b6101d3600c5481565b6101d360025481565b6101d361043e366004611a39565b610e71565b6101d3610451366004611a59565b611078565b6101d360015481565b6101d360085481565b6001600160a01b0381165f908152600d602052604081205460ff1661048e57505f919050565b6001600160a01b0382165f908152600e60205260409020546001546104b39082611af9565b4311159392505050565b5f6104c6611544565b6001600160a01b0382165f818152600d60209081526040808320805460ff19169055600e82528083209290925590514281527f4e0b191f7f5c32b1b5e3704b68874b1a3980147cae00be8ece271bfb5b92c07a910160405180910390a25060015b919050565b60106020525f90815260409020805460018201546002830154600384015460048501546005860180546001600160a01b0396871697959694959394831693600160a01b90930460ff1692919061058190611b0c565b80601f01602080910402602001604051908101604052809291908181526020018280546105ad90611b0c565b80156105f85780601f106105cf576101008083540402835291602001916105f8565b820191905f5260205f20905b8154815290600101906020018083116105db57829003601f168201915b5050505050908060060154908060070154905089565b610616611597565b6001600160a01b0381166106715760405162461bcd60e51b815260206004820152601a60248201527f496e76616c69642076616c696461746f7273206164647265737300000000000060448201526064015b60405180910390fd5b601380546001600160a01b0319166001600160a01b0383161790555f5b828110156107c1575f8484838181106106a9576106a9611b44565b90506020020160208101906106be9190611898565b6001600160a01b0316036107145760405162461bcd60e51b815260206004820152601960248201527f496e76616c69642076616c696461746f722061646472657373000000000000006044820152606401610668565b6001600d5f86868581811061072b5761072b611b44565b90506020020160208101906107409190611898565b6001600160a01b03166001600160a01b031681526020019081526020015f205f6101000a81548160ff02191690831515021790555043600e5f86868581811061078b5761078b611b44565b90506020020160208101906107a09190611898565b6001600160a01b0316815260208101919091526040015f205560010161068e565b505062093a8060018181556018600281905560306003556004556201518060058190556702c68af0bb14000060065560079290925560089190915569152d02c7e14af68000006009556015600a55678ac7230489e80000600b55670de0b6b3a7640000600c555f805460ff191690911790555050565b5f6108406115df565b610848611686565b5f8381526010602052604081206001015490036108a75760405162461bcd60e51b815260206004820152601760248201527f50726f706f73616c20646f6573206e6f742065786973740000000000000000006044820152606401610668565b335f9081526012602090815260408083208684529091529020600101541561091d5760405162461bcd60e51b815260206004820152602360248201527f596f752063616e277420766f746520666f7220612070726f706f73616c20747760448201526269636560e81b6064820152608401610668565b6001545f8481526010602052604090206002015461093b9190611af9565b431061097c5760405162461bcd60e51b815260206004820152601060248201526f141c9bdc1bdcd85b08195e1c1a5c995960821b6044820152606401610668565b335f818152601260209081526040808320878452825291829020426001820181905581546001600160a01b031916851782556002909101805460ff191687151590811790915583519081529182015285917f6c59bda68cac318717c60c7c9635a78a0f0613f9887cc18a7157f5745a86d14e910160405180910390a38115610a41575f83815260116020526040902054610a1b9061ffff166001611b58565b5f848152601160205260409020805461ffff191661ffff92909216919091179055610a8f565b5f83815260116020526040902054610a649062010000900461ffff166001611b58565b5f848152601160205260409020805461ffff92909216620100000263ffff0000199092169190911790555b5f83815260116020526040902054640100000000900460ff1615610ab557506001610e42565b6013546040805163037deea760e41b815290516002926001600160a01b0316916337deea709160048083019260209291908290030181865afa158015610afd573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b219190611b72565b610b2b9190611b89565b610b36906001611af9565b5f8481526011602052604090205461ffff1610610d43575f838152601160209081526040808320805464ff00000000191664010000000017905560109091529020600301546001819003610c90575f84815260106020526040902060040154600160a01b900460ff1615610bec575f84815260106020908152604080832060040180546001600160a01b039081168552600d8452828520805460ff191660011790559054168352600e9091529020439055610cff565b5f848152601060208181526040808420600490810180546001600160a01b039081168752600d8552838720805460ff19169055815481168752600e8552838720879055601354968b9052949093529154905163a1ff465560e01b81529083169181019190915291169063a1ff4655906024015f604051808303815f87803b158015610c75575f5ffd5b505af1158015610c87573d5f5f3e3d5ffd5b50505050610cff565b80600203610cbf575f8481526010602052604090206006810154600790910154610cba91906116b4565b610cff565b60405162461bcd60e51b8152602060048201526015602482015274496e76616c69642070726f706f73616c207479706560581b6044820152606401610668565b837f90d2e923947d9356c1c04391cb9e2e9c5d4ad6c165a849787b0c7569bbe99e2442604051610d3191815260200190565b60405180910390a26001915050610e42565b6013546040805163037deea760e41b815290516002926001600160a01b0316916337deea709160048083019260209291908290030181865afa158015610d8b573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610daf9190611b72565b610db99190611b89565b610dc4906001611af9565b5f8481526011602052604090205462010000900461ffff1610610e3e575f8381526011602052604090819020805464ff0000000019166401000000001790555183907f36bdb56d707cdf53eadffe319a71ddf97736be67b8caab47b7720201a6b65ca090610e359042815260200190565b60405180910390a25b5060015b610e6b60017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b92915050565b5f610e7a6115df565b610e8483836117a2565b610ed05760405162461bcd60e51b815260206004820152601860248201527f436f6e6669672076616c69646174696f6e206661696c656400000000000000006044820152606401610668565b335f818152600f60208181526040808420805482518085018890528084018b9052606081018a905260808082018390528451808303909101815260a0909101909352825192840192909220958552929091529291839190610f3083611ba8565b909155505060408051610120810182525f8082526020808301828152838501838152606085018481526080860185815260a0870186815288518087018a5287815260c0890190815233895260e089018f905261010089018e905242865243855260028085528b8952601090975298909620875181546001600160a01b039182166001600160a01b031990911617825594516001820155925194830194909455516003820155915160048301805494511515600160a01b026001600160a81b031990951691909216179290921790915591519091829160058201906110149082611c20565b5060e0820151600682015561010090910151600790910155604080518781526020810187905242818301529051339184917f8bfc061277ae1778974ada10db7f9664ab1d67c455c025c025b438c52c69d1819181900360600190a350949350505050565b5f6110816115df565b831561121c57601354604051634c71db1360e11b81526001600160a01b0387811660048301525f9216906398e3b62690602401602060405180830381865afa1580156110cf573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110f39190611cdb565b905080156111555760405162461bcd60e51b815260206004820152602960248201527f56616c696461746f7220697320616c726561647920696e20746f702076616c6960448201526819185d1bdc881cd95d60ba1b6064820152608401610668565b6001600160a01b0386165f908152600d602052604090205460ff161561121a576001600160a01b0386165f908152600e602052604090205460015461119a9082611af9565b4311156111d0576001600160a01b0387165f908152600d60209081526040808320805460ff19169055600e909152812055611218565b60405162461bcd60e51b815260206004820152601f60248201527f43616e27742061646420616e20616c72656164792070617373656420647374006044820152606401610668565b505b505b6001600160a01b0385165f908152600d602052604090205460ff161580156112415750835b8061124a575083155b6112965760405162461bcd60e51b815260206004820152601e60248201527f43616e27742061646420616e20616c72656164792065786973742064737400006044820152606401610668565b335f818152600f6020908152604080832054905190936112c29290918a918a918a918a91899101611cf6565b60408051601f1981840301815291905280516020909101209050610bb88411156113215760405162461bcd60e51b815260206004820152601060248201526f44657461696c7320746f6f206c6f6e6760801b6044820152606401610668565b5f818152601060205260409020600101541561137f5760405162461bcd60e51b815260206004820152601760248201527f50726f706f73616c20616c7265616479206578697374730000000000000000006044820152606401610668565b335f908152600f6020526040812080549161139983611ba8565b909155505060408051610120810182525f8082526020808301829052828401829052606083018290526080830182815260a084018381528551808401875284815260c086015260e085018490526101008501939093523384526001600160a01b038c1690528915159091528251601f880182900482028101820190935286835290919087908790819084018382808284375f92018290525060c08601948552426020808801918252436040808a01918252600160608b018181528c875260109094529420895181546001600160a01b0319166001600160a01b03918216178255935194810194909455516002840155516003830155608087015160048301805460a08a0151929093166001600160a81b031990931692909217600160a01b911515919091021790559351859493506005840192506114d8915082611c20565b5060e08201516006820155610100909101516007909101556040805188151581524260208201526001600160a01b038a1691339185917f1af05d46b8c1ec021d82b7128cff40e91a1c2337deffc010df48eeddef8da56c910160405180910390a4509695505050505050565b3361f010146115955760405162461bcd60e51b815260206004820152601860248201527f56616c696461746f727320636f6e7472616374206f6e6c7900000000000000006044820152606401610668565b565b5f5460ff16156115955760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b6044820152606401610668565b601354604051631015428760e21b81523360048201526001600160a01b03909116906340550a1c90602401602060405180830381865afa158015611625573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116499190611cdb565b6115955760405162461bcd60e51b815260206004820152600e60248201526d56616c696461746f72206f6e6c7960901b6044820152606401610668565b61168e611840565b60027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b6116be82826117a2565b50815f036116cc5760015550565b816001036116da5760025550565b816002036116e85760035550565b816003036116f65760045550565b816004036117045760055550565b816005036117125760065550565b816006036117205760075550565b8160070361172e5760085550565b8160080361173c5760095550565b8160090361174a57600a5550565b81600a0361175857600b5550565b81600b0361176657600c5550565b60405162461bcd60e51b8152602060048201526011602482015270155b9adb9bdddb8818dbdb999a59c81251607a1b6044820152606401610668565b5f600b8311156117e85760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a590818dbdb999a59c81251607a1b6044820152606401610668565b5f82116118375760405162461bcd60e51b815260206004820152601d60248201527f436f6e6669672076616c7565206d75737420626520706f7369746976650000006044820152606401610668565b50600192915050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f005460020361159557604051633ee5aeb560e01b815260040160405180910390fd5b80356001600160a01b0381168114610527575f5ffd5b5f602082840312156118a8575f5ffd5b6118b182611882565b9392505050565b5f5f604083850312156118c9575f5ffd5b6118d283611882565b946020939093013593505050565b5f602082840312156118f0575f5ffd5b5035919050565b60018060a01b038a16815288602082015287604082015286606082015260018060a01b038616608082015284151560a082015261012060c08201525f845180610120840152806020870161014085015e5f6101408285010152610140601f19601f8301168401019150508360e0830152826101008301529a9950505050505050505050565b5f5f5f6040848603121561198e575f5ffd5b833567ffffffffffffffff8111156119a4575f5ffd5b8401601f810186136119b4575f5ffd5b803567ffffffffffffffff8111156119ca575f5ffd5b8660208260051b84010111156119de575f5ffd5b6020918201945092506119f2908501611882565b90509250925092565b8015158114611a08575f5ffd5b50565b5f5f60408385031215611a1c575f5ffd5b823591506020830135611a2e816119fb565b809150509250929050565b5f5f60408385031215611a4a575f5ffd5b50508035926020909101359150565b5f5f5f5f60608587031215611a6c575f5ffd5b611a7585611882565b93506020850135611a85816119fb565b9250604085013567ffffffffffffffff811115611aa0575f5ffd5b8501601f81018713611ab0575f5ffd5b803567ffffffffffffffff811115611ac6575f5ffd5b876020828401011115611ad7575f5ffd5b949793965060200194505050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610e6b57610e6b611ae5565b600181811c90821680611b2057607f821691505b602082108103611b3e57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b61ffff8181168382160190811115610e6b57610e6b611ae5565b5f60208284031215611b82575f5ffd5b5051919050565b5f82611ba357634e487b7160e01b5f52601260045260245ffd5b500490565b5f60018201611bb957611bb9611ae5565b5060010190565b634e487b7160e01b5f52604160045260245ffd5b601f821115611c1b57805f5260205f20601f840160051c81016020851015611bf95750805b601f840160051c820191505b81811015611c18575f8155600101611c05565b50505b505050565b815167ffffffffffffffff811115611c3a57611c3a611bc0565b611c4e81611c488454611b0c565b84611bd4565b6020601f821160018114611c80575f8315611c695750848201515b5f19600385901b1c1916600184901b178455611c18565b5f84815260208120601f198516915b82811015611caf5787850151825560209485019460019092019101611c8f565b5084821015611ccc57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f60208284031215611ceb575f5ffd5b81516118b1816119fb565b6001600160a01b03878116825286166020820152841515604082015260a0606082018190528101839052828460c08301375f60c084830101525f60c0601f19601f860116830101905082608083015297965050505050505056fea264697066735822122093114184b66e2cad9f520c534061d5f404543b592c5a69683a8dd9392d35cd3e64736f6c634300081d0033",
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

// PROPOSALADDR is a free data retrieval call binding the contract method 0x437ccda8.
//
// Solidity: function PROPOSAL_ADDR() view returns(address)
func (_Proposal *ProposalCaller) PROPOSALADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "PROPOSAL_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PROPOSALADDR is a free data retrieval call binding the contract method 0x437ccda8.
//
// Solidity: function PROPOSAL_ADDR() view returns(address)
func (_Proposal *ProposalSession) PROPOSALADDR() (common.Address, error) {
	return _Proposal.Contract.PROPOSALADDR(&_Proposal.CallOpts)
}

// PROPOSALADDR is a free data retrieval call binding the contract method 0x437ccda8.
//
// Solidity: function PROPOSAL_ADDR() view returns(address)
func (_Proposal *ProposalCallerSession) PROPOSALADDR() (common.Address, error) {
	return _Proposal.Contract.PROPOSALADDR(&_Proposal.CallOpts)
}

// PUNISHADDR is a free data retrieval call binding the contract method 0x8c872d05.
//
// Solidity: function PUNISH_ADDR() view returns(address)
func (_Proposal *ProposalCaller) PUNISHADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "PUNISH_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PUNISHADDR is a free data retrieval call binding the contract method 0x8c872d05.
//
// Solidity: function PUNISH_ADDR() view returns(address)
func (_Proposal *ProposalSession) PUNISHADDR() (common.Address, error) {
	return _Proposal.Contract.PUNISHADDR(&_Proposal.CallOpts)
}

// PUNISHADDR is a free data retrieval call binding the contract method 0x8c872d05.
//
// Solidity: function PUNISH_ADDR() view returns(address)
func (_Proposal *ProposalCallerSession) PUNISHADDR() (common.Address, error) {
	return _Proposal.Contract.PUNISHADDR(&_Proposal.CallOpts)
}

// STAKINGADDR is a free data retrieval call binding the contract method 0x5d4f0cb6.
//
// Solidity: function STAKING_ADDR() view returns(address)
func (_Proposal *ProposalCaller) STAKINGADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "STAKING_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKINGADDR is a free data retrieval call binding the contract method 0x5d4f0cb6.
//
// Solidity: function STAKING_ADDR() view returns(address)
func (_Proposal *ProposalSession) STAKINGADDR() (common.Address, error) {
	return _Proposal.Contract.STAKINGADDR(&_Proposal.CallOpts)
}

// STAKINGADDR is a free data retrieval call binding the contract method 0x5d4f0cb6.
//
// Solidity: function STAKING_ADDR() view returns(address)
func (_Proposal *ProposalCallerSession) STAKINGADDR() (common.Address, error) {
	return _Proposal.Contract.STAKINGADDR(&_Proposal.CallOpts)
}

// VALIDATORADDR is a free data retrieval call binding the contract method 0x9f759dba.
//
// Solidity: function VALIDATOR_ADDR() view returns(address)
func (_Proposal *ProposalCaller) VALIDATORADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "VALIDATOR_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORADDR is a free data retrieval call binding the contract method 0x9f759dba.
//
// Solidity: function VALIDATOR_ADDR() view returns(address)
func (_Proposal *ProposalSession) VALIDATORADDR() (common.Address, error) {
	return _Proposal.Contract.VALIDATORADDR(&_Proposal.CallOpts)
}

// VALIDATORADDR is a free data retrieval call binding the contract method 0x9f759dba.
//
// Solidity: function VALIDATOR_ADDR() view returns(address)
func (_Proposal *ProposalCallerSession) VALIDATORADDR() (common.Address, error) {
	return _Proposal.Contract.VALIDATORADDR(&_Proposal.CallOpts)
}

// BlockReward is a free data retrieval call binding the contract method 0x0ac168a1.
//
// Solidity: function blockReward() view returns(uint256)
func (_Proposal *ProposalCaller) BlockReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "blockReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockReward is a free data retrieval call binding the contract method 0x0ac168a1.
//
// Solidity: function blockReward() view returns(uint256)
func (_Proposal *ProposalSession) BlockReward() (*big.Int, error) {
	return _Proposal.Contract.BlockReward(&_Proposal.CallOpts)
}

// BlockReward is a free data retrieval call binding the contract method 0x0ac168a1.
//
// Solidity: function blockReward() view returns(uint256)
func (_Proposal *ProposalCallerSession) BlockReward() (*big.Int, error) {
	return _Proposal.Contract.BlockReward(&_Proposal.CallOpts)
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

// IsProposalValidForStaking is a free data retrieval call binding the contract method 0x107a03a8.
//
// Solidity: function isProposalValidForStaking(address validator) view returns(bool)
func (_Proposal *ProposalCaller) IsProposalValidForStaking(opts *bind.CallOpts, validator common.Address) (bool, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "isProposalValidForStaking", validator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsProposalValidForStaking is a free data retrieval call binding the contract method 0x107a03a8.
//
// Solidity: function isProposalValidForStaking(address validator) view returns(bool)
func (_Proposal *ProposalSession) IsProposalValidForStaking(validator common.Address) (bool, error) {
	return _Proposal.Contract.IsProposalValidForStaking(&_Proposal.CallOpts, validator)
}

// IsProposalValidForStaking is a free data retrieval call binding the contract method 0x107a03a8.
//
// Solidity: function isProposalValidForStaking(address validator) view returns(bool)
func (_Proposal *ProposalCallerSession) IsProposalValidForStaking(validator common.Address) (bool, error) {
	return _Proposal.Contract.IsProposalValidForStaking(&_Proposal.CallOpts, validator)
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() view returns(uint256)
func (_Proposal *ProposalCaller) MaxValidators(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "maxValidators")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() view returns(uint256)
func (_Proposal *ProposalSession) MaxValidators() (*big.Int, error) {
	return _Proposal.Contract.MaxValidators(&_Proposal.CallOpts)
}

// MaxValidators is a free data retrieval call binding the contract method 0x08ac5256.
//
// Solidity: function maxValidators() view returns(uint256)
func (_Proposal *ProposalCallerSession) MaxValidators() (*big.Int, error) {
	return _Proposal.Contract.MaxValidators(&_Proposal.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() view returns(uint256)
func (_Proposal *ProposalCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "minDelegation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() view returns(uint256)
func (_Proposal *ProposalSession) MinDelegation() (*big.Int, error) {
	return _Proposal.Contract.MinDelegation(&_Proposal.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x02985992.
//
// Solidity: function minDelegation() view returns(uint256)
func (_Proposal *ProposalCallerSession) MinDelegation() (*big.Int, error) {
	return _Proposal.Contract.MinDelegation(&_Proposal.CallOpts)
}

// MinUndelegation is a free data retrieval call binding the contract method 0xb2b2732a.
//
// Solidity: function minUndelegation() view returns(uint256)
func (_Proposal *ProposalCaller) MinUndelegation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "minUndelegation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinUndelegation is a free data retrieval call binding the contract method 0xb2b2732a.
//
// Solidity: function minUndelegation() view returns(uint256)
func (_Proposal *ProposalSession) MinUndelegation() (*big.Int, error) {
	return _Proposal.Contract.MinUndelegation(&_Proposal.CallOpts)
}

// MinUndelegation is a free data retrieval call binding the contract method 0xb2b2732a.
//
// Solidity: function minUndelegation() view returns(uint256)
func (_Proposal *ProposalCallerSession) MinUndelegation() (*big.Int, error) {
	return _Proposal.Contract.MinUndelegation(&_Proposal.CallOpts)
}

// MinValidatorStake is a free data retrieval call binding the contract method 0x017ddd35.
//
// Solidity: function minValidatorStake() view returns(uint256)
func (_Proposal *ProposalCaller) MinValidatorStake(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "minValidatorStake")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinValidatorStake is a free data retrieval call binding the contract method 0x017ddd35.
//
// Solidity: function minValidatorStake() view returns(uint256)
func (_Proposal *ProposalSession) MinValidatorStake() (*big.Int, error) {
	return _Proposal.Contract.MinValidatorStake(&_Proposal.CallOpts)
}

// MinValidatorStake is a free data retrieval call binding the contract method 0x017ddd35.
//
// Solidity: function minValidatorStake() view returns(uint256)
func (_Proposal *ProposalCallerSession) MinValidatorStake() (*big.Int, error) {
	return _Proposal.Contract.MinValidatorStake(&_Proposal.CallOpts)
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

// ProposalPassedHeight is a free data retrieval call binding the contract method 0x7668dd24.
//
// Solidity: function proposalPassedHeight(address ) view returns(uint256)
func (_Proposal *ProposalCaller) ProposalPassedHeight(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "proposalPassedHeight", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalPassedHeight is a free data retrieval call binding the contract method 0x7668dd24.
//
// Solidity: function proposalPassedHeight(address ) view returns(uint256)
func (_Proposal *ProposalSession) ProposalPassedHeight(arg0 common.Address) (*big.Int, error) {
	return _Proposal.Contract.ProposalPassedHeight(&_Proposal.CallOpts, arg0)
}

// ProposalPassedHeight is a free data retrieval call binding the contract method 0x7668dd24.
//
// Solidity: function proposalPassedHeight(address ) view returns(uint256)
func (_Proposal *ProposalCallerSession) ProposalPassedHeight(arg0 common.Address) (*big.Int, error) {
	return _Proposal.Contract.ProposalPassedHeight(&_Proposal.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(address proposer, uint256 createTime, uint256 createBlock, uint256 proposalType, address dst, bool flag, string details, uint256 cid, uint256 newValue)
func (_Proposal *ProposalCaller) Proposals(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Proposer     common.Address
	CreateTime   *big.Int
	CreateBlock  *big.Int
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
		CreateBlock  *big.Int
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
	outstruct.CreateBlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ProposalType = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Dst = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Flag = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.Details = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.Cid = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.NewValue = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x32ed5b12.
//
// Solidity: function proposals(bytes32 ) view returns(address proposer, uint256 createTime, uint256 createBlock, uint256 proposalType, address dst, bool flag, string details, uint256 cid, uint256 newValue)
func (_Proposal *ProposalSession) Proposals(arg0 [32]byte) (struct {
	Proposer     common.Address
	CreateTime   *big.Int
	CreateBlock  *big.Int
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
// Solidity: function proposals(bytes32 ) view returns(address proposer, uint256 createTime, uint256 createBlock, uint256 proposalType, address dst, bool flag, string details, uint256 cid, uint256 newValue)
func (_Proposal *ProposalCallerSession) Proposals(arg0 [32]byte) (struct {
	Proposer     common.Address
	CreateTime   *big.Int
	CreateBlock  *big.Int
	ProposalType *big.Int
	Dst          common.Address
	Flag         bool
	Details      string
	Cid          *big.Int
	NewValue     *big.Int
}, error) {
	return _Proposal.Contract.Proposals(&_Proposal.CallOpts, arg0)
}

// ProposerNonces is a free data retrieval call binding the contract method 0x9e4353cc.
//
// Solidity: function proposerNonces(address ) view returns(uint256)
func (_Proposal *ProposalCaller) ProposerNonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "proposerNonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposerNonces is a free data retrieval call binding the contract method 0x9e4353cc.
//
// Solidity: function proposerNonces(address ) view returns(uint256)
func (_Proposal *ProposalSession) ProposerNonces(arg0 common.Address) (*big.Int, error) {
	return _Proposal.Contract.ProposerNonces(&_Proposal.CallOpts, arg0)
}

// ProposerNonces is a free data retrieval call binding the contract method 0x9e4353cc.
//
// Solidity: function proposerNonces(address ) view returns(uint256)
func (_Proposal *ProposalCallerSession) ProposerNonces(arg0 common.Address) (*big.Int, error) {
	return _Proposal.Contract.ProposerNonces(&_Proposal.CallOpts, arg0)
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

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (_Proposal *ProposalCaller) UnbondingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "unbondingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (_Proposal *ProposalSession) UnbondingPeriod() (*big.Int, error) {
	return _Proposal.Contract.UnbondingPeriod(&_Proposal.CallOpts)
}

// UnbondingPeriod is a free data retrieval call binding the contract method 0x6cf6d675.
//
// Solidity: function unbondingPeriod() view returns(uint256)
func (_Proposal *ProposalCallerSession) UnbondingPeriod() (*big.Int, error) {
	return _Proposal.Contract.UnbondingPeriod(&_Proposal.CallOpts)
}

// ValidatorUnjailPeriod is a free data retrieval call binding the contract method 0xf945b623.
//
// Solidity: function validatorUnjailPeriod() view returns(uint256)
func (_Proposal *ProposalCaller) ValidatorUnjailPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "validatorUnjailPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ValidatorUnjailPeriod is a free data retrieval call binding the contract method 0xf945b623.
//
// Solidity: function validatorUnjailPeriod() view returns(uint256)
func (_Proposal *ProposalSession) ValidatorUnjailPeriod() (*big.Int, error) {
	return _Proposal.Contract.ValidatorUnjailPeriod(&_Proposal.CallOpts)
}

// ValidatorUnjailPeriod is a free data retrieval call binding the contract method 0xf945b623.
//
// Solidity: function validatorUnjailPeriod() view returns(uint256)
func (_Proposal *ProposalCallerSession) ValidatorUnjailPeriod() (*big.Int, error) {
	return _Proposal.Contract.ValidatorUnjailPeriod(&_Proposal.CallOpts)
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
// Solidity: function createProposal(address dst, bool flag, string details) returns(bytes32)
func (_Proposal *ProposalTransactor) CreateProposal(opts *bind.TransactOpts, dst common.Address, flag bool, details string) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "createProposal", dst, flag, details)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xd51cade8.
//
// Solidity: function createProposal(address dst, bool flag, string details) returns(bytes32)
func (_Proposal *ProposalSession) CreateProposal(dst common.Address, flag bool, details string) (*types.Transaction, error) {
	return _Proposal.Contract.CreateProposal(&_Proposal.TransactOpts, dst, flag, details)
}

// CreateProposal is a paid mutator transaction binding the contract method 0xd51cade8.
//
// Solidity: function createProposal(address dst, bool flag, string details) returns(bytes32)
func (_Proposal *ProposalTransactorSession) CreateProposal(dst common.Address, flag bool, details string) (*types.Transaction, error) {
	return _Proposal.Contract.CreateProposal(&_Proposal.TransactOpts, dst, flag, details)
}

// CreateUpdateConfigProposal is a paid mutator transaction binding the contract method 0xd24806eb.
//
// Solidity: function createUpdateConfigProposal(uint256 cid, uint256 newValue) returns(bytes32)
func (_Proposal *ProposalTransactor) CreateUpdateConfigProposal(opts *bind.TransactOpts, cid *big.Int, newValue *big.Int) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "createUpdateConfigProposal", cid, newValue)
}

// CreateUpdateConfigProposal is a paid mutator transaction binding the contract method 0xd24806eb.
//
// Solidity: function createUpdateConfigProposal(uint256 cid, uint256 newValue) returns(bytes32)
func (_Proposal *ProposalSession) CreateUpdateConfigProposal(cid *big.Int, newValue *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.CreateUpdateConfigProposal(&_Proposal.TransactOpts, cid, newValue)
}

// CreateUpdateConfigProposal is a paid mutator transaction binding the contract method 0xd24806eb.
//
// Solidity: function createUpdateConfigProposal(uint256 cid, uint256 newValue) returns(bytes32)
func (_Proposal *ProposalTransactorSession) CreateUpdateConfigProposal(cid *big.Int, newValue *big.Int) (*types.Transaction, error) {
	return _Proposal.Contract.CreateUpdateConfigProposal(&_Proposal.TransactOpts, cid, newValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] vals, address validators_) returns()
func (_Proposal *ProposalTransactor) Initialize(opts *bind.TransactOpts, vals []common.Address, validators_ common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "initialize", vals, validators_)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] vals, address validators_) returns()
func (_Proposal *ProposalSession) Initialize(vals []common.Address, validators_ common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.Initialize(&_Proposal.TransactOpts, vals, validators_)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] vals, address validators_) returns()
func (_Proposal *ProposalTransactorSession) Initialize(vals []common.Address, validators_ common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.Initialize(&_Proposal.TransactOpts, vals, validators_)
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
