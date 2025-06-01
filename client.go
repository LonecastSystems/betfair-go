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

type LoginStatus string

const (
	LS_SUCCESS                                 LoginStatus = "SUCCESS"
	LS_INVALID_USERNAME_OR_PASSWORD            LoginStatus = "INVALID_USERNAME_OR_PASSWORD"
	LS_ACCOUNT_NOW_LOCKED                      LoginStatus = "ACCOUNT_NOW_LOCKED"
	LS_ACCOUNT_ALREADY_LOCKED                  LoginStatus = "ACCOUNT_ALREADY_LOCKED"
	LS_PENDING_AUTH                            LoginStatus = "PENDING_AUTH"
	LS_TELBET_TERMS_CONDITIONS_NA              LoginStatus = "TELBET_TERMS_CONDITIONS_NA"
	LS_DUPLICATE_CARDS                         LoginStatus = "DUPLICATE_CARDS"
	LS_SECURITY_QUESTION_WRONG_3X              LoginStatus = "SECURITY_QUESTION_WRONG_3X"
	LS_KYC_SUSPEND                             LoginStatus = "KYC_SUSPEND"
	LS_SUSPENDED                               LoginStatus = "SUSPENDED"
	LS_CLOSED                                  LoginStatus = "CLOSED"
	LS_SELF_EXCLUDED                           LoginStatus = "SELF_EXCLUDED"
	LS_INVALID_CONNECTIVITY_TO_REGULATOR_DK    LoginStatus = "INVALID_CONNECTIVITY_TO_REGULATOR_DK"
	LS_NOT_AUTHORIZED_BY_REGULATOR_DK          LoginStatus = "NOT_AUTHORIZED_BY_REGULATOR_DK"
	LS_INVALID_CONNECTIVITY_TO_REGULATOR_IT    LoginStatus = "INVALID_CONNECTIVITY_TO_REGULATOR_IT"
	LS_NOT_AUTHORIZED_BY_REGULATOR_IT          LoginStatus = "NOT_AUTHORIZED_BY_REGULATOR_IT"
	LS_SECURITY_RESTRICTED_LOCATION            LoginStatus = "SECURITY_RESTRICTED_LOCATION"
	LS_BETTING_RESTRICTED_LOCATION             LoginStatus = "BETTING_RESTRICTED_LOCATION"
	LS_TRADING_MASTER                          LoginStatus = "TRADING_MASTER"
	LS_TRADING_MASTER_SUSPENDED                LoginStatus = "TRADING_MASTER_SUSPENDED"
	LS_AGENT_CLIENT_MASTER                     LoginStatus = "AGENT_CLIENT_MASTER"
	LS_AGENT_CLIENT_MASTER_SUSPENDED           LoginStatus = "AGENT_CLIENT_MASTER_SUSPENDED"
	LS_DANISH_AUTHORIZATION_REQUIRED           LoginStatus = "DANISH_AUTHORIZATION_REQUIRED"
	LS_SPAIN_MIGRATION_REQUIRED                LoginStatus = "SPAIN_MIGRATION_REQUIRED"
	LS_DENMARK_MIGRATION_REQUIRED              LoginStatus = "DENMARK_MIGRATION_REQUIRED"
	LS_SPANISH_TERMS_ACCEPTANCE_REQUIRED       LoginStatus = "SPANISH_TERMS_ACCEPTANCE_REQUIRED"
	LS_ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED    LoginStatus = "ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED"
	LS_CERT_AUTH_REQUIRED                      LoginStatus = "CERT_AUTH_REQUIRED"
	LS_CHANGE_PASSWORD_REQUIRED                LoginStatus = "CHANGE_PASSWORD_REQUIRED"
	LS_PERSONAL_MESSAGE_REQUIRED               LoginStatus = "PERSONAL_MESSAGE_REQUIRED"
	LS_INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED LoginStatus = "INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED"
	LS_EMAIL_LOGIN_NOT_ALLOWED                 LoginStatus = "EMAIL_LOGIN_NOT_ALLOWED"
	LS_MULTIPLE_USERS_WITH_SAME_CREDENTIAL     LoginStatus = "MULTIPLE_USERS_WITH_SAME_CREDENTIAL"
	LS_ACCOUNT_PENDING_PASSWORD_CHANGE         LoginStatus = "ACCOUNT_PENDING_PASSWORD_CHANGE"
	LS_TEMPORARY_BAN_TOO_MANY_REQUESTS         LoginStatus = "TEMPORARY_BAN_TOO_MANY_REQUESTS"
	LS_ITALIAN_PROFILING_ACCEPTANCE_REQUIRED   LoginStatus = "ITALIAN_PROFILING_ACCEPTANCE_REQUIRED"
	LS_AUTHORIZED_ONLY_FOR_DOMAIN_RO           LoginStatus = "AUTHORIZED_ONLY_FOR_DOMAIN_RO"
	LS_AUTHORIZED_ONLY_FOR_DOMAIN_SE           LoginStatus = "AUTHORIZED_ONLY_FOR_DOMAIN_SE"
	LS_SWEDEN_NATIONAL_IDENTIFIER_REQUIRED     LoginStatus = "SWEDEN_NATIONAL_IDENTIFIER_REQUIRED"
	LS_SWEDEN_BANK_ID_VERIFICATION_REQUIRED    LoginStatus = "SWEDEN_BANK_ID_VERIFICATION_REQUIRED"
	LS_ACTIONS_REQUIRED                        LoginStatus = "ACTIONS_REQUIRED"
	LS_INPUT_VALIDATION_ERROR                  LoginStatus = "INPUT_VALIDATION_ERROR"
	LS_MIGRATION_REQUIRED                      LoginStatus = "MIGRATION_REQUIRED"
	LS_TERMS_AND_CONDITIONS                    LoginStatus = "TERMS_AND_CONDITIONS"
	LS_CONTACT_VERIFICATION_REQUIRED           LoginStatus = "CONTACT_VERIFICATION_REQUIRED"
)

type SessionStatus string

const (
	SS_SUCCESS SessionStatus = "SUCCESS"
	SS_FAIL    SessionStatus = "FAIL"
)

type SessionStatusError string

const (
	SS_ERR_INPUT_VALIDATION_ERROR SessionStatusError = "INPUT_VALIDATION_ERROR"
	SS_ERR_INTERNAL_ERROR         SessionStatusError = "INTERNAL_ERROR"
	SS_ERR_NO_SESSION             SessionStatusError = "NO_SESSION"
)

type Jurisdiction string

const (
	JD_GLOBAL    = "com"
	JD_AUSTRALIA = "au"
	JD_ITALY     = "it"
	JD_SPAIN     = "es"
	JD_ROMANIA   = "ro"
	JD_SWEDEN    = "se"
)

type (
	ClientConfig struct {
		Tls             *tls.Config
		ApplicationName string
		ApplicationKey  string
		Rest            bool
		Jurisdiction    Jurisdiction
	}

	Client struct {
		HttpClient   *http.Client
		Config       *ClientConfig
		SessionToken string
	}

	SessionResponse struct {
		SessionToken string      `json:"sessionToken"`
		LoginStatus  LoginStatus `json:"loginStatus"`
	}

	SessionStatusResponse struct {
		Token   string             `json:"token"`
		Product string             `json:"product"`
		Status  SessionStatus      `json:"status"`
		Error   SessionStatusError `json:"error"`
	}
)

func NewClient(config *ClientConfig) *Client {
	return &Client{Config: config}
}

func (client *Client) Do(req *http.Request) (*http.Response, error) {
	if client.HttpClient == nil {
		return nil, errors.New("client not initialised: please resume or create a new session")
	}

	req.Header.Add("X-Authentication", client.SessionToken)
	req.Header.Add("X-Application", client.Config.ApplicationKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("content-type", "application/json")

	return client.HttpClient.Do(req)
}

func (client *Client) Login(username string, password string) (*SessionResponse, error) {
	postUrl := url.URL{Path: "https://identitysso-cert.betfair.com/api/certlogin"}
	q := postUrl.Query()
	q.Set("username", username)
	q.Set("password", password)

	postUrl.RawQuery = q.Encode()

	req, _ := http.NewRequest("POST", postUrl.RequestURI(), nil)
	req.SetBasicAuth(username, password)

	clientConfig := client.Config

	req.Header.Add("X-Application", clientConfig.ApplicationName)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: clientConfig.Tls,
		},
	}

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		client.HttpClient = nil
		return nil, err
	}

	json := &SessionResponse{}
	if err := ReadJson(resp, &json); err != nil {
		client.HttpClient = nil
		return json, err
	}

	if json.LoginStatus != LS_SUCCESS {
		client.HttpClient = nil
		return json, errors.New(string(json.LoginStatus))
	}

	client.SessionToken = json.SessionToken
	return json, nil
}

func (client *Client) Logout() (*SessionStatusResponse, error) {
	postUrl := url.URL{Path: client.Config.GetIdentityUrl() + "/api/logout"}

	req, _ := http.NewRequest("POST", postUrl.RequestURI(), nil)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	json := &SessionStatusResponse{}
	if err := ReadJson(resp, &json); err != nil {
		return json, err
	}

	if json.Error != "" {
		return json, errors.New(string(json.Error))
	}

	client.HttpClient = nil
	client.SessionToken = ""

	return json, nil
}

func (client *Client) Resume(sessionToken string) (*SessionStatusResponse, error) {
	client.HttpClient = &http.Client{}
	client.SessionToken = sessionToken

	return client.KeepAlive()
}

func (client *Client) KeepAlive() (*SessionStatusResponse, error) {
	keepAliveUrl := url.URL{Path: client.Config.GetIdentityUrl() + "/api/keepAlive"}
	req, _ := http.NewRequest("POST", keepAliveUrl.RequestURI(), nil)

	resp, err := client.Do(req)
	if err != nil {
		client.HttpClient = nil
		client.SessionToken = ""
		return nil, err
	}

	json := &SessionStatusResponse{}
	if err := ReadJson(resp, &json); err != nil {
		client.HttpClient = nil
		client.SessionToken = ""
		return json, err
	}

	if json.Error != "" {
		client.HttpClient = nil
		client.SessionToken = ""
		return json, errors.New(string(json.Error))
	}

	return json, nil
}

func (config *ClientConfig) GetIdentityUrl() string {
	jurisdiction := (config.Jurisdiction)
	if jurisdiction == "" {
		jurisdiction = JD_GLOBAL
	}

	return fmt.Sprintf("https://identitysso.betfair.%v", jurisdiction)
}

func (client *Client) GetStream(config *StreamingClientConfig) (*StreamingClient, error) {
	if client.HttpClient == nil {
		return nil, errors.New("client not initialised: please resume or create a new session")
	}

	sc := NewStreamingClient(config)

	clientConfig := client.Config
	return sc, sc.Authenticate(clientConfig.Tls, clientConfig.ApplicationKey, client.SessionToken)
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
	if client.Config.Rest {
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
		ID:      int(rand.Uint32()),
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
