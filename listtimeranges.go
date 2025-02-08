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

	if err := client.GetSports("listTimeRanges", params, &json); err != nil {
		return []TimeRangeResult{}, err
	}

	return json, nil
}
