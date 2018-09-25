// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bike

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

// BikeABI is the input ABI used to generate the binding from.
const BikeABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rentalExpiryTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"renter\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rentalFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"returnBike\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_extraData\",\"type\":\"bytes\"}],\"name\":\"receiveApproval\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isRented\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"payee\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rentalTimeInMinutes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newRenter\",\"type\":\"address\"}],\"name\":\"transferBike\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rentalFee\",\"type\":\"uint256\"},{\"name\":\"_rentalTimeInMinutes\",\"type\":\"uint256\"},{\"name\":\"_tokenContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"renter\",\"type\":\"address\"}],\"name\":\"Rented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"renter\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"rentalExpiryTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"currentTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Returned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// BikeBin is the compiled bytecode used for deploying new contracts.
const BikeBin = `608060405234801561001057600080fd5b50604051606080610f77833981018060405281019080805190602001909291908051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550826001819055508160028190555080600360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050610e8f806100e86000396000f3006080604052600436106100c5576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063271ee8ad146100ca5780632e88ab0b146100f55780634876956b1461014c57806355a373d61461017757806360a6673c146101ce578063715018a6146101e55780638da5cb5b146101fc5780638f4ffcb114610253578063a630495214610306578063ae90b21314610335578063d8f3ef601461038c578063e16d1fbd146103b7578063f2fde38b146103fa575b600080fd5b3480156100d657600080fd5b506100df61043d565b6040518082815260200191505060405180910390f35b34801561010157600080fd5b5061010a610443565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561015857600080fd5b50610161610469565b6040518082815260200191505060405180910390f35b34801561018357600080fd5b5061018c61046f565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101da57600080fd5b506101e3610495565b005b3480156101f157600080fd5b506101fa610806565b005b34801561020857600080fd5b50610211610908565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561025f57600080fd5b50610304600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061092d565b005b34801561031257600080fd5b5061031b610a72565b604051808215151515815260200191505060405180910390f35b34801561034157600080fd5b5061034a610a85565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561039857600080fd5b506103a1610aab565b6040518082815260200191505060405180910390f35b3480156103c357600080fd5b506103f8600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610ab1565b005b34801561040657600080fd5b5061043b600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bd5565b005b60065481565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60015481565b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600080600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561055d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f53656e646572206973206e6f74207468652063757272656e742072656e74657281525060200191505060405180910390fd5b600560149054906101000a900460ff16151561057857600080fd5b600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691506002600154029050600654421115156106bc578173ffffffffffffffffffffffffffffffffffffffff1663a9059cbb600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561067557600080fd5b505af1158015610689573d6000803e3d6000fd5b505050506040513d602081101561069f57600080fd5b810190808051906020019092919050505015156106bb57600080fd5b5b7f089560b26c90954c01dd43421beac7892c9cf783a8e9089820df3647e6ea4176600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166006544284604051808573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200184815260200183815260200182815260200194505050505060405180910390a16000600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600560146101000a81548160ff0219169083151502179055506000600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060006006819055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561086157600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008290506003600154028414151561094557600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166323b872dd8630876040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610a1c57600080fd5b505af1158015610a30573d6000803e3d6000fd5b505050506040513d6020811015610a4657600080fd5b81019080805190602001909291905050501515610a6257600080fd5b610a6b85610c3c565b5050505050565b600560149054906101000a900460ff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b76576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f53656e646572206973206e6f74207468652063757272656e742072656e74657281525060200191505060405180910390fd5b600560149054906101000a900460ff161515610b9157600080fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610c3057600080fd5b610c3981610d69565b50565b600560149054906101000a900460ff16151515610c5857600080fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600560146101000a81548160ff021916908315150217905550603c6002540242016006819055507f18ae6638be53d8a225d874e2755866e292376637d786506bf60355f3d13654b933604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610da557600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505600a165627a7a7230582097a812d38728d36e8188a8fb53495237265af602620f85f45f2cf060abbd9cf90029`

// DeployBike deploys a new Ethereum contract, binding an instance of Bike to it.
func DeployBike(auth *bind.TransactOpts, backend bind.ContractBackend, _rentalFee *big.Int, _rentalTimeInMinutes *big.Int, _tokenContract common.Address) (common.Address, *types.Transaction, *Bike, error) {
	parsed, err := abi.JSON(strings.NewReader(BikeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BikeBin), backend, _rentalFee, _rentalTimeInMinutes, _tokenContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bike{BikeCaller: BikeCaller{contract: contract}, BikeTransactor: BikeTransactor{contract: contract}, BikeFilterer: BikeFilterer{contract: contract}}, nil
}

// Bike is an auto generated Go binding around an Ethereum contract.
type Bike struct {
	BikeCaller     // Read-only binding to the contract
	BikeTransactor // Write-only binding to the contract
	BikeFilterer   // Log filterer for contract events
}

// BikeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BikeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BikeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BikeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BikeSession struct {
	Contract     *Bike             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BikeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BikeCallerSession struct {
	Contract *BikeCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BikeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BikeTransactorSession struct {
	Contract     *BikeTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BikeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BikeRaw struct {
	Contract *Bike // Generic contract binding to access the raw methods on
}

// BikeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BikeCallerRaw struct {
	Contract *BikeCaller // Generic read-only contract binding to access the raw methods on
}

// BikeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BikeTransactorRaw struct {
	Contract *BikeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBike creates a new instance of Bike, bound to a specific deployed contract.
func NewBike(address common.Address, backend bind.ContractBackend) (*Bike, error) {
	contract, err := bindBike(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bike{BikeCaller: BikeCaller{contract: contract}, BikeTransactor: BikeTransactor{contract: contract}, BikeFilterer: BikeFilterer{contract: contract}}, nil
}

// NewBikeCaller creates a new read-only instance of Bike, bound to a specific deployed contract.
func NewBikeCaller(address common.Address, caller bind.ContractCaller) (*BikeCaller, error) {
	contract, err := bindBike(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BikeCaller{contract: contract}, nil
}

// NewBikeTransactor creates a new write-only instance of Bike, bound to a specific deployed contract.
func NewBikeTransactor(address common.Address, transactor bind.ContractTransactor) (*BikeTransactor, error) {
	contract, err := bindBike(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BikeTransactor{contract: contract}, nil
}

// NewBikeFilterer creates a new log filterer instance of Bike, bound to a specific deployed contract.
func NewBikeFilterer(address common.Address, filterer bind.ContractFilterer) (*BikeFilterer, error) {
	contract, err := bindBike(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BikeFilterer{contract: contract}, nil
}

// bindBike binds a generic wrapper to an already deployed contract.
func bindBike(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BikeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bike *BikeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bike.Contract.BikeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bike *BikeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bike.Contract.BikeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bike *BikeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bike.Contract.BikeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bike *BikeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Bike.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bike *BikeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bike.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bike *BikeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bike.Contract.contract.Transact(opts, method, params...)
}

// IsRented is a free data retrieval call binding the contract method 0xa6304952.
//
// Solidity: function isRented() constant returns(bool)
func (_Bike *BikeCaller) IsRented(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "isRented")
	return *ret0, err
}

// IsRented is a free data retrieval call binding the contract method 0xa6304952.
//
// Solidity: function isRented() constant returns(bool)
func (_Bike *BikeSession) IsRented() (bool, error) {
	return _Bike.Contract.IsRented(&_Bike.CallOpts)
}

// IsRented is a free data retrieval call binding the contract method 0xa6304952.
//
// Solidity: function isRented() constant returns(bool)
func (_Bike *BikeCallerSession) IsRented() (bool, error) {
	return _Bike.Contract.IsRented(&_Bike.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bike *BikeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bike *BikeSession) Owner() (common.Address, error) {
	return _Bike.Contract.Owner(&_Bike.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Bike *BikeCallerSession) Owner() (common.Address, error) {
	return _Bike.Contract.Owner(&_Bike.CallOpts)
}

// Payee is a free data retrieval call binding the contract method 0xae90b213.
//
// Solidity: function payee() constant returns(address)
func (_Bike *BikeCaller) Payee(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "payee")
	return *ret0, err
}

// Payee is a free data retrieval call binding the contract method 0xae90b213.
//
// Solidity: function payee() constant returns(address)
func (_Bike *BikeSession) Payee() (common.Address, error) {
	return _Bike.Contract.Payee(&_Bike.CallOpts)
}

// Payee is a free data retrieval call binding the contract method 0xae90b213.
//
// Solidity: function payee() constant returns(address)
func (_Bike *BikeCallerSession) Payee() (common.Address, error) {
	return _Bike.Contract.Payee(&_Bike.CallOpts)
}

// RentalExpiryTime is a free data retrieval call binding the contract method 0x271ee8ad.
//
// Solidity: function rentalExpiryTime() constant returns(uint256)
func (_Bike *BikeCaller) RentalExpiryTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "rentalExpiryTime")
	return *ret0, err
}

// RentalExpiryTime is a free data retrieval call binding the contract method 0x271ee8ad.
//
// Solidity: function rentalExpiryTime() constant returns(uint256)
func (_Bike *BikeSession) RentalExpiryTime() (*big.Int, error) {
	return _Bike.Contract.RentalExpiryTime(&_Bike.CallOpts)
}

// RentalExpiryTime is a free data retrieval call binding the contract method 0x271ee8ad.
//
// Solidity: function rentalExpiryTime() constant returns(uint256)
func (_Bike *BikeCallerSession) RentalExpiryTime() (*big.Int, error) {
	return _Bike.Contract.RentalExpiryTime(&_Bike.CallOpts)
}

// RentalFee is a free data retrieval call binding the contract method 0x4876956b.
//
// Solidity: function rentalFee() constant returns(uint256)
func (_Bike *BikeCaller) RentalFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "rentalFee")
	return *ret0, err
}

// RentalFee is a free data retrieval call binding the contract method 0x4876956b.
//
// Solidity: function rentalFee() constant returns(uint256)
func (_Bike *BikeSession) RentalFee() (*big.Int, error) {
	return _Bike.Contract.RentalFee(&_Bike.CallOpts)
}

// RentalFee is a free data retrieval call binding the contract method 0x4876956b.
//
// Solidity: function rentalFee() constant returns(uint256)
func (_Bike *BikeCallerSession) RentalFee() (*big.Int, error) {
	return _Bike.Contract.RentalFee(&_Bike.CallOpts)
}

// RentalTimeInMinutes is a free data retrieval call binding the contract method 0xd8f3ef60.
//
// Solidity: function rentalTimeInMinutes() constant returns(uint256)
func (_Bike *BikeCaller) RentalTimeInMinutes(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "rentalTimeInMinutes")
	return *ret0, err
}

// RentalTimeInMinutes is a free data retrieval call binding the contract method 0xd8f3ef60.
//
// Solidity: function rentalTimeInMinutes() constant returns(uint256)
func (_Bike *BikeSession) RentalTimeInMinutes() (*big.Int, error) {
	return _Bike.Contract.RentalTimeInMinutes(&_Bike.CallOpts)
}

// RentalTimeInMinutes is a free data retrieval call binding the contract method 0xd8f3ef60.
//
// Solidity: function rentalTimeInMinutes() constant returns(uint256)
func (_Bike *BikeCallerSession) RentalTimeInMinutes() (*big.Int, error) {
	return _Bike.Contract.RentalTimeInMinutes(&_Bike.CallOpts)
}

// Renter is a free data retrieval call binding the contract method 0x2e88ab0b.
//
// Solidity: function renter() constant returns(address)
func (_Bike *BikeCaller) Renter(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "renter")
	return *ret0, err
}

// Renter is a free data retrieval call binding the contract method 0x2e88ab0b.
//
// Solidity: function renter() constant returns(address)
func (_Bike *BikeSession) Renter() (common.Address, error) {
	return _Bike.Contract.Renter(&_Bike.CallOpts)
}

// Renter is a free data retrieval call binding the contract method 0x2e88ab0b.
//
// Solidity: function renter() constant returns(address)
func (_Bike *BikeCallerSession) Renter() (common.Address, error) {
	return _Bike.Contract.Renter(&_Bike.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_Bike *BikeCaller) TokenContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Bike.contract.Call(opts, out, "tokenContract")
	return *ret0, err
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_Bike *BikeSession) TokenContract() (common.Address, error) {
	return _Bike.Contract.TokenContract(&_Bike.CallOpts)
}

// TokenContract is a free data retrieval call binding the contract method 0x55a373d6.
//
// Solidity: function tokenContract() constant returns(address)
func (_Bike *BikeCallerSession) TokenContract() (common.Address, error) {
	return _Bike.Contract.TokenContract(&_Bike.CallOpts)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_Bike *BikeTransactor) ReceiveApproval(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "receiveApproval", _from, _value, _token, _extraData)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_Bike *BikeSession) ReceiveApproval(_from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Bike.Contract.ReceiveApproval(&_Bike.TransactOpts, _from, _value, _token, _extraData)
}

// ReceiveApproval is a paid mutator transaction binding the contract method 0x8f4ffcb1.
//
// Solidity: function receiveApproval(_from address, _value uint256, _token address, _extraData bytes) returns()
func (_Bike *BikeTransactorSession) ReceiveApproval(_from common.Address, _value *big.Int, _token common.Address, _extraData []byte) (*types.Transaction, error) {
	return _Bike.Contract.ReceiveApproval(&_Bike.TransactOpts, _from, _value, _token, _extraData)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bike *BikeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bike *BikeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bike.Contract.RenounceOwnership(&_Bike.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Bike *BikeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Bike.Contract.RenounceOwnership(&_Bike.TransactOpts)
}

// ReturnBike is a paid mutator transaction binding the contract method 0x60a6673c.
//
// Solidity: function returnBike() returns()
func (_Bike *BikeTransactor) ReturnBike(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "returnBike")
}

// ReturnBike is a paid mutator transaction binding the contract method 0x60a6673c.
//
// Solidity: function returnBike() returns()
func (_Bike *BikeSession) ReturnBike() (*types.Transaction, error) {
	return _Bike.Contract.ReturnBike(&_Bike.TransactOpts)
}

// ReturnBike is a paid mutator transaction binding the contract method 0x60a6673c.
//
// Solidity: function returnBike() returns()
func (_Bike *BikeTransactorSession) ReturnBike() (*types.Transaction, error) {
	return _Bike.Contract.ReturnBike(&_Bike.TransactOpts)
}

// TransferBike is a paid mutator transaction binding the contract method 0xe16d1fbd.
//
// Solidity: function transferBike(newRenter address) returns()
func (_Bike *BikeTransactor) TransferBike(opts *bind.TransactOpts, newRenter common.Address) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "transferBike", newRenter)
}

// TransferBike is a paid mutator transaction binding the contract method 0xe16d1fbd.
//
// Solidity: function transferBike(newRenter address) returns()
func (_Bike *BikeSession) TransferBike(newRenter common.Address) (*types.Transaction, error) {
	return _Bike.Contract.TransferBike(&_Bike.TransactOpts, newRenter)
}

// TransferBike is a paid mutator transaction binding the contract method 0xe16d1fbd.
//
// Solidity: function transferBike(newRenter address) returns()
func (_Bike *BikeTransactorSession) TransferBike(newRenter common.Address) (*types.Transaction, error) {
	return _Bike.Contract.TransferBike(&_Bike.TransactOpts, newRenter)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Bike *BikeTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Bike.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Bike *BikeSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Bike.Contract.TransferOwnership(&_Bike.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(_newOwner address) returns()
func (_Bike *BikeTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Bike.Contract.TransferOwnership(&_Bike.TransactOpts, _newOwner)
}

// BikeOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Bike contract.
type BikeOwnershipRenouncedIterator struct {
	Event *BikeOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *BikeOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeOwnershipRenounced)
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
		it.Event = new(BikeOwnershipRenounced)
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
func (it *BikeOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeOwnershipRenounced represents a OwnershipRenounced event raised by the Bike contract.
type BikeOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Bike *BikeFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*BikeOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Bike.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BikeOwnershipRenouncedIterator{contract: _Bike.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: e OwnershipRenounced(previousOwner indexed address)
func (_Bike *BikeFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *BikeOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Bike.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeOwnershipRenounced)
				if err := _Bike.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// BikeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Bike contract.
type BikeOwnershipTransferredIterator struct {
	Event *BikeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *BikeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeOwnershipTransferred)
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
		it.Event = new(BikeOwnershipTransferred)
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
func (it *BikeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeOwnershipTransferred represents a OwnershipTransferred event raised by the Bike contract.
type BikeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Bike *BikeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BikeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bike.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BikeOwnershipTransferredIterator{contract: _Bike.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Bike *BikeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BikeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Bike.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeOwnershipTransferred)
				if err := _Bike.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// BikeRentedIterator is returned from FilterRented and is used to iterate over the raw logs and unpacked data for Rented events raised by the Bike contract.
type BikeRentedIterator struct {
	Event *BikeRented // Event containing the contract specifics and raw log

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
func (it *BikeRentedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeRented)
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
		it.Event = new(BikeRented)
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
func (it *BikeRentedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeRentedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeRented represents a Rented event raised by the Bike contract.
type BikeRented struct {
	Renter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRented is a free log retrieval operation binding the contract event 0x18ae6638be53d8a225d874e2755866e292376637d786506bf60355f3d13654b9.
//
// Solidity: e Rented(renter address)
func (_Bike *BikeFilterer) FilterRented(opts *bind.FilterOpts) (*BikeRentedIterator, error) {

	logs, sub, err := _Bike.contract.FilterLogs(opts, "Rented")
	if err != nil {
		return nil, err
	}
	return &BikeRentedIterator{contract: _Bike.contract, event: "Rented", logs: logs, sub: sub}, nil
}

// WatchRented is a free log subscription operation binding the contract event 0x18ae6638be53d8a225d874e2755866e292376637d786506bf60355f3d13654b9.
//
// Solidity: e Rented(renter address)
func (_Bike *BikeFilterer) WatchRented(opts *bind.WatchOpts, sink chan<- *BikeRented) (event.Subscription, error) {

	logs, sub, err := _Bike.contract.WatchLogs(opts, "Rented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeRented)
				if err := _Bike.contract.UnpackLog(event, "Rented", log); err != nil {
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

// BikeReturnedIterator is returned from FilterReturned and is used to iterate over the raw logs and unpacked data for Returned events raised by the Bike contract.
type BikeReturnedIterator struct {
	Event *BikeReturned // Event containing the contract specifics and raw log

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
func (it *BikeReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeReturned)
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
		it.Event = new(BikeReturned)
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
func (it *BikeReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeReturned represents a Returned event raised by the Bike contract.
type BikeReturned struct {
	Renter           common.Address
	RentalExpiryTime *big.Int
	CurrentTime      *big.Int
	Amount           *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReturned is a free log retrieval operation binding the contract event 0x089560b26c90954c01dd43421beac7892c9cf783a8e9089820df3647e6ea4176.
//
// Solidity: e Returned(renter address, rentalExpiryTime uint256, currentTime uint256, amount uint256)
func (_Bike *BikeFilterer) FilterReturned(opts *bind.FilterOpts) (*BikeReturnedIterator, error) {

	logs, sub, err := _Bike.contract.FilterLogs(opts, "Returned")
	if err != nil {
		return nil, err
	}
	return &BikeReturnedIterator{contract: _Bike.contract, event: "Returned", logs: logs, sub: sub}, nil
}

// WatchReturned is a free log subscription operation binding the contract event 0x089560b26c90954c01dd43421beac7892c9cf783a8e9089820df3647e6ea4176.
//
// Solidity: e Returned(renter address, rentalExpiryTime uint256, currentTime uint256, amount uint256)
func (_Bike *BikeFilterer) WatchReturned(opts *bind.WatchOpts, sink chan<- *BikeReturned) (event.Subscription, error) {

	logs, sub, err := _Bike.contract.WatchLogs(opts, "Returned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeReturned)
				if err := _Bike.contract.UnpackLog(event, "Returned", log); err != nil {
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
