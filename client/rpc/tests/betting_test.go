package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client/rpc"
)

func TestEventTypes(t *testing.T) {
	c := CreateClient(t)

	params := rpc.MarketParams{Filter: rpc.MarketFilter{
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
	c := CreateClient(t)

	params := rpc.MarketParams{Filter: rpc.MarketFilter{
		CompetitionIds:  []string{"10932509"},
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
	c := CreateClient(t)

	params := rpc.MarketParams{Filter: rpc.MarketFilter{
		EventTypeIds:    []string{"1"},
		MarketCountries: []string{"GB"},
	}}

	marketTypes, err := c.ListMarketTypes(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(marketTypes) == 0 {
		t.Fatal("Empty market types")
	}
}

func TestEvents(t *testing.T) {
	c := CreateClient(t)

	params := rpc.MarketParams{Filter: rpc.MarketFilter{
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

func TestMarketCatalogue(t *testing.T) {
	c := CreateClient(t)

	params := rpc.MarketParams{Filter: rpc.MarketFilter{}, MaxResults: "10"}

	markets, err := c.ListMarketCatalogue(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(markets)
	if len == 0 {
		t.Fatal(len)
	}
}

func TestMarketBook(t *testing.T) {
	c := CreateClient(t)

	params := rpc.MarketBookParams{MarketIds: []string{"1.233455113"}, PriceProjection: rpc.PriceProjection{PriceData: []string{"EX_ALL_OFFERS"}}}

	marketDetails, err := c.ListMarketBook(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(marketDetails)
	if len == 0 {
		t.Fatal(len)
	}
}

func TestListCurrentOrders(t *testing.T) {
	c := CreateClient(t)

	_, err := c.ListCurrentOrders(rpc.CurrentOrdersParams{FromRecord: 0, RecordCount: 1})
	if err != nil {
		t.Fatal(err)
	}
}
