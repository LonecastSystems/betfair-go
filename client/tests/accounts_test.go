package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client"
)

func TestGetDeveloperAppKeys(t *testing.T) {
	c := NewTestBetfairClient(t)

	resp, err := c.GetDeveloperAppKeys()
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) == 0 {
		t.Fatal("No apps")
	}
}

func TestGetAccountFunds(t *testing.T) {
	c := NewTestBetfairClient(t)

	resp, err := c.GetAccountFunds(client.AccountDetailsParams{Wallet: "UK"})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Wallet != "UK" {
		t.Fatal("Wallet wrong")
	}
}

func TestGetAccountDetails(t *testing.T) {
	c := NewTestBetfairClient(t)

	resp, err := c.GetAccountDetails()
	if err != nil {
		t.Fatal(err)
	}

	if resp.CurrencyCode != "GBP" {
		t.Fatal("Details wrong")
	}
}

func TestGetAccountStatement(t *testing.T) {
	c := NewTestBetfairClient(t)

	resp, err := c.GetAccountStatement(client.AccountStatementParams{})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.AccountStatements) == 0 {
		t.Fatal("No statements")
	}
}

func TestListCurrencyRates(t *testing.T) {
	c := NewTestBetfairClient(t)

	resp, err := c.ListCurrencyRates(client.CurrencyRateParams{FromCurrency: "GBP"})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) == 0 {
		t.Fatal("No rates")
	}
}
