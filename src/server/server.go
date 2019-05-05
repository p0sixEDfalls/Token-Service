package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"

	api "../api"
)

var JwtKey = []byte("my_secret_key")

type credentials struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Claims struct {
	Login string `json:"login"`
	jwt.StandardClaims
}

type value struct {
	Value int64 `json:"value"`
}

type addr struct {
	Address string `json:"address"`
}

type tx struct {
	Tx string `json:"tx"`
}

type balanceStr struct {
	Balance string `json:"balance"`
}

type trans struct {
	Addres string `json:"address"`
	Value  int64  `json:"value"`
}

func checkToken(cookie *http.Cookie) bool {
	tknStr := cookie.Value

	claim := &Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claim, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if !tkn.Valid {
		return false
	}
	if err != nil {
		return false
	}

	return true
}

func GetClaimFromToken(cookie *http.Cookie) Claims {
	tknStr := cookie.Value

	claim := &Claims{}

	_, err := jwt.ParseWithClaims(tknStr, claim, func(token *jwt.Token) (interface{}, error) {
		return JwtKey, nil
	})

	if err != nil {
		return Claims{}
	}

	return *claim
}

func refreshToken(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !checkToken(c) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claim := GetClaimFromToken(c)

	if time.Unix(claim.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claim.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func regUser(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !api.AddInvestor(creds.Login, creds.Password) {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func authUser(w http.ResponseWriter, r *http.Request) {
	var creds credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !api.CheckCredits(creds.Login, creds.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claim := &Claims{
		Login: creds.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}

func generateNewAddress(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !checkToken(c) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	address := api.GenerateNewKey(GetClaimFromToken(c).Login)
	addrS := &addr{
		Address: address}

	addrD, _ := json.Marshal(addrS)
	fmt.Fprintf(w, string(addrD))
}

func getUserAddress(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !checkToken(c) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	address := api.GetUserAddress(GetClaimFromToken(c).Login)
	addrS := &addr{
		Address: address}

	addrD, _ := json.Marshal(addrS)
	fmt.Fprintf(w, string(addrD))
}

func buyToken(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !checkToken(c) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var val value
	err = json.NewDecoder(r.Body).Decode(&val)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	txHash := api.BuyTokenForUser(GetClaimFromToken(c).Login, val.Value)
	txHashS := &tx{
		Tx: txHash}

	txHashD, _ := json.Marshal(txHashS)
	fmt.Fprintf(w, string(txHashD))
}

func getUserBalance(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !checkToken(c) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	bal := api.GetUserBalance(GetClaimFromToken(c).Login)
	balanceStrS := &balanceStr{
		Balance: bal}

	balanceStrD, _ := json.Marshal(balanceStrS)
	fmt.Fprintf(w, string(balanceStrD))
}

func sendToken(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if !checkToken(c) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var tr trans
	err = json.NewDecoder(r.Body).Decode(&tr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	txHash := api.TransferUserToken(GetClaimFromToken(c).Login, tr.Addres, tr.Value)
	txHashS := &tx{
		Tx: txHash}

	txHashD, _ := json.Marshal(txHashS)
	fmt.Fprintf(w, string(txHashD))
}

func HandleRequests() bool {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/reguser", regUser).Methods("POST")
	router.HandleFunc("/auth", authUser).Methods("POST")
	router.HandleFunc("/refresh", refreshToken).Methods("GET")
	router.HandleFunc("/api/generate", generateNewAddress).Methods("GET")
	router.HandleFunc("/api/getaddress", getUserAddress).Methods("GET")
	router.HandleFunc("/api/buytoken", buyToken).Methods("POST")
	router.HandleFunc("/api/getbalance", getUserBalance).Methods("GET")
	router.HandleFunc("/api/sendtoken", sendToken).Methods("POST")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		return false
	}
	return true
}
