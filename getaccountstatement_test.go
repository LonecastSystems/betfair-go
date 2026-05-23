package betfair

import (
	"context"
	"testing"
)

func TestGetAccountStatement(t *testing.T) {
	c := newTestClient(t)

	_, err := c.GetAccountStatement(context.Background(), GetAccountStatementParams{})
	if err != nil {
		t.Fatal(err)
	}
}
