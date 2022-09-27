// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package user

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

// UserPlatform is an auto generated low-level Go binding around an user-defined struct.
type UserPlatform struct {
	Exists   bool
	Pid      string
	Endpoint string
}

// UserMetaData contains all meta data concerning the User contract.
var UserMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"user\",\"type\":\"string\"}],\"name\":\"LogAddUserToPlatform\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"}],\"name\":\"LogDeletePlatform\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"user\",\"type\":\"string\"}],\"name\":\"LogDeleteUserFromPlatform\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"managers\",\"type\":\"address[]\"}],\"name\":\"LogNewPlatform\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"oldEndpoint\",\"type\":\"string\"}],\"name\":\"LogUpdatePlatformEndPoint\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"},{\"internalType\":\"address[]\",\"name\":\"managers\",\"type\":\"address[]\"}],\"name\":\"addPlatform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"user\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"}],\"name\":\"addUserToPlatform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"}],\"name\":\"deletePlatform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"user\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"}],\"name\":\"deleteUserToPlatform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"user\",\"type\":\"string\"}],\"name\":\"getUser\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"internalType\":\"structUser.Platform[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"pid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"endpoint\",\"type\":\"string\"}],\"name\":\"updatePlatform\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"07c39975": "addPlatform(string,string,address[])",
		"5f5a706e": "addUserToPlatform(string,string)",
		"7fb81ed9": "deletePlatform(string)",
		"54f9942c": "deleteUserToPlatform(string,string)",
		"31feb671": "getUser(string)",
		"bb003ca2": "updatePlatform(string,string)",
	},
	Bin: "0x608060405234801561001057600080fd5b5061149a806100206000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806307c399751461006757806331feb6711461007c57806354f9942c146100a55780635f5a706e146100b85780637fb81ed9146100cb578063bb003ca2146100de575b600080fd5b61007a610075366004610f67565b6100f1565b005b61008f61008a366004611028565b610287565b60405161009c919061114e565b60405180910390f35b61007a6100b3366004610efe565b6104fb565b61007a6100c6366004610efe565b61089d565b61007a6100d9366004610ebe565b610b02565b61007a6100ec366004610efe565b610bf5565b6000868660405161010392919061113e565b604080519182900390912060008181526020819052919091205490915060ff16156101495760405162461bcd60e51b8152600401610140906113a5565b60405180910390fd5b6000818152602081905260409020805460ff19166001178155610170906002018686610d83565b50600081815260208190526040902061018d906001018888610d83565b5060008181526001602081815260408084203385529091528220805460ff191690911790555b8281101561023c576000828152600160208190526040822090918686858181106101ed57634e487b7160e01b600052603260045260246000fd5b90506020020160208101906102029190610e9d565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790556102358160016113ce565b90506101b3565b507fe56559d3938dc969745114a29557b3666a266e290a5bec734d3720750cb6a17c8787878787876040516102769695949392919061122c565b60405180910390a150505050505050565b80516020808301919091206000818152600390925260408220546060929067ffffffffffffffff8111156102cb57634e487b7160e01b600052604160045260246000fd5b60405190808252806020026020018201604052801561030457816020015b6102f1610e07565b8152602001906001900390816102e95790505b50905060005b6000838152600360205260409020548110156104f157600083815260036020526040812080548291908490811061035157634e487b7160e01b600052603260045260246000fd5b906000526020600020015481526020019081526020016000206040518060600160405290816000820160009054906101000a900460ff161515151581526020016001820180546103a0906113fd565b80601f01602080910402602001604051908101604052809291908181526020018280546103cc906113fd565b80156104195780601f106103ee57610100808354040283529160200191610419565b820191906000526020600020905b8154815290600101906020018083116103fc57829003601f168201915b50505050508152602001600282018054610432906113fd565b80601f016020809104026020016040519081016040528092919081815260200182805461045e906113fd565b80156104ab5780601f10610480576101008083540402835291602001916104ab565b820191906000526020600020905b81548152906001019060200180831161048e57829003601f168201915b5050505050815250508282815181106104d457634e487b7160e01b600052603260045260246000fd5b60209081029190910101526104ea8160016113ce565b905061030a565b509150505b919050565b81818080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525084516020808701919091208252819052604090205460ff161515600114925061056b9150505760405162461bcd60e51b815260040161014090611378565b82828080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250845160208087019190912082526001808252604080842033855290925291205460ff1615151492506105e49150505760405162461bcd60e51b815260040161014090611341565b85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f89018190048102820181019092528781529250879150869081908401838280828437600092018290525085516020808801919091208252600281526040808320875188840120845290915290205460ff16151560011492506106979150505760405162461bcd60e51b815260040161014090611313565b600086866040516106a992919061113e565b60405180910390209050600089896040516106c592919061113e565b60408051918290039091206000818152600260209081528382208683529052918220805460ff1916905591505b60008281526003602052604090205481101561085357600082815260036020526040812080548390811061073657634e487b7160e01b600052603260045260246000fd5b906000526020600020015490508381146107505750610841565b60008381526003602052604090205461076b906001906113e6565b8210156107fa576000838152600360205260408120805461078e906001906113e6565b815481106107ac57634e487b7160e01b600052603260045260246000fd5b90600052602060002001549050806003600086815260200190815260200160002084815481106107ec57634e487b7160e01b600052603260045260246000fd5b600091825260209091200155505b600083815260036020526040902080548061082557634e487b7160e01b600052603160045260246000fd5b6001900381819060005260206000200160009055905550610853565b61084c8160016113ce565b90506106f2565b507f1c01dd4435c12015d89d03cc0990b319db9623f7fbe28e7685efe4b873cc262a88888c8c60405161088994939291906111fa565b60405180910390a150505050505050505050565b81818080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525084516020808701919091208252819052604090205460ff161515600114925061090d9150505760405162461bcd60e51b815260040161014090611378565b82828080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250845160208087019190912082526001808252604080842033855290925291205460ff1615151492506109869150505760405162461bcd60e51b815260040161014090611341565b85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525050604080516020601f89018190048102820181019092528781529250879150869081908401838280828437600092018290525085516020808801919091208252600281526040808320875188840120845290915290205460ff16159250610a359150505760405162461bcd60e51b8152600401610140906112e8565b60008686604051610a4792919061113e565b6040518091039020905060008989604051610a6392919061113e565b60408051918290039091206000818152600460205291909120909150610a8a908b8b610d83565b506000818152600360209081526040808320805460018181018355918552838520018690558484526002835281842086855290925291829020805460ff19169091179055517fbbb5c5d953d020807aee259fda440c899c15d9f884c8e4c76b709fa12e1e12c990610889908a908a908e908e906111fa565b81818080601f016020809104026020016040519081016040528093929190818152602001838380828437600092018290525084516020808701919091208252819052604090205460ff1615156001149250610b729150505760405162461bcd60e51b815260040161014090611378565b60008060008585604051610b8792919061113e565b6040518091039020815260200190815260200160002060000160006101000a81548160ff0219169083151502179055507ff7b9f12638798d455d659a6f155bc9d0e19367158edcddf8a9890e1c1a8aec618383604051610be89291906111de565b60405180910390a1505050565b83838080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250845160208087019190912082526001808252604080842033855290925291205460ff161515149250610c6e9150505760405162461bcd60e51b815260040161014090611341565b60008585604051610c8092919061113e565b6040518091039020905060008060008381526020019081526020016000206002018054610cac906113fd565b80601f0160208091040260200160405190810160405280929190818152602001828054610cd8906113fd565b8015610d255780601f10610cfa57610100808354040283529160200191610d25565b820191906000526020600020905b815481529060010190602001808311610d0857829003601f168201915b5050506000858152602081905260409020929350610d4b92600201915087905086610d83565b507f258e152143486b3414d4226f552077b6382e077bef62e0384229e85bb3c5c86087878787856040516102769594939291906112a1565b828054610d8f906113fd565b90600052602060002090601f016020900481019282610db15760008555610df7565b82601f10610dca5782800160ff19823516178555610df7565b82800160010185558215610df7579182015b82811115610df7578235825591602001919060010190610ddc565b50610e03929150610e2a565b5090565b604051806060016040528060001515815260200160608152602001606081525090565b5b80821115610e035760008155600101610e2b565b80356001600160a01b03811681146104f657600080fd5b60008083601f840112610e67578182fd5b50813567ffffffffffffffff811115610e7e578182fd5b602083019150836020828501011115610e9657600080fd5b9250929050565b600060208284031215610eae578081fd5b610eb782610e3f565b9392505050565b60008060208385031215610ed0578081fd5b823567ffffffffffffffff811115610ee6578182fd5b610ef285828601610e56565b90969095509350505050565b60008060008060408587031215610f13578182fd5b843567ffffffffffffffff80821115610f2a578384fd5b610f3688838901610e56565b90965094506020870135915080821115610f4e578384fd5b50610f5b87828801610e56565b95989497509550505050565b60008060008060008060608789031215610f7f578182fd5b863567ffffffffffffffff80821115610f96578384fd5b610fa28a838b01610e56565b90985096506020890135915080821115610fba578384fd5b610fc68a838b01610e56565b90965094506040890135915080821115610fde578384fd5b818901915089601f830112610ff1578384fd5b813581811115610fff578485fd5b8a60208083028501011115611012578485fd5b6020830194508093505050509295509295509295565b6000602080838503121561103a578182fd5b823567ffffffffffffffff80821115611051578384fd5b818501915085601f830112611064578384fd5b8135818111156110765761107661144e565b604051601f8201601f19168101850183811182821017156110995761109961144e565b60405281815283820185018810156110af578586fd5b818585018683013790810190930193909352509392505050565b60008284528282602086013780602084860101526020601f19601f85011685010190509392505050565b60008151808452815b81811015611118576020818501810151868301820152016110fc565b818111156111295782602083870101525b50601f01601f19169290920160200192915050565b6000828483379101908152919050565b60208082528251828201819052600091906040908185019080840286018301878501865b838110156111d057603f198984030185528151606081511515855288820151818a8701526111a2828701826110f3565b915050878201519150848103888601526111bc81836110f3565b968901969450505090860190600101611172565b509098975050505050505050565b6000602082526111f26020830184866110c9565b949350505050565b60006040825261120e6040830186886110c9565b82810360208401526112218185876110c9565b979650505050505050565b60006060825261124060608301888a6110c9565b60208382038185015261125482888a6110c9565b84810360408601528581528692508101835b86811015611292576001600160a01b0361127f85610e3f565b1682529282019290820190600101611266565b509a9950505050505050505050565b6000606082526112b56060830187896110c9565b82810360208401526112c88186886110c9565b905082810360408401526112dc81856110f3565b98975050505050505050565b6020808252601190820152701c1b185d199bdc9b5cc80808195e1a5cdd607a1b604082015260600190565b6020808252601490820152731c1b185d199bdc9b5cc8081b9bdd08195e1a5cdd60621b604082015260600190565b6020808252601e908201527f73656e646572206973206e6f7420706c6174666f726d206d616e616765720000604082015260600190565b6020808252601390820152721c1b185d199bdc9b5cc81b9bdd08195e1a5cdd606a1b604082015260600190565b6020808252600f908201526e706c6174666f726d2065786973747360881b604082015260600190565b600082198211156113e1576113e1611438565b500190565b6000828210156113f8576113f8611438565b500390565b60028104600182168061141157607f821691505b6020821081141561143257634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052604160045260246000fdfea2646970667358221220b17143f1adce24dfb090da6aeea98a1ae7b0fb9a5fe81e336b70dba89921599064736f6c63430008000033",
}

// UserABI is the input ABI used to generate the binding from.
// Deprecated: Use UserMetaData.ABI instead.
var UserABI = UserMetaData.ABI

// Deprecated: Use UserMetaData.Sigs instead.
// UserFuncSigs maps the 4-byte function signature to its string representation.
var UserFuncSigs = UserMetaData.Sigs

// UserBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UserMetaData.Bin instead.
var UserBin = UserMetaData.Bin

// DeployUser deploys a new Ethereum contract, binding an instance of User to it.
func DeployUser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *User, error) {
	parsed, err := UserMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UserBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// User is an auto generated Go binding around an Ethereum contract.
type User struct {
	UserCaller     // Read-only binding to the contract
	UserTransactor // Write-only binding to the contract
	UserFilterer   // Log filterer for contract events
}

// UserCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserSession struct {
	Contract     *User             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCallerSession struct {
	Contract *UserCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserTransactorSession struct {
	Contract     *UserTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRaw struct {
	Contract *User // Generic contract binding to access the raw methods on
}

// UserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCallerRaw struct {
	Contract *UserCaller // Generic read-only contract binding to access the raw methods on
}

// UserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserTransactorRaw struct {
	Contract *UserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUser creates a new instance of User, bound to a specific deployed contract.
func NewUser(address common.Address, backend bind.ContractBackend) (*User, error) {
	contract, err := bindUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// NewUserCaller creates a new read-only instance of User, bound to a specific deployed contract.
func NewUserCaller(address common.Address, caller bind.ContractCaller) (*UserCaller, error) {
	contract, err := bindUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserCaller{contract: contract}, nil
}

// NewUserTransactor creates a new write-only instance of User, bound to a specific deployed contract.
func NewUserTransactor(address common.Address, transactor bind.ContractTransactor) (*UserTransactor, error) {
	contract, err := bindUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserTransactor{contract: contract}, nil
}

// NewUserFilterer creates a new log filterer instance of User, bound to a specific deployed contract.
func NewUserFilterer(address common.Address, filterer bind.ContractFilterer) (*UserFilterer, error) {
	contract, err := bindUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserFilterer{contract: contract}, nil
}

// bindUser binds a generic wrapper to an already deployed contract.
func bindUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.UserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.contract.Transact(opts, method, params...)
}

// GetUser is a free data retrieval call binding the contract method 0x31feb671.
//
// Solidity: function getUser(string user) view returns((bool,string,string)[])
func (_User *UserCaller) GetUser(opts *bind.CallOpts, user string) ([]UserPlatform, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getUser", user)

	if err != nil {
		return *new([]UserPlatform), err
	}

	out0 := *abi.ConvertType(out[0], new([]UserPlatform)).(*[]UserPlatform)

	return out0, err

}

// GetUser is a free data retrieval call binding the contract method 0x31feb671.
//
// Solidity: function getUser(string user) view returns((bool,string,string)[])
func (_User *UserSession) GetUser(user string) ([]UserPlatform, error) {
	return _User.Contract.GetUser(&_User.CallOpts, user)
}

// GetUser is a free data retrieval call binding the contract method 0x31feb671.
//
// Solidity: function getUser(string user) view returns((bool,string,string)[])
func (_User *UserCallerSession) GetUser(user string) ([]UserPlatform, error) {
	return _User.Contract.GetUser(&_User.CallOpts, user)
}

// AddPlatform is a paid mutator transaction binding the contract method 0x07c39975.
//
// Solidity: function addPlatform(string pid, string endpoint, address[] managers) returns()
func (_User *UserTransactor) AddPlatform(opts *bind.TransactOpts, pid string, endpoint string, managers []common.Address) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addPlatform", pid, endpoint, managers)
}

// AddPlatform is a paid mutator transaction binding the contract method 0x07c39975.
//
// Solidity: function addPlatform(string pid, string endpoint, address[] managers) returns()
func (_User *UserSession) AddPlatform(pid string, endpoint string, managers []common.Address) (*types.Transaction, error) {
	return _User.Contract.AddPlatform(&_User.TransactOpts, pid, endpoint, managers)
}

// AddPlatform is a paid mutator transaction binding the contract method 0x07c39975.
//
// Solidity: function addPlatform(string pid, string endpoint, address[] managers) returns()
func (_User *UserTransactorSession) AddPlatform(pid string, endpoint string, managers []common.Address) (*types.Transaction, error) {
	return _User.Contract.AddPlatform(&_User.TransactOpts, pid, endpoint, managers)
}

// AddUserToPlatform is a paid mutator transaction binding the contract method 0x5f5a706e.
//
// Solidity: function addUserToPlatform(string user, string pid) returns()
func (_User *UserTransactor) AddUserToPlatform(opts *bind.TransactOpts, user string, pid string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "addUserToPlatform", user, pid)
}

// AddUserToPlatform is a paid mutator transaction binding the contract method 0x5f5a706e.
//
// Solidity: function addUserToPlatform(string user, string pid) returns()
func (_User *UserSession) AddUserToPlatform(user string, pid string) (*types.Transaction, error) {
	return _User.Contract.AddUserToPlatform(&_User.TransactOpts, user, pid)
}

// AddUserToPlatform is a paid mutator transaction binding the contract method 0x5f5a706e.
//
// Solidity: function addUserToPlatform(string user, string pid) returns()
func (_User *UserTransactorSession) AddUserToPlatform(user string, pid string) (*types.Transaction, error) {
	return _User.Contract.AddUserToPlatform(&_User.TransactOpts, user, pid)
}

// DeletePlatform is a paid mutator transaction binding the contract method 0x7fb81ed9.
//
// Solidity: function deletePlatform(string pid) returns()
func (_User *UserTransactor) DeletePlatform(opts *bind.TransactOpts, pid string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "deletePlatform", pid)
}

// DeletePlatform is a paid mutator transaction binding the contract method 0x7fb81ed9.
//
// Solidity: function deletePlatform(string pid) returns()
func (_User *UserSession) DeletePlatform(pid string) (*types.Transaction, error) {
	return _User.Contract.DeletePlatform(&_User.TransactOpts, pid)
}

// DeletePlatform is a paid mutator transaction binding the contract method 0x7fb81ed9.
//
// Solidity: function deletePlatform(string pid) returns()
func (_User *UserTransactorSession) DeletePlatform(pid string) (*types.Transaction, error) {
	return _User.Contract.DeletePlatform(&_User.TransactOpts, pid)
}

// DeleteUserToPlatform is a paid mutator transaction binding the contract method 0x54f9942c.
//
// Solidity: function deleteUserToPlatform(string user, string pid) returns()
func (_User *UserTransactor) DeleteUserToPlatform(opts *bind.TransactOpts, user string, pid string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "deleteUserToPlatform", user, pid)
}

// DeleteUserToPlatform is a paid mutator transaction binding the contract method 0x54f9942c.
//
// Solidity: function deleteUserToPlatform(string user, string pid) returns()
func (_User *UserSession) DeleteUserToPlatform(user string, pid string) (*types.Transaction, error) {
	return _User.Contract.DeleteUserToPlatform(&_User.TransactOpts, user, pid)
}

// DeleteUserToPlatform is a paid mutator transaction binding the contract method 0x54f9942c.
//
// Solidity: function deleteUserToPlatform(string user, string pid) returns()
func (_User *UserTransactorSession) DeleteUserToPlatform(user string, pid string) (*types.Transaction, error) {
	return _User.Contract.DeleteUserToPlatform(&_User.TransactOpts, user, pid)
}

// UpdatePlatform is a paid mutator transaction binding the contract method 0xbb003ca2.
//
// Solidity: function updatePlatform(string pid, string endpoint) returns()
func (_User *UserTransactor) UpdatePlatform(opts *bind.TransactOpts, pid string, endpoint string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "updatePlatform", pid, endpoint)
}

// UpdatePlatform is a paid mutator transaction binding the contract method 0xbb003ca2.
//
// Solidity: function updatePlatform(string pid, string endpoint) returns()
func (_User *UserSession) UpdatePlatform(pid string, endpoint string) (*types.Transaction, error) {
	return _User.Contract.UpdatePlatform(&_User.TransactOpts, pid, endpoint)
}

// UpdatePlatform is a paid mutator transaction binding the contract method 0xbb003ca2.
//
// Solidity: function updatePlatform(string pid, string endpoint) returns()
func (_User *UserTransactorSession) UpdatePlatform(pid string, endpoint string) (*types.Transaction, error) {
	return _User.Contract.UpdatePlatform(&_User.TransactOpts, pid, endpoint)
}

// UserLogAddUserToPlatformIterator is returned from FilterLogAddUserToPlatform and is used to iterate over the raw logs and unpacked data for LogAddUserToPlatform events raised by the User contract.
type UserLogAddUserToPlatformIterator struct {
	Event *UserLogAddUserToPlatform // Event containing the contract specifics and raw log

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
func (it *UserLogAddUserToPlatformIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserLogAddUserToPlatform)
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
		it.Event = new(UserLogAddUserToPlatform)
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
func (it *UserLogAddUserToPlatformIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserLogAddUserToPlatformIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserLogAddUserToPlatform represents a LogAddUserToPlatform event raised by the User contract.
type UserLogAddUserToPlatform struct {
	Pid  string
	User string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogAddUserToPlatform is a free log retrieval operation binding the contract event 0xbbb5c5d953d020807aee259fda440c899c15d9f884c8e4c76b709fa12e1e12c9.
//
// Solidity: event LogAddUserToPlatform(string pid, string user)
func (_User *UserFilterer) FilterLogAddUserToPlatform(opts *bind.FilterOpts) (*UserLogAddUserToPlatformIterator, error) {

	logs, sub, err := _User.contract.FilterLogs(opts, "LogAddUserToPlatform")
	if err != nil {
		return nil, err
	}
	return &UserLogAddUserToPlatformIterator{contract: _User.contract, event: "LogAddUserToPlatform", logs: logs, sub: sub}, nil
}

// WatchLogAddUserToPlatform is a free log subscription operation binding the contract event 0xbbb5c5d953d020807aee259fda440c899c15d9f884c8e4c76b709fa12e1e12c9.
//
// Solidity: event LogAddUserToPlatform(string pid, string user)
func (_User *UserFilterer) WatchLogAddUserToPlatform(opts *bind.WatchOpts, sink chan<- *UserLogAddUserToPlatform) (event.Subscription, error) {

	logs, sub, err := _User.contract.WatchLogs(opts, "LogAddUserToPlatform")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserLogAddUserToPlatform)
				if err := _User.contract.UnpackLog(event, "LogAddUserToPlatform", log); err != nil {
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

// ParseLogAddUserToPlatform is a log parse operation binding the contract event 0xbbb5c5d953d020807aee259fda440c899c15d9f884c8e4c76b709fa12e1e12c9.
//
// Solidity: event LogAddUserToPlatform(string pid, string user)
func (_User *UserFilterer) ParseLogAddUserToPlatform(log types.Log) (*UserLogAddUserToPlatform, error) {
	event := new(UserLogAddUserToPlatform)
	if err := _User.contract.UnpackLog(event, "LogAddUserToPlatform", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserLogDeletePlatformIterator is returned from FilterLogDeletePlatform and is used to iterate over the raw logs and unpacked data for LogDeletePlatform events raised by the User contract.
type UserLogDeletePlatformIterator struct {
	Event *UserLogDeletePlatform // Event containing the contract specifics and raw log

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
func (it *UserLogDeletePlatformIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserLogDeletePlatform)
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
		it.Event = new(UserLogDeletePlatform)
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
func (it *UserLogDeletePlatformIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserLogDeletePlatformIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserLogDeletePlatform represents a LogDeletePlatform event raised by the User contract.
type UserLogDeletePlatform struct {
	Pid string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogDeletePlatform is a free log retrieval operation binding the contract event 0xf7b9f12638798d455d659a6f155bc9d0e19367158edcddf8a9890e1c1a8aec61.
//
// Solidity: event LogDeletePlatform(string pid)
func (_User *UserFilterer) FilterLogDeletePlatform(opts *bind.FilterOpts) (*UserLogDeletePlatformIterator, error) {

	logs, sub, err := _User.contract.FilterLogs(opts, "LogDeletePlatform")
	if err != nil {
		return nil, err
	}
	return &UserLogDeletePlatformIterator{contract: _User.contract, event: "LogDeletePlatform", logs: logs, sub: sub}, nil
}

// WatchLogDeletePlatform is a free log subscription operation binding the contract event 0xf7b9f12638798d455d659a6f155bc9d0e19367158edcddf8a9890e1c1a8aec61.
//
// Solidity: event LogDeletePlatform(string pid)
func (_User *UserFilterer) WatchLogDeletePlatform(opts *bind.WatchOpts, sink chan<- *UserLogDeletePlatform) (event.Subscription, error) {

	logs, sub, err := _User.contract.WatchLogs(opts, "LogDeletePlatform")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserLogDeletePlatform)
				if err := _User.contract.UnpackLog(event, "LogDeletePlatform", log); err != nil {
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

// ParseLogDeletePlatform is a log parse operation binding the contract event 0xf7b9f12638798d455d659a6f155bc9d0e19367158edcddf8a9890e1c1a8aec61.
//
// Solidity: event LogDeletePlatform(string pid)
func (_User *UserFilterer) ParseLogDeletePlatform(log types.Log) (*UserLogDeletePlatform, error) {
	event := new(UserLogDeletePlatform)
	if err := _User.contract.UnpackLog(event, "LogDeletePlatform", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserLogDeleteUserFromPlatformIterator is returned from FilterLogDeleteUserFromPlatform and is used to iterate over the raw logs and unpacked data for LogDeleteUserFromPlatform events raised by the User contract.
type UserLogDeleteUserFromPlatformIterator struct {
	Event *UserLogDeleteUserFromPlatform // Event containing the contract specifics and raw log

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
func (it *UserLogDeleteUserFromPlatformIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserLogDeleteUserFromPlatform)
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
		it.Event = new(UserLogDeleteUserFromPlatform)
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
func (it *UserLogDeleteUserFromPlatformIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserLogDeleteUserFromPlatformIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserLogDeleteUserFromPlatform represents a LogDeleteUserFromPlatform event raised by the User contract.
type UserLogDeleteUserFromPlatform struct {
	Pid  string
	User string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogDeleteUserFromPlatform is a free log retrieval operation binding the contract event 0x1c01dd4435c12015d89d03cc0990b319db9623f7fbe28e7685efe4b873cc262a.
//
// Solidity: event LogDeleteUserFromPlatform(string pid, string user)
func (_User *UserFilterer) FilterLogDeleteUserFromPlatform(opts *bind.FilterOpts) (*UserLogDeleteUserFromPlatformIterator, error) {

	logs, sub, err := _User.contract.FilterLogs(opts, "LogDeleteUserFromPlatform")
	if err != nil {
		return nil, err
	}
	return &UserLogDeleteUserFromPlatformIterator{contract: _User.contract, event: "LogDeleteUserFromPlatform", logs: logs, sub: sub}, nil
}

// WatchLogDeleteUserFromPlatform is a free log subscription operation binding the contract event 0x1c01dd4435c12015d89d03cc0990b319db9623f7fbe28e7685efe4b873cc262a.
//
// Solidity: event LogDeleteUserFromPlatform(string pid, string user)
func (_User *UserFilterer) WatchLogDeleteUserFromPlatform(opts *bind.WatchOpts, sink chan<- *UserLogDeleteUserFromPlatform) (event.Subscription, error) {

	logs, sub, err := _User.contract.WatchLogs(opts, "LogDeleteUserFromPlatform")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserLogDeleteUserFromPlatform)
				if err := _User.contract.UnpackLog(event, "LogDeleteUserFromPlatform", log); err != nil {
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

// ParseLogDeleteUserFromPlatform is a log parse operation binding the contract event 0x1c01dd4435c12015d89d03cc0990b319db9623f7fbe28e7685efe4b873cc262a.
//
// Solidity: event LogDeleteUserFromPlatform(string pid, string user)
func (_User *UserFilterer) ParseLogDeleteUserFromPlatform(log types.Log) (*UserLogDeleteUserFromPlatform, error) {
	event := new(UserLogDeleteUserFromPlatform)
	if err := _User.contract.UnpackLog(event, "LogDeleteUserFromPlatform", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserLogNewPlatformIterator is returned from FilterLogNewPlatform and is used to iterate over the raw logs and unpacked data for LogNewPlatform events raised by the User contract.
type UserLogNewPlatformIterator struct {
	Event *UserLogNewPlatform // Event containing the contract specifics and raw log

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
func (it *UserLogNewPlatformIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserLogNewPlatform)
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
		it.Event = new(UserLogNewPlatform)
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
func (it *UserLogNewPlatformIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserLogNewPlatformIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserLogNewPlatform represents a LogNewPlatform event raised by the User contract.
type UserLogNewPlatform struct {
	Pid      string
	Endpoint string
	Managers []common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogNewPlatform is a free log retrieval operation binding the contract event 0xe56559d3938dc969745114a29557b3666a266e290a5bec734d3720750cb6a17c.
//
// Solidity: event LogNewPlatform(string pid, string endpoint, address[] managers)
func (_User *UserFilterer) FilterLogNewPlatform(opts *bind.FilterOpts) (*UserLogNewPlatformIterator, error) {

	logs, sub, err := _User.contract.FilterLogs(opts, "LogNewPlatform")
	if err != nil {
		return nil, err
	}
	return &UserLogNewPlatformIterator{contract: _User.contract, event: "LogNewPlatform", logs: logs, sub: sub}, nil
}

// WatchLogNewPlatform is a free log subscription operation binding the contract event 0xe56559d3938dc969745114a29557b3666a266e290a5bec734d3720750cb6a17c.
//
// Solidity: event LogNewPlatform(string pid, string endpoint, address[] managers)
func (_User *UserFilterer) WatchLogNewPlatform(opts *bind.WatchOpts, sink chan<- *UserLogNewPlatform) (event.Subscription, error) {

	logs, sub, err := _User.contract.WatchLogs(opts, "LogNewPlatform")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserLogNewPlatform)
				if err := _User.contract.UnpackLog(event, "LogNewPlatform", log); err != nil {
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

// ParseLogNewPlatform is a log parse operation binding the contract event 0xe56559d3938dc969745114a29557b3666a266e290a5bec734d3720750cb6a17c.
//
// Solidity: event LogNewPlatform(string pid, string endpoint, address[] managers)
func (_User *UserFilterer) ParseLogNewPlatform(log types.Log) (*UserLogNewPlatform, error) {
	event := new(UserLogNewPlatform)
	if err := _User.contract.UnpackLog(event, "LogNewPlatform", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserLogUpdatePlatformEndPointIterator is returned from FilterLogUpdatePlatformEndPoint and is used to iterate over the raw logs and unpacked data for LogUpdatePlatformEndPoint events raised by the User contract.
type UserLogUpdatePlatformEndPointIterator struct {
	Event *UserLogUpdatePlatformEndPoint // Event containing the contract specifics and raw log

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
func (it *UserLogUpdatePlatformEndPointIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserLogUpdatePlatformEndPoint)
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
		it.Event = new(UserLogUpdatePlatformEndPoint)
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
func (it *UserLogUpdatePlatformEndPointIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserLogUpdatePlatformEndPointIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserLogUpdatePlatformEndPoint represents a LogUpdatePlatformEndPoint event raised by the User contract.
type UserLogUpdatePlatformEndPoint struct {
	Pid         string
	Endpoint    string
	OldEndpoint string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogUpdatePlatformEndPoint is a free log retrieval operation binding the contract event 0x258e152143486b3414d4226f552077b6382e077bef62e0384229e85bb3c5c860.
//
// Solidity: event LogUpdatePlatformEndPoint(string pid, string endpoint, string oldEndpoint)
func (_User *UserFilterer) FilterLogUpdatePlatformEndPoint(opts *bind.FilterOpts) (*UserLogUpdatePlatformEndPointIterator, error) {

	logs, sub, err := _User.contract.FilterLogs(opts, "LogUpdatePlatformEndPoint")
	if err != nil {
		return nil, err
	}
	return &UserLogUpdatePlatformEndPointIterator{contract: _User.contract, event: "LogUpdatePlatformEndPoint", logs: logs, sub: sub}, nil
}

// WatchLogUpdatePlatformEndPoint is a free log subscription operation binding the contract event 0x258e152143486b3414d4226f552077b6382e077bef62e0384229e85bb3c5c860.
//
// Solidity: event LogUpdatePlatformEndPoint(string pid, string endpoint, string oldEndpoint)
func (_User *UserFilterer) WatchLogUpdatePlatformEndPoint(opts *bind.WatchOpts, sink chan<- *UserLogUpdatePlatformEndPoint) (event.Subscription, error) {

	logs, sub, err := _User.contract.WatchLogs(opts, "LogUpdatePlatformEndPoint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserLogUpdatePlatformEndPoint)
				if err := _User.contract.UnpackLog(event, "LogUpdatePlatformEndPoint", log); err != nil {
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

// ParseLogUpdatePlatformEndPoint is a log parse operation binding the contract event 0x258e152143486b3414d4226f552077b6382e077bef62e0384229e85bb3c5c860.
//
// Solidity: event LogUpdatePlatformEndPoint(string pid, string endpoint, string oldEndpoint)
func (_User *UserFilterer) ParseLogUpdatePlatformEndPoint(log types.Log) (*UserLogUpdatePlatformEndPoint, error) {
	event := new(UserLogUpdatePlatformEndPoint)
	if err := _User.contract.UnpackLog(event, "LogUpdatePlatformEndPoint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
