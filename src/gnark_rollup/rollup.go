// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rollup

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

// RollupAccount is an auto generated low-level Go binding around an user-defined struct.
type RollupAccount struct {
	Index   uint64
	Nonce   *big.Int
	Balance *big.Int
	PubKey  RollupPublicKey
}

// RollupPublicKey is an auto generated low-level Go binding around an user-defined struct.
type RollupPublicKey struct {
	X *big.Int
	Y *big.Int
}

// MerkleTreeMetaData contains all meta data concerning the MerkleTree contract.
var MerkleTreeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_levels\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ZERO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"computeMerkleRootFromPath\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLevels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextLeafIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"}],\"name\":\"hashLeftRight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"levels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"zeros\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040516200111638038062001116833981810160405281019062000036919062000177565b5f81116200007b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000072906200022b565b60405180910390fd5b60098110620000c1576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000b89062000299565b60405180910390fd5b6004811462000107576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000fe906200032d565b60405180910390fd5b805f819055507f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd32267600181905550506200034d565b5f80fd5b5f819050919050565b62000153816200013f565b81146200015e575f80fd5b50565b5f81519050620001718162000148565b92915050565b5f602082840312156200018f576200018e6200013b565b5b5f6200019e8482850162000161565b91505092915050565b5f82825260208201905092915050565b7f5f6c6576656c732073686f756c642062652067726561746572207468616e207a5f8201527f65726f0000000000000000000000000000000000000000000000000000000000602082015250565b5f62000213602383620001a7565b91506200022082620001b7565b604082019050919050565b5f6020820190508181035f830152620002448162000205565b9050919050565b7f5f6c6576656c732073686f756c64206265206c657373207468616e20380000005f82015250565b5f62000281601d83620001a7565b91506200028e826200024b565b602082019050919050565b5f6020820190508181035f830152620002b28162000273565b9050919050565b7f5f6c6576656c732073686f756c6420626520342028325e34203d2031362061635f8201527f636f756e74732920666f722074657374696e6700000000000000000000000000602082015250565b5f62000315603383620001a7565b91506200032282620002b9565b604082019050919050565b5f6020820190508181035f830152620003468162000307565b9050919050565b610dbb806200035b5f395ff3fe608060405234801561000f575f80fd5b5060043610610091575f3560e01c80635bb93995116100645780635bb939951461011f5780635ca1e1651461014f57806369ed88ec1461016d578063e82955881461019d578063ec732959146101cd57610091565b80630c394a60146100955780633fd141b6146100b35780634ecf518b146100e357806350e9b92514610101575b5f80fd5b61009d6101eb565b6040516100aa9190610866565b60405180910390f35b6100cd60048036038101906100c89190610a47565b6101f3565b6040516100da9190610abb565b60405180910390f35b6100eb61039e565b6040516100f89190610866565b60405180910390f35b6101096103a3565b6040516101169190610866565b60405180910390f35b61013960048036038101906101349190610ad4565b6103ac565b6040516101469190610866565b60405180910390f35b6101576104ba565b6040516101649190610866565b60405180910390f35b61018760048036038101906101829190610a47565b6104c3565b6040516101949190610866565b60405180910390f35b6101b760048036038101906101b29190610b12565b61066a565b6040516101c49190610866565b60405180910390f35b6101d561082a565b6040516101e29190610866565b60405180910390f35b5f8054905090565b5f80600167ffffffffffffffff8111156102105761020f6108a4565b5b60405190808252806020026020018201604052801561023e5781602001602082028036833780820191505090505b509050835f8151811061025457610253610b3d565b5b6020026020010151815f8151811061026f5761026e610b3d565b5b6020026020010181815250505f73__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49836040518263ffffffff1660e01b81526004016102b39190610c21565b602060405180830381865af41580156102ce573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906102f29190610c55565b90505f600190505b855181101561038e575f6002866103119190610cad565b67ffffffffffffffff160361034b576103448287838151811061033757610336610b3d565b5b60200260200101516103ac565b9150610372565b61036f86828151811061036157610360610b3d565b5b6020026020010151836103ac565b91505b60028561037f9190610cdd565b945080806001019150506102fa565b5060015481149250505092915050565b5f5481565b5f600254905090565b5f80600267ffffffffffffffff8111156103c9576103c86108a4565b5b6040519080825280602002602001820160405280156103f75781602001602082028036833780820191505090505b50905083815f8151811061040e5761040d610b3d565b5b602002602001018181525050828160018151811061042f5761042e610b3d565b5b60200260200101818152505073__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49826040518263ffffffff1660e01b81526004016104729190610c21565b602060405180830381865af415801561048d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104b19190610c55565b91505092915050565b5f600154905090565b5f80600167ffffffffffffffff8111156104e0576104df6108a4565b5b60405190808252806020026020018201604052801561050e5781602001602082028036833780820191505090505b509050835f8151811061052457610523610b3d565b5b6020026020010151815f8151811061053f5761053e610b3d565b5b6020026020010181815250505f73__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49836040518263ffffffff1660e01b81526004016105839190610c21565b602060405180830381865af415801561059e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105c29190610c55565b90505f600190505b855181101561065e575f6002866105e19190610cad565b67ffffffffffffffff160361061b576106148287838151811061060757610606610b3d565b5b60200260200101516103ac565b9150610642565b61063f86828151811061063157610630610b3d565b5b6020026020010151836103ac565b91505b60028561064f9190610cdd565b945080806001019150506105ca565b50809250505092915050565b5f80820361069a577f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd322679050610825565b600182036106ca577f1541c2e47d6e125aea7a107e749299d88fad35bd227179d4fb2bb0f63296d7c79050610825565b600282036106fa577f22d9432cd2f62dedd322bef0f3be1c172e71b6a1afe7866aee3b031ffd89aad79050610825565b6003820361072a577f049889855d82d41769e27ca791a905e528d867095ce49a8ac961ef4722ba6a169050610825565b6004820361075a577f2d1764e964f231d9129b555e1a653a65fca6e20051a0f7a388b5fd15fbb594559050610825565b6005820361078a577f26d2dd833a220c13591b1a528c91e635044b3fac5b46f4133435ff4c87f046949050610825565b600682036107ba577f2c42c8bc5bb243c5fbf697f47b4f41ba191664a8724e5e370090d6d8817930a89050610825565b600782036107ea577f1aaca32cb69540012989bd902c04015555aa164980a490df8c1e6be64fa57fca9050610825565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161081c90610d67565b60405180910390fd5b919050565b7f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd3226781565b5f819050919050565b6108608161084e565b82525050565b5f6020820190506108795f830184610857565b92915050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6108da82610894565b810181811067ffffffffffffffff821117156108f9576108f86108a4565b5b80604052505050565b5f61090b61087f565b905061091782826108d1565b919050565b5f67ffffffffffffffff821115610936576109356108a4565b5b602082029050602081019050919050565b5f80fd5b6109548161084e565b811461095e575f80fd5b50565b5f8135905061096f8161094b565b92915050565b5f6109876109828461091c565b610902565b905080838252602082019050602084028301858111156109aa576109a9610947565b5b835b818110156109d357806109bf8882610961565b8452602084019350506020810190506109ac565b5050509392505050565b5f82601f8301126109f1576109f0610890565b5b8135610a01848260208601610975565b91505092915050565b5f67ffffffffffffffff82169050919050565b610a2681610a0a565b8114610a30575f80fd5b50565b5f81359050610a4181610a1d565b92915050565b5f8060408385031215610a5d57610a5c610888565b5b5f83013567ffffffffffffffff811115610a7a57610a7961088c565b5b610a86858286016109dd565b9250506020610a9785828601610a33565b9150509250929050565b5f8115159050919050565b610ab581610aa1565b82525050565b5f602082019050610ace5f830184610aac565b92915050565b5f8060408385031215610aea57610ae9610888565b5b5f610af785828601610961565b9250506020610b0885828601610961565b9150509250929050565b5f60208284031215610b2757610b26610888565b5b5f610b3484828501610961565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610b9c8161084e565b82525050565b5f610bad8383610b93565b60208301905092915050565b5f602082019050919050565b5f610bcf82610b6a565b610bd98185610b74565b9350610be483610b84565b805f5b83811015610c14578151610bfb8882610ba2565b9750610c0683610bb9565b925050600181019050610be7565b5085935050505092915050565b5f6020820190508181035f830152610c398184610bc5565b905092915050565b5f81519050610c4f8161094b565b92915050565b5f60208284031215610c6a57610c69610888565b5b5f610c7784828501610c41565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f610cb782610a0a565b9150610cc283610a0a565b925082610cd257610cd1610c80565b5b828206905092915050565b5f610ce782610a0a565b9150610cf283610a0a565b925082610d0257610d01610c80565b5b828204905092915050565b5f82825260208201905092915050565b7f696e646578206f7574206f6620626f756e6473000000000000000000000000005f82015250565b5f610d51601383610d0d565b9150610d5c82610d1d565b602082019050919050565b5f6020820190508181035f830152610d7e81610d45565b905091905056fea26469706673582212202cbbc178e6ed4408c4bd08b83d39155eab79694a5fac52cda01f1aa05ebf047264736f6c63430008180033",
}

// MerkleTreeABI is the input ABI used to generate the binding from.
// Deprecated: Use MerkleTreeMetaData.ABI instead.
var MerkleTreeABI = MerkleTreeMetaData.ABI

// MerkleTreeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MerkleTreeMetaData.Bin instead.
var MerkleTreeBin = MerkleTreeMetaData.Bin

// DeployMerkleTree deploys a new Ethereum contract, binding an instance of MerkleTree to it.
func DeployMerkleTree(auth *bind.TransactOpts, backend bind.ContractBackend, _levels *big.Int) (common.Address, *types.Transaction, *MerkleTree, error) {
	parsed, err := MerkleTreeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	miMCAddr, _, _, _ := DeployMiMC(auth, backend)
	MerkleTreeBin = strings.ReplaceAll(MerkleTreeBin, "__$7fbbcdf035b79e739e205dd289b9b9f179$__", miMCAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MerkleTreeBin), backend, _levels)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleTree{MerkleTreeCaller: MerkleTreeCaller{contract: contract}, MerkleTreeTransactor: MerkleTreeTransactor{contract: contract}, MerkleTreeFilterer: MerkleTreeFilterer{contract: contract}}, nil
}

// MerkleTree is an auto generated Go binding around an Ethereum contract.
type MerkleTree struct {
	MerkleTreeCaller     // Read-only binding to the contract
	MerkleTreeTransactor // Write-only binding to the contract
	MerkleTreeFilterer   // Log filterer for contract events
}

// MerkleTreeCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleTreeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleTreeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleTreeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleTreeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleTreeSession struct {
	Contract     *MerkleTree       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MerkleTreeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleTreeCallerSession struct {
	Contract *MerkleTreeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MerkleTreeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleTreeTransactorSession struct {
	Contract     *MerkleTreeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleTreeRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleTreeRaw struct {
	Contract *MerkleTree // Generic contract binding to access the raw methods on
}

// MerkleTreeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleTreeCallerRaw struct {
	Contract *MerkleTreeCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleTreeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleTreeTransactorRaw struct {
	Contract *MerkleTreeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleTree creates a new instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTree(address common.Address, backend bind.ContractBackend) (*MerkleTree, error) {
	contract, err := bindMerkleTree(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleTree{MerkleTreeCaller: MerkleTreeCaller{contract: contract}, MerkleTreeTransactor: MerkleTreeTransactor{contract: contract}, MerkleTreeFilterer: MerkleTreeFilterer{contract: contract}}, nil
}

// NewMerkleTreeCaller creates a new read-only instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTreeCaller(address common.Address, caller bind.ContractCaller) (*MerkleTreeCaller, error) {
	contract, err := bindMerkleTree(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeCaller{contract: contract}, nil
}

// NewMerkleTreeTransactor creates a new write-only instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTreeTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleTreeTransactor, error) {
	contract, err := bindMerkleTree(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeTransactor{contract: contract}, nil
}

// NewMerkleTreeFilterer creates a new log filterer instance of MerkleTree, bound to a specific deployed contract.
func NewMerkleTreeFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleTreeFilterer, error) {
	contract, err := bindMerkleTree(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleTreeFilterer{contract: contract}, nil
}

// bindMerkleTree binds a generic wrapper to an already deployed contract.
func bindMerkleTree(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MerkleTreeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTree *MerkleTreeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTree.Contract.MerkleTreeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTree *MerkleTreeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTree.Contract.MerkleTreeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTree *MerkleTreeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTree.Contract.MerkleTreeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleTree *MerkleTreeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MerkleTree.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleTree *MerkleTreeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleTree.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleTree *MerkleTreeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleTree.Contract.contract.Transact(opts, method, params...)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_MerkleTree *MerkleTreeCaller) ZEROVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTree.contract.Call(opts, &out, "ZERO_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_MerkleTree *MerkleTreeSession) ZEROVALUE() (*big.Int, error) {
	return _MerkleTree.Contract.ZEROVALUE(&_MerkleTree.CallOpts)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_MerkleTree *MerkleTreeCallerSession) ZEROVALUE() (*big.Int, error) {
	return _MerkleTree.Contract.ZEROVALUE(&_MerkleTree.CallOpts)
}

// ComputeMerkleRootFromPath is a free data retrieval call binding the contract method 0x69ed88ec.
//
// Solidity: function computeMerkleRootFromPath(uint256[] path, uint64 index) pure returns(uint256)
func (_MerkleTree *MerkleTreeCaller) ComputeMerkleRootFromPath(opts *bind.CallOpts, path []*big.Int, index uint64) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTree.contract.Call(opts, &out, "computeMerkleRootFromPath", path, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeMerkleRootFromPath is a free data retrieval call binding the contract method 0x69ed88ec.
//
// Solidity: function computeMerkleRootFromPath(uint256[] path, uint64 index) pure returns(uint256)
func (_MerkleTree *MerkleTreeSession) ComputeMerkleRootFromPath(path []*big.Int, index uint64) (*big.Int, error) {
	return _MerkleTree.Contract.ComputeMerkleRootFromPath(&_MerkleTree.CallOpts, path, index)
}

// ComputeMerkleRootFromPath is a free data retrieval call binding the contract method 0x69ed88ec.
//
// Solidity: function computeMerkleRootFromPath(uint256[] path, uint64 index) pure returns(uint256)
func (_MerkleTree *MerkleTreeCallerSession) ComputeMerkleRootFromPath(path []*big.Int, index uint64) (*big.Int, error) {
	return _MerkleTree.Contract.ComputeMerkleRootFromPath(&_MerkleTree.CallOpts, path, index)
}

// GetLevels is a free data retrieval call binding the contract method 0x0c394a60.
//
// Solidity: function getLevels() view returns(uint256)
func (_MerkleTree *MerkleTreeCaller) GetLevels(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTree.contract.Call(opts, &out, "getLevels")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLevels is a free data retrieval call binding the contract method 0x0c394a60.
//
// Solidity: function getLevels() view returns(uint256)
func (_MerkleTree *MerkleTreeSession) GetLevels() (*big.Int, error) {
	return _MerkleTree.Contract.GetLevels(&_MerkleTree.CallOpts)
}

// GetLevels is a free data retrieval call binding the contract method 0x0c394a60.
//
// Solidity: function getLevels() view returns(uint256)
func (_MerkleTree *MerkleTreeCallerSession) GetLevels() (*big.Int, error) {
	return _MerkleTree.Contract.GetLevels(&_MerkleTree.CallOpts)
}

// GetNextLeafIndex is a free data retrieval call binding the contract method 0x50e9b925.
//
// Solidity: function getNextLeafIndex() view returns(uint256)
func (_MerkleTree *MerkleTreeCaller) GetNextLeafIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTree.contract.Call(opts, &out, "getNextLeafIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextLeafIndex is a free data retrieval call binding the contract method 0x50e9b925.
//
// Solidity: function getNextLeafIndex() view returns(uint256)
func (_MerkleTree *MerkleTreeSession) GetNextLeafIndex() (*big.Int, error) {
	return _MerkleTree.Contract.GetNextLeafIndex(&_MerkleTree.CallOpts)
}

// GetNextLeafIndex is a free data retrieval call binding the contract method 0x50e9b925.
//
// Solidity: function getNextLeafIndex() view returns(uint256)
func (_MerkleTree *MerkleTreeCallerSession) GetNextLeafIndex() (*big.Int, error) {
	return _MerkleTree.Contract.GetNextLeafIndex(&_MerkleTree.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(uint256)
func (_MerkleTree *MerkleTreeCaller) GetRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTree.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(uint256)
func (_MerkleTree *MerkleTreeSession) GetRoot() (*big.Int, error) {
	return _MerkleTree.Contract.GetRoot(&_MerkleTree.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(uint256)
func (_MerkleTree *MerkleTreeCallerSession) GetRoot() (*big.Int, error) {
	return _MerkleTree.Contract.GetRoot(&_MerkleTree.CallOpts)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x5bb93995.
//
// Solidity: function hashLeftRight(uint256 left, uint256 right) pure returns(uint256)
func (_MerkleTree *MerkleTreeCaller) HashLeftRight(opts *bind.CallOpts, left *big.Int, right *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTree.contract.Call(opts, &out, "hashLeftRight", left, right)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashLeftRight is a free data retrieval call binding the contract method 0x5bb93995.
//
// Solidity: function hashLeftRight(uint256 left, uint256 right) pure returns(uint256)
func (_MerkleTree *MerkleTreeSession) HashLeftRight(left *big.Int, right *big.Int) (*big.Int, error) {
	return _MerkleTree.Contract.HashLeftRight(&_MerkleTree.CallOpts, left, right)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x5bb93995.
//
// Solidity: function hashLeftRight(uint256 left, uint256 right) pure returns(uint256)
func (_MerkleTree *MerkleTreeCallerSession) HashLeftRight(left *big.Int, right *big.Int) (*big.Int, error) {
	return _MerkleTree.Contract.HashLeftRight(&_MerkleTree.CallOpts, left, right)
}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint256)
func (_MerkleTree *MerkleTreeCaller) Levels(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTree.contract.Call(opts, &out, "levels")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint256)
func (_MerkleTree *MerkleTreeSession) Levels() (*big.Int, error) {
	return _MerkleTree.Contract.Levels(&_MerkleTree.CallOpts)
}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint256)
func (_MerkleTree *MerkleTreeCallerSession) Levels() (*big.Int, error) {
	return _MerkleTree.Contract.Levels(&_MerkleTree.CallOpts)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x3fd141b6.
//
// Solidity: function verifyMerkleProof(uint256[] path, uint64 index) view returns(bool)
func (_MerkleTree *MerkleTreeCaller) VerifyMerkleProof(opts *bind.CallOpts, path []*big.Int, index uint64) (bool, error) {
	var out []interface{}
	err := _MerkleTree.contract.Call(opts, &out, "verifyMerkleProof", path, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x3fd141b6.
//
// Solidity: function verifyMerkleProof(uint256[] path, uint64 index) view returns(bool)
func (_MerkleTree *MerkleTreeSession) VerifyMerkleProof(path []*big.Int, index uint64) (bool, error) {
	return _MerkleTree.Contract.VerifyMerkleProof(&_MerkleTree.CallOpts, path, index)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x3fd141b6.
//
// Solidity: function verifyMerkleProof(uint256[] path, uint64 index) view returns(bool)
func (_MerkleTree *MerkleTreeCallerSession) VerifyMerkleProof(path []*big.Int, index uint64) (bool, error) {
	return _MerkleTree.Contract.VerifyMerkleProof(&_MerkleTree.CallOpts, path, index)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(uint256)
func (_MerkleTree *MerkleTreeCaller) Zeros(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MerkleTree.contract.Call(opts, &out, "zeros", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(uint256)
func (_MerkleTree *MerkleTreeSession) Zeros(i *big.Int) (*big.Int, error) {
	return _MerkleTree.Contract.Zeros(&_MerkleTree.CallOpts, i)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(uint256)
func (_MerkleTree *MerkleTreeCallerSession) Zeros(i *big.Int) (*big.Int, error) {
	return _MerkleTree.Contract.Zeros(&_MerkleTree.CallOpts, i)
}

// MiMCMetaData contains all meta data concerning the MiMC contract.
var MiMCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"in_x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"in_k\",\"type\":\"uint256\"}],\"name\":\"encipher\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"out_x\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"in_msgs\",\"type\":\"uint256[]\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"in_msgs\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"in_key\",\"type\":\"uint256\"}],\"name\":\"hash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x61065261004e600b8282823980515f1a607314610042577f4e487b71000000000000000000000000000000000000000000000000000000005f525f60045260245ffd5b305f52607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061004a575f3560e01c806340ec6e491461004e578063c91286fd1461007e578063ef61dd25146100ae575b5f80fd5b6100686004803603810190610063919061042b565b6100de565b6040516100759190610481565b60405180910390f35b6100986004803603810190610093919061049a565b6100f0565b6040516100a59190610481565b60405180910390f35b6100c860048036038101906100c391906104f4565b610128565b6040516100d59190610481565b60405180910390f35b5f6100e9825f6100f0565b9050919050565b5f61012083837f66a80b61b29ec044d14c4c8c613e762ba1fb8eeb0c454d1ee00ed6dedaa5b5c55f1c606e610160565b905092915050565b5f61015883837f66a80b61b29ec044d14c4c8c613e762ba1fb8eeb0c454d1ee00ed6dedaa5b5c55f1c606e610212565b905092915050565b5f808490505f7f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000190505f5b875181101561020457816101bb8983815181106101ab576101aa610532565b5b6020026020010151858989610212565b8983815181106101ce576101cd610532565b5b6020026020010151856101e1919061058c565b6101eb919061058c565b6101f591906105ec565b9250808060010191505061018b565b508192505050949350505050565b5f6001821015610220575f80fd5b604051602081016040528381527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f00000015f80855b5f8111156102845760208520855283898587518d08089250838384099150838385848509099950600181039050610252565b5082888a08945050505050949350505050565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6102f2826102ac565b810181811067ffffffffffffffff82111715610311576103106102bc565b5b80604052505050565b5f610323610297565b905061032f82826102e9565b919050565b5f67ffffffffffffffff82111561034e5761034d6102bc565b5b602082029050602081019050919050565b5f80fd5b5f819050919050565b61037581610363565b811461037f575f80fd5b50565b5f813590506103908161036c565b92915050565b5f6103a86103a384610334565b61031a565b905080838252602082019050602084028301858111156103cb576103ca61035f565b5b835b818110156103f457806103e08882610382565b8452602084019350506020810190506103cd565b5050509392505050565b5f82601f830112610412576104116102a8565b5b8135610422848260208601610396565b91505092915050565b5f602082840312156104405761043f6102a0565b5b5f82013567ffffffffffffffff81111561045d5761045c6102a4565b5b610469848285016103fe565b91505092915050565b61047b81610363565b82525050565b5f6020820190506104945f830184610472565b92915050565b5f80604083850312156104b0576104af6102a0565b5b5f83013567ffffffffffffffff8111156104cd576104cc6102a4565b5b6104d9858286016103fe565b92505060206104ea85828601610382565b9150509250929050565b5f806040838503121561050a576105096102a0565b5b5f61051785828601610382565b925050602061052885828601610382565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61059682610363565b91506105a183610363565b92508282019050808211156105b9576105b861055f565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6105f682610363565b915061060183610363565b925082610611576106106105bf565b5b82820690509291505056fea26469706673582212204bd03f9345fc87e1f3357926cf41c13711340ad03b4542e2ef80e2f1bec8a5a064736f6c63430008180033",
}

// MiMCABI is the input ABI used to generate the binding from.
// Deprecated: Use MiMCMetaData.ABI instead.
var MiMCABI = MiMCMetaData.ABI

// MiMCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MiMCMetaData.Bin instead.
var MiMCBin = MiMCMetaData.Bin

// DeployMiMC deploys a new Ethereum contract, binding an instance of MiMC to it.
func DeployMiMC(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MiMC, error) {
	parsed, err := MiMCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MiMCBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MiMC{MiMCCaller: MiMCCaller{contract: contract}, MiMCTransactor: MiMCTransactor{contract: contract}, MiMCFilterer: MiMCFilterer{contract: contract}}, nil
}

// MiMC is an auto generated Go binding around an Ethereum contract.
type MiMC struct {
	MiMCCaller     // Read-only binding to the contract
	MiMCTransactor // Write-only binding to the contract
	MiMCFilterer   // Log filterer for contract events
}

// MiMCCaller is an auto generated read-only Go binding around an Ethereum contract.
type MiMCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiMCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MiMCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiMCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MiMCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MiMCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MiMCSession struct {
	Contract     *MiMC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiMCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MiMCCallerSession struct {
	Contract *MiMCCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MiMCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MiMCTransactorSession struct {
	Contract     *MiMCTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MiMCRaw is an auto generated low-level Go binding around an Ethereum contract.
type MiMCRaw struct {
	Contract *MiMC // Generic contract binding to access the raw methods on
}

// MiMCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MiMCCallerRaw struct {
	Contract *MiMCCaller // Generic read-only contract binding to access the raw methods on
}

// MiMCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MiMCTransactorRaw struct {
	Contract *MiMCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMiMC creates a new instance of MiMC, bound to a specific deployed contract.
func NewMiMC(address common.Address, backend bind.ContractBackend) (*MiMC, error) {
	contract, err := bindMiMC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MiMC{MiMCCaller: MiMCCaller{contract: contract}, MiMCTransactor: MiMCTransactor{contract: contract}, MiMCFilterer: MiMCFilterer{contract: contract}}, nil
}

// NewMiMCCaller creates a new read-only instance of MiMC, bound to a specific deployed contract.
func NewMiMCCaller(address common.Address, caller bind.ContractCaller) (*MiMCCaller, error) {
	contract, err := bindMiMC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MiMCCaller{contract: contract}, nil
}

// NewMiMCTransactor creates a new write-only instance of MiMC, bound to a specific deployed contract.
func NewMiMCTransactor(address common.Address, transactor bind.ContractTransactor) (*MiMCTransactor, error) {
	contract, err := bindMiMC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MiMCTransactor{contract: contract}, nil
}

// NewMiMCFilterer creates a new log filterer instance of MiMC, bound to a specific deployed contract.
func NewMiMCFilterer(address common.Address, filterer bind.ContractFilterer) (*MiMCFilterer, error) {
	contract, err := bindMiMC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MiMCFilterer{contract: contract}, nil
}

// bindMiMC binds a generic wrapper to an already deployed contract.
func bindMiMC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MiMCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiMC *MiMCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MiMC.Contract.MiMCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiMC *MiMCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiMC.Contract.MiMCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiMC *MiMCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiMC.Contract.MiMCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MiMC *MiMCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MiMC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MiMC *MiMCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MiMC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MiMC *MiMCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MiMC.Contract.contract.Transact(opts, method, params...)
}

// Encipher is a free data retrieval call binding the contract method 0xef61dd25.
//
// Solidity: function encipher(uint256 in_x, uint256 in_k) pure returns(uint256 out_x)
func (_MiMC *MiMCCaller) Encipher(opts *bind.CallOpts, in_x *big.Int, in_k *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MiMC.contract.Call(opts, &out, "encipher", in_x, in_k)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Encipher is a free data retrieval call binding the contract method 0xef61dd25.
//
// Solidity: function encipher(uint256 in_x, uint256 in_k) pure returns(uint256 out_x)
func (_MiMC *MiMCSession) Encipher(in_x *big.Int, in_k *big.Int) (*big.Int, error) {
	return _MiMC.Contract.Encipher(&_MiMC.CallOpts, in_x, in_k)
}

// Encipher is a free data retrieval call binding the contract method 0xef61dd25.
//
// Solidity: function encipher(uint256 in_x, uint256 in_k) pure returns(uint256 out_x)
func (_MiMC *MiMCCallerSession) Encipher(in_x *big.Int, in_k *big.Int) (*big.Int, error) {
	return _MiMC.Contract.Encipher(&_MiMC.CallOpts, in_x, in_k)
}

// Hash is a free data retrieval call binding the contract method 0x40ec6e49.
//
// Solidity: function hash(uint256[] in_msgs) pure returns(uint256)
func (_MiMC *MiMCCaller) Hash(opts *bind.CallOpts, in_msgs []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MiMC.contract.Call(opts, &out, "hash", in_msgs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hash is a free data retrieval call binding the contract method 0x40ec6e49.
//
// Solidity: function hash(uint256[] in_msgs) pure returns(uint256)
func (_MiMC *MiMCSession) Hash(in_msgs []*big.Int) (*big.Int, error) {
	return _MiMC.Contract.Hash(&_MiMC.CallOpts, in_msgs)
}

// Hash is a free data retrieval call binding the contract method 0x40ec6e49.
//
// Solidity: function hash(uint256[] in_msgs) pure returns(uint256)
func (_MiMC *MiMCCallerSession) Hash(in_msgs []*big.Int) (*big.Int, error) {
	return _MiMC.Contract.Hash(&_MiMC.CallOpts, in_msgs)
}

// Hash0 is a free data retrieval call binding the contract method 0xc91286fd.
//
// Solidity: function hash(uint256[] in_msgs, uint256 in_key) pure returns(uint256)
func (_MiMC *MiMCCaller) Hash0(opts *bind.CallOpts, in_msgs []*big.Int, in_key *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _MiMC.contract.Call(opts, &out, "hash0", in_msgs, in_key)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Hash0 is a free data retrieval call binding the contract method 0xc91286fd.
//
// Solidity: function hash(uint256[] in_msgs, uint256 in_key) pure returns(uint256)
func (_MiMC *MiMCSession) Hash0(in_msgs []*big.Int, in_key *big.Int) (*big.Int, error) {
	return _MiMC.Contract.Hash0(&_MiMC.CallOpts, in_msgs, in_key)
}

// Hash0 is a free data retrieval call binding the contract method 0xc91286fd.
//
// Solidity: function hash(uint256[] in_msgs, uint256 in_key) pure returns(uint256)
func (_MiMC *MiMCCallerSession) Hash0(in_msgs []*big.Int, in_key *big.Int) (*big.Int, error) {
	return _MiMC.Contract.Hash0(&_MiMC.CallOpts, in_msgs, in_key)
}

// RollupMetaData contains all meta data concerning the Rollup contract.
var RollupMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifierAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_treeLevels\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initRoot\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMerkleRoot\",\"type\":\"uint256\"}],\"name\":\"MerkleRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structRollup.PublicKey\",\"name\":\"pubkey\",\"type\":\"tuple\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TransactionData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"RollupToEth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"computeMerkleRootFromPath\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rollupAccountPKX\",\"type\":\"uint256\"}],\"name\":\"getEthAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLevels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextLeafIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structRollup.PublicKey\",\"name\":\"pubKey\",\"type\":\"tuple\"}],\"internalType\":\"structRollup.Account\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"hashAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"}],\"name\":\"hashLeftRight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"levels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"merkleProof\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ganacheAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structRollup.PublicKey\",\"name\":\"pubKey\",\"type\":\"tuple\"}],\"internalType\":\"structRollup.Account\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"processDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"zkproof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256\",\"name\":\"postStateRoot\",\"type\":\"uint256\"}],\"name\":\"submitVerifyBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"verifyZkProof\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"merkleProof\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structRollup.PublicKey\",\"name\":\"pubKey\",\"type\":\"tuple\"}],\"internalType\":\"structRollup.Account\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"zeros\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b5060405162002c1138038062002c118339818101604052810190620000369190620002a6565b815f81116200007c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000739062000383565b60405180910390fd5b60098110620000c2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000b990620003f1565b60405180910390fd5b6004811462000108576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000ff9062000485565b60405180910390fd5b805f819055507f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd3226760018190555050815f8190555082600760086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503360045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620001ce81620001ff60201b60201c565b5f60075f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050620004a5565b8060018190555050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000238826200020d565b9050919050565b6200024a816200022c565b811462000255575f80fd5b50565b5f8151905062000268816200023f565b92915050565b5f819050919050565b62000282816200026e565b81146200028d575f80fd5b50565b5f81519050620002a08162000277565b92915050565b5f805f60608486031215620002c057620002bf62000209565b5b5f620002cf8682870162000258565b9350506020620002e28682870162000290565b9250506040620002f58682870162000290565b9150509250925092565b5f82825260208201905092915050565b7f5f6c6576656c732073686f756c642062652067726561746572207468616e207a5f8201527f65726f0000000000000000000000000000000000000000000000000000000000602082015250565b5f6200036b602383620002ff565b915062000378826200030f565b604082019050919050565b5f6020820190508181035f8301526200039c816200035d565b9050919050565b7f5f6c6576656c732073686f756c64206265206c657373207468616e20380000005f82015250565b5f620003d9601d83620002ff565b9150620003e682620003a3565b602082019050919050565b5f6020820190508181035f8301526200040a81620003cb565b9050919050565b7f5f6c6576656c732073686f756c6420626520342028325e34203d2031362061635f8201527f636f756e74732920666f722074657374696e6700000000000000000000000000602082015250565b5f6200046d603383620002ff565b91506200047a8262000411565b604082019050919050565b5f6020820190508181035f8301526200049e816200045f565b9050919050565b61275e80620004b35f395ff3fe60806040526004361061011e575f3560e01c806369ed88ec1161009f578063bd4986a011610063578063bd4986a0146103fe578063cc0d802e1461043a578063d0e30db014610462578063e82955881461046c578063ec732959146104a85761011e565b806369ed88ec146102f857806383290a3c146103345780638da5cb5b14610370578063909ad5ba1461039a578063b994f687146103d65761011e565b806349590657116100e657806349590657146102145780634ecf518b1461023e57806350e9b925146102685780635bb93995146102925780635ca1e165146102ce5761011e565b806302722d20146101225780630c394a601461014a5780631174dc99146101745780631ea5f6eb1461019c5780633fd141b6146101d8575b5f80fd5b34801561012d575f80fd5b5061014860048036038101906101439190611bf0565b6104d2565b005b348015610155575f80fd5b5061015e6107fc565b60405161016b9190611c6b565b60405180910390f35b34801561017f575f80fd5b5061019a60048036038101906101959190611cde565b610804565b005b3480156101a7575f80fd5b506101c260048036038101906101bd9190611d5f565b610b84565b6040516101cf9190611d99565b60405180910390f35b3480156101e3575f80fd5b506101fe60048036038101906101f99190611db2565b610bb4565b60405161020b9190611e26565b60405180910390f35b34801561021f575f80fd5b50610228610d5f565b6040516102359190611c6b565b60405180910390f35b348015610249575f80fd5b50610252610d6d565b60405161025f9190611c6b565b60405180910390f35b348015610273575f80fd5b5061027c610d72565b6040516102899190611c6b565b60405180910390f35b34801561029d575f80fd5b506102b860048036038101906102b39190611e3f565b610d7b565b6040516102c59190611c6b565b60405180910390f35b3480156102d9575f80fd5b506102e2610e89565b6040516102ef9190611c6b565b60405180910390f35b348015610303575f80fd5b5061031e60048036038101906103199190611db2565b610e92565b60405161032b9190611c6b565b60405180910390f35b34801561033f575f80fd5b5061035a60048036038101906103559190611d5f565b611039565b6040516103679190611c6b565b60405180910390f35b34801561037b575f80fd5b50610384611059565b6040516103919190611d99565b60405180910390f35b3480156103a5575f80fd5b506103c060048036038101906103bb9190611e7d565b61107e565b6040516103cd9190611c6b565b60405180910390f35b3480156103e1575f80fd5b506103fc60048036038101906103f79190611ec9565b611212565b005b348015610409575f80fd5b50610424600480360381019061041f9190611d5f565b6113be565b6040516104319190611d99565b60405180910390f35b348015610445575f80fd5b50610460600480360381019061045b9190611f2a565b6113f7565b005b61046a611572565b005b348015610477575f80fd5b50610492600480360381019061048d9190611d5f565b611749565b60405161049f9190611c6b565b60405180910390f35b3480156104b3575f80fd5b506104bc611909565b6040516104c99190611c6b565b60405180910390f35b5f60055f83606001515f015181526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff160361057b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161057290611fc4565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e09061202c565b60405180910390fd5b5f6105f38361107e565b905080845f815181106106095761060861204a565b5b60200260200101818152505061062284845f0151610bb4565b610661576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610658906120c1565b60405180910390fd5b84836040015110156106a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069f90612129565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc8690811502906040515f60405180830381858888f193505050501580156106eb573d5f803e3d5ffd5b508173ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364866040516107329190611c6b565b60405180910390a2848360400181815161074c9190612174565b9150818152505084836040015161076391906121a7565b83604001511115610777576107766121da565b5b5f6107818461107e565b905080855f815181106107975761079661204a565b5b6020026020010181815250505f6107b186865f0151610e92565b90506107bc8161192d565b7f3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a816040516107eb9190611c6b565b60405180910390a150505050505050565b5f8054905090565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610893576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161088a90612251565b60405180910390fd5b6108a084825f0151610bb4565b6108df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108d6906122b9565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660055f83606001515f015181526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036109ea577f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd32267845f8151811061097d5761097c61204a565b5b602002602001015114610993576109926121da565b5b8160055f83606001515f015181526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600660075f9054906101000a900467ffffffffffffffff1667ffffffffffffffff1681548110610a1d57610a1c61204a565b5b905f5260205f200154610a46848473ffffffffffffffffffffffffffffffffffffffff16610d7b565b14610a86576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a7d90612321565b60405180910390fd5b8281604001818151610a9891906121a7565b915081815250508281604001511015610ab457610ab36121da565b5b5f610abe8261107e565b905080855f81518110610ad457610ad361204a565b5b6020026020010181815250505f610aee86845f0151610e92565b9050610af98161192d565b7f3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a81604051610b289190611c6b565b60405180910390a160075f81819054906101000a900467ffffffffffffffff1680929190610b559061233f565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050505050565b6005602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f80600167ffffffffffffffff811115610bd157610bd06119b1565b5b604051908082528060200260200182016040528015610bff5781602001602082028036833780820191505090505b509050835f81518110610c1557610c1461204a565b5b6020026020010151815f81518110610c3057610c2f61204a565b5b6020026020010181815250505f73__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49836040518263ffffffff1660e01b8152600401610c749190612425565b602060405180830381865af4158015610c8f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610cb39190612459565b90505f600190505b8551811015610d4f575f600286610cd291906124b1565b67ffffffffffffffff1603610d0c57610d0582878381518110610cf857610cf761204a565b5b6020026020010151610d7b565b9150610d33565b610d30868281518110610d2257610d2161204a565b5b602002602001015183610d7b565b91505b600285610d4091906124e1565b94508080600101915050610cbb565b5060015481149250505092915050565b5f610d68610e89565b905090565b5f5481565b5f600254905090565b5f80600267ffffffffffffffff811115610d9857610d976119b1565b5b604051908082528060200260200182016040528015610dc65781602001602082028036833780820191505090505b50905083815f81518110610ddd57610ddc61204a565b5b6020026020010181815250508281600181518110610dfe57610dfd61204a565b5b60200260200101818152505073__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49826040518263ffffffff1660e01b8152600401610e419190612425565b602060405180830381865af4158015610e5c573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e809190612459565b91505092915050565b5f600154905090565b5f80600167ffffffffffffffff811115610eaf57610eae6119b1565b5b604051908082528060200260200182016040528015610edd5781602001602082028036833780820191505090505b509050835f81518110610ef357610ef261204a565b5b6020026020010151815f81518110610f0e57610f0d61204a565b5b6020026020010181815250505f73__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49836040518263ffffffff1660e01b8152600401610f529190612425565b602060405180830381865af4158015610f6d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f919190612459565b90505f600190505b855181101561102d575f600286610fb091906124b1565b67ffffffffffffffff1603610fea57610fe382878381518110610fd657610fd561204a565b5b6020026020010151610d7b565b9150611011565b61100e86828151811061100057610fff61204a565b5b602002602001015183610d7b565b91505b60028561101e91906124e1565b94508080600101915050610f99565b50809250505092915050565b60068181548110611048575f80fd5b905f5260205f20015f915090505481565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f80600567ffffffffffffffff81111561109b5761109a6119b1565b5b6040519080825280602002602001820160405280156110c95781602001602082028036833780820191505090505b509050825f015167ffffffffffffffff16815f815181106110ed576110ec61204a565b5b6020026020010181815250508260200151816001815181106111125761111161204a565b5b6020026020010181815250508260400151816002815181106111375761113661204a565b5b60200260200101818152505082606001515f01518160038151811061115f5761115e61204a565b5b602002602001018181525050826060015160200151816004815181106111885761118761204a565b5b60200260200101818152505073__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49826040518263ffffffff1660e01b81526004016111cb9190612425565b602060405180830381865af41580156111e6573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061120a9190612459565b915050919050565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146112a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161129890612251565b60405180910390fd5b6112a9611937565b5f6112b2610e89565b905080825f600281106112c8576112c761204a565b5b60200201818152505082826001600281106112e6576112e561204a565b5b602002018181525050600760089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635fe24f2385846040518363ffffffff1660e01b815260040161134c9291906125d0565b5f6040518083038186803b158015611362575f80fd5b505afa158015611374573d5f803e3d5ffd5b505050506113818361192d565b7f3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a836040516113b09190611c6b565b60405180910390a150505050565b5f60055f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611486576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161147d90612251565b60405180910390fd5b600760089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16635fe24f2383836040518363ffffffff1660e01b81526004016114e3929190612609565b5f6040518083038186803b1580156114f9575f80fd5b505afa15801561150b573d5f803e3d5ffd5b50505050611530816001600281106115265761152561204a565b5b602002013561192d565b7f3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a611559610e89565b6040516115669190611c6b565b60405180910390a15050565b5f34116115b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115ab906126a2565b60405180910390fd5b5f600267ffffffffffffffff8111156115d0576115cf6119b1565b5b6040519080825280602002602001820160405280156115fe5781602001602082028036833780820191505090505b5090503373ffffffffffffffffffffffffffffffffffffffff16815f8151811061162b5761162a61204a565b5b602002602001018181525050348160018151811061164c5761164b61204a565b5b6020026020010181815250505f73__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49836040518263ffffffff1660e01b81526004016116909190612425565b602060405180830381865af41580156116ab573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116cf9190612459565b9050600681908060018154018082558091505060019003905f5260205f20015f90919091909150553373ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c3460405161173d9190611c6b565b60405180910390a25050565b5f808203611779577f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd322679050611904565b600182036117a9577f1541c2e47d6e125aea7a107e749299d88fad35bd227179d4fb2bb0f63296d7c79050611904565b600282036117d9577f22d9432cd2f62dedd322bef0f3be1c172e71b6a1afe7866aee3b031ffd89aad79050611904565b60038203611809577f049889855d82d41769e27ca791a905e528d867095ce49a8ac961ef4722ba6a169050611904565b60048203611839577f2d1764e964f231d9129b555e1a653a65fca6e20051a0f7a388b5fd15fbb594559050611904565b60058203611869577f26d2dd833a220c13591b1a528c91e635044b3fac5b46f4133435ff4c87f046949050611904565b60068203611899577f2c42c8bc5bb243c5fbf697f47b4f41ba191664a8724e5e370090d6d8817930a89050611904565b600782036118c9577f1aaca32cb69540012989bd902c04015555aa164980a490df8c1e6be64fa57fca9050611904565b6040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016118fb9061270a565b60405180910390fd5b919050565b7f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd3226781565b8060018190555050565b6040518060400160405280600290602082028036833780820191505090505090565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b61197c8161196a565b8114611986575f80fd5b50565b5f8135905061199781611973565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6119e7826119a1565b810181811067ffffffffffffffff82111715611a0657611a056119b1565b5b80604052505050565b5f611a18611959565b9050611a2482826119de565b919050565b5f67ffffffffffffffff821115611a4357611a426119b1565b5b602082029050602081019050919050565b5f80fd5b5f611a6a611a6584611a29565b611a0f565b90508083825260208201905060208402830185811115611a8d57611a8c611a54565b5b835b81811015611ab65780611aa28882611989565b845260208401935050602081019050611a8f565b5050509392505050565b5f82601f830112611ad457611ad361199d565b5b8135611ae4848260208601611a58565b91505092915050565b5f80fd5b5f67ffffffffffffffff82169050919050565b611b0d81611af1565b8114611b17575f80fd5b50565b5f81359050611b2881611b04565b92915050565b5f60408284031215611b4357611b42611aed565b5b611b4d6040611a0f565b90505f611b5c84828501611989565b5f830152506020611b6f84828501611989565b60208301525092915050565b5f60a08284031215611b9057611b8f611aed565b5b611b9a6080611a0f565b90505f611ba984828501611b1a565b5f830152506020611bbc84828501611989565b6020830152506040611bd084828501611989565b6040830152506060611be484828501611b2e565b60608301525092915050565b5f805f60e08486031215611c0757611c06611962565b5b5f611c1486828701611989565b935050602084013567ffffffffffffffff811115611c3557611c34611966565b5b611c4186828701611ac0565b9250506040611c5286828701611b7b565b9150509250925092565b611c658161196a565b82525050565b5f602082019050611c7e5f830184611c5c565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611cad82611c84565b9050919050565b611cbd81611ca3565b8114611cc7575f80fd5b50565b5f81359050611cd881611cb4565b92915050565b5f805f806101008587031215611cf757611cf6611962565b5b5f85013567ffffffffffffffff811115611d1457611d13611966565b5b611d2087828801611ac0565b9450506020611d3187828801611989565b9350506040611d4287828801611cca565b9250506060611d5387828801611b7b565b91505092959194509250565b5f60208284031215611d7457611d73611962565b5b5f611d8184828501611989565b91505092915050565b611d9381611ca3565b82525050565b5f602082019050611dac5f830184611d8a565b92915050565b5f8060408385031215611dc857611dc7611962565b5b5f83013567ffffffffffffffff811115611de557611de4611966565b5b611df185828601611ac0565b9250506020611e0285828601611b1a565b9150509250929050565b5f8115159050919050565b611e2081611e0c565b82525050565b5f602082019050611e395f830184611e17565b92915050565b5f8060408385031215611e5557611e54611962565b5b5f611e6285828601611989565b9250506020611e7385828601611989565b9150509250929050565b5f60a08284031215611e9257611e91611962565b5b5f611e9f84828501611b7b565b91505092915050565b5f81905082602060080282011115611ec357611ec2611a54565b5b92915050565b5f806101208385031215611ee057611edf611962565b5b5f611eed85828601611ea8565b925050610100611eff85828601611989565b9150509250929050565b5f81905082602060020282011115611f2457611f23611a54565b5b92915050565b5f806101408385031215611f4157611f40611962565b5b5f611f4e85828601611ea8565b925050610100611f6085828601611f09565b9150509250929050565b5f82825260208201905092915050565b7f526f6c6c7570206164647265737320646f6573206e6f742065786973740000005f82015250565b5f611fae601d83611f6a565b9150611fb982611f7a565b602082019050919050565b5f6020820190508181035f830152611fdb81611fa2565b9050919050565b7f4e6f7420726967687466756c206163636f756e742061646472657373000000005f82015250565b5f612016601c83611f6a565b915061202182611fe2565b602082019050919050565b5f6020820190508181035f8301526120438161200a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f496e76616c6964204d65726b6c652070726f6f660000000000000000000000005f82015250565b5f6120ab601483611f6a565b91506120b682612077565b602082019050919050565b5f6020820190508181035f8301526120d88161209f565b9050919050565b7f42616c616e6365206e6f742073756666696369656e74000000000000000000005f82015250565b5f612113601683611f6a565b915061211e826120df565b602082019050919050565b5f6020820190508181035f83015261214081612107565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61217e8261196a565b91506121898361196a565b92508282039050818111156121a1576121a0612147565b5b92915050565b5f6121b18261196a565b91506121bc8361196a565b92508282019050808211156121d4576121d3612147565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffd5b7f4f6e6c79206f776e65722063616e207375626d697420626174636865730000005f82015250565b5f61223b601d83611f6a565b915061224682612207565b602082019050919050565b5f6020820190508181035f8301526122688161222f565b9050919050565b7f4d65726b6c652070726f6f6620696e76616c69640000000000000000000000005f82015250565b5f6122a3601483611f6a565b91506122ae8261226f565b602082019050919050565b5f6020820190508181035f8301526122d081612297565b9050919050565b7f4465706f736974206e6f7420636f7272656374206f6e6520696e2071756575655f82015250565b5f61230b602083611f6a565b9150612316826122d7565b602082019050919050565b5f6020820190508181035f830152612338816122ff565b9050919050565b5f61234982611af1565b915067ffffffffffffffff820361236357612362612147565b5b600182019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b6123a08161196a565b82525050565b5f6123b18383612397565b60208301905092915050565b5f602082019050919050565b5f6123d38261236e565b6123dd8185612378565b93506123e883612388565b805f5b838110156124185781516123ff88826123a6565b975061240a836123bd565b9250506001810190506123eb565b5085935050505092915050565b5f6020820190508181035f83015261243d81846123c9565b905092915050565b5f8151905061245381611973565b92915050565b5f6020828403121561246e5761246d611962565b5b5f61247b84828501612445565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6124bb82611af1565b91506124c683611af1565b9250826124d6576124d5612484565b5b828206905092915050565b5f6124eb82611af1565b91506124f683611af1565b92508261250657612505612484565b5b828204905092915050565b82818337505050565b6125276101008383612511565b5050565b5f60029050919050565b5f81905092915050565b5f819050919050565b6125518161196a565b82525050565b5f6125628383612548565b60208301905092915050565b5f602082019050919050565b6125838161252b565b61258d8184612535565b92506125988261253f565b805f5b838110156125c85781516125af8782612557565b96506125ba8361256e565b92505060018101905061259b565b505050505050565b5f610140820190506125e45f83018561251a565b6125f261010083018461257a565b9392505050565b61260560408383612511565b5050565b5f6101408201905061261d5f83018561251a565b61262b6101008301846125f9565b9392505050565b7f4465706f7369742076616c7565206d75737420626520677265617465722074685f8201527f616e203000000000000000000000000000000000000000000000000000000000602082015250565b5f61268c602483611f6a565b915061269782612632565b604082019050919050565b5f6020820190508181035f8301526126b981612680565b9050919050565b7f696e646578206f7574206f6620626f756e6473000000000000000000000000005f82015250565b5f6126f4601383611f6a565b91506126ff826126c0565b602082019050919050565b5f6020820190508181035f830152612721816126e8565b905091905056fea26469706673582212202fb39a8626f79c335722e8ddd8ca26e14449cb0315e70d75293795ed42269cf064736f6c63430008180033",
}

// RollupABI is the input ABI used to generate the binding from.
// Deprecated: Use RollupMetaData.ABI instead.
var RollupABI = RollupMetaData.ABI

// RollupBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RollupMetaData.Bin instead.
var RollupBin = RollupMetaData.Bin

// DeployRollup deploys a new Ethereum contract, binding an instance of Rollup to it.
func DeployRollup(auth *bind.TransactOpts, backend bind.ContractBackend, _verifierAddress common.Address, _treeLevels *big.Int, initRoot *big.Int) (common.Address, *types.Transaction, *Rollup, error) {
	parsed, err := RollupMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	miMCAddr, _, _, _ := DeployMiMC(auth, backend)
	RollupBin = strings.ReplaceAll(RollupBin, "__$7fbbcdf035b79e739e205dd289b9b9f179$__", miMCAddr.String()[2:])

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RollupBin), backend, _verifierAddress, _treeLevels, initRoot)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// Rollup is an auto generated Go binding around an Ethereum contract.
type Rollup struct {
	RollupCaller     // Read-only binding to the contract
	RollupTransactor // Write-only binding to the contract
	RollupFilterer   // Log filterer for contract events
}

// RollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupSession struct {
	Contract     *Rollup           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCallerSession struct {
	Contract *RollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTransactorSession struct {
	Contract     *RollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupRaw struct {
	Contract *Rollup // Generic contract binding to access the raw methods on
}

// RollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCallerRaw struct {
	Contract *RollupCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTransactorRaw struct {
	Contract *RollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollup creates a new instance of Rollup, bound to a specific deployed contract.
func NewRollup(address common.Address, backend bind.ContractBackend) (*Rollup, error) {
	contract, err := bindRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// NewRollupCaller creates a new read-only instance of Rollup, bound to a specific deployed contract.
func NewRollupCaller(address common.Address, caller bind.ContractCaller) (*RollupCaller, error) {
	contract, err := bindRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCaller{contract: contract}, nil
}

// NewRollupTransactor creates a new write-only instance of Rollup, bound to a specific deployed contract.
func NewRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTransactor, error) {
	contract, err := bindRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTransactor{contract: contract}, nil
}

// NewRollupFilterer creates a new log filterer instance of Rollup, bound to a specific deployed contract.
func NewRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupFilterer, error) {
	contract, err := bindRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupFilterer{contract: contract}, nil
}

// bindRollup binds a generic wrapper to an already deployed contract.
func bindRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RollupMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.RollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transact(opts, method, params...)
}

// RollupToEth is a free data retrieval call binding the contract method 0x1ea5f6eb.
//
// Solidity: function RollupToEth(uint256 ) view returns(address)
func (_Rollup *RollupCaller) RollupToEth(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "RollupToEth", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RollupToEth is a free data retrieval call binding the contract method 0x1ea5f6eb.
//
// Solidity: function RollupToEth(uint256 ) view returns(address)
func (_Rollup *RollupSession) RollupToEth(arg0 *big.Int) (common.Address, error) {
	return _Rollup.Contract.RollupToEth(&_Rollup.CallOpts, arg0)
}

// RollupToEth is a free data retrieval call binding the contract method 0x1ea5f6eb.
//
// Solidity: function RollupToEth(uint256 ) view returns(address)
func (_Rollup *RollupCallerSession) RollupToEth(arg0 *big.Int) (common.Address, error) {
	return _Rollup.Contract.RollupToEth(&_Rollup.CallOpts, arg0)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_Rollup *RollupCaller) ZEROVALUE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "ZERO_VALUE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_Rollup *RollupSession) ZEROVALUE() (*big.Int, error) {
	return _Rollup.Contract.ZEROVALUE(&_Rollup.CallOpts)
}

// ZEROVALUE is a free data retrieval call binding the contract method 0xec732959.
//
// Solidity: function ZERO_VALUE() view returns(uint256)
func (_Rollup *RollupCallerSession) ZEROVALUE() (*big.Int, error) {
	return _Rollup.Contract.ZEROVALUE(&_Rollup.CallOpts)
}

// ComputeMerkleRootFromPath is a free data retrieval call binding the contract method 0x69ed88ec.
//
// Solidity: function computeMerkleRootFromPath(uint256[] path, uint64 index) pure returns(uint256)
func (_Rollup *RollupCaller) ComputeMerkleRootFromPath(opts *bind.CallOpts, path []*big.Int, index uint64) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "computeMerkleRootFromPath", path, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeMerkleRootFromPath is a free data retrieval call binding the contract method 0x69ed88ec.
//
// Solidity: function computeMerkleRootFromPath(uint256[] path, uint64 index) pure returns(uint256)
func (_Rollup *RollupSession) ComputeMerkleRootFromPath(path []*big.Int, index uint64) (*big.Int, error) {
	return _Rollup.Contract.ComputeMerkleRootFromPath(&_Rollup.CallOpts, path, index)
}

// ComputeMerkleRootFromPath is a free data retrieval call binding the contract method 0x69ed88ec.
//
// Solidity: function computeMerkleRootFromPath(uint256[] path, uint64 index) pure returns(uint256)
func (_Rollup *RollupCallerSession) ComputeMerkleRootFromPath(path []*big.Int, index uint64) (*big.Int, error) {
	return _Rollup.Contract.ComputeMerkleRootFromPath(&_Rollup.CallOpts, path, index)
}

// DepositQueue is a free data retrieval call binding the contract method 0x83290a3c.
//
// Solidity: function depositQueue(uint256 ) view returns(uint256)
func (_Rollup *RollupCaller) DepositQueue(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "depositQueue", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositQueue is a free data retrieval call binding the contract method 0x83290a3c.
//
// Solidity: function depositQueue(uint256 ) view returns(uint256)
func (_Rollup *RollupSession) DepositQueue(arg0 *big.Int) (*big.Int, error) {
	return _Rollup.Contract.DepositQueue(&_Rollup.CallOpts, arg0)
}

// DepositQueue is a free data retrieval call binding the contract method 0x83290a3c.
//
// Solidity: function depositQueue(uint256 ) view returns(uint256)
func (_Rollup *RollupCallerSession) DepositQueue(arg0 *big.Int) (*big.Int, error) {
	return _Rollup.Contract.DepositQueue(&_Rollup.CallOpts, arg0)
}

// GetEthAddress is a free data retrieval call binding the contract method 0xbd4986a0.
//
// Solidity: function getEthAddress(uint256 rollupAccountPKX) view returns(address)
func (_Rollup *RollupCaller) GetEthAddress(opts *bind.CallOpts, rollupAccountPKX *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getEthAddress", rollupAccountPKX)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEthAddress is a free data retrieval call binding the contract method 0xbd4986a0.
//
// Solidity: function getEthAddress(uint256 rollupAccountPKX) view returns(address)
func (_Rollup *RollupSession) GetEthAddress(rollupAccountPKX *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetEthAddress(&_Rollup.CallOpts, rollupAccountPKX)
}

// GetEthAddress is a free data retrieval call binding the contract method 0xbd4986a0.
//
// Solidity: function getEthAddress(uint256 rollupAccountPKX) view returns(address)
func (_Rollup *RollupCallerSession) GetEthAddress(rollupAccountPKX *big.Int) (common.Address, error) {
	return _Rollup.Contract.GetEthAddress(&_Rollup.CallOpts, rollupAccountPKX)
}

// GetLevels is a free data retrieval call binding the contract method 0x0c394a60.
//
// Solidity: function getLevels() view returns(uint256)
func (_Rollup *RollupCaller) GetLevels(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getLevels")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLevels is a free data retrieval call binding the contract method 0x0c394a60.
//
// Solidity: function getLevels() view returns(uint256)
func (_Rollup *RollupSession) GetLevels() (*big.Int, error) {
	return _Rollup.Contract.GetLevels(&_Rollup.CallOpts)
}

// GetLevels is a free data retrieval call binding the contract method 0x0c394a60.
//
// Solidity: function getLevels() view returns(uint256)
func (_Rollup *RollupCallerSession) GetLevels() (*big.Int, error) {
	return _Rollup.Contract.GetLevels(&_Rollup.CallOpts)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(uint256)
func (_Rollup *RollupCaller) GetMerkleRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getMerkleRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(uint256)
func (_Rollup *RollupSession) GetMerkleRoot() (*big.Int, error) {
	return _Rollup.Contract.GetMerkleRoot(&_Rollup.CallOpts)
}

// GetMerkleRoot is a free data retrieval call binding the contract method 0x49590657.
//
// Solidity: function getMerkleRoot() view returns(uint256)
func (_Rollup *RollupCallerSession) GetMerkleRoot() (*big.Int, error) {
	return _Rollup.Contract.GetMerkleRoot(&_Rollup.CallOpts)
}

// GetNextLeafIndex is a free data retrieval call binding the contract method 0x50e9b925.
//
// Solidity: function getNextLeafIndex() view returns(uint256)
func (_Rollup *RollupCaller) GetNextLeafIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getNextLeafIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextLeafIndex is a free data retrieval call binding the contract method 0x50e9b925.
//
// Solidity: function getNextLeafIndex() view returns(uint256)
func (_Rollup *RollupSession) GetNextLeafIndex() (*big.Int, error) {
	return _Rollup.Contract.GetNextLeafIndex(&_Rollup.CallOpts)
}

// GetNextLeafIndex is a free data retrieval call binding the contract method 0x50e9b925.
//
// Solidity: function getNextLeafIndex() view returns(uint256)
func (_Rollup *RollupCallerSession) GetNextLeafIndex() (*big.Int, error) {
	return _Rollup.Contract.GetNextLeafIndex(&_Rollup.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(uint256)
func (_Rollup *RollupCaller) GetRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "getRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(uint256)
func (_Rollup *RollupSession) GetRoot() (*big.Int, error) {
	return _Rollup.Contract.GetRoot(&_Rollup.CallOpts)
}

// GetRoot is a free data retrieval call binding the contract method 0x5ca1e165.
//
// Solidity: function getRoot() view returns(uint256)
func (_Rollup *RollupCallerSession) GetRoot() (*big.Int, error) {
	return _Rollup.Contract.GetRoot(&_Rollup.CallOpts)
}

// HashAccount is a free data retrieval call binding the contract method 0x909ad5ba.
//
// Solidity: function hashAccount((uint64,uint256,uint256,(uint256,uint256)) account) pure returns(uint256)
func (_Rollup *RollupCaller) HashAccount(opts *bind.CallOpts, account RollupAccount) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "hashAccount", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashAccount is a free data retrieval call binding the contract method 0x909ad5ba.
//
// Solidity: function hashAccount((uint64,uint256,uint256,(uint256,uint256)) account) pure returns(uint256)
func (_Rollup *RollupSession) HashAccount(account RollupAccount) (*big.Int, error) {
	return _Rollup.Contract.HashAccount(&_Rollup.CallOpts, account)
}

// HashAccount is a free data retrieval call binding the contract method 0x909ad5ba.
//
// Solidity: function hashAccount((uint64,uint256,uint256,(uint256,uint256)) account) pure returns(uint256)
func (_Rollup *RollupCallerSession) HashAccount(account RollupAccount) (*big.Int, error) {
	return _Rollup.Contract.HashAccount(&_Rollup.CallOpts, account)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x5bb93995.
//
// Solidity: function hashLeftRight(uint256 left, uint256 right) pure returns(uint256)
func (_Rollup *RollupCaller) HashLeftRight(opts *bind.CallOpts, left *big.Int, right *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "hashLeftRight", left, right)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashLeftRight is a free data retrieval call binding the contract method 0x5bb93995.
//
// Solidity: function hashLeftRight(uint256 left, uint256 right) pure returns(uint256)
func (_Rollup *RollupSession) HashLeftRight(left *big.Int, right *big.Int) (*big.Int, error) {
	return _Rollup.Contract.HashLeftRight(&_Rollup.CallOpts, left, right)
}

// HashLeftRight is a free data retrieval call binding the contract method 0x5bb93995.
//
// Solidity: function hashLeftRight(uint256 left, uint256 right) pure returns(uint256)
func (_Rollup *RollupCallerSession) HashLeftRight(left *big.Int, right *big.Int) (*big.Int, error) {
	return _Rollup.Contract.HashLeftRight(&_Rollup.CallOpts, left, right)
}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint256)
func (_Rollup *RollupCaller) Levels(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "levels")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint256)
func (_Rollup *RollupSession) Levels() (*big.Int, error) {
	return _Rollup.Contract.Levels(&_Rollup.CallOpts)
}

// Levels is a free data retrieval call binding the contract method 0x4ecf518b.
//
// Solidity: function levels() view returns(uint256)
func (_Rollup *RollupCallerSession) Levels() (*big.Int, error) {
	return _Rollup.Contract.Levels(&_Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupSession) Owner() (common.Address, error) {
	return _Rollup.Contract.Owner(&_Rollup.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rollup *RollupCallerSession) Owner() (common.Address, error) {
	return _Rollup.Contract.Owner(&_Rollup.CallOpts)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x3fd141b6.
//
// Solidity: function verifyMerkleProof(uint256[] path, uint64 index) view returns(bool)
func (_Rollup *RollupCaller) VerifyMerkleProof(opts *bind.CallOpts, path []*big.Int, index uint64) (bool, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "verifyMerkleProof", path, index)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x3fd141b6.
//
// Solidity: function verifyMerkleProof(uint256[] path, uint64 index) view returns(bool)
func (_Rollup *RollupSession) VerifyMerkleProof(path []*big.Int, index uint64) (bool, error) {
	return _Rollup.Contract.VerifyMerkleProof(&_Rollup.CallOpts, path, index)
}

// VerifyMerkleProof is a free data retrieval call binding the contract method 0x3fd141b6.
//
// Solidity: function verifyMerkleProof(uint256[] path, uint64 index) view returns(bool)
func (_Rollup *RollupCallerSession) VerifyMerkleProof(path []*big.Int, index uint64) (bool, error) {
	return _Rollup.Contract.VerifyMerkleProof(&_Rollup.CallOpts, path, index)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(uint256)
func (_Rollup *RollupCaller) Zeros(opts *bind.CallOpts, i *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Rollup.contract.Call(opts, &out, "zeros", i)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(uint256)
func (_Rollup *RollupSession) Zeros(i *big.Int) (*big.Int, error) {
	return _Rollup.Contract.Zeros(&_Rollup.CallOpts, i)
}

// Zeros is a free data retrieval call binding the contract method 0xe8295588.
//
// Solidity: function zeros(uint256 i) pure returns(uint256)
func (_Rollup *RollupCallerSession) Zeros(i *big.Int) (*big.Int, error) {
	return _Rollup.Contract.Zeros(&_Rollup.CallOpts, i)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Rollup *RollupTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Rollup *RollupSession) Deposit() (*types.Transaction, error) {
	return _Rollup.Contract.Deposit(&_Rollup.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_Rollup *RollupTransactorSession) Deposit() (*types.Transaction, error) {
	return _Rollup.Contract.Deposit(&_Rollup.TransactOpts)
}

// ProcessDeposit is a paid mutator transaction binding the contract method 0x1174dc99.
//
// Solidity: function processDeposit(uint256[] merkleProof, uint256 amount, address ganacheAddress, (uint64,uint256,uint256,(uint256,uint256)) account) returns()
func (_Rollup *RollupTransactor) ProcessDeposit(opts *bind.TransactOpts, merkleProof []*big.Int, amount *big.Int, ganacheAddress common.Address, account RollupAccount) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "processDeposit", merkleProof, amount, ganacheAddress, account)
}

// ProcessDeposit is a paid mutator transaction binding the contract method 0x1174dc99.
//
// Solidity: function processDeposit(uint256[] merkleProof, uint256 amount, address ganacheAddress, (uint64,uint256,uint256,(uint256,uint256)) account) returns()
func (_Rollup *RollupSession) ProcessDeposit(merkleProof []*big.Int, amount *big.Int, ganacheAddress common.Address, account RollupAccount) (*types.Transaction, error) {
	return _Rollup.Contract.ProcessDeposit(&_Rollup.TransactOpts, merkleProof, amount, ganacheAddress, account)
}

// ProcessDeposit is a paid mutator transaction binding the contract method 0x1174dc99.
//
// Solidity: function processDeposit(uint256[] merkleProof, uint256 amount, address ganacheAddress, (uint64,uint256,uint256,(uint256,uint256)) account) returns()
func (_Rollup *RollupTransactorSession) ProcessDeposit(merkleProof []*big.Int, amount *big.Int, ganacheAddress common.Address, account RollupAccount) (*types.Transaction, error) {
	return _Rollup.Contract.ProcessDeposit(&_Rollup.TransactOpts, merkleProof, amount, ganacheAddress, account)
}

// SubmitVerifyBatch is a paid mutator transaction binding the contract method 0xb994f687.
//
// Solidity: function submitVerifyBatch(uint256[8] zkproof, uint256 postStateRoot) returns()
func (_Rollup *RollupTransactor) SubmitVerifyBatch(opts *bind.TransactOpts, zkproof [8]*big.Int, postStateRoot *big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "submitVerifyBatch", zkproof, postStateRoot)
}

// SubmitVerifyBatch is a paid mutator transaction binding the contract method 0xb994f687.
//
// Solidity: function submitVerifyBatch(uint256[8] zkproof, uint256 postStateRoot) returns()
func (_Rollup *RollupSession) SubmitVerifyBatch(zkproof [8]*big.Int, postStateRoot *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitVerifyBatch(&_Rollup.TransactOpts, zkproof, postStateRoot)
}

// SubmitVerifyBatch is a paid mutator transaction binding the contract method 0xb994f687.
//
// Solidity: function submitVerifyBatch(uint256[8] zkproof, uint256 postStateRoot) returns()
func (_Rollup *RollupTransactorSession) SubmitVerifyBatch(zkproof [8]*big.Int, postStateRoot *big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitVerifyBatch(&_Rollup.TransactOpts, zkproof, postStateRoot)
}

// VerifyZkProof is a paid mutator transaction binding the contract method 0xcc0d802e.
//
// Solidity: function verifyZkProof(uint256[8] proof, uint256[2] input) returns()
func (_Rollup *RollupTransactor) VerifyZkProof(opts *bind.TransactOpts, proof [8]*big.Int, input [2]*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "verifyZkProof", proof, input)
}

// VerifyZkProof is a paid mutator transaction binding the contract method 0xcc0d802e.
//
// Solidity: function verifyZkProof(uint256[8] proof, uint256[2] input) returns()
func (_Rollup *RollupSession) VerifyZkProof(proof [8]*big.Int, input [2]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.VerifyZkProof(&_Rollup.TransactOpts, proof, input)
}

// VerifyZkProof is a paid mutator transaction binding the contract method 0xcc0d802e.
//
// Solidity: function verifyZkProof(uint256[8] proof, uint256[2] input) returns()
func (_Rollup *RollupTransactorSession) VerifyZkProof(proof [8]*big.Int, input [2]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.VerifyZkProof(&_Rollup.TransactOpts, proof, input)
}

// Withdraw is a paid mutator transaction binding the contract method 0x02722d20.
//
// Solidity: function withdraw(uint256 amount, uint256[] merkleProof, (uint64,uint256,uint256,(uint256,uint256)) account) returns()
func (_Rollup *RollupTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int, merkleProof []*big.Int, account RollupAccount) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "withdraw", amount, merkleProof, account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x02722d20.
//
// Solidity: function withdraw(uint256 amount, uint256[] merkleProof, (uint64,uint256,uint256,(uint256,uint256)) account) returns()
func (_Rollup *RollupSession) Withdraw(amount *big.Int, merkleProof []*big.Int, account RollupAccount) (*types.Transaction, error) {
	return _Rollup.Contract.Withdraw(&_Rollup.TransactOpts, amount, merkleProof, account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x02722d20.
//
// Solidity: function withdraw(uint256 amount, uint256[] merkleProof, (uint64,uint256,uint256,(uint256,uint256)) account) returns()
func (_Rollup *RollupTransactorSession) Withdraw(amount *big.Int, merkleProof []*big.Int, account RollupAccount) (*types.Transaction, error) {
	return _Rollup.Contract.Withdraw(&_Rollup.TransactOpts, amount, merkleProof, account)
}

// RollupDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Rollup contract.
type RollupDepositIterator struct {
	Event *RollupDeposit // Event containing the contract specifics and raw log

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
func (it *RollupDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupDeposit)
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
		it.Event = new(RollupDeposit)
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
func (it *RollupDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupDeposit represents a Deposit event raised by the Rollup contract.
type RollupDeposit struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Rollup *RollupFilterer) FilterDeposit(opts *bind.FilterOpts, user []common.Address) (*RollupDepositIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return &RollupDepositIterator{contract: _Rollup.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Rollup *RollupFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *RollupDeposit, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Deposit", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupDeposit)
				if err := _Rollup.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c.
//
// Solidity: event Deposit(address indexed user, uint256 amount)
func (_Rollup *RollupFilterer) ParseDeposit(log types.Log) (*RollupDeposit, error) {
	event := new(RollupDeposit)
	if err := _Rollup.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupMerkleRootUpdatedIterator is returned from FilterMerkleRootUpdated and is used to iterate over the raw logs and unpacked data for MerkleRootUpdated events raised by the Rollup contract.
type RollupMerkleRootUpdatedIterator struct {
	Event *RollupMerkleRootUpdated // Event containing the contract specifics and raw log

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
func (it *RollupMerkleRootUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupMerkleRootUpdated)
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
		it.Event = new(RollupMerkleRootUpdated)
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
func (it *RollupMerkleRootUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupMerkleRootUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupMerkleRootUpdated represents a MerkleRootUpdated event raised by the Rollup contract.
type RollupMerkleRootUpdated struct {
	NewMerkleRoot *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootUpdated is a free log retrieval operation binding the contract event 0x3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a.
//
// Solidity: event MerkleRootUpdated(uint256 newMerkleRoot)
func (_Rollup *RollupFilterer) FilterMerkleRootUpdated(opts *bind.FilterOpts) (*RollupMerkleRootUpdatedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "MerkleRootUpdated")
	if err != nil {
		return nil, err
	}
	return &RollupMerkleRootUpdatedIterator{contract: _Rollup.contract, event: "MerkleRootUpdated", logs: logs, sub: sub}, nil
}

// WatchMerkleRootUpdated is a free log subscription operation binding the contract event 0x3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a.
//
// Solidity: event MerkleRootUpdated(uint256 newMerkleRoot)
func (_Rollup *RollupFilterer) WatchMerkleRootUpdated(opts *bind.WatchOpts, sink chan<- *RollupMerkleRootUpdated) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "MerkleRootUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupMerkleRootUpdated)
				if err := _Rollup.contract.UnpackLog(event, "MerkleRootUpdated", log); err != nil {
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

// ParseMerkleRootUpdated is a log parse operation binding the contract event 0x3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a.
//
// Solidity: event MerkleRootUpdated(uint256 newMerkleRoot)
func (_Rollup *RollupFilterer) ParseMerkleRootUpdated(log types.Log) (*RollupMerkleRootUpdated, error) {
	event := new(RollupMerkleRootUpdated)
	if err := _Rollup.contract.UnpackLog(event, "MerkleRootUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Rollup contract.
type RollupRegisteredIterator struct {
	Event *RollupRegistered // Event containing the contract specifics and raw log

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
func (it *RollupRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupRegistered)
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
		it.Event = new(RollupRegistered)
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
func (it *RollupRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupRegistered represents a Registered event raised by the Rollup contract.
type RollupRegistered struct {
	Sender  common.Address
	Index   *big.Int
	Nonce   *big.Int
	Balance *big.Int
	Pubkey  RollupPublicKey
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x698a7ff468a195128113518f28a19ee3c5fdc665599eb479592ef44d324063e6.
//
// Solidity: event Registered(address sender, uint256 index, uint256 nonce, uint256 balance, (uint256,uint256) pubkey)
func (_Rollup *RollupFilterer) FilterRegistered(opts *bind.FilterOpts) (*RollupRegisteredIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &RollupRegisteredIterator{contract: _Rollup.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x698a7ff468a195128113518f28a19ee3c5fdc665599eb479592ef44d324063e6.
//
// Solidity: event Registered(address sender, uint256 index, uint256 nonce, uint256 balance, (uint256,uint256) pubkey)
func (_Rollup *RollupFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *RollupRegistered) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupRegistered)
				if err := _Rollup.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x698a7ff468a195128113518f28a19ee3c5fdc665599eb479592ef44d324063e6.
//
// Solidity: event Registered(address sender, uint256 index, uint256 nonce, uint256 balance, (uint256,uint256) pubkey)
func (_Rollup *RollupFilterer) ParseRegistered(log types.Log) (*RollupRegistered, error) {
	event := new(RollupRegistered)
	if err := _Rollup.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupTransactionDataIterator is returned from FilterTransactionData and is used to iterate over the raw logs and unpacked data for TransactionData events raised by the Rollup contract.
type RollupTransactionDataIterator struct {
	Event *RollupTransactionData // Event containing the contract specifics and raw log

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
func (it *RollupTransactionDataIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupTransactionData)
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
		it.Event = new(RollupTransactionData)
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
func (it *RollupTransactionDataIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupTransactionDataIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupTransactionData represents a TransactionData event raised by the Rollup contract.
type RollupTransactionData struct {
	BatchId *big.Int
	Data    common.Hash
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransactionData is a free log retrieval operation binding the contract event 0xa67113ee6649fb1f8a86e178a2c6214e4e5036f3862c7496ce7df28f41fff0d7.
//
// Solidity: event TransactionData(uint256 indexed batchId, bytes indexed data)
func (_Rollup *RollupFilterer) FilterTransactionData(opts *bind.FilterOpts, batchId []*big.Int, data [][]byte) (*RollupTransactionDataIterator, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "TransactionData", batchIdRule, dataRule)
	if err != nil {
		return nil, err
	}
	return &RollupTransactionDataIterator{contract: _Rollup.contract, event: "TransactionData", logs: logs, sub: sub}, nil
}

// WatchTransactionData is a free log subscription operation binding the contract event 0xa67113ee6649fb1f8a86e178a2c6214e4e5036f3862c7496ce7df28f41fff0d7.
//
// Solidity: event TransactionData(uint256 indexed batchId, bytes indexed data)
func (_Rollup *RollupFilterer) WatchTransactionData(opts *bind.WatchOpts, sink chan<- *RollupTransactionData, batchId []*big.Int, data [][]byte) (event.Subscription, error) {

	var batchIdRule []interface{}
	for _, batchIdItem := range batchId {
		batchIdRule = append(batchIdRule, batchIdItem)
	}
	var dataRule []interface{}
	for _, dataItem := range data {
		dataRule = append(dataRule, dataItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "TransactionData", batchIdRule, dataRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupTransactionData)
				if err := _Rollup.contract.UnpackLog(event, "TransactionData", log); err != nil {
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

// ParseTransactionData is a log parse operation binding the contract event 0xa67113ee6649fb1f8a86e178a2c6214e4e5036f3862c7496ce7df28f41fff0d7.
//
// Solidity: event TransactionData(uint256 indexed batchId, bytes indexed data)
func (_Rollup *RollupFilterer) ParseTransactionData(log types.Log) (*RollupTransactionData, error) {
	event := new(RollupTransactionData)
	if err := _Rollup.contract.UnpackLog(event, "TransactionData", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RollupWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Rollup contract.
type RollupWithdrawIterator struct {
	Event *RollupWithdraw // Event containing the contract specifics and raw log

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
func (it *RollupWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupWithdraw)
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
		it.Event = new(RollupWithdraw)
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
func (it *RollupWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupWithdraw represents a Withdraw event raised by the Rollup contract.
type RollupWithdraw struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Rollup *RollupFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address) (*RollupWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return &RollupWithdrawIterator{contract: _Rollup.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Rollup *RollupFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *RollupWithdraw, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Withdraw", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupWithdraw)
				if err := _Rollup.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed user, uint256 amount)
func (_Rollup *RollupFilterer) ParseWithdraw(log types.Log) (*RollupWithdraw, error) {
	event := new(RollupWithdraw)
	if err := _Rollup.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VerifierMetaData contains all meta data concerning the Verifier contract.
var VerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[2]\",\"name\":\"input\",\"type\":\"uint256[2]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b506121908061001d5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c806344f63692146100435780635fe24f2314610073578063f11817b21461008f575b5f80fd5b61005d60048036038101906100589190611eed565b6100ab565b60405161006a9190611fc7565b60405180910390f35b61008d60048036038101906100889190612001565b61020a565b005b6100a960048036038101906100a49190612062565b6104a5565b005b6100b3611e5d565b6100ec825f600881106100c9576100c86120a0565b5b6020020135836001600881106100e2576100e16120a0565b5b6020020135610a61565b815f600481106100ff576100fe6120a0565b5b6020020181815250506101748260036008811061011f5761011e6120a0565b5b602002013583600260088110610138576101376120a0565b5b602002013584600560088110610151576101506120a0565b5b60200201358560046008811061016a576101696120a0565b5b6020020135610c0f565b82600260048110610188576101876120a0565b5b60200201836001600481106101a05761019f6120a0565b5b60200201828152508281525050506101e8826006600881106101c5576101c46120a0565b5b6020020135836007600881106101de576101dd6120a0565b5b6020020135610a61565b816003600481106101fc576101fb6120a0565b5b602002018181525050919050565b5f80610215836111db565b915091505f6040516101008682377f2a0a395e749a6d8065c3aa6f56a8e21054ab08950fbb5396aa7b23548a0750996101008201527f03cd85ef4ea213c5f35ebd04de895b36f6d3a9a908c2c2ace34b1173cbeac3706101208201527f0ce840007a44fb67829beb581e09550635ddf1a5d44e6402a987f6310c54fb796101408201527f1156053259115281dcdd831be5ac033f9e77376ea1eaa5315eb726a800f59c446101608201527f28f39096ab99c023cfe9bcce11b088d12a41abd5a9c085f336a6789b5b5c0d236101808201527f1d446325d8a76607a162bc82fcfb50374970e844fec05a54154726c7a883aae86101a08201527f0fb7b7b4599f852b6ed861009e46e2d90b21ac3c4ac991d53a181b04399695fc6101c08201527f13b8f14cd90dd4abf42b4091e3d70e3eaf0e1229e2ac08362de30b85f73eeea56101e08201527f1dce2366eacb1561041dd7456ac5a56ae7503cf5b7b2d8c493e2a692f4af81216102008201527f0c7fda83e78eb0e0ac38fb9f95135896ab908a0d428422fbca6d366fbedebd5461022082015283610240820152826102608201527f0401b3d84d318fff9615e514d291d232c36afec9481ec2f5575a4aaee8833b5c6102808201527f1b9463d5943c466f0ebe906fa237c020b3e39fdfd3930c2feb9130c84eb46cf86102a08201527f12f660ad5e9ec1b81638fb9d80878b5c9d80c1f3d8144fff40da70687102f4ef6102c08201527f185758338d20958ad3e89dec68c6d608ec21e0083ea08aa23be0165b54bfe9836102e08201526020816103008360085afa9150805182169150508061049e576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b5f806104c7845f600481106104bd576104bc6120a0565b5b60200201356113ad565b915091505f805f80610509886002600481106104e6576104e56120a0565b5b6020020135896001600481106104ff576104fe6120a0565b5b60200201356114e2565b93509350935093505f806105348a60036004811061052a576105296120a0565b5b60200201356113ad565b915091505f806105438b6111db565b9150915061054f611e7f565b8a815f60188110610563576105626120a0565b5b6020020181815250508981600160188110610581576105806120a0565b5b602002018181525050878160026018811061059f5761059e6120a0565b5b60200201818152505088816003601881106105bd576105bc6120a0565b5b60200201818152505085816004601881106105db576105da6120a0565b5b60200201818152505086816005601881106105f9576105f86120a0565b5b6020020181815250508481600660188110610617576106166120a0565b5b6020020181815250508381600760188110610635576106346120a0565b5b6020020181815250507f2a0a395e749a6d8065c3aa6f56a8e21054ab08950fbb5396aa7b23548a07509981600860188110610673576106726120a0565b5b6020020181815250507f03cd85ef4ea213c5f35ebd04de895b36f6d3a9a908c2c2ace34b1173cbeac370816009601881106106b1576106b06120a0565b5b6020020181815250507f0ce840007a44fb67829beb581e09550635ddf1a5d44e6402a987f6310c54fb7981600a601881106106ef576106ee6120a0565b5b6020020181815250507f1156053259115281dcdd831be5ac033f9e77376ea1eaa5315eb726a800f59c4481600b6018811061072d5761072c6120a0565b5b6020020181815250507f28f39096ab99c023cfe9bcce11b088d12a41abd5a9c085f336a6789b5b5c0d2381600c6018811061076b5761076a6120a0565b5b6020020181815250507f1d446325d8a76607a162bc82fcfb50374970e844fec05a54154726c7a883aae881600d601881106107a9576107a86120a0565b5b6020020181815250507f0fb7b7b4599f852b6ed861009e46e2d90b21ac3c4ac991d53a181b04399695fc81600e601881106107e7576107e66120a0565b5b6020020181815250507f13b8f14cd90dd4abf42b4091e3d70e3eaf0e1229e2ac08362de30b85f73eeea581600f60188110610825576108246120a0565b5b6020020181815250507f1dce2366eacb1561041dd7456ac5a56ae7503cf5b7b2d8c493e2a692f4af812181601060188110610863576108626120a0565b5b6020020181815250507f0c7fda83e78eb0e0ac38fb9f95135896ab908a0d428422fbca6d366fbedebd54816011601881106108a1576108a06120a0565b5b60200201818152505082816012601881106108bf576108be6120a0565b5b60200201818152505081816013601881106108dd576108dc6120a0565b5b6020020181815250507f0401b3d84d318fff9615e514d291d232c36afec9481ec2f5575a4aaee8833b5c8160146018811061091b5761091a6120a0565b5b6020020181815250507f1b9463d5943c466f0ebe906fa237c020b3e39fdfd3930c2feb9130c84eb46cf881601560188110610959576109586120a0565b5b6020020181815250507f12f660ad5e9ec1b81638fb9d80878b5c9d80c1f3d8144fff40da70687102f4ef81601660188110610997576109966120a0565b5b6020020181815250507f185758338d20958ad3e89dec68c6d608ec21e0083ea08aa23be0165b54bfe983816017601881106109d5576109d46120a0565b5b6020020181815250505f6109e7611ea2565b6020816103008560085afa9150811580610a1957506001815f60018110610a1157610a106120a0565b5b602002015114155b15610a50576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050505050505050505050565b5f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4783101580610ab157507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478210155b15610ae8576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f83148015610af657505f82145b15610b03575f9050610c09565b5f610ba17f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610b3657610b356120cd565b5b60037f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610b6757610b666120cd565b5b877f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610b9757610b966120cd565b5b898a0909086118bd565b9050808303610bb9575f600185901b17915050610c09565b610bc281611959565b8303610bd75760018085901b17915050610c09565b6040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b92915050565b5f807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4786101580610c6057507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478510155b80610c8b57507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478410155b80610cb657507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478310155b15610ced576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8385878917171703610d05575f80915091506111d2565b5f805f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610d3757610d366120cd565b5b60037f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47610d649190612127565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610d9357610d926120cd565b5b8a8c090990505f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610dc957610dc86120cd565b5b8a7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610df957610df86120cd565b5b8c8d090990505f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610e2f57610e2e6120cd565b5b8a7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610e5f57610e5e6120cd565b5b8c8d090990507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610e9457610e936120cd565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610ec357610ec26120cd565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610ef257610ef16120cd565b5b8c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450610fd77f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610f4d57610f4c6120cd565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610f7c57610f7b6120cd565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780610fab57610faa6120cd565b5b8e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611959565b93505050505f8061107a7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806110105761100f6120cd565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061103f5761103e6120cd565b5b8586097f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611071576110706120cd565b5b878809086118bd565b90506111077f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806110ae576110ad6120cd565b5b7f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea47f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806110fe576110fd6120cd565b5b848808096119c4565b15915050611116838383611a2e565b8093508194505050828714801561112c57508186145b15611154575f8161113d575f611140565b60025b60ff1660028b901b171794508793506111ce565b61115d83611959565b87148015611172575061116f82611959565b86145b1561119b57600181611184575f611187565b60025b60ff1660028b901b171794508793506111cd565b6040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b5050505b94509492505050565b5f805f60019050604051604081015f7f01ec5bbb485ef1a3414f7591f7912c6ade47e3cd3f4fedbf62bd8ce01bddac5083527f20a77c62e6088ed1c4d7a2076b2e393274b6fcb43b508803bbe2674c839754b660208401527f297782c7af6cb27fd80dd4e4d07add9b1abf3fb77e8286ea82f7e1c6aca4c4a682527f2e50ffd913db6cb5bf2d85ab1eeded735f66016fdb47a6433e0df86f4e33d72d6020830152863590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f018570f62856ecbd85aa7b2367100418440fb9ab57ffbd982185393e8bf8fd1582527f0682bdd37018f03e0f554dbf688cddb2c399a5af1e3123155f9d7e64de743e786020830152602087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693508251955060208301519450505050806113a7576040517fa54f8e2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50915091565b5f805f83036113c1575f80915091506114dd565b5f6001808516149050600184901c92507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47831061142a576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6114c77f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061145c5761145b6120cd565b5b60037f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061148d5761148c6120cd565b5b867f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806114bd576114bc6120cd565b5b88890909086118bd565b915080156114db576114d882611959565b91505b505b915091565b5f805f805f861480156114f457505f85145b1561150a575f805f8093509350935093506118b4565b5f60018088161490505f6002808916149050600288901c95508694507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478610158061157557507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478510155b156115ac576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806115dc576115db6120cd565b5b60037f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd476116099190612127565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611638576116376120cd565b5b888a090990505f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061166e5761166d6120cd565b5b887f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061169e5761169d6120cd565b5b8a8b090990505f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806116d4576116d36120cd565b5b887f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611704576117036120cd565b5b8a8b090990507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611739576117386120cd565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611768576117676120cd565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611797576117966120cd565b5b8a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508965061187c7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806117f2576117f16120cd565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611821576118206120cd565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806118505761184f6120cd565b5b8c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e77508611959565b9550611889878786611a2e565b809750819850505084156118ae576118a087611959565b96506118ab86611959565b95505b50505050505b92959194509250565b5f6118e8827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611d29565b9050817f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061191a576119196120cd565b5b82830914611954576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b5f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780838161198b5761198a6120cd565b5b067f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4703816119bc576119bb6120cd565b5b069050919050565b5f806119f0837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f52611d29565b9050827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611a2257611a216120cd565b5b82830914915050919050565b5f805f611acd7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611a6357611a626120cd565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611a9257611a916120cd565b5b8788097f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611ac457611ac36120cd565b5b898a09086118bd565b90508315611ae157611ade81611959565b90505b611b6c7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611b1357611b126120cd565b5b7f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea47f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611b6357611b626120cd565b5b848a08096118bd565b92507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611b9d57611b9c6120cd565b5b611bd87f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611bcf57611bce6120cd565b5b60028609611dc0565b860991507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611c0b57611c0a6120cd565b5b611c457f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611c3d57611c3c6120cd565b5b848509611959565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611c7457611c736120cd565b5b8586090886141580611ce957507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611cb057611caf6120cd565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611cdf57611cde6120cd565b5b8385096002098514155b15611d20576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50935093915050565b5f8060405160208152602080820152602060408201528460608201528360808201527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760a082015260208160c08360055afa9150805192505080611db9576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5092915050565b5f611deb827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd45611d29565b905060017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780611e1e57611e1d6120cd565b5b82840914611e58576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040518060800160405280600490602082028036833780820191505090505090565b604051806103000160405280601890602082028036833780820191505090505090565b6040518060200160405280600190602082028036833780820191505090505090565b5f80fd5b5f80fd5b5f81905082602060080282011115611ee757611ee6611ec8565b5b92915050565b5f6101008284031215611f0357611f02611ec4565b5b5f611f1084828501611ecc565b91505092915050565b5f60049050919050565b5f81905092915050565b5f819050919050565b5f819050919050565b611f4881611f36565b82525050565b5f611f598383611f3f565b60208301905092915050565b5f602082019050919050565b611f7a81611f19565b611f848184611f23565b9250611f8f82611f2d565b805f5b83811015611fbf578151611fa68782611f4e565b9650611fb183611f65565b925050600181019050611f92565b505050505050565b5f608082019050611fda5f830184611f71565b92915050565b5f81905082602060020282011115611ffb57611ffa611ec8565b5b92915050565b5f80610140838503121561201857612017611ec4565b5b5f61202585828601611ecc565b92505061010061203785828601611fe0565b9150509250929050565b5f8190508260206004028201111561205c5761205b611ec8565b5b92915050565b5f8060c0838503121561207857612077611ec4565b5b5f61208585828601612041565b925050608061209685828601611fe0565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61213182611f36565b915061213c83611f36565b9250828203905081811115612154576121536120fa565b5b9291505056fea26469706673582212201dfbd54cc100dcd28c35a8ab39482a1956479980104e2274498f3fae7d6cfaca64736f6c63430008180033",
}

// VerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use VerifierMetaData.ABI instead.
var VerifierABI = VerifierMetaData.ABI

// VerifierBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VerifierMetaData.Bin instead.
var VerifierBin = VerifierMetaData.Bin

// DeployVerifier deploys a new Ethereum contract, binding an instance of Verifier to it.
func DeployVerifier(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Verifier, error) {
	parsed, err := VerifierMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VerifierBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// Verifier is an auto generated Go binding around an Ethereum contract.
type Verifier struct {
	VerifierCaller     // Read-only binding to the contract
	VerifierTransactor // Write-only binding to the contract
	VerifierFilterer   // Log filterer for contract events
}

// VerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type VerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VerifierSession struct {
	Contract     *Verifier         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VerifierCallerSession struct {
	Contract *VerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// VerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VerifierTransactorSession struct {
	Contract     *VerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// VerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type VerifierRaw struct {
	Contract *Verifier // Generic contract binding to access the raw methods on
}

// VerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VerifierCallerRaw struct {
	Contract *VerifierCaller // Generic read-only contract binding to access the raw methods on
}

// VerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VerifierTransactorRaw struct {
	Contract *VerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVerifier creates a new instance of Verifier, bound to a specific deployed contract.
func NewVerifier(address common.Address, backend bind.ContractBackend) (*Verifier, error) {
	contract, err := bindVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Verifier{VerifierCaller: VerifierCaller{contract: contract}, VerifierTransactor: VerifierTransactor{contract: contract}, VerifierFilterer: VerifierFilterer{contract: contract}}, nil
}

// NewVerifierCaller creates a new read-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierCaller(address common.Address, caller bind.ContractCaller) (*VerifierCaller, error) {
	contract, err := bindVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierCaller{contract: contract}, nil
}

// NewVerifierTransactor creates a new write-only instance of Verifier, bound to a specific deployed contract.
func NewVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*VerifierTransactor, error) {
	contract, err := bindVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VerifierTransactor{contract: contract}, nil
}

// NewVerifierFilterer creates a new log filterer instance of Verifier, bound to a specific deployed contract.
func NewVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*VerifierFilterer, error) {
	contract, err := bindVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VerifierFilterer{contract: contract}, nil
}

// bindVerifier binds a generic wrapper to an already deployed contract.
func bindVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.VerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.VerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Verifier *VerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Verifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Verifier *VerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Verifier *VerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Verifier.Contract.contract.Transact(opts, method, params...)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_Verifier *VerifierCaller) CompressProof(opts *bind.CallOpts, proof [8]*big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "compressProof", proof)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_Verifier *VerifierSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _Verifier.Contract.CompressProof(&_Verifier.CallOpts, proof)
}

// CompressProof is a free data retrieval call binding the contract method 0x44f63692.
//
// Solidity: function compressProof(uint256[8] proof) view returns(uint256[4] compressed)
func (_Verifier *VerifierCallerSession) CompressProof(proof [8]*big.Int) ([4]*big.Int, error) {
	return _Verifier.Contract.CompressProof(&_Verifier.CallOpts, proof)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf11817b2.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[2] input) view returns()
func (_Verifier *VerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [2]*big.Int) error {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf11817b2.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[2] input) view returns()
func (_Verifier *VerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [2]*big.Int) error {
	return _Verifier.Contract.VerifyCompressedProof(&_Verifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xf11817b2.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[2] input) view returns()
func (_Verifier *VerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [2]*big.Int) error {
	return _Verifier.Contract.VerifyCompressedProof(&_Verifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe24f23.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] input) view returns()
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [2]*big.Int) error {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe24f23.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] input) view returns()
func (_Verifier *VerifierSession) VerifyProof(proof [8]*big.Int, input [2]*big.Int) error {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x5fe24f23.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[2] input) view returns()
func (_Verifier *VerifierCallerSession) VerifyProof(proof [8]*big.Int, input [2]*big.Int) error {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, proof, input)
}
