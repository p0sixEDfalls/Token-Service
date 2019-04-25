package api

import (
	"testing"
)

func TestTokenForUser(t *testing.T) {
	var val int64
	val = -123
	login := "exmple"
	addr := "0x123"
	if BuyTokenForUser(login, val) != "null" {
		t.Error("Error: TestTokenForUser:TestBuyTokenForUser")
	}

	if TransferUserToken(login, addr, val) != "null" {
		t.Error("Error: TestTokenForUser:TransferUserToken")
	}
}
