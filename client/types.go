package client

import "net/http"

type (
	JsonRpcClient struct {
		Client         *http.Client
		ApplicationKey string
		SessionToken   string
	}
)

type (
	SessionResponse struct {
		SessionToken string `json:"sessionToken"`
		LoginStatus  string `json:"loginStatus"`
	}
	SessionLogoutResponse struct {
		Token   string `json:"token"`
		Product string `json:"product"`
		Status  string `json:"status"`
		Error   string `json:"error"`
	}
)
