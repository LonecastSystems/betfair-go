package betfair

import (
	"time"
)

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

	CancelInstructionReport struct {
		Status        InstructionReportStatus    `json:"status"`
		ErrorCode     InstructionReportErrorCode `json:"errorCode,omitempty"`
		Instruction   CancelInstruction          `json:"instruction,omitempty"`
		SizeCancelled float64                    `json:"sizeCancelled"`
		CanceledDate  time.Time                  `json:"canceledDate,omitempty"`
	}
)

func (client *Client) CancelOrders(params CancelOrdersParams) (CancelExecutionReport, error) {
	json := CancelExecutionReport{}

	if err := client.GetSports("cancelOrders", params, &json); err != nil {
		return CancelExecutionReport{}, err
	}

	return json, nil
}
