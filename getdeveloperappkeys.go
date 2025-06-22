package betfair

import "context"

type DeveloperApp struct {
	AppName string `json:"appName"`
	AppID   int64  `json:"appId"`
}

func (client *Client) GetDeveloperAppKeys(ctx context.Context) ([]DeveloperApp, error) {
	json := []DeveloperApp{}

	return json, client.GetAccounts(ctx, "getDeveloperAppKeys", nil, &json)
}
