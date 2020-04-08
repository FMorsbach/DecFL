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
const ContractABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_configuration\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_weightsAddress\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LocalUpdatesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configurationAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"localUpdates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"trainer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"storageAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"updateAddress\",\"type\":\"string\"}],\"name\":\"submitLocalAggregation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"updateAddress\",\"type\":\"string\"}],\"name\":\"submitLocalUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weightsAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x608060405234801561001057600080fd5b5060405162000af738038062000af78339818101604052604081101561003557600080fd5b810190808051604051939291908464010000000082111561005557600080fd5b8382019150602082018581111561006b57600080fd5b825186600182028301116401000000008211171561008857600080fd5b8083526020830192505050908051906020019080838360005b838110156100bc5780820151818401526020810190506100a1565b50505050905090810190601f1680156100e95780820380516001836020036101000a031916815260200191505b506040526020018051604051939291908464010000000082111561010c57600080fd5b8382019150602082018581111561012257600080fd5b825186600182028301116401000000008211171561013f57600080fd5b8083526020830192505050908051906020019080838360005b83811015610173578082015181840152602081019050610158565b50505050905090810190601f1680156101a05780820380516001836020036101000a031916815260200191505b5060405250505081600090805190602001906101bd9291906101dc565b5080600190805190602001906101d49291906101dc565b505050610281565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061021d57805160ff191683800117855561024b565b8280016001018555821561024b579182015b8281111561024a57825182559160200191906001019061022f565b5b509050610258919061025c565b5090565b61027e91905b8082111561027a576000816000905550600101610262565b5090565b90565b61086680620002916000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c8063482222101461006757806357c5986d146100ea5780635ba71fc1146101a55780635e82448114610260578063730a004f1461027e578063f67eb81014610358575b600080fd5b61006f6103db565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100af578082015181840152602081019050610094565b50505050905090810190601f1680156100dc5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101a36004803603602081101561010057600080fd5b810190808035906020019064010000000081111561011d57600080fd5b82018360208201111561012f57600080fd5b8035906020019184600183028401116401000000008311171561015157600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610479565b005b61025e600480360360208110156101bb57600080fd5b81019080803590602001906401000000008111156101d857600080fd5b8201836020820111156101ea57600080fd5b8035906020019184600183028401116401000000008311171561020c57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610538565b005b6102686105f7565b6040518082815260200191505060405180910390f35b6102aa6004803603602081101561029457600080fd5b8101908080359060200190929190505050610604565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561031c578082015181840152602081019050610301565b50505050905090810190601f1680156103495780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6103606106ed565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156103a0578082015181840152602081019050610385565b50505050905090810190601f1680156103cd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104715780601f1061044657610100808354040283529160200191610471565b820191906000526020600020905b81548152906001019060200180831161045457829003601f168201915b505050505081565b600260405180604001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200183815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908051906020019061053292919061078b565b50505050565b600360405180604001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200183815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190805190602001906105f192919061078b565b50505050565b6000600280549050905090565b6002818154811061061157fe5b90600052602060002090600202016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106e35780601f106106b8576101008083540402835291602001916106e3565b820191906000526020600020905b8154815290600101906020018083116106c657829003601f168201915b5050505050905082565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107835780601f1061075857610100808354040283529160200191610783565b820191906000526020600020905b81548152906001019060200180831161076657829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106107cc57805160ff19168380011785556107fa565b828001600101855582156107fa579182015b828111156107f95782518255916020019190600101906107de565b5b509050610807919061080b565b5090565b61082d91905b80821115610829576000816000905550600101610811565b5090565b9056fea26469706673582212206e4f76d1cef01ca51904329d7b962267ea45185f20494b6091bee2f5c03b06e364736f6c63430006040033"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _configuration string, _weightsAddress string) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend, _configuration, _weightsAddress)
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
