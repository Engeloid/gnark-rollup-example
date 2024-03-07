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

// IVerifierMetaData contains all meta data concerning the IVerifier contract.
var IVerifierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"proof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[8]\",\"name\":\"input\",\"type\":\"uint256[8]\"}],\"name\":\"verifyProof\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IVerifierABI is the input ABI used to generate the binding from.
// Deprecated: Use IVerifierMetaData.ABI instead.
var IVerifierABI = IVerifierMetaData.ABI

// IVerifier is an auto generated Go binding around an Ethereum contract.
type IVerifier struct {
	IVerifierCaller     // Read-only binding to the contract
	IVerifierTransactor // Write-only binding to the contract
	IVerifierFilterer   // Log filterer for contract events
}

// IVerifierCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVerifierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVerifierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVerifierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVerifierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVerifierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVerifierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVerifierSession struct {
	Contract     *IVerifier        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IVerifierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVerifierCallerSession struct {
	Contract *IVerifierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IVerifierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVerifierTransactorSession struct {
	Contract     *IVerifierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IVerifierRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVerifierRaw struct {
	Contract *IVerifier // Generic contract binding to access the raw methods on
}

// IVerifierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVerifierCallerRaw struct {
	Contract *IVerifierCaller // Generic read-only contract binding to access the raw methods on
}

// IVerifierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVerifierTransactorRaw struct {
	Contract *IVerifierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVerifier creates a new instance of IVerifier, bound to a specific deployed contract.
func NewIVerifier(address common.Address, backend bind.ContractBackend) (*IVerifier, error) {
	contract, err := bindIVerifier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVerifier{IVerifierCaller: IVerifierCaller{contract: contract}, IVerifierTransactor: IVerifierTransactor{contract: contract}, IVerifierFilterer: IVerifierFilterer{contract: contract}}, nil
}

// NewIVerifierCaller creates a new read-only instance of IVerifier, bound to a specific deployed contract.
func NewIVerifierCaller(address common.Address, caller bind.ContractCaller) (*IVerifierCaller, error) {
	contract, err := bindIVerifier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVerifierCaller{contract: contract}, nil
}

// NewIVerifierTransactor creates a new write-only instance of IVerifier, bound to a specific deployed contract.
func NewIVerifierTransactor(address common.Address, transactor bind.ContractTransactor) (*IVerifierTransactor, error) {
	contract, err := bindIVerifier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVerifierTransactor{contract: contract}, nil
}

// NewIVerifierFilterer creates a new log filterer instance of IVerifier, bound to a specific deployed contract.
func NewIVerifierFilterer(address common.Address, filterer bind.ContractFilterer) (*IVerifierFilterer, error) {
	contract, err := bindIVerifier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVerifierFilterer{contract: contract}, nil
}

// bindIVerifier binds a generic wrapper to an already deployed contract.
func bindIVerifier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IVerifierMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVerifier *IVerifierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVerifier.Contract.IVerifierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVerifier *IVerifierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVerifier.Contract.IVerifierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVerifier *IVerifierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVerifier.Contract.IVerifierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVerifier *IVerifierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVerifier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVerifier *IVerifierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVerifier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVerifier *IVerifierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVerifier.Contract.contract.Transact(opts, method, params...)
}

// VerifyProof is a free data retrieval call binding the contract method 0xa6047e6c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[8] input) view returns()
func (_IVerifier *IVerifierCaller) VerifyProof(opts *bind.CallOpts, proof [8]*big.Int, input [8]*big.Int) error {
	var out []interface{}
	err := _IVerifier.contract.Call(opts, &out, "verifyProof", proof, input)

	if err != nil {
		return err
	}

	return err

}

// VerifyProof is a free data retrieval call binding the contract method 0xa6047e6c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[8] input) view returns()
func (_IVerifier *IVerifierSession) VerifyProof(proof [8]*big.Int, input [8]*big.Int) error {
	return _IVerifier.Contract.VerifyProof(&_IVerifier.CallOpts, proof, input)
}

// VerifyProof is a free data retrieval call binding the contract method 0xa6047e6c.
//
// Solidity: function verifyProof(uint256[8] proof, uint256[8] input) view returns()
func (_IVerifier *IVerifierCallerSession) VerifyProof(proof [8]*big.Int, input [8]*big.Int) error {
	return _IVerifier.Contract.VerifyProof(&_IVerifier.CallOpts, proof, input)
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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_verifierAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_treeLevels\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initRoot\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMerkleRoot\",\"type\":\"uint256\"}],\"name\":\"MerkleRootUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structRollup.PublicKey\",\"name\":\"pubkey\",\"type\":\"tuple\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"batchId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"TransactionData\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"RollupToEth\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ZERO_VALUE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"computeMerkleRootFromPath\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rollupAccountPKX\",\"type\":\"uint256\"}],\"name\":\"getEthAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLevels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMerkleRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextLeafIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structRollup.PublicKey\",\"name\":\"pubKey\",\"type\":\"tuple\"}],\"internalType\":\"structRollup.Account\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"hashAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"left\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"right\",\"type\":\"uint256\"}],\"name\":\"hashLeftRight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"levels\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"merkleProof\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"ganacheAddress\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structRollup.PublicKey\",\"name\":\"pubKey\",\"type\":\"tuple\"}],\"internalType\":\"structRollup.Account\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"processDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"zkproof\",\"type\":\"uint256[8]\"},{\"internalType\":\"uint256[8]\",\"name\":\"input\",\"type\":\"uint256[8]\"}],\"name\":\"submitVerifyBatch\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"path\",\"type\":\"uint256[]\"},{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"}],\"name\":\"verifyMerkleProof\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"merkleProof\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint64\",\"name\":\"index\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"y\",\"type\":\"uint256\"}],\"internalType\":\"structRollup.PublicKey\",\"name\":\"pubKey\",\"type\":\"tuple\"}],\"internalType\":\"structRollup.Account\",\"name\":\"account\",\"type\":\"tuple\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"i\",\"type\":\"uint256\"}],\"name\":\"zeros\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x608060405234801562000010575f80fd5b50604051620027fa380380620027fa8339818101604052810190620000369190620002a6565b815f81116200007c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000739062000383565b60405180910390fd5b60098110620000c2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000b990620003f1565b60405180910390fd5b6004811462000108576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000ff9062000485565b60405180910390fd5b805f819055507f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd3226760018190555050815f8190555082600760086101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055503360045f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550620001ce81620001ff60201b60201c565b5f60075f6101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905550505050620004a5565b8060018190555050565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f62000238826200020d565b9050919050565b6200024a816200022c565b811462000255575f80fd5b50565b5f8151905062000268816200023f565b92915050565b5f819050919050565b62000282816200026e565b81146200028d575f80fd5b50565b5f81519050620002a08162000277565b92915050565b5f805f60608486031215620002c057620002bf62000209565b5b5f620002cf8682870162000258565b9350506020620002e28682870162000290565b9250506040620002f58682870162000290565b9150509250925092565b5f82825260208201905092915050565b7f5f6c6576656c732073686f756c642062652067726561746572207468616e207a5f8201527f65726f0000000000000000000000000000000000000000000000000000000000602082015250565b5f6200036b602383620002ff565b915062000378826200030f565b604082019050919050565b5f6020820190508181035f8301526200039c816200035d565b9050919050565b7f5f6c6576656c732073686f756c64206265206c657373207468616e20380000005f82015250565b5f620003d9601d83620002ff565b9150620003e682620003a3565b602082019050919050565b5f6020820190508181035f8301526200040a81620003cb565b9050919050565b7f5f6c6576656c732073686f756c6420626520342028325e34203d2031362061635f8201527f636f756e74732920666f722074657374696e6700000000000000000000000000602082015250565b5f6200046d603383620002ff565b91506200047a8262000411565b604082019050919050565b5f6020820190508181035f8301526200049e816200045f565b9050919050565b61234780620004b35f395ff3fe608060405260043610610113575f3560e01c80635ca1e1651161009f578063bd4986a011610063578063bd4986a0146103cb578063d0e30db014610407578063e829558814610411578063ec7329591461044d578063f139d33e1461047757610113565b80635ca1e165146102c357806369ed88ec146102ed57806383290a3c146103295780638da5cb5b14610365578063909ad5ba1461038f57610113565b80633fd141b6116100e65780633fd141b6146101cd57806349590657146102095780634ecf518b1461023357806350e9b9251461025d5780635bb939951461028757610113565b806302722d20146101175780630c394a601461013f5780631174dc99146101695780631ea5f6eb14610191575b5f80fd5b348015610122575f80fd5b5061013d60048036038101906101389190611918565b61049f565b005b34801561014a575f80fd5b506101536107c9565b6040516101609190611993565b60405180910390f35b348015610174575f80fd5b5061018f600480360381019061018a9190611a06565b6107d1565b005b34801561019c575f80fd5b506101b760048036038101906101b29190611a87565b610b38565b6040516101c49190611ac1565b60405180910390f35b3480156101d8575f80fd5b506101f360048036038101906101ee9190611ada565b610b68565b6040516102009190611b4e565b60405180910390f35b348015610214575f80fd5b5061021d610d13565b60405161022a9190611993565b60405180910390f35b34801561023e575f80fd5b50610247610d21565b6040516102549190611993565b60405180910390f35b348015610268575f80fd5b50610271610d26565b60405161027e9190611993565b60405180910390f35b348015610292575f80fd5b506102ad60048036038101906102a89190611b67565b610d2f565b6040516102ba9190611993565b60405180910390f35b3480156102ce575f80fd5b506102d7610e3d565b6040516102e49190611993565b60405180910390f35b3480156102f8575f80fd5b50610313600480360381019061030e9190611ada565b610e46565b6040516103209190611993565b60405180910390f35b348015610334575f80fd5b5061034f600480360381019061034a9190611a87565b610fed565b60405161035c9190611993565b60405180910390f35b348015610370575f80fd5b5061037961100d565b6040516103869190611ac1565b60405180910390f35b34801561039a575f80fd5b506103b560048036038101906103b09190611ba5565b611032565b6040516103c29190611993565b60405180910390f35b3480156103d6575f80fd5b506103f160048036038101906103ec9190611a87565b6111c6565b6040516103fe9190611ac1565b60405180910390f35b61040f6111ff565b005b34801561041c575f80fd5b5061043760048036038101906104329190611a87565b6112db565b6040516104449190611993565b60405180910390f35b348015610458575f80fd5b5061046161149b565b60405161046e9190611993565b60405180910390f35b348015610482575f80fd5b5061049d60048036038101906104989190611bf1565b6114bf565b005b5f60055f83606001515f015181526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610548576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053f90611c8b565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146105b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105ad90611cf3565b60405180910390fd5b83826040015110156105fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105f490611d5b565b60405180910390fd5b5f61060783611032565b905080845f8151811061061d5761061c611d79565b5b60200260200101818152505061063684845f0151610b68565b610675576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161066c90611df0565b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff166108fc8690811502906040515f60405180830381858888f193505050501580156106b8573d5f803e3d5ffd5b508173ffffffffffffffffffffffffffffffffffffffff167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364866040516106ff9190611993565b60405180910390a284836040018181516107199190611e3b565b915081815250508483604001516107309190611e6e565b8360400151111561074457610743611ea1565b5b5f61074e84611032565b905080855f8151811061076457610763611d79565b5b6020026020010181815250505f61077e86865f0151610e46565b905061078981611677565b7f3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a816040516107b89190611993565b60405180910390a150505050505050565b5f8054905090565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610860576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161085790611f18565b60405180910390fd5b61086d84825f0151610b68565b6108ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016108a390611f80565b60405180910390fd5b5f73ffffffffffffffffffffffffffffffffffffffff1660055f83606001515f015181526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16036109b7577f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd32267845f8151811061094a57610949611d79565b5b6020026020010151146109605761095f611ea1565b5b8160055f83606001515f015181526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b600660075f9054906101000a900467ffffffffffffffff1667ffffffffffffffff16815481106109ea576109e9611d79565b5b905f5260205f200154610a13848473ffffffffffffffffffffffffffffffffffffffff16610d2f565b14610a53576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a4a90611fe8565b60405180910390fd5b8281604001511015610a6857610a67611ea1565b5b5f610a7282611032565b905080855f81518110610a8857610a87611d79565b5b6020026020010181815250505f610aa286845f0151610e46565b9050610aad81611677565b7f3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a81604051610adc9190611993565b60405180910390a160075f81819054906101000a900467ffffffffffffffff1680929190610b0990612006565b91906101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555050505050505050565b6005602052805f5260405f205f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f80600167ffffffffffffffff811115610b8557610b846116d9565b5b604051908082528060200260200182016040528015610bb35781602001602082028036833780820191505090505b509050835f81518110610bc957610bc8611d79565b5b6020026020010151815f81518110610be457610be3611d79565b5b6020026020010181815250505f73__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49836040518263ffffffff1660e01b8152600401610c2891906120ec565b602060405180830381865af4158015610c43573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c679190612120565b90505f600190505b8551811015610d03575f600286610c869190612178565b67ffffffffffffffff1603610cc057610cb982878381518110610cac57610cab611d79565b5b6020026020010151610d2f565b9150610ce7565b610ce4868281518110610cd657610cd5611d79565b5b602002602001015183610d2f565b91505b600285610cf491906121a8565b94508080600101915050610c6f565b5060015481149250505092915050565b5f610d1c610e3d565b905090565b5f5481565b5f600254905090565b5f80600267ffffffffffffffff811115610d4c57610d4b6116d9565b5b604051908082528060200260200182016040528015610d7a5781602001602082028036833780820191505090505b50905083815f81518110610d9157610d90611d79565b5b6020026020010181815250508281600181518110610db257610db1611d79565b5b60200260200101818152505073__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49826040518263ffffffff1660e01b8152600401610df591906120ec565b602060405180830381865af4158015610e10573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e349190612120565b91505092915050565b5f600154905090565b5f80600167ffffffffffffffff811115610e6357610e626116d9565b5b604051908082528060200260200182016040528015610e915781602001602082028036833780820191505090505b509050835f81518110610ea757610ea6611d79565b5b6020026020010151815f81518110610ec257610ec1611d79565b5b6020026020010181815250505f73__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49836040518263ffffffff1660e01b8152600401610f0691906120ec565b602060405180830381865af4158015610f21573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f459190612120565b90505f600190505b8551811015610fe1575f600286610f649190612178565b67ffffffffffffffff1603610f9e57610f9782878381518110610f8a57610f89611d79565b5b6020026020010151610d2f565b9150610fc5565b610fc2868281518110610fb457610fb3611d79565b5b602002602001015183610d2f565b91505b600285610fd291906121a8565b94508080600101915050610f4d565b50809250505092915050565b60068181548110610ffc575f80fd5b905f5260205f20015f915090505481565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f80600567ffffffffffffffff81111561104f5761104e6116d9565b5b60405190808252806020026020018201604052801561107d5781602001602082028036833780820191505090505b509050825f015167ffffffffffffffff16815f815181106110a1576110a0611d79565b5b6020026020010181815250508260200151816001815181106110c6576110c5611d79565b5b6020026020010181815250508260400151816002815181106110eb576110ea611d79565b5b60200260200101818152505082606001515f01518160038151811061111357611112611d79565b5b6020026020010181815250508260600151602001518160048151811061113c5761113b611d79565b5b60200260200101818152505073__$7fbbcdf035b79e739e205dd289b9b9f179$__6340ec6e49826040518263ffffffff1660e01b815260040161117f91906120ec565b602060405180830381865af415801561119a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111be9190612120565b915050919050565b5f60055f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f3411611241576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161123890612248565b60405180910390fd5b5f611262343373ffffffffffffffffffffffffffffffffffffffff16610d2f565b9050600681908060018154018082558091505060019003905f5260205f20015f90919091909150553373ffffffffffffffffffffffffffffffffffffffff167fe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c346040516112d09190611993565b60405180910390a250565b5f80820361130b577f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd322679050611496565b6001820361133b577f1541c2e47d6e125aea7a107e749299d88fad35bd227179d4fb2bb0f63296d7c79050611496565b6002820361136b577f22d9432cd2f62dedd322bef0f3be1c172e71b6a1afe7866aee3b031ffd89aad79050611496565b6003820361139b577f049889855d82d41769e27ca791a905e528d867095ce49a8ac961ef4722ba6a169050611496565b600482036113cb577f2d1764e964f231d9129b555e1a653a65fca6e20051a0f7a388b5fd15fbb594559050611496565b600582036113fb577f26d2dd833a220c13591b1a528c91e635044b3fac5b46f4133435ff4c87f046949050611496565b6006820361142b577f2c42c8bc5bb243c5fbf697f47b4f41ba191664a8724e5e370090d6d8817930a89050611496565b6007820361145b577f1aaca32cb69540012989bd902c04015555aa164980a490df8c1e6be64fa57fca9050611496565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161148d906122b0565b60405180910390fd5b919050565b7f04eeee01e61bc107da94dbecba316a16b023ccf7f001d076848442e4dfd3226781565b60045f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461154e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161154590611f18565b60405180910390fd5b5f611557610e3d565b9050815f6008811061156c5761156b611d79565b5b602002013581146115805761157f611ea1565b5b600760089054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a6047e6c84846040518363ffffffff1660e01b81526004016115dd9291906122e8565b5f6040518083038186803b1580156115f3575f80fd5b505afa158015611605573d5f803e3d5ffd5b505050505f82600160086116199190611e3b565b6008811061162a57611629611d79565b5b6020020135905061163a81611677565b7f3ee94a330b3a2d0451c3863014f4175fee871947fc526006b5c7fc72973f674a816040516116699190611993565b60405180910390a150505050565b8060018190555050565b5f604051905090565b5f80fd5b5f80fd5b5f819050919050565b6116a481611692565b81146116ae575f80fd5b50565b5f813590506116bf8161169b565b92915050565b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61170f826116c9565b810181811067ffffffffffffffff8211171561172e5761172d6116d9565b5b80604052505050565b5f611740611681565b905061174c8282611706565b919050565b5f67ffffffffffffffff82111561176b5761176a6116d9565b5b602082029050602081019050919050565b5f80fd5b5f61179261178d84611751565b611737565b905080838252602082019050602084028301858111156117b5576117b461177c565b5b835b818110156117de57806117ca88826116b1565b8452602084019350506020810190506117b7565b5050509392505050565b5f82601f8301126117fc576117fb6116c5565b5b813561180c848260208601611780565b91505092915050565b5f80fd5b5f67ffffffffffffffff82169050919050565b61183581611819565b811461183f575f80fd5b50565b5f813590506118508161182c565b92915050565b5f6040828403121561186b5761186a611815565b5b6118756040611737565b90505f611884848285016116b1565b5f830152506020611897848285016116b1565b60208301525092915050565b5f60a082840312156118b8576118b7611815565b5b6118c26080611737565b90505f6118d184828501611842565b5f8301525060206118e4848285016116b1565b60208301525060406118f8848285016116b1565b604083015250606061190c84828501611856565b60608301525092915050565b5f805f60e0848603121561192f5761192e61168a565b5b5f61193c868287016116b1565b935050602084013567ffffffffffffffff81111561195d5761195c61168e565b5b611969868287016117e8565b925050604061197a868287016118a3565b9150509250925092565b61198d81611692565b82525050565b5f6020820190506119a65f830184611984565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6119d5826119ac565b9050919050565b6119e5816119cb565b81146119ef575f80fd5b50565b5f81359050611a00816119dc565b92915050565b5f805f806101008587031215611a1f57611a1e61168a565b5b5f85013567ffffffffffffffff811115611a3c57611a3b61168e565b5b611a48878288016117e8565b9450506020611a59878288016116b1565b9350506040611a6a878288016119f2565b9250506060611a7b878288016118a3565b91505092959194509250565b5f60208284031215611a9c57611a9b61168a565b5b5f611aa9848285016116b1565b91505092915050565b611abb816119cb565b82525050565b5f602082019050611ad45f830184611ab2565b92915050565b5f8060408385031215611af057611aef61168a565b5b5f83013567ffffffffffffffff811115611b0d57611b0c61168e565b5b611b19858286016117e8565b9250506020611b2a85828601611842565b9150509250929050565b5f8115159050919050565b611b4881611b34565b82525050565b5f602082019050611b615f830184611b3f565b92915050565b5f8060408385031215611b7d57611b7c61168a565b5b5f611b8a858286016116b1565b9250506020611b9b858286016116b1565b9150509250929050565b5f60a08284031215611bba57611bb961168a565b5b5f611bc7848285016118a3565b91505092915050565b5f81905082602060080282011115611beb57611bea61177c565b5b92915050565b5f806102008385031215611c0857611c0761168a565b5b5f611c1585828601611bd0565b925050610100611c2785828601611bd0565b9150509250929050565b5f82825260208201905092915050565b7f526f6c6c7570206164647265737320646f6573206e6f742065786973740000005f82015250565b5f611c75601d83611c31565b9150611c8082611c41565b602082019050919050565b5f6020820190508181035f830152611ca281611c69565b9050919050565b7f4e6f7420726967687466756c206163636f756e742061646472657373000000005f82015250565b5f611cdd601c83611c31565b9150611ce882611ca9565b602082019050919050565b5f6020820190508181035f830152611d0a81611cd1565b9050919050565b7f42616c616e6365206e6f742073756666696369656e74000000000000000000005f82015250565b5f611d45601683611c31565b9150611d5082611d11565b602082019050919050565b5f6020820190508181035f830152611d7281611d39565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52603260045260245ffd5b7f496e76616c6964204d65726b6c652070726f6f660000000000000000000000005f82015250565b5f611dda601483611c31565b9150611de582611da6565b602082019050919050565b5f6020820190508181035f830152611e0781611dce565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f611e4582611692565b9150611e5083611692565b9250828203905081811115611e6857611e67611e0e565b5b92915050565b5f611e7882611692565b9150611e8383611692565b9250828201905080821115611e9b57611e9a611e0e565b5b92915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52600160045260245ffd5b7f4f6e6c79206f776e65722063616e207375626d697420626174636865730000005f82015250565b5f611f02601d83611c31565b9150611f0d82611ece565b602082019050919050565b5f6020820190508181035f830152611f2f81611ef6565b9050919050565b7f4d65726b6c652070726f6f6620696e76616c69640000000000000000000000005f82015250565b5f611f6a601483611c31565b9150611f7582611f36565b602082019050919050565b5f6020820190508181035f830152611f9781611f5e565b9050919050565b7f4465706f736974206e6f7420636f7272656374206f6e6520696e2071756575655f82015250565b5f611fd2602083611c31565b9150611fdd82611f9e565b602082019050919050565b5f6020820190508181035f830152611fff81611fc6565b9050919050565b5f61201082611819565b915067ffffffffffffffff820361202a57612029611e0e565b5b600182019050919050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b61206781611692565b82525050565b5f612078838361205e565b60208301905092915050565b5f602082019050919050565b5f61209a82612035565b6120a4818561203f565b93506120af8361204f565b805f5b838110156120df5781516120c6888261206d565b97506120d183612084565b9250506001810190506120b2565b5085935050505092915050565b5f6020820190508181035f8301526121048184612090565b905092915050565b5f8151905061211a8161169b565b92915050565b5f602082840312156121355761213461168a565b5b5f6121428482850161210c565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffd5b5f61218282611819565b915061218d83611819565b92508261219d5761219c61214b565b5b828206905092915050565b5f6121b282611819565b91506121bd83611819565b9250826121cd576121cc61214b565b5b828204905092915050565b7f4465706f7369742076616c7565206d75737420626520677265617465722074685f8201527f616e203000000000000000000000000000000000000000000000000000000000602082015250565b5f612232602483611c31565b915061223d826121d8565b604082019050919050565b5f6020820190508181035f83015261225f81612226565b9050919050565b7f696e646578206f7574206f6620626f756e6473000000000000000000000000005f82015250565b5f61229a601383611c31565b91506122a582612266565b602082019050919050565b5f6020820190508181035f8301526122c78161228e565b9050919050565b82818337505050565b6122e461010083836122ce565b5050565b5f610200820190506122fc5f8301856122d7565b61230a6101008301846122d7565b939250505056fea2646970667358221220f88185c735c5eb5fc768badb062bb4f0bf772344caddf402df2da3dea4475d6764736f6c63430008180033",
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

// SubmitVerifyBatch is a paid mutator transaction binding the contract method 0xf139d33e.
//
// Solidity: function submitVerifyBatch(uint256[8] zkproof, uint256[8] input) returns()
func (_Rollup *RollupTransactor) SubmitVerifyBatch(opts *bind.TransactOpts, zkproof [8]*big.Int, input [8]*big.Int) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "submitVerifyBatch", zkproof, input)
}

// SubmitVerifyBatch is a paid mutator transaction binding the contract method 0xf139d33e.
//
// Solidity: function submitVerifyBatch(uint256[8] zkproof, uint256[8] input) returns()
func (_Rollup *RollupSession) SubmitVerifyBatch(zkproof [8]*big.Int, input [8]*big.Int) (*types.Transaction, error) {
	return _Rollup.Contract.SubmitVerifyBatch(&_Rollup.TransactOpts, zkproof, input)
}

// SubmitVerifyBatch is a paid mutator transaction binding the contract method 0xf139d33e.
//
// Solidity: function submitVerifyBatch(uint256[8] zkproof, uint256[8] input) returns()
func (_Rollup *RollupTransactorSession) SubmitVerifyBatch(zkproof [8]*big.Int, input [8]*big.Int) (*types.Transaction, error) {
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
