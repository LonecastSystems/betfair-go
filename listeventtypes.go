package betfair

import "context"

type (
	ListEventTypesParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	ListEventTypesResult struct {
		EventType   EventType `json:"eventType,omitempty"`
		MarketCount int       `json:"marketCount,omitempty"`
	}
)

func (client *Client) ListEventTypes(ctx context.Context, params ListEventTypesParams) ([]ListEventTypesResult, error) {
	json := []ListEventTypesResult{}

	return json, client.GetSports(ctx, "listEventTypes", params, &json)
}
