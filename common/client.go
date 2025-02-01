package common

import (
	"crypto/tls"
	"errors"
	"net/http"
	"net/url"

	"github.com/LonecastSystems/betfair-go/helpers"
)

type SessionStatus string

const (
	KAS_SUCCESS = "SUCCESS"
	KAS_FAIL    = "FAIL"
)

type SessionError string

const (
	KAE_INPUT_VALIDATION_ERROR = "INPUT_VALIDATION_ERROR"
	KAE_INTERNAL_ERROR         = "INTERNAL_ERROR"
	KAE_NO_SESSION             = "NO_SESSION"
)

type (
	JsonClient struct {
		Client         *http.Client
		ApplicationKey string
		SessionToken   string
	}

	SessionResponse struct {
		SessionToken string `json:"sessionToken"`
		LoginStatus  string `json:"loginStatus"`
	}

	SessionStatusResponse struct {
		Token   string        `json:"token"`
		Product string        `json:"product"`
		Status  SessionStatus `json:"status"`
		Error   SessionError  `json:"error"`
	}
)

func NewJsonClient(app_key string) *JsonClient {
	return &JsonClient{Client: &http.Client{}, ApplicationKey: app_key}
}

func (client *JsonClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Authentication", client.SessionToken)
	req.Header.Add("X-Application", client.ApplicationKey)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("content-type", "application/json")

	return client.Client.Do(req)
}

func (jsonClient *JsonClient) ResumeSession(sessionToken string) (*http.Response, error) {
	jsonClient.SessionToken = sessionToken

	keepAliveUrl := url.URL{Path: "https://identitysso.betfair.com/api/keepAlive"}
	req, _ := http.NewRequest("POST", keepAliveUrl.RequestURI(), nil)

	resp, err := jsonClient.Do(req)
	if err != nil {
		jsonClient.SessionToken = ""
		return resp, err
	}

	json := SessionStatusResponse{}
	if err := helpers.ReadJson(resp, &json); err != nil {
		jsonClient.SessionToken = ""
		return resp, err
	}

	if json.Error != "" {
		jsonClient.SessionToken = ""
		return resp, errors.New(string(json.Error))
	}

	return resp, nil
}

func (jsonClient *JsonClient) NewSession(tls *tls.Config, apiKey string, applicationName string, username string, password string) (*http.Response, error) {
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
	if err != nil {
		return resp, err
	}

	json := SessionResponse{}
	if err := helpers.ReadJson(resp, &json); err != nil {
		return resp, err
	}

	jsonClient.SessionToken = json.SessionToken
	return resp, nil
}

func (jsonClient *JsonClient) ClearSession() (*http.Response, error) {
	postUrl := url.URL{Path: "https://identitysso.betfair.com/api/logout"}

	req, _ := http.NewRequest("POST", postUrl.RequestURI(), nil)

	resp, err := jsonClient.Do(req)
	if err != nil {
		return resp, err
	}

	json := SessionStatusResponse{}
	if err := helpers.ReadJson(resp, &json); err != nil {
		return resp, err
	}

	if json.Error != "" {
		return resp, errors.New(string(json.Error))
	}

	jsonClient.SessionToken = ""
	return resp, err
}
