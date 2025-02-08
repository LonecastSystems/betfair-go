package betfair

import "testing"

func TestGetAccountFunds(t *testing.T) {
	c := NewTestClient(t)

	resp, err := c.GetAccountFunds(AccountDetailsParams{Wallet: "UK"})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Wallet != "UK" {
		t.Fatal("Wallet wrong")
	}
}
