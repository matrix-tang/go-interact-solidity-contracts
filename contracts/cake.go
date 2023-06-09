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

// CakeMetaData contains all meta data concerning the Cake contract.
var CakeMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromDelegate\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"toDelegate\",\"type\":\"address\"}],\"name\":\"DelegateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBalance\",\"type\":\"uint256\"}],\"name\":\"DelegateVotesChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DELEGATION_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"checkpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"fromBlock\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"}],\"name\":\"delegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegatee\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"delegateBySig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"delegator\",\"type\":\"address\"}],\"name\":\"delegates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getCurrentVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"getPriorVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"numCheckpoints\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// CakeABI is the input ABI used to generate the binding from.
// Deprecated: Use CakeMetaData.ABI instead.
var CakeABI = CakeMetaData.ABI

// Cake is an auto generated Go binding around an Ethereum contract.
type Cake struct {
	CakeCaller     // Read-only binding to the contract
	CakeTransactor // Write-only binding to the contract
	CakeFilterer   // Log filterer for contract events
}

// CakeCaller is an auto generated read-only Go binding around an Ethereum contract.
type CakeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CakeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CakeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CakeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CakeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CakeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CakeSession struct {
	Contract     *Cake             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CakeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CakeCallerSession struct {
	Contract *CakeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CakeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CakeTransactorSession struct {
	Contract     *CakeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CakeRaw is an auto generated low-level Go binding around an Ethereum contract.
type CakeRaw struct {
	Contract *Cake // Generic contract binding to access the raw methods on
}

// CakeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CakeCallerRaw struct {
	Contract *CakeCaller // Generic read-only contract binding to access the raw methods on
}

// CakeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CakeTransactorRaw struct {
	Contract *CakeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCake creates a new instance of Cake, bound to a specific deployed contract.
func NewCake(address common.Address, backend bind.ContractBackend) (*Cake, error) {
	contract, err := bindCake(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Cake{CakeCaller: CakeCaller{contract: contract}, CakeTransactor: CakeTransactor{contract: contract}, CakeFilterer: CakeFilterer{contract: contract}}, nil
}

// NewCakeCaller creates a new read-only instance of Cake, bound to a specific deployed contract.
func NewCakeCaller(address common.Address, caller bind.ContractCaller) (*CakeCaller, error) {
	contract, err := bindCake(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CakeCaller{contract: contract}, nil
}

// NewCakeTransactor creates a new write-only instance of Cake, bound to a specific deployed contract.
func NewCakeTransactor(address common.Address, transactor bind.ContractTransactor) (*CakeTransactor, error) {
	contract, err := bindCake(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CakeTransactor{contract: contract}, nil
}

// NewCakeFilterer creates a new log filterer instance of Cake, bound to a specific deployed contract.
func NewCakeFilterer(address common.Address, filterer bind.ContractFilterer) (*CakeFilterer, error) {
	contract, err := bindCake(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CakeFilterer{contract: contract}, nil
}

// bindCake binds a generic wrapper to an already deployed contract.
func bindCake(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CakeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cake *CakeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cake.Contract.CakeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cake *CakeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cake.Contract.CakeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cake *CakeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cake.Contract.CakeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Cake *CakeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Cake.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Cake *CakeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cake.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Cake *CakeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Cake.Contract.contract.Transact(opts, method, params...)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Cake *CakeCaller) DELEGATIONTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "DELEGATION_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Cake *CakeSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _Cake.Contract.DELEGATIONTYPEHASH(&_Cake.CallOpts)
}

// DELEGATIONTYPEHASH is a free data retrieval call binding the contract method 0xe7a324dc.
//
// Solidity: function DELEGATION_TYPEHASH() view returns(bytes32)
func (_Cake *CakeCallerSession) DELEGATIONTYPEHASH() ([32]byte, error) {
	return _Cake.Contract.DELEGATIONTYPEHASH(&_Cake.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Cake *CakeCaller) DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Cake *CakeSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Cake.Contract.DOMAINTYPEHASH(&_Cake.CallOpts)
}

// DOMAINTYPEHASH is a free data retrieval call binding the contract method 0x20606b70.
//
// Solidity: function DOMAIN_TYPEHASH() view returns(bytes32)
func (_Cake *CakeCallerSession) DOMAINTYPEHASH() ([32]byte, error) {
	return _Cake.Contract.DOMAINTYPEHASH(&_Cake.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cake *CakeCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cake *CakeSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cake.Contract.Allowance(&_Cake.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Cake *CakeCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Cake.Contract.Allowance(&_Cake.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cake *CakeCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cake *CakeSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Cake.Contract.BalanceOf(&_Cake.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Cake *CakeCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Cake.Contract.BalanceOf(&_Cake.CallOpts, account)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_Cake *CakeCaller) Checkpoints(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "checkpoints", arg0, arg1)

	outstruct := new(struct {
		FromBlock uint32
		Votes     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromBlock = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Votes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_Cake *CakeSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _Cake.Contract.Checkpoints(&_Cake.CallOpts, arg0, arg1)
}

// Checkpoints is a free data retrieval call binding the contract method 0xf1127ed8.
//
// Solidity: function checkpoints(address , uint32 ) view returns(uint32 fromBlock, uint256 votes)
func (_Cake *CakeCallerSession) Checkpoints(arg0 common.Address, arg1 uint32) (struct {
	FromBlock uint32
	Votes     *big.Int
}, error) {
	return _Cake.Contract.Checkpoints(&_Cake.CallOpts, arg0, arg1)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cake *CakeCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cake *CakeSession) Decimals() (uint8, error) {
	return _Cake.Contract.Decimals(&_Cake.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Cake *CakeCallerSession) Decimals() (uint8, error) {
	return _Cake.Contract.Decimals(&_Cake.CallOpts)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Cake *CakeCaller) Delegates(opts *bind.CallOpts, delegator common.Address) (common.Address, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "delegates", delegator)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Cake *CakeSession) Delegates(delegator common.Address) (common.Address, error) {
	return _Cake.Contract.Delegates(&_Cake.CallOpts, delegator)
}

// Delegates is a free data retrieval call binding the contract method 0x587cde1e.
//
// Solidity: function delegates(address delegator) view returns(address)
func (_Cake *CakeCallerSession) Delegates(delegator common.Address) (common.Address, error) {
	return _Cake.Contract.Delegates(&_Cake.CallOpts, delegator)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_Cake *CakeCaller) GetCurrentVotes(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "getCurrentVotes", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_Cake *CakeSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _Cake.Contract.GetCurrentVotes(&_Cake.CallOpts, account)
}

// GetCurrentVotes is a free data retrieval call binding the contract method 0xb4b5ea57.
//
// Solidity: function getCurrentVotes(address account) view returns(uint256)
func (_Cake *CakeCallerSession) GetCurrentVotes(account common.Address) (*big.Int, error) {
	return _Cake.Contract.GetCurrentVotes(&_Cake.CallOpts, account)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Cake *CakeCaller) GetOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "getOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Cake *CakeSession) GetOwner() (common.Address, error) {
	return _Cake.Contract.GetOwner(&_Cake.CallOpts)
}

// GetOwner is a free data retrieval call binding the contract method 0x893d20e8.
//
// Solidity: function getOwner() view returns(address)
func (_Cake *CakeCallerSession) GetOwner() (common.Address, error) {
	return _Cake.Contract.GetOwner(&_Cake.CallOpts)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Cake *CakeCaller) GetPriorVotes(opts *bind.CallOpts, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "getPriorVotes", account, blockNumber)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Cake *CakeSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _Cake.Contract.GetPriorVotes(&_Cake.CallOpts, account, blockNumber)
}

// GetPriorVotes is a free data retrieval call binding the contract method 0x782d6fe1.
//
// Solidity: function getPriorVotes(address account, uint256 blockNumber) view returns(uint256)
func (_Cake *CakeCallerSession) GetPriorVotes(account common.Address, blockNumber *big.Int) (*big.Int, error) {
	return _Cake.Contract.GetPriorVotes(&_Cake.CallOpts, account, blockNumber)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cake *CakeCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cake *CakeSession) Name() (string, error) {
	return _Cake.Contract.Name(&_Cake.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Cake *CakeCallerSession) Name() (string, error) {
	return _Cake.Contract.Name(&_Cake.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Cake *CakeCaller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Cake *CakeSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Cake.Contract.Nonces(&_Cake.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Cake *CakeCallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Cake.Contract.Nonces(&_Cake.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Cake *CakeCaller) NumCheckpoints(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "numCheckpoints", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Cake *CakeSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _Cake.Contract.NumCheckpoints(&_Cake.CallOpts, arg0)
}

// NumCheckpoints is a free data retrieval call binding the contract method 0x6fcfff45.
//
// Solidity: function numCheckpoints(address ) view returns(uint32)
func (_Cake *CakeCallerSession) NumCheckpoints(arg0 common.Address) (uint32, error) {
	return _Cake.Contract.NumCheckpoints(&_Cake.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cake *CakeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cake *CakeSession) Owner() (common.Address, error) {
	return _Cake.Contract.Owner(&_Cake.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Cake *CakeCallerSession) Owner() (common.Address, error) {
	return _Cake.Contract.Owner(&_Cake.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cake *CakeCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cake *CakeSession) Symbol() (string, error) {
	return _Cake.Contract.Symbol(&_Cake.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Cake *CakeCallerSession) Symbol() (string, error) {
	return _Cake.Contract.Symbol(&_Cake.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cake *CakeCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Cake.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cake *CakeSession) TotalSupply() (*big.Int, error) {
	return _Cake.Contract.TotalSupply(&_Cake.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Cake *CakeCallerSession) TotalSupply() (*big.Int, error) {
	return _Cake.Contract.TotalSupply(&_Cake.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cake *CakeTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cake *CakeSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Approve(&_Cake.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Cake *CakeTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Approve(&_Cake.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cake *CakeTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cake *CakeSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.DecreaseAllowance(&_Cake.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Cake *CakeTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.DecreaseAllowance(&_Cake.TransactOpts, spender, subtractedValue)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Cake *CakeTransactor) Delegate(opts *bind.TransactOpts, delegatee common.Address) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "delegate", delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Cake *CakeSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Cake.Contract.Delegate(&_Cake.TransactOpts, delegatee)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address delegatee) returns()
func (_Cake *CakeTransactorSession) Delegate(delegatee common.Address) (*types.Transaction, error) {
	return _Cake.Contract.Delegate(&_Cake.TransactOpts, delegatee)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Cake *CakeTransactor) DelegateBySig(opts *bind.TransactOpts, delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "delegateBySig", delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Cake *CakeSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Cake.Contract.DelegateBySig(&_Cake.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// DelegateBySig is a paid mutator transaction binding the contract method 0xc3cda520.
//
// Solidity: function delegateBySig(address delegatee, uint256 nonce, uint256 expiry, uint8 v, bytes32 r, bytes32 s) returns()
func (_Cake *CakeTransactorSession) DelegateBySig(delegatee common.Address, nonce *big.Int, expiry *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Cake.Contract.DelegateBySig(&_Cake.TransactOpts, delegatee, nonce, expiry, v, r, s)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cake *CakeTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cake *CakeSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.IncreaseAllowance(&_Cake.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Cake *CakeTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.IncreaseAllowance(&_Cake.TransactOpts, spender, addedValue)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Cake *CakeTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Cake *CakeSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Mint(&_Cake.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address _to, uint256 _amount) returns()
func (_Cake *CakeTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Mint(&_Cake.TransactOpts, _to, _amount)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_Cake *CakeTransactor) Mint0(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "mint0", amount)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_Cake *CakeSession) Mint0(amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Mint0(&_Cake.TransactOpts, amount)
}

// Mint0 is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 amount) returns(bool)
func (_Cake *CakeTransactorSession) Mint0(amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Mint0(&_Cake.TransactOpts, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cake *CakeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cake *CakeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cake.Contract.RenounceOwnership(&_Cake.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Cake *CakeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Cake.Contract.RenounceOwnership(&_Cake.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cake *CakeTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cake *CakeSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Transfer(&_Cake.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Cake *CakeTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.Transfer(&_Cake.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cake *CakeTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cake *CakeSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.TransferFrom(&_Cake.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Cake *CakeTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Cake.Contract.TransferFrom(&_Cake.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cake *CakeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Cake.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cake *CakeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cake.Contract.TransferOwnership(&_Cake.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Cake *CakeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Cake.Contract.TransferOwnership(&_Cake.TransactOpts, newOwner)
}

// CakeApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Cake contract.
type CakeApprovalIterator struct {
	Event *CakeApproval // Event containing the contract specifics and raw log

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
func (it *CakeApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakeApproval)
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
		it.Event = new(CakeApproval)
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
func (it *CakeApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakeApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakeApproval represents a Approval event raised by the Cake contract.
type CakeApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Cake *CakeFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CakeApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cake.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CakeApprovalIterator{contract: _Cake.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Cake *CakeFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CakeApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Cake.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakeApproval)
				if err := _Cake.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Cake *CakeFilterer) ParseApproval(log types.Log) (*CakeApproval, error) {
	event := new(CakeApproval)
	if err := _Cake.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakeDelegateChangedIterator is returned from FilterDelegateChanged and is used to iterate over the raw logs and unpacked data for DelegateChanged events raised by the Cake contract.
type CakeDelegateChangedIterator struct {
	Event *CakeDelegateChanged // Event containing the contract specifics and raw log

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
func (it *CakeDelegateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakeDelegateChanged)
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
		it.Event = new(CakeDelegateChanged)
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
func (it *CakeDelegateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakeDelegateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakeDelegateChanged represents a DelegateChanged event raised by the Cake contract.
type CakeDelegateChanged struct {
	Delegator    common.Address
	FromDelegate common.Address
	ToDelegate   common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDelegateChanged is a free log retrieval operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Cake *CakeFilterer) FilterDelegateChanged(opts *bind.FilterOpts, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (*CakeDelegateChangedIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _Cake.contract.FilterLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return &CakeDelegateChangedIterator{contract: _Cake.contract, event: "DelegateChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateChanged is a free log subscription operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Cake *CakeFilterer) WatchDelegateChanged(opts *bind.WatchOpts, sink chan<- *CakeDelegateChanged, delegator []common.Address, fromDelegate []common.Address, toDelegate []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var fromDelegateRule []interface{}
	for _, fromDelegateItem := range fromDelegate {
		fromDelegateRule = append(fromDelegateRule, fromDelegateItem)
	}
	var toDelegateRule []interface{}
	for _, toDelegateItem := range toDelegate {
		toDelegateRule = append(toDelegateRule, toDelegateItem)
	}

	logs, sub, err := _Cake.contract.WatchLogs(opts, "DelegateChanged", delegatorRule, fromDelegateRule, toDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakeDelegateChanged)
				if err := _Cake.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
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

// ParseDelegateChanged is a log parse operation binding the contract event 0x3134e8a2e6d97e929a7e54011ea5485d7d196dd5f0ba4d4ef95803e8e3fc257f.
//
// Solidity: event DelegateChanged(address indexed delegator, address indexed fromDelegate, address indexed toDelegate)
func (_Cake *CakeFilterer) ParseDelegateChanged(log types.Log) (*CakeDelegateChanged, error) {
	event := new(CakeDelegateChanged)
	if err := _Cake.contract.UnpackLog(event, "DelegateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakeDelegateVotesChangedIterator is returned from FilterDelegateVotesChanged and is used to iterate over the raw logs and unpacked data for DelegateVotesChanged events raised by the Cake contract.
type CakeDelegateVotesChangedIterator struct {
	Event *CakeDelegateVotesChanged // Event containing the contract specifics and raw log

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
func (it *CakeDelegateVotesChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakeDelegateVotesChanged)
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
		it.Event = new(CakeDelegateVotesChanged)
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
func (it *CakeDelegateVotesChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakeDelegateVotesChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakeDelegateVotesChanged represents a DelegateVotesChanged event raised by the Cake contract.
type CakeDelegateVotesChanged struct {
	Delegate        common.Address
	PreviousBalance *big.Int
	NewBalance      *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDelegateVotesChanged is a free log retrieval operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Cake *CakeFilterer) FilterDelegateVotesChanged(opts *bind.FilterOpts, delegate []common.Address) (*CakeDelegateVotesChangedIterator, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Cake.contract.FilterLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return &CakeDelegateVotesChangedIterator{contract: _Cake.contract, event: "DelegateVotesChanged", logs: logs, sub: sub}, nil
}

// WatchDelegateVotesChanged is a free log subscription operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Cake *CakeFilterer) WatchDelegateVotesChanged(opts *bind.WatchOpts, sink chan<- *CakeDelegateVotesChanged, delegate []common.Address) (event.Subscription, error) {

	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _Cake.contract.WatchLogs(opts, "DelegateVotesChanged", delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakeDelegateVotesChanged)
				if err := _Cake.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
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

// ParseDelegateVotesChanged is a log parse operation binding the contract event 0xdec2bacdd2f05b59de34da9b523dff8be42e5e38e818c82fdb0bae774387a724.
//
// Solidity: event DelegateVotesChanged(address indexed delegate, uint256 previousBalance, uint256 newBalance)
func (_Cake *CakeFilterer) ParseDelegateVotesChanged(log types.Log) (*CakeDelegateVotesChanged, error) {
	event := new(CakeDelegateVotesChanged)
	if err := _Cake.contract.UnpackLog(event, "DelegateVotesChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Cake contract.
type CakeOwnershipTransferredIterator struct {
	Event *CakeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CakeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakeOwnershipTransferred)
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
		it.Event = new(CakeOwnershipTransferred)
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
func (it *CakeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakeOwnershipTransferred represents a OwnershipTransferred event raised by the Cake contract.
type CakeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cake *CakeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CakeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cake.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CakeOwnershipTransferredIterator{contract: _Cake.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cake *CakeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CakeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Cake.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakeOwnershipTransferred)
				if err := _Cake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Cake *CakeFilterer) ParseOwnershipTransferred(log types.Log) (*CakeOwnershipTransferred, error) {
	event := new(CakeOwnershipTransferred)
	if err := _Cake.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CakeTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Cake contract.
type CakeTransferIterator struct {
	Event *CakeTransfer // Event containing the contract specifics and raw log

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
func (it *CakeTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CakeTransfer)
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
		it.Event = new(CakeTransfer)
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
func (it *CakeTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CakeTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CakeTransfer represents a Transfer event raised by the Cake contract.
type CakeTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cake *CakeFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CakeTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cake.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CakeTransferIterator{contract: _Cake.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cake *CakeFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CakeTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Cake.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CakeTransfer)
				if err := _Cake.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Cake *CakeFilterer) ParseTransfer(log types.Log) (*CakeTransfer, error) {
	event := new(CakeTransfer)
	if err := _Cake.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
