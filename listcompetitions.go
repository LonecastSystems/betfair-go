package betfair

import "context"

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

func (client *Client) ListCompetitions(ctx context.Context, params CompetitionParams) ([]CompetitionResult, error) {
	json := []CompetitionResult{}

	return json, client.GetSports(ctx, "listCompetitions", params, &json)
}
