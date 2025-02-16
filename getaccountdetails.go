package betfair

type AccountDetailsResponse struct {
	CurrencyCode  string  `json:"currencyCode"`
	FirstName     string  `json:"firstName"`
	LastName      string  `json:"lastName"`
	LocaleCode    string  `json:"localeCode"`
	Region        string  `json:"region"`
	Timezone      string  `json:"timezone"`
	DiscountRate  float64 `json:"discountRate"`
	PointsBalance int     `json:"pointsBalance"`
	CountryCode   string  `json:"countryCode"`
}

func (client *Client) GetAccountDetails() (AccountDetailsResponse, error) {
	json := AccountDetailsResponse{}

	return json, client.GetAccounts("getAccountDetails", nil, &json)
}
