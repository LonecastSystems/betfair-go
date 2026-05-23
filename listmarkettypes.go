package betfair

import "context"

type (
	// ListMarketTypesParams are the request parameters for listMarketTypes.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687525/listMarketTypes
	ListMarketTypesParams struct {
		Filter MarketFilter `json:"filter"`           // Filter selects desired markets. All markets that match the criteria in the filter are selected.
		Locale string       `json:"locale,omitempty"` // Locale is the language used for the response. If not specified, the default is returned.
	}

	// ListMarketTypesResult holds a market type (MarketTypeResult).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#MarketTypeResult
	ListMarketTypesResult struct {
		MarketType  string `json:"marketType,omitempty"`  // MarketType is the market type code (e.g. MATCH_ODDS, NEXT_GOAL).
		MarketCount int    `json:"marketCount,omitempty"` // MarketCount is the number of markets associated with this market type.
	}
)

// ListMarketTypes returns a list of market types (i.e. MATCH_ODDS, NEXT_GOAL) associated with the markets selected by the MarketFilter. The market types are always the same, regardless of locale.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687525/listMarketTypes
func (client *Client) ListMarketTypes(ctx context.Context, params ListMarketTypesParams) ([]ListMarketTypesResult, error) {
	json := []ListMarketTypesResult{}

	return json, client.GetSports(ctx, "listMarketTypes", params, &json)
}
