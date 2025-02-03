package client

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
	"net/http"
	"net/url"

	"github.com/LonecastSystems/betfair-go/streaming"
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
	BetfairClient struct {
		Client          *http.Client
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

func NewBetfairClient(tls *tls.Config, app_key string, applicationName string) *BetfairClient {
	return &BetfairClient{Tls: tls, ApplicationKey: app_key, ApplicationName: applicationName}
}

func (client *BetfairClient) Do(req *http.Request) (*http.Response, error) {
	if client.Client == nil {
		return nil, errors.New("client not initialised: please resume or create a new session")
	}

	req.Header.Add("X-Authentication", client.SessionToken)
	req.Header.Add("X-Application", client.ApplicationKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("content-type", "application/json")

	return client.Client.Do(req)
}

func (jsonClient *BetfairClient) Resume(sessionToken string) (*http.Response, error) {
	jsonClient.Client = &http.Client{}
	jsonClient.SessionToken = sessionToken

	keepAliveUrl := url.URL{Path: "https://identitysso.betfair.com/api/keepAlive"}
	req, _ := http.NewRequest("POST", keepAliveUrl.RequestURI(), nil)

	resp, err := jsonClient.Do(req)
	if err != nil {
		jsonClient.Client = nil
		jsonClient.SessionToken = ""
		return resp, err
	}

	json := SessionStatusResponse{}
	if err := readJson(resp, &json); err != nil {
		jsonClient.Client = nil
		jsonClient.SessionToken = ""
		return resp, err
	}

	if json.Error != "" {
		jsonClient.Client = nil
		jsonClient.SessionToken = ""
		return resp, errors.New(string(json.Error))
	}

	return resp, nil
}

func (jsonClient *BetfairClient) Login(username string, password string) (*http.Response, error) {
	postUrl := url.URL{Path: "https://identitysso-cert.betfair.com/api/certlogin"}
	q := postUrl.Query()
	q.Set("username", username)
	q.Set("password", password)

	postUrl.RawQuery = q.Encode()

	req, _ := http.NewRequest("POST", postUrl.RequestURI(), nil)
	req.SetBasicAuth(username, password)

	req.Header.Add("X-Application", jsonClient.ApplicationName)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	jsonClient.Client = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: jsonClient.Tls,
		},
	}

	resp, err := jsonClient.Client.Do(req)
	if err != nil {
		jsonClient.Client = nil
		return resp, err
	}

	json := SessionResponse{}
	if err := readJson(resp, &json); err != nil {
		jsonClient.Client = nil
		return resp, err
	}

	jsonClient.SessionToken = json.SessionToken
	return resp, nil
}

func (jsonClient *BetfairClient) Logout() (*http.Response, error) {
	postUrl := url.URL{Path: "https://identitysso.betfair.com/api/logout"}

	req, _ := http.NewRequest("POST", postUrl.RequestURI(), nil)

	resp, err := jsonClient.Do(req)
	if err != nil {
		return resp, err
	}

	json := SessionStatusResponse{}
	if err := readJson(resp, &json); err != nil {
		return resp, err
	}

	if json.Error != "" {
		return resp, errors.New(string(json.Error))
	}

	jsonClient.Client = nil
	jsonClient.SessionToken = ""

	return resp, nil
}

func (client *BetfairClient) LoginStream(sc *streaming.StreamingClient) error {
	if client.Client == nil {
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

func (client *BetfairClient) GetAccounts(method string, params any, response any) error {
	return client.get(api_account, method, params, response)
}

func (client *BetfairClient) GetSports(method string, params any, response any) error {
	return client.get(api_betting, method, params, response)
}

func (client *BetfairClient) GetHeartbeats(method string, params any, response any) error {
	return client.get(api_heartbeat, method, params, response)
}

func (client *BetfairClient) GetScores(method string, params any, response any) error {
	return client.getRPC(api_scores, method, params, response) //Only supported by RPC for now.
}

func (client *BetfairClient) get(api string, method string, params any, response any) error {
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

func (client *BetfairClient) getRPC(api string, method string, params any, response any) error {
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
	if err = readJson(res, &jsonRpc); err != nil {
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

func (client *BetfairClient) getRest(api string, method string, params any, response any) error {
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
		if err = readJson(res, &jsonRestError); err != nil {
			return err
		}

		return fmt.Errorf("%v: %v", jsonRestError.FaultCode, jsonRestError.FaultString)
	}

	if err = readJson(res, &response); err != nil {
		return err
	}

	return nil
}

func readJson(res *http.Response, response any) error {
	defer res.Body.Close()

	if body, err := io.ReadAll(res.Body); err == nil {
		return json.Unmarshal(body, &response)
	}

	return nil
}

func GetTLSConfig(certFilePath string, keyFilePath string) (*tls.Config, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	clientTLSCert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientTLSCert},
	}, nil
}
