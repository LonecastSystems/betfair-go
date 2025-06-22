package betfair

import (
	"context"
	"time"
)

type SortDir string

const (
	SD_EARLIEST_TO_LATEST SortDir = "EARLIEST_TO_LATEST"
	SD_LATEST_TO_EARLIEST SortDir = "LATEST_TO_EARLIEST"
)

type OrderBy string

const (
	OB_BY_BET          OrderBy = "BY_BET"
	OB_BY_MARKET       OrderBy = "BY_MARKET"
	OB_BY_MATCH_TIME   OrderBy = "BY_MATCH_TIME"
	OB_BY_PLACE_TIME   OrderBy = "BY_PLACE_TIME"
	OB_BY_SETTLED_TIME OrderBy = "BY_SETTLED_TIME"
	OB_BY_VOID_TIME    OrderBy = "BY_VOID_TIME"
)

type (
	CurrentOrdersParams struct {
		BetIDs                 []string        `json:"betIds,omitempty"`
		MarketIDs              []string        `json:"marketIds,omitempty"`
		OrderProjection        OrderProjection `json:"orderProjection,omitempty"`
		CustomerOrderRefs      []string        `json:"customerOrderRefs,omitempty"`
		CustomerStrategyRefs   []string        `json:"customerStrategyRefs,omitempty"`
		DateRange              TimeRange       `json:"dateRange,omitempty"`
		OrderBy                OrderBy         `json:"orderBy,omitempty"`
		SortDir                SortDir         `json:"sortDir,omitempty"`
		FromDate               time.Time       `json:"fromDate,omitempty"`
		FromRecord             int             `json:"fromRecord,omitempty"`
		RecordCount            int             `json:"recordCount,omitempty"`
		IncludeItemDescription bool            `json:"includeItemDescription,omitempty"`
		IncludeSourceID        bool            `json:"includeSourceId,omitempty"`
	}

	CurrentOrderSummaryReport struct {
		CurrentOrders []CurrentOrderSummary `json:"currentOrders"`
		MoreAvailable bool                  `json:"moreAvailable"`
	}

	CurrentOrderSummary struct {
		BetID                  string                 `json:"betId"`
		MarketID               string                 `json:"marketId"`
		SelectionID            int64                  `json:"selectionId"`
		Handicap               float64                `json:"handicap"`
		PriceSize              PriceSize              `json:"priceSize"`
		BspLiability           float64                `json:"bspLiability"`
		Side                   Side                   `json:"side"`
		OrderStatus            OrderStatus            `json:"status"`
		PersistenceType        PersistenceType        `json:"persistenceType"`
		OrderType              OrderType              `json:"orderType"`
		PlacedDate             time.Time              `json:"placedDate"`
		MatchedDate            time.Time              `json:"matchedDate"`
		AveragePriceMatched    float64                `json:"averagePriceMatched,omitempty"`
		SizeMatched            float64                `json:"sizeMatched,omitempty"`
		SizeRemaining          float64                `json:"sizeRemaining,omitempty"`
		SizeLapsed             float64                `json:"sizeLapsed,omitempty"`
		SizeCancelled          float64                `json:"sizeCancelled,omitempty"`
		SizeVoided             float64                `json:"sizeVoided,omitempty"`
		RegulatorAuthCode      string                 `json:"regulatorAuthCode,omitempty"`
		RegulatorCode          string                 `json:"regulatorCode,omitempty"`
		CustomerOrderRef       string                 `json:"customerOrderRef,omitempty"`
		CustomerStrategyRef    string                 `json:"customerStrategyRef,omitempty"`
		CurrentItemDescription CurrentItemDescription `json:"currentItemDescription,omitempty"`
		SourceIdKey            string                 `json:"sourceIdKey,omitempty"`
		SourceIdDescription    string                 `json:"sourceIdDescription,omitempty"`
	}

	CurrentItemDescription struct {
		MarketVersion MarketVersion `json:"marketVersion,omitempty"`
	}
)

func (client *Client) ListCurrentOrders(ctx context.Context, params CurrentOrdersParams) (CurrentOrderSummaryReport, error) {
	json := CurrentOrderSummaryReport{}

	return json, client.GetSports(ctx, "listCurrentOrders", params, &json)
}
