package contract

import (
	"context"
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	token "./token"
)

func GetUserBalance(userAddress common.Address) string {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		return "null"
	}

	address := common.HexToAddress("0xA2E92bF423314928CD5Afa44AA4Ad001D6bfE065")
	instance, err := token.NewToken(address, client)
	if err != nil {
		return "null"
	}

	balance, err := instance.BalanceOf(&bind.CallOpts{}, userAddress)
	if err != nil {
		return "null"

	} else {
		balanceStr := balance.String()
		return balanceStr
	}

}

func ByTokenForUser(userAddress common.Address, value int64) string {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		return "null"
	}

	privateKey, err := crypto.HexToECDSA("e06b4df63d710bb26cce84eb39f13ef2247157b3b6ba383c4f84ef40b12b746a")
	if err != nil {
		return "null"
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "null"
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "null"
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "null"
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xA2E92bF423314928CD5Afa44AA4Ad001D6bfE065")
	instance, err := token.NewToken(address, client)
	if err != nil {
		return "null"
	}

	tx, err := instance.Transfer(auth, userAddress, big.NewInt(value))
	if err != nil {
		return "null"
	} else {
		return tx.Hash().String()
	}
}

func TransferUserToken(userAddressFrom, userAddressTo common.Address, value int64) string {
	client, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		return "null"
	}

	privateKey, err := crypto.HexToECDSA("e06b4df63d710bb26cce84eb39f13ef2247157b3b6ba383c4f84ef40b12b746a")
	if err != nil {
		return "null"
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "null"
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return "null"
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return "null"
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	address := common.HexToAddress("0xA2E92bF423314928CD5Afa44AA4Ad001D6bfE065")
	instance, err := token.NewToken(address, client)
	if err != nil {
		return "null"
	}

	tx, err := instance.TransferFrom(auth, userAddressFrom, userAddressTo, big.NewInt(value))
	if err != nil {
		return "null"
	} else {
		return tx.Hash().String()
	}
}
