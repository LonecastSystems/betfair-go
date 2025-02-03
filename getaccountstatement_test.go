package betfairgo

import "testing"

func TestGetAccountStatement(t *testing.T) {
	c := NewTestBetfairClient(t)

	resp, err := c.GetAccountStatement(AccountStatementParams{})
	if err != nil {
		t.Fatal(err)
	}

	if len(resp.AccountStatements) == 0 {
		t.Fatal("No statements")
	}
}
