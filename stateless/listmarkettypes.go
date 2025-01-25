package stateless

type (
	MarketTypeParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale"`
	}

	MarketTypeResult struct {
		MarketType  string `json:"marketType,omitempty"`
		MarketCount int    `json:"marketCount,omitempty"`
	}
)

func (client *StatelessClient) ListMarketTypes(params MarketTypeParams) ([]MarketTypeResult, error) {
	json := []MarketTypeResult{}

	if err := GetSports(client, "listMarketTypes", params, &json); err != nil {
		return []MarketTypeResult{}, err
	}

	return json, nil
}
