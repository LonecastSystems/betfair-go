package betfair

type DeveloperApp struct {
	AppName string `json:"appName"`
	AppID   int64  `json:"appId"`
}

func (client *Client) GetDeveloperAppKeys() ([]DeveloperApp, error) {
	json := []DeveloperApp{}

	return json, client.GetAccounts("getDeveloperAppKeys", nil, &json)
}
