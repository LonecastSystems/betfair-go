package stateless

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

func (client *StatelessClient) UpdateOrders(params PlaceOrdersParams) (PlaceExecutionReport, error) {
	json := PlaceExecutionReport{}

	if err := GetSports(client, "placeOrders", params, &json); err != nil {
		return PlaceExecutionReport{}, err
	}

	return json, nil
}
