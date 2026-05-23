package betfair

import (
	"context"
	"testing"
)

func TestGetAccountFunds(t *testing.T) {
	c := newTestClient(t)

	resp, err := c.GetAccountFunds(context.Background(), GetAccountFundsParams{Wallet: W_UK})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Wallet != W_UK {
		t.Fatal("Wallet wrong")
	}
}
