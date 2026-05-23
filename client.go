package betfair

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/url"
)

// LoginStatus is returned by the Non-Interactive (bot) login and Interactive Login API endpoints.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687915/Non-Interactive+bot+login#LoginStatus
type LoginStatus string

const (
	LS_SUCCESS                                 LoginStatus = "SUCCESS"                                 // Successful login.
	LS_INVALID_USERNAME_OR_PASSWORD            LoginStatus = "INVALID_USERNAME_OR_PASSWORD"            // The username or password are invalid.
	LS_ACCOUNT_NOW_LOCKED                      LoginStatus = "ACCOUNT_NOW_LOCKED"                      // The account was just locked.
	LS_ACCOUNT_ALREADY_LOCKED                  LoginStatus = "ACCOUNT_ALREADY_LOCKED"                  // The account is already locked.
	LS_PENDING_AUTH                            LoginStatus = "PENDING_AUTH"                            // Pending authentication.
	LS_TELBET_TERMS_CONDITIONS_NA              LoginStatus = "TELBET_TERMS_CONDITIONS_NA"              // Telbet terms and conditions rejected.
	LS_DUPLICATE_CARDS                         LoginStatus = "DUPLICATE_CARDS"                         // Duplicate cards.
	LS_SECURITY_QUESTION_WRONG_3X              LoginStatus = "SECURITY_QUESTION_WRONG_3X"              // The user has entered wrong the security answer 3 times.
	LS_KYC_SUSPEND                             LoginStatus = "KYC_SUSPEND"                             // KYC suspended.
	LS_SUSPENDED                               LoginStatus = "SUSPENDED"                               // The account is suspended.
	LS_CLOSED                                  LoginStatus = "CLOSED"                                  // The account is closed.
	LS_SELF_EXCLUDED                           LoginStatus = "SELF_EXCLUDED"                           // The account has been self-excluded.
	LS_INVALID_CONNECTIVITY_TO_REGULATOR_DK    LoginStatus = "INVALID_CONNECTIVITY_TO_REGULATOR_DK"    // The DK regulator cannot be accessed due to some internal problems in the system behind or in at regulator; timeout cases included.
	LS_NOT_AUTHORIZED_BY_REGULATOR_DK          LoginStatus = "NOT_AUTHORIZED_BY_REGULATOR_DK"          // The user identified by the given credentials is not authorized in the DK's jurisdictions due to the regulators' policies. Ex: the user for which this session should be created is not allowed to act(play, bet) in the DK's jurisdiction.
	LS_INVALID_CONNECTIVITY_TO_REGULATOR_IT    LoginStatus = "INVALID_CONNECTIVITY_TO_REGULATOR_IT"    // The IT regulator cannot be accessed due to some internal problems in the system behind or in at regulator; timeout cases included.
	LS_NOT_AUTHORIZED_BY_REGULATOR_IT          LoginStatus = "NOT_AUTHORIZED_BY_REGULATOR_IT"          // The user identified by the given credentials is not authorized in the IT's jurisdictions due to the regulators' policies. Ex: the user for which this session should be created is not allowed to act(play, bet) in the IT's jurisdiction.
	LS_SECURITY_RESTRICTED_LOCATION            LoginStatus = "SECURITY_RESTRICTED_LOCATION"            // The account is restricted due to security concerns.
	LS_BETTING_RESTRICTED_LOCATION             LoginStatus = "BETTING_RESTRICTED_LOCATION"             // The account is accessed from a location where betting is restricted.
	LS_TRADING_MASTER                          LoginStatus = "TRADING_MASTER"                          // Trading Master Account.
	LS_TRADING_MASTER_SUSPENDED                LoginStatus = "TRADING_MASTER_SUSPENDED"                // Suspended Trading Master Account.
	LS_AGENT_CLIENT_MASTER                     LoginStatus = "AGENT_CLIENT_MASTER"                     // Agent Client Master.
	LS_AGENT_CLIENT_MASTER_SUSPENDED           LoginStatus = "AGENT_CLIENT_MASTER_SUSPENDED"           // Suspended Agent Client Master.
	LS_DANISH_AUTHORIZATION_REQUIRED           LoginStatus = "DANISH_AUTHORIZATION_REQUIRED"           // Danish authorization required.
	LS_SPAIN_MIGRATION_REQUIRED                LoginStatus = "SPAIN_MIGRATION_REQUIRED"                // Spain migration required.
	LS_DENMARK_MIGRATION_REQUIRED              LoginStatus = "DENMARK_MIGRATION_REQUIRED"              // Denmark migration required.
	LS_SPANISH_TERMS_ACCEPTANCE_REQUIRED       LoginStatus = "SPANISH_TERMS_ACCEPTANCE_REQUIRED"       // The latest Spanish terms and conditions version must be accepted. You must login to the website to accept the new conditions.
	LS_ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED    LoginStatus = "ITALIAN_CONTRACT_ACCEPTANCE_REQUIRED"    // The latest Italian contract version must be accepted. You must login to the website to accept the new conditions.
	LS_CERT_AUTH_REQUIRED                      LoginStatus = "CERT_AUTH_REQUIRED"                      // Certificate required or certificate present but could not authenticate with it. Please check that the correct file path is specified and ensure you are entering the correct password.
	LS_CHANGE_PASSWORD_REQUIRED                LoginStatus = "CHANGE_PASSWORD_REQUIRED"                // Change password required.
	LS_PERSONAL_MESSAGE_REQUIRED               LoginStatus = "PERSONAL_MESSAGE_REQUIRED"               // Personal message required for the user.
	LS_INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED LoginStatus = "INTERNATIONAL_TERMS_ACCEPTANCE_REQUIRED" // The latest international terms and conditions must be accepted prior to logging in.
	LS_EMAIL_LOGIN_NOT_ALLOWED                 LoginStatus = "EMAIL_LOGIN_NOT_ALLOWED"                 // This account has not opted in to log in with the email.
	LS_MULTIPLE_USERS_WITH_SAME_CREDENTIAL     LoginStatus = "MULTIPLE_USERS_WITH_SAME_CREDENTIAL"     // There is more than one account with the same credential.
	LS_ACCOUNT_PENDING_PASSWORD_CHANGE         LoginStatus = "ACCOUNT_PENDING_PASSWORD_CHANGE"         // The account must undergo password recovery to reactivate via https://identitysso.betfair.com/view/recoverpassword.
	LS_TEMPORARY_BAN_TOO_MANY_REQUESTS         LoginStatus = "TEMPORARY_BAN_TOO_MANY_REQUESTS"         // The limit for successful login requests per minute has been exceeded. New login attempts will be banned for 20 minutes.
	LS_ITALIAN_PROFILING_ACCEPTANCE_REQUIRED   LoginStatus = "ITALIAN_PROFILING_ACCEPTANCE_REQUIRED"   // You must login to the website to accept the new conditions.
	LS_AUTHORIZED_ONLY_FOR_DOMAIN_RO           LoginStatus = "AUTHORIZED_ONLY_FOR_DOMAIN_RO"           // You are attempting to login to the Betfair Romania domain with a non .ro account.
	LS_AUTHORIZED_ONLY_FOR_DOMAIN_SE           LoginStatus = "AUTHORIZED_ONLY_FOR_DOMAIN_SE"           // You are attempting to login to the Betfair Swedish domain with a non .se account.
	LS_SWEDEN_NATIONAL_IDENTIFIER_REQUIRED     LoginStatus = "SWEDEN_NATIONAL_IDENTIFIER_REQUIRED"     // You must provided your Swedish National identifier via Betfair.se before proceeding.
	LS_SWEDEN_BANK_ID_VERIFICATION_REQUIRED    LoginStatus = "SWEDEN_BANK_ID_VERIFICATION_REQUIRED"    // You must provided your Swedish bank id via Betfair.se before proceeding.
	LS_ACTIONS_REQUIRED                        LoginStatus = "ACTIONS_REQUIRED"                        // You must login to https://www.betfair.com to provide the missing information.
	LS_INPUT_VALIDATION_ERROR                  LoginStatus = "INPUT_VALIDATION_ERROR"                  // There is a problem with the data validity contained within the request. Please check that the request (including headers) is in the correct format.
	LS_MIGRATION_REQUIRED                      LoginStatus = "MIGRATION_REQUIRED"                      // Brazil customers only - you must login to betfair.bet.br to migrate your account (from 1st January 2025).
	LS_TERMS_AND_CONDITIONS                    LoginStatus = "TERMS_AND_CONDITIONS"                    // Brazil customers only - you must login to betfair.bet.br to accept the new terms and conditions (from 1st January 2025).
	LS_CONTACT_VERIFICATION_REQUIRED           LoginStatus = "CONTACT_VERIFICATION_REQUIRED"           // You must login via Betfair website (https://www.betfair.com) and complete KYC document verification.
)

// SessionStatus is returned by Keep Alive and Logout operations.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687869/Login+Session+Management#SessionStatus
type SessionStatus string

const (
	SS_SUCCESS SessionStatus = "SUCCESS" // Keep Alive or Logout succeeded.
	SS_FAIL    SessionStatus = "FAIL"    // Keep Alive or Logout failed.
)

// SessionStatusError is returned when SessionStatus is FAIL.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687869/Login+Session+Management#SessionStatusError
type SessionStatusError string

const (
	SS_ERR_INPUT_VALIDATION_ERROR SessionStatusError = "INPUT_VALIDATION_ERROR" // There is a problem with the data validity contained within the request.
	SS_ERR_INTERNAL_ERROR         SessionStatusError = "INTERNAL_ERROR"         // An internal error occurred processing the request.
	SS_ERR_NO_SESSION             SessionStatusError = "NO_SESSION"             // The session token header is missing, invalid, or the session has expired.
)

// Jurisdiction selects the Betfair identity and API domain suffix.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687915/Non-Interactive+bot+login#Other-Jurisdictions
type Jurisdiction string

const (
	JD_GLOBAL    Jurisdiction = "com" // Global Exchange (.com).
	JD_AUSTRALIA Jurisdiction = "au"  // Australian & New Zealand Exchange (.com.au).
	JD_ITALY     Jurisdiction = "it"  // Italian Exchange (.it).
	JD_SPAIN     Jurisdiction = "es"  // Spanish Exchange (.es).
	JD_ROMANIA   Jurisdiction = "ro"  // Romanian Exchange (.ro).
	JD_SWEDEN    Jurisdiction = "se"  // Swedish Exchange (.se).
)

// ClientConfig configures the Betfair API client.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687915/Non-Interactive+bot+login#Certificate-Login-Interface-Details
type (
	ClientConfig struct {
		Tls             *tls.Config  // TLS client configuration for certificate-based non-interactive login.
		ApplicationName string       // Optional application name identifier.
		ApplicationKey  string       // Application key sent in the X-Application header.
		Rest            bool         // Use REST API instead of JSON-RPC when true.
		Jurisdiction    Jurisdiction // Betfair identity and API domain jurisdiction suffix.
	}

	Client struct {
		HttpClient   *http.Client
		Config       *ClientConfig
		SessionToken string
	}

	// LoginResponse is returned by the Non-Interactive (bot) login endpoint.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687915/Non-Interactive+bot+login#Response
	LoginResponse struct {
		SessionToken string      `json:"sessionToken"` // Session token created at login.
		LoginStatus  LoginStatus `json:"loginStatus"`  // Login status (SUCCESS or failure reason).
	}

	// SessionStatusResponse is returned by Keep Alive and Logout operations.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687869/Login+Session+Management#Response-structure
	SessionStatusResponse struct {
		Token   string             `json:"token"`   // Session token.
		Product string             `json:"product"` // Application key passed as header.
		Status  SessionStatus      `json:"status"`  // SUCCESS or FAIL.
		Error   SessionStatusError `json:"error"`   // Error code if status is FAIL.
	}

	LogoutResponse    = SessionStatusResponse
	KeepAliveResponse = SessionStatusResponse
	ResumeResponse    = SessionStatusResponse
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

// Login submits a HTTP POST request to certlogin with username/password and certificate.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687915/Non-Interactive+bot+login#Details-of-a-Login-Request
func (client *Client) Login(ctx context.Context, username string, password string) (*LoginResponse, error) {
	loginURL, err := url.Parse(client.Config.GetCertIdentityUrl() + "/api/certlogin")
	if err != nil {
		return nil, err
	}
	q := loginURL.Query()
	q.Set("username", username)
	q.Set("password", password)
	loginURL.RawQuery = q.Encode()
	req, err := http.NewRequestWithContext(ctx, "POST", loginURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("X-Application", client.Config.ApplicationKey)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client.HttpClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: client.Config.Tls,
		},
	}

	resp, err := client.HttpClient.Do(req)
	if err != nil {
		client.HttpClient = nil
		return nil, err
	}

	json := &LoginResponse{}
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

// Logout terminates your existing session.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687869/Login+Session+Management#Logout
func (client *Client) Logout(ctx context.Context) (*LogoutResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", client.Config.GetIdentityUrl()+"/api/logout", nil)
	if err != nil {
		return nil, err
	}

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

// Resume extends the session with an existing token by calling Keep Alive.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687869/Login+Session+Management#Keep-Alive
func (client *Client) Resume(ctx context.Context, sessionToken string) (*ResumeResponse, error) {
	client.HttpClient = &http.Client{}
	client.SessionToken = sessionToken

	return client.KeepAlive(ctx)
}

// KeepAlive extends the session timeout period.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687869/Login+Session+Management#Keep-Alive
func (client *Client) KeepAlive(ctx context.Context) (*KeepAliveResponse, error) {
	req, err := http.NewRequestWithContext(ctx, "POST", client.Config.GetIdentityUrl()+"/api/keepAlive", nil)
	if err != nil {
		return nil, err
	}

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

func (config *ClientConfig) jurisdiction() Jurisdiction {
	if config.Jurisdiction == "" {
		return JD_GLOBAL
	}
	return config.Jurisdiction
}

func (config *ClientConfig) GetIdentityUrl() string {
	return fmt.Sprintf("https://identitysso.betfair.%v", config.jurisdiction())
}

func (config *ClientConfig) GetCertIdentityUrl() string {
	return fmt.Sprintf("https://identitysso-cert.betfair.%v", config.jurisdiction())
}

func (config *ClientConfig) GetApiUrl() string {
	jurisdiction := (config.Jurisdiction)
	if jurisdiction == JD_AUSTRALIA {
		return fmt.Sprintf("https://api.betfair.com.%v", jurisdiction)
	}

	return "https://api.betfair.com"
}

func (client *Client) GetStream(config *StreamingClientConfig) (*StreamingClient, error) {
	if client.HttpClient == nil {
		return nil, errors.New("client not initialised: please resume or create a new session")
	}

	sc := NewStreamingClient(config)

	return sc, sc.Authenticate(client.Config.Tls, client.Config.ApplicationKey, client.SessionToken)
}

const (
	api_account   = "account"
	api_betting   = "betting"
	api_heartbeat = "heartbeat"
	api_scores    = "scores"
)

func (client *Client) GetAccounts(ctx context.Context, method string, params any, response any) error {
	return client.get(ctx, api_account, method, params, response)
}

func (client *Client) GetSports(ctx context.Context, method string, params any, response any) error {
	return client.get(ctx, api_betting, method, params, response)
}

func (client *Client) GetHeartbeats(ctx context.Context, method string, params any, response any) error {
	return client.get(ctx, api_heartbeat, method, params, response)
}

func (client *Client) GetScores(ctx context.Context, method string, params any, response any) error {
	return client.getRPC(ctx, api_scores, method, params, response) //Only supported by RPC for now.
}

func (client *Client) get(ctx context.Context, api string, method string, params any, response any) error {
	if client.Config.Rest {
		return client.getRest(ctx, api, method, params, response)
	} else {
		return client.getRPC(ctx, api, method, params, response)
	}
}

// JsonAPINGExceptionErrorCode is returned in APINGException when a Betting API operation fails.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687450/Betting+Exceptions#APINGException
type JsonAPINGExceptionErrorCode string

const (
	EC_UNEXPECTED_ERROR            JsonAPINGExceptionErrorCode = "UNEXPECTED_ERROR"            // An unexpected internal error occurred that prevented successful request processing.
	EC_INVALID_INPUT_DATA          JsonAPINGExceptionErrorCode = "INVALID_INPUT_DATA"          // The data input is invalid. A specific description is returned via errorDetails. Please note: if the number of placeOrders, updateOrders, replaceOrders, or cancelOrders instructions exceeds the documented limit you will also receive this error.
	EC_INVALID_SESSION_INFORMATION JsonAPINGExceptionErrorCode = "INVALID_SESSION_INFORMATION" // The session token hasn't been provided, is invalid or has expired. Login again to create a new session.
	EC_INVALID_APP_KEY             JsonAPINGExceptionErrorCode = "INVALID_APP_KEY"             // The application key passed is invalid or is not present.
	EC_SERVICE_BUSY                JsonAPINGExceptionErrorCode = "SERVICE_BUSY"                // The service is currently too busy to service this request.
	EC_TIMEOUT_ERROR               JsonAPINGExceptionErrorCode = "TIMEOUT_ERROR"               // The Internal call to downstream service timed out. Please note: If a TIMEOUT error occurs on a placeOrders/replaceOrders request, you should check listCurrentOrders to verify the status of your bets before placing further orders. Please Note: Timeouts will occur after 5 seconds of attempting to process the bet but please allow up to 15 seconds for a timed out order to appear. After this time any unprocessed bets will automatically be Lapsed and no longer be available on the Exchange.
	EC_NO_SESSION                  JsonAPINGExceptionErrorCode = "NO_SESSION"                  // A session token header ('X-Authentication') has not been provided in the request.
	EC_NO_APP_KEY                  JsonAPINGExceptionErrorCode = "NO_APP_KEY"                  // An application key header ('X-Application') has not been provided in the request.
	EC_TOO_MANY_REQUESTS           JsonAPINGExceptionErrorCode = "TOO_MANY_REQUESTS"           // There are too many pending (in-flight) requests e.g. a listMarketBook with Order/Match projections is limited to 3 concurrent requests. The error also applies to listCurrentOrders, listMarketProfitAndLoss and listClearedOrders if you have 3 or more requests currently in execution.
	EC_SERVICE_UNAVAILABLE         JsonAPINGExceptionErrorCode = "SERVICE_UNAVAILABLE"         // The requested service is unavailable.
	EC_REQUEST_SIZE_EXCEEDS_LIMIT  JsonAPINGExceptionErrorCode = "REQUEST_SIZE_EXCEEDS_LIMIT"  // The request exceeds the request size limit. Requests are limited to a total of 250 betId's/marketId's (or a combination of both).
	EC_TOO_MUCH_DATA               JsonAPINGExceptionErrorCode = "TOO_MUCH_DATA"               // The operation requested too much data, exceeding the Market Data Request Limits. You must adjust your request parameters to stay with the documented limits.
	EC_ACCESS_DENIED               JsonAPINGExceptionErrorCode = "ACCESS_DENIED"               // The calling client is not permitted to perform the specific action e.g. they have an App Key restriction in place or attempting to place a bet from a restricted jurisdiction.
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

func (client *Client) getRPC(ctx context.Context, api string, method string, params any, response any) error {
	query := JsonRPC{
		JsonRPC: "2.0",
		Method:  fmt.Sprintf("%v/v1.0/%v", apis[api], method),
		Params:  params,
		ID:      int(rand.UintN(16)),
	}

	body, err := json.Marshal(&query)
	if err != nil {
		return err
	}

	apiUrl := fmt.Sprintf("%v/exchange/%v/json-rpc/v1/", client.Config.GetApiUrl(), api)
	req, err := http.NewRequestWithContext(ctx, "POST", apiUrl, bytes.NewBuffer(body))
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

func (client *Client) getRest(ctx context.Context, api string, method string, params any, response any) error {
	body, err := json.Marshal(&params)
	if err != nil {
		return err
	}

	apiUrl := fmt.Sprintf("%v/exchange/%v/rest/v1.0/%v/", client.Config.GetApiUrl(), api, method)
	req, err := http.NewRequestWithContext(ctx, "POST", apiUrl, bytes.NewBuffer(body))
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
