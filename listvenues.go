package betfair

import "context"

type (
	ListVenuesParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	ListVenuesResult struct {
		Venue       string `json:"venue"`
		MarketCount int    `json:"marketCount"`
	}
)

func (client *Client) ListVenues(ctx context.Context, params ListVenuesParams) ([]ListVenuesResult, error) {
	json := []ListVenuesResult{}

	return json, client.GetSports(ctx, "listVenues", params, &json)
}
