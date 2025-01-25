package stateless

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

func (client *StatelessClient) ListCountries(params CountryParams) ([]CountryCodeResult, error) {
	json := []CountryCodeResult{}

	if err := GetSports(client, "listCountries", params, &json); err != nil {
		return []CountryCodeResult{}, err
	}

	return json, nil
}
