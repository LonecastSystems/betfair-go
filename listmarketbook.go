package betfair

import (
	"time"
)

type RollupModel string

const (
	RM_STAKE  = "STAKE"
	RM_PAYOUT = "PAYOUT"
	RM_NONE   = "NONE"
)

type RunnerStatus string

const (
	RS_ACTIVE         RunnerStatus = "ACTIVE"
	RS_WINNER         RunnerStatus = "WINNER"
	RS_LOSER          RunnerStatus = "LOSER"
	RS_PLACED         RunnerStatus = "PLACED"
	RS_REMOVED_VACANT RunnerStatus = "REMOVED_VACANT"
	RS_REMOVED        RunnerStatus = "REMOVED"
	RS_HIDDEN         RunnerStatus = "HIDDEN"
)

type MarketStatus string

const (
	MS_INACTIVE  MarketStatus = "INACTIVE"
	MS_OPEN      MarketStatus = "OPEN"
	MS_SUSPENDED MarketStatus = "SUSPENDED"
	MS_CLOSED    MarketStatus = "CLOSED"
)

type (
	MarketBookParams struct {
		MarketIDs                     []string        `json:"marketIds"`
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

func (client *Client) ListMarketBook(params MarketBookParams) ([]MarketBook, error) {
	json := []MarketBook{}

	if err := client.GetSports("listMarketBook", params, &json); err != nil {
		return json, err
	}

	return json, nil
}
