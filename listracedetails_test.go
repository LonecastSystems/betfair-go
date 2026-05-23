package betfair

import (
	"context"
	"testing"
)

func TestListRaceDetails(t *testing.T) {
	c := newTestClient(t)

	params := ListRaceDetailsParams{}

	races, err := c.ListRaceDetails(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(races)
	if len == 0 {
		t.Fatal("No races")
	}
}
