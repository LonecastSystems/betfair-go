package client

import "time"

type (
	EventParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale"`
	}

	EventResult struct {
		Event       Event `json:"event"`
		MarketCount int   `json:"marketCount"`
	}

	Event struct {
		ID          string    `json:"id,omitempty"`
		Name        string    `json:"name,omitempty"`
		CountryCode string    `json:"countryCode,omitempty"`
		Timezone    string    `json:"timezone,omitempty"`
		Venue       string    `json:"venue,omitempty"`
		OpenDate    time.Time `json:"openDate,omitempty"`
	}
)

func (client *BetfairClient) ListEvents(params EventParams) ([]EventResult, error) {
	json := []EventResult{}

	if err := client.GetSports("listEvents", params, &json); err != nil {
		return []EventResult{}, err
	}

	return json, nil
}
