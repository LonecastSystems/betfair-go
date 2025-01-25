package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client/stateless"
)

func TestGetDeveloperAppKeys(t *testing.T) {
	c := NewTestStatelessClient(t)

	resp, err := c.GetDeveloperAppKeys()
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) == 0 {
		t.Fatal("No apps")
	}
}

func TestGetAccountFunds(t *testing.T) {
	c := NewTestStatelessClient(t)

	resp, err := c.GetAccountFunds(stateless.AccountDetailsParams{Wallet: "UK"})
	if err != nil {
		t.Fatal(err)
	}

	if resp.Wallet != "UK" {
		t.Fatal("Wallet wrong")
	}
}

func TestGetAccountDetails(t *testing.T) {
	c := NewTestStatelessClient(t)

	resp, err := c.GetAccountDetails()
	if err != nil {
		t.Fatal(err)
	}

	if resp.CurrencyCode != "GBP" {
		t.Fatal("Details wrong")
	}
}

func TestGetAccountStatement(t *testing.T) {
	c := NewTestStatelessClient(t)

	resp, err := c.GetAccountStatement(stateless.AccountStatementParams{})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.AccountStatements) == 0 {
		t.Fatal("No statements")
	}
}

func TestListCurrencyRates(t *testing.T) {
	c := NewTestStatelessClient(t)

	resp, err := c.ListCurrencyRates(stateless.CurrencyRateParams{FromCurrency: "GBP"})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) == 0 {
		t.Fatal("No rates")
	}
}
