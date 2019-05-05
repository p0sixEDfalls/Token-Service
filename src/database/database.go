package database

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type investor struct {
	gorm.Model
	Login    string
	Password string
	PrivKey  string
}

func AddInvestor(login, password string) bool {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		return false
	}
	defer db.Close()

	if db.Create(&investor{Login: login, Password: password}).Error != nil {
		return false
	}

	return true
}

func AddUserPrivKey(privkey, login string) bool {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		return false
	}
	defer db.Close()

	var inv investor
	db.Where("login = ?", login).First(&inv)

	if inv.Login == "" {
		return false
	}

	if len(inv.PrivKey) == 0 {
		inv.PrivKey = privkey
		db.Save(&inv)
		return true
	}

	return false
}

func CheckCredits(login, password string) bool {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		return false
	}
	defer db.Close()

	var inv investor
	db.Where("login = ?", login).First(&inv)

	if inv.Login == "" {
		return false
	}

	if inv.Password != password {
		return false
	}

	return true
}

func GetUserAddress(login string) string {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		return "null"
	}
	defer db.Close()

	var inv investor
	db.Where("login = ?", login).First(&inv)

	if len(inv.Login) == 0 {
		return "null"
	}

	if len(inv.PrivKey) == 0 {
		return "null"
	}

	privateKey, err := crypto.HexToECDSA(inv.PrivKey[2:])
	if err != nil {
		return "null"
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, e := publicKey.(*ecdsa.PublicKey)
	if !e {
		return "null"
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address
}

func Init() bool {
	db, err := gorm.Open("postgres", "host=db port=5432 user=postgres dbname=postgres password=password sslmode=disable")
	if err != nil {
		return false
	}
	defer db.Close()
	db.AutoMigrate(&investor{})
	return true
}
