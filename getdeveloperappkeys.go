package betfair

type DeveloperApp struct {
	AppName string `json:"appName"`
	AppID   int64  `json:"appId"`
}

func (client *Client) GetDeveloperAppKeys() ([]DeveloperApp, error) {
	json := []DeveloperApp{}

	if err := client.GetAccounts("getDeveloperAppKeys", nil, &json); err != nil {
		return json, err
	}

	return json, nil
}
