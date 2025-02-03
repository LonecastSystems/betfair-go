package betfairgo

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

func (client *BetfairClient) ListCountries(params CountryParams) ([]CountryCodeResult, error) {
	json := []CountryCodeResult{}

	if err := client.GetSports("listCountries", params, &json); err != nil {
		return []CountryCodeResult{}, err
	}

	return json, nil
}
