package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client"
)

func TestListCompetitions(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.CompetitionParams{Filter: client.MarketFilter{
		CompetitionIDs:  []string{"10932509"},
		EventTypeIDs:    []string{"1"},
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

func TestListCountries(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.CountryParams{Filter: client.MarketFilter{
		MarketCountries: []string{"GB"},
	}}

	countries, err := c.ListCountries(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(countries)
	if len != 1 {
		t.Fatal(len)
	}

	countryCode := countries[0].CountryCode
	if countryCode != "GB" {
		t.Fatal(countryCode)
	}
}

func TestListCurrentOrders(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.CurrentOrdersParams{
		FromRecord:  0,
		RecordCount: 1}

	report, err := c.ListCurrentOrders(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(report.CurrentOrders) == 0 {
		t.Fatal("No orders")
	}
}

func TestEvents(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.EventParams{Filter: client.MarketFilter{
		EventTypeIDs:    []string{"1"},
		MarketCountries: []string{"GB"},
	}}

	events, err := c.ListEvents(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(events)
	if len == 0 {
		t.Fatal("No events")
	}
}

func TestEventTypes(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.EventTypeParams{Filter: client.MarketFilter{
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

func TestMarketBook(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.MarketBookParams{
		MarketIDs: []string{"1.233455113"},
		PriceProjection: client.PriceProjection{
			PriceData: []client.PriceData{"EX_ALL_OFFERS"}}}

	marketBooks, err := c.ListMarketBook(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(marketBooks)
	if len == 0 {
		t.Fatal("No market books")
	}
}

func TestMarketCatalogue(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.MarketCatalogueParams{
		Filter:     client.MarketFilter{},
		MaxResults: 10}

	markets, err := c.ListMarketCatalogue(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(markets)
	if len == 0 {
		t.Fatal("No markets")
	}
}

func TestListMarketProfitAndLoss(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.MarketProfitAndLossParams{
		MarketIDs: []string{"1.233455113"}}

	pnl, err := c.ListMarketProfitAndLoss(params)
	if err != nil {
		t.Fatal(err)
	}

	len := len(pnl)
	if len == 0 {
		t.Fatal("No PnL")
	}
}

func TestMarketTypes(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.MarketTypeParams{Filter: client.MarketFilter{
		EventTypeIDs:    []string{"1"},
		MarketCountries: []string{"GB"},
	}}

	marketTypes, err := c.ListMarketTypes(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(marketTypes) == 0 {
		t.Fatal("No market types")
	}
}

func TestListRunnerBook(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.RunnerBookParams{
		MarketID: "1.233455113",
		PriceProjection: client.PriceProjection{
			PriceData: []client.PriceData{"EX_ALL_OFFERS"}}}

	runners, err := c.ListRunnerBook(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(runners) == 0 {
		t.Fatal("No runner books")
	}
}

func TestListTimeRanges(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.TimeRangesParams{Filter: client.MarketFilter{
		EventTypeIDs: []string{"1"}},
		Granularity: "DAYS"}

	timeRanges, err := c.ListTimeRanges(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(timeRanges) == 0 {
		t.Fatal("No time ranges")
	}
}

func TestListVenues(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.VenueParams{Filter: client.MarketFilter{
		MarketCountries: []string{"GB"},
	}}

	venues, err := c.ListVenues(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(venues) == 0 {
		t.Fatal("No venues")
	}
}

func TestListClearedOrders(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := client.ClearedOrdersParams{BetStatus: "SETTLED"}

	report, err := c.ListClearedOrders(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(report.ClearedOrders) == 0 {
		t.Fatal("No cleared orders")
	}
}
