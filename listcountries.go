package betfair

import "context"

type (
	ListCountriesParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	ListCountriesResult struct {
		CountryCode string `json:"countryCode"`
		MarketCount int    `json:"marketCount"`
	}
)

func (client *Client) ListCountries(ctx context.Context, params ListCountriesParams) ([]ListCountriesResult, error) {
	json := []ListCountriesResult{}

	return json, client.GetSports(ctx, "listCountries", params, &json)
}
