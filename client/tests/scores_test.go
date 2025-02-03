package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client"
)

func TestListRaceDetails(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.RaceDetailsParams{}

	races, err := c.ListRaceDetails(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(races)
	if len == 0 {
		t.Fatal("No races")
	}
}
