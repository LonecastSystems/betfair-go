package betfair

import "context"

type (
	CurrencyRateParams struct {
		FromCurrency string `json:"fromCurrency,omitempty"`
	}

	CurrencyRate struct {
		CurrencyCode string  `json:"currencyCode"`
		Rate         float64 `json:"rate"`
	}
)

func (client *Client) ListCurrencyRates(ctx context.Context, params CurrencyRateParams) ([]CurrencyRate, error) {
	json := []CurrencyRate{}

	return json, client.GetAccounts(ctx, "listCurrencyRates", params, &json)
}
