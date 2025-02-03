package betfairgo

import "testing"

func TestGetDeveloperAppKeys(t *testing.T) {
	c := NewTestBetfairClient(t)

	resp, err := c.GetDeveloperAppKeys()
	if err != nil {
		t.Fatal(err)
	}

	if len(resp) == 0 {
		t.Fatal("No apps")
	}
}
