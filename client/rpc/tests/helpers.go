package tests

import (
	"testing"

	"github.com/LonecastSystems/betfair-go/client/rpc"
)

var sessionKey, appKey = "", ""

func NewTestRpcClient(t *testing.T) *rpc.RpcClient {
	if sessionKey == "" || appKey == "" {
		t.Skip("Invalid credentials")
	}

	return rpc.NewRpcClient(sessionKey, appKey)
}
