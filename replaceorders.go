package betfair

import "context"

type (
	// ReplaceOrdersParams are the request parameters for replaceOrders.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687487/replaceOrders
	ReplaceOrdersParams struct {
		MarketID      string               `json:"marketId"`                // MarketID is the market id these orders are to be placed on.
		Instructions  []ReplaceInstruction `json:"instructions"`            // Instructions are the replace instructions. The limit of replace instructions per request is 60.
		CustomerRef   string               `json:"customerRef,omitempty"`   // CustomerRef is an optional parameter allowing the client to pass a unique string (up to 32 chars) that is used to de-dupe mistaken re-submissions.
		MarketVersion MarketVersion        `json:"marketVersion,omitempty"` // MarketVersion is an optional parameter allowing the client to specify which version of the market the orders should be placed on.
		Async         bool                 `json:"async,omitempty"`         // Async is an optional flag which specifies if the orders should be replaced asynchronously.
	}

	// ReplaceInstruction is an instruction to replace a bet.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#ReplaceInstruction
	ReplaceInstruction struct {
		BetID    string  `json:"betId"`    // BetID is the bet ID to replace.
		NewPrice float64 `json:"newPrice"` // NewPrice is the new price for the bet.
	}

	// ReplaceOrdersReport is the response from replaceOrders (ReplaceExecutionReport).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#ReplaceExecutionReport
	ReplaceOrdersReport struct {
		CustomerRef        string                           `json:"customerRef,omitempty"`        // CustomerRef is the customer reference passed in the request.
		Status             ExecutionReportStatus            `json:"status"`                       // Status is the status of the execution report.
		ErrorCode          ExecutionReportErrorCode         `json:"errorCode,omitempty"`          // ErrorCode is the error code if the operation failed.
		MarketID           string                           `json:"marketId,omitempty"`           // MarketID is the market id these orders were replaced on.
		InstructionReports []ReplaceOrdersInstructionReport `json:"instructionReports,omitempty"` // InstructionReports is the list of instruction reports.
	}

	// ReplaceOrdersInstructionReport is the report for a replace instruction.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#ReplaceInstructionReport
	ReplaceOrdersInstructionReport struct {
		Status                        InstructionReportStatus       `json:"status"`                            // Status is the status of the instruction report.
		ErrorCode                     InstructionReportErrorCode    `json:"errorCode,omitempty"`               // ErrorCode is the error code if the instruction failed.
		CancelOrdersInstructionReport CancelOrdersInstructionReport `json:"cancelInstructionReport,omitempty"` // CancelOrdersInstructionReport is the cancel instruction report.
		PlaceOrdersInstructionReport  PlaceOrdersInstructionReport  `json:"placeInstructionReport,omitempty"`  // PlaceOrdersInstructionReport is the place instruction report.
	}
)

// ReplaceOrders is logically a bulk cancel followed by a bulk place. The cancel is completed first then the new orders are placed. The new orders will be placed atomically in that they will all be placed or none will be placed. In the case where the new orders cannot be placed the cancellations will not be rolled back.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687487/replaceOrders
func (client *Client) ReplaceOrders(ctx context.Context, params ReplaceOrdersParams) (ReplaceOrdersReport, error) {
	json := ReplaceOrdersReport{}

	return json, client.GetSports(ctx, "replaceOrders", params, &json)
}
