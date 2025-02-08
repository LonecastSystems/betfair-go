package betfair

import "testing"

func TestListRunnerBook(t *testing.T) {
	c := NewTestClient(t)

	params := RunnerBookParams{
		MarketID: "1.233455113",
		PriceProjection: PriceProjection{
			PriceData: []PriceData{"EX_ALL_OFFERS"}}}

	runners, err := c.ListRunnerBook(params)
	if err != nil {
		t.Fatal(err)
	}

	if len(runners) == 0 {
		t.Fatal("No runner books")
	}
}
