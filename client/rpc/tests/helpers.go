package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client/rpc"
)

var sessionKey, appKey = "", ""

func CreateClient(t *testing.T) rpc.RpcBettingClient {
	if sessionKey == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	return rpc.RpcBettingClient(rpc.CreateClient(sessionKey, appKey))
}
