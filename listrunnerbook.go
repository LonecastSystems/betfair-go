package betfair

import (
	"context"
	"time"
)

type (
	// ListRunnerBookParams are the request parameters for listRunnerBook.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687847/listRunnerBook
	ListRunnerBookParams struct {
		MarketID                      string          `json:"marketId"`                                // MarketID is the unique id for the market.
		SelectionID                   int64           `json:"selectionId"`                             // SelectionID is the unique id for the selection in the market.
		Handicap                      float64         `json:"handicap,omitempty"`                      // Handicap is the handicap associated with the runner in case of Asian handicap market.
		PriceProjection               PriceProjection `json:"priceProjection,omitempty"`               // PriceProjection is the projection of price data you want to receive in the response.
		OrderProjection               OrderProjection `json:"orderProjection,omitempty"`               // OrderProjection is the orders you want to receive in the response.
		MatchProjection               MatchProjection `json:"matchProjection,omitempty"`               // MatchProjection specifies the representation of matches if you ask for orders.
		IncludeOverallPosition        bool            `json:"includeOverallPosition,omitempty"`        // IncludeOverallPosition returns matches for each selection if you ask for orders. Defaults to true if unspecified.
		PartitionMatchedByStrategyRef bool            `json:"partitionMatchedByStrategyRef,omitempty"` // PartitionMatchedByStrategyRef returns the breakdown of matches by strategy for each selection if you ask for orders. Defaults to false if unspecified.
		CustomerStrategyRefs          []string        `json:"customerStrategyRefs,omitempty"`          // CustomerStrategyRefs restricts the results to orders matching any of the specified set of customer defined strategies.
		CurrencyCode                  string          `json:"currencyCode,omitempty"`                  // CurrencyCode is a Betfair standard currency code. If not specified, the default currency code is used.
		Locale                        string          `json:"locale,omitempty"`                        // Locale is the language used for the response. If not specified, the default is returned.
		MatchedSince                  time.Time       `json:"matchedSince,omitempty"`                  // MatchedSince restricts the results to orders that have at least one fragment matched since the specified date.
		BetIDs                        []string        `json:"betIds,omitempty"`                        // BetIDs restricts the results to orders with the specified bet IDs. A maximum of 250 betId's can be provided at a time.
	}
)

// ListRunnerBookResult is the dynamic data in a market (MarketBook).
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#MarketBook
type ListRunnerBookResult = MarketBook

// ListRunnerBook returns a list of dynamic data about a market and a specified runner. Dynamic data includes prices, the status of the market, the status of selections, the traded volume, and the status of any orders you have placed in the market.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687847/listRunnerBook
func (client *Client) ListRunnerBook(ctx context.Context, params ListRunnerBookParams) ([]ListRunnerBookResult, error) {
	json := []ListRunnerBookResult{}

	return json, client.GetSports(ctx, "listRunnerBook", params, &json)
}
