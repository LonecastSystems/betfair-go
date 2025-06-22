package betfair

import "context"

type (
	MarketTypeParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	MarketTypeResult struct {
		MarketType  string `json:"marketType,omitempty"`
		MarketCount int    `json:"marketCount,omitempty"`
	}
)

func (client *Client) ListMarketTypes(ctx context.Context, params MarketTypeParams) ([]MarketTypeResult, error) {
	json := []MarketTypeResult{}

	return json, client.GetSports(ctx, "listMarketTypes", params, &json)
}
