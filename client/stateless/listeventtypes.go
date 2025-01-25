package stateless

type (
	EventTypeParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale"`
	}

	EventTypeResult struct {
		EventType   EventType `json:"eventType,omitempty"`
		MarketCount int       `json:"marketCount,omitempty"`
	}

	EventType struct {
		ID   string `json:"id,omitempty"`
		Name string `json:"name,omitempty"`
	}
)

func (client *StatelessClient) ListEventTypes(params EventTypeParams) ([]EventTypeResult, error) {
	json := []EventTypeResult{}

	if err := GetSports(client, "listEventTypes", params, &json); err != nil {
		return []EventTypeResult{}, err
	}

	return json, nil
}
