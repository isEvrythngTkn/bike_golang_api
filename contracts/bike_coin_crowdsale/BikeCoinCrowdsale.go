// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bike_coin_crowdsale

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

// BikeCoinCrowdsaleABI is the input ABI used to generate the binding from.
const BikeCoinCrowdsaleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"weiRaised\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"buyTokens\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_rate\",\"type\":\"uint256\"},{\"name\":\"_wallet\",\"type\":\"address\"},{\"name\":\"tokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"purchaser\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"beneficiary\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokenPurchase\",\"type\":\"event\"}]"

// BikeCoinCrowdsaleBin is the compiled bytecode used for deploying new contracts.
const BikeCoinCrowdsaleBin = `608060405234801561001057600080fd5b506040516060806106c283398101806040528101908080519060200190929190805190602001909291908051906020019092919050505082828260008311151561005957600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561009557600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156100d157600080fd5b8260028190555081600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050505050506105548061016e6000396000f30060806040526004361061006d576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680632c4e722e146100785780634042b66f146100a3578063521eb273146100ce578063ec8ac4d814610125578063fc0c546a1461015b575b610076336101b2565b005b34801561008457600080fd5b5061008d610280565b6040518082815260200191505060405180910390f35b3480156100af57600080fd5b506100b8610286565b6040518082815260200191505060405180910390f35b3480156100da57600080fd5b506100e361028c565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610159600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506101b2565b005b34801561016757600080fd5b506101706102b2565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6000803491506101c283836102d7565b6101cb82610327565b90506101e28260035461034590919063ffffffff16565b6003819055506101f28382610361565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad188484604051808381526020018281526020019250505060405180910390a3610269838361036f565b610271610373565b61027b83836103de565b505050565b60025481565b60035481565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561031357600080fd5b6000811415151561032357600080fd5b5050565b600061033e600254836103e290919063ffffffff16565b9050919050565b6000818301905082811015151561035857fe5b80905092915050565b61036b828261041a565b5050565b5050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050501580156103db573d6000803e3d6000fd5b50565b5050565b6000808314156103f55760009050610414565b818302905081838281151561040657fe5b0414151561041057fe5b8090505b92915050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166340c10f1983836040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b1580156104de57600080fd5b505af11580156104f2573d6000803e3d6000fd5b505050506040513d602081101561050857600080fd5b8101908080519060200190929190505050151561052457600080fd5b50505600a165627a7a7230582079b0ea96bc80eba601ac566f51d297b6188e17fcb76c7f4e18731f55c3ff94610029`

// DeployBikeCoinCrowdsale deploys a new Ethereum contract, binding an instance of BikeCoinCrowdsale to it.
func DeployBikeCoinCrowdsale(auth *bind.TransactOpts, backend bind.ContractBackend, _rate *big.Int, _wallet common.Address, tokenAddress common.Address) (common.Address, *types.Transaction, *BikeCoinCrowdsale, error) {
	parsed, err := abi.JSON(strings.NewReader(BikeCoinCrowdsaleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BikeCoinCrowdsaleBin), backend, _rate, _wallet, tokenAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BikeCoinCrowdsale{BikeCoinCrowdsaleCaller: BikeCoinCrowdsaleCaller{contract: contract}, BikeCoinCrowdsaleTransactor: BikeCoinCrowdsaleTransactor{contract: contract}, BikeCoinCrowdsaleFilterer: BikeCoinCrowdsaleFilterer{contract: contract}}, nil
}

// BikeCoinCrowdsale is an auto generated Go binding around an Ethereum contract.
type BikeCoinCrowdsale struct {
	BikeCoinCrowdsaleCaller     // Read-only binding to the contract
	BikeCoinCrowdsaleTransactor // Write-only binding to the contract
	BikeCoinCrowdsaleFilterer   // Log filterer for contract events
}

// BikeCoinCrowdsaleCaller is an auto generated read-only Go binding around an Ethereum contract.
type BikeCoinCrowdsaleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeCoinCrowdsaleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BikeCoinCrowdsaleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeCoinCrowdsaleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BikeCoinCrowdsaleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BikeCoinCrowdsaleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BikeCoinCrowdsaleSession struct {
	Contract     *BikeCoinCrowdsale // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// BikeCoinCrowdsaleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BikeCoinCrowdsaleCallerSession struct {
	Contract *BikeCoinCrowdsaleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// BikeCoinCrowdsaleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BikeCoinCrowdsaleTransactorSession struct {
	Contract     *BikeCoinCrowdsaleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// BikeCoinCrowdsaleRaw is an auto generated low-level Go binding around an Ethereum contract.
type BikeCoinCrowdsaleRaw struct {
	Contract *BikeCoinCrowdsale // Generic contract binding to access the raw methods on
}

// BikeCoinCrowdsaleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BikeCoinCrowdsaleCallerRaw struct {
	Contract *BikeCoinCrowdsaleCaller // Generic read-only contract binding to access the raw methods on
}

// BikeCoinCrowdsaleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BikeCoinCrowdsaleTransactorRaw struct {
	Contract *BikeCoinCrowdsaleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBikeCoinCrowdsale creates a new instance of BikeCoinCrowdsale, bound to a specific deployed contract.
func NewBikeCoinCrowdsale(address common.Address, backend bind.ContractBackend) (*BikeCoinCrowdsale, error) {
	contract, err := bindBikeCoinCrowdsale(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BikeCoinCrowdsale{BikeCoinCrowdsaleCaller: BikeCoinCrowdsaleCaller{contract: contract}, BikeCoinCrowdsaleTransactor: BikeCoinCrowdsaleTransactor{contract: contract}, BikeCoinCrowdsaleFilterer: BikeCoinCrowdsaleFilterer{contract: contract}}, nil
}

// NewBikeCoinCrowdsaleCaller creates a new read-only instance of BikeCoinCrowdsale, bound to a specific deployed contract.
func NewBikeCoinCrowdsaleCaller(address common.Address, caller bind.ContractCaller) (*BikeCoinCrowdsaleCaller, error) {
	contract, err := bindBikeCoinCrowdsale(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BikeCoinCrowdsaleCaller{contract: contract}, nil
}

// NewBikeCoinCrowdsaleTransactor creates a new write-only instance of BikeCoinCrowdsale, bound to a specific deployed contract.
func NewBikeCoinCrowdsaleTransactor(address common.Address, transactor bind.ContractTransactor) (*BikeCoinCrowdsaleTransactor, error) {
	contract, err := bindBikeCoinCrowdsale(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BikeCoinCrowdsaleTransactor{contract: contract}, nil
}

// NewBikeCoinCrowdsaleFilterer creates a new log filterer instance of BikeCoinCrowdsale, bound to a specific deployed contract.
func NewBikeCoinCrowdsaleFilterer(address common.Address, filterer bind.ContractFilterer) (*BikeCoinCrowdsaleFilterer, error) {
	contract, err := bindBikeCoinCrowdsale(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BikeCoinCrowdsaleFilterer{contract: contract}, nil
}

// bindBikeCoinCrowdsale binds a generic wrapper to an already deployed contract.
func bindBikeCoinCrowdsale(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BikeCoinCrowdsaleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BikeCoinCrowdsale.Contract.BikeCoinCrowdsaleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BikeCoinCrowdsale.Contract.BikeCoinCrowdsaleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BikeCoinCrowdsale.Contract.BikeCoinCrowdsaleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BikeCoinCrowdsale.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BikeCoinCrowdsale.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BikeCoinCrowdsale.Contract.contract.Transact(opts, method, params...)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleCaller) Rate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeCoinCrowdsale.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleSession) Rate() (*big.Int, error) {
	return _BikeCoinCrowdsale.Contract.Rate(&_BikeCoinCrowdsale.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() constant returns(uint256)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleCallerSession) Rate() (*big.Int, error) {
	return _BikeCoinCrowdsale.Contract.Rate(&_BikeCoinCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BikeCoinCrowdsale.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleSession) Token() (common.Address, error) {
	return _BikeCoinCrowdsale.Contract.Token(&_BikeCoinCrowdsale.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleCallerSession) Token() (common.Address, error) {
	return _BikeCoinCrowdsale.Contract.Token(&_BikeCoinCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BikeCoinCrowdsale.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleSession) Wallet() (common.Address, error) {
	return _BikeCoinCrowdsale.Contract.Wallet(&_BikeCoinCrowdsale.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleCallerSession) Wallet() (common.Address, error) {
	return _BikeCoinCrowdsale.Contract.Wallet(&_BikeCoinCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleCaller) WeiRaised(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BikeCoinCrowdsale.contract.Call(opts, out, "weiRaised")
	return *ret0, err
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleSession) WeiRaised() (*big.Int, error) {
	return _BikeCoinCrowdsale.Contract.WeiRaised(&_BikeCoinCrowdsale.CallOpts)
}

// WeiRaised is a free data retrieval call binding the contract method 0x4042b66f.
//
// Solidity: function weiRaised() constant returns(uint256)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleCallerSession) WeiRaised() (*big.Int, error) {
	return _BikeCoinCrowdsale.Contract.WeiRaised(&_BikeCoinCrowdsale.CallOpts)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleTransactor) BuyTokens(opts *bind.TransactOpts, _beneficiary common.Address) (*types.Transaction, error) {
	return _BikeCoinCrowdsale.contract.Transact(opts, "buyTokens", _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _BikeCoinCrowdsale.Contract.BuyTokens(&_BikeCoinCrowdsale.TransactOpts, _beneficiary)
}

// BuyTokens is a paid mutator transaction binding the contract method 0xec8ac4d8.
//
// Solidity: function buyTokens(_beneficiary address) returns()
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleTransactorSession) BuyTokens(_beneficiary common.Address) (*types.Transaction, error) {
	return _BikeCoinCrowdsale.Contract.BuyTokens(&_BikeCoinCrowdsale.TransactOpts, _beneficiary)
}

// BikeCoinCrowdsaleTokenPurchaseIterator is returned from FilterTokenPurchase and is used to iterate over the raw logs and unpacked data for TokenPurchase events raised by the BikeCoinCrowdsale contract.
type BikeCoinCrowdsaleTokenPurchaseIterator struct {
	Event *BikeCoinCrowdsaleTokenPurchase // Event containing the contract specifics and raw log

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
func (it *BikeCoinCrowdsaleTokenPurchaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BikeCoinCrowdsaleTokenPurchase)
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
		it.Event = new(BikeCoinCrowdsaleTokenPurchase)
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
func (it *BikeCoinCrowdsaleTokenPurchaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BikeCoinCrowdsaleTokenPurchaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BikeCoinCrowdsaleTokenPurchase represents a TokenPurchase event raised by the BikeCoinCrowdsale contract.
type BikeCoinCrowdsaleTokenPurchase struct {
	Purchaser   common.Address
	Beneficiary common.Address
	Value       *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenPurchase is a free log retrieval operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: e TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleFilterer) FilterTokenPurchase(opts *bind.FilterOpts, purchaser []common.Address, beneficiary []common.Address) (*BikeCoinCrowdsaleTokenPurchaseIterator, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BikeCoinCrowdsale.contract.FilterLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return &BikeCoinCrowdsaleTokenPurchaseIterator{contract: _BikeCoinCrowdsale.contract, event: "TokenPurchase", logs: logs, sub: sub}, nil
}

// WatchTokenPurchase is a free log subscription operation binding the contract event 0x623b3804fa71d67900d064613da8f94b9617215ee90799290593e1745087ad18.
//
// Solidity: e TokenPurchase(purchaser indexed address, beneficiary indexed address, value uint256, amount uint256)
func (_BikeCoinCrowdsale *BikeCoinCrowdsaleFilterer) WatchTokenPurchase(opts *bind.WatchOpts, sink chan<- *BikeCoinCrowdsaleTokenPurchase, purchaser []common.Address, beneficiary []common.Address) (event.Subscription, error) {

	var purchaserRule []interface{}
	for _, purchaserItem := range purchaser {
		purchaserRule = append(purchaserRule, purchaserItem)
	}
	var beneficiaryRule []interface{}
	for _, beneficiaryItem := range beneficiary {
		beneficiaryRule = append(beneficiaryRule, beneficiaryItem)
	}

	logs, sub, err := _BikeCoinCrowdsale.contract.WatchLogs(opts, "TokenPurchase", purchaserRule, beneficiaryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BikeCoinCrowdsaleTokenPurchase)
				if err := _BikeCoinCrowdsale.contract.UnpackLog(event, "TokenPurchase", log); err != nil {
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
