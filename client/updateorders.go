package client

type (
	UpdateOrdersParams struct {
		MarketID     string              `json:"marketId"`
		Instructions []UpdateInstruction `json:"instructions"`
		CustomerRef  string              `json:"customerRef,omitempty"`
	}

	UpdateInstruction struct {
		BetID              string          `json:"betId"`
		NewPersistenceType PersistenceType `json:"newPersistenceType,omitempty"`
	}
)

func (client *BetfairClient) UpdateOrders(params PlaceOrdersParams) (PlaceExecutionReport, error) {
	json := PlaceExecutionReport{}

	if err := client.GetSports("placeOrders", params, &json); err != nil {
		return PlaceExecutionReport{}, err
	}

	return json, nil
}
