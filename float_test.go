package betfair

import (
	"encoding/json"
	"math"
	"testing"
)

func TestFloatUnmarshalAcceptsNaNAndInfinity(t *testing.T) {
	t.Parallel()

	var prices StartingPrices
	err := json.Unmarshal([]byte(`{
		"nearPrice": "NaN",
		"farPrice": 12.5,
		"actualSP": "Infinity",
		"backStakeTaken": [{"price": 1.01, "size": 2}]
	}`), &prices)
	if err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if !math.IsNaN(float64(prices.NearPrice)) {
		t.Fatalf("NearPrice = %v, want NaN", prices.NearPrice)
	}
	if float64(prices.FarPrice) != 12.5 {
		t.Fatalf("FarPrice = %v, want 12.5", prices.FarPrice)
	}
	if !math.IsInf(float64(prices.ActualSP), 1) {
		t.Fatalf("ActualSP = %v, want +Inf", prices.ActualSP)
	}
	if len(prices.BackStakeTaken) != 1 || prices.BackStakeTaken[0].Size != 2 {
		t.Fatalf("BackStakeTaken = %#v", prices.BackStakeTaken)
	}
}

func TestPriceSizeUnmarshalAcceptsInfinitySize(t *testing.T) {
	t.Parallel()

	var offer PriceSize
	if err := json.Unmarshal([]byte(`{"price":2.1,"size":"Infinity"}`), &offer); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if float64(offer.Price) != 2.1 || !math.IsInf(float64(offer.Size), 1) {
		t.Fatalf("PriceSize = %#v", offer)
	}
}

func TestRunnerChangeUnmarshalAcceptsStreamSpecials(t *testing.T) {
	t.Parallel()

	var change RunnerChange
	err := json.Unmarshal([]byte(`{
		"spn":"Infinity",
		"spf":"NaN",
		"bdatb":[[0,2.2,"Infinity"]],
		"atb":[[1.5,10]]
	}`), &change)
	if err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if change.StartingPriceNear == nil || !math.IsInf(float64(*change.StartingPriceNear), 1) {
		t.Fatalf("StartingPriceNear = %#v", change.StartingPriceNear)
	}
	if change.StartingPriceFar == nil || !math.IsNaN(float64(*change.StartingPriceFar)) {
		t.Fatalf("StartingPriceFar = %#v", change.StartingPriceFar)
	}
	if len(change.BestDisplayAvailableToBack) != 1 ||
		!math.IsInf(float64(change.BestDisplayAvailableToBack[0][2]), 1) {
		t.Fatalf("BestDisplayAvailableToBack = %#v", change.BestDisplayAvailableToBack)
	}
	if len(change.AvailableToBack) != 1 || float64(change.AvailableToBack[0][1]) != 10 {
		t.Fatalf("AvailableToBack = %#v", change.AvailableToBack)
	}
}
