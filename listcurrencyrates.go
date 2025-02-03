package betfairgo

type (
	CurrencyRateParams struct {
		FromCurrency string `json:"fromCurrency,omitempty"`
	}

	CurrencyRate struct {
		CurrencyCode string  `json:"currencyCode"`
		Rate         float64 `json:"rate"`
	}
)

func (client *BetfairClient) ListCurrencyRates(params CurrencyRateParams) ([]CurrencyRate, error) {
	json := []CurrencyRate{}

	if err := client.GetAccounts("listCurrencyRates", params, &json); err != nil {
		return []CurrencyRate{}, err
	}

	return json, nil
}
