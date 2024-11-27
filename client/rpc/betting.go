package rpc

func (client *RpcClient) ListEventTypes(params MarketParams) ([]EventTypeResult, error) {
	json := []EventTypeResult{}

	if err := GetSports(client, 1, "listEventTypes", params, &json); err != nil {
		return []EventTypeResult{}, err
	}

	return json, nil
}

func (client *RpcClient) ListEvents(params MarketParams) ([]EventResult, error) {
	json := []EventResult{}

	if err := GetSports(client, 2, "listEvents", params, &json); err != nil {
		return []EventResult{}, err
	}

	return json, nil
}

func (client *RpcClient) ListCompetitions(params MarketParams) ([]CompetitionResult, error) {
	json := []CompetitionResult{}

	if err := GetSports(client, 3, "listCompetitions", params, &json); err != nil {
		return []CompetitionResult{}, err
	}

	return json, nil
}

func (client *RpcClient) ListMarketTypes(params MarketParams) ([]MarketTypeResult, error) {
	json := []MarketTypeResult{}

	if err := GetSports(client, 4, "listMarketTypes", params, &json); err != nil {
		return []MarketTypeResult{}, err
	}

	return json, nil
}

func (client *RpcClient) ListMarketCatalogue(params MarketParams) ([]MarketCatalogueResult, error) {
	json := []MarketCatalogueResult{}

	if err := GetSports(client, 5, "listMarketCatalogue", params, &json); err != nil {
		return []MarketCatalogueResult{}, err
	}

	return json, nil
}

func (client *RpcClient) ListMarketBook(params MarketBookParams) ([]MarketBookResult, error) {
	json := []MarketBookResult{}

	if err := GetSports(client, 6, "listMarketBook", params, &json); err != nil {
		return []MarketBookResult{}, err
	}

	return json, nil
}

func (client *RpcClient) ListCurrentOrders(params CurrentOrdersParams) (CurrentOrderSummaryReport, error) {
	json := CurrentOrderSummaryReport{}

	if err := GetSports(client, 7, "listCurrentOrders", params, &json); err != nil {
		return CurrentOrderSummaryReport{}, err
	}

	return json, nil
}
