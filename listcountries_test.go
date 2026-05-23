package betfair

import (
	"context"
	"testing"
)

func TestListCountries(t *testing.T) {
	c := newTestClient(t)

	params := ListCountriesParams{Filter: MarketFilter{
		MarketCountries: []string{"GB"},
	}}

	countries, err := c.ListCountries(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(countries)
	if len != 1 {
		t.Fatal(len)
	}

	countryCode := countries[0].CountryCode
	if countryCode != "GB" {
		t.Fatal(countryCode)
	}
}
