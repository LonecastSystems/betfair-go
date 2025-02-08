package betfair

import "testing"

func TestListVenues(t *testing.T) {
	c := NewTestClient(t)

	params := VenueParams{Filter: MarketFilter{
		MarketCountries: []string{"GB"},
	}}

	venues, err := c.ListVenues(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(venues) == 0 {
		t.Fatal("No venues")
	}
}
