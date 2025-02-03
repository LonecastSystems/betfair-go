package betfairgo

type (
	CompetitionParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale"`
	}

	CompetitionResult struct {
		Competition       Competition `json:"competition"`
		MarketCount       int         `json:"marketCount"`
		CompetitionRegion string      `json:"competitionRegion"`
	}

	Competition struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)

func (client *BetfairClient) ListCompetitions(params CompetitionParams) ([]CompetitionResult, error) {
	json := []CompetitionResult{}

	if err := client.GetSports("listCompetitions", params, &json); err != nil {
		return []CompetitionResult{}, err
	}

	return json, nil
}
