package stateless

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"

	"github.com/LonecastSystems/betfair-go/common"
	"github.com/LonecastSystems/betfair-go/helpers"
)

type (
	StatelessClient struct {
		Client *common.JsonClient
		Rest   bool
	}
)

func NewStatelessClient(app_key string, rest bool) *StatelessClient {
	return &StatelessClient{Client: common.NewJsonClient(app_key), Rest: rest}
}

func (client *StatelessClient) Do(req *http.Request) (*http.Response, error) {
	return client.Client.Do(req)
}

func (client *StatelessClient) ResumeSession(sessionToken string) (*http.Response, error) {
	return client.Client.ResumeSession(sessionToken)
}

func (client *StatelessClient) NewSession(tls *tls.Config, applicationName string, username string, password string) (*http.Response, error) {
	return client.Client.NewSession(tls, applicationName, username, password)
}

func (client *StatelessClient) ClearSession() (*http.Response, error) {
	return client.Client.ClearSession()
}

const (
	api_account   = "account"
	api_betting   = "betting"
	api_heartbeat = "heartbeat"
	api_scores    = "scores"
)

func GetAccounts[T any](client *StatelessClient, method string, params any, response *T) error {
	return get(client, api_account, method, params, response)
}

func GetSports[T any](client *StatelessClient, method string, params any, response *T) error {
	return get(client, api_betting, method, params, response)
}

func GetHeartbeats[T any](client *StatelessClient, method string, params any, response *T) error {
	return get(client, api_heartbeat, method, params, response)
}

func GetScores[T any](client *StatelessClient, method string, params any, response *T) error {
	return getRPC(client, api_scores, method, params, response) //Only supported by RPC for now.
}

func get[T any](client *StatelessClient, api string, method string, params any, response *T) error {
	if client.Rest {
		return getRest(client, api, method, params, response)
	} else {
		return getRPC(client, api, method, params, response)
	}
}

type JsonAPINGExceptionErrorCode string

const (
	EC_UNEXPECTED_ERROR            = "UNEXPECTED_ERROR"
	EC_INVALID_INPUT_DATA          = "INVALID_INPUT_DATA"
	EC_INVALID_SESSION_INFORMATION = "INVALID_SESSION_INFORMATION"
	EC_INVALID_APP_KEY             = "INVALID_APP_KEY"
	EC_SERVICE_BUSY                = "SERVICE_BUSY"
	EC_TIMEOUT_ERROR               = "TIMEOUT_ERROR"
	EC_NO_SESSION                  = "NO_SESSION"
	EC_NO_APP_KEY                  = "NO_APP_KEY"
	EC_TOO_MANY_REQUESTS           = "TOO_MANY_REQUESTS"
	EC_SERVICE_UNAVAILABLE         = "SERVICE_UNAVAILABLE"
	EC_REQUEST_SIZE_EXCEEDS_LIMIT  = "REQUEST_SIZE_EXCEEDS_LIMIT"
	EC_TOO_MUCH_DATA               = "TOO_MUCH_DATA"
	EC_ACCESS_DENIED               = "ACCESS_DENIED"
)

type (
	JsonRpcResponse struct {
		JsonRPC string      `json:"jsonrpc"`
		Result  interface{} `json:"result"`
		Error   JsonError   `json:"error,omitempty"`
		ID      int         `json:"id"`
	}

	JsonError struct {
		Code    int      `json:"code"`
		Message string   `json:"message"`
		Data    JsonData `json:"data"`
	}

	JsonData struct {
		APINGException JsonAPINGException `json:"APINGException"`
		ExceptionName  string             `json:"exceptionname"`
	}

	JsonAPINGException struct {
		RequestUUID  string                      `json:"requestUUID"`
		ErrorCode    JsonAPINGExceptionErrorCode `json:"errorCode"`
		ErrorDetails string                      `json:"errorDetails"`
	}

	JsonRPC struct {
		JsonRPC string `json:"jsonrpc"`
		Method  string `json:"method"`
		Params  any    `json:"params"`
		ID      int    `json:"id"`
	}
)

var apis = map[string]string{
	api_account:   "AccountAPING",
	api_betting:   "SportsAPING",
	api_heartbeat: "HeartbeatAPING",
	api_scores:    "ScoresAPING",
}

func getRPC[T any](client *StatelessClient, api string, method string, params any, response *T) error {
	query := JsonRPC{
		JsonRPC: "2.0",
		Method:  fmt.Sprintf("%v/v1.0/%v", apis[api], method),
		Params:  params,
		ID:      int(rand.Uint64()),
	}

	body, err := json.Marshal(&query)
	if err != nil {
		return err
	}

	apiUrl := fmt.Sprintf("https://api.betfair.com/exchange/%v/json-rpc/v1/", api)
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

	if jsonError, errorCode := jsonRpc.Error, jsonRpc.Error.Code; errorCode < 0 {
		ex := jsonError.Data.APINGException

		return fmt.Errorf("%v -> %v: %v (%v)", ex.RequestUUID, jsonError.Code, jsonError.Message, ex.ErrorCode)
	}

	if m, err := json.Marshal(jsonRpc.Result); err == nil {
		return json.Unmarshal(m, &response)
	}

	return nil
}

type JsonRestErrorResponse struct {
	FaultCode   string `json:"faultcode"`
	FaultString string `json:"faultstring"`
	Detail      struct {
	} `json:"detail"`
}

func getRest[T any](client *StatelessClient, api string, method string, params any, response *T) error {
	body, err := json.Marshal(&params)
	if err != nil {
		return err
	}

	apiUrl := fmt.Sprintf("https://api.betfair.com/exchange/%v/rest/v1.0/%v/", api, method)
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		jsonRestError := JsonRestErrorResponse{}
		if err = helpers.ReadJson(res, &jsonRestError); err != nil {
			return err
		}

		return fmt.Errorf("%v: %v", jsonRestError.FaultCode, jsonRestError.FaultString)
	}

	if err = helpers.ReadJson(res, &response); err != nil {
		return err
	}

	return nil
}
