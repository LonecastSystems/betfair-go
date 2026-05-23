package betfair

import "context"

type (
	// ListMarketProfitAndLossParams are the request parameters for listMarketProfitAndLoss.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687667/listMarketProfitAndLoss
	ListMarketProfitAndLossParams struct {
		MarketIDs          []string `json:"marketIds,omitempty"`          // MarketIDs is the list of markets to calculate profit and loss.
		IncludeSettledBets bool     `json:"includeSettledBets,omitempty"` // IncludeSettledBets is the option to include settled bets (partially settled markets only). Defaults to false if not specified.
		IncludeBspBets     bool     `json:"includeBspBets,omitempty"`     // IncludeBspBets is the option to include BSP bets. Defaults to false if not specified.
		NetOfCommission    bool     `json:"netOfCommission,omitempty"`    // NetOfCommission is the option to return profit and loss net of users current commission rate for this market including any special tariffs. Defaults to false if not specified.
	}

	// MarketProfitAndLoss holds profit and loss for a market.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#MarketProfitAndLoss
	MarketProfitAndLoss struct {
		MarketID          string                `json:"marketId,omitempty"`          // MarketID is the unique identifier for the market.
		CommissionApplied float64               `json:"commissionApplied,omitempty"` // CommissionApplied is the commission for this market.
		ProfitAndLosses   []RunnerProfitAndLoss `json:"profitAndLosses,omitempty"`   // ProfitAndLosses is the list of profit and loss per runner.
	}

	// RunnerProfitAndLoss holds profit and loss for a runner.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#RunnerProfitAndLoss
	RunnerProfitAndLoss struct {
		SelectionID int64   `json:"selectionId,omitempty"` // SelectionID is the selectionId.
		IfWin       float64 `json:"ifWin,omitempty"`       // IfWin is the profit or loss incurred should the selection be a winner.
		IfLose      float64 `json:"ifLose,omitempty"`      // IfLose is the profit or loss incurred should the selection be a loser.
		IfPlace     float64 `json:"ifPlace,omitempty"`     // IfPlace is the profit or loss incurred should the selection be placed (EachWay markets only).
	}
)

// ListMarketProfitAndLossResult is profit and loss for a market.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#MarketProfitAndLoss
type ListMarketProfitAndLossResult = MarketProfitAndLoss

// ListMarketProfitAndLoss retrieves profit and loss for a given list of OPEN markets. The values are calculated using matched bets and optionally settled bets. Only odds (MarketBettingType= ODDS) markets are implemented; markets of other types are silently ignored. To retrieve your profit and loss for CLOSED markets, please use the listClearedOrders request. Please note: Market Data Request Limits apply to requests made to listMarketProfitAndLoss.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687667/listMarketProfitAndLoss
func (client *Client) ListMarketProfitAndLoss(ctx context.Context, params ListMarketProfitAndLossParams) ([]ListMarketProfitAndLossResult, error) {
	json := []ListMarketProfitAndLossResult{}

	return json, client.GetSports(ctx, "listMarketProfitAndLoss", params, &json)
}
