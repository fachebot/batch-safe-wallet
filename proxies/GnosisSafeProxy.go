// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxies

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
)

// ProxiesMetaData contains all meta data concerning the Proxies contract.
var ProxiesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_singleton\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"}]",
}

// ProxiesABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxiesMetaData.ABI instead.
var ProxiesABI = ProxiesMetaData.ABI

// Proxies is an auto generated Go binding around an Ethereum contract.
type Proxies struct {
	ProxiesCaller     // Read-only binding to the contract
	ProxiesTransactor // Write-only binding to the contract
	ProxiesFilterer   // Log filterer for contract events
}

// ProxiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxiesSession struct {
	Contract     *Proxies          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxiesCallerSession struct {
	Contract *ProxiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ProxiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxiesTransactorSession struct {
	Contract     *ProxiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ProxiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxiesRaw struct {
	Contract *Proxies // Generic contract binding to access the raw methods on
}

// ProxiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxiesCallerRaw struct {
	Contract *ProxiesCaller // Generic read-only contract binding to access the raw methods on
}

// ProxiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxiesTransactorRaw struct {
	Contract *ProxiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxies creates a new instance of Proxies, bound to a specific deployed contract.
func NewProxies(address common.Address, backend bind.ContractBackend) (*Proxies, error) {
	contract, err := bindProxies(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxies{ProxiesCaller: ProxiesCaller{contract: contract}, ProxiesTransactor: ProxiesTransactor{contract: contract}, ProxiesFilterer: ProxiesFilterer{contract: contract}}, nil
}

// NewProxiesCaller creates a new read-only instance of Proxies, bound to a specific deployed contract.
func NewProxiesCaller(address common.Address, caller bind.ContractCaller) (*ProxiesCaller, error) {
	contract, err := bindProxies(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxiesCaller{contract: contract}, nil
}

// NewProxiesTransactor creates a new write-only instance of Proxies, bound to a specific deployed contract.
func NewProxiesTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxiesTransactor, error) {
	contract, err := bindProxies(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxiesTransactor{contract: contract}, nil
}

// NewProxiesFilterer creates a new log filterer instance of Proxies, bound to a specific deployed contract.
func NewProxiesFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxiesFilterer, error) {
	contract, err := bindProxies(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxiesFilterer{contract: contract}, nil
}

// bindProxies binds a generic wrapper to an already deployed contract.
func bindProxies(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxiesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxies *ProxiesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxies.Contract.ProxiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxies *ProxiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxies.Contract.ProxiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxies *ProxiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxies.Contract.ProxiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxies *ProxiesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxies.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxies *ProxiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxies.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxies *ProxiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxies.Contract.contract.Transact(opts, method, params...)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxies *ProxiesTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Proxies.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxies *ProxiesSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxies.Contract.Fallback(&_Proxies.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Proxies *ProxiesTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Proxies.Contract.Fallback(&_Proxies.TransactOpts, calldata)
}
