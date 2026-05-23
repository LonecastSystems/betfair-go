package betfair

import "context"

type (
	ListCompetitionsParams struct {
		Filter MarketFilter `json:"filter"`
		Locale string       `json:"locale,omitempty"`
	}

	ListCompetitionsResult struct {
		Competition       Competition `json:"competition"`
		MarketCount       int         `json:"marketCount"`
		CompetitionRegion string      `json:"competitionRegion"`
	}
)

func (client *Client) ListCompetitions(ctx context.Context, params ListCompetitionsParams) ([]ListCompetitionsResult, error) {
	json := []ListCompetitionsResult{}

	return json, client.GetSports(ctx, "listCompetitions", params, &json)
}
