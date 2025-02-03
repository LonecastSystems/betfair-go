package betfairgo

import "testing"

func TestListRaceDetails(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := RaceDetailsParams{}

	races, err := c.ListRaceDetails(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(races)
	if len == 0 {
		t.Fatal("No races")
	}
}
