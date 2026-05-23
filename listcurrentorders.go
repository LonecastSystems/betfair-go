package betfair

import (
	"context"
	"time"
)

// SortDir specifies the direction the results will be sorted in. If no value is passed in, it defaults to EARLIEST_TO_LATEST.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#SortDir
type SortDir string

const (
	SD_EARLIEST_TO_LATEST SortDir = "EARLIEST_TO_LATEST" // SD_EARLIEST_TO_LATEST orders from earliest value to latest e.g. lowest betId is first in the results.
	SD_LATEST_TO_EARLIEST SortDir = "LATEST_TO_EARLIEST" // SD_LATEST_TO_EARLIEST orders from the latest value to the earliest e.g. highest betId is first in the results.
)

// OrderBy specifies how the results will be ordered. If no value is passed in, it defaults to BY_BET. Also acts as a filter such that only orders with a valid value in the field being ordered by will be returned (i.e. BY_VOID_TIME returns only voided orders, BY_SETTLED_TIME (applies to partially settled markets) returns only settled orders and BY_MATCH_TIME returns only orders with a matched date (voided, settled, matched orders)). Note that specifying an orderBy parameter defines the context of the date filter applied by the dateRange parameter (placed, matched, voided or settled date).
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#OrderBy
type OrderBy string

const (
	OB_BY_BET          OrderBy = "BY_BET"          // OB_BY_BET is deprecated; use BY_PLACE_TIME instead. Order by placed time, then bet id.
	OB_BY_MARKET       OrderBy = "BY_MARKET"       // OB_BY_MARKET orders by market id, then placed time, then bet id.
	OB_BY_MATCH_TIME   OrderBy = "BY_MATCH_TIME"   // OB_BY_MATCH_TIME orders by time of last matched fragment (if any), then placed time, then bet id. Filters out orders which have no matched date.
	OB_BY_PLACE_TIME   OrderBy = "BY_PLACE_TIME"   // OB_BY_PLACE_TIME orders by placed time, then bet id. This is an alias of to be deprecated BY_BET.
	OB_BY_SETTLED_TIME OrderBy = "BY_SETTLED_TIME" // OB_BY_SETTLED_TIME orders by time of last settled fragment (if any due to partial market settlement), then by last match time, then placed time, then bet id. Filters out orders which have not been settled.
	OB_BY_VOID_TIME    OrderBy = "BY_VOID_TIME"    // OB_BY_VOID_TIME orders by time of last voided fragment (if any), then by last match time, then placed time, then bet id. Filters out orders which have not been voided.
)

type (
	// ListCurrentOrdersParams are the request parameters for listCurrentOrders.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687504/listCurrentOrders
	ListCurrentOrdersParams struct {
		BetIDs                 []string        `json:"betIds,omitempty"`                 // BetIDs optionally restricts the results to the specified bet IDs. A maximum of 250 betId's, or a combination of 250 betId's & marketId's are permitted.
		MarketIDs              []string        `json:"marketIds,omitempty"`              // MarketIDs optionally restricts the results to the specified market IDs. A maximum of 250 marketId's, or a combination of 250 marketId's & betId's are permitted.
		OrderProjection        OrderProjection `json:"orderProjection,omitempty"`        // OrderProjection optionally restricts the results to the specified order status.
		CustomerOrderRefs      []string        `json:"customerOrderRefs,omitempty"`      // CustomerOrderRefs optionally restricts the results to the specified customer order references.
		CustomerStrategyRefs   []string        `json:"customerStrategyRefs,omitempty"`   // CustomerStrategyRefs optionally restricts the results to the specified customer strategy references.
		DateRange              TimeRange       `json:"dateRange,omitempty"`              // DateRange optionally restricts the results to be from/to the specified date. The dates used to filter on will change to placed, matched, voided or settled dates depending on the orderBy.
		OrderBy                OrderBy         `json:"orderBy,omitempty"`                // OrderBy specifies how the results will be ordered.
		SortDir                SortDir         `json:"sortDir,omitempty"`                // SortDir specifies the direction the results will be sorted in.
		FromRecord             int             `json:"fromRecord,omitempty"`             // FromRecord specifies the first record that will be returned. Records start at index zero, not at index one.
		RecordCount            int             `json:"recordCount,omitempty"`            // RecordCount specifies how many records will be returned from the index position 'fromRecord'. Note that there is a page size limit of 1000.
		IncludeItemDescription bool            `json:"includeItemDescription,omitempty"` // IncludeItemDescription if true then extra description parameters are included in the CurrentOrderSummaryReport.
		IncludeSourceID        bool            `json:"includeSourceId,omitempty"`        // IncludeSourceID specifies if sourceIdKey and sourceIdDescription should be returned. Defaults to false.
	}

	// ListCurrentOrdersReport is the response from listCurrentOrders (CurrentOrderSummaryReport).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#CurrentOrderSummaryReport
	ListCurrentOrdersReport struct {
		CurrentOrders []ListCurrentOrdersResult `json:"currentOrders"` // CurrentOrders is the list of current orders.
		MoreAvailable bool                      `json:"moreAvailable"` // MoreAvailable indicates whether more records are available.
	}

	// ListCurrentOrdersResult holds a current order (CurrentOrderSummary).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#CurrentOrderSummary
	ListCurrentOrdersResult struct {
		BetID                  string                 `json:"betId"`                            // BetID is the bet ID of this order.
		MarketID               string                 `json:"marketId"`                         // MarketID is the market ID these orders are to be placed on.
		SelectionID            int64                  `json:"selectionId"`                      // SelectionID is the selection ID these orders are to be placed on.
		Handicap               float64                `json:"handicap"`                         // Handicap is the handicap associated with the runner in case of Asian handicap market.
		PriceSize              PriceSize              `json:"priceSize"`                        // PriceSize is the price and size of this bet.
		BspLiability           float64                `json:"bspLiability"`                     // BspLiability is the BSP liability.
		Side                   Side                   `json:"side"`                             // Side is the side of the bet.
		OrderStatus            OrderStatus            `json:"status"`                           // OrderStatus is the status of the order.
		PersistenceType        PersistenceType        `json:"persistenceType"`                  // PersistenceType is what to do with the order at turn in play.
		OrderType              OrderType              `json:"orderType"`                        // OrderType is the type of the order.
		PlacedDate             time.Time              `json:"placedDate"`                       // PlacedDate is the date the order was placed.
		MatchedDate            time.Time              `json:"matchedDate"`                      // MatchedDate is the date the order was matched.
		AveragePriceMatched    float64                `json:"averagePriceMatched,omitempty"`    // AveragePriceMatched is the average price matched at.
		SizeMatched            float64                `json:"sizeMatched,omitempty"`            // SizeMatched is the amount matched.
		SizeRemaining          float64                `json:"sizeRemaining,omitempty"`          // SizeRemaining is the amount remaining.
		SizeLapsed             float64                `json:"sizeLapsed,omitempty"`             // SizeLapsed is the amount lapsed.
		SizeCancelled          float64                `json:"sizeCancelled,omitempty"`          // SizeCancelled is the amount cancelled.
		SizeVoided             float64                `json:"sizeVoided,omitempty"`             // SizeVoided is the amount voided.
		RegulatorAuthCode      string                 `json:"regulatorAuthCode,omitempty"`      // RegulatorAuthCode is the regulator auth code.
		RegulatorCode          string                 `json:"regulatorCode,omitempty"`          // RegulatorCode is the regulator code.
		CustomerOrderRef       string                 `json:"customerOrderRef,omitempty"`       // CustomerOrderRef is an optional reference customers can use to specify which order has sent the order.
		CustomerStrategyRef    string                 `json:"customerStrategyRef,omitempty"`    // CustomerStrategyRef is an optional reference customers can use to specify which strategy has sent the order.
		CurrentItemDescription CurrentItemDescription `json:"currentItemDescription,omitempty"` // CurrentItemDescription contains extra description parameters.
		SourceIdKey            string                 `json:"sourceIdKey,omitempty"`            // SourceIdKey is the source id key.
		SourceIdDescription    string                 `json:"sourceIdDescription,omitempty"`    // SourceIdDescription is the source id description.
	}

	// CurrentItemDescription contains extra description parameters for a current order.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#CurrentItemDescription
	CurrentItemDescription struct {
		MarketVersion MarketVersion `json:"marketVersion,omitempty"` // MarketVersion is the version of the market.
	}
)

// ListCurrentOrders returns a list of your current orders. Optionally you can filter and sort your current orders using the various parameters, setting none of the parameters will return all of your current orders up to a maximum of 1000 bets, ordered BY_BET and sorted EARLIEST_TO_LATEST. To retrieve more than 1000 orders, you need to make use of the fromRecord and recordCount parameters.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687504/listCurrentOrders
func (client *Client) ListCurrentOrders(ctx context.Context, params ListCurrentOrdersParams) (ListCurrentOrdersReport, error) {
	json := ListCurrentOrdersReport{}

	return json, client.GetSports(ctx, "listCurrentOrders", params, &json)
}
