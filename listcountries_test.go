package betfairgo

import "testing"

func TestListCountries(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := CountryParams{Filter: MarketFilter{
		MarketCountries: []string{"GB"},
	}}

	countries, err := c.ListCountries(params)
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
