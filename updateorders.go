package betfair

import "context"

type (
	// UpdateOrdersParams are the request parameters for updateOrders.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687485/updateOrders
	UpdateOrdersParams struct {
		MarketID     string              `json:"marketId"`              // MarketID is the market id these orders are to be placed on.
		Instructions []UpdateInstruction `json:"instructions"`          // Instructions are the update instructions. The limit of update instructions per request is 60.
		CustomerRef  string              `json:"customerRef,omitempty"` // CustomerRef is an optional parameter allowing the client to pass a unique string (up to 32 chars) that is used to de-dupe mistaken re-submissions.
	}

	// UpdateInstruction is an instruction to update a bet.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#UpdateInstruction
	UpdateInstruction struct {
		BetID              string          `json:"betId"`                        // BetID is the bet ID to update.
		NewPersistenceType PersistenceType `json:"newPersistenceType,omitempty"` // NewPersistenceType is the new persistence type for the bet.
	}

	// UpdateOrdersReport is the response from updateOrders (UpdateExecutionReport).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#UpdateExecutionReport
	UpdateOrdersReport struct {
		CustomerRef        string                          `json:"customerRef,omitempty"`        // CustomerRef is the customer reference passed in the request.
		Status             ExecutionReportStatus           `json:"status"`                       // Status is the status of the execution report.
		ErrorCode          ExecutionReportErrorCode        `json:"errorCode,omitempty"`          // ErrorCode is the error code if the operation failed.
		MarketID           string                          `json:"marketId,omitempty"`           // MarketID is the market id these orders were updated on.
		InstructionReports []UpdateOrdersInstructionReport `json:"instructionReports,omitempty"` // InstructionReports is the list of instruction reports.
	}

	// UpdateOrdersInstructionReport is the report for an update instruction.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#UpdateInstructionReport
	UpdateOrdersInstructionReport struct {
		Status            InstructionReportStatus    `json:"status"`              // Status is the status of the instruction report.
		ErrorCode         InstructionReportErrorCode `json:"errorCode,omitempty"` // ErrorCode is the error code if the instruction failed.
		UpdateInstruction UpdateInstruction          `json:"instruction"`         // UpdateInstruction is the update instruction that was processed.
	}
)

// UpdateOrders updates non-exposure changing fields.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687485/updateOrders
func (client *Client) UpdateOrders(ctx context.Context, params UpdateOrdersParams) (UpdateOrdersReport, error) {
	json := UpdateOrdersReport{}

	return json, client.GetSports(ctx, "updateOrders", params, &json)
}
