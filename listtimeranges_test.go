package betfair

import "testing"

func TestListTimeRanges(t *testing.T) {
	c := NewTestClient(t)

	params := TimeRangesParams{Filter: MarketFilter{
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
