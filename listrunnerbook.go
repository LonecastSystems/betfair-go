package betfair

import (
	"context"
	"time"
)

type (
	ListRunnerBookParams struct {
		MarketID                      string          `json:"marketId"`
		SelectionID                   int64           `json:"selectionId"`
		Handicap                      float64         `json:"handicap,omitempty"`
		PriceProjection               PriceProjection `json:"priceProjection,omitempty"`
		OrderProjection               OrderProjection `json:"orderProjection,omitempty"`
		MatchProjection               MatchProjection `json:"matchProjection,omitempty"`
		IncludeOverallPosition        bool            `json:"includeOverallPosition,omitempty"`
		PartitionMatchedByStrategyRef bool            `json:"partitionMatchedByStrategyRef,omitempty"`
		CustomerStrategyRefs          []string        `json:"customerStrategyRefs,omitempty"`
		CurrencyCode                  string          `json:"currencyCode,omitempty"`
		Locale                        string          `json:"locale,omitempty"`
		MatchedSince                  time.Time       `json:"matchedSince,omitempty"`
		BetIDs                        []string        `json:"betIds,omitempty"`
	}
)

type ListRunnerBookResult = MarketBook

func (client *Client) ListRunnerBook(ctx context.Context, params ListRunnerBookParams) ([]ListRunnerBookResult, error) {
	json := []ListRunnerBookResult{}

	return json, client.GetSports(ctx, "listRunnerBook", params, &json)
}
