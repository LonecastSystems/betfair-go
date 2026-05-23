package betfair

import "context"

// GetAccountDetailsResponse is the response for retrieving account details.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687852/Accounts+TypeDefinitions#AccountDetailsResponse
type GetAccountDetailsResponse struct {
	CurrencyCode  string  `json:"currencyCode"`  // ISO currency code for the account.
	FirstName     string  `json:"firstName"`     // The first name of the user.
	LastName      string  `json:"lastName"`      // The last name of the user.
	LocaleCode    string  `json:"localeCode"`    // The locale code of the user.
	Region        string  `json:"region"`        // The region of the user.
	Timezone      string  `json:"timezone"`      // The timezone of the user.
	DiscountRate  float64 `json:"discountRate"`  // User Discount Rate.
	PointsBalance int     `json:"pointsBalance"` // The Betfair points balance.
	CountryCode   string  `json:"countryCode"`   // The ISO-2 country code for the account.
}

// GetAccountDetails returns the details relating your account, including your discount rate and Betfair point balance.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2699900/getAccountDetails
func (client *Client) GetAccountDetails(ctx context.Context) (GetAccountDetailsResponse, error) {
	json := GetAccountDetailsResponse{}

	return json, client.GetAccounts(ctx, "getAccountDetails", nil, &json)
}
