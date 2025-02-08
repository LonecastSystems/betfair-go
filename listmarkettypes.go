package betfair

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

func (client *Client) ListMarketTypes(params MarketTypeParams) ([]MarketTypeResult, error) {
	json := []MarketTypeResult{}

	if err := client.GetSports("listMarketTypes", params, &json); err != nil {
		return []MarketTypeResult{}, err
	}

	return json, nil
}
