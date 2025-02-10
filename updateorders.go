package betfair

type (
	UpdateOrdersParams struct {
		MarketID     string              `json:"marketId"`
		Instructions []UpdateInstruction `json:"instructions"`
		CustomerRef  string              `json:"customerRef,omitempty"`
	}

	UpdateInstruction struct {
		BetID              string          `json:"betId"`
		NewPersistenceType PersistenceType `json:"newPersistenceType,omitempty"`
	}

	UpdateExecutionReport struct {
		CustomerRef        string                    `json:"customerRef,omitempty"`
		Status             ExecutionReportStatus     `json:"status"`
		ErrorCode          ExecutionReportErrorCode  `json:"errorCode,omitempty"`
		MarketID           string                    `json:"marketId,omitempty"`
		InstructionReports []UpdateInstructionReport `json:"instructionReports,omitempty"`
	}

	UpdateInstructionReport struct {
		Status            InstructionReportStatus    `json:"status"`
		ErrorCode         InstructionReportErrorCode `json:"errorCode,omitempty"`
		UpdateInstruction UpdateInstruction          `json:"instruction"`
	}
)

func (client *Client) UpdateOrders(params UpdateOrdersParams) (UpdateExecutionReport, error) {
	json := UpdateExecutionReport{}

	if err := client.GetSports("updateOrders", params, &json); err != nil {
		return UpdateExecutionReport{}, err
	}

	return json, nil
}
