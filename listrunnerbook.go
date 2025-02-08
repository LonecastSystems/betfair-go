package betfair

import (
	"time"
)

type (
	RunnerBookParams struct {
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

func (client *Client) ListRunnerBook(params RunnerBookParams) ([]MarketBook, error) {
	json := []MarketBook{}

	if err := client.GetSports("listRunnerBook", params, &json); err != nil {
		return []MarketBook{}, err
	}

	return json, nil
}
