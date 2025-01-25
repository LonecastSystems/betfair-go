package stateless

import (
	"time"
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

	MarketVersion struct {
		Version int64 `json:"version,omitempty"`
	}
)

func (client *StatelessClient) ListCurrentOrders(params CurrentOrdersParams) (CurrentOrderSummaryReport, error) {
	json := CurrentOrderSummaryReport{}

	if err := GetSports(client, "listCurrentOrders", params, &json); err != nil {
		return CurrentOrderSummaryReport{}, err
	}

	return json, nil
}
