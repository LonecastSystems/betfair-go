package betfair

import (
	"context"
	"testing"
)

func TestGetAccountFunds(t *testing.T) {
	c := NewTestClient(t)

	resp, err := c.GetAccountFunds(context.Background(), AccountDetailsParams{Wallet: W_UK})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Wallet != W_UK {
		t.Fatal("Wallet wrong")
	}
}
