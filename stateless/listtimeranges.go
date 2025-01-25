package stateless

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

func (client *StatelessClient) ListTimeRanges(params TimeRangesParams) ([]TimeRangeResult, error) {
	json := []TimeRangeResult{}

	if err := GetSports(client, "listTimeRanges", params, &json); err != nil {
		return []TimeRangeResult{}, err
	}

	return json, nil
}
