package betfair

type (
	EventParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	EventResult struct {
		Event       Event `json:"event"`
		MarketCount int   `json:"marketCount"`
	}
)

func (client *Client) ListEvents(params EventParams) ([]EventResult, error) {
	json := []EventResult{}

	return json, client.GetSports("listEvents", params, &json)
}
