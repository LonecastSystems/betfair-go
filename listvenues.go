package betfair

import "context"

type (
	// ListVenuesParams are the request parameters for listVenues.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687521/listVenues
	ListVenuesParams struct {
		Filter MarketFilter `json:"filter"`           // Filter selects desired markets. All markets that match the criteria in the filter are selected.
		Locale string       `json:"locale,omitempty"` // Locale is the language used for the response. If not specified, the default is returned.
	}

	// ListVenuesResult holds a venue (VenueResult).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#VenueResult
	ListVenuesResult struct {
		Venue       string `json:"venue"`       // Venue is the venue associated with the markets (e.g. Cheltenham, Ascot).
		MarketCount int    `json:"marketCount"` // MarketCount is the number of markets associated with this Venue.
	}
)

// ListVenues returns a list of Venues (i.e. Cheltenham, Ascot) associated with the markets selected by the MarketFilter. Only Horse Racing & Greyhound markets are associated with a Venue.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687521/listVenues
func (client *Client) ListVenues(ctx context.Context, params ListVenuesParams) ([]ListVenuesResult, error) {
	json := []ListVenuesResult{}

	return json, client.GetSports(ctx, "listVenues", params, &json)
}
