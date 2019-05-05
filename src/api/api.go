package api

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"

	contract "../contract"
	database "../database"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func AddInvestor(login, password string) bool {
	if database.AddInvestor(login, password) {
		return true
	}
	return false
}

func GenerateNewKey(login string) string {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return "null"
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)
	strPrivKey := hexutil.Encode(privateKeyBytes)
	if !database.AddUserPrivKey(strPrivKey, login) {
		return "null"
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, _ := publicKey.(*ecdsa.PublicKey)
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address
}

func CheckCredits(login, password string) bool {
	if database.CheckCredits(login, password) {
		return true
	}

	return false
}

func GetUserAddress(login string) string {
	address := database.GetUserAddress(login)
	return address
}

func GetUserBalance(login string) string {
	addressStr := database.GetUserAddress(login)
	balanceStr := contract.GetUserBalance(common.HexToAddress(addressStr))
	return balanceStr
}

func BuyTokenForUser(login string, value int64) string {
	if value < 0 {
		return "null"
	}
	addressStr := database.GetUserAddress(login)
	if addressStr == "null" {
		return addressStr
	}
	txHash := contract.ByTokenForUser(common.HexToAddress(addressStr), value)
	return txHash
}

func TransferUserToken(login, userAddressTo string, value int64) string {
	if value < 0 {
		return "null"
	}
	userAddressFrom := database.GetUserAddress(login)
	if userAddressFrom == "null" {
		return userAddressFrom
	}
	txHash := contract.TransferUserToken(common.HexToAddress(userAddressFrom), common.HexToAddress(userAddressTo), value)
	return txHash
}
