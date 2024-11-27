package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/LonecastSystems/betfair-go/client/common"
	"github.com/LonecastSystems/betfair-go/helpers"
)

const betfairUrl = "https://api.betfair.com/exchange/%v/json-rpc/v1/"

const (
	api_account = "account"
	api_betting = "betting"
)

var apis = map[string]string{
	api_account: "AccountAPING",
	api_betting: "SportsAPING",
}

func CreateClient(sessionToken string, app_key string) *RpcClient {
	return &RpcClient{Client: common.CreateClient(sessionToken, app_key)}
}

func (client *RpcClient) Do(req *http.Request) (*http.Response, error) {
	return client.Client.Do(req)
}

func GetAccounts[T any, TParams any](client *RpcClient, id int, method string, params TParams, response *T) error {
	return get(client, api_account, id, method, params, response)
}

func GetSports[T any, TParams any](client *RpcClient, id int, method string, params TParams, response *T) error {
	return get(client, api_betting, id, method, params, response)
}

func get[T any, TParams any](client *RpcClient, api string, id int, method string, params TParams, response *T) error {
	query := JsonRPC[TParams]{
		JsonRPC: "2.0",
		Method:  fmt.Sprintf("%v/v1.0/%v", apis[api], method),
		Params:  params,
		ID:      id,
	}

	body, err := json.Marshal(&query)
	if err != nil {
		return err
	}

	apiUrl := fmt.Sprintf(betfairUrl, api)

	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(body))
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

	if errorCode := jsonRpc.Error.Code; errorCode < 0 {
		if errorCodex := jsonRpc.Error.Data.APINGException.ErrorCode; errorCodex != "" {
			return errors.New(errorCodex)
		} else {
			return errors.New(strconv.Itoa(errorCode))
		}
	}

	if m, err := json.Marshal(jsonRpc.Result); err == nil {
		return json.Unmarshal(m, &response)
	}

	return nil
}
