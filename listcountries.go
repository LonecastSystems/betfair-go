package betfair

import "context"

type (
	// ListCountriesParams are the request parameters for listCountries.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687523/listCountries
	ListCountriesParams struct {
		Filter MarketFilter `json:"filter"`           // Filter selects desired markets. All markets that match the criteria in the filter are selected.
		Locale string       `json:"locale,omitempty"` // Locale is the language used for the response. If not specified, the default is returned.
	}

	// ListCountriesResult holds a country (CountryCodeResult).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#CountryCodeResult
	ListCountriesResult struct {
		CountryCode string `json:"countryCode"` // CountryCode is the ISO-2 code for the country.
		MarketCount int    `json:"marketCount"` // MarketCount is the number of markets associated with this Country.
	}
)

// ListCountries returns a list of Countries associated with the markets selected by the MarketFilter.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687523/listCountries
func (client *Client) ListCountries(ctx context.Context, params ListCountriesParams) ([]ListCountriesResult, error) {
	json := []ListCountriesResult{}

	return json, client.GetSports(ctx, "listCountries", params, &json)
}
