package betfair

type (
	VenueParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	VenueResult struct {
		Venue       string `json:"venue"`
		MarketCount int    `json:"marketCount"`
	}
)

func (client *Client) ListVenues(params VenueParams) ([]VenueResult, error) {
	json := []VenueResult{}

	return json, client.GetSports("listVenues", params, &json)
}
