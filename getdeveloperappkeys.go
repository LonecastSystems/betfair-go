package betfair

import "context"

// DeveloperApp describes developer/vendor specific application.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687852/Accounts+TypeDefinitions#DeveloperApp
type DeveloperApp struct {
	AppName     string                `json:"appName"`     // The unique name of the application.
	AppID       int64                 `json:"appId"`       // A unique id of this application.
	AppVersions []DeveloperAppVersion `json:"appVersions"` // The application versions (including application keys).
}

// DeveloperAppVersion describes a version of an external application.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687852/Accounts+TypeDefinitions#DeveloperAppVersion
type DeveloperAppVersion struct {
	Owner                string `json:"owner"`                  // The user who owns the specific version of the application.
	VersionID            int64  `json:"versionId"`              // The unique Id of the application version.
	Version              string `json:"version"`                // The version identifier string such as 1.0, 2.0. Unique for a given application.
	ApplicationKey       string `json:"applicationKey"`         // The unqiue application key associated with this application version.
	DelayData            bool   `json:"delayData"`              // Indicates whether the data exposed by platform services as seen by this application key is delayed or realtime.
	SubscriptionRequired bool   `json:"subscriptionRequired"`   // Indicates whether the application version needs explicit subscription.
	OwnerManaged         bool   `json:"ownerManaged"`           // Indicates whether the application version needs explicit management by the software owner. A value of false indicates, this is a version meant for personal developer use.
	Active               bool   `json:"active"`                 // Indicates whether the application version is currently active.
	VendorID             string `json:"vendorId,omitempty"`     // Public unique string provided to the Vendor that they can use to pass to the Betfair API in order to identify themselves.
	VendorSecret         string `json:"vendorSecret,omitempty"` // Private unique string provided to the Vendor that they pass with certain calls to confirm their identity. Linked to a particular App Key.
}

type GetDeveloperAppKeysResult = DeveloperApp

// GetDeveloperAppKeys gets all application keys owned by the given developer/vendor.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687893/getDeveloperAppKeys
func (client *Client) GetDeveloperAppKeys(ctx context.Context) ([]GetDeveloperAppKeysResult, error) {
	json := []GetDeveloperAppKeysResult{}

	return json, client.GetAccounts(ctx, "getDeveloperAppKeys", nil, &json)
}
