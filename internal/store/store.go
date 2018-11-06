// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package store

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

// CtpsABI is the input ABI used to generate the binding from.
const CtpsABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"livro\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nome\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"fls\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nascimento\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"profissao\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"estado\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"filiacao\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"drt\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"serie\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numero\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newDataRegistro\",\"type\":\"string\"},{\"name\":\"newProf\",\"type\":\"string\"},{\"name\":\"newNumeroReg\",\"type\":\"string\"},{\"name\":\"newLivro\",\"type\":\"string\"},{\"name\":\"newFls\",\"type\":\"string\"},{\"name\":\"newDRT\",\"type\":\"string\"}],\"name\":\"setProfissao\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"numeroRegistro\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newSerie\",\"type\":\"string\"},{\"name\":\"newNome\",\"type\":\"string\"},{\"name\":\"newNasc\",\"type\":\"string\"},{\"name\":\"newEst\",\"type\":\"string\"},{\"name\":\"newData\",\"type\":\"string\"},{\"name\":\"newFil\",\"type\":\"string\"},{\"name\":\"newRG\",\"type\":\"string\"}],\"name\":\"setFormulario\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"dataRegistro\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// CtpsBin is the compiled bytecode used for deploying new contracts.
const CtpsBin = `608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550611565806100606000396000f3006080604052600436106100db576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630f43de7c146100e05780632deb124b1461017057806334431f3d1461020057806359dbe458146102905780635cbfa37914610320578063688e3467146103b057806373d4a13a146104405780637e44bf1b146104d05780638260b50c14610560578063916d6e00146105f057806391f5de3814610680578063a01f558c146106d7578063b87424951461089e578063bd55ed391461092e578063dca8d6e314610b3b575b600080fd5b3480156100ec57600080fd5b506100f5610bcb565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561013557808201518184015260208101905061011a565b50505050905090810190601f1680156101625780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561017c57600080fd5b50610185610c69565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101c55780820151818401526020810190506101aa565b50505050905090810190601f1680156101f25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561020c57600080fd5b50610215610d07565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561025557808201518184015260208101905061023a565b50505050905090810190601f1680156102825780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561029c57600080fd5b506102a5610da5565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102e55780820151818401526020810190506102ca565b50505050905090810190601f1680156103125780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561032c57600080fd5b50610335610e43565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561037557808201518184015260208101905061035a565b50505050905090810190601f1680156103a25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156103bc57600080fd5b506103c5610ee1565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104055780820151818401526020810190506103ea565b50505050905090810190601f1680156104325780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561044c57600080fd5b50610455610f7f565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561049557808201518184015260208101905061047a565b50505050905090810190601f1680156104c25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156104dc57600080fd5b506104e561101d565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561052557808201518184015260208101905061050a565b50505050905090810190601f1680156105525780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561056c57600080fd5b506105756110bb565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105b557808201518184015260208101905061059a565b50505050905090810190601f1680156105e25780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156105fc57600080fd5b50610605611159565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561064557808201518184015260208101905061062a565b50505050905090810190601f1680156106725780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561068c57600080fd5b506106956111f7565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156106e357600080fd5b5061089c600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061121c565b005b3480156108aa57600080fd5b506108b36112ae565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156108f35780820151818401526020810190506108d8565b50505050905090810190601f1680156109205780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561093a57600080fd5b50610b39600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061134c565b005b348015610b4757600080fd5b50610b506113f6565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610b90578082015181840152602081019050610b75565b50505050905090810190601f168015610bbd5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b600b8054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c615780601f10610c3657610100808354040283529160200191610c61565b820191906000526020600020905b815481529060010190602001808311610c4457829003601f168201915b505050505081565b60028054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cff5780601f10610cd457610100808354040283529160200191610cff565b820191906000526020600020905b815481529060010190602001808311610ce257829003601f168201915b505050505081565b600c8054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d9d5780601f10610d7257610100808354040283529160200191610d9d565b820191906000526020600020905b815481529060010190602001808311610d8057829003601f168201915b505050505081565b60038054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e3b5780601f10610e1057610100808354040283529160200191610e3b565b820191906000526020600020905b815481529060010190602001808311610e1e57829003601f168201915b505050505081565b60098054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ed95780601f10610eae57610100808354040283529160200191610ed9565b820191906000526020600020905b815481529060010190602001808311610ebc57829003601f168201915b505050505081565b60048054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f775780601f10610f4c57610100808354040283529160200191610f77565b820191906000526020600020905b815481529060010190602001808311610f5a57829003601f168201915b505050505081565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110155780601f10610fea57610100808354040283529160200191611015565b820191906000526020600020905b815481529060010190602001808311610ff857829003601f168201915b505050505081565b60068054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110b35780601f10611088576101008083540402835291602001916110b3565b820191906000526020600020905b81548152906001019060200180831161109657829003601f168201915b505050505081565b600d8054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111515780601f1061112657610100808354040283529160200191611151565b820191906000526020600020905b81548152906001019060200180831161113457829003601f168201915b505050505081565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111ef5780601f106111c4576101008083540402835291602001916111ef565b820191906000526020600020905b8154815290600101906020018083116111d257829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8560089080519060200190611232929190611494565b508460099080519060200190611249929190611494565b5083600a9080519060200190611260929190611494565b5081600c9080519060200190611277929190611494565b5080600d908051906020019061128e929190611494565b5082600b90805190602001906112a5929190611494565b50505050505050565b600a8054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113445780601f1061131957610100808354040283529160200191611344565b820191906000526020600020905b81548152906001019060200180831161132757829003601f168201915b505050505081565b8660019080519060200190611362929190611494565b508560029080519060200190611379929190611494565b508460039080519060200190611390929190611494565b5083600490805190602001906113a7929190611494565b5082600590805190602001906113be929190611494565b5081600690805190602001906113d5929190611494565b5080600790805190602001906113ec929190611494565b5050505050505050565b60088054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561148c5780601f106114615761010080835404028352916020019161148c565b820191906000526020600020905b81548152906001019060200180831161146f57829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106114d557805160ff1916838001178555611503565b82800160010185558215611503579182015b828111156115025782518255916020019190600101906114e7565b5b5090506115109190611514565b5090565b61153691905b8082111561153257600081600090555060010161151a565b5090565b905600a165627a7a72305820ba39f2225583a365c04584167e400457d3baeef9ef59c13e2f67ee4d508daf020029`

// DeployCtps deploys a new Ethereum contract, binding an instance of Ctps to it.
func DeployCtps(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ctps, error) {
	parsed, err := abi.JSON(strings.NewReader(CtpsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CtpsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ctps{CtpsCaller: CtpsCaller{contract: contract}, CtpsTransactor: CtpsTransactor{contract: contract}, CtpsFilterer: CtpsFilterer{contract: contract}}, nil
}

// Ctps is an auto generated Go binding around an Ethereum contract.
type Ctps struct {
	CtpsCaller     // Read-only binding to the contract
	CtpsTransactor // Write-only binding to the contract
	CtpsFilterer   // Log filterer for contract events
}

// CtpsCaller is an auto generated read-only Go binding around an Ethereum contract.
type CtpsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CtpsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CtpsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CtpsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CtpsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CtpsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CtpsSession struct {
	Contract     *Ctps             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CtpsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CtpsCallerSession struct {
	Contract *CtpsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CtpsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CtpsTransactorSession struct {
	Contract     *CtpsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CtpsRaw is an auto generated low-level Go binding around an Ethereum contract.
type CtpsRaw struct {
	Contract *Ctps // Generic contract binding to access the raw methods on
}

// CtpsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CtpsCallerRaw struct {
	Contract *CtpsCaller // Generic read-only contract binding to access the raw methods on
}

// CtpsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CtpsTransactorRaw struct {
	Contract *CtpsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCtps creates a new instance of Ctps, bound to a specific deployed contract.
func NewCtps(address common.Address, backend bind.ContractBackend) (*Ctps, error) {
	contract, err := bindCtps(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ctps{CtpsCaller: CtpsCaller{contract: contract}, CtpsTransactor: CtpsTransactor{contract: contract}, CtpsFilterer: CtpsFilterer{contract: contract}}, nil
}

// NewCtpsCaller creates a new read-only instance of Ctps, bound to a specific deployed contract.
func NewCtpsCaller(address common.Address, caller bind.ContractCaller) (*CtpsCaller, error) {
	contract, err := bindCtps(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CtpsCaller{contract: contract}, nil
}

// NewCtpsTransactor creates a new write-only instance of Ctps, bound to a specific deployed contract.
func NewCtpsTransactor(address common.Address, transactor bind.ContractTransactor) (*CtpsTransactor, error) {
	contract, err := bindCtps(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CtpsTransactor{contract: contract}, nil
}

// NewCtpsFilterer creates a new log filterer instance of Ctps, bound to a specific deployed contract.
func NewCtpsFilterer(address common.Address, filterer bind.ContractFilterer) (*CtpsFilterer, error) {
	contract, err := bindCtps(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CtpsFilterer{contract: contract}, nil
}

// bindCtps binds a generic wrapper to an already deployed contract.
func bindCtps(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CtpsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ctps *CtpsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ctps.Contract.CtpsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ctps *CtpsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ctps.Contract.CtpsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ctps *CtpsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ctps.Contract.CtpsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ctps *CtpsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ctps.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ctps *CtpsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ctps.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ctps *CtpsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ctps.Contract.contract.Transact(opts, method, params...)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Ctps *CtpsCaller) Data(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Ctps *CtpsSession) Data() (string, error) {
	return _Ctps.Contract.Data(&_Ctps.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(string)
func (_Ctps *CtpsCallerSession) Data() (string, error) {
	return _Ctps.Contract.Data(&_Ctps.CallOpts)
}

// DataRegistro is a free data retrieval call binding the contract method 0xdca8d6e3.
//
// Solidity: function dataRegistro() constant returns(string)
func (_Ctps *CtpsCaller) DataRegistro(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "dataRegistro")
	return *ret0, err
}

// DataRegistro is a free data retrieval call binding the contract method 0xdca8d6e3.
//
// Solidity: function dataRegistro() constant returns(string)
func (_Ctps *CtpsSession) DataRegistro() (string, error) {
	return _Ctps.Contract.DataRegistro(&_Ctps.CallOpts)
}

// DataRegistro is a free data retrieval call binding the contract method 0xdca8d6e3.
//
// Solidity: function dataRegistro() constant returns(string)
func (_Ctps *CtpsCallerSession) DataRegistro() (string, error) {
	return _Ctps.Contract.DataRegistro(&_Ctps.CallOpts)
}

// Drt is a free data retrieval call binding the contract method 0x8260b50c.
//
// Solidity: function drt() constant returns(string)
func (_Ctps *CtpsCaller) Drt(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "drt")
	return *ret0, err
}

// Drt is a free data retrieval call binding the contract method 0x8260b50c.
//
// Solidity: function drt() constant returns(string)
func (_Ctps *CtpsSession) Drt() (string, error) {
	return _Ctps.Contract.Drt(&_Ctps.CallOpts)
}

// Drt is a free data retrieval call binding the contract method 0x8260b50c.
//
// Solidity: function drt() constant returns(string)
func (_Ctps *CtpsCallerSession) Drt() (string, error) {
	return _Ctps.Contract.Drt(&_Ctps.CallOpts)
}

// Estado is a free data retrieval call binding the contract method 0x688e3467.
//
// Solidity: function estado() constant returns(string)
func (_Ctps *CtpsCaller) Estado(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "estado")
	return *ret0, err
}

// Estado is a free data retrieval call binding the contract method 0x688e3467.
//
// Solidity: function estado() constant returns(string)
func (_Ctps *CtpsSession) Estado() (string, error) {
	return _Ctps.Contract.Estado(&_Ctps.CallOpts)
}

// Estado is a free data retrieval call binding the contract method 0x688e3467.
//
// Solidity: function estado() constant returns(string)
func (_Ctps *CtpsCallerSession) Estado() (string, error) {
	return _Ctps.Contract.Estado(&_Ctps.CallOpts)
}

// Filiacao is a free data retrieval call binding the contract method 0x7e44bf1b.
//
// Solidity: function filiacao() constant returns(string)
func (_Ctps *CtpsCaller) Filiacao(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "filiacao")
	return *ret0, err
}

// Filiacao is a free data retrieval call binding the contract method 0x7e44bf1b.
//
// Solidity: function filiacao() constant returns(string)
func (_Ctps *CtpsSession) Filiacao() (string, error) {
	return _Ctps.Contract.Filiacao(&_Ctps.CallOpts)
}

// Filiacao is a free data retrieval call binding the contract method 0x7e44bf1b.
//
// Solidity: function filiacao() constant returns(string)
func (_Ctps *CtpsCallerSession) Filiacao() (string, error) {
	return _Ctps.Contract.Filiacao(&_Ctps.CallOpts)
}

// Fls is a free data retrieval call binding the contract method 0x34431f3d.
//
// Solidity: function fls() constant returns(string)
func (_Ctps *CtpsCaller) Fls(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "fls")
	return *ret0, err
}

// Fls is a free data retrieval call binding the contract method 0x34431f3d.
//
// Solidity: function fls() constant returns(string)
func (_Ctps *CtpsSession) Fls() (string, error) {
	return _Ctps.Contract.Fls(&_Ctps.CallOpts)
}

// Fls is a free data retrieval call binding the contract method 0x34431f3d.
//
// Solidity: function fls() constant returns(string)
func (_Ctps *CtpsCallerSession) Fls() (string, error) {
	return _Ctps.Contract.Fls(&_Ctps.CallOpts)
}

// Livro is a free data retrieval call binding the contract method 0x0f43de7c.
//
// Solidity: function livro() constant returns(string)
func (_Ctps *CtpsCaller) Livro(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "livro")
	return *ret0, err
}

// Livro is a free data retrieval call binding the contract method 0x0f43de7c.
//
// Solidity: function livro() constant returns(string)
func (_Ctps *CtpsSession) Livro() (string, error) {
	return _Ctps.Contract.Livro(&_Ctps.CallOpts)
}

// Livro is a free data retrieval call binding the contract method 0x0f43de7c.
//
// Solidity: function livro() constant returns(string)
func (_Ctps *CtpsCallerSession) Livro() (string, error) {
	return _Ctps.Contract.Livro(&_Ctps.CallOpts)
}

// Nascimento is a free data retrieval call binding the contract method 0x59dbe458.
//
// Solidity: function nascimento() constant returns(string)
func (_Ctps *CtpsCaller) Nascimento(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "nascimento")
	return *ret0, err
}

// Nascimento is a free data retrieval call binding the contract method 0x59dbe458.
//
// Solidity: function nascimento() constant returns(string)
func (_Ctps *CtpsSession) Nascimento() (string, error) {
	return _Ctps.Contract.Nascimento(&_Ctps.CallOpts)
}

// Nascimento is a free data retrieval call binding the contract method 0x59dbe458.
//
// Solidity: function nascimento() constant returns(string)
func (_Ctps *CtpsCallerSession) Nascimento() (string, error) {
	return _Ctps.Contract.Nascimento(&_Ctps.CallOpts)
}

// Nome is a free data retrieval call binding the contract method 0x2deb124b.
//
// Solidity: function nome() constant returns(string)
func (_Ctps *CtpsCaller) Nome(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "nome")
	return *ret0, err
}

// Nome is a free data retrieval call binding the contract method 0x2deb124b.
//
// Solidity: function nome() constant returns(string)
func (_Ctps *CtpsSession) Nome() (string, error) {
	return _Ctps.Contract.Nome(&_Ctps.CallOpts)
}

// Nome is a free data retrieval call binding the contract method 0x2deb124b.
//
// Solidity: function nome() constant returns(string)
func (_Ctps *CtpsCallerSession) Nome() (string, error) {
	return _Ctps.Contract.Nome(&_Ctps.CallOpts)
}

// Numero is a free data retrieval call binding the contract method 0x91f5de38.
//
// Solidity: function numero() constant returns(address)
func (_Ctps *CtpsCaller) Numero(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "numero")
	return *ret0, err
}

// Numero is a free data retrieval call binding the contract method 0x91f5de38.
//
// Solidity: function numero() constant returns(address)
func (_Ctps *CtpsSession) Numero() (common.Address, error) {
	return _Ctps.Contract.Numero(&_Ctps.CallOpts)
}

// Numero is a free data retrieval call binding the contract method 0x91f5de38.
//
// Solidity: function numero() constant returns(address)
func (_Ctps *CtpsCallerSession) Numero() (common.Address, error) {
	return _Ctps.Contract.Numero(&_Ctps.CallOpts)
}

// NumeroRegistro is a free data retrieval call binding the contract method 0xb8742495.
//
// Solidity: function numeroRegistro() constant returns(string)
func (_Ctps *CtpsCaller) NumeroRegistro(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "numeroRegistro")
	return *ret0, err
}

// NumeroRegistro is a free data retrieval call binding the contract method 0xb8742495.
//
// Solidity: function numeroRegistro() constant returns(string)
func (_Ctps *CtpsSession) NumeroRegistro() (string, error) {
	return _Ctps.Contract.NumeroRegistro(&_Ctps.CallOpts)
}

// NumeroRegistro is a free data retrieval call binding the contract method 0xb8742495.
//
// Solidity: function numeroRegistro() constant returns(string)
func (_Ctps *CtpsCallerSession) NumeroRegistro() (string, error) {
	return _Ctps.Contract.NumeroRegistro(&_Ctps.CallOpts)
}

// Profissao is a free data retrieval call binding the contract method 0x5cbfa379.
//
// Solidity: function profissao() constant returns(string)
func (_Ctps *CtpsCaller) Profissao(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "profissao")
	return *ret0, err
}

// Profissao is a free data retrieval call binding the contract method 0x5cbfa379.
//
// Solidity: function profissao() constant returns(string)
func (_Ctps *CtpsSession) Profissao() (string, error) {
	return _Ctps.Contract.Profissao(&_Ctps.CallOpts)
}

// Profissao is a free data retrieval call binding the contract method 0x5cbfa379.
//
// Solidity: function profissao() constant returns(string)
func (_Ctps *CtpsCallerSession) Profissao() (string, error) {
	return _Ctps.Contract.Profissao(&_Ctps.CallOpts)
}

// Serie is a free data retrieval call binding the contract method 0x916d6e00.
//
// Solidity: function serie() constant returns(string)
func (_Ctps *CtpsCaller) Serie(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Ctps.contract.Call(opts, out, "serie")
	return *ret0, err
}

// Serie is a free data retrieval call binding the contract method 0x916d6e00.
//
// Solidity: function serie() constant returns(string)
func (_Ctps *CtpsSession) Serie() (string, error) {
	return _Ctps.Contract.Serie(&_Ctps.CallOpts)
}

// Serie is a free data retrieval call binding the contract method 0x916d6e00.
//
// Solidity: function serie() constant returns(string)
func (_Ctps *CtpsCallerSession) Serie() (string, error) {
	return _Ctps.Contract.Serie(&_Ctps.CallOpts)
}

// SetFormulario is a paid mutator transaction binding the contract method 0xbd55ed39.
//
// Solidity: function setFormulario(newSerie string, newNome string, newNasc string, newEst string, newData string, newFil string, newRG string) returns()
func (_Ctps *CtpsTransactor) SetFormulario(opts *bind.TransactOpts, newSerie string, newNome string, newNasc string, newEst string, newData string, newFil string, newRG string) (*types.Transaction, error) {
	return _Ctps.contract.Transact(opts, "setFormulario", newSerie, newNome, newNasc, newEst, newData, newFil, newRG)
}

// SetFormulario is a paid mutator transaction binding the contract method 0xbd55ed39.
//
// Solidity: function setFormulario(newSerie string, newNome string, newNasc string, newEst string, newData string, newFil string, newRG string) returns()
func (_Ctps *CtpsSession) SetFormulario(newSerie string, newNome string, newNasc string, newEst string, newData string, newFil string, newRG string) (*types.Transaction, error) {
	return _Ctps.Contract.SetFormulario(&_Ctps.TransactOpts, newSerie, newNome, newNasc, newEst, newData, newFil, newRG)
}

// SetFormulario is a paid mutator transaction binding the contract method 0xbd55ed39.
//
// Solidity: function setFormulario(newSerie string, newNome string, newNasc string, newEst string, newData string, newFil string, newRG string) returns()
func (_Ctps *CtpsTransactorSession) SetFormulario(newSerie string, newNome string, newNasc string, newEst string, newData string, newFil string, newRG string) (*types.Transaction, error) {
	return _Ctps.Contract.SetFormulario(&_Ctps.TransactOpts, newSerie, newNome, newNasc, newEst, newData, newFil, newRG)
}

// SetProfissao is a paid mutator transaction binding the contract method 0xa01f558c.
//
// Solidity: function setProfissao(newDataRegistro string, newProf string, newNumeroReg string, newLivro string, newFls string, newDRT string) returns()
func (_Ctps *CtpsTransactor) SetProfissao(opts *bind.TransactOpts, newDataRegistro string, newProf string, newNumeroReg string, newLivro string, newFls string, newDRT string) (*types.Transaction, error) {
	return _Ctps.contract.Transact(opts, "setProfissao", newDataRegistro, newProf, newNumeroReg, newLivro, newFls, newDRT)
}

// SetProfissao is a paid mutator transaction binding the contract method 0xa01f558c.
//
// Solidity: function setProfissao(newDataRegistro string, newProf string, newNumeroReg string, newLivro string, newFls string, newDRT string) returns()
func (_Ctps *CtpsSession) SetProfissao(newDataRegistro string, newProf string, newNumeroReg string, newLivro string, newFls string, newDRT string) (*types.Transaction, error) {
	return _Ctps.Contract.SetProfissao(&_Ctps.TransactOpts, newDataRegistro, newProf, newNumeroReg, newLivro, newFls, newDRT)
}

// SetProfissao is a paid mutator transaction binding the contract method 0xa01f558c.
//
// Solidity: function setProfissao(newDataRegistro string, newProf string, newNumeroReg string, newLivro string, newFls string, newDRT string) returns()
func (_Ctps *CtpsTransactorSession) SetProfissao(newDataRegistro string, newProf string, newNumeroReg string, newLivro string, newFls string, newDRT string) (*types.Transaction, error) {
	return _Ctps.Contract.SetProfissao(&_Ctps.TransactOpts, newDataRegistro, newProf, newNumeroReg, newLivro, newFls, newDRT)
}
