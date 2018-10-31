// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

import (
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
	"math/big"
	"strings"
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

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"livro\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nome\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fls\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nascimento\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"profissao\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"estado\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"filiacao\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"drt\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"serie\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numero\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newDataRegistro\",\"type\":\"string\"},{\"name\":\"newProf\",\"type\":\"string\"},{\"name\":\"newNumeroReg\",\"type\":\"string\"},{\"name\":\"newLivro\",\"type\":\"string\"},{\"name\":\"newFls\",\"type\":\"string\"},{\"name\":\"newDRT\",\"type\":\"string\"}],\"name\":\"setProfissao\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numeroRegistro\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newSerie\",\"type\":\"string\"},{\"name\":\"newNome\",\"type\":\"string\"},{\"name\":\"newNasc\",\"type\":\"string\"},{\"name\":\"newEst\",\"type\":\"string\"},{\"name\":\"newData\",\"type\":\"string\"},{\"name\":\"newFil\",\"type\":\"string\"},{\"name\":\"newRG\",\"type\":\"string\"}],\"name\":\"setFormulario\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dataRegistro\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Store *StoreCaller) Data(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Store *StoreSession) Data() (string, error) {
	return _Store.Contract.Data(&_Store.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Store *StoreCallerSession) Data() (string, error) {
	return _Store.Contract.Data(&_Store.CallOpts)
}

// DataRegistro is a free data retrieval call binding the contract method 0xdca8d6e3.
//
// Solidity: function dataRegistro() constant returns(string)
func (_Store *StoreCaller) DataRegistro(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "dataRegistro")
	return *ret0, err
}

// DataRegistro is a free data retrieval call binding the contract method 0xdca8d6e3.
//
// Solidity: function dataRegistro() constant returns(string)
func (_Store *StoreSession) DataRegistro() (string, error) {
	return _Store.Contract.DataRegistro(&_Store.CallOpts)
}

// DataRegistro is a free data retrieval call binding the contract method 0xdca8d6e3.
//
// Solidity: function dataRegistro() constant returns(string)
func (_Store *StoreCallerSession) DataRegistro() (string, error) {
	return _Store.Contract.DataRegistro(&_Store.CallOpts)
}

// Drt is a free data retrieval call binding the contract method 0x8260b50c.
//
// Solidity: function drt() constant returns(string)
func (_Store *StoreCaller) Drt(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "drt")
	return *ret0, err
}

// Drt is a free data retrieval call binding the contract method 0x8260b50c.
//
// Solidity: function drt() constant returns(string)
func (_Store *StoreSession) Drt() (string, error) {
	return _Store.Contract.Drt(&_Store.CallOpts)
}

// Drt is a free data retrieval call binding the contract method 0x8260b50c.
//
// Solidity: function drt() constant returns(string)
func (_Store *StoreCallerSession) Drt() (string, error) {
	return _Store.Contract.Drt(&_Store.CallOpts)
}

// Estado is a free data retrieval call binding the contract method 0x688e3467.
//
// Solidity: function estado() constant returns(string)
func (_Store *StoreCaller) Estado(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "estado")
	return *ret0, err
}

// Estado is a free data retrieval call binding the contract method 0x688e3467.
//
// Solidity: function estado() constant returns(string)
func (_Store *StoreSession) Estado() (string, error) {
	return _Store.Contract.Estado(&_Store.CallOpts)
}

// Estado is a free data retrieval call binding the contract method 0x688e3467.
//
// Solidity: function estado() constant returns(string)
func (_Store *StoreCallerSession) Estado() (string, error) {
	return _Store.Contract.Estado(&_Store.CallOpts)
}

// Filiacao is a free data retrieval call binding the contract method 0x7e44bf1b.
//
// Solidity: function filiacao() constant returns(string)
func (_Store *StoreCaller) Filiacao(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "filiacao")
	return *ret0, err
}

// Filiacao is a free data retrieval call binding the contract method 0x7e44bf1b.
//
// Solidity: function filiacao() constant returns(string)
func (_Store *StoreSession) Filiacao() (string, error) {
	return _Store.Contract.Filiacao(&_Store.CallOpts)
}

// Filiacao is a free data retrieval call binding the contract method 0x7e44bf1b.
//
// Solidity: function filiacao() constant returns(string)
func (_Store *StoreCallerSession) Filiacao() (string, error) {
	return _Store.Contract.Filiacao(&_Store.CallOpts)
}

// Fls is a free data retrieval call binding the contract method 0x34431f3d.
//
// Solidity: function fls() constant returns(string)
func (_Store *StoreCaller) Fls(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "fls")
	return *ret0, err
}

// Fls is a free data retrieval call binding the contract method 0x34431f3d.
//
// Solidity: function fls() constant returns(string)
func (_Store *StoreSession) Fls() (string, error) {
	return _Store.Contract.Fls(&_Store.CallOpts)
}

// Fls is a free data retrieval call binding the contract method 0x34431f3d.
//
// Solidity: function fls() constant returns(string)
func (_Store *StoreCallerSession) Fls() (string, error) {
	return _Store.Contract.Fls(&_Store.CallOpts)
}

// Livro is a free data retrieval call binding the contract method 0x0f43de7c.
//
// Solidity: function livro() constant returns(string)
func (_Store *StoreCaller) Livro(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "livro")
	return *ret0, err
}

// Livro is a free data retrieval call binding the contract method 0x0f43de7c.
//
// Solidity: function livro() constant returns(string)
func (_Store *StoreSession) Livro() (string, error) {
	return _Store.Contract.Livro(&_Store.CallOpts)
}

// Livro is a free data retrieval call binding the contract method 0x0f43de7c.
//
// Solidity: function livro() constant returns(string)
func (_Store *StoreCallerSession) Livro() (string, error) {
	return _Store.Contract.Livro(&_Store.CallOpts)
}

// Nascimento is a free data retrieval call binding the contract method 0x59dbe458.
//
// Solidity: function nascimento() constant returns(string)
func (_Store *StoreCaller) Nascimento(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "nascimento")
	return *ret0, err
}

// Nascimento is a free data retrieval call binding the contract method 0x59dbe458.
//
// Solidity: function nascimento() constant returns(string)
func (_Store *StoreSession) Nascimento() (string, error) {
	return _Store.Contract.Nascimento(&_Store.CallOpts)
}

// Nascimento is a free data retrieval call binding the contract method 0x59dbe458.
//
// Solidity: function nascimento() constant returns(string)
func (_Store *StoreCallerSession) Nascimento() (string, error) {
	return _Store.Contract.Nascimento(&_Store.CallOpts)
}

// Nome is a free data retrieval call binding the contract method 0x2deb124b.
//
// Solidity: function nome() constant returns(string)
func (_Store *StoreCaller) Nome(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "nome")
	return *ret0, err
}

// Nome is a free data retrieval call binding the contract method 0x2deb124b.
//
// Solidity: function nome() constant returns(string)
func (_Store *StoreSession) Nome() (string, error) {
	return _Store.Contract.Nome(&_Store.CallOpts)
}

// Nome is a free data retrieval call binding the contract method 0x2deb124b.
//
// Solidity: function nome() constant returns(string)
func (_Store *StoreCallerSession) Nome() (string, error) {
	return _Store.Contract.Nome(&_Store.CallOpts)
}

// Numero is a free data retrieval call binding the contract method 0x91f5de38.
//
// Solidity: function numero() constant returns(address)
func (_Store *StoreCaller) Numero(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "numero")
	return *ret0, err
}

// Numero is a free data retrieval call binding the contract method 0x91f5de38.
//
// Solidity: function numero() constant returns(address)
func (_Store *StoreSession) Numero() (common.Address, error) {
	return _Store.Contract.Numero(&_Store.CallOpts)
}

// Numero is a free data retrieval call binding the contract method 0x91f5de38.
//
// Solidity: function numero() constant returns(address)
func (_Store *StoreCallerSession) Numero() (common.Address, error) {
	return _Store.Contract.Numero(&_Store.CallOpts)
}

// NumeroRegistro is a free data retrieval call binding the contract method 0xb8742495.
//
// Solidity: function numeroRegistro() constant returns(string)
func (_Store *StoreCaller) NumeroRegistro(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "numeroRegistro")
	return *ret0, err
}

// NumeroRegistro is a free data retrieval call binding the contract method 0xb8742495.
//
// Solidity: function numeroRegistro() constant returns(string)
func (_Store *StoreSession) NumeroRegistro() (string, error) {
	return _Store.Contract.NumeroRegistro(&_Store.CallOpts)
}

// NumeroRegistro is a free data retrieval call binding the contract method 0xb8742495.
//
// Solidity: function numeroRegistro() constant returns(string)
func (_Store *StoreCallerSession) NumeroRegistro() (string, error) {
	return _Store.Contract.NumeroRegistro(&_Store.CallOpts)
}

// Profissao is a free data retrieval call binding the contract method 0x5cbfa379.
//
// Solidity: function profissao() constant returns(string)
func (_Store *StoreCaller) Profissao(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "profissao")
	return *ret0, err
}

// Profissao is a free data retrieval call binding the contract method 0x5cbfa379.
//
// Solidity: function profissao() constant returns(string)
func (_Store *StoreSession) Profissao() (string, error) {
	return _Store.Contract.Profissao(&_Store.CallOpts)
}

// Profissao is a free data retrieval call binding the contract method 0x5cbfa379.
//
// Solidity: function profissao() constant returns(string)
func (_Store *StoreCallerSession) Profissao() (string, error) {
	return _Store.Contract.Profissao(&_Store.CallOpts)
}

// Serie is a free data retrieval call binding the contract method 0x916d6e00.
//
// Solidity: function serie() constant returns(string)
func (_Store *StoreCaller) Serie(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Store.contract.Call(opts, out, "serie")
	return *ret0, err
}

// Serie is a free data retrieval call binding the contract method 0x916d6e00.
//
// Solidity: function serie() constant returns(string)
func (_Store *StoreSession) Serie() (string, error) {
	return _Store.Contract.Serie(&_Store.CallOpts)
}

// Serie is a free data retrieval call binding the contract method 0x916d6e00.
//
// Solidity: function serie() constant returns(string)
func (_Store *StoreCallerSession) Serie() (string, error) {
	return _Store.Contract.Serie(&_Store.CallOpts)
}

// SetFormulario is a paid mutator transaction binding the contract method 0xbd55ed39.
//
// Solidity: function setFormulario(newSerie string, newNome string, newNasc string, newEst string, newData string, newFil string, newRG string) returns()
func (_Store *StoreTransactor) SetFormulario(opts *bind.TransactOpts, newSerie string, newNome string, newNasc string, newEst string, newData string, newFil string, newRG string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setFormulario", newSerie, newNome, newNasc, newEst, newData, newFil, newRG)
}

// SetFormulario is a paid mutator transaction binding the contract method 0xbd55ed39.
//
// Solidity: function setFormulario(newSerie string, newNome string, newNasc string, newEst string, newData string, newFil string, newRG string) returns()
func (_Store *StoreSession) SetFormulario(newSerie string, newNome string, newNasc string, newEst string, newData string, newFil string, newRG string) (*types.Transaction, error) {
	return _Store.Contract.SetFormulario(&_Store.TransactOpts, newSerie, newNome, newNasc, newEst, newData, newFil, newRG)
}

// SetFormulario is a paid mutator transaction binding the contract method 0xbd55ed39.
//
// Solidity: function setFormulario(newSerie string, newNome string, newNasc string, newEst string, newData string, newFil string, newRG string) returns()
func (_Store *StoreTransactorSession) SetFormulario(newSerie string, newNome string, newNasc string, newEst string, newData string, newFil string, newRG string) (*types.Transaction, error) {
	return _Store.Contract.SetFormulario(&_Store.TransactOpts, newSerie, newNome, newNasc, newEst, newData, newFil, newRG)
}

// SetProfissao is a paid mutator transaction binding the contract method 0xa01f558c.
//
// Solidity: function setProfissao(newDataRegistro string, newProf string, newNumeroReg string, newLivro string, newFls string, newDRT string) returns()
func (_Store *StoreTransactor) SetProfissao(opts *bind.TransactOpts, newDataRegistro string, newProf string, newNumeroReg string, newLivro string, newFls string, newDRT string) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setProfissao", newDataRegistro, newProf, newNumeroReg, newLivro, newFls, newDRT)
}

// SetProfissao is a paid mutator transaction binding the contract method 0xa01f558c.
//
// Solidity: function setProfissao(newDataRegistro string, newProf string, newNumeroReg string, newLivro string, newFls string, newDRT string) returns()
func (_Store *StoreSession) SetProfissao(newDataRegistro string, newProf string, newNumeroReg string, newLivro string, newFls string, newDRT string) (*types.Transaction, error) {
	return _Store.Contract.SetProfissao(&_Store.TransactOpts, newDataRegistro, newProf, newNumeroReg, newLivro, newFls, newDRT)
}

// SetProfissao is a paid mutator transaction binding the contract method 0xa01f558c.
//
// Solidity: function setProfissao(newDataRegistro string, newProf string, newNumeroReg string, newLivro string, newFls string, newDRT string) returns()
func (_Store *StoreTransactorSession) SetProfissao(newDataRegistro string, newProf string, newNumeroReg string, newLivro string, newFls string, newDRT string) (*types.Transaction, error) {
	return _Store.Contract.SetProfissao(&_Store.TransactOpts, newDataRegistro, newProf, newNumeroReg, newLivro, newFls, newDRT)
}
