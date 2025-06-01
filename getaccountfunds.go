package betfair

type (
	AccountDetailsParams struct {
		Wallet Wallet `json:"wallet,omitempty"`
	}

	AccountFundsResponse struct {
		AvailableToBetBalance float64 `json:"availableToBetBalance"`
		Exposure              float64 `json:"exposure"`
		RetainedCommission    float64 `json:"retainedCommission"`
		ExposureLimit         float64 `json:"exposureLimit"`
		DiscountRate          float64 `json:"discountRate"`
		PointsBalance         int     `json:"pointsBalance"`
		Wallet                string  `json:"wallet"`
	}
)

func (client *Client) GetAccountFunds(params AccountDetailsParams) (AccountFundsResponse, error) {
	json := AccountFundsResponse{}

	return json, client.GetAccounts("getAccountFunds", params, &json)
}
