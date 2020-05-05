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
const ContractABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_configurationAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_weightsAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_scriptsAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_updatesTillAggregation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_target_epoch\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LocalUpdatesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"trainer\",\"type\":\"address\"}],\"name\":\"addTrainer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configurationAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current_epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"localUpdates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"trainer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"storageAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scriptsAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumModel.states\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"updateAddress\",\"type\":\"string\"}],\"name\":\"submitLocalAggregation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"updateAddress\",\"type\":\"string\"}],\"name\":\"submitLocalUpdate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"target_epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weightsAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x60806040523480156200001157600080fd5b50604051620016d7380380620016d7833981810160405260a08110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660018202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000c3578082015181840152602081019050620000a6565b50505050905090810190601f168015620000f15780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011557600080fd5b838201915060208201858111156200012c57600080fd5b82518660018202830111640100000000821117156200014a57600080fd5b8083526020830192505050908051906020019080838360005b838110156200018057808201518184015260208101905062000163565b50505050905090810190601f168015620001ae5780820380516001836020036101000a031916815260200191505b5060405260200180516040519392919084640100000000821115620001d257600080fd5b83820191506020820185811115620001e957600080fd5b82518660018202830111640100000000821117156200020757600080fd5b8083526020830192505050908051906020019080838360005b838110156200023d57808201518184015260208101905062000220565b50505050905090810190601f1680156200026b5780820380516001836020036101000a031916815260200191505b5060405260200180519060200190929190805190602001909291905050508460019080519060200190620002a192919062000371565b508360029080519060200190620002ba92919062000371565b508260039080519060200190620002d392919062000371565b508160068190555060006004819055508060058190555060008060006101000a81548160ff021916908360028111156200030957fe5b02179055506001600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550505050505062000420565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003b457805160ff1916838001178555620003e5565b82800160010185558215620003e5579182015b82811115620003e4578251825591602001919060010190620003c7565b5b509050620003f49190620003f8565b5090565b6200041d91905b8082111562000419576000816000905550600101620003ff565b5090565b90565b6112a780620004306000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c806367efec9d1161007157806367efec9d14610309578063730a004f1461038c5780639372b4e414610466578063c19d93fb14610484578063cf560221146104b0578063f67eb810146104ce576100a9565b806348222210146100ae5780635117a8401461013157806357c5986d146101755780635ba71fc1146102305780635e824481146102eb575b600080fd5b6100b6610551565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100f65780820151818401526020810190506100db565b50505050905090810190601f1680156101235780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101736004803603602081101561014757600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105ef565b005b61022e6004803603602081101561018b57600080fd5b81019080803590602001906401000000008111156101a857600080fd5b8201836020820111156101ba57600080fd5b803590602001918460018302840111640100000000831117156101dc57600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610709565b005b6102e96004803603602081101561024657600080fd5b810190808035906020019064010000000081111561026357600080fd5b82018360208201111561027557600080fd5b8035906020019184600183028401116401000000008311171561029757600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610955565b005b6102f3610ead565b6040518082815260200191505060405180910390f35b610311610eba565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610351578082015181840152602081019050610336565b50505050905090810190601f16801561037e5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103b8600480360360208110156103a257600080fd5b8101908080359060200190929190505050610f58565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561042a57808201518184015260208101905061040f565b50505050905090810190601f1680156104575780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b61046e611041565b6040518082815260200191505060405180910390f35b61048c611047565b6040518082600281111561049c57fe5b60ff16815260200191505060405180910390f35b6104b8611059565b6040518082815260200191505060405180910390f35b6104d661105f565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105165780820151818401526020810190506104fb565b50505050905090810190601f1680156105435780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105e75780601f106105bc576101008083540402835291602001916105e7565b820191906000526020600020905b8154815290600101906020018083116105ca57829003601f168201915b505050505081565b600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166106ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4e6f7420616e20617574686f72697a656420747261696e65720000000000000081525060200191505060405180910390fd5b6001600c60008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166107c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4e6f7420616e20617574686f72697a656420747261696e65720000000000000081525060200191505060405180910390fd5b60008060009054906101000a900460ff1660028111156107e457fe5b8160028111156107f057fe5b14610863576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f4e6f742076616c6964206174207468697320737461746500000000000000000081525060200191505060405180910390fd5b600b60405180604001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200184815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550602082015181600101908051906020019061091c9291906110fd565b505050600654600b80549050106109515760016000806101000a81548160ff0219169083600281111561094b57fe5b02179055505b5050565b600c60003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610a14576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260198152602001807f4e6f7420616e20617574686f72697a656420747261696e65720000000000000081525060200191505060405180910390fd5b60016000809054906101000a900460ff166002811115610a3057fe5b816002811115610a3c57fe5b14610aaf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260178152602001807f4e6f742076616c6964206174207468697320737461746500000000000000000081525060200191505060405180910390fd5b60016007836040518082805190602001908083835b60208310610ae75780518252602082019150602081019050602083039250610ac4565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054016007836040518082805190602001908083835b60208310610b525780518252602082019150602081019050602083039250610b2f565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902081905550600160095401600981905550600a82908060018154018082558091505060019003906000526020600020016000909190919091509080519060200190610bd09291906110fd565b50600760086040518082805460018160011615610100020316600290048015610c305780601f10610c0e576101008083540402835291820191610c30565b820191906000526020600020905b815481529060010190602001808311610c1c575b50509150509081526020016040518091039020546007836040518082805190602001908083835b60208310610c7a5780518252602082019150602081019050602083039250610c57565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020541115610ccc578160089080519060200190610cca9291906110fd565b505b60065460095410610ea957600860029080546001816001161561010002031660029004610cfa92919061117d565b5060016004540160048190555060006009819055505b6000600a805490501115610de15760006007600a6001600a805490500381548110610d3757fe5b906000526020600020016040518082805460018160011615610100020316600290048015610d9c5780601f10610d7a576101008083540402835291820191610d9c565b820191906000526020600020905b815481529060010190602001808311610d88575b5050915050908152602001604051809103902081905550600a805480610dbe57fe5b600190038181906000526020600020016000610dda9190611204565b9055610d10565b5b6000600b805490501115610e5057600b805480610dfb57fe5b6001900381819060005260206000209060020201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182016000610e479190611204565b50509055610de2565b6005546004541015610e845760008060006101000a81548160ff02191690836002811115610e7a57fe5b0217905550610ea8565b60026000806101000a81548160ff02191690836002811115610ea257fe5b02179055505b5b5050565b6000600b80549050905090565b60038054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f505780601f10610f2557610100808354040283529160200191610f50565b820191906000526020600020905b815481529060010190602001808311610f3357829003601f168201915b505050505081565b600b8181548110610f6557fe5b90600052602060002090600202016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110375780601f1061100c57610100808354040283529160200191611037565b820191906000526020600020905b81548152906001019060200180831161101a57829003601f168201915b5050505050905082565b60045481565b6000809054906101000a900460ff1681565b60055481565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110f55780601f106110ca576101008083540402835291602001916110f5565b820191906000526020600020905b8154815290600101906020018083116110d857829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061113e57805160ff191683800117855561116c565b8280016001018555821561116c579182015b8281111561116b578251825591602001919060010190611150565b5b509050611179919061124c565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106111b657805485556111f3565b828001600101855582156111f357600052602060002091601f016020900482015b828111156111f25782548255916001019190600101906111d7565b5b509050611200919061124c565b5090565b50805460018160011615610100020316600290046000825580601f1061122a5750611249565b601f016020900490600052602060002090810190611248919061124c565b5b50565b61126e91905b8082111561126a576000816000905550600101611252565b5090565b9056fea2646970667358221220b27e505a9906fb35032de5caa01337c3ec4b25be979f20283f96d68b2cd85d5c64736f6c63430006060033"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _configurationAddress string, _weightsAddress string, _scriptsAddress string, _updatesTillAggregation *big.Int, _target_epoch *big.Int) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend, _configurationAddress, _weightsAddress, _scriptsAddress, _updatesTillAggregation, _target_epoch)
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
// Solidity: function LocalUpdatesCount() view returns(uint256)
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
// Solidity: function LocalUpdatesCount() view returns(uint256)
func (_Contract *ContractSession) LocalUpdatesCount() (*big.Int, error) {
	return _Contract.Contract.LocalUpdatesCount(&_Contract.CallOpts)
}

// LocalUpdatesCount is a free data retrieval call binding the contract method 0x5e824481.
//
// Solidity: function LocalUpdatesCount() view returns(uint256)
func (_Contract *ContractCallerSession) LocalUpdatesCount() (*big.Int, error) {
	return _Contract.Contract.LocalUpdatesCount(&_Contract.CallOpts)
}

// ConfigurationAddress is a free data retrieval call binding the contract method 0x48222210.
//
// Solidity: function configurationAddress() view returns(string)
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
// Solidity: function configurationAddress() view returns(string)
func (_Contract *ContractSession) ConfigurationAddress() (string, error) {
	return _Contract.Contract.ConfigurationAddress(&_Contract.CallOpts)
}

// ConfigurationAddress is a free data retrieval call binding the contract method 0x48222210.
//
// Solidity: function configurationAddress() view returns(string)
func (_Contract *ContractCallerSession) ConfigurationAddress() (string, error) {
	return _Contract.Contract.ConfigurationAddress(&_Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x9372b4e4.
//
// Solidity: function current_epoch() view returns(uint256)
func (_Contract *ContractCaller) CurrentEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "current_epoch")
	return *ret0, err
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x9372b4e4.
//
// Solidity: function current_epoch() view returns(uint256)
func (_Contract *ContractSession) CurrentEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// CurrentEpoch is a free data retrieval call binding the contract method 0x9372b4e4.
//
// Solidity: function current_epoch() view returns(uint256)
func (_Contract *ContractCallerSession) CurrentEpoch() (*big.Int, error) {
	return _Contract.Contract.CurrentEpoch(&_Contract.CallOpts)
}

// LocalUpdates is a free data retrieval call binding the contract method 0x730a004f.
//
// Solidity: function localUpdates(uint256 ) view returns(address trainer, string storageAddress)
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
// Solidity: function localUpdates(uint256 ) view returns(address trainer, string storageAddress)
func (_Contract *ContractSession) LocalUpdates(arg0 *big.Int) (struct {
	Trainer        common.Address
	StorageAddress string
}, error) {
	return _Contract.Contract.LocalUpdates(&_Contract.CallOpts, arg0)
}

// LocalUpdates is a free data retrieval call binding the contract method 0x730a004f.
//
// Solidity: function localUpdates(uint256 ) view returns(address trainer, string storageAddress)
func (_Contract *ContractCallerSession) LocalUpdates(arg0 *big.Int) (struct {
	Trainer        common.Address
	StorageAddress string
}, error) {
	return _Contract.Contract.LocalUpdates(&_Contract.CallOpts, arg0)
}

// ScriptsAddress is a free data retrieval call binding the contract method 0x67efec9d.
//
// Solidity: function scriptsAddress() view returns(string)
func (_Contract *ContractCaller) ScriptsAddress(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "scriptsAddress")
	return *ret0, err
}

// ScriptsAddress is a free data retrieval call binding the contract method 0x67efec9d.
//
// Solidity: function scriptsAddress() view returns(string)
func (_Contract *ContractSession) ScriptsAddress() (string, error) {
	return _Contract.Contract.ScriptsAddress(&_Contract.CallOpts)
}

// ScriptsAddress is a free data retrieval call binding the contract method 0x67efec9d.
//
// Solidity: function scriptsAddress() view returns(string)
func (_Contract *ContractCallerSession) ScriptsAddress() (string, error) {
	return _Contract.Contract.ScriptsAddress(&_Contract.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Contract *ContractCaller) State(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "state")
	return *ret0, err
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Contract *ContractSession) State() (uint8, error) {
	return _Contract.Contract.State(&_Contract.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns(uint8)
func (_Contract *ContractCallerSession) State() (uint8, error) {
	return _Contract.Contract.State(&_Contract.CallOpts)
}

// TargetEpoch is a free data retrieval call binding the contract method 0xcf560221.
//
// Solidity: function target_epoch() view returns(uint256)
func (_Contract *ContractCaller) TargetEpoch(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "target_epoch")
	return *ret0, err
}

// TargetEpoch is a free data retrieval call binding the contract method 0xcf560221.
//
// Solidity: function target_epoch() view returns(uint256)
func (_Contract *ContractSession) TargetEpoch() (*big.Int, error) {
	return _Contract.Contract.TargetEpoch(&_Contract.CallOpts)
}

// TargetEpoch is a free data retrieval call binding the contract method 0xcf560221.
//
// Solidity: function target_epoch() view returns(uint256)
func (_Contract *ContractCallerSession) TargetEpoch() (*big.Int, error) {
	return _Contract.Contract.TargetEpoch(&_Contract.CallOpts)
}

// WeightsAddress is a free data retrieval call binding the contract method 0xf67eb810.
//
// Solidity: function weightsAddress() view returns(string)
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
// Solidity: function weightsAddress() view returns(string)
func (_Contract *ContractSession) WeightsAddress() (string, error) {
	return _Contract.Contract.WeightsAddress(&_Contract.CallOpts)
}

// WeightsAddress is a free data retrieval call binding the contract method 0xf67eb810.
//
// Solidity: function weightsAddress() view returns(string)
func (_Contract *ContractCallerSession) WeightsAddress() (string, error) {
	return _Contract.Contract.WeightsAddress(&_Contract.CallOpts)
}

// AddTrainer is a paid mutator transaction binding the contract method 0x5117a840.
//
// Solidity: function addTrainer(address trainer) returns()
func (_Contract *ContractTransactor) AddTrainer(opts *bind.TransactOpts, trainer common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addTrainer", trainer)
}

// AddTrainer is a paid mutator transaction binding the contract method 0x5117a840.
//
// Solidity: function addTrainer(address trainer) returns()
func (_Contract *ContractSession) AddTrainer(trainer common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddTrainer(&_Contract.TransactOpts, trainer)
}

// AddTrainer is a paid mutator transaction binding the contract method 0x5117a840.
//
// Solidity: function addTrainer(address trainer) returns()
func (_Contract *ContractTransactorSession) AddTrainer(trainer common.Address) (*types.Transaction, error) {
	return _Contract.Contract.AddTrainer(&_Contract.TransactOpts, trainer)
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
