package betfair

import "context"

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

func (client *Client) ListEvents(ctx context.Context, params EventParams) ([]EventResult, error) {
	json := []EventResult{}

	return json, client.GetSports(ctx, "listEvents", params, &json)
}
