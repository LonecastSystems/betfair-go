package betfair

import "context"

type (
	GetAccountFundsParams struct {
		Wallet Wallet `json:"wallet,omitempty"`
	}

	GetAccountFundsResponse struct {
		AvailableToBetBalance float64 `json:"availableToBetBalance"`
		Exposure              float64 `json:"exposure"`
		RetainedCommission    float64 `json:"retainedCommission"`
		ExposureLimit         float64 `json:"exposureLimit"`
		DiscountRate          float64 `json:"discountRate"`
		PointsBalance         int     `json:"pointsBalance"`
		Wallet                Wallet  `json:"wallet"`
	}
)

func (client *Client) GetAccountFunds(ctx context.Context, params GetAccountFundsParams) (GetAccountFundsResponse, error) {
	json := GetAccountFundsResponse{}

	return json, client.GetAccounts(ctx, "getAccountFunds", params, &json)
}
