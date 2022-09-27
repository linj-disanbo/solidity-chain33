package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/linj-disanbo/solidity-chain33/user"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

var (
	// 0xf7e7a1eaa30589315fb4b9317c6a3b65166a3a5b
	PrivateKey, _ = crypto.HexToECDSA("1130d8e8b921cc1815048f1ef69a1b9cf1b2d8283f45aa8100b55b4bb6510fee")
	FromAddr      = common.HexToAddress("0x89f4931f1E14b3908DBC33C9927d8EfbA3b2fEdf")
	ToAddr        = common.HexToAddress("0x8f47c360cae278e743072263734b35c8f0e48a1d")
	ContractCdd   = "0xCbDCE717e290aAdb09fbd009DD14dBCAC30Bf8B1"
	//ChainRpc      = "http://124.71.110.109:8547"
	ChainRpc = "http://localhost:8545"
)

// fix sendTransaction get result
type ethclientFix struct {
	conn *rpc.Client
	*ethclient.Client
	lastTxHash string
}

func NewClient(conn *rpc.Client) *ethclientFix {
	client := &ethclientFix{
		conn: conn,
	}
	client.Client = ethclient.NewClient(conn)
	return client
}

func (cli *ethclientFix) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	data, err := tx.MarshalBinary()
	if err != nil {
		return err
	}
	return cli.conn.CallContext(ctx, &cli.lastTxHash, "eth_sendRawTransaction", hexutil.Encode(data))
}

func main() {
	chainRpc := os.Getenv("CHAINRPC")
	if chainRpc == "" {
		chainRpc = ChainRpc
	}
	rpcConn, err := rpc.DialContext(context.Background(), chainRpc)
	if err != nil {
		fmt.Println("eth client: ", err)
		os.Exit(1)
	}
	client := NewClient(rpcConn)
	if err != nil {
		fmt.Println("eth client: ", err)
		os.Exit(1)
	}

	cid, err := client.ChainID(context.Background())
	if err != nil {
		fmt.Println("eth client ChainID: ", err)
		os.Exit(1)
	}
	fmt.Println("chainid", cid)

	//privkey1, err := crypto.HexToECDSA(PrivateKey)
	pubkey := PrivateKey.Public()
	pubkeyEcDSA, ok := pubkey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("get pubkey failed")
		os.Exit(1)
	}

	fromAddr := crypto.PubkeyToAddress(*pubkeyEcDSA)
	_ = fromAddr

	nonce, err := client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		fmt.Println("get PendingNonceAt failed", err)
		os.Exit(1)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		fmt.Println("SuggestGasPrice ", err)
	}
	fmt.Printf("Suggested gasPrice: %d \n", gasPrice)

	auth, err := bind.NewKeyedTransactorWithChainID(PrivateKey, cid)
	if err != nil {
		fmt.Println("NewKeyedTransactorWithChainID ", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(10000000) // in units
	auth.GasPrice = gasPrice

	// 返回的 contract address , tx hash 都是不对的。
	// 1. SendTransaction 接口不返回 区块链返回值
	// 2. 本地生成的合约地址字母有大小写， 区块链上 是全小写

	//func DeployUser(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *User, error) {
	userAddress, tx, instance, err := user.DeployUser(auth, client)
	if err != nil {
		fmt.Println("DeployUser ", err)
		os.Exit(1)
	}

	fmt.Println("user contract address", userAddress.Hex(), "\t tx hash: ", tx.Hash().Hex(), "clientFixHash: ", client.lastTxHash)

	for false {
		rs, err := client.TransactionReceipt(context.Background(), common.HexToHash(client.lastTxHash))
		if err != nil {
			fmt.Println("TransactionReceipt  ", err)
			time.Sleep(time.Second)
			continue
		}
		fmt.Printf("tx: %v\n", rs)
		break
	}
	time.Sleep(5 * time.Second) // 保证上链
	userContractAddr := userAddress
	_ = userContractAddr
	_ = instance

	nonce, err = client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		fmt.Println("get PendingNonceAt failed", err)
		os.Exit(1)
	}
	auth.Nonce = big.NewInt(int64(nonce))

	txAddP, err := instance.AddPlatform(auth, "pid1", "http://127.0.0.1:6666", []common.Address{})
	if err != nil {
		fmt.Println("AddPlatform  ", err)
		os.Exit(1)
	}
	fmt.Println("AddPlatform tx  ", "tx-return: ", txAddP.Hash().Hex(), "clientFixHash: ", client.lastTxHash)
	time.Sleep(5 * time.Second) // 保证上链

	nonce, err = client.PendingNonceAt(context.Background(), fromAddr)
	if err != nil {
		fmt.Println("get PendingNonceAt failed", err)
		os.Exit(1)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	txAddU, err := instance.AddUserToPlatform(auth, "user1", "pid1")
	if err != nil {
		fmt.Println("AddUserToPlatform  ", err)
		os.Exit(1)
	}
	fmt.Println("AddUserToPlatform tx  ", "tx-return: ", txAddU.Hash().Hex(), "clientFixHash: ", client.lastTxHash)
	time.Sleep(5 * time.Second) // 保证上链

	userInfo, err := instance.GetUser(nil, "user1")
	if err != nil {
		fmt.Println("GetUser  ", err)
		os.Exit(1)
	}
	fmt.Println("GetUser    ", " return: ", userInfo)
}
