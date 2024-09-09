package betting

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/LonecastSystems/betfair-go/client"
	"github.com/LonecastSystems/betfair-go/helpers"
)

var postUrl = url.URL{Path: "https://api.betfair.com/exchange/betting/json-rpc/v1/"}

func (client *BettingClient) Do(req *http.Request) (*http.Response, error) {
	return client.Client.Do(req)
}

func CreateClient(sessionToken string, app_key string) *BettingClient {
	return &BettingClient{Client: *client.CreateClient(sessionToken, app_key)}
}

func Get[T any](client *BettingClient, id int, method string, params RPCParams, response *T) error {
	query := JsonRPC{
		JsonRPC: "2.0",
		Method:  fmt.Sprintf("SportsAPING/v1.0/%v", method),
		Params:  params,
		ID:      1,
	}

	body, err := json.Marshal(&query)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", postUrl.RequestURI(), bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	jsonRpc := JsonRpcResponse{}
	if err = helpers.ReadJson(res, &jsonRpc); err != nil {
		return err
	}

	if m, err := json.Marshal(jsonRpc.Result); err == nil {
		return json.Unmarshal(m, &response)
	}

	return nil
}

func (client *BettingClient) ListCompetitions(params RPCParams) ([]CompetitionResult, error) {
	json := []CompetitionResult{}

	if err := Get(client, 1, "listCompetitions", params, &json); err != nil {
		return []CompetitionResult{}, err
	}

	return json, nil
}

func (client *BettingClient) ListEventTypes(params RPCParams) ([]EventTypeResult, error) {
	json := []EventTypeResult{}

	if err := Get(client, 1, "listEventTypes", params, &json); err != nil {
		return []EventTypeResult{}, err
	}

	return json, nil
}

func (client *BettingClient) ListEvents(params RPCParams) ([]EventResult, error) {
	json := []EventResult{}

	if err := Get(client, 1, "listEvents", params, &json); err != nil {
		return []EventResult{}, err
	}

	return json, nil
}

func (client *BettingClient) ListMarketTypes(params RPCParams) ([]MarketTypeResult, error) {
	json := []MarketTypeResult{}

	if err := Get(client, 1, "listMarketTypes", params, &json); err != nil {
		return []MarketTypeResult{}, err
	}

	return json, nil
}
