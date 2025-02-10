package betfair

type (
	EventParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale"`
	}

	EventResult struct {
		Event       Event `json:"event"`
		MarketCount int   `json:"marketCount"`
	}
)

func (client *Client) ListEvents(params EventParams) ([]EventResult, error) {
	json := []EventResult{}

	if err := client.GetSports("listEvents", params, &json); err != nil {
		return []EventResult{}, err
	}

	return json, nil
}
