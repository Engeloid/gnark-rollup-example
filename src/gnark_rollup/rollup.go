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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifierAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_treeLevels\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initRoot\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMerkleRoot\",\"type\":\"uint256\"}],\"name\":\"MerkleRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structRollup.PublicKey\",\"name\":\"pubkey\",\"type\":\"tuple\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TransactionData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"RollupToEth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"computeMerkleRootFromPath\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rollupAccountPKX\",\"type\":\"uint256\"}],\"name\":\"getEthAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLevels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextLeafIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structRollup.PublicKey\",\"name\":\"pubKey\",\"type\":\"tuple\"}],\"internalType\":\"structRollup.Account\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"hashAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"}],\"name\":\"hashLeftRight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"levels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"merkleProof\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ganacheAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structRollup.PublicKey\",\"name\":\"pubKey\",\"type\":\"tuple\"}],\"internalType\":\"structRollup.Account\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"processDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"zkproof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[128]\",\"name\":\"input\",\"type\":\"uint256[128]\"}],\"name\":\"submitVerifyBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"merkleProof\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structRollup.PublicKey\",\"name\":\"pubKey\",\"type\":\"tuple\"}],\"internalType\":\"structRollup.Account\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"zeros\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b506040516200282c3803806200282c8339818101604052810190620000369190620002a6565b815f81116200007c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000739062000383565b60405180910390fd5b60098110620000c2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000b990620003f1565b60405180910390fd5b6004811462000108576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000ff9062000485565b60405180910390fd5b805f819055507f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd3226760018190555050815f8190555082600760086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503360045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620001ce81620001ff60201b60201c565b5f60075f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050620004a5565b8060018190555050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000238826200020d565b9050919050565b6200024a816200022c565b811462000255575f80fd5b50565b5f8151905062000268816200023f565b92915050565b5f819050919050565b62000282816200026e565b81146200028d575f80fd5b50565b5f81519050620002a08162000277565b92915050565b5f805f60608486031215620002c057620002bf62000209565b5b5f620002cf8682870162000258565b9350506020620002e28682870162000290565b9250506040620002f58682870162000290565b9150509250925092565b5f82825260208201905092915050565b7f5f6c6576656c732073686f756c642062652067726561746572207468616e207a5f8201527f65726f0000000000000000000000000000000000000000000000000000000000602082015250565b5f6200036b602383620002ff565b915062000378826200030f565b604082019050919050565b5f6020820190508181035f8301526200039c816200035d565b9050919050565b7f5f6c6576656c732073686f756c64206265206c657373207468616e20380000005f82015250565b5f620003d9601d83620002ff565b9150620003e682620003a3565b602082019050919050565b5f6020820190508181035f8301526200040a81620003cb565b9050919050565b7f5f6c6576656c732073686f756c6420626520342028325e34203d2031362061635f8201527f636f756e74732920666f722074657374696e6700000000000000000000000000602082015250565b5f6200046d603383620002ff565b91506200047a8262000411565b604082019050919050565b5f6020820190508181035f8301526200049e816200045f565b9050919050565b61237980620004b35f395ff3fe608060405260043610610113575f3560e01c80635ca1e1651161009f578063bd4986a011610063578063bd4986a0146103cb578063d0e30db014610407578063d32fec6a14610411578063e829558814610439578063ec7329591461047557610113565b80635ca1e165146102c357806369ed88ec146102ed57806383290a3c146103295780638da5cb5b14610365578063909ad5ba1461038f57610113565b80633fd141b6116100e65780633fd141b6146101cd57806349590657146102095780634ecf518b1461023357806350e9b9251461025d5780635bb939951461028757610113565b806302722d20146101175780630c394a601461013f5780631174dc99146101695780631ea5f6eb14610191575b5f80fd5b348015610122575f80fd5b5061013d60048036038101906101389190611918565b61049f565b005b34801561014a575f80fd5b506101536107c9565b6040516101609190611993565b60405180910390f35b348015610174575f80fd5b5061018f600480360381019061018a9190611a06565b6107d1565b005b34801561019c575f80fd5b506101b760048036038101906101b29190611a87565b610b38565b6040516101c49190611ac1565b60405180910390f35b3480156101d8575f80fd5b506101f360048036038101906101ee9190611ada565b610b68565b6040516102009190611b4e565b60405180910390f35b348015610214575f80fd5b5061021d610d13565b60405161022a9190611993565b60405180910390f35b34801561023e575f80fd5b50610247610d21565b6040516102549190611993565b60405180910390f35b348015610268575f80fd5b50610271610d26565b60405161027e9190611993565b60405180910390f35b348015610292575f80fd5b506102ad60048036038101906102a89190611b67565b610d2f565b6040516102ba9190611993565b60405180910390f35b3480156102ce575f80fd5b506102d7610e3d565b6040516102e49190611993565b60405180910390f35b3480156102f8575f80fd5b50610313600480360381019061030e9190611ada565b610e46565b6040516103209190611993565b60405180910390f35b348015610334575f80fd5b5061034f600480360381019061034a9190611a87565b610fed565b60405161035c9190611993565b60405180910390f35b348015610370575f80fd5b5061037961100d565b6040516103869190611ac1565b60405180910390f35b34801561039a575f80fd5b506103b560048036038101906103b09190611ba5565b611032565b6040516103c29190611993565b60405180910390f35b3480156103d6575f80fd5b506103f160048036038101906103ec9190611a87565b6111c6565b6040516103fe9190611ac1565b60405180910390f35b61040f6111ff565b005b34801561041c575f80fd5b5061043760048036038101906104329190611c12565b6112db565b005b348015610444575f80fd5b5061045f600480360381019061045a9190611a87565b611493565b60405161046c9190611993565b60405180910390f35b348015610480575f80fd5b50610489611653565b6040516104969190611993565b60405180910390f35b5f60055f83606001515f015181526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610548576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053f90611cac565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ad90611d14565b60405180910390fd5b83826040015110156105fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f490611d7c565b60405180910390fd5b5f61060783611032565b905080845f8151811061061d5761061c611d9a565b5b60200260200101818152505061063684845f0151610b68565b610675576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066c90611e11565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc8690811502906040515f60405180830381858888f193505050501580156106b8573d5f803e3d5ffd5b508173ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364866040516106ff9190611993565b60405180910390a284836040018181516107199190611e5c565b915081815250508483604001516107309190611e8f565b8360400151111561074457610743611ec2565b5b5f61074e84611032565b905080855f8151811061076457610763611d9a565b5b6020026020010181815250505f61077e86865f0151610e46565b905061078981611677565b7f3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a816040516107b89190611993565b60405180910390a150505050505050565b5f8054905090565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610860576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085790611f39565b60405180910390fd5b61086d84825f0151610b68565b6108ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a390611fa1565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660055f83606001515f015181526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036109b7577f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd32267845f8151811061094a57610949611d9a565b5b6020026020010151146109605761095f611ec2565b5b8160055f83606001515f015181526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600660075f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16815481106109ea576109e9611d9a565b5b905f5260205f200154610a13848473ffffffffffffffffffffffffffffffffffffffff16610d2f565b14610a53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4a90612009565b60405180910390fd5b8281604001511015610a6857610a67611ec2565b5b5f610a7282611032565b905080855f81518110610a8857610a87611d9a565b5b6020026020010181815250505f610aa286845f0151610e46565b9050610aad81611677565b7f3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a81604051610adc9190611993565b60405180910390a160075f81819054906101000a900467ffffffffffffffff1680929190610b0990612027565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050505050565b6005602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f80600167ffffffffffffffff811115610b8557610b846116d9565b5b604051908082528060200260200182016040528015610bb35781602001602082028036833780820191505090505b509050835f81518110610bc957610bc8611d9a565b5b6020026020010151815f81518110610be457610be3611d9a565b5b6020026020010181815250505f73__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49836040518263ffffffff1660e01b8152600401610c28919061210d565b602060405180830381865af4158015610c43573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c679190612141565b90505f600190505b8551811015610d03575f600286610c869190612199565b67ffffffffffffffff1603610cc057610cb982878381518110610cac57610cab611d9a565b5b6020026020010151610d2f565b9150610ce7565b610ce4868281518110610cd657610cd5611d9a565b5b602002602001015183610d2f565b91505b600285610cf491906121c9565b94508080600101915050610c6f565b5060015481149250505092915050565b5f610d1c610e3d565b905090565b5f5481565b5f600254905090565b5f80600267ffffffffffffffff811115610d4c57610d4b6116d9565b5b604051908082528060200260200182016040528015610d7a5781602001602082028036833780820191505090505b50905083815f81518110610d9157610d90611d9a565b5b6020026020010181815250508281600181518110610db257610db1611d9a565b5b60200260200101818152505073__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49826040518263ffffffff1660e01b8152600401610df5919061210d565b602060405180830381865af4158015610e10573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e349190612141565b91505092915050565b5f600154905090565b5f80600167ffffffffffffffff811115610e6357610e626116d9565b5b604051908082528060200260200182016040528015610e915781602001602082028036833780820191505090505b509050835f81518110610ea757610ea6611d9a565b5b6020026020010151815f81518110610ec257610ec1611d9a565b5b6020026020010181815250505f73__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49836040518263ffffffff1660e01b8152600401610f06919061210d565b602060405180830381865af4158015610f21573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f459190612141565b90505f600190505b8551811015610fe1575f600286610f649190612199565b67ffffffffffffffff1603610f9e57610f9782878381518110610f8a57610f89611d9a565b5b6020026020010151610d2f565b9150610fc5565b610fc2868281518110610fb457610fb3611d9a565b5b602002602001015183610d2f565b91505b600285610fd291906121c9565b94508080600101915050610f4d565b50809250505092915050565b60068181548110610ffc575f80fd5b905f5260205f20015f915090505481565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f80600567ffffffffffffffff81111561104f5761104e6116d9565b5b60405190808252806020026020018201604052801561107d5781602001602082028036833780820191505090505b509050825f015167ffffffffffffffff16815f815181106110a1576110a0611d9a565b5b6020026020010181815250508260200151816001815181106110c6576110c5611d9a565b5b6020026020010181815250508260400151816002815181106110eb576110ea611d9a565b5b60200260200101818152505082606001515f01518160038151811061111357611112611d9a565b5b6020026020010181815250508260600151602001518160048151811061113c5761113b611d9a565b5b60200260200101818152505073__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49826040518263ffffffff1660e01b815260040161117f919061210d565b602060405180830381865af415801561119a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111be9190612141565b915050919050565b5f60055f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f3411611241576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123890612269565b60405180910390fd5b5f611262343373ffffffffffffffffffffffffffffffffffffffff16610d2f565b9050600681908060018154018082558091505060019003905f5260205f20015f90919091909150553373ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c346040516112d09190611993565b60405180910390a250565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461136a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161136190611f39565b60405180910390fd5b5f611373610e3d565b9050815f6080811061138857611387611d9a565b5b6020020135811461139c5761139b611ec2565b5b600760089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630698b54284846040518363ffffffff1660e01b81526004016113f99291906122b2565b5f6040518083038186803b15801561140f575f80fd5b505afa158015611421573d5f803e3d5ffd5b505050505f82600160806114359190611e5c565b6080811061144657611445611d9a565b5b6020020135905061145681611677565b7f3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a816040516114859190611993565b60405180910390a150505050565b5f8082036114c3577f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd32267905061164e565b600182036114f3577f1541c2e47d6e125aea7a107e749299d88fad35bd227179d4fb2bb0f63296d7c7905061164e565b60028203611523577f22d9432cd2f62dedd322bef0f3be1c172e71b6a1afe7866aee3b031ffd89aad7905061164e565b60038203611553577f049889855d82d41769e27ca791a905e528d867095ce49a8ac961ef4722ba6a16905061164e565b60048203611583577f2d1764e964f231d9129b555e1a653a65fca6e20051a0f7a388b5fd15fbb59455905061164e565b600582036115b3577f26d2dd833a220c13591b1a528c91e635044b3fac5b46f4133435ff4c87f04694905061164e565b600682036115e3577f2c42c8bc5bb243c5fbf697f47b4f41ba191664a8724e5e370090d6d8817930a8905061164e565b60078203611613577f1aaca32cb69540012989bd902c04015555aa164980a490df8c1e6be64fa57fca905061164e565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161164590612325565b60405180910390fd5b919050565b7f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd3226781565b8060018190555050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6116a481611692565b81146116ae575f80fd5b50565b5f813590506116bf8161169b565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61170f826116c9565b810181811067ffffffffffffffff8211171561172e5761172d6116d9565b5b80604052505050565b5f611740611681565b905061174c8282611706565b919050565b5f67ffffffffffffffff82111561176b5761176a6116d9565b5b602082029050602081019050919050565b5f80fd5b5f61179261178d84611751565b611737565b905080838252602082019050602084028301858111156117b5576117b461177c565b5b835b818110156117de57806117ca88826116b1565b8452602084019350506020810190506117b7565b5050509392505050565b5f82601f8301126117fc576117fb6116c5565b5b813561180c848260208601611780565b91505092915050565b5f80fd5b5f67ffffffffffffffff82169050919050565b61183581611819565b811461183f575f80fd5b50565b5f813590506118508161182c565b92915050565b5f6040828403121561186b5761186a611815565b5b6118756040611737565b90505f611884848285016116b1565b5f830152506020611897848285016116b1565b60208301525092915050565b5f60a082840312156118b8576118b7611815565b5b6118c26080611737565b90505f6118d184828501611842565b5f8301525060206118e4848285016116b1565b60208301525060406118f8848285016116b1565b604083015250606061190c84828501611856565b60608301525092915050565b5f805f60e0848603121561192f5761192e61168a565b5b5f61193c868287016116b1565b935050602084013567ffffffffffffffff81111561195d5761195c61168e565b5b611969868287016117e8565b925050604061197a868287016118a3565b9150509250925092565b61198d81611692565b82525050565b5f6020820190506119a65f830184611984565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6119d5826119ac565b9050919050565b6119e5816119cb565b81146119ef575f80fd5b50565b5f81359050611a00816119dc565b92915050565b5f805f806101008587031215611a1f57611a1e61168a565b5b5f85013567ffffffffffffffff811115611a3c57611a3b61168e565b5b611a48878288016117e8565b9450506020611a59878288016116b1565b9350506040611a6a878288016119f2565b9250506060611a7b878288016118a3565b91505092959194509250565b5f60208284031215611a9c57611a9b61168a565b5b5f611aa9848285016116b1565b91505092915050565b611abb816119cb565b82525050565b5f602082019050611ad45f830184611ab2565b92915050565b5f8060408385031215611af057611aef61168a565b5b5f83013567ffffffffffffffff811115611b0d57611b0c61168e565b5b611b19858286016117e8565b9250506020611b2a85828601611842565b9150509250929050565b5f8115159050919050565b611b4881611b34565b82525050565b5f602082019050611b615f830184611b3f565b92915050565b5f8060408385031215611b7d57611b7c61168a565b5b5f611b8a858286016116b1565b9250506020611b9b858286016116b1565b9150509250929050565b5f60a08284031215611bba57611bb961168a565b5b5f611bc7848285016118a3565b91505092915050565b5f81905082602060080282011115611beb57611bea61177c565b5b92915050565b5f81905082602060800282011115611c0c57611c0b61177c565b5b92915050565b5f806111008385031215611c2957611c2861168a565b5b5f611c3685828601611bd0565b925050610100611c4885828601611bf1565b9150509250929050565b5f82825260208201905092915050565b7f526f6c6c7570206164647265737320646f6573206e6f742065786973740000005f82015250565b5f611c96601d83611c52565b9150611ca182611c62565b602082019050919050565b5f6020820190508181035f830152611cc381611c8a565b9050919050565b7f4e6f7420726967687466756c206163636f756e742061646472657373000000005f82015250565b5f611cfe601c83611c52565b9150611d0982611cca565b602082019050919050565b5f6020820190508181035f830152611d2b81611cf2565b9050919050565b7f42616c616e6365206e6f742073756666696369656e74000000000000000000005f82015250565b5f611d66601683611c52565b9150611d7182611d32565b602082019050919050565b5f6020820190508181035f830152611d9381611d5a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f496e76616c6964204d65726b6c652070726f6f660000000000000000000000005f82015250565b5f611dfb601483611c52565b9150611e0682611dc7565b602082019050919050565b5f6020820190508181035f830152611e2881611def565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611e6682611692565b9150611e7183611692565b9250828203905081811115611e8957611e88611e2f565b5b92915050565b5f611e9982611692565b9150611ea483611692565b9250828201905080821115611ebc57611ebb611e2f565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffd5b7f4f6e6c79206f776e65722063616e207375626d697420626174636865730000005f82015250565b5f611f23601d83611c52565b9150611f2e82611eef565b602082019050919050565b5f6020820190508181035f830152611f5081611f17565b9050919050565b7f4d65726b6c652070726f6f6620696e76616c69640000000000000000000000005f82015250565b5f611f8b601483611c52565b9150611f9682611f57565b602082019050919050565b5f6020820190508181035f830152611fb881611f7f565b9050919050565b7f4465706f736974206e6f7420636f7272656374206f6e6520696e2071756575655f82015250565b5f611ff3602083611c52565b9150611ffe82611fbf565b602082019050919050565b5f6020820190508181035f83015261202081611fe7565b9050919050565b5f61203182611819565b915067ffffffffffffffff820361204b5761204a611e2f565b5b600182019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61208881611692565b82525050565b5f612099838361207f565b60208301905092915050565b5f602082019050919050565b5f6120bb82612056565b6120c58185612060565b93506120d083612070565b805f5b838110156121005781516120e7888261208e565b97506120f2836120a5565b9250506001810190506120d3565b5085935050505092915050565b5f6020820190508181035f83015261212581846120b1565b905092915050565b5f8151905061213b8161169b565b92915050565b5f602082840312156121565761215561168a565b5b5f6121638482850161212d565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f6121a382611819565b91506121ae83611819565b9250826121be576121bd61216c565b5b828206905092915050565b5f6121d382611819565b91506121de83611819565b9250826121ee576121ed61216c565b5b828204905092915050565b7f4465706f7369742076616c7565206d75737420626520677265617465722074685f8201527f616e203000000000000000000000000000000000000000000000000000000000602082015250565b5f612253602483611c52565b915061225e826121f9565b604082019050919050565b5f6020820190508181035f83015261228081612247565b9050919050565b82818337505050565b61229d6101008383612287565b5050565b6122ae6110008383612287565b5050565b5f611100820190506122c65f830185612290565b6122d46101008301846122a1565b9392505050565b7f696e646578206f7574206f6620626f756e6473000000000000000000000000005f82015250565b5f61230f601383611c52565b915061231a826122db565b602082019050919050565b5f6020820190508181035f83015261233c81612303565b905091905056fea2646970667358221220d6b495db11ea89ba88494b7915776387fd3078b1f5b93491e6a048298b93818664736f6c63430008180033",
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

// SubmitVerifyBatch is a paid mutator transaction binding the contract method 0xd32fec6a.
//
// Solidity: function submitVerifyBatch(uint256[8] zkproof, uint256[128] input) returns()
func (_Rollup *RollupTransactor) SubmitVerifyBatch(opts *bind.TransactOpts, zkproof [8]*big.Int, input [128]*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "submitVerifyBatch", zkproof, input)
}

// SubmitVerifyBatch is a paid mutator transaction binding the contract method 0xd32fec6a.
//
// Solidity: function submitVerifyBatch(uint256[8] zkproof, uint256[128] input) returns()
func (_Rollup *RollupSession) SubmitVerifyBatch(zkproof [8]*big.Int, input [128]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitVerifyBatch(&_Rollup.TransactOpts, zkproof, input)
}

// SubmitVerifyBatch is a paid mutator transaction binding the contract method 0xd32fec6a.
//
// Solidity: function submitVerifyBatch(uint256[8] zkproof, uint256[128] input) returns()
func (_Rollup *RollupTransactorSession) SubmitVerifyBatch(zkproof [8]*big.Int, input [128]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitVerifyBatch(&_Rollup.TransactOpts, zkproof, input)
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
	ABI: "[{\"inputs\":[],\"name\":\"ProofInvalid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PublicInputNotInField\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"}],\"name\":\"compressProof\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressed\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"compressedProof\",\"type\":\"uint256[4]\"},{\"internalType\":\"uint256[128]\",\"name\":\"input\",\"type\":\"uint256[128]\"}],\"name\":\"verifyCompressedProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[128]\",\"name\":\"input\",\"type\":\"uint256[128]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50616d52806200001e5f395ff3fe608060405234801561000f575f80fd5b506004361061003f575f3560e01c80630698b5421461004357806344f636921461005f578063d0eff68f1461008f575b5f80fd5b61005d60048036038101906100589190616acf565b6100ab565b005b61007960048036038101906100749190616b0f565b610346565b6040516100869190616be9565b60405180910390f35b6100a960048036038101906100a49190616c23565b6104a5565b005b5f806100b683610a61565b915091505f6040516101008682377f10b6c381ef69d708661b933885c1f66e08724d9001c135ee1c7b1b0da08498fb6101008201527f0452d2815cc5d0d4190d27ebe866886a81af0042d8a40091286a160c05efb8c06101208201527f27e5ca7d3c6a0485578f48e3c1bd5f84e2b32204b40d5bf7f5141844289044f16101408201527f124f33a5c7f3ca32aaf25c81a4fa5932cd8bc08b09d40d2d1027aa72d5c84c9a6101608201527f2d686e8598465a029d0f854888dbf622ff2be2a06748ef10db44f579b1426d216101808201527f1f7ea897a6112ad5bc7f8765815f62d3ab0d4547763865000bac252e71905ae66101a08201527f0a8e869f6f733346e8733a4d3ec42d86e72251138bf945b566b8fcd3c08543166101c08201527f0d5c64d136a42af58a77d8602c8412f4a1e1e80567207e8e6cd07c661e8644766101e08201527f1c24c1e49253789a3434be06aad50f3af92cb10e7fbf6bf16fbd500e30b21fc66102008201527f237ef84b03c0349b8c7a09397d827c5a36de45562c6a61a574e75d5c3c4bf19661022082015283610240820152826102608201527f1a0ffdf59bcb4310a2ed23818d833cd032f6e412b5b8a252b92df4ef4e4c51b26102808201527f0c468ab825ab4e5a1149fd74a54a945a57e266be74be00975cb61eb2babc968e6102a08201527f14a67f28a957603c39a0a9a97f40871d8c004eab0722a17675851d0ad7b8f1d96102c08201527f1707e5a7e714cc1880a630fc91f81afc7d8d3c952566971aa83486cbf656a16a6102e08201526020816103008360085afa9150805182169150508061033f576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5050505050565b61034e616a1e565b610387825f6008811061036457610363616c62565b5b60200201358360016008811061037d5761037c616c62565b5b60200201356157f4565b815f6004811061039a57610399616c62565b5b60200201818152505061040f826003600881106103ba576103b9616c62565b5b6020020135836002600881106103d3576103d2616c62565b5b6020020135846005600881106103ec576103eb616c62565b5b60200201358560046008811061040557610404616c62565b5b60200201356159a2565b8260026004811061042357610422616c62565b5b602002018360016004811061043b5761043a616c62565b5b6020020182815250828152505050610483826006600881106104605761045f616c62565b5b60200201358360076008811061047957610478616c62565b5b60200201356157f4565b8160036004811061049757610496616c62565b5b602002018181525050919050565b5f806104c7845f600481106104bd576104bc616c62565b5b6020020135615f6e565b915091505f805f80610509886002600481106104e6576104e5616c62565b5b6020020135896001600481106104ff576104fe616c62565b5b60200201356160a3565b93509350935093505f806105348a60036004811061052a57610529616c62565b5b6020020135615f6e565b915091505f806105438b610a61565b9150915061054f616a40565b8a815f6018811061056357610562616c62565b5b602002018181525050898160016018811061058157610580616c62565b5b602002018181525050878160026018811061059f5761059e616c62565b5b60200201818152505088816003601881106105bd576105bc616c62565b5b60200201818152505085816004601881106105db576105da616c62565b5b60200201818152505086816005601881106105f9576105f8616c62565b5b602002018181525050848160066018811061061757610616616c62565b5b602002018181525050838160076018811061063557610634616c62565b5b6020020181815250507f10b6c381ef69d708661b933885c1f66e08724d9001c135ee1c7b1b0da08498fb8160086018811061067357610672616c62565b5b6020020181815250507f0452d2815cc5d0d4190d27ebe866886a81af0042d8a40091286a160c05efb8c0816009601881106106b1576106b0616c62565b5b6020020181815250507f27e5ca7d3c6a0485578f48e3c1bd5f84e2b32204b40d5bf7f5141844289044f181600a601881106106ef576106ee616c62565b5b6020020181815250507f124f33a5c7f3ca32aaf25c81a4fa5932cd8bc08b09d40d2d1027aa72d5c84c9a81600b6018811061072d5761072c616c62565b5b6020020181815250507f2d686e8598465a029d0f854888dbf622ff2be2a06748ef10db44f579b1426d2181600c6018811061076b5761076a616c62565b5b6020020181815250507f1f7ea897a6112ad5bc7f8765815f62d3ab0d4547763865000bac252e71905ae681600d601881106107a9576107a8616c62565b5b6020020181815250507f0a8e869f6f733346e8733a4d3ec42d86e72251138bf945b566b8fcd3c085431681600e601881106107e7576107e6616c62565b5b6020020181815250507f0d5c64d136a42af58a77d8602c8412f4a1e1e80567207e8e6cd07c661e86447681600f6018811061082557610824616c62565b5b6020020181815250507f1c24c1e49253789a3434be06aad50f3af92cb10e7fbf6bf16fbd500e30b21fc68160106018811061086357610862616c62565b5b6020020181815250507f237ef84b03c0349b8c7a09397d827c5a36de45562c6a61a574e75d5c3c4bf196816011601881106108a1576108a0616c62565b5b60200201818152505082816012601881106108bf576108be616c62565b5b60200201818152505081816013601881106108dd576108dc616c62565b5b6020020181815250507f1a0ffdf59bcb4310a2ed23818d833cd032f6e412b5b8a252b92df4ef4e4c51b28160146018811061091b5761091a616c62565b5b6020020181815250507f0c468ab825ab4e5a1149fd74a54a945a57e266be74be00975cb61eb2babc968e8160156018811061095957610958616c62565b5b6020020181815250507f14a67f28a957603c39a0a9a97f40871d8c004eab0722a17675851d0ad7b8f1d98160166018811061099757610996616c62565b5b6020020181815250507f1707e5a7e714cc1880a630fc91f81afc7d8d3c952566971aa83486cbf656a16a816017601881106109d5576109d4616c62565b5b6020020181815250505f6109e7616a63565b6020816103008560085afa9150811580610a1957506001815f60018110610a1157610a10616c62565b5b602002015114155b15610a50576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505050505050505050505050505050565b5f805f60019050604051604081015f7f293f591a5dcecada3d8151843a24456ae43bc4ec8397a260fd1eff01090aa47383527f1626cade05032b19c6c42b811f2fcea236f124828aa94f87afe3fa50c019384c60208401527f132e5f51bfcf1ad13e4f12896dd7f62c3e1c5cc0db85304c32242f48b84466cf82527f25a65800920dd2d5c3fdd889057ac0775b8b5ae73d76bbb9e0c6ae4f43f76b1f6020830152863590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f18a00ae787c28f7b7e5d2581faab4c25a8d72c3b4f982063414a7d90672e960682527f2c85d4d10542bf8978ca1e05f87bd1c4931131aca28fbcb7f68b357de91101be6020830152602087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f150a2cfadeb18d487e88b4dc35d5565e29f046adce7893a4e7712af250c93c2c82527f234ad385d815a3b6f5b278025c88879c84763ea7dd2df09b1777ce032978209b6020830152604087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0a0c437a693c25b2d10321a15b2d1f60cb85dd409a214d26a96ac624eb845a8182527f2fcd45164ce9c7c794f2d061a5f8b6b8d2d215727afaf63b20d22ae688917cce6020830152606087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0df8e475bb538e79dbf5222afc71867e9aa74469b5fb8cb9176795fea3869bf382527edc94293214605b839b0175d272ecd7f666dc4871305f0a0066f35065226c476020830152608087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f18da5c67641dc4374c48bb61dd584b4a0ba42d663613a12d077ce45c3c590e3382527f1228f70a2129412149dc1fd8b92535c0cc40a362ccaad9efd1bbe963673538e8602083015260a087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1fe6e902bc6dee654534b77cf5dbb6139a0a5c74087aa20c7f31c8f30408e9de82527f1b3bb0fed9b42d453602f79c3a2ea4268603dde16ce77b881092f3e30873f54d602083015260c087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0415f3eb44f91337ac778698fe89716d383054cb24be93bb813c51ff81cc95e782527f18f866cc1841ced310762df9653d5392fe270481899eddaa82fb69ec564f382b602083015260e087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f19e426f5e131ef280e478a11d732e8c43b1c7d263c1d827402d16742b10ee73a82527f143ed0f7f397fcd9335291e1a3eddd177290f3c471db5b51339e0e4ea2cc1682602083015261010087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0ab654db0ffc09bc92c95e055c5470c3fe243986500c6d11adbced940585a59682527f1edd9779dbf658096880292e362ebc67f7b8ae18eeee7167db1ae375092a2a4a602083015261012087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0c2127bd09f7038b3031470e12c70b277ff01e0c11ad1fd7b9ba267274e0df0b82527f180c65ca8f55dda733f7260942a53f6440c58f0b4b8d0ae443215b68de14c187602083015261014087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f23a6068881fae5127fdaa00dbb371a34864db57849d609b1e8703107ce4d263082527f06a99cc4a28ac16825d02beceaf784985023b0dca1fd561c4f05c766798abfed602083015261016087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f043610e15160510516ced4a8425bace68ec94f8f342d2b2cd9545b24caf266f782527f17e2996ebc4af4fa67d56a01b6801d41aa39c8f51c28e1d16b8cd74fd991eb04602083015261018087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1a51f6fa7678ba8165a04e8a4a778c4606d9c570d86d88f51b5fb2b46ab9247382527f26259646c88dea2e1736650ec4ceb85105a76253e5aa236c11e93eb99feb7df060208301526101a087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f06b7fe9879f2ab88372ac91f1674deb19441a6939043e9f2dc84cca1ea1bdb9282527f2055e2eeebf4e5fc8fddc61c7eff56ec2ce6090fabc23d53b27402adef21e48960208301526101c087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f16b5561644d9c17b4ab870a45df313cf8f8a375047ec17a4f7631066f01e4cbb82527f1d265df96e5ca23ea8a68884ee7ee9a96e423bda34fed203c48caadf518883e360208301526101e087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f22aab6b6b0d00eb733013d19492f74e5eb95074cf9f8c1902484366f0f89d4e182527f1a4bd6efcc5f7d2c41ceb19d5c3801aa951149b260f5dd1211b93d656d569925602083015261020087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f19367c09d4b7c5caeb6ee60af0221d59fb73eafb67e188476cc96b533af6f01a82527f0115d668dc8eac6dab9f3e3571f4ed879170f0aa01171888e557fccd462722f3602083015261022087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f23aafc67a4772442ce06cefe40a50bf0c805b556d782994e4ddf5f09b1f0977882527f0d53731c40675eee521b4adf3e1690a10c5f00498761d7378dd55d414efbfc33602083015261024087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f23c3332e89923429426892098d4a5d27950467d26dd696296cbee06a454ebbec82527f041e8e14a7a32e07cbd453db0a9e21c2d95fe079a44e414e8ea825a30bad946e602083015261026087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1e0f6ac255859f4dfb0940e1e2166c805d051785c61368dcbaff2a6ea8d9c62682527f05631977cccb328b0a307ea6def8af68a42f28e37f59b6fc9c07befe3f6f39a8602083015261028087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2b741590f995e9f7fc4f25656999e6c36cb81c49b7e6c67469749744733414b882527f25f375785f8e3988bb3f8f1b9db4ab782f4d803f015b51233ccf391493874cde60208301526102a087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1752cf957d78d62d34980d1f6718cfbd774e63db66f08deff76426d0271fd76f82527f29344f79e47f64bc2766d7144e9d9a748e799aed068c9c195427cc1d207d314560208301526102c087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f27034afafa83b6ae704111ee45883cd433551baeae992b944ba688baf9c9cf7c82527f0f4704e87f39d0fd410da2dbaaca59889b690199c11be9b5b2be85910caeba6d60208301526102e087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f255b9cbd6137fcf2a0e531c67b74cf28a22bb7530e9ac1a8b964e57ebeb0efde82527f2a74c5c6c8a6987cdc8e7bac5ab3e10434620718b7731c9af9e68e3acc3ed61c602083015261030087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f216de622bc7672b1f2e84f199f4a36ec1fa29ae525e976a601eab24152c8fa8e82527f26d944801e32e1cb78957c3866ba8efba061597771c20ea9de68d91c54e28751602083015261032087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0234136e7d831d7ddea29fe58d520315f70bac65a1d85e1cbfbe2fa302c9f74682527e80ec26262c4d465b1fbb3c8a3df5af08ff689a16f0df5eb7a8f3b5a9d4bfe5602083015261034087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1953a542eba13fbc08f5f4dbd7e101de4a8d4c90efb6e6440fac00d36463b66b82527f099eaa24d252c150bb4d4f97246dd4b5f4f5792709e31c0ad18fbe1d2fa522eb602083015261036087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f230ec0d70f9e4f193cad78f014ab990f3998c626171295e72de2c6b1ffdddc5682527f09e2e7b30cf76f6f2ea7b53957ecb4abfa6e26ce6551eff96e9eeb660923e3e1602083015261038087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0c134f0e1c999cfa8353041f82d4ec1afe9ad8d09ba99cadc2e5840eb876f46b82527f0acc28248ed99ae8426a4ef4c0838da4a371b98f2b22b23941b6ea0ef2e0c43560208301526103a087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f279cfe0255ffaf9659756a5cc76c6a27b6cde9aac49aad2a9151366e44cec5a482527f0bc47c474584e888a9959a2d3c38b42dc0cb5fc5da518a24212268b1838bf40060208301526103c087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2fef468589c84f80f0190c6dc4de1d3c400e2499a955a37407accddd5d863c2082527f13e9cf6f2882a01c2fccfed0bbbcc47f0494989f55f5cbd2bb9df0e82a83bbbe60208301526103e087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f199ac13305889a7df1077b7327d2ddc9ad45dc7a8aa2d8cb74d16a259b80857e82527f088309c9a4114dc67357b1109c954b1d1b29dee601dcc64772b4115523844158602083015261040087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f25f6954bb8ee4103bde18f6dd825aaf6da26d46e9f639f0656113c759dfb3c7482527f1bc1c1b90ccc5aac1ebd142344a29fe11215755ba31bfd8cedacb090d876b561602083015261042087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2c2a3da55d5a2bc31b34b8f7a645034a629c8df355698f2b2ada4f4430a1aa9682527f1e0701a9abbba97591fef3ae09520cba2558c86d43e8e01e32f8a7e502f015d7602083015261044087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f06423f6366ab058a680bc7cd28c2d0056cbc2e09162200d868535f4c6223c52d82527f0863e78cd41fbeb73e5ebc308cf2a982f5b8c7fd2a656a7ad3487a356999a8b9602083015261046087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f29e6ee526492d6a6622162b9a01dfe2558db4fb22482f4242cf711f912e0b41582527f06613d137ca8b20cd7adda94bc12ee1849d9dce6e97d5be9b2a65e66b845d156602083015261048087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f109d9b6f2f02d132bf52fcdad1ec2652768e9e770d1103a87836bc300ddba6c182527f2810ea035b2f71cc545f9e4cf9708b3e9ef2d8cee89c6b047870ff9b20070cfd60208301526104a087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f21372b365480a83aa993b5de8cd93c1c0220f2c1748a92faa4ac5e4cc353099982527f1c3a9f084654797dc8bff74e5b03e0bfd1229964e0d538da3025d940dfd14f6a60208301526104c087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f25f36eeedeeba791493dc9798c0948b182b7497be2072372e5e65f7ae5da27d782527f29240f4e09c77431089ff7a867de934a55d449e7fb16bb9991001d3f0afa3d7d60208301526104e087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2e0494d37e1befe1825315893c0600ef4d3715aaebf1099df7072ffbe10169bb82527f2c9d439dc00926045e4248c5115636b7b600a3e2a8b2d6af7de878ac390b5854602083015261050087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f199aeab5c0346b09f871f458bf1e3e50491afac1f7f990465860f590d391fe1182527f15f8ffeb1efc3ff97062b74e7fa9f3b4c8eb3a1d308c2ae4fc863ec80adb854c602083015261052087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f14c0382cb9135e03df333c6143d38d251382e6466e01baefeaf6aa966dfe530c82527f12fea330dae281119c756ac8ccb72b98ce3799a7bbb9d8028a1c7c30acbf29d6602083015261054087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0b7c35dde9b827f51b8bfeb844149b2c63ca7cafe9f550758e7b149afe79b32f82527f1c29b53db5a4c41ffeb979690d746639c71f22ac35da5eba14fd6437f8faa104602083015261056087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2c3c6f7b099d6048dcc5e6798a23228bbc3c1d7f66551820f7cf770acc7a13b582527f1f142938e43f63fb923b3e88e45ad39f70f08b7a6a60c4e207b736265174de40602083015261058087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507e4710976640633fb8dfaf0eadfe4cbfa9369cd68614b5df5a51ce91b7c9e64582527f0dd433aa0b53c52621d771854dc9578d67f29447029093dba8cfa20d5f66a72760208301526105a087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f05a58d8f06db3dee1538623a593f42f3368305cf96e2d4f9f6324cb02df1ef0782527f26e662bbf3648c5eb061ef28aa73522e93c9c984a39bbc18f7f4956dcd05602360208301526105c087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0a610107ddccd355d3141f37d7638725b219d5e00a12ad6fc671da0da2db85ec82527f27b4e3505fe7d41ff59b7471ea1dce902884f1513f9ec6b4329733589a6cec1860208301526105e087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1b4bed93c2f04adc5a2679bef62dd88b19a58fc6d0502538ab7685f334a62b3782527f206cf1e266f169e7c79dd217576fe6bc16844517a22d95e0333e9c2741a8c25d602083015261060087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f270158b004ae7127f628ecccc95f0ba1c013b10f1c05c1305cbc3a51d5cdaa2a82527f1fa29e5b965deadf77248b83d598407787ffd4becd13c95ebc459e086a7f425c602083015261062087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0c30511fe6c315a6dd039cc1ed95744f569c2e31f230fd6438aafa4686b64c4682527f1cd37e2599f8ae245763011c8dfe78730ca9b066e551e26094d08f47170b49cc602083015261064087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f14a47e6515ece69b69513a07a6318c5490546cedb3248e8e01df4b3586bdb0d582527f022fc6cd8403e7eedb1aeb3c29183cb1f6d088a0032b66da5cedd0ecce746ce8602083015261066087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1886c760d4ef4dc9b212034cbfc3006f0ac18b3ce41ea97617332966eefbc07c82527e60642a5f3fa930481eefeacd0f3e7cbc6001f51b960562ec5ea55f06490b84602083015261068087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f302654cdd735516a487249c4526df615e93e70256ba46449cf7ce44ca40d979082527f1a9e4790cee0af6e69e4077aab53bb82ab5ab1481baef4e2c196d074db5cd05b60208301526106a087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0fb8da9b9b8f5542aac2b8cc667434e0b2d73742f2d074d11e37beea0edda45f82527f18a80eb4132b2dae8c679d925e734ab377056fd9bcc8a97f93b253950c515be660208301526106c087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1abfe288e33658028967c8f34598f35dfbb0c80d03b886d8fd37f7d1821c6e4782527f1069710517e3e7b787f6025017ed3482361ce10ba54d17edcc55ad492d84aab060208301526106e087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f14fa38022770ecc66bb2101c91ae98998a91ec1af943a281df338e60af14d0d682527f0fc937b124aaba57c0aca0d56e55c980a6030e3d64096cfa290ae8ccdd06e17d602083015261070087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f07555d8fb5e83252af527806d6e5a73a73169075087824bd22294d415a69651382527f11b81a6a5a517caacc246bb4918fbd76c49931c2e9fe49833cac212d7314213a602083015261072087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f033a220c02716f5106170cfc7d5cfd2a52164e13ed05588f28c6319c8ea7088982527f155151cc3a894e17bee43b9233b1b18415367332cb21eb419ec5e5d9394fa914602083015261074087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f18833e968ee3b41697308b523384883b574cf4bce8017ca2895f2417438acc0d82527f0f217046734e70cf4d732182a4871f263399c54eec03d880ec59f8a467f82a7b602083015261076087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0b3e157165a552e23ce4ebbe022a8422102a507320ef54595a01bccff559156d82527f06c016e99fffd9d79fdc59efdc2f12e0f96553b875581e31d9e6b5e92ae4995c602083015261078087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f046092cc5f89b17cdd149051711153fb08291176c416e11b0480cdacaee756f682527f138742c5b57d12a5258be6c17fdd761d213f51d1cbb0a25c4ef6fa994f91ffad60208301526107a087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2a3cbbfb93e037130d04005b8adc876e1951e7e18112664f9e254e668da422d682527f045c2885f8b459e804a6dab5e3340f48ca37e8210c14d9e82dd3f31e2424d74460208301526107c087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f15988ed9890f4a78c7c2d3d1ad7304b463af6f8627a7caa9b17384a8b493363082527f0b0d24ff7b36d3a4f4e39b858c2bb395a367751d4b46e5b47d688ceeaed59aca60208301526107e087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f092a864f91d9ebceb463b4db5ebfe14427441e3045c8685062b7708955459a9182527f1c84e5b71b94f26e64533b6039f92a27248d30a29aea8d5dfbdbd8c1456a1f66602083015261080087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f041f592b25c235f7cf686ab6715c811fc9e51f93df4c2fdffa7a9d731ced05e582527f09120434b6d3485ef1f073b58ba28b242639b80caff119dc1e12920bec006753602083015261082087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2046bacb205af340f6c5f6f8b6ef0825bae4920735840ad3fcfe2c70b5caf9d782527f08b4262c6d08cc0fd8b697ec45502b5aa5b2f1c8fdf39b7aeb8efd6a4825166d602083015261084087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1510b3d309afca2677e160c490dca27e9828831cd5f16c435a0c6880bc01d06682527f218fca3e930c336f69c7c90c903f84e1dca50bbc76ac578af14fcbca99ea8ad9602083015261086087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f058bcde29a9601058de32587dfa58638de22fadacb3e776eb3ed2d0366cddf7a82527f0d522fd8635d1928fa55751901676673dbdb505890c4f542d87c98ca3c354595602083015261088087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0ef24d25e5bda320b0a58689726e29fda42ac53a685c217003ac0abc95d4c65b82527f23aefc6fd90ab5a6d0883b923085e6dfbe4fa898262c9a48336aad5517c95d6760208301526108a087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0d9021e176f39eb0faf7431b541785b58883fc4d30ede26577b4f7dd1a52169b82527f0fe6ff001df356e8dbc80cb2c01671748e5bcf81f0b170f79e19fa6101dea57760208301526108c087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2510288baafc264ee3cacc3f183df909b67336a5872bf4321a51524c581c5f7482527f03725943e2419833ad7ce858dcaadf6a012d92d95e8cd7055e4319e8eb1374a560208301526108e087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f101b3ae7aaa875c8abea52499393311d84fa7b7682b9ddb4b556fb99bf8ae89482527f02566091b48312735d90a8f02fb3e839bc4635072c8740c30912ad729824b9c2602083015261090087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f05d6c3f592453eca8352efe10bebc8bc9e7d90de34b74fad5f8bef8c52c6803e82527f26d8ee6c9a85b2f846f23283553517bbf17f8716c4b18b2a8571042936367847602083015261092087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f22606faa39d4f4daf6065f07e87853bf62e57d6a8d0cc2450568bde70e3010c882527f16c3df6df913dd2c3113300f83b96ffd60dd2359bb9cdc5199c6f0a93d83cd1f602083015261094087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f15f6df75d738058e5c9dd31c582c66a68053e19d73a09090a2de34e92d3ca57382527f1a0f17cbb374721c2a5d0fc6f4bf0645e6a43aca090698aa21969c7cfb25241c602083015261096087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0c74eeeee33b5b1bb98e94b2b146dca805e569fe5aa06a86896d5e6e23d7afce82527f09202e3b8f07907941e3de14cc8921b5a56c7a59e2d8aba8a02ca7e3aab703ed602083015261098087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f18eb856474e71a87cec1c3daa9c8843815ac6e7accecde11ab200424a32f9d6282527f1679055c9fd8c1400a05eea28b47d103e9da66d8921fd7f3870de6250348a28060208301526109a087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1eed2d7bf55d74b77bb5ae88d998641c8ad28730f975af69c49eb41b4a256ed482527f2283d3b87420766f6b00ce7d59ecb31bfabbf4f724eb0b4c7a34d68c4ea80c5560208301526109c087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2de1a67cd96979355022edb9e019c8678ab49dc9b85c002f09cc9da29cc9705682527f17f1815b83739720403a53c57d8c3b060f57be0a3a8e2abc56ab50ecdcf2cb3260208301526109e087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1ccd0a31e80c547de2e83c12d23987e01b8efef33bace2b4f179a9d4583424fd82527f2aa346713d320702ea399d9e53bb3f302518c5b0906fdc02a97109775578f6ce6020830152610a0087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f012f60c67661880823f6f364815bc9e6084cb4c5ba98aa8b554a1a8220f14a7882527f11060252ab8acb5ccba8a82cd27443b574c71ecea136bf59410ba25c5c817c3d6020830152610a2087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f14395f8b19c1d2d4376a0abaecf7520075499892798d33a4791c082cb0bdba6682527f1fc31b478ad394c76815fb5d87b6dd14b37b5275430bacf047500cc0863b03946020830152610a4087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f23a7e95d9a8bcf3151dc1799da80e9e93c123bb6e1ee92e7c24dab747fd5261c82527f292b564e67a83eb43e3c001e9a64ea10143f8aad11ba13fd19bd8a3ee5021f4f6020830152610a6087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2fa49183a9a81ccebfcf70eb5e302b1ff974cd52d7e4f00c5f60778828eab33582527f01155dc7ad68b21da7d64f5936e17e0d68954725416735083861174b70a894896020830152610a8087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1d26d9fd881dafa623b9bd9e8a7ba10625cf1bd016912bb2dd61c92175d73c5182527f2613f15ee4fbe01a28b2da9c90e333cc634e4ae5ea2e41dd5d8814360fd5e67c6020830152610aa087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f17949d2091e51d7c61066f2b70777130669e66151d01e33313e32eec2899c87a82527f29d4d36c3fea308c8d53d36d36d0a5978dd5ccb82ee4f500078fee04b22537266020830152610ac087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0f4aa4d8a8703fe23d479025a25a5af7cde1d91df9a1ce2f74b8f20ad5aeda3382527f21edeb5dfebc52ab5af50f3229b1c0abfde837e83aee1e39f66ada40ba002acb6020830152610ae087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1bfe72e15ba8caaee0af70d30ee0ae8b353e0c143b1da1a594476490bdbab6df82527f208f72a3d63473f793e5a673f35fcffd91d4a35b704f7763e337cd989f3f9ce96020830152610b0087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f11e01eed812b6ba1a4e59fb7fc16f2c5b2679f7ff665b03a02595ac950f37e3d82527f0b2d3095a20757c50a1727afe7e42911b36ae05ab5823608c3b1a3a5ed49e2756020830152610b2087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0b8739ecb0e946d04af069f269be8249ad696af3bebc10076e853e1a8612c93c82527f257ec0180c862c9eee0462c4e1281db727feded3850c1b299114e33af2c7ded76020830152610b4087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f25189673a5ac215bfc62dafe30bf4f78b6ae2da82c8790cf7e28db31a1c468c482527f26ca90551ec5b4e5b6169503c2bd16373f6f432c355b5b65ee27e4301d583b556020830152610b6087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f08fb36d078db8433706be424bf838b25712724c63bd29d5ae7120621a34b703582527f0466326f8c029d8e70d7f75ec3dc23c4ac4cff53fd9585d5803ddd0e4d57712b6020830152610b8087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1ae13e15abc6f94e674447316cdf7c1bfa25617de5d680d1af7bf0be95150ac882527f0ffa8f314ecaa8929df2f41be706e0d38a5a7ec5661048fcad139dd28b201bfd6020830152610ba087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2090e734ed9a247092d022700c6218d0e5caf71601c53bf27ef25a2249493ead82527f2c8060279c1888ddd2b36713626fa33229d8b3f9ca14cfd178dcc73ce8e726f36020830152610bc087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0cdb3e038d50ab8843c3fe36c77412ae698f0ee6c8fb2db1d1d140313a82726882527f2574f9461c297b4afecfec6c1d1440245fab11ff655399113d9646dd114adae66020830152610be087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f253fa6574eb878475c38c30582a10c17168da5ca22fdc83e4be52474dfa93f6382527f159d5d803364182ec961930854da792c605a8b3b9c44a3c07b4eee3eb3dfc1ac6020830152610c0087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f294c027c99e5e19aafc8498d03b1839d8bb72408250be179b12195cf3ca55d4f82527f25f36957440394e57cc3bbafb6490364008333bc5fbd523904f83a88025850066020830152610c2087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2c7834e458171f543408418e8a1d21ce64a50c4a09a63913107069ee3a6df93e82527f2d467f0cebc08b4219b8cae0e5017a5fedf209d338ee59b0b04fb9ec14486ac06020830152610c4087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2485f0e00dc4054ed001e72869229a296adb71787c2c2c9ed18fd2496cdb24a682527f1f743743390ed578d28516d80b3ba9a40e3f2103d5d47cc2f1e83e7fc8b8c6456020830152610c6087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1193988ddd1c7daea31f876393c0318ae4c2d9564e8256831579628c5376bbbe82527f0919372a5a08ea340bc66103fe503b7af5de7ceedd90a331994d10ccb3ee398e6020830152610c8087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2aeea117ebc1669afccf6d66a45b657a53166285d07fd6db22b7a6a66903c35582527f0f4c754ca8ff15aec90ecb0d57dbed2fcb8162bd5b729f3116b819f050a766ee6020830152610ca087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f182e949ebbba0d303221ab2deb47089e949c86d1487a5e5a8229c1348cb58b8682527f275202967d9ba268cdc1449a99e4f844fc955f5b399c16fc8fc62e0301c22ce56020830152610cc087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f153c5925b256b55a67992854a91615c1680017d6982a8fe1d2da10c89af3d37182527f1172f5ecc357a7bd7770bfc8b9ed6fcbbbe18c8eab8248357e0050227a836b9c6020830152610ce087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2e465b7e94d79f4c670cc170cabea0807aac9b18cae40c14d9860db0b7b0441f82527f0d9f096c275ab28d1786b1617d25105eccb171ec901230744aeb3cfc27353b5a6020830152610d0087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2f2845ddedd5a7976d3fec91499a2dcdfe1e1d292fda2f28bf5288478bda8c7082527f2d6209fd0ab7b2a46999c7243ae9971a4758a1fe626f46971d35a9fa7ff6da496020830152610d2087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f26ae2dd21c717e002b5825357c57a92b76f7cd97ba202385d53ed03c036d7f5d82527f23cef2348c4f7e233d7bfec32a91f7e545a839d0830816d6d40a35141713a0be6020830152610d4087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f05f31cacbaece10af1bbc9416751b4f277ae1accdc288c6f86c6e6c8a030012d82527f12365fd05e7cd73b973afe53ae14f5b5170cfdcc8b9d6eb64651ede7844269256020830152610d6087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f182f8516b25ac2d70dc6359b10de1c1fe99a4a250bb73f280c529a71729a2a9682527f2273253eedb549df93bd00e7b10c8d247208c2456c643055d043f72ffcc6be136020830152610d8087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1f63029d8684630346075b799e483fdec640c62e0a7bdbb694206b2634dd1f8d82527f2aa5d983ea3f65066df9e5b089f647088bb3105603539e5c016cee6ec9c37b166020830152610da087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f27fb26670c3abc286d2d6d7ebccfbbed6c994584b52d58c9e384d85f0c0003bd82527f1c4a77a8112663436fab5a59d79eb6937a5b6f40478c3615f6057c44a8175e356020830152610dc087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2a690681eea0606d2a1837b71960b6f541fae5e2e3dd721ef87806151db6c87082527f159d25cfa83735a4a2e5a99199152e64ea32c633cbac18e36edec1f3ee1e52776020830152610de087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0b9ac7c20562aef3912aaafc8938b3436c3fed265f5e458d5c634d53be7732b082527f16c168abd90456cef6ab1a48fa484a6402d30c2f52d6a79d0516e60ad924258c6020830152610e0087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1ac968c2cb22cc76bec1e900a57e24f7205aa056c8c860d16d11b84429aa80e682527f0cc5f5486dfa43475541b6bd031c3e7408596ad5b78dca60ada0561976697e066020830152610e2087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0cbbd1dff03d083283919bdb8a2dbad0355a3411123e6cab71494966d58b0be582527f058d80ccc9e11b50879e266fa3c9f5e95867cdf6094a0f18b3f7a15aaf3453346020830152610e4087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0542195ad5e6a5d5c1993d07fdf695406e208012a18b097e3c440a8f91feedfa82527f06dd4c515edb71614efa7c1909940c7967a43b4e3b2fc0e39fd22173b75ad0116020830152610e6087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0d424b62ae3c595f4fc04372c1c456fcdb78816643dd2d683af444004192cad182527f16cca6dc7134eb3dbcc79a70683d48e6c14d27e01417636b09ecfd3f4d4cf3f96020830152610e8087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2f64f5470d7621166f3efe6d8066a5865dbfbc9f60101cc09c6f2ce2de5ad0ef82527f1bb6fb34d3067b6694b5bd75f6e0d8bf54b7165a3bd15e5462cc46caf33148576020830152610ea087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0db67575b0eafaff7cf1fe918e5a1c78e290c767b741daf236f6a4ee55fa260582527e24efcebe815922f1dd1157f7f09c57688c73922ad5de2257deb207e42170b96020830152610ec087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f15b6a24861eccc95eeabec6060dc7a07d42968f2ba93142a5d7e48873e43f19682527f0342cdb58dfdf2dd0f9d6c814350ae46b7b25f5884e8d1e70c94a7f4650355f96020830152610ee087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0847ba9dc1a3ff868dd27f4012eb458542ca58498b9bdc3b43469d216dec4a6f82527f10d4110cbc8a788dbf3d1f75f6a6d198a4606dfd9921f7e532e3093b5380229d6020830152610f0087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f197c8acaa828897988fe349b951a47d249eb753100aafc23de4de4ea11858f0282527f2866ec3a3243b10ee1271189e20bbbe6cd9d64568d914519ff04d8294a51de176020830152610f2087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f0bc04a377c5e260656f6e98edf4958523b1f38a43e38ba4705269d071a1962ef82527f2dc0c1a0824de1ba8588eedb099f153591fb7e0567a04267d069accd75494e6a6020830152610f4087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f1d2dc9dd8049712eb953670a2de54119c398e99da95238f07c421f54dff6ac8082527f289e1607164ab3f0582e6c718684dcabdf1b9a78133f4353628be9591dea10296020830152610f6087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f2eb673f8ad8bd44c97fdc4e2805f92088dbc9e7a4895f3a232fc8b401bcbda8f82527f15553e43541a396c074ce94ba7f36cf24405e2ed26d97fe0bb7fdb963861a1b76020830152610f8087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f065dc91e86bb07b3aa71609ab93e67a3215dfdd90fe80294d843fa5ac1b26ac782527f0aeda45ede9c469a5aa1ab328d1bba362c7b6725bda9117ef73e8b87dfd190f26020830152610fa087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f297bc4b5ace73461e817e0611dde224074858c6d2589753227e029b6591f3d0282527f25c3344d89492dd73be547f2fe380203f2186d1549619f66dda7aeb015bfbe5a6020830152610fc087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693507f22c96a80b29ebcd4a4d93f33931b87e456009190db8505d11b0d12f0a81213a682527f06d21eb660a048331ccc2f00d2fb1db540d7592f307876708211f93beb0066466020830152610fe087013590508060408301527f30644e72e131a029b85045b68181585d2833e84879b9709143e1f593f000000181108416935060408260608460075afa8416935060408360808560065afa841693508251955060208301519450505050806157ee576040517fa54f8e2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50915091565b5f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478310158061584457507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478210155b1561587b576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8314801561588957505f82145b15615896575f905061599c565b5f6159347f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806158c9576158c8616c8f565b5b60037f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806158fa576158f9616c8f565b5b877f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061592a57615929616c8f565b5b898a09090861647e565b905080830361594c575f600185901b1791505061599c565b6159558161651a565b830361596a5760018085901b1791505061599c565b6040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b92915050565b5f807f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47861015806159f357507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478510155b80615a1e57507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478410155b80615a4957507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478310155b15615a80576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8385878917171703615a98575f8091509150615f65565b5f805f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615aca57615ac9616c8f565b5b60037f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47615af79190616ce9565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615b2657615b25616c8f565b5b8a8c090990505f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615b5c57615b5b616c8f565b5b8a7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615b8c57615b8b616c8f565b5b8c8d090990505f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615bc257615bc1616c8f565b5b8a7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615bf257615bf1616c8f565b5b8c8d090990507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615c2757615c26616c8f565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615c5657615c55616c8f565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615c8557615c84616c8f565b5b8c860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e5089450615d6a7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615ce057615cdf616c8f565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615d0f57615d0e616c8f565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615d3e57615d3d616c8f565b5b8e870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750861651a565b93505050505f80615e0d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615da357615da2616c8f565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615dd257615dd1616c8f565b5b8586097f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615e0457615e03616c8f565b5b8788090861647e565b9050615e9a7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615e4157615e40616c8f565b5b7f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea47f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780615e9157615e90616c8f565b5b84880809616585565b15915050615ea98383836165ef565b80935081945050508287148015615ebf57508186145b15615ee7575f81615ed0575f615ed3565b60025b60ff1660028b901b17179450879350615f61565b615ef08361651a565b87148015615f055750615f028261651a565b86145b15615f2e57600181615f17575f615f1a565b60025b60ff1660028b901b17179450879350615f60565b6040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5b5050505b94509492505050565b5f805f8303615f82575f809150915061609e565b5f6001808516149050600184901c92507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478310615feb576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6160887f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061601d5761601c616c8f565b5b60037f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061604e5761604d616c8f565b5b867f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061607e5761607d616c8f565b5b888909090861647e565b9150801561609c576160998261651a565b91505b505b915091565b5f805f805f861480156160b557505f85145b156160cb575f805f809350935093509350616475565b5f60018088161490505f6002808916149050600288901c95508694507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478610158061613657507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478510155b1561616d576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061619d5761619c616c8f565b5b60037f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd476161ca9190616ce9565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806161f9576161f8616c8f565b5b888a090990505f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061622f5761622e616c8f565b5b887f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061625f5761625e616c8f565b5b8a8b090990505f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061629557616294616c8f565b5b887f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806162c5576162c4616c8f565b5b8a8b090990507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806162fa576162f9616c8f565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061632957616328616c8f565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061635857616357616c8f565b5b8a860984087f2b149d40ceb8aaae81be18991be06ac3b5b4c5e559dbefa33267e6dc24a138e508965061643d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806163b3576163b2616c8f565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806163e2576163e1616c8f565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061641157616410616c8f565b5b8c870984087f2fcd3ac2a640a154eb23960892a85a68f031ca0c8344b23a577dcf1052b9e7750861651a565b955061644a8787866165ef565b8097508198505050841561646f576164618761651a565b965061646c8661651a565b95505b50505050505b92959194509250565b5f6164a9827f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526168ea565b9050817f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806164db576164da616c8f565b5b82830914616515576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b5f7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4780838161654c5761654b616c8f565b5b067f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47038161657d5761657c616c8f565b5b069050919050565b5f806165b1837f0c19139cb84c680a6e14116da060561765e05aa45a1c72a34f082305b61f3f526168ea565b9050827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806165e3576165e2616c8f565b5b82830914915050919050565b5f805f61668e7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061662457616623616c8f565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061665357616652616c8f565b5b8788097f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061668557616684616c8f565b5b898a090861647e565b905083156166a25761669f8161651a565b90505b61672d7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806166d4576166d3616c8f565b5b7f183227397098d014dc2822db40c0ac2ecbc0b548b438e5469e10460b6c3e7ea47f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061672457616723616c8f565b5b848a080961647e565b92507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061675e5761675d616c8f565b5b6167997f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806167905761678f616c8f565b5b60028609616981565b860991507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806167cc576167cb616c8f565b5b6168067f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806167fe576167fd616c8f565b5b84850961651a565b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061683557616834616c8f565b5b85860908861415806168aa57507f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd478061687157616870616c8f565b5b7f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806168a05761689f616c8f565b5b8385096002098514155b156168e1576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50935093915050565b5f8060405160208152602080820152602060408201528460608201528360808201527f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd4760a082015260208160c08360055afa915080519250508061697a576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5092915050565b5f6169ac827f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd456168ea565b905060017f30644e72e131a029b85045b68181585d97816a916871ca8d3c208c16d87cfd47806169df576169de616c8f565b5b82840914616a19576040517f7fcdd1f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6040518060800160405280600490602082028036833780820191505090505090565b604051806103000160405280601890602082028036833780820191505090505090565b6040518060200160405280600190602082028036833780820191505090505090565b5f80fd5b5f80fd5b5f81905082602060080282011115616aa857616aa7616a89565b5b92915050565b5f81905082602060800282011115616ac957616ac8616a89565b5b92915050565b5f806111008385031215616ae657616ae5616a85565b5b5f616af385828601616a8d565b925050610100616b0585828601616aae565b9150509250929050565b5f6101008284031215616b2557616b24616a85565b5b5f616b3284828501616a8d565b91505092915050565b5f60049050919050565b5f81905092915050565b5f819050919050565b5f819050919050565b616b6a81616b58565b82525050565b5f616b7b8383616b61565b60208301905092915050565b5f602082019050919050565b616b9c81616b3b565b616ba68184616b45565b9250616bb182616b4f565b805f5b83811015616be1578151616bc88782616b70565b9650616bd383616b87565b925050600181019050616bb4565b505050505050565b5f608082019050616bfc5f830184616b93565b92915050565b5f81905082602060040282011115616c1d57616c1c616a89565b5b92915050565b5f806110808385031215616c3a57616c39616a85565b5b5f616c4785828601616c02565b9250506080616c5885828601616aae565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f616cf382616b58565b9150616cfe83616b58565b9250828203905081811115616d1657616d15616cbc565b5b9291505056fea26469706673582212204a5572ede2b949f02a3f9d89ec09146abffa4fe4bed980c4e066f0bb2a85861a64736f6c63430008180033",
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

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xd0eff68f.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[128] input) view returns()
func (_Verifier *VerifierCaller) VerifyCompressedProof(opts *bind.CallOpts, compressedProof [4]*big.Int, input [128]*big.Int) error {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyCompressedProof", compressedProof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xd0eff68f.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[128] input) view returns()
func (_Verifier *VerifierSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [128]*big.Int) error {
	return _Verifier.Contract.VerifyCompressedProof(&_Verifier.CallOpts, compressedProof, input)
}

// VerifyCompressedProof is a free data retrieval call binding the contract method 0xd0eff68f.
//
// Solidity: function verifyCompressedProof(uint256[4] compressedProof, uint256[128] input) view returns()
func (_Verifier *VerifierCallerSession) VerifyCompressedProof(compressedProof [4]*big.Int, input [128]*big.Int) error {
	return _Verifier.Contract.VerifyCompressedProof(&_Verifier.CallOpts, compressedProof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x0698b542.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[128] input) view returns()
func (_Verifier *VerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [128]*big.Int) error {
	var out []interface{}
	err := _Verifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0x0698b542.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[128] input) view returns()
func (_Verifier *VerifierSession) VerifyProof(proof [8]*big.Int, input [128]*big.Int) error {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0x0698b542.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[128] input) view returns()
func (_Verifier *VerifierCallerSession) VerifyProof(proof [8]*big.Int, input [128]*big.Int) error {
	return _Verifier.Contract.VerifyProof(&_Verifier.CallOpts, proof, input)
}
