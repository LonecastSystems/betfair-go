package betfairgo

type DeveloperApp struct {
	AppName string `json:"appName"`
	AppID   int64  `json:"appId"`
}

func (client *BetfairClient) GetDeveloperAppKeys() ([]DeveloperApp, error) {
	json := []DeveloperApp{}

	if err := client.GetAccounts("getDeveloperAppKeys", nil, &json); err != nil {
		return []DeveloperApp{}, err
	}

	return json, nil
}
