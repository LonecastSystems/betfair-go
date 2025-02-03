package betfairgo

import "testing"

func TestEvents(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := EventParams{Filter: MarketFilter{
		EventTypeIDs:    []string{"1"},
		MarketCountries: []string{"GB"},
	}}

	events, err := c.ListEvents(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(events)
	if len == 0 {
		t.Fatal("No events")
	}
}
