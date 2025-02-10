package betfair

type (
	EventTypeParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale"`
	}

	EventTypeResult struct {
		EventType   EventType `json:"eventType,omitempty"`
		MarketCount int       `json:"marketCount,omitempty"`
	}
)

func (client *Client) ListEventTypes(params EventTypeParams) ([]EventTypeResult, error) {
	json := []EventTypeResult{}

	if err := client.GetSports("listEventTypes", params, &json); err != nil {
		return []EventTypeResult{}, err
	}

	return json, nil
}
