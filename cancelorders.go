package betfair

import "context"

type (
	CancelOrdersParams struct {
		MarketID     string              `json:"marketId,omitempty"`
		Instructions []CancelInstruction `json:"instructions,omitempty"`
		CustomerRef  string              `json:"customerRef,omitempty"`
	}

	CancelInstruction struct {
		BetID         string  `json:"betId,omitempty"`
		SizeReduction float64 `json:"sizeReduction,omitempty"`
	}

	CancelOrdersReport struct {
		CustomerRef        string                          `json:"customerRef,omitempty"`
		Status             ExecutionReportStatus           `json:"status"`
		ErrorCode          ExecutionReportErrorCode        `json:"errorCode,omitempty"`
		MarketID           string                          `json:"marketId,omitempty"`
		InstructionReports []CancelOrdersInstructionReport `json:"instructionReports,omitempty"`
	}
)

func (client *Client) CancelOrders(ctx context.Context, params CancelOrdersParams) (CancelOrdersReport, error) {
	json := CancelOrdersReport{}

	return json, client.GetSports(ctx, "cancelOrders", params, &json)
}
