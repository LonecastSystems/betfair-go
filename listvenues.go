package betfair

import "context"

type (
	VenueParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	VenueResult struct {
		Venue       string `json:"venue"`
		MarketCount int    `json:"marketCount"`
	}
)

func (client *Client) ListVenues(ctx context.Context, params VenueParams) ([]VenueResult, error) {
	json := []VenueResult{}

	return json, client.GetSports(ctx, "listVenues", params, &json)
}
