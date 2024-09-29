package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client/rpc"
)

func TestGetAccountFunds(t *testing.T) {
	c := CreateClient(t)

	wallet, err := c.GetAccountFunds(rpc.AccountDetailsParams{Wallet: "UK"})
	if err != nil {
		t.Fatal(err)
	}

	if wallet.Wallet != "UK" {
		t.Fatal("Wallet null")
	}
}
