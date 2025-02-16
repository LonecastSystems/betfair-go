package betfair

type (
	CountryParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale"`
	}

	CountryCodeResult struct {
		CountryCode string `json:"countryCode"`
		MarketCount int    `json:"marketCount"`
	}
)

func (client *Client) ListCountries(params CountryParams) ([]CountryCodeResult, error) {
	json := []CountryCodeResult{}

	return json, client.GetSports("listCountries", params, &json)
}
