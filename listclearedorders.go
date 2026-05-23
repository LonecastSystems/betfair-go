package betfair

import (
	"context"
	"time"
)

// BetStatus restricts the results to the specified status.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#BetStatus
type BetStatus string

const (
	BS_SETTLED   BetStatus = "SETTLED"   // BS_SETTLED is a matched bet that was settled normally.
	BS_VOIDED    BetStatus = "VOIDED"    // BS_VOIDED is a matched bet that was subsequently voided by Betfair, before, during or after settlement.
	BS_LAPSED    BetStatus = "LAPSED"    // BS_LAPSED is an unmatched bet that was cancelled by Betfair (for example at turn in play).
	BS_CANCELLED BetStatus = "CANCELLED" // BS_CANCELLED is an unmatched bet that was cancelled by an explicit customer action.
)

// GroupBy specifies how to aggregate the lines, if not supplied then the lowest level is returned, i.e. bet by bet. This is only applicable to SETTLED BetStatus.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#GroupBy
type GroupBy string

const (
	GB_EVENT_TYPE GroupBy = "EVENT_TYPE" // GB_EVENT_TYPE is a roll up of settled P&L, commission paid and number of bet orders, on a specified event type.
	GB_EVENT      GroupBy = "EVENT"      // GB_EVENT is a roll up of settled P&L, commission paid and number of bet orders, on a specified event.
	GB_MARKET     GroupBy = "MARKET"     // GB_MARKET is a roll up of settled P&L, commission paid and number of bet orders, on a specified market.
	GB_SIDE       GroupBy = "SIDE"       // GB_SIDE is an averaged roll up of settled P&L, and number of bets, on the specified side of a specified selection within a specified market, that are either settled or voided.
	GB_BET        GroupBy = "BET"        // GB_BET is the P&L, side and regulatory information etc, about each individual bet order.
)

type (
	// ListClearedOrdersParams are the request parameters for listClearedOrders.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687749/listClearedOrders
	ListClearedOrdersParams struct {
		BetStatus              BetStatus `json:"betStatus"`                        // BetStatus restricts the results to the specified status.
		EventTypeIDs           []string  `json:"eventTypeIds,omitempty"`           // EventTypeIDs optionally restricts the results to the specified Event Type IDs.
		EventIDs               []string  `json:"eventIds,omitempty"`               // EventIDs optionally restricts the results to the specified Event IDs.
		MarketIDs              []string  `json:"marketIds,omitempty"`              // MarketIDs optionally restricts the results to the specified market IDs.
		RunnerIDs              []int64   `json:"runnerIds,omitempty"`              // RunnerIDs optionally restricts the results to the specified Runners.
		BetIDs                 []string  `json:"betIds,omitempty"`                 // BetIDs optionally restricts the results to the specified bet IDs. A maximum of 1000 betId's are allowed in a single request.
		CustomerOrderRefs      []string  `json:"customerOrderRefs,omitempty"`      // CustomerOrderRefs optionally restricts the results to the specified customer order references.
		CustomerStrategyRefs   []string  `json:"customerStrategyRefs,omitempty"`   // CustomerStrategyRefs optionally restricts the results to the specified customer strategy references.
		Side                   Side      `json:"side,omitempty"`                   // Side optionally restricts the results to the specified side.
		SettledDateRange       TimeRange `json:"settledDateRange,omitempty"`       // SettledDateRange optionally restricts the results to be from/to the specified settled date.
		GroupBy                GroupBy   `json:"groupBy,omitempty"`                // GroupBy specifies how to aggregate the lines. This is only applicable to SETTLED BetStatus.
		IncludeItemDescription bool      `json:"includeItemDescription,omitempty"` // IncludeItemDescription if true then an ItemDescription object is included in the response.
		Locale                 string    `json:"locale,omitempty"`                 // Locale is the language used for the itemDescription. If not specified, the customer account default is returned.
		FromRecord             int       `json:"fromRecord,omitempty"`             // FromRecord specifies the first record that will be returned. Records start at index zero.
		RecordCount            int       `json:"recordCount,omitempty"`            // RecordCount specifies how many records will be returned, from the index position 'fromRecord'. Note that there is a page size limit of 1000.
		IncludeSourceID        bool      `json:"includeSourceId,omitempty"`        // IncludeSourceID specifies if sourceIdKey and sourceIdDescription should be returned. Defaults to false.
	}

	// ListClearedOrdersReport is the response from listClearedOrders (ClearedOrderSummaryReport).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#ClearedOrderSummaryReport
	ListClearedOrdersReport struct {
		ClearedOrders []ListClearedOrdersResult `json:"clearedOrders"` // ClearedOrders is the list of cleared orders.
		MoreAvailable bool                      `json:"moreAvailable"` // MoreAvailable indicates whether more records are available.
	}

	// ItemDescription contains descriptive information about a cleared order item.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#ItemDescription
	ItemDescription struct {
		EventTypeDesc   string    `json:"eventTypeDesc,omitempty"`   // EventTypeDesc is the event type description.
		EventDesc       string    `json:"eventDesc,omitempty"`       // EventDesc is the event description.
		MarketDesc      string    `json:"marketDesc,omitempty"`      // MarketDesc is the market description.
		MarketType      string    `json:"marketType,omitempty"`      // MarketType is the market type.
		MarketStartTime time.Time `json:"marketStartTime,omitempty"` // MarketStartTime is the market start time.
		RunnerDesc      string    `json:"runnerDesc,omitempty"`      // RunnerDesc is the runner description.
		NumberOfWinners int       `json:"numberOfWinners,omitempty"` // NumberOfWinners is the number of winners.
		EachWayDivisor  float64   `json:"eachWayDivisor,omitempty"`  // EachWayDivisor is the each way divisor.
	}

	// ListClearedOrdersResult holds a cleared order (ClearedOrderSummary).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#ClearedOrderSummary
	ListClearedOrdersResult struct {
		EventTypeID         string          `json:"eventTypeId,omitempty"`         // EventTypeID is the event type ID.
		EventID             string          `json:"eventId,omitempty"`             // EventID is the event ID.
		MarketID            string          `json:"marketId,omitempty"`            // MarketID is the market ID.
		SelectionID         int64           `json:"selectionId,omitempty"`         // SelectionID is the selection ID.
		Handicap            float64         `json:"handicap,omitempty"`            // Handicap is the handicap associated with the runner in case of Asian handicap market.
		BetID               string          `json:"betId,omitempty"`               // BetID is the bet ID.
		PlacedDate          time.Time       `json:"placedDate,omitempty"`          // PlacedDate is the date the order was placed.
		PersistenceType     PersistenceType `json:"persistenceType,omitempty"`     // PersistenceType is what to do with the order at turn in play.
		OrderType           OrderType       `json:"orderType,omitempty"`           // OrderType is the type of the order.
		Side                Side            `json:"side,omitempty"`                // Side is the side of the bet.
		ItemDescription     ItemDescription `json:"itemDescription,omitempty"`     // ItemDescription contains descriptive information about the cleared order item.
		PriceRequested      float64         `json:"priceRequested,omitempty"`      // PriceRequested is the price requested.
		SettledDate         time.Time       `json:"settledDate,omitempty"`         // SettledDate is the date the order was settled.
		LastMatchedDate     time.Time       `json:"lastMatchedDate,omitempty"`     // LastMatchedDate is the date the order was last matched.
		BetCount            int             `json:"betCount,omitempty"`            // BetCount is the number of bets in this group.
		Commission          float64         `json:"commission,omitempty"`          // Commission is the commission paid.
		PriceMatched        float64         `json:"priceMatched,omitempty"`        // PriceMatched is the price matched at.
		PriceReduced        bool            `json:"priceReduced,omitempty"`        // PriceReduced indicates whether the price was reduced.
		SizeSettled         float64         `json:"sizeSettled,omitempty"`         // SizeSettled is the amount settled.
		Profit              float64         `json:"profit,omitempty"`              // Profit is the profit or loss incurred.
		SizeCancelled       float64         `json:"sizeCancelled,omitempty"`       // SizeCancelled is the amount cancelled.
		CustomerOrderRef    string          `json:"customerOrderRef,omitempty"`    // CustomerOrderRef is an optional reference customers can use to specify which order has sent the order.
		CustomerStrategyRef string          `json:"customerStrategyRef,omitempty"` // CustomerStrategyRef is an optional reference customers can use to specify which strategy has sent the order.
		SourceIDKey         string          `json:"sourceIdKey,omitempty"`         // SourceIDKey is the source id key.
		SourceIDDescription string          `json:"sourceIdDescription,omitempty"` // SourceIDDescription is the source id description.
	}
)

// ListClearedOrders returns a list of settled bets based on the bet status, ordered by settled date. To retrieve more than 1000 records, you need to make use of the fromRecord and recordCount parameters.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687749/listClearedOrders
func (client *Client) ListClearedOrders(ctx context.Context, params ListClearedOrdersParams) (ListClearedOrdersReport, error) {
	json := ListClearedOrdersReport{}

	return json, client.GetSports(ctx, "listClearedOrders", params, &json)
}
