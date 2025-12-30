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
	ABI: "[{\"type\":\"function\",\"name\":\"PROPOSAL_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUNISH_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blockReward\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createProposal\",\"inputs\":[{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createUpdateConfigProposal\",\"inputs\":[{\"name\":\"cid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreaseRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"vals\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"validators_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isProposalValidForStaking\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minValidatorStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pass\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalLastingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalPassedHeight\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposals\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"createBlock\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposalType\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"cid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposerNonces\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"punishThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"results\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"agree\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"reject\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"resultExist\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUnpassed\",\"inputs\":[{\"name\":\"val\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbondingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorUnjailPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voteProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"auth\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"votes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"voteTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"auth\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawProfitPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"LogCreateConfigProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogCreateProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dst\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogPassProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogRejectProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogSetUnpassed\",\"inputs\":[{\"name\":\"val\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogVote\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"auth\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055611bbd806100405f395ff3fe608060405234801561000f575f5ffd5b50600436106101a1575f3560e01c80635d4f0cb6116100f35780639f759dba11610093578063d24806eb1161006e578063d24806eb146103f9578063d51cade81461040c578063e823c8141461041f578063f945b62314610428575f5ffd5b80639f759dba146103d4578063a4c4d922146103dd578063cb1ea725146103f0575f5ffd5b806382c4b3b2116100ce57806382c4b3b2146103815780638c872d05146103a357806394522b6d146103ac5780639e4353cc146103b5575f5ffd5b80635d4f0cb6146103505780636cf6d675146103595780637668dd2414610362575f5ffd5b80631db5ade81161015e578063437ccda811610139578063437ccda8146102b357806344c1aa99146102d4578063462d0b2e146102dd5780634c6b25b1146102f2575f5ffd5b80631db5ade8146102155780632897183d1461028257806332ed5b121461028b575f5ffd5b8063017ddd35146101a557806308ac5256146101c15780630ac168a1146101ca578063107a03a8146101d3578063158ef93e146101f657806315ea278114610202575b5f5ffd5b6101ae60095481565b6040519081526020015b60405180910390f35b6101ae600a5481565b6101ae60065481565b6101e66101e13660046116cf565b610431565b60405190151581526020016101b8565b5f546101e69060ff1681565b6101e66102103660046116cf565b610486565b61025b6102233660046116ef565b601060209081525f92835260408084209091529082529020805460018201546002909201546001600160a01b03909116919060ff1683565b604080516001600160a01b03909416845260208401929092521515908201526060016101b8565b6101ae60045481565b61029e610299366004611717565b6104f5565b6040516101b89998979695949392919061172e565b6102bc61f01281565b6040516001600160a01b0390911681526020016101b8565b6101ae60035481565b6102f06102eb3660046117b3565b6105d7565b005b61032d610300366004611717565b600f6020525f908152604090205461ffff8082169162010000810490911690640100000000900460ff1683565b6040805161ffff94851681529390921660208401521515908201526060016101b8565b6102bc61f01381565b6101ae60075481565b6101ae6103703660046116cf565b600c6020525f908152604090205481565b6101e661038f3660046116cf565b600b6020525f908152604090205460ff1681565b6102bc61f01181565b6101ae60055481565b6101ae6103c33660046116cf565b600d6020525f908152604090205481565b6102bc61f01081565b6101e66103eb366004611842565b6107e8565b6101ae60025481565b6101ae610407366004611870565b610e17565b6101ae61041a366004611890565b61101e565b6101ae60015481565b6101ae60085481565b6001600160a01b0381165f908152600b602052604081205460ff1661045757505f919050565b6001600160a01b0382165f908152600c602052604090205460015461047c9082611930565b4311159392505050565b5f61048f611397565b6001600160a01b0382165f818152600b60209081526040808320805460ff19169055600c82528083209290925590514281527f4e0b191f7f5c32b1b5e3704b68874b1a3980147cae00be8ece271bfb5b92c07a910160405180910390a25060015b919050565b600e6020525f90815260409020805460018201546002830154600384015460048501546005860180546001600160a01b0396871697959694959394831693600160a01b90930460ff1692919061054a90611943565b80601f016020809104026020016040519081016040528092919081815260200182805461057690611943565b80156105c15780601f10610598576101008083540402835291602001916105c1565b820191905f5260205f20905b8154815290600101906020018083116105a457829003601f168201915b5050505050908060060154908060070154905089565b6105df6113ea565b6001600160a01b03811661063a5760405162461bcd60e51b815260206004820152601a60248201527f496e76616c69642076616c696461746f7273206164647265737300000000000060448201526064015b60405180910390fd5b601180546001600160a01b0319166001600160a01b0383161790555f5b8281101561078a575f8484838181106106725761067261197b565b905060200201602081019061068791906116cf565b6001600160a01b0316036106dd5760405162461bcd60e51b815260206004820152601960248201527f496e76616c69642076616c696461746f722061646472657373000000000000006044820152606401610631565b6001600b5f8686858181106106f4576106f461197b565b905060200201602081019061070991906116cf565b6001600160a01b03166001600160a01b031681526020019081526020015f205f6101000a81548160ff02191690831515021790555043600c5f8686858181106107545761075461197b565b905060200201602081019061076991906116cf565b6001600160a01b0316815260208101919091526040015f2055600101610657565b505062093a8060018181556018600281905560306003556004556201518060058190556702c68af0bb14000060065560079290925560089190915569152d02c7e14af68000006009556015600a555f805460ff191690911790555050565b5f6107f1611432565b6107f96114d9565b5f838152600e6020526040812060010154900361084d5760405162461bcd60e51b8152602060048201526012602482015271141c9bdc1bdcd85b081b9bdd08195e1a5cdd60721b6044820152606401610631565b335f908152601060209081526040808320868452909152902060010154156108c35760405162461bcd60e51b815260206004820152602360248201527f596f752063616e277420766f746520666f7220612070726f706f73616c20747760448201526269636560e81b6064820152608401610631565b6001545f848152600e60205260409020600201546108e19190611930565b43106109225760405162461bcd60e51b815260206004820152601060248201526f141c9bdc1bdcd85b08195e1c1a5c995960821b6044820152606401610631565b335f818152601060209081526040808320878452825291829020426001820181905581546001600160a01b031916851782556002909101805460ff191687151590811790915583519081529182015285917f6c59bda68cac318717c60c7c9635a78a0f0613f9887cc18a7157f5745a86d14e910160405180910390a381156109e7575f838152600f60205260409020546109c19061ffff16600161198f565b5f848152600f60205260409020805461ffff191661ffff92909216919091179055610a35565b5f838152600f6020526040902054610a0a9062010000900461ffff16600161198f565b5f848152600f60205260409020805461ffff92909216620100000263ffff0000199092169190911790555b5f838152600f6020526040902054640100000000900460ff1615610a5b57506001610de8565b6011546040805163037deea760e41b815290516002926001600160a01b0316916337deea709160048083019260209291908290030181865afa158015610aa3573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ac791906119a9565b610ad191906119c0565b610adc906001611930565b5f848152600f602052604090205461ffff1610610ce9575f838152600f60209081526040808320805464ff000000001916640100000000179055600e9091529020600301546001819003610c36575f848152600e6020526040902060040154600160a01b900460ff1615610b92575f848152600e6020908152604080832060040180546001600160a01b039081168552600b8452828520805460ff191660011790559054168352600c9091529020439055610ca5565b5f848152600e60208181526040808420600490810180546001600160a01b039081168752600b8552838720805460ff19169055815481168752600c8552838720879055601154968b9052949093529154905163a1ff465560e01b81529083169181019190915291169063a1ff4655906024015f604051808303815f87803b158015610c1b575f5ffd5b505af1158015610c2d573d5f5f3e3d5ffd5b50505050610ca5565b80600203610c65575f848152600e602052604090206006810154600790910154610c609190611507565b610ca5565b60405162461bcd60e51b8152602060048201526015602482015274496e76616c69642070726f706f73616c207479706560581b6044820152606401610631565b837f90d2e923947d9356c1c04391cb9e2e9c5d4ad6c165a849787b0c7569bbe99e2442604051610cd791815260200190565b60405180910390a26001915050610de8565b6011546040805163037deea760e41b815290516002926001600160a01b0316916337deea709160048083019260209291908290030181865afa158015610d31573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d5591906119a9565b610d5f91906119c0565b610d6a906001611930565b5f848152600f602052604090205462010000900461ffff1610610de4575f838152600f602052604090819020805464ff0000000019166401000000001790555183907f36bdb56d707cdf53eadffe319a71ddf97736be67b8caab47b7720201a6b65ca090610ddb9042815260200190565b60405180910390a25b5060015b610e1160017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b92915050565b5f610e20611432565b610e2a83836115d9565b610e765760405162461bcd60e51b815260206004820152601860248201527f436f6e6669672076616c69646174696f6e206661696c656400000000000000006044820152606401610631565b335f818152600d60208181526040808420805482518085018890528084018b9052606081018a905260808082018390528451808303909101815260a0909101909352825192840192909220958552929091529291839190610ed6836119df565b909155505060408051610120810182525f8082526020808301828152838501838152606085018481526080860185815260a0870186815288518087018a5287815260c0890190815233895260e089018f905261010089018e905242865243855260028085528b8952600e90975298909620875181546001600160a01b039182166001600160a01b031990911617825594516001820155925194830194909455516003820155915160048301805494511515600160a01b026001600160a81b03199095169190921617929092179091559151909182916005820190610fba9082611a57565b5060e0820151600682015561010090910151600790910155604080518781526020810187905242818301529051339184917f8bfc061277ae1778974ada10db7f9664ab1d67c455c025c025b438c52c69d1819181900360600190a350949350505050565b5f611027611432565b6001600160a01b0385165f908152600b602052604090205460ff1615801561104c5750835b8061107757506001600160a01b0385165f908152600b602052604090205460ff168015611077575083155b6110e95760405162461bcd60e51b815260206004820152603f60248201527f43616e27742061646420616e20616c726561647920657869737420647374206f60448201527f722043616e27742072656d6f76652061206e6f742070617373656420647374006064820152608401610631565b335f818152600d6020908152604080832054905190936111159290918a918a918a918a91899101611b12565b60408051601f1981840301815291905280516020909101209050610bb88411156111745760405162461bcd60e51b815260206004820152601060248201526f44657461696c7320746f6f206c6f6e6760801b6044820152606401610631565b5f818152600e6020526040902060010154156111d25760405162461bcd60e51b815260206004820152601760248201527f50726f706f73616c20616c7265616479206578697374730000000000000000006044820152606401610631565b335f908152600d602052604081208054916111ec836119df565b909155505060408051610120810182525f8082526020808301829052828401829052606083018290526080830182815260a084018381528551808401875284815260c086015260e085018490526101008501939093523384526001600160a01b038c1690528915159091528251601f880182900482028101820190935286835290919087908790819084018382808284375f92018290525060c08601948552426020808801918252436040808a01918252600160608b018181528c8752600e9094529420895181546001600160a01b0319166001600160a01b03918216178255935194810194909455516002840155516003830155608087015160048301805460a08a0151929093166001600160a81b031990931692909217600160a01b9115159190910217905593518594935060058401925061132b915082611a57565b5060e08201516006820155610100909101516007909101556040805188151581524260208201526001600160a01b038a1691339185917f1af05d46b8c1ec021d82b7128cff40e91a1c2337deffc010df48eeddef8da56c910160405180910390a4509695505050505050565b3361f010146113e85760405162461bcd60e51b815260206004820152601860248201527f56616c696461746f727320636f6e7472616374206f6e6c7900000000000000006044820152606401610631565b565b5f5460ff16156113e85760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b6044820152606401610631565b601154604051631015428760e21b81523360048201526001600160a01b03909116906340550a1c90602401602060405180830381865afa158015611478573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061149c9190611b6c565b6113e85760405162461bcd60e51b815260206004820152600e60248201526d56616c696461746f72206f6e6c7960901b6044820152606401610631565b6114e1611677565b60027f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b61151182826115d9565b50815f0361151f5760015550565b8160010361152d5760025550565b8160020361153b5760035550565b816003036115495760045550565b816004036115575760055550565b816005036115655760065550565b816006036115735760075550565b816007036115815760085550565b8160080361158f5760095550565b8160090361159d57600a5550565b60405162461bcd60e51b8152602060048201526011602482015270155b9adb9bdddb8818dbdb999a59c81251607a1b6044820152606401610631565b5f600983111561161f5760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a590818dbdb999a59c81251607a1b6044820152606401610631565b5f821161166e5760405162461bcd60e51b815260206004820152601d60248201527f436f6e6669672076616c7565206d75737420626520706f7369746976650000006044820152606401610631565b50600192915050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00546002036113e857604051633ee5aeb560e01b815260040160405180910390fd5b80356001600160a01b03811681146104f0575f5ffd5b5f602082840312156116df575f5ffd5b6116e8826116b9565b9392505050565b5f5f60408385031215611700575f5ffd5b611709836116b9565b946020939093013593505050565b5f60208284031215611727575f5ffd5b5035919050565b60018060a01b038a16815288602082015287604082015286606082015260018060a01b038616608082015284151560a082015261012060c08201525f845180610120840152806020870161014085015e5f6101408285010152610140601f19601f8301168401019150508360e0830152826101008301529a9950505050505050505050565b5f5f5f604084860312156117c5575f5ffd5b833567ffffffffffffffff8111156117db575f5ffd5b8401601f810186136117eb575f5ffd5b803567ffffffffffffffff811115611801575f5ffd5b8660208260051b8401011115611815575f5ffd5b6020918201945092506118299085016116b9565b90509250925092565b801515811461183f575f5ffd5b50565b5f5f60408385031215611853575f5ffd5b82359150602083013561186581611832565b809150509250929050565b5f5f60408385031215611881575f5ffd5b50508035926020909101359150565b5f5f5f5f606085870312156118a3575f5ffd5b6118ac856116b9565b935060208501356118bc81611832565b9250604085013567ffffffffffffffff8111156118d7575f5ffd5b8501601f810187136118e7575f5ffd5b803567ffffffffffffffff8111156118fd575f5ffd5b87602082840101111561190e575f5ffd5b949793965060200194505050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610e1157610e1161191c565b600181811c9082168061195757607f821691505b60208210810361197557634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b61ffff8181168382160190811115610e1157610e1161191c565b5f602082840312156119b9575f5ffd5b5051919050565b5f826119da57634e487b7160e01b5f52601260045260245ffd5b500490565b5f600182016119f0576119f061191c565b5060010190565b634e487b7160e01b5f52604160045260245ffd5b601f821115611a5257805f5260205f20601f840160051c81016020851015611a305750805b601f840160051c820191505b81811015611a4f575f8155600101611a3c565b50505b505050565b815167ffffffffffffffff811115611a7157611a716119f7565b611a8581611a7f8454611943565b84611a0b565b6020601f821160018114611ab7575f8315611aa05750848201515b5f19600385901b1c1916600184901b178455611a4f565b5f84815260208120601f198516915b82811015611ae65787850151825560209485019460019092019101611ac6565b5084821015611b0357868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b6001600160a01b03878116825286166020820152841515604082015260a0606082018190528101839052828460c08301375f60c084830101525f60c0601f19601f8601168301019050826080830152979650505050505050565b5f60208284031215611b7c575f5ffd5b81516116e88161183256fea26469706673582212202aa56d5693ca7686bda7dcb75055b81016e8fd4c6f6c6ccba5108eadbc18a35f64736f6c634300081d0033",
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
