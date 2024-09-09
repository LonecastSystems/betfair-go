package rpc

import (
	"slices"
	"testing"

	"github.com/LonecastSystems/betfair-go/client/rpc"
)

var sessionKey, appKey = "", ""

func CreateClient() rpc.RpcBettingClient {
	return rpc.RpcBettingClient(rpc.CreateClient(sessionKey, appKey))
}

func TestEventTypes(t *testing.T) {
	if sessionKey == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	c := CreateClient()

	params := rpc.RPCParams{Filter: rpc.MarketFilter{
		MarketTypeCodes: []string{"OVER_UNDER_25"},
	}}

	eventTypes, err := c.ListEventTypes(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(eventTypes)
	if len != 1 {
		t.Fatal(len)
	}

	eventTypeName := eventTypes[0].EventType.Name
	if eventTypeName != "Soccer" {
		t.Fatal(eventTypeName)
	}
}

func TestListCompetitions(t *testing.T) {
	if sessionKey == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	c := CreateClient()

	params := rpc.RPCParams{Filter: rpc.MarketFilter{
		TextQuery:       "Premier League",
		EventTypeIds:    []string{"1"},
		MarketCountries: []string{"GB"},
	}}

	competitions, err := c.ListCompetitions(params)
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

func TestMarketTypes(t *testing.T) {
	if sessionKey == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	c := CreateClient()

	params := rpc.RPCParams{Filter: rpc.MarketFilter{
		EventTypeIds:    []string{"1"},
		MarketCountries: []string{"GB"},
	}}

	marketTypes, err := c.ListMarketTypes(params)
	if err != nil {
		t.Fatal(err)
	}

	marketCodes := []string{
		"ALT_TOTAL_GOALS",
		"ASIAN_HANDICAP",
		"BOTH_TEAMS_TO_SCORE",
		"CORNER_ODDS",
		"CORRECT_SCORE",
		"DOUBLE_CHANCE",
		"DRAW_NO_BET",
		"FIRST_HALF_GOALS_05",
		"FIRST_HALF_GOALS_15",
		"FIRST_HALF_GOALS_25",
		"HALF_TIME",
		"HALF_TIME_FULL_TIME",
		"HALF_TIME_SCORE",
		"HANDICAP",
		"MATCH_ODDS",
		"OVER_UNDER_05",
		"OVER_UNDER_15",
		"OVER_UNDER_25",
		"OVER_UNDER_35",
		"OVER_UNDER_45",
		"OVER_UNDER_55",
		"OVER_UNDER_65",
		"OVER_UNDER_75",
		"OVER_UNDER_85",
		"PROMOTION",
		"RELEGATION",
		"ROCK_BOTTOM",
		"SPECIALS_NEXT_MGR",
		"TEAM_A_1",
		"TEAM_B_1",
		"TOP_10_FINISH",
		"TOP_2_FINISH",
		"TOP_4_FINISH_FT",
		"TOP_6_FINISH",
		"TOP_GOALSCORER",
		"TOP_N_FINISH",
		"UNDIFFERENTIATED",
		"WINNER",
		"WINNER_WITHOUT",
	}

	for _, marketType := range marketTypes {
		if !slices.Contains(marketCodes, marketType.MarketType) {
			t.Fatal(marketType)
		}
	}
}

func TestEvents(t *testing.T) {
	if sessionKey == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	c := CreateClient()

	params := rpc.RPCParams{Filter: rpc.MarketFilter{
		EventTypeIds:    []string{"1"},
		MarketCountries: []string{"GB"},
	}}

	events, err := c.ListEvents(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(events)
	if len == 0 {
		t.Fatal(len)
	}
}
