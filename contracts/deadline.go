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

// DeadlineMetaData contains all meta data concerning the Deadline go-interact-solidity-contracts.
var DeadlineMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"dl\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"show\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DeadlineABI is the input ABI used to generate the binding from.
// Deprecated: Use DeadlineMetaData.ABI instead.
var DeadlineABI = DeadlineMetaData.ABI

// Deadline is an auto generated Go binding around an Ethereum go-interact-solidity-contracts.
type Deadline struct {
	DeadlineCaller     // Read-only binding to the go-interact-solidity-contracts
	DeadlineTransactor // Write-only binding to the go-interact-solidity-contracts
	DeadlineFilterer   // Log filterer for go-interact-solidity-contracts events
}

// DeadlineCaller is an auto generated read-only Go binding around an Ethereum go-interact-solidity-contracts.
type DeadlineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeadlineTransactor is an auto generated write-only Go binding around an Ethereum go-interact-solidity-contracts.
type DeadlineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeadlineFilterer is an auto generated log filtering Go binding around an Ethereum go-interact-solidity-contracts events.
type DeadlineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeadlineSession is an auto generated Go binding around an Ethereum go-interact-solidity-contracts,
// with pre-set call and transact options.
type DeadlineSession struct {
	Contract     *Deadline         // Generic go-interact-solidity-contracts binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeadlineCallerSession is an auto generated read-only Go binding around an Ethereum go-interact-solidity-contracts,
// with pre-set call options.
type DeadlineCallerSession struct {
	Contract *DeadlineCaller // Generic go-interact-solidity-contracts caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DeadlineTransactorSession is an auto generated write-only Go binding around an Ethereum go-interact-solidity-contracts,
// with pre-set transact options.
type DeadlineTransactorSession struct {
	Contract     *DeadlineTransactor // Generic go-interact-solidity-contracts transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DeadlineRaw is an auto generated low-level Go binding around an Ethereum go-interact-solidity-contracts.
type DeadlineRaw struct {
	Contract *Deadline // Generic go-interact-solidity-contracts binding to access the raw methods on
}

// DeadlineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum go-interact-solidity-contracts.
type DeadlineCallerRaw struct {
	Contract *DeadlineCaller // Generic read-only go-interact-solidity-contracts binding to access the raw methods on
}

// DeadlineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum go-interact-solidity-contracts.
type DeadlineTransactorRaw struct {
	Contract *DeadlineTransactor // Generic write-only go-interact-solidity-contracts binding to access the raw methods on
}

// NewDeadline creates a new instance of Deadline, bound to a specific deployed go-interact-solidity-contracts.
func NewDeadline(address common.Address, backend bind.ContractBackend) (*Deadline, error) {
	contract, err := bindDeadline(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deadline{DeadlineCaller: DeadlineCaller{contract: contract}, DeadlineTransactor: DeadlineTransactor{contract: contract}, DeadlineFilterer: DeadlineFilterer{contract: contract}}, nil
}

// NewDeadlineCaller creates a new read-only instance of Deadline, bound to a specific deployed go-interact-solidity-contracts.
func NewDeadlineCaller(address common.Address, caller bind.ContractCaller) (*DeadlineCaller, error) {
	contract, err := bindDeadline(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeadlineCaller{contract: contract}, nil
}

// NewDeadlineTransactor creates a new write-only instance of Deadline, bound to a specific deployed go-interact-solidity-contracts.
func NewDeadlineTransactor(address common.Address, transactor bind.ContractTransactor) (*DeadlineTransactor, error) {
	contract, err := bindDeadline(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeadlineTransactor{contract: contract}, nil
}

// NewDeadlineFilterer creates a new log filterer instance of Deadline, bound to a specific deployed go-interact-solidity-contracts.
func NewDeadlineFilterer(address common.Address, filterer bind.ContractFilterer) (*DeadlineFilterer, error) {
	contract, err := bindDeadline(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeadlineFilterer{contract: contract}, nil
}

// bindDeadline binds a generic wrapper to an already deployed go-interact-solidity-contracts.
func bindDeadline(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DeadlineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) go-interact-solidity-contracts method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deadline *DeadlineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deadline.Contract.DeadlineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the go-interact-solidity-contracts, calling
// its default method if one is available.
func (_Deadline *DeadlineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deadline.Contract.DeadlineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) go-interact-solidity-contracts method with params as input values.
func (_Deadline *DeadlineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deadline.Contract.DeadlineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) go-interact-solidity-contracts method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deadline *DeadlineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deadline.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the go-interact-solidity-contracts, calling
// its default method if one is available.
func (_Deadline *DeadlineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deadline.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) go-interact-solidity-contracts method with params as input values.
func (_Deadline *DeadlineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deadline.Contract.contract.Transact(opts, method, params...)
}

// Dl is a free data retrieval call binding the go-interact-solidity-contracts method 0x24a85aac.
//
// Solidity: function dl() view returns(uint256)
func (_Deadline *DeadlineCaller) Dl(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deadline.contract.Call(opts, &out, "dl")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Dl is a free data retrieval call binding the go-interact-solidity-contracts method 0x24a85aac.
//
// Solidity: function dl() view returns(uint256)
func (_Deadline *DeadlineSession) Dl() (*big.Int, error) {
	return _Deadline.Contract.Dl(&_Deadline.CallOpts)
}

// Dl is a free data retrieval call binding the go-interact-solidity-contracts method 0x24a85aac.
//
// Solidity: function dl() view returns(uint256)
func (_Deadline *DeadlineCallerSession) Dl() (*big.Int, error) {
	return _Deadline.Contract.Dl(&_Deadline.CallOpts)
}

// Show is a free data retrieval call binding the go-interact-solidity-contracts method 0xcc80f6f3.
//
// Solidity: function show() view returns(uint256)
func (_Deadline *DeadlineCaller) Show(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Deadline.contract.Call(opts, &out, "show")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Show is a free data retrieval call binding the go-interact-solidity-contracts method 0xcc80f6f3.
//
// Solidity: function show() view returns(uint256)
func (_Deadline *DeadlineSession) Show() (*big.Int, error) {
	return _Deadline.Contract.Show(&_Deadline.CallOpts)
}

// Show is a free data retrieval call binding the go-interact-solidity-contracts method 0xcc80f6f3.
//
// Solidity: function show() view returns(uint256)
func (_Deadline *DeadlineCallerSession) Show() (*big.Int, error) {
	return _Deadline.Contract.Show(&_Deadline.CallOpts)
}
