// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package trip

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

// TripABI is the input ABI used to generate the binding from.
const TripABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"charge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_buyer\",\"type\":\"address\"},{\"name\":\"_time\",\"type\":\"uint256\"},{\"name\":\"_count\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"fromLat\",\"type\":\"bytes10\"},{\"name\":\"fromLng\",\"type\":\"bytes11\"},{\"name\":\"toLat\",\"type\":\"bytes10\"},{\"name\":\"toLng\",\"type\":\"bytes11\"}],\"name\":\"_tripLoc\",\"type\":\"tuple\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_schedule\",\"type\":\"bytes[]\"},{\"name\":\"_places\",\"type\":\"uint8\"},{\"name\":\"_description\",\"type\":\"bytes\"}],\"name\":\"setArribute\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setEnabled\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_routeType\",\"type\":\"uint8\"}],\"name\":\"setVehicle\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"tokenFallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_carrierId\",\"type\":\"uint256\"},{\"name\":\"_tripId\",\"type\":\"uint256\"},{\"components\":[{\"name\":\"fromLat\",\"type\":\"bytes10\"},{\"name\":\"fromLng\",\"type\":\"bytes11\"},{\"name\":\"toLat\",\"type\":\"bytes10\"},{\"name\":\"toLng\",\"type\":\"bytes11\"}],\"name\":\"_tripLoc\",\"type\":\"tuple\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_schedule\",\"type\":\"bytes[]\"},{\"name\":\"_places\",\"type\":\"uint8\"},{\"name\":\"_description\",\"type\":\"bytes\"},{\"name\":\"_routeType\",\"type\":\"uint8\"},{\"name\":\"_enabled\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"carrierAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"tripContract\",\"type\":\"address\"}],\"name\":\"Charged\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"availableChargeAmount\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"beforeBuy\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCarrierAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCarrierId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"getTickets\",\"outputs\":[{\"name\":\"ticketsArray\",\"type\":\"uint8\"},{\"name\":\"indexesArray\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTickets\",\"outputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"},{\"name\":\"times\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTripId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"info\",\"outputs\":[{\"name\":\"carrierAddress\",\"type\":\"address\"},{\"name\":\"_carrierId\",\"type\":\"uint256\"},{\"components\":[{\"name\":\"fromLat\",\"type\":\"bytes10\"},{\"name\":\"fromLng\",\"type\":\"bytes11\"},{\"name\":\"toLat\",\"type\":\"bytes10\"},{\"name\":\"toLng\",\"type\":\"bytes11\"}],\"name\":\"_tripLoc\",\"type\":\"tuple\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_schedule\",\"type\":\"bytes[]\"},{\"name\":\"_places\",\"type\":\"uint8\"},{\"name\":\"_description\",\"type\":\"bytes\"},{\"name\":\"_enabled\",\"type\":\"bool\"},{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_routeType\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Trip is an auto generated Go binding around an Ethereum contract.
type Trip struct {
	TripCaller     // Read-only binding to the contract
	TripTransactor // Write-only binding to the contract
	TripFilterer   // Log filterer for contract events
}

// TripCaller is an auto generated read-only Go binding around an Ethereum contract.
type TripCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TripTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TripTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TripFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TripFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TripSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TripSession struct {
	Contract     *Trip             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TripCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TripCallerSession struct {
	Contract *TripCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TripTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TripTransactorSession struct {
	Contract     *TripTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TripRaw is an auto generated low-level Go binding around an Ethereum contract.
type TripRaw struct {
	Contract *Trip // Generic contract binding to access the raw methods on
}

// TripCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TripCallerRaw struct {
	Contract *TripCaller // Generic read-only contract binding to access the raw methods on
}

// TripTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TripTransactorRaw struct {
	Contract *TripTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTrip creates a new instance of Trip, bound to a specific deployed contract.
func NewTrip(address common.Address, backend bind.ContractBackend) (*Trip, error) {
	contract, err := bindTrip(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Trip{TripCaller: TripCaller{contract: contract}, TripTransactor: TripTransactor{contract: contract}, TripFilterer: TripFilterer{contract: contract}}, nil
}

// NewTripCaller creates a new read-only instance of Trip, bound to a specific deployed contract.
func NewTripCaller(address common.Address, caller bind.ContractCaller) (*TripCaller, error) {
	contract, err := bindTrip(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TripCaller{contract: contract}, nil
}

// NewTripTransactor creates a new write-only instance of Trip, bound to a specific deployed contract.
func NewTripTransactor(address common.Address, transactor bind.ContractTransactor) (*TripTransactor, error) {
	contract, err := bindTrip(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TripTransactor{contract: contract}, nil
}

// NewTripFilterer creates a new log filterer instance of Trip, bound to a specific deployed contract.
func NewTripFilterer(address common.Address, filterer bind.ContractFilterer) (*TripFilterer, error) {
	contract, err := bindTrip(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TripFilterer{contract: contract}, nil
}

// bindTrip binds a generic wrapper to an already deployed contract.
func bindTrip(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TripABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trip *TripRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trip.Contract.TripCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trip *TripRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trip.Contract.TripTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trip *TripRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trip.Contract.TripTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Trip *TripCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Trip.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Trip *TripTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Trip.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Trip *TripTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Trip.Contract.contract.Transact(opts, method, params...)
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	FromLat [10]byte
	FromLng [11]byte
	ToLat   [10]byte
	ToLng   [11]byte
}

// AvailableChargeAmount is a free data retrieval call binding the contract method 0x63c192ce.
//
// Solidity: function availableChargeAmount() constant returns(uint256 amount)
func (_Trip *TripCaller) AvailableChargeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Trip.contract.Call(opts, out, "availableChargeAmount")
	return *ret0, err
}

// AvailableChargeAmount is a free data retrieval call binding the contract method 0x63c192ce.
//
// Solidity: function availableChargeAmount() constant returns(uint256 amount)
func (_Trip *TripSession) AvailableChargeAmount() (*big.Int, error) {
	return _Trip.Contract.AvailableChargeAmount(&_Trip.CallOpts)
}

// AvailableChargeAmount is a free data retrieval call binding the contract method 0x63c192ce.
//
// Solidity: function availableChargeAmount() constant returns(uint256 amount)
func (_Trip *TripCallerSession) AvailableChargeAmount() (*big.Int, error) {
	return _Trip.Contract.AvailableChargeAmount(&_Trip.CallOpts)
}

// BeforeBuy is a free data retrieval call binding the contract method 0x102b9b26.
//
// Solidity: function beforeBuy(uint256 _value, bytes _data) constant returns(bool, string)
func (_Trip *TripCaller) BeforeBuy(opts *bind.CallOpts, _value *big.Int, _data []byte) (bool, string, error) {
	var (
		ret0 = new(bool)
		ret1 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Trip.contract.Call(opts, out, "beforeBuy", _value, _data)
	return *ret0, *ret1, err
}

// BeforeBuy is a free data retrieval call binding the contract method 0x102b9b26.
//
// Solidity: function beforeBuy(uint256 _value, bytes _data) constant returns(bool, string)
func (_Trip *TripSession) BeforeBuy(_value *big.Int, _data []byte) (bool, string, error) {
	return _Trip.Contract.BeforeBuy(&_Trip.CallOpts, _value, _data)
}

// BeforeBuy is a free data retrieval call binding the contract method 0x102b9b26.
//
// Solidity: function beforeBuy(uint256 _value, bytes _data) constant returns(bool, string)
func (_Trip *TripCallerSession) BeforeBuy(_value *big.Int, _data []byte) (bool, string, error) {
	return _Trip.Contract.BeforeBuy(&_Trip.CallOpts, _value, _data)
}

// GetCarrierAddress is a free data retrieval call binding the contract method 0xabc14e61.
//
// Solidity: function getCarrierAddress() constant returns(address)
func (_Trip *TripCaller) GetCarrierAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Trip.contract.Call(opts, out, "getCarrierAddress")
	return *ret0, err
}

// GetCarrierAddress is a free data retrieval call binding the contract method 0xabc14e61.
//
// Solidity: function getCarrierAddress() constant returns(address)
func (_Trip *TripSession) GetCarrierAddress() (common.Address, error) {
	return _Trip.Contract.GetCarrierAddress(&_Trip.CallOpts)
}

// GetCarrierAddress is a free data retrieval call binding the contract method 0xabc14e61.
//
// Solidity: function getCarrierAddress() constant returns(address)
func (_Trip *TripCallerSession) GetCarrierAddress() (common.Address, error) {
	return _Trip.Contract.GetCarrierAddress(&_Trip.CallOpts)
}

// GetCarrierId is a free data retrieval call binding the contract method 0x21a88559.
//
// Solidity: function getCarrierId() constant returns(uint256)
func (_Trip *TripCaller) GetCarrierId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Trip.contract.Call(opts, out, "getCarrierId")
	return *ret0, err
}

// GetCarrierId is a free data retrieval call binding the contract method 0x21a88559.
//
// Solidity: function getCarrierId() constant returns(uint256)
func (_Trip *TripSession) GetCarrierId() (*big.Int, error) {
	return _Trip.Contract.GetCarrierId(&_Trip.CallOpts)
}

// GetCarrierId is a free data retrieval call binding the contract method 0x21a88559.
//
// Solidity: function getCarrierId() constant returns(uint256)
func (_Trip *TripCallerSession) GetCarrierId() (*big.Int, error) {
	return _Trip.Contract.GetCarrierId(&_Trip.CallOpts)
}

// GetTickets is a free data retrieval call binding the contract method 0x14f2979f.
//
// Solidity: function getTickets(uint256 _time) constant returns(uint8 ticketsArray, uint256[] indexesArray)
func (_Trip *TripCaller) GetTickets(opts *bind.CallOpts, _time *big.Int) (struct {
	TicketsArray uint8
	IndexesArray []*big.Int
}, error) {
	ret := new(struct {
		TicketsArray uint8
		IndexesArray []*big.Int
	})
	out := ret
	err := _Trip.contract.Call(opts, out, "getTickets", _time)
	return *ret, err
}

// GetTickets is a free data retrieval call binding the contract method 0x14f2979f.
//
// Solidity: function getTickets(uint256 _time) constant returns(uint8 ticketsArray, uint256[] indexesArray)
func (_Trip *TripSession) GetTickets(_time *big.Int) (struct {
	TicketsArray uint8
	IndexesArray []*big.Int
}, error) {
	return _Trip.Contract.GetTickets(&_Trip.CallOpts, _time)
}

// GetTickets is a free data retrieval call binding the contract method 0x14f2979f.
//
// Solidity: function getTickets(uint256 _time) constant returns(uint8 ticketsArray, uint256[] indexesArray)
func (_Trip *TripCallerSession) GetTickets(_time *big.Int) (struct {
	TicketsArray uint8
	IndexesArray []*big.Int
}, error) {
	return _Trip.Contract.GetTickets(&_Trip.CallOpts, _time)
}

// GetTickets0 is a free data retrieval call binding the contract method 0x4ed02622.
//
// Solidity: function getTickets() constant returns(address[] addresses, uint256[] times)
func (_Trip *TripCaller) GetTickets0(opts *bind.CallOpts) (struct {
	Addresses []common.Address
	Times     []*big.Int
}, error) {
	ret := new(struct {
		Addresses []common.Address
		Times     []*big.Int
	})
	out := ret
	err := _Trip.contract.Call(opts, out, "getTickets0")
	return *ret, err
}

// GetTickets0 is a free data retrieval call binding the contract method 0x4ed02622.
//
// Solidity: function getTickets() constant returns(address[] addresses, uint256[] times)
func (_Trip *TripSession) GetTickets0() (struct {
	Addresses []common.Address
	Times     []*big.Int
}, error) {
	return _Trip.Contract.GetTickets0(&_Trip.CallOpts)
}

// GetTickets0 is a free data retrieval call binding the contract method 0x4ed02622.
//
// Solidity: function getTickets() constant returns(address[] addresses, uint256[] times)
func (_Trip *TripCallerSession) GetTickets0() (struct {
	Addresses []common.Address
	Times     []*big.Int
}, error) {
	return _Trip.Contract.GetTickets0(&_Trip.CallOpts)
}

// GetTripId is a free data retrieval call binding the contract method 0xf1d5659e.
//
// Solidity: function getTripId() constant returns(uint256)
func (_Trip *TripCaller) GetTripId(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Trip.contract.Call(opts, out, "getTripId")
	return *ret0, err
}

// GetTripId is a free data retrieval call binding the contract method 0xf1d5659e.
//
// Solidity: function getTripId() constant returns(uint256)
func (_Trip *TripSession) GetTripId() (*big.Int, error) {
	return _Trip.Contract.GetTripId(&_Trip.CallOpts)
}

// GetTripId is a free data retrieval call binding the contract method 0xf1d5659e.
//
// Solidity: function getTripId() constant returns(uint256)
func (_Trip *TripCallerSession) GetTripId() (*big.Int, error) {
	return _Trip.Contract.GetTripId(&_Trip.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() constant returns(address carrierAddress, uint256 _carrierId, Struct0 _tripLoc, uint256 _price, bytes[] _schedule, uint8 _places, bytes _description, bool _enabled, address _token, uint8 _routeType)
func (_Trip *TripCaller) Info(opts *bind.CallOpts) (struct {
	CarrierAddress common.Address
	CarrierId      *big.Int
	TripLoc        Struct0
	Price          *big.Int
	Schedule       [][]byte
	Places         uint8
	Description    []byte
	Enabled        bool
	Token          common.Address
	RouteType      uint8
}, error) {
	ret := new(struct {
		CarrierAddress common.Address
		CarrierId      *big.Int
		TripLoc        Struct0
		Price          *big.Int
		Schedule       [][]byte
		Places         uint8
		Description    []byte
		Enabled        bool
		Token          common.Address
		RouteType      uint8
	})
	out := ret
	err := _Trip.contract.Call(opts, out, "info")
	return *ret, err
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() constant returns(address carrierAddress, uint256 _carrierId, Struct0 _tripLoc, uint256 _price, bytes[] _schedule, uint8 _places, bytes _description, bool _enabled, address _token, uint8 _routeType)
func (_Trip *TripSession) Info() (struct {
	CarrierAddress common.Address
	CarrierId      *big.Int
	TripLoc        Struct0
	Price          *big.Int
	Schedule       [][]byte
	Places         uint8
	Description    []byte
	Enabled        bool
	Token          common.Address
	RouteType      uint8
}, error) {
	return _Trip.Contract.Info(&_Trip.CallOpts)
}

// Info is a free data retrieval call binding the contract method 0x370158ea.
//
// Solidity: function info() constant returns(address carrierAddress, uint256 _carrierId, Struct0 _tripLoc, uint256 _price, bytes[] _schedule, uint8 _places, bytes _description, bool _enabled, address _token, uint8 _routeType)
func (_Trip *TripCallerSession) Info() (struct {
	CarrierAddress common.Address
	CarrierId      *big.Int
	TripLoc        Struct0
	Price          *big.Int
	Schedule       [][]byte
	Places         uint8
	Description    []byte
	Enabled        bool
	Token          common.Address
	RouteType      uint8
}, error) {
	return _Trip.Contract.Info(&_Trip.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Trip *TripCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Trip.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Trip *TripSession) Owner() (common.Address, error) {
	return _Trip.Contract.Owner(&_Trip.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Trip *TripCallerSession) Owner() (common.Address, error) {
	return _Trip.Contract.Owner(&_Trip.CallOpts)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(address _to) returns()
func (_Trip *TripTransactor) Charge(opts *bind.TransactOpts, _to common.Address) (*types.Transaction, error) {
	return _Trip.contract.Transact(opts, "charge", _to)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(address _to) returns()
func (_Trip *TripSession) Charge(_to common.Address) (*types.Transaction, error) {
	return _Trip.Contract.Charge(&_Trip.TransactOpts, _to)
}

// Charge is a paid mutator transaction binding the contract method 0xfc6bd76a.
//
// Solidity: function charge(address _to) returns()
func (_Trip *TripTransactorSession) Charge(_to common.Address) (*types.Transaction, error) {
	return _Trip.Contract.Charge(&_Trip.TransactOpts, _to)
}

// Refund is a paid mutator transaction binding the contract method 0x18fd8903.
//
// Solidity: function refund(address _buyer, uint256 _time, uint256 _count) returns()
func (_Trip *TripTransactor) Refund(opts *bind.TransactOpts, _buyer common.Address, _time *big.Int, _count *big.Int) (*types.Transaction, error) {
	return _Trip.contract.Transact(opts, "refund", _buyer, _time, _count)
}

// Refund is a paid mutator transaction binding the contract method 0x18fd8903.
//
// Solidity: function refund(address _buyer, uint256 _time, uint256 _count) returns()
func (_Trip *TripSession) Refund(_buyer common.Address, _time *big.Int, _count *big.Int) (*types.Transaction, error) {
	return _Trip.Contract.Refund(&_Trip.TransactOpts, _buyer, _time, _count)
}

// Refund is a paid mutator transaction binding the contract method 0x18fd8903.
//
// Solidity: function refund(address _buyer, uint256 _time, uint256 _count) returns()
func (_Trip *TripTransactorSession) Refund(_buyer common.Address, _time *big.Int, _count *big.Int) (*types.Transaction, error) {
	return _Trip.Contract.Refund(&_Trip.TransactOpts, _buyer, _time, _count)
}

// SetArribute is a paid mutator transaction binding the contract method 0x4a7b590e.
//
// Solidity: function setArribute(Struct0 _tripLoc, uint256 _price, bytes[] _schedule, uint8 _places, bytes _description) returns()
func (_Trip *TripTransactor) SetArribute(opts *bind.TransactOpts, _tripLoc Struct0, _price *big.Int, _schedule [][]byte, _places uint8, _description []byte) (*types.Transaction, error) {
	return _Trip.contract.Transact(opts, "setArribute", _tripLoc, _price, _schedule, _places, _description)
}

// SetArribute is a paid mutator transaction binding the contract method 0x4a7b590e.
//
// Solidity: function setArribute(Struct0 _tripLoc, uint256 _price, bytes[] _schedule, uint8 _places, bytes _description) returns()
func (_Trip *TripSession) SetArribute(_tripLoc Struct0, _price *big.Int, _schedule [][]byte, _places uint8, _description []byte) (*types.Transaction, error) {
	return _Trip.Contract.SetArribute(&_Trip.TransactOpts, _tripLoc, _price, _schedule, _places, _description)
}

// SetArribute is a paid mutator transaction binding the contract method 0x4a7b590e.
//
// Solidity: function setArribute(Struct0 _tripLoc, uint256 _price, bytes[] _schedule, uint8 _places, bytes _description) returns()
func (_Trip *TripTransactorSession) SetArribute(_tripLoc Struct0, _price *big.Int, _schedule [][]byte, _places uint8, _description []byte) (*types.Transaction, error) {
	return _Trip.Contract.SetArribute(&_Trip.TransactOpts, _tripLoc, _price, _schedule, _places, _description)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool _enabled) returns()
func (_Trip *TripTransactor) SetEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Trip.contract.Transact(opts, "setEnabled", _enabled)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool _enabled) returns()
func (_Trip *TripSession) SetEnabled(_enabled bool) (*types.Transaction, error) {
	return _Trip.Contract.SetEnabled(&_Trip.TransactOpts, _enabled)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool _enabled) returns()
func (_Trip *TripTransactorSession) SetEnabled(_enabled bool) (*types.Transaction, error) {
	return _Trip.Contract.SetEnabled(&_Trip.TransactOpts, _enabled)
}

// SetVehicle is a paid mutator transaction binding the contract method 0xb1ab1f50.
//
// Solidity: function setVehicle(uint8 _routeType) returns()
func (_Trip *TripTransactor) SetVehicle(opts *bind.TransactOpts, _routeType uint8) (*types.Transaction, error) {
	return _Trip.contract.Transact(opts, "setVehicle", _routeType)
}

// SetVehicle is a paid mutator transaction binding the contract method 0xb1ab1f50.
//
// Solidity: function setVehicle(uint8 _routeType) returns()
func (_Trip *TripSession) SetVehicle(_routeType uint8) (*types.Transaction, error) {
	return _Trip.Contract.SetVehicle(&_Trip.TransactOpts, _routeType)
}

// SetVehicle is a paid mutator transaction binding the contract method 0xb1ab1f50.
//
// Solidity: function setVehicle(uint8 _routeType) returns()
func (_Trip *TripTransactorSession) SetVehicle(_routeType uint8) (*types.Transaction, error) {
	return _Trip.Contract.SetVehicle(&_Trip.TransactOpts, _routeType)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) returns()
func (_Trip *TripTransactor) TokenFallback(opts *bind.TransactOpts, _from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Trip.contract.Transact(opts, "tokenFallback", _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) returns()
func (_Trip *TripSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Trip.Contract.TokenFallback(&_Trip.TransactOpts, _from, _value, _data)
}

// TokenFallback is a paid mutator transaction binding the contract method 0xc0ee0b8a.
//
// Solidity: function tokenFallback(address _from, uint256 _value, bytes _data) returns()
func (_Trip *TripTransactorSession) TokenFallback(_from common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Trip.Contract.TokenFallback(&_Trip.TransactOpts, _from, _value, _data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trip *TripTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Trip.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trip *TripSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trip.Contract.TransferOwnership(&_Trip.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Trip *TripTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Trip.Contract.TransferOwnership(&_Trip.TransactOpts, newOwner)
}

// TripChargedIterator is returned from FilterCharged and is used to iterate over the raw logs and unpacked data for Charged events raised by the Trip contract.
type TripChargedIterator struct {
	Event *TripCharged // Event containing the contract specifics and raw log

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
func (it *TripChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TripCharged)
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
		it.Event = new(TripCharged)
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
func (it *TripChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TripChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TripCharged represents a Charged event raised by the Trip contract.
type TripCharged struct {
	CarrierAddress common.Address
	Amount         *big.Int
	TripContract   common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCharged is a free log retrieval operation binding the contract event 0x654be7454f73d17698cc6328c334c6b369ce05eeb009e2fc3189178ed31cf8b1.
//
// Solidity: event Charged(address carrierAddress, uint256 amount, address tripContract)
func (_Trip *TripFilterer) FilterCharged(opts *bind.FilterOpts) (*TripChargedIterator, error) {

	logs, sub, err := _Trip.contract.FilterLogs(opts, "Charged")
	if err != nil {
		return nil, err
	}
	return &TripChargedIterator{contract: _Trip.contract, event: "Charged", logs: logs, sub: sub}, nil
}

// WatchCharged is a free log subscription operation binding the contract event 0x654be7454f73d17698cc6328c334c6b369ce05eeb009e2fc3189178ed31cf8b1.
//
// Solidity: event Charged(address carrierAddress, uint256 amount, address tripContract)
func (_Trip *TripFilterer) WatchCharged(opts *bind.WatchOpts, sink chan<- *TripCharged) (event.Subscription, error) {

	logs, sub, err := _Trip.contract.WatchLogs(opts, "Charged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TripCharged)
				if err := _Trip.contract.UnpackLog(event, "Charged", log); err != nil {
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

// ParseCharged is a log parse operation binding the contract event 0x654be7454f73d17698cc6328c334c6b369ce05eeb009e2fc3189178ed31cf8b1.
//
// Solidity: event Charged(address carrierAddress, uint256 amount, address tripContract)
func (_Trip *TripFilterer) ParseCharged(log types.Log) (*TripCharged, error) {
	event := new(TripCharged)
	if err := _Trip.contract.UnpackLog(event, "Charged", log); err != nil {
		return nil, err
	}
	return event, nil
}
