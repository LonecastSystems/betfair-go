package betfair

import "context"

type (
	ListMarketTypesParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	ListMarketTypesResult struct {
		MarketType  string `json:"marketType,omitempty"`
		MarketCount int    `json:"marketCount,omitempty"`
	}
)

func (client *Client) ListMarketTypes(ctx context.Context, params ListMarketTypesParams) ([]ListMarketTypesResult, error) {
	json := []ListMarketTypesResult{}

	return json, client.GetSports(ctx, "listMarketTypes", params, &json)
}
