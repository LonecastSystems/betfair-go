package stateless

type (
	CurrencyRateParams struct {
		FromCurrency string `json:"fromCurrency,omitempty"`
	}

	CurrencyRate struct {
		CurrencyCode string  `json:"currencyCode"`
		Rate         float64 `json:"rate"`
	}
)

func (client *StatelessClient) ListCurrencyRates(params CurrencyRateParams) ([]CurrencyRate, error) {
	json := []CurrencyRate{}

	if err := GetAccounts(client, "listCurrencyRates", params, &json); err != nil {
		return []CurrencyRate{}, err
	}

	return json, nil
}
