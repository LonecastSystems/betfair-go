package betfair

type TimeGranularity string

const (
	TG_DAYS    TimeGranularity = "DAYS"
	TG_HOURS   TimeGranularity = "HOURS"
	TG_MINUTES TimeGranularity = "MINUTES"
)

type (
	TimeRangesParams struct {
		Filter      MarketFilter    `json:"filter"`
		Granularity TimeGranularity `json:"granularity"`
	}

	TimeRangeResult struct {
		TimeRange   TimeRange `json:"timeRange,omitempty"`
		MarketCount int       `json:"marketCount,omitempty"`
	}
)

func (client *Client) ListTimeRanges(params TimeRangesParams) ([]TimeRangeResult, error) {
	json := []TimeRangeResult{}

	return json, client.GetSports("listTimeRanges", params, &json)
}
