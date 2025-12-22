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
	ABI: "[{\"type\":\"function\",\"name\":\"PROPOSAL_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PUNISH_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"STAKING_DEADLINE_PERIOD\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VALIDATOR_CONTRACT_ADDR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blockReward\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createProposal\",\"inputs\":[{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createUpdateConfigProposal\",\"inputs\":[{\"name\":\"cid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decreaseRate\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"vals\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_validators\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialized\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isProposalValidForStaking\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maxValidators\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minValidatorStake\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pass\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalLastingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposalPassedTime\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposals\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"createTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposalType\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"dst\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"details\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"cid\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"punishThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"removeThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"results\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"agree\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"reject\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"resultExist\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setUnpassed\",\"inputs\":[{\"name\":\"val\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unbondingPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"validatorUnjailPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"voteProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"auth\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"votes\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"voter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"voteTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"auth\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawProfitPeriod\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"LogCreateConfigProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"cid\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newValue\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogCreateProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"dst\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"flag\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogPassProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogRejectProposal\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogSetUnpassed\",\"inputs\":[{\"name\":\"val\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LogVote\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"auth\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"time\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x608060405234801561000f575f80fd5b50611a418061001d5f395ff3fe608060405234801561000f575f80fd5b50600436106101a1575f3560e01c8063462d0b2e116100f3578063cb1ea72511610093578063e823c8141161006e578063e823c814146103f6578063ebddb3cf146103ff578063f945b62314610409578063f9a2bbc714610412575f80fd5b8063cb1ea725146103c7578063d24806eb146103d0578063d51cade8146103e3575f80fd5b806371cbeb3a116100ce57806371cbeb3a1461036a57806382c4b3b21461038957806394522b6d146103ab578063a4c4d922146103b4575f80fd5b8063462d0b2e146102ee5780634c6b25b1146103035780636cf6d67514610361575f80fd5b806315ea27811161015e57806332ed5b121161013957806332ed5b12146102ac5780633bbf0865146102d3578063437ccda8146102dc57806344c1aa99146102e5575f80fd5b806315ea2781146102235780631db5ade8146102365780632897183d146102a3575f80fd5b8063017ddd35146101a557806308ac5256146101c15780630ac168a1146101ca5780630e2374a5146101d3578063107a03a8146101f4578063158ef93e14610217575b5f80fd5b6101ae60095481565b6040519081526020015b60405180910390f35b6101ae600a5481565b6101ae60065481565b6101dc61f01381565b6040516001600160a01b0390911681526020016101b8565b61020761020236600461154b565b61041b565b60405190151581526020016101b8565b5f546102079060ff1681565b61020761023136600461154b565b610470565b61027c61024436600461156b565b600f60209081525f92835260408084209091529082529020805460018201546002909201546001600160a01b03909116919060ff1683565b604080516001600160a01b03909416845260208401929092521515908201526060016101b8565b6101ae60045481565b6102bf6102ba366004611593565b6104df565b6040516101b89897969594939291906115aa565b6101dc61f01181565b6101dc61f01281565b6101ae60035481565b6103016102fc366004611642565b6105bb565b005b61033e610311366004611593565b600e6020525f908152604090205461ffff8082169162010000810490911690640100000000900460ff1683565b6040805161ffff94851681529390921660208401521515908201526060016101b8565b6101ae60075481565b6101ae61037836600461154b565b600c6020525f908152604090205481565b61020761039736600461154b565b600b6020525f908152604090205460ff1681565b6101ae60055481565b6102076103c23660046116cf565b6107d6565b6101ae60025481565b6102076103de3660046116fd565b610d99565b6102076103f136600461171d565b610f09565b6101ae60015481565b6101ae62093a8081565b6101ae60085481565b6101dc61f01081565b6001600160a01b0381165f908152600b602052604081205460ff1661044157505f919050565b6001600160a01b0382165f908152600c602052604090205461046662093a80826117ba565b4211159392505050565b5f610479611205565b6001600160a01b0382165f818152600b60209081526040808320805460ff19169055600c82528083209290925590514281527f4e0b191f7f5c32b1b5e3704b68874b1a3980147cae00be8ece271bfb5b92c07a910160405180910390a25060015b919050565b600d6020525f9081526040902080546001820154600283015460038401546004850180546001600160a01b03958616969495939493831693600160a01b90930460ff1692919061052e906117cd565b80601f016020809104026020016040519081016040528092919081815260200182805461055a906117cd565b80156105a55780601f1061057c576101008083540402835291602001916105a5565b820191905f5260205f20905b81548152906001019060200180831161058857829003601f168201915b5050505050908060050154908060060154905088565b6105c3611258565b6001600160a01b03811661061e5760405162461bcd60e51b815260206004820152601a60248201527f496e76616c69642076616c696461746f7273206164647265737300000000000060448201526064015b60405180910390fd5b601080546001600160a01b0319166001600160a01b0383161790555f5b82811015610778575f84848381811061065657610656611805565b905060200201602081019061066b919061154b565b6001600160a01b0316036106c15760405162461bcd60e51b815260206004820152601960248201527f496e76616c69642076616c696461746f722061646472657373000000000000006044820152606401610615565b6001600b5f8686858181106106d8576106d8611805565b90506020020160208101906106ed919061154b565b6001600160a01b03166001600160a01b031681526020019081526020015f205f6101000a81548160ff02191690831515021790555042600c5f86868581811061073857610738611805565b905060200201602081019061074d919061154b565b6001600160a01b0316815260208101919091526040015f20558061077081611819565b91505061063b565b505062093a8060018181556018600281905560306003556004556201518060058190556702c68af0bb14000060065560079290925560089190915569152d02c7e14af68000006009556015600a555f805460ff191690911790555050565b5f6107df6112a0565b5f838152600d602052604081206001015490036108335760405162461bcd60e51b8152602060048201526012602482015271141c9bdc1bdcd85b081b9bdd08195e1a5cdd60721b6044820152606401610615565b335f908152600f60209081526040808320868452909152902060010154156108a95760405162461bcd60e51b815260206004820152602360248201527f596f752063616e277420766f746520666f7220612070726f706f73616c20747760448201526269636560e81b6064820152608401610615565b600180545f858152600d60205260409020909101546108c891906117ba565b42106109095760405162461bcd60e51b815260206004820152601060248201526f141c9bdc1bdcd85b08195e1c1a5c995960821b6044820152606401610615565b335f818152600f60209081526040808320878452825291829020426001820181905581546001600160a01b031916851782556002909101805460ff191687151590811790915583519081529182015285917f6c59bda68cac318717c60c7c9635a78a0f0613f9887cc18a7157f5745a86d14e910160405180910390a381156109ce575f838152600e60205260409020546109a89061ffff166001611831565b5f848152600e60205260409020805461ffff191661ffff92909216919091179055610a1c565b5f838152600e60205260409020546109f19062010000900461ffff166001611831565b5f848152600e60205260409020805461ffff92909216620100000263ffff0000199092169190911790555b5f838152600e6020526040902054640100000000900460ff1615610a4257506001610d93565b6010546040805163037deea760e41b815290516002926001600160a01b0316916337deea709160048083019260209291908290030181865afa158015610a8a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aae9190611853565b610ab8919061186a565b610ac39060016117ba565b5f848152600e602052604090205461ffff1610610c94575f838152600e60209081526040808320805464ff000000001916640100000000179055600d909152902060020154600103610c17575f838152600d6020526040902060030154600160a01b900460ff1615610b77575f838152600d6020908152604080832060030180546001600160a01b039081168552600b8452828520805460ff191660011790559054168352600c9091529020429055610c52565b5f838152600d6020818152604080842060030180546001600160a01b039081168652600b8452828620805460ff19169055815481168652600c845282862086905560105495899052939092529054905163a1ff465560e01b8152908216600482015291169063a1ff4655906024015f604051808303815f87803b158015610bfc575f80fd5b505af1158015610c0e573d5f803e3d5ffd5b50505050610c52565b5f838152600d602052604090206002908101549003610c52575f838152600d602052604090206005810154600690910154610c529190611347565b827f90d2e923947d9356c1c04391cb9e2e9c5d4ad6c165a849787b0c7569bbe99e2442604051610c8491815260200190565b60405180910390a2506001610d93565b6010546040805163037deea760e41b815290516002926001600160a01b0316916337deea709160048083019260209291908290030181865afa158015610cdc573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d009190611853565b610d0a919061186a565b610d159060016117ba565b5f848152600e602052604090205462010000900461ffff1610610d8f575f838152600e602052604090819020805464ff0000000019166401000000001790555183907f36bdb56d707cdf53eadffe319a71ddf97736be67b8caab47b7720201a6b65ca090610d869042815260200190565b60405180910390a25b5060015b92915050565b5f610da26112a0565b610dac83836113e1565b506040516bffffffffffffffffffffffff193360601b16602082015260348101849052605481018390524260748201525f90609401604051602081830303815290604052805190602001209050610e016114e3565b33815260c0810185905260e08101849052426020808301918252600260408085018281525f878152600d9094529220845181546001600160a01b0319166001600160a01b039182161782559351600182015591519082015560608301516003820180546080860151929094166001600160a81b031990941693909317600160a01b911515919091021790915560a08201518291906004820190610ea490826118eb565b5060c0820151600582015560e090910151600690910155604080518681526020810186905242818301529051339184917f8bfc061277ae1778974ada10db7f9664ab1d67c455c025c025b438c52c69d1819181900360600190a3506001949350505050565b5f610f126112a0565b6001600160a01b0385165f908152600b602052604090205460ff16158015610f375750835b80610f6257506001600160a01b0385165f908152600b602052604090205460ff168015610f62575083155b610fd45760405162461bcd60e51b815260206004820152603f60248201527f43616e27742061646420616e20616c726561647920657869737420647374206f60448201527f722043616e27742072656d6f76652061206e6f742070617373656420647374006064820152608401610615565b5f338686868642604051602001610ff0969594939291906119a7565b60408051601f1981840301815291905280516020909101209050610bb883111561104f5760405162461bcd60e51b815260206004820152601060248201526f44657461696c7320746f6f206c6f6e6760801b6044820152606401610615565b5f818152600d6020526040902060010154156110ad5760405162461bcd60e51b815260206004820152601760248201527f50726f706f73616c20616c7265616479206578697374730000000000000000006044820152606401610615565b6110b56114e3565b3381526001600160a01b03871660608201528515156080820152604080516020601f87018190048102820181019092528581529086908690819084018382808284375f92018290525060a0860194855242602080880191825260016040808a018281528b8652600d909352909320885181546001600160a01b0319166001600160a01b03918216178255925193810193909355516002830155606087015160038301805460808a0151929093166001600160a81b031990931692909217600160a01b911515919091021790559351859493506004840192506111989150826118eb565b5060c0820151600582015560e0909101516006909101556040805187151581524260208201526001600160a01b03891691339185917f1af05d46b8c1ec021d82b7128cff40e91a1c2337deffc010df48eeddef8da56c910160405180910390a45060019695505050505050565b3361f010146112565760405162461bcd60e51b815260206004820152601860248201527f56616c696461746f727320636f6e7472616374206f6e6c7900000000000000006044820152606401610615565b565b5f5460ff16156112565760405162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b6044820152606401610615565b601054604051631015428760e21b81523360048201526001600160a01b03909116906340550a1c90602401602060405180830381865afa1580156112e6573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061130a91906119f0565b6112565760405162461bcd60e51b815260206004820152600e60248201526d56616c696461746f72206f6e6c7960901b6044820152606401610615565b61135182826113e1565b50815f0361135f5760018190555b8160010361136d5760028190555b8160020361137b5760038190555b816003036113895760048190555b816004036113975760058190555b816005036113a55760068190555b816006036113b35760078190555b816007036113c15760088190555b816008036113cf5760098190555b816009036113dd57600a8190555b5050565b5f60098311156114275760405162461bcd60e51b8152602060048201526011602482015270125b9d985b1a590818dbdb999a59c81251607a1b6044820152606401610615565b825f0361149457610e108210158015611443575062278d008211155b61148f5760405162461bcd60e51b815260206004820152601760248201527f496e76616c69642070726f706f73616c20706572696f640000000000000000006044820152606401610615565b610d8f565b5f8211610d8f5760405162461bcd60e51b815260206004820152601d60248201527f436f6e6669672076616c7565206d75737420626520706f7369746976650000006044820152606401610615565b6040518061010001604052805f6001600160a01b031681526020015f81526020015f81526020015f6001600160a01b031681526020015f15158152602001606081526020015f81526020015f81525090565b80356001600160a01b03811681146104da575f80fd5b5f6020828403121561155b575f80fd5b61156482611535565b9392505050565b5f806040838503121561157c575f80fd5b61158583611535565b946020939093013593505050565b5f602082840312156115a3575f80fd5b5035919050565b6001600160a01b03898116825260208083018a9052604083018990529087166060830152851515608083015261010060a0830181905285519083018190525f918291905b8183101561160d578783018101518584016101200152918201916115ee565b505f61012082860181019190915260c085019690965260e08401949094525050601f909101601f191601019695505050505050565b5f805f60408486031215611654575f80fd5b833567ffffffffffffffff8082111561166b575f80fd5b818601915086601f83011261167e575f80fd5b81358181111561168c575f80fd5b8760208260051b85010111156116a0575f80fd5b6020928301955093506116b69186019050611535565b90509250925092565b80151581146116cc575f80fd5b50565b5f80604083850312156116e0575f80fd5b8235915060208301356116f2816116bf565b809150509250929050565b5f806040838503121561170e575f80fd5b50508035926020909101359150565b5f805f8060608587031215611730575f80fd5b61173985611535565b93506020850135611749816116bf565b9250604085013567ffffffffffffffff80821115611765575f80fd5b818701915087601f830112611778575f80fd5b813581811115611786575f80fd5b886020828501011115611797575f80fd5b95989497505060200194505050565b634e487b7160e01b5f52601160045260245ffd5b80820180821115610d9357610d936117a6565b600181811c908216806117e157607f821691505b6020821081036117ff57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b5f6001820161182a5761182a6117a6565b5060010190565b61ffff81811683821601908082111561184c5761184c6117a6565b5092915050565b5f60208284031215611863575f80fd5b5051919050565b5f8261188457634e487b7160e01b5f52601260045260245ffd5b500490565b634e487b7160e01b5f52604160045260245ffd5b601f8211156118e6575f81815260208120601f850160051c810160208610156118c35750805b601f850160051c820191505b818110156118e2578281556001016118cf565b5050505b505050565b815167ffffffffffffffff81111561190557611905611889565b6119198161191384546117cd565b8461189d565b602080601f83116001811461194c575f84156119355750858301515b5f19600386901b1c1916600185901b1785556118e2565b5f85815260208120601f198616915b8281101561197a5788860151825594840194600190910190840161195b565b508582101561199757878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f6bffffffffffffffffffffffff19808960601b168352808860601b1660148401525085151560f81b602883015283856029840137506029920191820152604901949350505050565b5f60208284031215611a00575f80fd5b8151611564816116bf56fea2646970667358221220d042826b39568d40f25ea1058d39ffd521db96cbe59c368bad8149bafafce4ec64736f6c63430008140033",
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

// PUNISHCONTRACTADDR is a free data retrieval call binding the contract method 0x3bbf0865.
//
// Solidity: function PUNISH_CONTRACT_ADDR() view returns(address)
func (_Proposal *ProposalCaller) PUNISHCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "PUNISH_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PUNISHCONTRACTADDR is a free data retrieval call binding the contract method 0x3bbf0865.
//
// Solidity: function PUNISH_CONTRACT_ADDR() view returns(address)
func (_Proposal *ProposalSession) PUNISHCONTRACTADDR() (common.Address, error) {
	return _Proposal.Contract.PUNISHCONTRACTADDR(&_Proposal.CallOpts)
}

// PUNISHCONTRACTADDR is a free data retrieval call binding the contract method 0x3bbf0865.
//
// Solidity: function PUNISH_CONTRACT_ADDR() view returns(address)
func (_Proposal *ProposalCallerSession) PUNISHCONTRACTADDR() (common.Address, error) {
	return _Proposal.Contract.PUNISHCONTRACTADDR(&_Proposal.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Proposal *ProposalCaller) STAKINGCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "STAKING_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Proposal *ProposalSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Proposal.Contract.STAKINGCONTRACTADDR(&_Proposal.CallOpts)
}

// STAKINGCONTRACTADDR is a free data retrieval call binding the contract method 0x0e2374a5.
//
// Solidity: function STAKING_CONTRACT_ADDR() view returns(address)
func (_Proposal *ProposalCallerSession) STAKINGCONTRACTADDR() (common.Address, error) {
	return _Proposal.Contract.STAKINGCONTRACTADDR(&_Proposal.CallOpts)
}

// STAKINGDEADLINEPERIOD is a free data retrieval call binding the contract method 0xebddb3cf.
//
// Solidity: function STAKING_DEADLINE_PERIOD() view returns(uint256)
func (_Proposal *ProposalCaller) STAKINGDEADLINEPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "STAKING_DEADLINE_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STAKINGDEADLINEPERIOD is a free data retrieval call binding the contract method 0xebddb3cf.
//
// Solidity: function STAKING_DEADLINE_PERIOD() view returns(uint256)
func (_Proposal *ProposalSession) STAKINGDEADLINEPERIOD() (*big.Int, error) {
	return _Proposal.Contract.STAKINGDEADLINEPERIOD(&_Proposal.CallOpts)
}

// STAKINGDEADLINEPERIOD is a free data retrieval call binding the contract method 0xebddb3cf.
//
// Solidity: function STAKING_DEADLINE_PERIOD() view returns(uint256)
func (_Proposal *ProposalCallerSession) STAKINGDEADLINEPERIOD() (*big.Int, error) {
	return _Proposal.Contract.STAKINGDEADLINEPERIOD(&_Proposal.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Proposal *ProposalCaller) VALIDATORCONTRACTADDR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "VALIDATOR_CONTRACT_ADDR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Proposal *ProposalSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Proposal.Contract.VALIDATORCONTRACTADDR(&_Proposal.CallOpts)
}

// VALIDATORCONTRACTADDR is a free data retrieval call binding the contract method 0xf9a2bbc7.
//
// Solidity: function VALIDATOR_CONTRACT_ADDR() view returns(address)
func (_Proposal *ProposalCallerSession) VALIDATORCONTRACTADDR() (common.Address, error) {
	return _Proposal.Contract.VALIDATORCONTRACTADDR(&_Proposal.CallOpts)
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

// ProposalPassedTime is a free data retrieval call binding the contract method 0x71cbeb3a.
//
// Solidity: function proposalPassedTime(address ) view returns(uint256)
func (_Proposal *ProposalCaller) ProposalPassedTime(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Proposal.contract.Call(opts, &out, "proposalPassedTime", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalPassedTime is a free data retrieval call binding the contract method 0x71cbeb3a.
//
// Solidity: function proposalPassedTime(address ) view returns(uint256)
func (_Proposal *ProposalSession) ProposalPassedTime(arg0 common.Address) (*big.Int, error) {
	return _Proposal.Contract.ProposalPassedTime(&_Proposal.CallOpts, arg0)
}

// ProposalPassedTime is a free data retrieval call binding the contract method 0x71cbeb3a.
//
// Solidity: function proposalPassedTime(address ) view returns(uint256)
func (_Proposal *ProposalCallerSession) ProposalPassedTime(arg0 common.Address) (*big.Int, error) {
	return _Proposal.Contract.ProposalPassedTime(&_Proposal.CallOpts, arg0)
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

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] vals, address _validators) returns()
func (_Proposal *ProposalTransactor) Initialize(opts *bind.TransactOpts, vals []common.Address, _validators common.Address) (*types.Transaction, error) {
	return _Proposal.contract.Transact(opts, "initialize", vals, _validators)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] vals, address _validators) returns()
func (_Proposal *ProposalSession) Initialize(vals []common.Address, _validators common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.Initialize(&_Proposal.TransactOpts, vals, _validators)
}

// Initialize is a paid mutator transaction binding the contract method 0x462d0b2e.
//
// Solidity: function initialize(address[] vals, address _validators) returns()
func (_Proposal *ProposalTransactorSession) Initialize(vals []common.Address, _validators common.Address) (*types.Transaction, error) {
	return _Proposal.Contract.Initialize(&_Proposal.TransactOpts, vals, _validators)
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
