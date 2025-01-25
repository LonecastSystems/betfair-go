package stateless

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

func (client *StatelessClient) GetAccountDetails() (AccountDetailsResponse, error) {
	json := AccountDetailsResponse{}

	if err := GetAccounts(client, "getAccountDetails", nil, &json); err != nil {
		return AccountDetailsResponse{}, err
	}

	return json, nil
}
