package stateless

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

func (client *StatelessClient) ListVenues(params VenueParams) ([]TimeRangeResult, error) {
	json := []TimeRangeResult{}

	if err := GetSports(client, "listVenues", params, &json); err != nil {
		return []TimeRangeResult{}, err
	}

	return json, nil
}
