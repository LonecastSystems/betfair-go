package betfair

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/url"
)

type SessionStatus string

const (
	SS_SUCCESS = "SUCCESS"
	SS_FAIL    = "FAIL"
)

type SessionStatusError string

const (
	SS_ERR_INPUT_VALIDATION_ERROR = "INPUT_VALIDATION_ERROR"
	SS_ERR_INTERNAL_ERROR         = "INTERNAL_ERROR"
	SS_ERR_NO_SESSION             = "NO_SESSION"
)

type (
	Client struct {
		HttpClient      *http.Client
		Tls             *tls.Config
		ApplicationName string
		ApplicationKey  string
		SessionToken    string
		Rest            bool
	}

	SessionResponse struct {
		SessionToken string `json:"sessionToken"`
		LoginStatus  string `json:"loginStatus"`
	}

	SessionStatusResponse struct {
		Token   string             `json:"token"`
		Product string             `json:"product"`
		Status  SessionStatus      `json:"status"`
		Error   SessionStatusError `json:"error"`
	}
)

func NewClient(tls *tls.Config, app_key string, applicationName string) *Client {
	return &Client{Tls: tls, ApplicationKey: app_key, ApplicationName: applicationName}
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	if client.HttpClient == nil {
		return nil, errors.New("client not initialised: please resume or create a new session")
	}

	req.Header.Add("X-Authentication", client.SessionToken)
	req.Header.Add("X-Application", client.ApplicationKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("content-type", "application/json")

	return client.HttpClient.Do(req)
}

func (client *Client) Resume(sessionToken string) (*http.Response, error) {
	client.HttpClient = &http.Client{}
	client.SessionToken = sessionToken

	keepAliveUrl := url.URL{Path: "https://identitysso.betfair.com/api/keepAlive"}
	req, _ := http.NewRequest("POST", keepAliveUrl.RequestURI(), nil)

	resp, err := client.Do(req)
	if err != nil {
		client.HttpClient = nil
		client.SessionToken = ""
		return resp, err
	}

	json := SessionStatusResponse{}
	if err := ReadJson(resp, &json); err != nil {
		client.HttpClient = nil
		client.SessionToken = ""
		return resp, err
	}

	if json.Error != "" {
		client.HttpClient = nil
		client.SessionToken = ""
		return resp, errors.New(string(json.Error))
	}

	return resp, nil
}

func (client *Client) Login(username string, password string) (*http.Response, error) {
	postUrl := url.URL{Path: "https://identitysso-cert.betfair.com/api/certlogin"}
	q := postUrl.Query()
	q.Set("username", username)
	q.Set("password", password)

	postUrl.RawQuery = q.Encode()

	req, _ := http.NewRequest("POST", postUrl.RequestURI(), nil)
	req.SetBasicAuth(username, password)

	req.Header.Add("X-Application", client.ApplicationName)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: client.Tls,
		},
	}

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		client.HttpClient = nil
		return resp, err
	}

	json := SessionResponse{}
	if err := ReadJson(resp, &json); err != nil {
		client.HttpClient = nil
		return resp, err
	}

	client.SessionToken = json.SessionToken
	return resp, nil
}

func (client *Client) Logout() (*http.Response, error) {
	postUrl := url.URL{Path: "https://identitysso.betfair.com/api/logout"}

	req, _ := http.NewRequest("POST", postUrl.RequestURI(), nil)

	resp, err := client.Do(req)
	if err != nil {
		return resp, err
	}

	json := SessionStatusResponse{}
	if err := ReadJson(resp, &json); err != nil {
		return resp, err
	}

	if json.Error != "" {
		return resp, errors.New(string(json.Error))
	}

	client.HttpClient = nil
	client.SessionToken = ""

	return resp, nil
}

func (client *Client) GetStream(sc *StreamingClient) error {
	if client.HttpClient == nil {
		return errors.New("client not initialised: please resume or create a new session")
	}

	return sc.Authenticate(client.Tls, client.ApplicationKey, client.SessionToken)
}

const (
	api_account   = "account"
	api_betting   = "betting"
	api_heartbeat = "heartbeat"
	api_scores    = "scores"
)

func (client *Client) GetAccounts(method string, params any, response any) error {
	return client.get(api_account, method, params, response)
}

func (client *Client) GetSports(method string, params any, response any) error {
	return client.get(api_betting, method, params, response)
}

func (client *Client) GetHeartbeats(method string, params any, response any) error {
	return client.get(api_heartbeat, method, params, response)
}

func (client *Client) GetScores(method string, params any, response any) error {
	return client.getRPC(api_scores, method, params, response) //Only supported by RPC for now.
}

func (client *Client) get(api string, method string, params any, response any) error {
	if client.Rest {
		return client.getRest(api, method, params, response)
	} else {
		return client.getRPC(api, method, params, response)
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
		JsonRPC string    `json:"jsonrpc"`
		Result  any       `json:"result"`
		Error   JsonError `json:"error,omitempty"`
		ID      int       `json:"id"`
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

func (client *Client) getRPC(api string, method string, params any, response any) error {
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
	if err = ReadJson(res, &jsonRpc); err != nil {
		return err
	}

	if jsonError, errorCode := jsonRpc.Error, jsonRpc.Error.Code; errorCode < 0 {
		ex := jsonError.Data.APINGException

		return fmt.Errorf("%v -> %v: %v (%v)", jsonRpc.ID, jsonError.Code, jsonError.Message, ex.ErrorCode)
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

func (client *Client) getRest(api string, method string, params any, response any) error {
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
		if err = ReadJson(res, &jsonRestError); err != nil {
			return err
		}

		return fmt.Errorf("%v: %v", jsonRestError.FaultCode, jsonRestError.FaultString)
	}

	if err = ReadJson(res, &response); err != nil {
		return err
	}

	return nil
}
