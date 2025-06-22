package betfair

import (
	"context"
	"testing"
)

func TestListCompetitions(t *testing.T) {
	c := NewTestClient(t)

	params := CompetitionParams{Filter: MarketFilter{
		CompetitionIDs:  []string{"10932509"},
		EventTypeIDs:    []string{"1"},
		MarketCountries: []string{"GB"},
	}}

	competitions, err := c.ListCompetitions(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(competitions)
	if len != 1 {
		t.Fatal(len)
	}

	competitionName := competitions[0].Competition.Name
	if competitionName != "English Premier League" {
		t.Fatal(competitionName)
	}
}
