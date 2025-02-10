package betfair

type (
	VenueParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale"`
	}

	VenueResult struct {
		Venue       string `json:"venue"`
		MarketCount int    `json:"marketCount"`
	}
)

func (client *Client) ListVenues(params VenueParams) ([]VenueResult, error) {
	json := []VenueResult{}

	if err := client.GetSports("listVenues", params, &json); err != nil {
		return []VenueResult{}, err
	}

	return json, nil
}
