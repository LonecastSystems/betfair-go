package rpc

import (
	"net/http"
	"time"
)

type (
	JsonRpcClient struct {
		Client         *http.Client
		ApplicationKey string
		SessionToken   string
	}

	RpcBettingClient interface {
		ListCompetitions(params RPCParams) ([]CompetitionResult, error)
		ListEventTypes(params RPCParams) ([]EventTypeResult, error)
		ListEvents(params RPCParams) ([]EventResult, error)
		ListMarketTypes(params RPCParams) ([]MarketTypeResult, error)
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

// Requests
type (
	CompetitionResult struct {
		Competition struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"competition"`
		MarketCount       int    `json:"marketCount"`
		CompetitionRegion string `json:"competitionRegion"`
	}

	EventTypeResult struct {
		EventType struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"eventType"`
		MarketCount int `json:"marketCount"`
	}

	EventResult struct {
		Event struct {
			ID          string    `json:"id"`
			Name        string    `json:"name"`
			CountryCode string    `json:"countryCode"`
			Timezone    string    `json:"timezone"`
			OpenDate    time.Time `json:"openDate"`
		} `json:"event"`
		MarketCount int `json:"marketCount"`
	}

	MarketTypeResult struct {
		MarketType  string `json:"marketType"`
		MarketCount int    `json:"marketCount"`
	}
)

// RPC
type (
	JsonRpcResponse struct {
		JsonRPC string      `json:"jsonrpc"`
		Result  interface{} `json:"result"`
		Error   JsonError   `json:"error,omitempty"`
		ID      int         `json:"id"`
	}

	JsonError struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
		Data    struct {
			APINGException struct {
				RequestUUID  string `json:"requestUUID"`
				ErrorCode    string `json:"errorCode"`
				ErrorDetails string `json:"errorDetails"`
			} `json:"APINGException"`
			Exceptionname string `json:"exceptionname"`
		} `json:"data"`
	}

	JsonRPC struct {
		JsonRPC string    `json:"jsonrpc"`
		Method  string    `json:"method"`
		Params  RPCParams `json:"params"`
		ID      int       `json:"id"`
	}

	RPCParams struct {
		Filter MarketFilter `json:"filter"`
	}
)

// Filters
type (
	MarketFilter struct {
		TextQuery          string    `json:"textQuery,omitempty"`
		EventTypeIds       []string  `json:"eventTypeIds,omitempty"`
		EventIds           []string  `json:"eventIds,omitempty"`
		CompetitionIds     []string  `json:"competitionIds,omitempty"`
		MarketIds          []string  `json:"marketIds,omitempty"`
		Venues             []string  `json:"venues,omitempty"`
		BspOnly            bool      `json:"bspOnly,omitempty"`
		TurnInPlayEnabled  bool      `json:"turnInPlayEnabled,omitempty"`
		InPlayOnly         bool      `json:"inPlayOnly,omitempty"`
		MarketBettingTypes []string  `json:"marketBettingTypes,omitempty"`
		MarketTypeCodes    []string  `json:"marketTypeCodes,omitempty"`
		MarketCountries    []string  `json:"marketCountries"`
		MarketStartTime    TimeRange `json:"marketStartTime,omitempty"`
		WithOrders         []string  `json:"withOrders,omitempty"`
		RaceTypes          []string  `json:"raceTypes,omitempty"`
	}

	TimeRange struct {
		From string `json:"from,omitempty"`
		To   string `json:"to,omitempty"`
	}
)
