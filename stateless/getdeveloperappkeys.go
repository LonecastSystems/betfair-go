package stateless

type DeveloperApp struct {
	AppName string `json:"appName"`
	AppID   int64  `json:"appId"`
}

func (client *StatelessClient) GetDeveloperAppKeys() ([]DeveloperApp, error) {
	json := []DeveloperApp{}

	if err := GetAccounts(client, "getDeveloperAppKeys", nil, &json); err != nil {
		return []DeveloperApp{}, err
	}

	return json, nil
}
