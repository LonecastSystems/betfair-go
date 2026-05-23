package betfair

import "context"

type (
	// ListCurrencyRatesParams are the parameters for listCurrencyRates.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687935/listCurrencyRates
	ListCurrencyRatesParams struct {
		FromCurrency string `json:"fromCurrency,omitempty"` // The currency from which the rates are computed. Please note: GBP is currently the only based currency support.
	}

	// CurrencyRate is a currency conversion rate.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687852/Accounts+TypeDefinitions#CurrencyRate
	CurrencyRate struct {
		CurrencyCode string  `json:"currencyCode"` // The currency code.
		Rate         float64 `json:"rate"`         // The conversion rate.
	}
)

type ListCurrencyRatesResult = CurrencyRate

// ListCurrencyRates returns a list of currency rates based on given currency. Please note: the currency rates are updated once every hour a few seconds after the hour.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687935/listCurrencyRates
func (client *Client) ListCurrencyRates(ctx context.Context, params ListCurrencyRatesParams) ([]ListCurrencyRatesResult, error) {
	json := []ListCurrencyRatesResult{}

	return json, client.GetAccounts(ctx, "listCurrencyRates", params, &json)
}
