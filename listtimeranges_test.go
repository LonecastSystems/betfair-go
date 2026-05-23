package betfair

import (
	"context"
	"testing"
)

func TestListTimeRanges(t *testing.T) {
	c := newTestClient(t)

	params := ListTimeRangesParams{Filter: MarketFilter{
		EventTypeIDs: []string{"1"}},
		Granularity: TG_DAYS}

	timeRanges, err := c.ListTimeRanges(context.Background(), params)
	if err != nil {
		t.Fatal(err)
	}

	if len(timeRanges) == 0 {
		t.Fatal("No time ranges")
	}
}
