package betfairgo

type (
	MarketProfitAndLossParams struct {
		MarketIDs          []string `json:"marketIds"`
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

func (client *BetfairClient) ListMarketProfitAndLoss(params MarketProfitAndLossParams) ([]MarketProfitAndLoss, error) {
	json := []MarketProfitAndLoss{}

	if err := client.GetSports("listMarketProfitAndLoss", params, &json); err != nil {
		return []MarketProfitAndLoss{}, err
	}

	return json, nil
}
