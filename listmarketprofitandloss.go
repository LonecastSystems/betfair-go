package betfair

import "context"

type (
	ListMarketProfitAndLossParams struct {
		MarketIDs          []string `json:"marketIds,omitempty"`
		IncludeSettledBets bool     `json:"includeSettledBets,omitempty"`
		IncludeBspBets     bool     `json:"includeBspBets,omitempty"`
		NetOfCommission    bool     `json:"netOfCommission,omitempty"`
	}

	MarketProfitAndLoss struct {
		MarketID          string                `json:"marketId,omitempty"`
		CommissionApplied float64               `json:"commissionApplied,omitempty"`
		ProfitAndLosses   []RunnerProfitAndLoss `json:"profitAndLosses,omitempty"`
	}

	RunnerProfitAndLoss struct {
		SelectionID int64   `json:"selectionId,omitempty"`
		IfWin       float64 `json:"ifWin,omitempty"`
		IfLose      float64 `json:"ifLose,omitempty"`
		IfPlace     float64 `json:"ifPlace,omitempty"`
	}
)

type ListMarketProfitAndLossResult = MarketProfitAndLoss

func (client *Client) ListMarketProfitAndLoss(ctx context.Context, params ListMarketProfitAndLossParams) ([]ListMarketProfitAndLossResult, error) {
	json := []ListMarketProfitAndLossResult{}

	return json, client.GetSports(ctx, "listMarketProfitAndLoss", params, &json)
}
