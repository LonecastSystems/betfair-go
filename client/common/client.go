package common

import (
	"crypto/tls"
	"net/http"
	"net/url"

	"github.com/LonecastSystems/betfair-go/helpers"
)

type (
	JsonClient struct {
		Client         *http.Client
		ApplicationKey string
		SessionToken   string
	}
)

func CreateClient(sessionToken string, app_key string) *JsonClient {
	return &JsonClient{Client: &http.Client{}, SessionToken: sessionToken, ApplicationKey: app_key}
}

func (client *JsonClient) Do(req *http.Request) (*http.Response, error) {
	req.Header.Add("X-Authentication", client.SessionToken)
	req.Header.Add("X-Application", client.ApplicationKey)
	req.Header.Add("Accept", "application/json")

	return client.Client.Do(req)
}

func (jsonClient *JsonClient) Login(tls *tls.Config, apiKey string, applicationName string, username string, password string) (response *http.Response, err error) {
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

func (jsonClient *JsonClient) Logout() (jsonResponse SessionLogoutResponse, response *http.Response, err error) {
	postUrl := url.URL{Path: "https://identitysso.betfair.com/api/logout"}

	req, _ := http.NewRequest("POST", postUrl.RequestURI(), nil)

	json := SessionLogoutResponse{}

	resp, err := jsonClient.Do(req)
	if err != nil {
		return json, resp, err
	}

	if err := helpers.ReadJson(resp, &json); err != nil {
		return json, resp, err
	}

	jsonClient.SessionToken = ""
	return json, resp, err
}
