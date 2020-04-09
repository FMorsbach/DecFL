// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_configuration\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_weightsAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_updatesTillAggregation\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LocalUpdatesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configurationAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"localUpdates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"trainer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"storageAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"updateAddress\",\"type\":\"string\"}],\"name\":\"submitLocalAggregation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"updateAddress\",\"type\":\"string\"}],\"name\":\"submitLocalUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weightsAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x60806040523480156200001157600080fd5b5060405162000f3638038062000f36833981810160405260608110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660018202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000c3578082015181840152602081019050620000a6565b50505050905090810190601f168015620000f15780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011557600080fd5b838201915060208201858111156200012c57600080fd5b82518660018202830111640100000000821117156200014a57600080fd5b8083526020830192505050908051906020019080838360005b838110156200018057808201518184015260208101905062000163565b50505050905090810190601f168015620001ae5780820380516001836020036101000a031916815260200191505b50604052602001805190602001909291905050508260009080519060200190620001da9291906200020c565b508160019080519060200190620001f39291906200020c565b50806003819055506000600281905550505050620002bb565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200024f57805160ff191683800117855562000280565b8280016001018555821562000280579182015b828111156200027f57825182559160200191906001019062000262565b5b5090506200028f919062000293565b5090565b620002b891905b80821115620002b45760008160009055506001016200029a565b5090565b90565b610c6b80620002cb6000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80635e8244811161005b5780635e8244811461027b578063730a004f14610299578063900cf0cf14610373578063f67eb810146103915761007d565b8063482222101461008257806357c5986d146101055780635ba71fc1146101c0575b600080fd5b61008a610414565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100ca5780820151818401526020810190506100af565b50505050905090810190601f1680156100f75780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101be6004803603602081101561011b57600080fd5b810190808035906020019064010000000081111561013857600080fd5b82018360208201111561014a57600080fd5b8035906020019184600183028401116401000000008311171561016c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505091929192905050506104b2565b005b610279600480360360208110156101d657600080fd5b81019080803590602001906401000000008111156101f357600080fd5b82018360208201111561020557600080fd5b8035906020019184600183028401116401000000008311171561022757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610571565b005b6102836108a7565b6040518082815260200191505060405180910390f35b6102c5600480360360208110156102af57600080fd5b81019080803590602001909291905050506108b4565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561033757808201518184015260208101905061031c565b50505050905090810190601f1680156103645780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b61037b61099d565b6040518082815260200191505060405180910390f35b6103996109a3565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103d95780820151818401526020810190506103be565b50505050905090810190601f1680156104065780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104aa5780601f1061047f576101008083540402835291602001916104aa565b820191906000526020600020905b81548152906001019060200180831161048d57829003601f168201915b505050505081565b600860405180604001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200183815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908051906020019061056b929190610a41565b50505050565b60016004826040518082805190602001908083835b602083106105a95780518252602082019150602081019050602083039250610586565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054016004826040518082805190602001908083835b6020831061061457805182526020820191506020810190506020830392506105f1565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902081905550600160065401600681905550600781908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610692929190610ac1565b506004600560405180828054600181600116156101000203166002900480156106f25780601f106106d05761010080835404028352918201916106f2565b820191906000526020600020905b8154815290600101906020018083116106de575b50509150509081526020016040518091039020546004826040518082805190602001908083835b6020831061073c5780518252602082019150602081019050602083039250610719565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054111561078e57806005908051906020019061078c929190610ac1565b505b600354600654106108a4576005600190805460018160011615610100020316600290046107bc929190610b41565b5060016002540160028190555060006006819055505b600060078054905011156108a357600060046007600160078054905003815481106107f957fe5b90600052602060002001604051808280546001816001161561010002031660029004801561085e5780601f1061083c57610100808354040283529182019161085e565b820191906000526020600020905b81548152906001019060200180831161084a575b5050915050908152602001604051809103902081905550600780548061088057fe5b60019003818190600052602060002001600061089c9190610bc8565b90556107d2565b5b50565b6000600880549050905090565b600881815481106108c157fe5b90600052602060002090600202016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109935780601f1061096857610100808354040283529160200191610993565b820191906000526020600020905b81548152906001019060200180831161097657829003601f168201915b5050505050905082565b60025481565b60018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a395780601f10610a0e57610100808354040283529160200191610a39565b820191906000526020600020905b815481529060010190602001808311610a1c57829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610a8257805160ff1916838001178555610ab0565b82800160010185558215610ab0579182015b82811115610aaf578251825591602001919060010190610a94565b5b509050610abd9190610c10565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b0257805160ff1916838001178555610b30565b82800160010185558215610b30579182015b82811115610b2f578251825591602001919060010190610b14565b5b509050610b3d9190610c10565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b7a5780548555610bb7565b82800160010185558215610bb757600052602060002091601f016020900482015b82811115610bb6578254825591600101919060010190610b9b565b5b509050610bc49190610c10565b5090565b50805460018160011615610100020316600290046000825580601f10610bee5750610c0d565b601f016020900490600052602060002090810190610c0c9190610c10565b5b50565b610c3291905b80821115610c2e576000816000905550600101610c16565b5090565b9056fea264697066735822122070f3fcadb8e7beef649e02bf1829f714b80a621d84c4e1698f0ba325bb0b907364736f6c63430006040033"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _configuration string, _weightsAddress string, _updatesTillAggregation *big.Int) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend, _configuration, _weightsAddress, _updatesTillAggregation)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// LocalUpdatesCount is a free data retrieval call binding the contract method 0x5e824481.
//
// Solidity: function LocalUpdatesCount() constant returns(uint256)
func (_Contract *ContractCaller) LocalUpdatesCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "LocalUpdatesCount")
	return *ret0, err
}

// LocalUpdatesCount is a free data retrieval call binding the contract method 0x5e824481.
//
// Solidity: function LocalUpdatesCount() constant returns(uint256)
func (_Contract *ContractSession) LocalUpdatesCount() (*big.Int, error) {
	return _Contract.Contract.LocalUpdatesCount(&_Contract.CallOpts)
}

// LocalUpdatesCount is a free data retrieval call binding the contract method 0x5e824481.
//
// Solidity: function LocalUpdatesCount() constant returns(uint256)
func (_Contract *ContractCallerSession) LocalUpdatesCount() (*big.Int, error) {
	return _Contract.Contract.LocalUpdatesCount(&_Contract.CallOpts)
}

// ConfigurationAddress is a free data retrieval call binding the contract method 0x48222210.
//
// Solidity: function configurationAddress() constant returns(string)
func (_Contract *ContractCaller) ConfigurationAddress(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "configurationAddress")
	return *ret0, err
}

// ConfigurationAddress is a free data retrieval call binding the contract method 0x48222210.
//
// Solidity: function configurationAddress() constant returns(string)
func (_Contract *ContractSession) ConfigurationAddress() (string, error) {
	return _Contract.Contract.ConfigurationAddress(&_Contract.CallOpts)
}

// ConfigurationAddress is a free data retrieval call binding the contract method 0x48222210.
//
// Solidity: function configurationAddress() constant returns(string)
func (_Contract *ContractCallerSession) ConfigurationAddress() (string, error) {
	return _Contract.Contract.ConfigurationAddress(&_Contract.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() constant returns(uint256)
func (_Contract *ContractCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "epoch")
	return *ret0, err
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() constant returns(uint256)
func (_Contract *ContractSession) Epoch() (*big.Int, error) {
	return _Contract.Contract.Epoch(&_Contract.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() constant returns(uint256)
func (_Contract *ContractCallerSession) Epoch() (*big.Int, error) {
	return _Contract.Contract.Epoch(&_Contract.CallOpts)
}

// LocalUpdates is a free data retrieval call binding the contract method 0x730a004f.
//
// Solidity: function localUpdates(uint256 ) constant returns(address trainer, string storageAddress)
func (_Contract *ContractCaller) LocalUpdates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Trainer        common.Address
	StorageAddress string
}, error) {
	ret := new(struct {
		Trainer        common.Address
		StorageAddress string
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "localUpdates", arg0)
	return *ret, err
}

// LocalUpdates is a free data retrieval call binding the contract method 0x730a004f.
//
// Solidity: function localUpdates(uint256 ) constant returns(address trainer, string storageAddress)
func (_Contract *ContractSession) LocalUpdates(arg0 *big.Int) (struct {
	Trainer        common.Address
	StorageAddress string
}, error) {
	return _Contract.Contract.LocalUpdates(&_Contract.CallOpts, arg0)
}

// LocalUpdates is a free data retrieval call binding the contract method 0x730a004f.
//
// Solidity: function localUpdates(uint256 ) constant returns(address trainer, string storageAddress)
func (_Contract *ContractCallerSession) LocalUpdates(arg0 *big.Int) (struct {
	Trainer        common.Address
	StorageAddress string
}, error) {
	return _Contract.Contract.LocalUpdates(&_Contract.CallOpts, arg0)
}

// WeightsAddress is a free data retrieval call binding the contract method 0xf67eb810.
//
// Solidity: function weightsAddress() constant returns(string)
func (_Contract *ContractCaller) WeightsAddress(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "weightsAddress")
	return *ret0, err
}

// WeightsAddress is a free data retrieval call binding the contract method 0xf67eb810.
//
// Solidity: function weightsAddress() constant returns(string)
func (_Contract *ContractSession) WeightsAddress() (string, error) {
	return _Contract.Contract.WeightsAddress(&_Contract.CallOpts)
}

// WeightsAddress is a free data retrieval call binding the contract method 0xf67eb810.
//
// Solidity: function weightsAddress() constant returns(string)
func (_Contract *ContractCallerSession) WeightsAddress() (string, error) {
	return _Contract.Contract.WeightsAddress(&_Contract.CallOpts)
}

// SubmitLocalAggregation is a paid mutator transaction binding the contract method 0x5ba71fc1.
//
// Solidity: function submitLocalAggregation(string updateAddress) returns()
func (_Contract *ContractTransactor) SubmitLocalAggregation(opts *bind.TransactOpts, updateAddress string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submitLocalAggregation", updateAddress)
}

// SubmitLocalAggregation is a paid mutator transaction binding the contract method 0x5ba71fc1.
//
// Solidity: function submitLocalAggregation(string updateAddress) returns()
func (_Contract *ContractSession) SubmitLocalAggregation(updateAddress string) (*types.Transaction, error) {
	return _Contract.Contract.SubmitLocalAggregation(&_Contract.TransactOpts, updateAddress)
}

// SubmitLocalAggregation is a paid mutator transaction binding the contract method 0x5ba71fc1.
//
// Solidity: function submitLocalAggregation(string updateAddress) returns()
func (_Contract *ContractTransactorSession) SubmitLocalAggregation(updateAddress string) (*types.Transaction, error) {
	return _Contract.Contract.SubmitLocalAggregation(&_Contract.TransactOpts, updateAddress)
}

// SubmitLocalUpdate is a paid mutator transaction binding the contract method 0x57c5986d.
//
// Solidity: function submitLocalUpdate(string updateAddress) returns()
func (_Contract *ContractTransactor) SubmitLocalUpdate(opts *bind.TransactOpts, updateAddress string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submitLocalUpdate", updateAddress)
}

// SubmitLocalUpdate is a paid mutator transaction binding the contract method 0x57c5986d.
//
// Solidity: function submitLocalUpdate(string updateAddress) returns()
func (_Contract *ContractSession) SubmitLocalUpdate(updateAddress string) (*types.Transaction, error) {
	return _Contract.Contract.SubmitLocalUpdate(&_Contract.TransactOpts, updateAddress)
}

// SubmitLocalUpdate is a paid mutator transaction binding the contract method 0x57c5986d.
//
// Solidity: function submitLocalUpdate(string updateAddress) returns()
func (_Contract *ContractTransactorSession) SubmitLocalUpdate(updateAddress string) (*types.Transaction, error) {
	return _Contract.Contract.SubmitLocalUpdate(&_Contract.TransactOpts, updateAddress)
}
