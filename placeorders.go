package betfair

import "context"

// BetTargetType is an optional field to allow betting to a targeted PAYOUT or BACKERS_PROFIT.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#BetTargetType
type BetTargetType string

const (
	BTT_BACKER_PROFIT BetTargetType = "BACKERS_PROFIT" // BTT_BACKER_PROFIT is the payout requested minus the calculated size at which this LimitOrder is to be placed. BetTargetType bets are invalid for LINE markets.
	BTT_PAYOUT        BetTargetType = "PAYOUT"         // BTT_PAYOUT is the total payout requested on a LimitOrder.
)

// TimeInForce is the type of TimeInForce value to use. This value takes precedence over any PersistenceType value chosen.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#TimeInForce
type TimeInForce string

const (
	TIF_FILL_OR_KILL TimeInForce = "FILL_OR_KILL" // TIF_FILL_OR_KILL executes the transaction immediately and completely (filled to size or between minFillSize and size) or not at all (cancelled).
)

type (
	// PlaceOrdersParams are the request parameters for placeOrders.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687496/placeOrders
	PlaceOrdersParams struct {
		MarketID            string             `json:"marketId"`                      // MarketID is the market id these orders are to be placed on.
		Instructions        []PlaceInstruction `json:"instructions"`                  // Instructions are the place instructions. The limit of place instructions per request is 200 for the Global Exchange and 50 for the Italian Exchange.
		CustomerRef         string             `json:"customerRef,omitempty"`         // CustomerRef is an optional parameter allowing the client to pass a unique string (up to 32 chars) that is used to de-dupe mistaken re-submissions.
		MarketVersion       MarketVersion      `json:"marketVersion,omitempty"`       // MarketVersion is an optional parameter allowing the client to specify which version of the market the orders should be placed on.
		CustomerStrategyRef string             `json:"customerStrategyRef,omitempty"` // CustomerStrategyRef is an optional reference customers can use to specify which strategy has sent the order. The string is limited to 15 characters.
		Async               bool               `json:"async,omitempty"`               // Async is an optional flag which specifies if the orders should be placed asynchronously.
	}

	// PlaceInstruction is an instruction to place a bet.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#PlaceInstruction
	PlaceInstruction struct {
		OrderType          OrderType          `json:"orderType"`                    // OrderType is the type of order.
		SelectionID        int64              `json:"selectionId"`                  // SelectionID is the selection id.
		Handicap           float64            `json:"handicap,omitempty"`           // Handicap is the handicap associated with the runner in case of Asian handicap market.
		Side               Side               `json:"side"`                         // Side is the side of the bet.
		LimitOrder         LimitOrder         `json:"limitOrder,omitempty"`         // LimitOrder is the details of a limit order.
		LimitOnCloseOrder  LimitOnCloseOrder  `json:"limitOnCloseOrder,omitempty"`  // LimitOnCloseOrder is the details of a limit on close order.
		MarketOnCloseOrder MarketOnCloseOrder `json:"marketOnCloseOrder,omitempty"` // MarketOnCloseOrder is the details of a market on close order.
		CustomerOrderRef   string             `json:"customerOrderRef,omitempty"`   // CustomerOrderRef is an optional reference customers can use to specify which order has sent the order.
	}

	// LimitOrder is the details of a limit order.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#LimitOrder
	LimitOrder struct {
		Size            float64         `json:"size"`                    // Size is the bet size.
		Price           float64         `json:"price"`                   // Price is the bet price.
		PersistenceType PersistenceType `json:"persistenceType"`         // PersistenceType is what to do with the order at turn in play.
		TimeInForce     TimeInForce     `json:"timeInForce,omitempty"`   // TimeInForce is the type of TimeInForce value to use.
		MinFillSize     float64         `json:"minFillSize,omitempty"`   // MinFillSize is the minimum fill size for FILL_OR_KILL orders.
		BetTargetType   BetTargetType   `json:"betTargetType,omitempty"` // BetTargetType is an optional field to allow betting to a targeted PAYOUT or BACKERS_PROFIT.
		BetTargetSize   float64         `json:"betTargetSize,omitempty"` // BetTargetSize is the value of the payout or profit.
	}

	// LimitOnCloseOrder is the details of a limit on close order.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#LimitOnCloseOrder
	LimitOnCloseOrder struct {
		Liability float64 `json:"liability"` // Liability is the size of the bet.
		Price     float64 `json:"price"`     // Price is the limit price of the bet.
	}

	// MarketOnCloseOrder is the details of a market on close order.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#MarketOnCloseOrder
	MarketOnCloseOrder struct {
		Liability float64 `json:"liability"` // Liability is the size of the bet.
	}

	// PlaceOrdersReport is the response from placeOrders (PlaceExecutionReport).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#PlaceExecutionReport
	PlaceOrdersReport struct {
		CustomerRef        string                         `json:"customerRef,omitempty"`        // CustomerRef is the customer reference passed in the request.
		Status             ExecutionReportStatus          `json:"status"`                       // Status is the status of the execution report.
		ErrorCode          ExecutionReportErrorCode       `json:"errorCode,omitempty"`          // ErrorCode is the error code if the operation failed.
		MarketID           string                         `json:"marketId,omitempty"`           // MarketID is the market id these orders were placed on.
		InstructionReports []PlaceOrdersInstructionReport `json:"instructionReports,omitempty"` // InstructionReports is the list of instruction reports.
	}
)

// PlaceOrders places new orders into market. Please note that additional bet sizing rules apply to bets placed into the Italian Exchange. In normal circumstances the placeOrders is an atomic operation. PLEASE NOTE: if the 'Best Execution' features is switched off, placeOrders can return 'PROCESSED_WITH_ERRORS' meaning that some bets can be rejected and other placed when submitted in the same PlaceInstruction.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687496/placeOrders
func (client *Client) PlaceOrders(ctx context.Context, params PlaceOrdersParams) (PlaceOrdersReport, error) {
	json := PlaceOrdersReport{}

	return json, client.GetSports(ctx, "placeOrders", params, &json)
}
