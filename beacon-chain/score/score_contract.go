// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package score

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

// ScoreContractMetaData contains all meta data concerning the ScoreContract contract.
var ScoreContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"name\":\"ScoreUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"ValidatorRegistered\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endBlock\",\"type\":\"uint256\"}],\"name\":\"GetEpochRangeScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"GetScore\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"GetScoreData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"IsValidatorRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"}],\"name\":\"RegisterValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYSTEM\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"SetTargetValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"pubKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"}],\"name\":\"UpdateScore\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"blockAggregates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalScore\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"scores\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"score\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastUpdatedBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"targetValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ScoreContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ScoreContractMetaData.ABI instead.
var ScoreContractABI = ScoreContractMetaData.ABI

// ScoreContract is an auto generated Go binding around an Ethereum contract.
type ScoreContract struct {
	ScoreContractCaller     // Read-only binding to the contract
	ScoreContractTransactor // Write-only binding to the contract
	ScoreContractFilterer   // Log filterer for contract events
}

// ScoreContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ScoreContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScoreContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ScoreContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScoreContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ScoreContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ScoreContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ScoreContractSession struct {
	Contract     *ScoreContract    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ScoreContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ScoreContractCallerSession struct {
	Contract *ScoreContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// ScoreContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ScoreContractTransactorSession struct {
	Contract     *ScoreContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// ScoreContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ScoreContractRaw struct {
	Contract *ScoreContract // Generic contract binding to access the raw methods on
}

// ScoreContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ScoreContractCallerRaw struct {
	Contract *ScoreContractCaller // Generic read-only contract binding to access the raw methods on
}

// ScoreContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ScoreContractTransactorRaw struct {
	Contract *ScoreContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewScoreContract creates a new instance of ScoreContract, bound to a specific deployed contract.
func NewScoreContract(address common.Address, backend bind.ContractBackend) (*ScoreContract, error) {
	contract, err := bindScoreContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ScoreContract{ScoreContractCaller: ScoreContractCaller{contract: contract}, ScoreContractTransactor: ScoreContractTransactor{contract: contract}, ScoreContractFilterer: ScoreContractFilterer{contract: contract}}, nil
}

// NewScoreContractCaller creates a new read-only instance of ScoreContract, bound to a specific deployed contract.
func NewScoreContractCaller(address common.Address, caller bind.ContractCaller) (*ScoreContractCaller, error) {
	contract, err := bindScoreContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ScoreContractCaller{contract: contract}, nil
}

// NewScoreContractTransactor creates a new write-only instance of ScoreContract, bound to a specific deployed contract.
func NewScoreContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ScoreContractTransactor, error) {
	contract, err := bindScoreContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ScoreContractTransactor{contract: contract}, nil
}

// NewScoreContractFilterer creates a new log filterer instance of ScoreContract, bound to a specific deployed contract.
func NewScoreContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ScoreContractFilterer, error) {
	contract, err := bindScoreContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ScoreContractFilterer{contract: contract}, nil
}

// bindScoreContract binds a generic wrapper to an already deployed contract.
func bindScoreContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ScoreContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScoreContract *ScoreContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScoreContract.Contract.ScoreContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScoreContract *ScoreContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScoreContract.Contract.ScoreContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScoreContract *ScoreContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScoreContract.Contract.ScoreContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ScoreContract *ScoreContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ScoreContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ScoreContract *ScoreContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ScoreContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ScoreContract *ScoreContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ScoreContract.Contract.contract.Transact(opts, method, params...)
}

// GetEpochRangeScore is a free data retrieval call binding the contract method 0xb7e3707f.
//
// Solidity: function GetEpochRangeScore(uint256 startBlock, uint256 endBlock) view returns(uint256)
func (_ScoreContract *ScoreContractCaller) GetEpochRangeScore(opts *bind.CallOpts, startBlock *big.Int, endBlock *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ScoreContract.contract.Call(opts, &out, "GetEpochRangeScore", startBlock, endBlock)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEpochRangeScore is a free data retrieval call binding the contract method 0xb7e3707f.
//
// Solidity: function GetEpochRangeScore(uint256 startBlock, uint256 endBlock) view returns(uint256)
func (_ScoreContract *ScoreContractSession) GetEpochRangeScore(startBlock *big.Int, endBlock *big.Int) (*big.Int, error) {
	return _ScoreContract.Contract.GetEpochRangeScore(&_ScoreContract.CallOpts, startBlock, endBlock)
}

// GetEpochRangeScore is a free data retrieval call binding the contract method 0xb7e3707f.
//
// Solidity: function GetEpochRangeScore(uint256 startBlock, uint256 endBlock) view returns(uint256)
func (_ScoreContract *ScoreContractCallerSession) GetEpochRangeScore(startBlock *big.Int, endBlock *big.Int) (*big.Int, error) {
	return _ScoreContract.Contract.GetEpochRangeScore(&_ScoreContract.CallOpts, startBlock, endBlock)
}

// GetScore is a free data retrieval call binding the contract method 0xc274e070.
//
// Solidity: function GetScore(bytes pubKey) view returns(uint256)
func (_ScoreContract *ScoreContractCaller) GetScore(opts *bind.CallOpts, pubKey []byte) (*big.Int, error) {
	var out []interface{}
	err := _ScoreContract.contract.Call(opts, &out, "GetScore", pubKey)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetScore is a free data retrieval call binding the contract method 0xc274e070.
//
// Solidity: function GetScore(bytes pubKey) view returns(uint256)
func (_ScoreContract *ScoreContractSession) GetScore(pubKey []byte) (*big.Int, error) {
	return _ScoreContract.Contract.GetScore(&_ScoreContract.CallOpts, pubKey)
}

// GetScore is a free data retrieval call binding the contract method 0xc274e070.
//
// Solidity: function GetScore(bytes pubKey) view returns(uint256)
func (_ScoreContract *ScoreContractCallerSession) GetScore(pubKey []byte) (*big.Int, error) {
	return _ScoreContract.Contract.GetScore(&_ScoreContract.CallOpts, pubKey)
}

// GetScoreData is a free data retrieval call binding the contract method 0xd4c71f48.
//
// Solidity: function GetScoreData(bytes pubKey) view returns(uint256 score, uint256 lastUpdatedBlock)
func (_ScoreContract *ScoreContractCaller) GetScoreData(opts *bind.CallOpts, pubKey []byte) (struct {
	Score            *big.Int
	LastUpdatedBlock *big.Int
}, error) {
	var out []interface{}
	err := _ScoreContract.contract.Call(opts, &out, "GetScoreData", pubKey)

	outstruct := new(struct {
		Score            *big.Int
		LastUpdatedBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Score = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastUpdatedBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetScoreData is a free data retrieval call binding the contract method 0xd4c71f48.
//
// Solidity: function GetScoreData(bytes pubKey) view returns(uint256 score, uint256 lastUpdatedBlock)
func (_ScoreContract *ScoreContractSession) GetScoreData(pubKey []byte) (struct {
	Score            *big.Int
	LastUpdatedBlock *big.Int
}, error) {
	return _ScoreContract.Contract.GetScoreData(&_ScoreContract.CallOpts, pubKey)
}

// GetScoreData is a free data retrieval call binding the contract method 0xd4c71f48.
//
// Solidity: function GetScoreData(bytes pubKey) view returns(uint256 score, uint256 lastUpdatedBlock)
func (_ScoreContract *ScoreContractCallerSession) GetScoreData(pubKey []byte) (struct {
	Score            *big.Int
	LastUpdatedBlock *big.Int
}, error) {
	return _ScoreContract.Contract.GetScoreData(&_ScoreContract.CallOpts, pubKey)
}

// IsValidatorRegistered is a free data retrieval call binding the contract method 0x715ca973.
//
// Solidity: function IsValidatorRegistered(bytes pubKey) view returns(bool)
func (_ScoreContract *ScoreContractCaller) IsValidatorRegistered(opts *bind.CallOpts, pubKey []byte) (bool, error) {
	var out []interface{}
	err := _ScoreContract.contract.Call(opts, &out, "IsValidatorRegistered", pubKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorRegistered is a free data retrieval call binding the contract method 0x715ca973.
//
// Solidity: function IsValidatorRegistered(bytes pubKey) view returns(bool)
func (_ScoreContract *ScoreContractSession) IsValidatorRegistered(pubKey []byte) (bool, error) {
	return _ScoreContract.Contract.IsValidatorRegistered(&_ScoreContract.CallOpts, pubKey)
}

// IsValidatorRegistered is a free data retrieval call binding the contract method 0x715ca973.
//
// Solidity: function IsValidatorRegistered(bytes pubKey) view returns(bool)
func (_ScoreContract *ScoreContractCallerSession) IsValidatorRegistered(pubKey []byte) (bool, error) {
	return _ScoreContract.Contract.IsValidatorRegistered(&_ScoreContract.CallOpts, pubKey)
}

// SYSTEM is a free data retrieval call binding the contract method 0x51351d53.
//
// Solidity: function SYSTEM() view returns(address)
func (_ScoreContract *ScoreContractCaller) SYSTEM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ScoreContract.contract.Call(opts, &out, "SYSTEM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SYSTEM is a free data retrieval call binding the contract method 0x51351d53.
//
// Solidity: function SYSTEM() view returns(address)
func (_ScoreContract *ScoreContractSession) SYSTEM() (common.Address, error) {
	return _ScoreContract.Contract.SYSTEM(&_ScoreContract.CallOpts)
}

// SYSTEM is a free data retrieval call binding the contract method 0x51351d53.
//
// Solidity: function SYSTEM() view returns(address)
func (_ScoreContract *ScoreContractCallerSession) SYSTEM() (common.Address, error) {
	return _ScoreContract.Contract.SYSTEM(&_ScoreContract.CallOpts)
}

// BlockAggregates is a free data retrieval call binding the contract method 0x5152c058.
//
// Solidity: function blockAggregates(uint256 ) view returns(uint256 totalScore, uint256 count)
func (_ScoreContract *ScoreContractCaller) BlockAggregates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TotalScore *big.Int
	Count      *big.Int
}, error) {
	var out []interface{}
	err := _ScoreContract.contract.Call(opts, &out, "blockAggregates", arg0)

	outstruct := new(struct {
		TotalScore *big.Int
		Count      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalScore = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Count = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// BlockAggregates is a free data retrieval call binding the contract method 0x5152c058.
//
// Solidity: function blockAggregates(uint256 ) view returns(uint256 totalScore, uint256 count)
func (_ScoreContract *ScoreContractSession) BlockAggregates(arg0 *big.Int) (struct {
	TotalScore *big.Int
	Count      *big.Int
}, error) {
	return _ScoreContract.Contract.BlockAggregates(&_ScoreContract.CallOpts, arg0)
}

// BlockAggregates is a free data retrieval call binding the contract method 0x5152c058.
//
// Solidity: function blockAggregates(uint256 ) view returns(uint256 totalScore, uint256 count)
func (_ScoreContract *ScoreContractCallerSession) BlockAggregates(arg0 *big.Int) (struct {
	TotalScore *big.Int
	Count      *big.Int
}, error) {
	return _ScoreContract.Contract.BlockAggregates(&_ScoreContract.CallOpts, arg0)
}

// Scores is a free data retrieval call binding the contract method 0x9c2aebe0.
//
// Solidity: function scores(bytes ) view returns(uint256 score, uint256 lastUpdatedBlock)
func (_ScoreContract *ScoreContractCaller) Scores(opts *bind.CallOpts, arg0 []byte) (struct {
	Score            *big.Int
	LastUpdatedBlock *big.Int
}, error) {
	var out []interface{}
	err := _ScoreContract.contract.Call(opts, &out, "scores", arg0)

	outstruct := new(struct {
		Score            *big.Int
		LastUpdatedBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Score = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LastUpdatedBlock = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Scores is a free data retrieval call binding the contract method 0x9c2aebe0.
//
// Solidity: function scores(bytes ) view returns(uint256 score, uint256 lastUpdatedBlock)
func (_ScoreContract *ScoreContractSession) Scores(arg0 []byte) (struct {
	Score            *big.Int
	LastUpdatedBlock *big.Int
}, error) {
	return _ScoreContract.Contract.Scores(&_ScoreContract.CallOpts, arg0)
}

// Scores is a free data retrieval call binding the contract method 0x9c2aebe0.
//
// Solidity: function scores(bytes ) view returns(uint256 score, uint256 lastUpdatedBlock)
func (_ScoreContract *ScoreContractCallerSession) Scores(arg0 []byte) (struct {
	Score            *big.Int
	LastUpdatedBlock *big.Int
}, error) {
	return _ScoreContract.Contract.Scores(&_ScoreContract.CallOpts, arg0)
}

// TargetValidatorsCount is a free data retrieval call binding the contract method 0x235d16ac.
//
// Solidity: function targetValidatorsCount() view returns(uint256)
func (_ScoreContract *ScoreContractCaller) TargetValidatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ScoreContract.contract.Call(opts, &out, "targetValidatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TargetValidatorsCount is a free data retrieval call binding the contract method 0x235d16ac.
//
// Solidity: function targetValidatorsCount() view returns(uint256)
func (_ScoreContract *ScoreContractSession) TargetValidatorsCount() (*big.Int, error) {
	return _ScoreContract.Contract.TargetValidatorsCount(&_ScoreContract.CallOpts)
}

// TargetValidatorsCount is a free data retrieval call binding the contract method 0x235d16ac.
//
// Solidity: function targetValidatorsCount() view returns(uint256)
func (_ScoreContract *ScoreContractCallerSession) TargetValidatorsCount() (*big.Int, error) {
	return _ScoreContract.Contract.TargetValidatorsCount(&_ScoreContract.CallOpts)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xe8d0a894.
//
// Solidity: function RegisterValidator(bytes pubKey) returns()
func (_ScoreContract *ScoreContractTransactor) RegisterValidator(opts *bind.TransactOpts, pubKey []byte) (*types.Transaction, error) {
	return _ScoreContract.contract.Transact(opts, "RegisterValidator", pubKey)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xe8d0a894.
//
// Solidity: function RegisterValidator(bytes pubKey) returns()
func (_ScoreContract *ScoreContractSession) RegisterValidator(pubKey []byte) (*types.Transaction, error) {
	return _ScoreContract.Contract.RegisterValidator(&_ScoreContract.TransactOpts, pubKey)
}

// RegisterValidator is a paid mutator transaction binding the contract method 0xe8d0a894.
//
// Solidity: function RegisterValidator(bytes pubKey) returns()
func (_ScoreContract *ScoreContractTransactorSession) RegisterValidator(pubKey []byte) (*types.Transaction, error) {
	return _ScoreContract.Contract.RegisterValidator(&_ScoreContract.TransactOpts, pubKey)
}

// SetTargetValidatorsCount is a paid mutator transaction binding the contract method 0x4f0e1b08.
//
// Solidity: function SetTargetValidatorsCount(uint256 _count) returns()
func (_ScoreContract *ScoreContractTransactor) SetTargetValidatorsCount(opts *bind.TransactOpts, _count *big.Int) (*types.Transaction, error) {
	return _ScoreContract.contract.Transact(opts, "SetTargetValidatorsCount", _count)
}

// SetTargetValidatorsCount is a paid mutator transaction binding the contract method 0x4f0e1b08.
//
// Solidity: function SetTargetValidatorsCount(uint256 _count) returns()
func (_ScoreContract *ScoreContractSession) SetTargetValidatorsCount(_count *big.Int) (*types.Transaction, error) {
	return _ScoreContract.Contract.SetTargetValidatorsCount(&_ScoreContract.TransactOpts, _count)
}

// SetTargetValidatorsCount is a paid mutator transaction binding the contract method 0x4f0e1b08.
//
// Solidity: function SetTargetValidatorsCount(uint256 _count) returns()
func (_ScoreContract *ScoreContractTransactorSession) SetTargetValidatorsCount(_count *big.Int) (*types.Transaction, error) {
	return _ScoreContract.Contract.SetTargetValidatorsCount(&_ScoreContract.TransactOpts, _count)
}

// UpdateScore is a paid mutator transaction binding the contract method 0x9774d1dd.
//
// Solidity: function UpdateScore(bytes pubKey, uint256 score) returns()
func (_ScoreContract *ScoreContractTransactor) UpdateScore(opts *bind.TransactOpts, pubKey []byte, score *big.Int) (*types.Transaction, error) {
	return _ScoreContract.contract.Transact(opts, "UpdateScore", pubKey, score)
}

// UpdateScore is a paid mutator transaction binding the contract method 0x9774d1dd.
//
// Solidity: function UpdateScore(bytes pubKey, uint256 score) returns()
func (_ScoreContract *ScoreContractSession) UpdateScore(pubKey []byte, score *big.Int) (*types.Transaction, error) {
	return _ScoreContract.Contract.UpdateScore(&_ScoreContract.TransactOpts, pubKey, score)
}

// UpdateScore is a paid mutator transaction binding the contract method 0x9774d1dd.
//
// Solidity: function UpdateScore(bytes pubKey, uint256 score) returns()
func (_ScoreContract *ScoreContractTransactorSession) UpdateScore(pubKey []byte, score *big.Int) (*types.Transaction, error) {
	return _ScoreContract.Contract.UpdateScore(&_ScoreContract.TransactOpts, pubKey, score)
}

// ScoreContractScoreUpdatedIterator is returned from FilterScoreUpdated and is used to iterate over the raw logs and unpacked data for ScoreUpdated events raised by the ScoreContract contract.
type ScoreContractScoreUpdatedIterator struct {
	Event *ScoreContractScoreUpdated // Event containing the contract specifics and raw log

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
func (it *ScoreContractScoreUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScoreContractScoreUpdated)
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
		it.Event = new(ScoreContractScoreUpdated)
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
func (it *ScoreContractScoreUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScoreContractScoreUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScoreContractScoreUpdated represents a ScoreUpdated event raised by the ScoreContract contract.
type ScoreContractScoreUpdated struct {
	PubKey      []byte
	Score       *big.Int
	BlockNumber *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterScoreUpdated is a free log retrieval operation binding the contract event 0x76e12d9737c0ef8254d9f11b62f7c00ae0a3bd138a42e1d25011ee7f03d47b37.
//
// Solidity: event ScoreUpdated(bytes pubKey, uint256 score, uint256 blockNumber)
func (_ScoreContract *ScoreContractFilterer) FilterScoreUpdated(opts *bind.FilterOpts) (*ScoreContractScoreUpdatedIterator, error) {

	logs, sub, err := _ScoreContract.contract.FilterLogs(opts, "ScoreUpdated")
	if err != nil {
		return nil, err
	}
	return &ScoreContractScoreUpdatedIterator{contract: _ScoreContract.contract, event: "ScoreUpdated", logs: logs, sub: sub}, nil
}

// WatchScoreUpdated is a free log subscription operation binding the contract event 0x76e12d9737c0ef8254d9f11b62f7c00ae0a3bd138a42e1d25011ee7f03d47b37.
//
// Solidity: event ScoreUpdated(bytes pubKey, uint256 score, uint256 blockNumber)
func (_ScoreContract *ScoreContractFilterer) WatchScoreUpdated(opts *bind.WatchOpts, sink chan<- *ScoreContractScoreUpdated) (event.Subscription, error) {

	logs, sub, err := _ScoreContract.contract.WatchLogs(opts, "ScoreUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScoreContractScoreUpdated)
				if err := _ScoreContract.contract.UnpackLog(event, "ScoreUpdated", log); err != nil {
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

// ParseScoreUpdated is a log parse operation binding the contract event 0x76e12d9737c0ef8254d9f11b62f7c00ae0a3bd138a42e1d25011ee7f03d47b37.
//
// Solidity: event ScoreUpdated(bytes pubKey, uint256 score, uint256 blockNumber)
func (_ScoreContract *ScoreContractFilterer) ParseScoreUpdated(log types.Log) (*ScoreContractScoreUpdated, error) {
	event := new(ScoreContractScoreUpdated)
	if err := _ScoreContract.contract.UnpackLog(event, "ScoreUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ScoreContractValidatorRegisteredIterator is returned from FilterValidatorRegistered and is used to iterate over the raw logs and unpacked data for ValidatorRegistered events raised by the ScoreContract contract.
type ScoreContractValidatorRegisteredIterator struct {
	Event *ScoreContractValidatorRegistered // Event containing the contract specifics and raw log

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
func (it *ScoreContractValidatorRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ScoreContractValidatorRegistered)
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
		it.Event = new(ScoreContractValidatorRegistered)
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
func (it *ScoreContractValidatorRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ScoreContractValidatorRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ScoreContractValidatorRegistered represents a ValidatorRegistered event raised by the ScoreContract contract.
type ScoreContractValidatorRegistered struct {
	PubKey []byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValidatorRegistered is a free log retrieval operation binding the contract event 0x64b6e61d93b7a91e8cc4376183ede0997a27b44fd9dd2f30a866b2a5730efdb1.
//
// Solidity: event ValidatorRegistered(bytes pubKey)
func (_ScoreContract *ScoreContractFilterer) FilterValidatorRegistered(opts *bind.FilterOpts) (*ScoreContractValidatorRegisteredIterator, error) {

	logs, sub, err := _ScoreContract.contract.FilterLogs(opts, "ValidatorRegistered")
	if err != nil {
		return nil, err
	}
	return &ScoreContractValidatorRegisteredIterator{contract: _ScoreContract.contract, event: "ValidatorRegistered", logs: logs, sub: sub}, nil
}

// WatchValidatorRegistered is a free log subscription operation binding the contract event 0x64b6e61d93b7a91e8cc4376183ede0997a27b44fd9dd2f30a866b2a5730efdb1.
//
// Solidity: event ValidatorRegistered(bytes pubKey)
func (_ScoreContract *ScoreContractFilterer) WatchValidatorRegistered(opts *bind.WatchOpts, sink chan<- *ScoreContractValidatorRegistered) (event.Subscription, error) {

	logs, sub, err := _ScoreContract.contract.WatchLogs(opts, "ValidatorRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ScoreContractValidatorRegistered)
				if err := _ScoreContract.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
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

// ParseValidatorRegistered is a log parse operation binding the contract event 0x64b6e61d93b7a91e8cc4376183ede0997a27b44fd9dd2f30a866b2a5730efdb1.
//
// Solidity: event ValidatorRegistered(bytes pubKey)
func (_ScoreContract *ScoreContractFilterer) ParseValidatorRegistered(log types.Log) (*ScoreContractValidatorRegistered, error) {
	event := new(ScoreContractValidatorRegistered)
	if err := _ScoreContract.contract.UnpackLog(event, "ValidatorRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
