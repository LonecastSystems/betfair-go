package betfair

import "context"

type (
	CountryParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	CountryCodeResult struct {
		CountryCode string `json:"countryCode"`
		MarketCount int    `json:"marketCount"`
	}
)

func (client *Client) ListCountries(ctx context.Context, params CountryParams) ([]CountryCodeResult, error) {
	json := []CountryCodeResult{}

	return json, client.GetSports(ctx, "listCountries", params, &json)
}
