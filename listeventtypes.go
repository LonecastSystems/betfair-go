package betfair

import "context"

type (
	// ListEventTypesParams are the request parameters for listEventTypes.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687448/listEventTypes
	ListEventTypesParams struct {
		Filter MarketFilter `json:"filter"`           // Filter selects desired markets. All markets that match the criteria in the filter are selected.
		Locale string       `json:"locale,omitempty"` // Locale is the language used for the response. If not specified, the default is returned.
	}

	// ListEventTypesResult holds an Event Type (EventTypeResult).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#EventTypeResult
	ListEventTypesResult struct {
		EventType   EventType `json:"eventType,omitempty"`   // EventType is the Event Type (i.e. Sport).
		MarketCount int       `json:"marketCount,omitempty"` // MarketCount is the number of markets associated with this Event Type.
	}
)

// ListEventTypes returns a list of Event Types (i.e. Sports) associated with the markets selected by the MarketFilter.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687448/listEventTypes
func (client *Client) ListEventTypes(ctx context.Context, params ListEventTypesParams) ([]ListEventTypesResult, error) {
	json := []ListEventTypesResult{}

	return json, client.GetSports(ctx, "listEventTypes", params, &json)
}
