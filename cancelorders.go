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

	CancelExecutionReport struct {
		CustomerRef        string                    `json:"customerRef,omitempty"`
		Status             ExecutionReportStatus     `json:"status"`
		ErrorCode          ExecutionReportErrorCode  `json:"errorCode,omitempty"`
		MarketID           string                    `json:"marketId,omitempty"`
		InstructionReports []CancelInstructionReport `json:"instructionReports,omitempty"`
	}
)

func (client *Client) CancelOrders(ctx context.Context, params CancelOrdersParams) (CancelExecutionReport, error) {
	json := CancelExecutionReport{}

	return json, client.GetSports(ctx, "cancelOrders", params, &json)
}
