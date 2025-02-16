package betfair

type (
	CurrencyRateParams struct {
		FromCurrency string `json:"fromCurrency,omitempty"`
	}

	CurrencyRate struct {
		CurrencyCode string  `json:"currencyCode"`
		Rate         float64 `json:"rate"`
	}
)

func (client *Client) ListCurrencyRates(params CurrencyRateParams) ([]CurrencyRate, error) {
	json := []CurrencyRate{}

	return json, client.GetAccounts("listCurrencyRates", params, &json)
}
