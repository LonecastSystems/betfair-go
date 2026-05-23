package betfair

import "context"

type (
	// GetAccountFundsParams are the parameters for getAccountFunds.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687888/getAccountFunds
	GetAccountFundsParams struct {
		Wallet Wallet `json:"wallet,omitempty"` // Name of the wallet in question. Global wallet is returned by default.
	}

	// GetAccountFundsResponse is the response for retrieving available to bet.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687852/Accounts+TypeDefinitions#AccountFundsResponse
	GetAccountFundsResponse struct {
		AvailableToBetBalance float64 `json:"availableToBetBalance"` // Amount available to bet.
		Exposure              float64 `json:"exposure"`              // Current exposure.
		RetainedCommission    float64 `json:"retainedCommission"`    // Sum of retained commission.
		ExposureLimit         float64 `json:"exposureLimit"`         // Exposure limit.
		DiscountRate          float64 `json:"discountRate"`          // User Discount Rate. Please note: Betfair AUS/NZ customers should not rely on this to determine their discount rates which are now applied at the account level.
		PointsBalance         int     `json:"pointsBalance"`         // The Betfair points balance.
		Wallet                Wallet  `json:"wallet"`                // Name of the wallet in question.
	}
)

// GetAccountFunds returns the available to bet amount, exposure and commission information.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687888/getAccountFunds
func (client *Client) GetAccountFunds(ctx context.Context, params GetAccountFundsParams) (GetAccountFundsResponse, error) {
	json := GetAccountFundsResponse{}

	return json, client.GetAccounts(ctx, "getAccountFunds", params, &json)
}
