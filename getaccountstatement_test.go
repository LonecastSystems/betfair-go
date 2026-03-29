package betfair

import (
	"context"
	"testing"
)

func TestGetAccountStatement(t *testing.T) {
	c := NewTestClient(t)

	_, err := c.GetAccountStatement(context.Background(), AccountStatementParams{})
	if err != nil {
		t.Fatal(err)
	}
}
