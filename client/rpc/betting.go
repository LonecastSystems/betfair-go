package rpc

func (client *JsonRpcClient) ListEventTypes(params RPCParams) ([]EventTypeResult, error) {
	json := []EventTypeResult{}

	if err := Get(client, 1, "listEventTypes", params, &json); err != nil {
		return []EventTypeResult{}, err
	}

	return json, nil
}

func (client *JsonRpcClient) ListEvents(params RPCParams) ([]EventResult, error) {
	json := []EventResult{}

	if err := Get(client, 1, "listEvents", params, &json); err != nil {
		return []EventResult{}, err
	}

	return json, nil
}

func (client *JsonRpcClient) ListCompetitions(params RPCParams) ([]CompetitionResult, error) {
	json := []CompetitionResult{}

	if err := Get(client, 1, "listCompetitions", params, &json); err != nil {
		return []CompetitionResult{}, err
	}

	return json, nil
}

func (client *JsonRpcClient) ListMarketTypes(params RPCParams) ([]MarketTypeResult, error) {
	json := []MarketTypeResult{}

	if err := Get(client, 1, "listMarketTypes", params, &json); err != nil {
		return []MarketTypeResult{}, err
	}

	return json, nil
}
