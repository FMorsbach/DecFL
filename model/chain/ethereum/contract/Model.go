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
const ContractABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_configurationAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_weightsAddress\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_updatesTillAggregation\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_target_epoch\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LocalUpdatesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"configurationAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"current_epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"localUpdates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"trainer\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"storageAddress\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumModel.states\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"updateAddress\",\"type\":\"string\"}],\"name\":\"submitLocalAggregation\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"updateAddress\",\"type\":\"string\"}],\"name\":\"submitLocalUpdate\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"target_epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"weightsAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractBin is the compiled bytecode used for deploying new contracts.
var ContractBin = "0x60806040523480156200001157600080fd5b506040516200110838038062001108833981810160405260808110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b838201915060208201858111156200006f57600080fd5b82518660018202830111640100000000821117156200008d57600080fd5b8083526020830192505050908051906020019080838360005b83811015620000c3578082015181840152602081019050620000a6565b50505050905090810190601f168015620000f15780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200011557600080fd5b838201915060208201858111156200012c57600080fd5b82518660018202830111640100000000821117156200014a57600080fd5b8083526020830192505050908051906020019080838360005b838110156200018057808201518184015260208101905062000163565b50505050905090810190601f168015620001ae5780820380516001836020036101000a031916815260200191505b5060405260200180519060200190929190805190602001909291905050508360019080519060200190620001e492919062000242565b508260029080519060200190620001fd92919062000242565b508160058190555060006003819055508060048190555060008060006101000a81548160ff021916908360028111156200023357fe5b021790555050505050620002f1565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200028557805160ff1916838001178555620002b6565b82800160010185558215620002b6579182015b82811115620002b557825182559160200191906001019062000298565b5b509050620002c59190620002c9565b5090565b620002ee91905b80821115620002ea576000816000905550600101620002d0565b5090565b90565b610e0780620003016000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063730a004f11610066578063730a004f146102df5780639372b4e4146103b9578063c19d93fb146103d7578063cf56022114610403578063f67eb8101461042157610093565b8063482222101461009857806357c5986d1461011b5780635ba71fc1146101ee5780635e824481146102c1575b600080fd5b6100a06104a4565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156100e05780820151818401526020810190506100c5565b50505050905090810190601f16801561010d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101d46004803603602081101561013157600080fd5b810190808035906020019064010000000081111561014e57600080fd5b82018360208201111561016057600080fd5b8035906020019184600183028401116401000000008311171561018257600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610542565b604051808215151515815260200191505060405180910390f35b6102a76004803603602081101561020457600080fd5b810190808035906020019064010000000081111561022157600080fd5b82018360208201111561023357600080fd5b8035906020019184600183028401116401000000008311171561025557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050509192919290505050610670565b604051808215151515815260200191505060405180910390f35b6102c9610aab565b6040518082815260200191505060405180910390f35b61030b600480360360208110156102f557600080fd5b8101908080359060200190929190505050610ab8565b604051808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200180602001828103825283818151815260200191508051906020019080838360005b8381101561037d578082015181840152602081019050610362565b50505050905090810190601f1680156103aa5780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b6103c1610ba1565b6040518082815260200191505060405180910390f35b6103df610ba7565b604051808260028111156103ef57fe5b60ff16815260200191505060405180910390f35b61040b610bb9565b6040518082815260200191505060405180910390f35b610429610bbf565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561046957808201518184015260208101905061044e565b50505050905090810190601f1680156104965780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561053a5780601f1061050f5761010080835404028352916020019161053a565b820191906000526020600020905b81548152906001019060200180831161051d57829003601f168201915b505050505081565b600080600281111561055057fe5b6000809054906101000a900460ff16600281111561056a57fe5b14610578576000905061066b565b600a60405180604001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200184815250908060018154018082558091505060019003906000526020600020906002020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019080519060200190610631929190610c5d565b505050600554600a80549050106106665760016000806101000a81548160ff0219169083600281111561066057fe5b02179055505b600190505b919050565b60006001600281111561067f57fe5b6000809054906101000a900460ff16600281111561069957fe5b146106a75760009050610aa6565b60016006836040518082805190602001908083835b602083106106df57805182526020820191506020810190506020830392506106bc565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902054016006836040518082805190602001908083835b6020831061074a5780518252602082019150602081019050602083039250610727565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020819055506001600854016008819055506009829080600181540180825580915050600190039060005260206000200160009091909190915090805190602001906107c8929190610c5d565b506006600760405180828054600181600116156101000203166002900480156108285780601f10610806576101008083540402835291820191610828565b820191906000526020600020905b815481529060010190602001808311610814575b50509150509081526020016040518091039020546006836040518082805190602001908083835b60208310610872578051825260208201915060208101905060208303925061084f565b6001836020036101000a03801982511681845116808217855250505050505090500191505090815260200160405180910390205411156108c45781600790805190602001906108c2929190610c5d565b505b60055460085410610aa1576007600290805460018160011615610100020316600290046108f2929190610cdd565b5060016003540160038190555060006008819055505b600060098054905011156109d9576000600660096001600980549050038154811061092f57fe5b9060005260206000200160405180828054600181600116156101000203166002900480156109945780601f10610972576101008083540402835291820191610994565b820191906000526020600020905b815481529060010190602001808311610980575b505091505090815260200160405180910390208190555060098054806109b657fe5b6001900381819060005260206000200160006109d29190610d64565b9055610908565b5b6000600a805490501115610a4857600a8054806109f357fe5b6001900381819060005260206000209060020201600080820160006101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055600182016000610a3f9190610d64565b505090556109da565b6004546003541015610a7c5760008060006101000a81548160ff02191690836002811115610a7257fe5b0217905550610aa0565b60026000806101000a81548160ff02191690836002811115610a9a57fe5b02179055505b5b600190505b919050565b6000600a80549050905090565b600a8181548110610ac557fe5b90600052602060002090600202016000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b975780601f10610b6c57610100808354040283529160200191610b97565b820191906000526020600020905b815481529060010190602001808311610b7a57829003601f168201915b5050505050905082565b60035481565b6000809054906101000a900460ff1681565b60045481565b60028054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c555780601f10610c2a57610100808354040283529160200191610c55565b820191906000526020600020905b815481529060010190602001808311610c3857829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c9e57805160ff1916838001178555610ccc565b82800160010185558215610ccc579182015b82811115610ccb578251825591602001919060010190610cb0565b5b509050610cd99190610dac565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610d165780548555610d53565b82800160010185558215610d5357600052602060002091601f016020900482015b82811115610d52578254825591600101919060010190610d37565b5b509050610d609190610dac565b5090565b50805460018160011615610100020316600290046000825580601f10610d8a5750610da9565b601f016020900490600052602060002090810190610da89190610dac565b5b50565b610dce91905b80821115610dca576000816000905550600101610db2565b5090565b9056fea26469706673582212201aa0608ad2a0a78b8a1f2da4ffeeece578ba31fa21f59c12a775ab310ee0f44364736f6c63430006060033"

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend, _configurationAddress string, _weightsAddress string, _updatesTillAggregation *big.Int, _target_epoch *big.Int) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractBin), backend, _configurationAddress, _weightsAddress, _updatesTillAggregation, _target_epoch)
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

// SubmitLocalAggregation is a paid mutator transaction binding the contract method 0x5ba71fc1.
//
// Solidity: function submitLocalAggregation(string updateAddress) returns(bool)
func (_Contract *ContractTransactor) SubmitLocalAggregation(opts *bind.TransactOpts, updateAddress string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submitLocalAggregation", updateAddress)
}

// SubmitLocalAggregation is a paid mutator transaction binding the contract method 0x5ba71fc1.
//
// Solidity: function submitLocalAggregation(string updateAddress) returns(bool)
func (_Contract *ContractSession) SubmitLocalAggregation(updateAddress string) (*types.Transaction, error) {
	return _Contract.Contract.SubmitLocalAggregation(&_Contract.TransactOpts, updateAddress)
}

// SubmitLocalAggregation is a paid mutator transaction binding the contract method 0x5ba71fc1.
//
// Solidity: function submitLocalAggregation(string updateAddress) returns(bool)
func (_Contract *ContractTransactorSession) SubmitLocalAggregation(updateAddress string) (*types.Transaction, error) {
	return _Contract.Contract.SubmitLocalAggregation(&_Contract.TransactOpts, updateAddress)
}

// SubmitLocalUpdate is a paid mutator transaction binding the contract method 0x57c5986d.
//
// Solidity: function submitLocalUpdate(string updateAddress) returns(bool)
func (_Contract *ContractTransactor) SubmitLocalUpdate(opts *bind.TransactOpts, updateAddress string) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "submitLocalUpdate", updateAddress)
}

// SubmitLocalUpdate is a paid mutator transaction binding the contract method 0x57c5986d.
//
// Solidity: function submitLocalUpdate(string updateAddress) returns(bool)
func (_Contract *ContractSession) SubmitLocalUpdate(updateAddress string) (*types.Transaction, error) {
	return _Contract.Contract.SubmitLocalUpdate(&_Contract.TransactOpts, updateAddress)
}

// SubmitLocalUpdate is a paid mutator transaction binding the contract method 0x57c5986d.
//
// Solidity: function submitLocalUpdate(string updateAddress) returns(bool)
func (_Contract *ContractTransactorSession) SubmitLocalUpdate(updateAddress string) (*types.Transaction, error) {
	return _Contract.Contract.SubmitLocalUpdate(&_Contract.TransactOpts, updateAddress)
}
