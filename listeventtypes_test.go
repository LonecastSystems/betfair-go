package betfairgo

import "testing"

func TestEventTypes(t *testing.T) {
	c := NewTestBetfairClient(t)

	params := EventTypeParams{Filter: MarketFilter{
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
