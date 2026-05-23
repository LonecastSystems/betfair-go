package betfair

import "context"

type (
	// ListCompetitionsParams are the request parameters for listCompetitions.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687446/listCompetitions
	ListCompetitionsParams struct {
		Filter MarketFilter `json:"filter"`           // Filter selects desired markets. All markets that match the criteria in the filter are selected.
		Locale string       `json:"locale,omitempty"` // Locale is the language used for the response. If not specified, the default is returned.
	}

	// ListCompetitionsResult holds a Competition (CompetitionResult).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687465/Betting+Type+Definitions#CompetitionResult
	ListCompetitionsResult struct {
		Competition       Competition `json:"competition"`       // Competition is the competition associated with the markets.
		MarketCount       int         `json:"marketCount"`       // MarketCount is the number of markets associated with this Competition.
		CompetitionRegion string      `json:"competitionRegion"` // CompetitionRegion is the region in which this competition is taking place.
	}
)

// ListCompetitions returns a list of Competitions (i.e., World Cup 2013) associated with the markets selected by the MarketFilter. Please Note: for horse and greyhounds racing, please use listVenues
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687446/listCompetitions
func (client *Client) ListCompetitions(ctx context.Context, params ListCompetitionsParams) ([]ListCompetitionsResult, error) {
	json := []ListCompetitionsResult{}

	return json, client.GetSports(ctx, "listCompetitions", params, &json)
}
