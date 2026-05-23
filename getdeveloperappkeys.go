package betfair

import "context"

type DeveloperApp struct {
	AppName     string                `json:"appName"`
	AppID       int64                 `json:"appId"`
	AppVersions []DeveloperAppVersion `json:"appVersions"`
}

type DeveloperAppVersion struct {
	Owner                string `json:"owner"`
	VersionID            int64  `json:"versionId"`
	Version              string `json:"version"`
	ApplicationKey       string `json:"applicationKey"`
	DelayData            bool   `json:"delayData"`
	SubscriptionRequired bool   `json:"subscriptionRequired"`
	OwnerManaged         bool   `json:"ownerManaged"`
	Active               bool   `json:"active"`
	VendorID             string `json:"vendorId,omitempty"`
	VendorSecret         string `json:"vendorSecret,omitempty"`
}

type GetDeveloperAppKeysResult = DeveloperApp

func (client *Client) GetDeveloperAppKeys(ctx context.Context) ([]GetDeveloperAppKeysResult, error) {
	json := []GetDeveloperAppKeysResult{}

	return json, client.GetAccounts(ctx, "getDeveloperAppKeys", nil, &json)
}
