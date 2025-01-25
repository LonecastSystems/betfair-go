package stateless

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

func (client *StatelessClient) ListCompetitions(params CompetitionParams) ([]CompetitionResult, error) {
	json := []CompetitionResult{}

	if err := GetSports(client, "listCompetitions", params, &json); err != nil {
		return []CompetitionResult{}, err
	}

	return json, nil
}
