package stateless

import (
	"time"
)

type (
	ClearedOrdersParams struct {
		BetStatus              BetStatus `json:"betStatus"`
		EventTypeIDs           []string  `json:"eventTypeIds,omitempty"`
		EventIDs               []string  `json:"eventIds,omitempty"`
		MarketIDs              []string  `json:"marketIds,omitempty"`
		RunnerIDs              []string  `json:"runnerIds,omitempty"`
		BetIDs                 []string  `json:"betIds,omitempty"`
		CustomerOrderRefs      []string  `json:"customerOrderRefs,omitempty"`
		CustomerStrategyRefs   []string  `json:"customerStrategyRefs,omitempty"`
		Side                   Side      `json:"side,omitempty"`
		SettledDateRange       TimeRange `json:"settledDateRange,omitempty"`
		GroupBy                GroupBy   `json:"groupBy,omitempty"`
		IncludeItemDescription bool      `json:"includeItemDescription,omitempty"`
		Locale                 string    `json:"locale,omitempty"`
		FromRecord             int       `json:"fromRecord,omitempty"`
		RecordCount            int       `json:"recordCount,omitempty"`
		IncludeSourceID        bool      `json:"includeSourceId,omitempty"`
	}

	ClearedOrderSummaryReport struct {
		ClearedOrders []ClearedOrderSummary `json:"clearedOrders"`
		MoreAvailable bool                  `json:"moreAvailable"`
	}

	ItemDescription struct {
		EventTypeDesc   string    `json:"eventTypeDesc,omitempty"`
		EventDesc       string    `json:"eventDesc,omitempty"`
		MarketDesc      string    `json:"marketDesc,omitempty"`
		MarketType      string    `json:"marketType,omitempty"`
		MarketStartTime time.Time `json:"marketStartTime,omitempty"`
		RunnerDesc      string    `json:"runnerDesc,omitempty"`
		NumberOfWinners int       `json:"numberOfWinners,omitempty"`
		EachWayDivisor  float64   `json:"eachWayDivisor,omitempty"`
	}

	ClearedOrderSummary struct {
		EventTypeID         string          `json:"eventTypeId,omitempty"`
		EventID             string          `json:"eventId,omitempty"`
		MarketID            string          `json:"marketId,omitempty"`
		SelectionID         int64           `json:"selectionId,omitempty"`
		Handicap            float64         `json:"handicap,omitempty"`
		BetID               string          `json:"betId,omitempty"`
		PlacedDate          time.Time       `json:"placedDate,omitempty"`
		PersistenceType     PersistenceType `json:"persistenceType,omitempty"`
		OrderType           OrderType       `json:"orderType,omitempty"`
		Side                Side            `json:"side,omitempty"`
		ItemDescription     ItemDescription `json:"itemDescription,omitempty"`
		PriceRequested      float64         `json:"priceRequested,omitempty"`
		SettledDate         time.Time       `json:"settledDate,omitempty"`
		LastMatchedDate     time.Time       `json:"lastMatchedDate,omitempty"`
		BetCount            int             `json:"betCount,omitempty"`
		Commission          float64         `json:"commission,omitempty"`
		PriceMatched        float64         `json:"priceMatched,omitempty"`
		PriceReduced        bool            `json:"priceReduced,omitempty"`
		SizeSettled         float64         `json:"sizeSettled,omitempty"`
		Profit              float64         `json:"profit,omitempty"`
		SizeCancelled       float64         `json:"sizeCancelled,omitempty"`
		CustomerOrderRef    string          `json:"customerOrderRef,omitempty"`
		CustomerStrategyRef string          `json:"customerStrategyRef,omitempty"`
		SourceIDKey         string          `json:"sourceIdKey,omitempty"`
		SourceIDDescription string          `json:"sourceIdDescription,omitempty"`
	}
)

func (client *StatelessClient) ListClearedOrders(params ClearedOrdersParams) (ClearedOrderSummaryReport, error) {
	json := ClearedOrderSummaryReport{}

	if err := GetSports(client, "listClearedOrders", params, &json); err != nil {
		return ClearedOrderSummaryReport{}, err
	}

	return json, nil
}
