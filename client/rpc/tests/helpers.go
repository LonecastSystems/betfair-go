package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client/rpc"
)

var sessionKey, appKey = "", ""

func CreateClient(t *testing.T) *rpc.JsonRpcClient {
	if sessionKey == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	return rpc.CreateClient(sessionKey, appKey)
}
