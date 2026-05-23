package betfair

import "context"

type (
	ListCurrencyRatesParams struct {
		FromCurrency string `json:"fromCurrency,omitempty"`
	}

	CurrencyRate struct {
		CurrencyCode string  `json:"currencyCode"`
		Rate         float64 `json:"rate"`
	}
)

type ListCurrencyRatesResult = CurrencyRate

func (client *Client) ListCurrencyRates(ctx context.Context, params ListCurrencyRatesParams) ([]ListCurrencyRatesResult, error) {
	json := []ListCurrencyRatesResult{}

	return json, client.GetAccounts(ctx, "listCurrencyRates", params, &json)
}
