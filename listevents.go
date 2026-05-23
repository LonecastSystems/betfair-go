package betfair

import "context"

type (
	ListEventsParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	ListEventsResult struct {
		Event       Event `json:"event"`
		MarketCount int   `json:"marketCount"`
	}
)

func (client *Client) ListEvents(ctx context.Context, params ListEventsParams) ([]ListEventsResult, error) {
	json := []ListEventsResult{}

	return json, client.GetSports(ctx, "listEvents", params, &json)
}
