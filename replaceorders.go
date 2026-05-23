package betfair

import "context"

type (
	ReplaceOrdersParams struct {
		MarketID      string               `json:"marketId"`
		Instructions  []ReplaceInstruction `json:"instructions"`
		CustomerRef   string               `json:"customerRef,omitempty"`
		MarketVersion MarketVersion        `json:"marketVersion,omitempty"`
		Async         bool                 `json:"async,omitempty"`
	}

	ReplaceInstruction struct {
		BetID    string  `json:"betId"`
		NewPrice float64 `json:"newPrice"`
	}

	ReplaceOrdersReport struct {
		CustomerRef        string                           `json:"customerRef,omitempty"`
		Status             ExecutionReportStatus            `json:"status"`
		ErrorCode          ExecutionReportErrorCode         `json:"errorCode,omitempty"`
		MarketID           string                           `json:"marketId,omitempty"`
		InstructionReports []ReplaceOrdersInstructionReport `json:"instructionReports,omitempty"`
	}

	ReplaceOrdersInstructionReport struct {
		Status                        InstructionReportStatus       `json:"status"`
		ErrorCode                     InstructionReportErrorCode    `json:"errorCode,omitempty"`
		CancelOrdersInstructionReport CancelOrdersInstructionReport `json:"cancelInstructionReport,omitempty"`
		PlaceOrdersInstructionReport  PlaceOrdersInstructionReport  `json:"placeInstructionReport,omitempty"`
	}
)

func (client *Client) ReplaceOrders(ctx context.Context, params ReplaceOrdersParams) (ReplaceOrdersReport, error) {
	json := ReplaceOrdersReport{}

	return json, client.GetSports(ctx, "replaceOrders", params, &json)
}
