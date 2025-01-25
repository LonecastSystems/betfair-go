package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/stateless"
)

func TestListRaceDetails(t *testing.T) {
	c := NewTestStatelessClient(t)

	params := stateless.RaceDetailsParams{}

	races, err := c.ListRaceDetails(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(races)
	if len == 0 {
		t.Fatal("No races")
	}
}
