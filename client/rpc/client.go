package rpc

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/LonecastSystems/betfair-go/helpers"
)

var betfairUrl = "https://api.betfair.com/exchange/%v/json-rpc/v1/"

const (
	api_account = "account"
	api_betting = "betting"
)

var apis = map[string]string{
	api_account: "AccountAPING",
	api_betting: "SportsAPING",
}

func (client *JsonRpcClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Authentication", client.SessionToken)
	req.Header.Add("X-Application", client.ApplicationKey)
	req.Header.Add("Accept", "application/json")

	return client.Client.Do(req)
}

func CreateClient(sessionToken string, app_key string) *JsonRpcClient {
	return &JsonRpcClient{Client: &http.Client{}, SessionToken: sessionToken, ApplicationKey: app_key}
}

func Login(tls *tls.Config, apiKey string, applicationName string, username string, password string) (jsonClient *JsonRpcClient, response *http.Response, err error) {
	postUrl := url.URL{Path: "https://identitysso-cert.betfair.com/api/certlogin"}
	q := postUrl.Query()
	q.Set("username", username)
	q.Set("password", password)

	postUrl.RawQuery = q.Encode()

	req, _ := http.NewRequest("POST", postUrl.RequestURI(), nil)
	req.SetBasicAuth(username, password)

	req.Header.Add("X-Application", applicationName)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tls,
		},
	}

	resp, err := client.Do(req)
	if err == nil {
		json := SessionResponse{}
		if err := helpers.ReadJson(resp, &json); err == nil {
			return CreateClient(json.SessionToken, apiKey), resp, nil
		}
	}

	return nil, resp, err
}

func (client *JsonRpcClient) Logout() (response *http.Response, err error) {
	postUrl := url.URL{Path: "https://identitysso.betfair.com/api/logout"}

	req, _ := http.NewRequest("POST", postUrl.RequestURI(), nil)

	resp, err := client.Do(req)
	if err == nil {
		json := SessionLogoutResponse{}
		return resp, helpers.ReadJson(resp, &json)
	}

	return resp, err
}

func GetAccounts[T any, TParams any](client *JsonRpcClient, id int, method string, params TParams, response *T) error {
	return Get(client, api_account, id, method, params, response)
}

func GetSports[T any, TParams any](client *JsonRpcClient, id int, method string, params TParams, response *T) error {
	return Get(client, api_betting, id, method, params, response)
}

func Get[T any, TParams any](client *JsonRpcClient, api string, id int, method string, params TParams, response *T) error {
	query := JsonRPC[TParams]{
		JsonRPC: "2.0",
		Method:  fmt.Sprintf("%v/v1.0/%v", apis[api], method),
		Params:  params,
		ID:      1,
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
