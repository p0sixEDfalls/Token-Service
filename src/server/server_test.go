package server

import (
	"net/http"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestgetClaimFromToken(t *testing.T) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claim := &Claims{
		Login: "Example",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, _ := token.SignedString(JwtKey)

	cookie := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}

	claim2 := GetClaimFromToken(cookie)

	if claim2.Login != claim.Login {
		t.Error("Error: getClaimFromToken:claim.Login")
	}

	if claim2.ExpiresAt != claim.ExpiresAt {
		t.Error("Error: getClaimFromToken:claim.ExpiresAt")
	}
}
