package betfair

import (
	"context"
	"testing"
)

func TestListVenues(t *testing.T) {
	c := newTestClient(t)

	params := ListVenuesParams{Filter: MarketFilter{
		MarketCountries: []string{"GB"},
	}}

	venues, err := c.ListVenues(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	if len(venues) == 0 {
		t.Fatal("No venues")
	}
}
