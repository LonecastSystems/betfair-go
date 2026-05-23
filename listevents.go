package betfair

import "context"

type (
	// ListEventsParams are the request parameters for listEvents.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687442/listEvents
	ListEventsParams struct {
		Filter MarketFilter `json:"filter"`           // Filter selects desired markets. All markets that match the criteria in the filter are selected.
		Locale string       `json:"locale,omitempty"` // Locale is the language used for the response. If not specified, the default is returned.
	}

	// ListEventsResult holds an Event (EventResult).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#EventResult
	ListEventsResult struct {
		Event       Event `json:"event"`       // Event is the event associated with the markets.
		MarketCount int   `json:"marketCount"` // MarketCount is the number of markets associated with this Event.
	}
)

// ListEvents returns a list of Events (i.e, Reading vs. Man United) associated with the markets selected by the MarketFilter.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687442/listEvents
func (client *Client) ListEvents(ctx context.Context, params ListEventsParams) ([]ListEventsResult, error) {
	json := []ListEventsResult{}

	return json, client.GetSports(ctx, "listEvents", params, &json)
}
