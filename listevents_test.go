package betfair

import (
	"context"
	"testing"
)

func TestEvents(t *testing.T) {
	c := NewTestClient(t)

	params := EventParams{Filter: MarketFilter{
		EventTypeIDs:    []string{"1"},
		MarketCountries: []string{"GB"},
	}}

	events, err := c.ListEvents(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(events)
	if len == 0 {
		t.Fatal("No events")
	}
}
