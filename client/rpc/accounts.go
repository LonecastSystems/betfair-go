package rpc

type (
	AccountDetailsParams struct {
		Wallet string `json:"wallet"`
	}

	WalletResult struct {
		AvailableToBetBalance float64 `json:"availableToBetBalance"`
		Exposure              float64 `json:"exposure"`
		RetainedCommission    float64 `json:"retainedCommission"`
		ExposureLimit         float64 `json:"exposureLimit"`
		DiscountRate          float64 `json:"discountRate"`
		PointsBalance         int     `json:"pointsBalance"`
		Wallet                string  `json:"wallet"`
	}
)

func (client *RpcClient) GetAccountFunds(params AccountDetailsParams) (WalletResult, error) {
	json := WalletResult{}

	if err := GetAccounts(client, 1, "getAccountFunds", params, &json); err != nil {
		return WalletResult{}, err
	}

	return json, nil
}
