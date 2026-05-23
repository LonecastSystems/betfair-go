package betfair

import "context"

type (
	// CancelOrdersParams are the request parameters for cancelOrders.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687491/cancelOrders
	CancelOrdersParams struct {
		MarketID     string              `json:"marketId,omitempty"`     // MarketID if marketId and betId aren't supplied all bets are cancelled.
		Instructions []CancelInstruction `json:"instructions,omitempty"` // Instructions are the cancel instructions. The limit of cancel instructions per request is 60.
		CustomerRef  string              `json:"customerRef,omitempty"`  // CustomerRef is an optional parameter allowing the client to pass a unique string (up to 32 chars) that is used to de-dupe mistaken re-submissions.
	}

	// CancelInstruction is an instruction to cancel a bet.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#CancelInstruction
	CancelInstruction struct {
		BetID         string  `json:"betId,omitempty"`         // BetID is the bet ID to cancel.
		SizeReduction float64 `json:"sizeReduction,omitempty"` // SizeReduction if supplied this is a partial cancel.
	}

	// CancelOrdersReport is the response from cancelOrders (CancelExecutionReport).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#CancelExecutionReport
	CancelOrdersReport struct {
		CustomerRef        string                          `json:"customerRef,omitempty"`        // CustomerRef is the customer reference passed in the request.
		Status             ExecutionReportStatus           `json:"status"`                       // Status is the status of the execution report.
		ErrorCode          ExecutionReportErrorCode        `json:"errorCode,omitempty"`          // ErrorCode is the error code if the operation failed.
		MarketID           string                          `json:"marketId,omitempty"`           // MarketID is the market id these orders were cancelled on.
		InstructionReports []CancelOrdersInstructionReport `json:"instructionReports,omitempty"` // InstructionReports is the list of instruction reports.
	}
)

// CancelOrders cancels all bets OR cancel all bets on a market OR fully or partially cancel particular orders on a market. Only LIMIT orders can be cancelled or partially cancelled once placed.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687491/cancelOrders
func (client *Client) CancelOrders(ctx context.Context, params CancelOrdersParams) (CancelOrdersReport, error) {
	json := CancelOrdersReport{}

	return json, client.GetSports(ctx, "cancelOrders", params, &json)
}
