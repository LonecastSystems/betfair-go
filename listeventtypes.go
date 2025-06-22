package betfair

import "context"

type (
	EventTypeParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	EventTypeResult struct {
		EventType   EventType `json:"eventType,omitempty"`
		MarketCount int       `json:"marketCount,omitempty"`
	}
)

func (client *Client) ListEventTypes(ctx context.Context, params EventTypeParams) ([]EventTypeResult, error) {
	json := []EventTypeResult{}

	return json, client.GetSports(ctx, "listEventTypes", params, &json)
}
