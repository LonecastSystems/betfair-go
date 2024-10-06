package rpc

func (client *JsonRpcClient) ListEventTypes(params MarketParams) ([]EventTypeResult, error) {
	json := []EventTypeResult{}

	if err := GetSports(client, 1, "listEventTypes", params, &json); err != nil {
		return []EventTypeResult{}, err
	}

	return json, nil
}

func (client *JsonRpcClient) ListEvents(params MarketParams) ([]EventResult, error) {
	json := []EventResult{}

	if err := GetSports(client, 1, "listEvents", params, &json); err != nil {
		return []EventResult{}, err
	}

	return json, nil
}

func (client *JsonRpcClient) ListCompetitions(params MarketParams) ([]CompetitionResult, error) {
	json := []CompetitionResult{}

	if err := GetSports(client, 1, "listCompetitions", params, &json); err != nil {
		return []CompetitionResult{}, err
	}

	return json, nil
}

func (client *JsonRpcClient) ListMarketTypes(params MarketParams) ([]MarketTypeResult, error) {
	json := []MarketTypeResult{}

	if err := GetSports(client, 1, "listMarketTypes", params, &json); err != nil {
		return []MarketTypeResult{}, err
	}

	return json, nil
}

func (client *JsonRpcClient) ListMarketCatalogue(params MarketParams) ([]MarketCatalogueResult, error) {
	json := []MarketCatalogueResult{}

	if err := GetSports(client, 1, "listMarketCatalogue", params, &json); err != nil {
		return []MarketCatalogueResult{}, err
	}

	return json, nil
}

func (client *JsonRpcClient) ListMarketBook(params MarketBookParams) ([]MarketBookResult, error) {
	json := []MarketBookResult{}

	if err := GetSports(client, 1, "listMarketBook", params, &json); err != nil {
		return []MarketBookResult{}, err
	}

	return json, nil
}

func (client *JsonRpcClient) GetCurrentOrders(params CurrentOrdersParams) (CurrentOrderResult, error) {
	json := CurrentOrderResult{}

	if err := GetSports(client, 1, "listCurrentOrders", params, &json); err != nil {
		return CurrentOrderResult{}, err
	}

	return json, nil
}
