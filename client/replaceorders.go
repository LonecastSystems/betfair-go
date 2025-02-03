package client

type (
	ReplaceOrdersParams struct {
		MarketID     string               `json:"marketId"`
		Instructions []ReplaceInstruction `json:"instructions"`
	}

	ReplaceInstruction struct {
		BetID         string        `json:"betId"`
		NewPrice      float64       `json:"newPrice"`
		CustomerRef   string        `json:"customerRef,omitempty"`
		MarketVersion MarketVersion `json:"marketVersion,omitempty"`
		Async         bool          `json:"async,omitempty"`
	}

	ReplaceExecutionReport struct {
		CustomerRef        string                     `json:"customerRef,omitempty"`
		Status             ExecutionReportStatus      `json:"status"`
		ErrorCode          ExecutionReportErrorCode   `json:"errorCode,omitempty"`
		MarketID           string                     `json:"marketId,omitempty"`
		InstructionReports []ReplaceInstructionReport `json:"instructionReports,omitempty"`
	}

	ReplaceInstructionReport struct {
		Status                  InstructionReportStatus    `json:"status"`
		ErrorCode               InstructionReportErrorCode `json:"errorCode,omitempty"`
		CancelInstructionReport CancelInstructionReport    `json:"cancelInstructionReport,omitempty"`
		PlaceInstructionReport  PlaceInstructionReport     `json:"placeInstructionReport,omitempty"`
	}
)

func (client *BetfairClient) ReplaceOrders(params ReplaceOrdersParams) (ReplaceExecutionReport, error) {
	json := ReplaceExecutionReport{}

	if err := client.GetSports("replaceOrders", params, &json); err != nil {
		return ReplaceExecutionReport{}, err
	}

	return json, nil
}
