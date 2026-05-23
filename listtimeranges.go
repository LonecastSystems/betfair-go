package betfair

import "context"

// TimeGranularity is the granularity of time periods that correspond to markets selected by the market filter.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687455/Betting+Enums#TimeGranularity
type TimeGranularity string

const (
	TG_DAYS    TimeGranularity = "DAYS"    // TG_DAYS is DAYS.
	TG_HOURS   TimeGranularity = "HOURS"   // TG_HOURS is HOURS.
	TG_MINUTES TimeGranularity = "MINUTES" // TG_MINUTES is MINUTES.
)

type (
	// ListTimeRangesParams are the request parameters for listTimeRanges.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687444/listTimeRanges
	ListTimeRangesParams struct {
		Filter      MarketFilter    `json:"filter"`      // Filter selects desired markets. All markets that match the criteria in the filter are selected.
		Granularity TimeGranularity `json:"granularity"` // Granularity is the granularity of time periods that correspond to markets selected by the market filter.
	}

	// ListTimeRangesResult holds a time range (TimeRangeResult).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#TimeRangeResult
	ListTimeRangesResult struct {
		TimeRange   TimeRange `json:"timeRange,omitempty"`   // TimeRange is the time range associated with the markets.
		MarketCount int       `json:"marketCount,omitempty"` // MarketCount is the number of markets associated with this time range.
	}
)

// ListTimeRanges returns a list of time ranges in the granularity specified in the request (i.e. 3PM to 4PM, Aug 14th to Aug 15th) associated with the markets selected by the MarketFilter.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687444/listTimeRanges
func (client *Client) ListTimeRanges(ctx context.Context, params ListTimeRangesParams) ([]ListTimeRangesResult, error) {
	json := []ListTimeRangesResult{}

	return json, client.GetSports(ctx, "listTimeRanges", params, &json)
}
