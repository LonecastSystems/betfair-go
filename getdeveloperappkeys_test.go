package betfair

import (
	"context"
	"testing"
)

func TestGetDeveloperAppKeys(t *testing.T) {
	c := NewTestClient(t)

	resp, err := c.GetDeveloperAppKeys(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) == 0 {
		t.Fatal("No apps")
	}
}
