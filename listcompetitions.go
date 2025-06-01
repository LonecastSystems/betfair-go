package betfair

type (
	CompetitionParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	CompetitionResult struct {
		Competition       Competition `json:"competition"`
		MarketCount       int         `json:"marketCount"`
		CompetitionRegion string      `json:"competitionRegion"`
	}
)

func (client *Client) ListCompetitions(params CompetitionParams) ([]CompetitionResult, error) {
	json := []CompetitionResult{}

	return json, client.GetSports("listCompetitions", params, &json)
}
