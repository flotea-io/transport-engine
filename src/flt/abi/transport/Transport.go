// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package transport

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

// TransportABI is the input ABI used to generate the binding from.
const TransportABI = "[{\"constant\":false,\"inputs\":[{\"components\":[{\"name\":\"fromLat\",\"type\":\"bytes10\"},{\"name\":\"fromLng\",\"type\":\"bytes11\"},{\"name\":\"toLat\",\"type\":\"bytes10\"},{\"name\":\"toLng\",\"type\":\"bytes11\"}],\"name\":\"_tripLoc\",\"type\":\"tuple\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_schedule\",\"type\":\"bytes[]\"},{\"name\":\"_places\",\"type\":\"uint8\"},{\"name\":\"_description\",\"type\":\"bytes\"},{\"name\":\"_routeType\",\"type\":\"uint8\"},{\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"createTrip\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isNew\",\"type\":\"bool\"},{\"name\":\"_companyWallet\",\"type\":\"address\"},{\"name\":\"_company\",\"type\":\"bytes32\"},{\"name\":\"_web\",\"type\":\"bytes32\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"emitCarrier\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tripId\",\"type\":\"uint256\"},{\"name\":\"_tickets\",\"type\":\"uint256\"},{\"name\":\"_buyerAddr\",\"type\":\"address\"},{\"name\":\"_price\",\"type\":\"uint256\"},{\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"emitPurchasedTicket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tripId\",\"type\":\"uint256\"},{\"name\":\"_tickets\",\"type\":\"uint256\"},{\"name\":\"_buyerAddr\",\"type\":\"address\"}],\"name\":\"emitRefundedTickets\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tripId\",\"type\":\"uint256\"},{\"name\":\"updateType\",\"type\":\"string\"}],\"name\":\"emitTripUpdateEvent\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"setEnabled\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"name\":\"_votingCarrierAddress\",\"type\":\"address\"},{\"name\":\"_carriersAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"votingCarrier\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"carriers\",\"type\":\"address\"}],\"name\":\"TransportInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_companyWallet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_company\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_web\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"NewCarrier\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_companyWallet\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_company\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_web\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"CarrierUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_trip\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tripId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_eventType\",\"type\":\"string\"}],\"name\":\"TripEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_trip\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tripId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tickets\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_buyerAddr\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_time\",\"type\":\"uint256\"}],\"name\":\"PurchasedTickets\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_trip\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tripId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_tickets\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_buyerAddr\",\"type\":\"address\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCarriers\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"trips\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"},{\"name\":\"enabled\",\"type\":\"bool\"},{\"name\":\"exist\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"tripsId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tripsLength\",\"outputs\":[{\"name\":\"length\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Transport is an auto generated Go binding around an Ethereum contract.
type Transport struct {
	TransportCaller     // Read-only binding to the contract
	TransportTransactor // Write-only binding to the contract
	TransportFilterer   // Log filterer for contract events
}

// TransportCaller is an auto generated read-only Go binding around an Ethereum contract.
type TransportCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransportTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TransportTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransportFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TransportFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TransportSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TransportSession struct {
	Contract     *Transport        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TransportCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TransportCallerSession struct {
	Contract *TransportCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TransportTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TransportTransactorSession struct {
	Contract     *TransportTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TransportRaw is an auto generated low-level Go binding around an Ethereum contract.
type TransportRaw struct {
	Contract *Transport // Generic contract binding to access the raw methods on
}

// TransportCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TransportCallerRaw struct {
	Contract *TransportCaller // Generic read-only contract binding to access the raw methods on
}

// TransportTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TransportTransactorRaw struct {
	Contract *TransportTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTransport creates a new instance of Transport, bound to a specific deployed contract.
func NewTransport(address common.Address, backend bind.ContractBackend) (*Transport, error) {
	contract, err := bindTransport(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Transport{TransportCaller: TransportCaller{contract: contract}, TransportTransactor: TransportTransactor{contract: contract}, TransportFilterer: TransportFilterer{contract: contract}}, nil
}

// NewTransportCaller creates a new read-only instance of Transport, bound to a specific deployed contract.
func NewTransportCaller(address common.Address, caller bind.ContractCaller) (*TransportCaller, error) {
	contract, err := bindTransport(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TransportCaller{contract: contract}, nil
}

// NewTransportTransactor creates a new write-only instance of Transport, bound to a specific deployed contract.
func NewTransportTransactor(address common.Address, transactor bind.ContractTransactor) (*TransportTransactor, error) {
	contract, err := bindTransport(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TransportTransactor{contract: contract}, nil
}

// NewTransportFilterer creates a new log filterer instance of Transport, bound to a specific deployed contract.
func NewTransportFilterer(address common.Address, filterer bind.ContractFilterer) (*TransportFilterer, error) {
	contract, err := bindTransport(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TransportFilterer{contract: contract}, nil
}

// bindTransport binds a generic wrapper to an already deployed contract.
func bindTransport(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TransportABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transport *TransportRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transport.Contract.TransportCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transport *TransportRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transport.Contract.TransportTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transport *TransportRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transport.Contract.TransportTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Transport *TransportCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Transport.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Transport *TransportTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Transport.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Transport *TransportTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Transport.Contract.contract.Transact(opts, method, params...)
}

// Struct0 is an auto generated low-level Go binding around an user-defined struct.
type Struct0 struct {
	FromLat [10]byte
	FromLng [11]byte
	ToLat   [10]byte
	ToLng   [11]byte
}

// GetCarriers is a free data retrieval call binding the contract method 0x0323d5ce.
//
// Solidity: function getCarriers() constant returns(address)
func (_Transport *TransportCaller) GetCarriers(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Transport.contract.Call(opts, out, "getCarriers")
	return *ret0, err
}

// GetCarriers is a free data retrieval call binding the contract method 0x0323d5ce.
//
// Solidity: function getCarriers() constant returns(address)
func (_Transport *TransportSession) GetCarriers() (common.Address, error) {
	return _Transport.Contract.GetCarriers(&_Transport.CallOpts)
}

// GetCarriers is a free data retrieval call binding the contract method 0x0323d5ce.
//
// Solidity: function getCarriers() constant returns(address)
func (_Transport *TransportCallerSession) GetCarriers() (common.Address, error) {
	return _Transport.Contract.GetCarriers(&_Transport.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_Transport *TransportCaller) TokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Transport.contract.Call(opts, out, "tokenAddress")
	return *ret0, err
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_Transport *TransportSession) TokenAddress() (common.Address, error) {
	return _Transport.Contract.TokenAddress(&_Transport.CallOpts)
}

// TokenAddress is a free data retrieval call binding the contract method 0x9d76ea58.
//
// Solidity: function tokenAddress() constant returns(address)
func (_Transport *TransportCallerSession) TokenAddress() (common.Address, error) {
	return _Transport.Contract.TokenAddress(&_Transport.CallOpts)
}

// Trips is a free data retrieval call binding the contract method 0xf6824a7c.
//
// Solidity: function trips(uint256 ) constant returns(address addr, bool enabled, bool exist)
func (_Transport *TransportCaller) Trips(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr    common.Address
	Enabled bool
	Exist   bool
}, error) {
	ret := new(struct {
		Addr    common.Address
		Enabled bool
		Exist   bool
	})
	out := ret
	err := _Transport.contract.Call(opts, out, "trips", arg0)
	return *ret, err
}

// Trips is a free data retrieval call binding the contract method 0xf6824a7c.
//
// Solidity: function trips(uint256 ) constant returns(address addr, bool enabled, bool exist)
func (_Transport *TransportSession) Trips(arg0 *big.Int) (struct {
	Addr    common.Address
	Enabled bool
	Exist   bool
}, error) {
	return _Transport.Contract.Trips(&_Transport.CallOpts, arg0)
}

// Trips is a free data retrieval call binding the contract method 0xf6824a7c.
//
// Solidity: function trips(uint256 ) constant returns(address addr, bool enabled, bool exist)
func (_Transport *TransportCallerSession) Trips(arg0 *big.Int) (struct {
	Addr    common.Address
	Enabled bool
	Exist   bool
}, error) {
	return _Transport.Contract.Trips(&_Transport.CallOpts, arg0)
}

// TripsId is a free data retrieval call binding the contract method 0xffbb6bbb.
//
// Solidity: function tripsId(address ) constant returns(uint256)
func (_Transport *TransportCaller) TripsId(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Transport.contract.Call(opts, out, "tripsId", arg0)
	return *ret0, err
}

// TripsId is a free data retrieval call binding the contract method 0xffbb6bbb.
//
// Solidity: function tripsId(address ) constant returns(uint256)
func (_Transport *TransportSession) TripsId(arg0 common.Address) (*big.Int, error) {
	return _Transport.Contract.TripsId(&_Transport.CallOpts, arg0)
}

// TripsId is a free data retrieval call binding the contract method 0xffbb6bbb.
//
// Solidity: function tripsId(address ) constant returns(uint256)
func (_Transport *TransportCallerSession) TripsId(arg0 common.Address) (*big.Int, error) {
	return _Transport.Contract.TripsId(&_Transport.CallOpts, arg0)
}

// TripsLength is a free data retrieval call binding the contract method 0x454a496d.
//
// Solidity: function tripsLength() constant returns(uint256 length)
func (_Transport *TransportCaller) TripsLength(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Transport.contract.Call(opts, out, "tripsLength")
	return *ret0, err
}

// TripsLength is a free data retrieval call binding the contract method 0x454a496d.
//
// Solidity: function tripsLength() constant returns(uint256 length)
func (_Transport *TransportSession) TripsLength() (*big.Int, error) {
	return _Transport.Contract.TripsLength(&_Transport.CallOpts)
}

// TripsLength is a free data retrieval call binding the contract method 0x454a496d.
//
// Solidity: function tripsLength() constant returns(uint256 length)
func (_Transport *TransportCallerSession) TripsLength() (*big.Int, error) {
	return _Transport.Contract.TripsLength(&_Transport.CallOpts)
}

// CreateTrip is a paid mutator transaction binding the contract method 0x21dfa7d2.
//
// Solidity: function createTrip(Struct0 _tripLoc, uint256 _price, bytes[] _schedule, uint8 _places, bytes _description, uint8 _routeType, bool _enabled) returns()
func (_Transport *TransportTransactor) CreateTrip(opts *bind.TransactOpts, _tripLoc Struct0, _price *big.Int, _schedule [][]byte, _places uint8, _description []byte, _routeType uint8, _enabled bool) (*types.Transaction, error) {
	return _Transport.contract.Transact(opts, "createTrip", _tripLoc, _price, _schedule, _places, _description, _routeType, _enabled)
}

// CreateTrip is a paid mutator transaction binding the contract method 0x21dfa7d2.
//
// Solidity: function createTrip(Struct0 _tripLoc, uint256 _price, bytes[] _schedule, uint8 _places, bytes _description, uint8 _routeType, bool _enabled) returns()
func (_Transport *TransportSession) CreateTrip(_tripLoc Struct0, _price *big.Int, _schedule [][]byte, _places uint8, _description []byte, _routeType uint8, _enabled bool) (*types.Transaction, error) {
	return _Transport.Contract.CreateTrip(&_Transport.TransactOpts, _tripLoc, _price, _schedule, _places, _description, _routeType, _enabled)
}

// CreateTrip is a paid mutator transaction binding the contract method 0x21dfa7d2.
//
// Solidity: function createTrip(Struct0 _tripLoc, uint256 _price, bytes[] _schedule, uint8 _places, bytes _description, uint8 _routeType, bool _enabled) returns()
func (_Transport *TransportTransactorSession) CreateTrip(_tripLoc Struct0, _price *big.Int, _schedule [][]byte, _places uint8, _description []byte, _routeType uint8, _enabled bool) (*types.Transaction, error) {
	return _Transport.Contract.CreateTrip(&_Transport.TransactOpts, _tripLoc, _price, _schedule, _places, _description, _routeType, _enabled)
}

// EmitCarrier is a paid mutator transaction binding the contract method 0x806523ba.
//
// Solidity: function emitCarrier(bool isNew, address _companyWallet, bytes32 _company, bytes32 _web, uint256 _index) returns()
func (_Transport *TransportTransactor) EmitCarrier(opts *bind.TransactOpts, isNew bool, _companyWallet common.Address, _company [32]byte, _web [32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Transport.contract.Transact(opts, "emitCarrier", isNew, _companyWallet, _company, _web, _index)
}

// EmitCarrier is a paid mutator transaction binding the contract method 0x806523ba.
//
// Solidity: function emitCarrier(bool isNew, address _companyWallet, bytes32 _company, bytes32 _web, uint256 _index) returns()
func (_Transport *TransportSession) EmitCarrier(isNew bool, _companyWallet common.Address, _company [32]byte, _web [32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Transport.Contract.EmitCarrier(&_Transport.TransactOpts, isNew, _companyWallet, _company, _web, _index)
}

// EmitCarrier is a paid mutator transaction binding the contract method 0x806523ba.
//
// Solidity: function emitCarrier(bool isNew, address _companyWallet, bytes32 _company, bytes32 _web, uint256 _index) returns()
func (_Transport *TransportTransactorSession) EmitCarrier(isNew bool, _companyWallet common.Address, _company [32]byte, _web [32]byte, _index *big.Int) (*types.Transaction, error) {
	return _Transport.Contract.EmitCarrier(&_Transport.TransactOpts, isNew, _companyWallet, _company, _web, _index)
}

// EmitPurchasedTicket is a paid mutator transaction binding the contract method 0x2dcac642.
//
// Solidity: function emitPurchasedTicket(uint256 _tripId, uint256 _tickets, address _buyerAddr, uint256 _price, uint256 _time) returns()
func (_Transport *TransportTransactor) EmitPurchasedTicket(opts *bind.TransactOpts, _tripId *big.Int, _tickets *big.Int, _buyerAddr common.Address, _price *big.Int, _time *big.Int) (*types.Transaction, error) {
	return _Transport.contract.Transact(opts, "emitPurchasedTicket", _tripId, _tickets, _buyerAddr, _price, _time)
}

// EmitPurchasedTicket is a paid mutator transaction binding the contract method 0x2dcac642.
//
// Solidity: function emitPurchasedTicket(uint256 _tripId, uint256 _tickets, address _buyerAddr, uint256 _price, uint256 _time) returns()
func (_Transport *TransportSession) EmitPurchasedTicket(_tripId *big.Int, _tickets *big.Int, _buyerAddr common.Address, _price *big.Int, _time *big.Int) (*types.Transaction, error) {
	return _Transport.Contract.EmitPurchasedTicket(&_Transport.TransactOpts, _tripId, _tickets, _buyerAddr, _price, _time)
}

// EmitPurchasedTicket is a paid mutator transaction binding the contract method 0x2dcac642.
//
// Solidity: function emitPurchasedTicket(uint256 _tripId, uint256 _tickets, address _buyerAddr, uint256 _price, uint256 _time) returns()
func (_Transport *TransportTransactorSession) EmitPurchasedTicket(_tripId *big.Int, _tickets *big.Int, _buyerAddr common.Address, _price *big.Int, _time *big.Int) (*types.Transaction, error) {
	return _Transport.Contract.EmitPurchasedTicket(&_Transport.TransactOpts, _tripId, _tickets, _buyerAddr, _price, _time)
}

// EmitRefundedTickets is a paid mutator transaction binding the contract method 0xca52055e.
//
// Solidity: function emitRefundedTickets(uint256 _tripId, uint256 _tickets, address _buyerAddr) returns()
func (_Transport *TransportTransactor) EmitRefundedTickets(opts *bind.TransactOpts, _tripId *big.Int, _tickets *big.Int, _buyerAddr common.Address) (*types.Transaction, error) {
	return _Transport.contract.Transact(opts, "emitRefundedTickets", _tripId, _tickets, _buyerAddr)
}

// EmitRefundedTickets is a paid mutator transaction binding the contract method 0xca52055e.
//
// Solidity: function emitRefundedTickets(uint256 _tripId, uint256 _tickets, address _buyerAddr) returns()
func (_Transport *TransportSession) EmitRefundedTickets(_tripId *big.Int, _tickets *big.Int, _buyerAddr common.Address) (*types.Transaction, error) {
	return _Transport.Contract.EmitRefundedTickets(&_Transport.TransactOpts, _tripId, _tickets, _buyerAddr)
}

// EmitRefundedTickets is a paid mutator transaction binding the contract method 0xca52055e.
//
// Solidity: function emitRefundedTickets(uint256 _tripId, uint256 _tickets, address _buyerAddr) returns()
func (_Transport *TransportTransactorSession) EmitRefundedTickets(_tripId *big.Int, _tickets *big.Int, _buyerAddr common.Address) (*types.Transaction, error) {
	return _Transport.Contract.EmitRefundedTickets(&_Transport.TransactOpts, _tripId, _tickets, _buyerAddr)
}

// EmitTripUpdateEvent is a paid mutator transaction binding the contract method 0x8e91e611.
//
// Solidity: function emitTripUpdateEvent(uint256 _tripId, string updateType) returns()
func (_Transport *TransportTransactor) EmitTripUpdateEvent(opts *bind.TransactOpts, _tripId *big.Int, updateType string) (*types.Transaction, error) {
	return _Transport.contract.Transact(opts, "emitTripUpdateEvent", _tripId, updateType)
}

// EmitTripUpdateEvent is a paid mutator transaction binding the contract method 0x8e91e611.
//
// Solidity: function emitTripUpdateEvent(uint256 _tripId, string updateType) returns()
func (_Transport *TransportSession) EmitTripUpdateEvent(_tripId *big.Int, updateType string) (*types.Transaction, error) {
	return _Transport.Contract.EmitTripUpdateEvent(&_Transport.TransactOpts, _tripId, updateType)
}

// EmitTripUpdateEvent is a paid mutator transaction binding the contract method 0x8e91e611.
//
// Solidity: function emitTripUpdateEvent(uint256 _tripId, string updateType) returns()
func (_Transport *TransportTransactorSession) EmitTripUpdateEvent(_tripId *big.Int, updateType string) (*types.Transaction, error) {
	return _Transport.Contract.EmitTripUpdateEvent(&_Transport.TransactOpts, _tripId, updateType)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool _enabled) returns()
func (_Transport *TransportTransactor) SetEnabled(opts *bind.TransactOpts, _enabled bool) (*types.Transaction, error) {
	return _Transport.contract.Transact(opts, "setEnabled", _enabled)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool _enabled) returns()
func (_Transport *TransportSession) SetEnabled(_enabled bool) (*types.Transaction, error) {
	return _Transport.Contract.SetEnabled(&_Transport.TransactOpts, _enabled)
}

// SetEnabled is a paid mutator transaction binding the contract method 0x328d8f72.
//
// Solidity: function setEnabled(bool _enabled) returns()
func (_Transport *TransportTransactorSession) SetEnabled(_enabled bool) (*types.Transaction, error) {
	return _Transport.Contract.SetEnabled(&_Transport.TransactOpts, _enabled)
}

// TransportCarrierUpdatedIterator is returned from FilterCarrierUpdated and is used to iterate over the raw logs and unpacked data for CarrierUpdated events raised by the Transport contract.
type TransportCarrierUpdatedIterator struct {
	Event *TransportCarrierUpdated // Event containing the contract specifics and raw log

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
func (it *TransportCarrierUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransportCarrierUpdated)
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
		it.Event = new(TransportCarrierUpdated)
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
func (it *TransportCarrierUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransportCarrierUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransportCarrierUpdated represents a CarrierUpdated event raised by the Transport contract.
type TransportCarrierUpdated struct {
	CompanyWallet common.Address
	Company       [32]byte
	Web           [32]byte
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCarrierUpdated is a free log retrieval operation binding the contract event 0xb8a1ffc97098af18c5712ef62fd620b94caefb188f07e05ca8d268fdcfb8450e.
//
// Solidity: event CarrierUpdated(address _companyWallet, bytes32 _company, bytes32 _web, uint256 _index)
func (_Transport *TransportFilterer) FilterCarrierUpdated(opts *bind.FilterOpts) (*TransportCarrierUpdatedIterator, error) {

	logs, sub, err := _Transport.contract.FilterLogs(opts, "CarrierUpdated")
	if err != nil {
		return nil, err
	}
	return &TransportCarrierUpdatedIterator{contract: _Transport.contract, event: "CarrierUpdated", logs: logs, sub: sub}, nil
}

// WatchCarrierUpdated is a free log subscription operation binding the contract event 0xb8a1ffc97098af18c5712ef62fd620b94caefb188f07e05ca8d268fdcfb8450e.
//
// Solidity: event CarrierUpdated(address _companyWallet, bytes32 _company, bytes32 _web, uint256 _index)
func (_Transport *TransportFilterer) WatchCarrierUpdated(opts *bind.WatchOpts, sink chan<- *TransportCarrierUpdated) (event.Subscription, error) {

	logs, sub, err := _Transport.contract.WatchLogs(opts, "CarrierUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransportCarrierUpdated)
				if err := _Transport.contract.UnpackLog(event, "CarrierUpdated", log); err != nil {
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

// ParseCarrierUpdated is a log parse operation binding the contract event 0xb8a1ffc97098af18c5712ef62fd620b94caefb188f07e05ca8d268fdcfb8450e.
//
// Solidity: event CarrierUpdated(address _companyWallet, bytes32 _company, bytes32 _web, uint256 _index)
func (_Transport *TransportFilterer) ParseCarrierUpdated(log types.Log) (*TransportCarrierUpdated, error) {
	event := new(TransportCarrierUpdated)
	if err := _Transport.contract.UnpackLog(event, "CarrierUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransportNewCarrierIterator is returned from FilterNewCarrier and is used to iterate over the raw logs and unpacked data for NewCarrier events raised by the Transport contract.
type TransportNewCarrierIterator struct {
	Event *TransportNewCarrier // Event containing the contract specifics and raw log

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
func (it *TransportNewCarrierIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransportNewCarrier)
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
		it.Event = new(TransportNewCarrier)
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
func (it *TransportNewCarrierIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransportNewCarrierIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransportNewCarrier represents a NewCarrier event raised by the Transport contract.
type TransportNewCarrier struct {
	CompanyWallet common.Address
	Company       [32]byte
	Web           [32]byte
	Index         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewCarrier is a free log retrieval operation binding the contract event 0x7da3acbab78e2cfa0aac005e5b53955fce84566e521fa040c614279d2593d9a6.
//
// Solidity: event NewCarrier(address _companyWallet, bytes32 _company, bytes32 _web, uint256 _index)
func (_Transport *TransportFilterer) FilterNewCarrier(opts *bind.FilterOpts) (*TransportNewCarrierIterator, error) {

	logs, sub, err := _Transport.contract.FilterLogs(opts, "NewCarrier")
	if err != nil {
		return nil, err
	}
	return &TransportNewCarrierIterator{contract: _Transport.contract, event: "NewCarrier", logs: logs, sub: sub}, nil
}

// WatchNewCarrier is a free log subscription operation binding the contract event 0x7da3acbab78e2cfa0aac005e5b53955fce84566e521fa040c614279d2593d9a6.
//
// Solidity: event NewCarrier(address _companyWallet, bytes32 _company, bytes32 _web, uint256 _index)
func (_Transport *TransportFilterer) WatchNewCarrier(opts *bind.WatchOpts, sink chan<- *TransportNewCarrier) (event.Subscription, error) {

	logs, sub, err := _Transport.contract.WatchLogs(opts, "NewCarrier")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransportNewCarrier)
				if err := _Transport.contract.UnpackLog(event, "NewCarrier", log); err != nil {
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

// ParseNewCarrier is a log parse operation binding the contract event 0x7da3acbab78e2cfa0aac005e5b53955fce84566e521fa040c614279d2593d9a6.
//
// Solidity: event NewCarrier(address _companyWallet, bytes32 _company, bytes32 _web, uint256 _index)
func (_Transport *TransportFilterer) ParseNewCarrier(log types.Log) (*TransportNewCarrier, error) {
	event := new(TransportNewCarrier)
	if err := _Transport.contract.UnpackLog(event, "NewCarrier", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransportPurchasedTicketsIterator is returned from FilterPurchasedTickets and is used to iterate over the raw logs and unpacked data for PurchasedTickets events raised by the Transport contract.
type TransportPurchasedTicketsIterator struct {
	Event *TransportPurchasedTickets // Event containing the contract specifics and raw log

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
func (it *TransportPurchasedTicketsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransportPurchasedTickets)
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
		it.Event = new(TransportPurchasedTickets)
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
func (it *TransportPurchasedTicketsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransportPurchasedTicketsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransportPurchasedTickets represents a PurchasedTickets event raised by the Transport contract.
type TransportPurchasedTickets struct {
	Trip      common.Address
	TripId    *big.Int
	Tickets   *big.Int
	BuyerAddr common.Address
	Price     *big.Int
	Time      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPurchasedTickets is a free log retrieval operation binding the contract event 0x08cdff2d5568fedc99785db826cc06fdaa0e53e1fd32271717d8dd1ee73dd500.
//
// Solidity: event PurchasedTickets(address _trip, uint256 _tripId, uint256 _tickets, address _buyerAddr, uint256 _price, uint256 _time)
func (_Transport *TransportFilterer) FilterPurchasedTickets(opts *bind.FilterOpts) (*TransportPurchasedTicketsIterator, error) {

	logs, sub, err := _Transport.contract.FilterLogs(opts, "PurchasedTickets")
	if err != nil {
		return nil, err
	}
	return &TransportPurchasedTicketsIterator{contract: _Transport.contract, event: "PurchasedTickets", logs: logs, sub: sub}, nil
}

// WatchPurchasedTickets is a free log subscription operation binding the contract event 0x08cdff2d5568fedc99785db826cc06fdaa0e53e1fd32271717d8dd1ee73dd500.
//
// Solidity: event PurchasedTickets(address _trip, uint256 _tripId, uint256 _tickets, address _buyerAddr, uint256 _price, uint256 _time)
func (_Transport *TransportFilterer) WatchPurchasedTickets(opts *bind.WatchOpts, sink chan<- *TransportPurchasedTickets) (event.Subscription, error) {

	logs, sub, err := _Transport.contract.WatchLogs(opts, "PurchasedTickets")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransportPurchasedTickets)
				if err := _Transport.contract.UnpackLog(event, "PurchasedTickets", log); err != nil {
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

// ParsePurchasedTickets is a log parse operation binding the contract event 0x08cdff2d5568fedc99785db826cc06fdaa0e53e1fd32271717d8dd1ee73dd500.
//
// Solidity: event PurchasedTickets(address _trip, uint256 _tripId, uint256 _tickets, address _buyerAddr, uint256 _price, uint256 _time)
func (_Transport *TransportFilterer) ParsePurchasedTickets(log types.Log) (*TransportPurchasedTickets, error) {
	event := new(TransportPurchasedTickets)
	if err := _Transport.contract.UnpackLog(event, "PurchasedTickets", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransportRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the Transport contract.
type TransportRefundedIterator struct {
	Event *TransportRefunded // Event containing the contract specifics and raw log

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
func (it *TransportRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransportRefunded)
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
		it.Event = new(TransportRefunded)
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
func (it *TransportRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransportRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransportRefunded represents a Refunded event raised by the Transport contract.
type TransportRefunded struct {
	Trip      common.Address
	TripId    *big.Int
	Tickets   *big.Int
	BuyerAddr common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0x3a15507d0ec0c19700dcd5dc99368206f7bf8f5edbe079fb848fe12671420bd0.
//
// Solidity: event Refunded(address _trip, uint256 _tripId, uint256 _tickets, address _buyerAddr)
func (_Transport *TransportFilterer) FilterRefunded(opts *bind.FilterOpts) (*TransportRefundedIterator, error) {

	logs, sub, err := _Transport.contract.FilterLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return &TransportRefundedIterator{contract: _Transport.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0x3a15507d0ec0c19700dcd5dc99368206f7bf8f5edbe079fb848fe12671420bd0.
//
// Solidity: event Refunded(address _trip, uint256 _tripId, uint256 _tickets, address _buyerAddr)
func (_Transport *TransportFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *TransportRefunded) (event.Subscription, error) {

	logs, sub, err := _Transport.contract.WatchLogs(opts, "Refunded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransportRefunded)
				if err := _Transport.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0x3a15507d0ec0c19700dcd5dc99368206f7bf8f5edbe079fb848fe12671420bd0.
//
// Solidity: event Refunded(address _trip, uint256 _tripId, uint256 _tickets, address _buyerAddr)
func (_Transport *TransportFilterer) ParseRefunded(log types.Log) (*TransportRefunded, error) {
	event := new(TransportRefunded)
	if err := _Transport.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransportTransportInitializedIterator is returned from FilterTransportInitialized and is used to iterate over the raw logs and unpacked data for TransportInitialized events raised by the Transport contract.
type TransportTransportInitializedIterator struct {
	Event *TransportTransportInitialized // Event containing the contract specifics and raw log

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
func (it *TransportTransportInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransportTransportInitialized)
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
		it.Event = new(TransportTransportInitialized)
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
func (it *TransportTransportInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransportTransportInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransportTransportInitialized represents a TransportInitialized event raised by the Transport contract.
type TransportTransportInitialized struct {
	VotingCarrier common.Address
	Carriers      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterTransportInitialized is a free log retrieval operation binding the contract event 0xa2dfc5aa015f3928070c6d5ecff4b64908b5a661bc39ae3b76f2fbef4b187d4b.
//
// Solidity: event TransportInitialized(address votingCarrier, address carriers)
func (_Transport *TransportFilterer) FilterTransportInitialized(opts *bind.FilterOpts) (*TransportTransportInitializedIterator, error) {

	logs, sub, err := _Transport.contract.FilterLogs(opts, "TransportInitialized")
	if err != nil {
		return nil, err
	}
	return &TransportTransportInitializedIterator{contract: _Transport.contract, event: "TransportInitialized", logs: logs, sub: sub}, nil
}

// WatchTransportInitialized is a free log subscription operation binding the contract event 0xa2dfc5aa015f3928070c6d5ecff4b64908b5a661bc39ae3b76f2fbef4b187d4b.
//
// Solidity: event TransportInitialized(address votingCarrier, address carriers)
func (_Transport *TransportFilterer) WatchTransportInitialized(opts *bind.WatchOpts, sink chan<- *TransportTransportInitialized) (event.Subscription, error) {

	logs, sub, err := _Transport.contract.WatchLogs(opts, "TransportInitialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransportTransportInitialized)
				if err := _Transport.contract.UnpackLog(event, "TransportInitialized", log); err != nil {
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

// ParseTransportInitialized is a log parse operation binding the contract event 0xa2dfc5aa015f3928070c6d5ecff4b64908b5a661bc39ae3b76f2fbef4b187d4b.
//
// Solidity: event TransportInitialized(address votingCarrier, address carriers)
func (_Transport *TransportFilterer) ParseTransportInitialized(log types.Log) (*TransportTransportInitialized, error) {
	event := new(TransportTransportInitialized)
	if err := _Transport.contract.UnpackLog(event, "TransportInitialized", log); err != nil {
		return nil, err
	}
	return event, nil
}

// TransportTripEventIterator is returned from FilterTripEvent and is used to iterate over the raw logs and unpacked data for TripEvent events raised by the Transport contract.
type TransportTripEventIterator struct {
	Event *TransportTripEvent // Event containing the contract specifics and raw log

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
func (it *TransportTripEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TransportTripEvent)
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
		it.Event = new(TransportTripEvent)
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
func (it *TransportTripEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TransportTripEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TransportTripEvent represents a TripEvent event raised by the Transport contract.
type TransportTripEvent struct {
	Trip      common.Address
	TripId    *big.Int
	EventType string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTripEvent is a free log retrieval operation binding the contract event 0xac79b139c3a0deca65a4720eeaa2fa0375948035a7d6706026a955e46fbd0b12.
//
// Solidity: event TripEvent(address _trip, uint256 _tripId, string _eventType)
func (_Transport *TransportFilterer) FilterTripEvent(opts *bind.FilterOpts) (*TransportTripEventIterator, error) {

	logs, sub, err := _Transport.contract.FilterLogs(opts, "TripEvent")
	if err != nil {
		return nil, err
	}
	return &TransportTripEventIterator{contract: _Transport.contract, event: "TripEvent", logs: logs, sub: sub}, nil
}

// WatchTripEvent is a free log subscription operation binding the contract event 0xac79b139c3a0deca65a4720eeaa2fa0375948035a7d6706026a955e46fbd0b12.
//
// Solidity: event TripEvent(address _trip, uint256 _tripId, string _eventType)
func (_Transport *TransportFilterer) WatchTripEvent(opts *bind.WatchOpts, sink chan<- *TransportTripEvent) (event.Subscription, error) {

	logs, sub, err := _Transport.contract.WatchLogs(opts, "TripEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TransportTripEvent)
				if err := _Transport.contract.UnpackLog(event, "TripEvent", log); err != nil {
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

// ParseTripEvent is a log parse operation binding the contract event 0xac79b139c3a0deca65a4720eeaa2fa0375948035a7d6706026a955e46fbd0b12.
//
// Solidity: event TripEvent(address _trip, uint256 _tripId, string _eventType)
func (_Transport *TransportFilterer) ParseTripEvent(log types.Log) (*TransportTripEvent, error) {
	event := new(TransportTripEvent)
	if err := _Transport.contract.UnpackLog(event, "TripEvent", log); err != nil {
		return nil, err
	}
	return event, nil
}
