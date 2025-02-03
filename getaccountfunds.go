package betfairgo

type (
	AccountDetailsParams struct {
		Wallet Wallet `json:"wallet"`
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

func (client *BetfairClient) GetAccountFunds(params AccountDetailsParams) (AccountFundsResponse, error) {
	json := AccountFundsResponse{}

	if err := client.GetAccounts("getAccountFunds", params, &json); err != nil {
		return AccountFundsResponse{}, err
	}

	return json, nil
}
