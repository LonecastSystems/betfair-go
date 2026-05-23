package betfair

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand/v2"
)

// StatusMessageErrorCode categorizes Exchange Stream API status message failures.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#StatusMessage
type StatusMessageErrorCode string

const (
	SMEC_ERR_INVALID_INPUT                 StatusMessageErrorCode = "INVALID_INPUT"                 // Failure code returned when an invalid input is provided (could not deserialize the message).
	SMEC_ERR_TIMEOUT                       StatusMessageErrorCode = "TIMEOUT"                       // Failure code when a client times out (i.e. too slow sending data).
	SMEC_ERR_NO_APP_KEY                    StatusMessageErrorCode = "NO_APP_KEY"                    // Failure code returned when an application key is not found in the message.
	SMEC_ERR_INVALID_APP_KEY               StatusMessageErrorCode = "INVALID_APP_KEY"               // Failure code returned when an invalid application key is received.
	SMEC_ERR_NO_SESSION                    StatusMessageErrorCode = "NO_SESSION"                    // Failure code returned when a session token is not found in the message.
	SMEC_ERR_INVALID_SESSION_INFORMATION   StatusMessageErrorCode = "INVALID_SESSION_INFORMATION"   // Failure code returned when an invalid session token is received.
	SMEC_ERR_MAX_CONNECTION_LIMIT_EXCEEDED StatusMessageErrorCode = "MAX_CONNECTION_LIMIT_EXCEEDED" // Failure code returned when a client tries to create more connections than allowed to.
	SMEC_ERR_TOO_MANY_REQUESTS             StatusMessageErrorCode = "TOO_MANY_REQUESTS"             // Failure code is returned when a client makes too many requests within a short time period.
	SMEC_ERR_SUBSCRIPTION_LIMIT_EXCEEDED   StatusMessageErrorCode = "SUBSCRIPTION_LIMIT_EXCEEDED"   // Thrown when subscribed to more markets than allowed to - set to 200 markets by default.
	SMEC_ERR_INVALID_CLOCK                 StatusMessageErrorCode = "INVALID_CLOCK"                 // Failure code returned when an invalid clock is provided on re-subscription (check initialClk / clk supplied).
	SMEC_ERR_UNEXPECTED_ERROR              StatusMessageErrorCode = "UNEXPECTED_ERROR"              // Failure code returned when an internal error occurred on the server. Slow or unstable internet connectivity on the client side is one of the most common root causes of this issue.
	SMEC_ERR_CONNECTION_FAILED             StatusMessageErrorCode = "CONNECTION_FAILED"             // Failure code used when the client/server connection is terminated.
)

type (
	// StreamingClientConfig configures Exchange Stream API connection and subscription defaults.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#MarketSubscriptionMessage
	StreamingClientConfig struct {
		SegmentationEnabled bool   // Segmentation Enabled - allow the server to send large sets of data in segments, instead of a single block.
		ConflateMs          int    // Conflate Milliseconds - the conflation rate (looped back on initial image after validation: bounds are 0 to 120000).
		HeartbeatMs         int    // Heartbeat Milliseconds - the heartbeat rate (looped back on initial image after validation: bounds are 500 to 5000).
		InitialClk          string // Token value (received in initial MarketChangeMessage) that should be passed to resume a subscription.
		Clk                 string // Token value delta (received in MarketChangeMessage) that should be passed to resume a subscription.
		PreProduction       bool   // Connect to the integration stream endpoint instead of production.
	}

	StreamingClient struct {
		Connection *tls.Conn
		Config     *StreamingClientConfig
	}

	// StatusMessage is the response to a stream request.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687396/Exchange+Stream+API#StatusMessage
	StatusMessage struct {
		ID                   int                    `json:"id"`                   // Client generated unique id to link request with response (like json rpc).
		StatusCode           string                 `json:"statusCode"`           // The status of the last request (SUCCESS or FAILURE).
		ConnectionClosed     bool                   `json:"connectionClosed"`     // Is the connection now closed.
		ErrorCode            StatusMessageErrorCode `json:"errorCode"`            // The type of error in case of a failure.
		ErrorMessage         string                 `json:"errorMessage"`         // Additional message in case of a failure.
		ConnectionsAvailable int                    `json:"connectionsAvailable"` // The number of connections available for this account at this moment in time. Present on responses to Authentication messages only.
	}
)

type (
	ConnectionMessage struct {
		Op           string `json:"op"`
		ConnectionID string `json:"connectionId"`
	}

	AuthenticationMessage struct {
		ID      int    `json:"id"`
		Op      string `json:"op"`
		AppKey  string `json:"appKey"`
		Session string `json:"session"`
	}
)

func NewStreamingClient(config *StreamingClientConfig) (client *StreamingClient) {
	return &StreamingClient{Config: config}
}

func (config *StreamingClientConfig) GetStreamUrl() string {
	if config.PreProduction {
		return "stream-api-integration.betfair.com"
	}

	return "stream-api.betfair.com:443"
}

func (client *StreamingClient) Authenticate(config *tls.Config, applicationKey, sessionToken string) (err error) {
	if client.Connection, err = tls.Dial("tcp", client.Config.GetStreamUrl(), config); err != nil {
		return err
	}

	cm := &ConnectionMessage{}
	if err = client.Read(cm); err != nil {
		client.Connection = nil
		return err
	}

	am := AuthenticationMessage{
		ID:      int(rand.UintN(16)),
		Op:      "authentication",
		AppKey:  applicationKey,
		Session: sessionToken,
	}

	if err = client.Write(am, true); err != nil {
		client.Connection = nil
	}

	return err
}

func (client *StreamingClient) Close() error {
	if client.Connection != nil {
		return client.Connection.Close()
	}

	return nil
}

func (client *StreamingClient) Read(response any) error {
	if client.Connection == nil {
		return errors.New("connection not established: please authenticate")
	}
	if err := json.NewDecoder(client.Connection).Decode(&response); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (client *StreamingClient) Write(request any, isRequest bool) error {
	if client.Connection == nil {
		return errors.New("connection not established: please authenticate")
	}

	if err := json.NewEncoder(client.Connection).Encode(request); err != nil {
		return err
	}

	if !isRequest {
		return nil
	}

	status := &StatusMessage{}
	if err := client.Read(status); err != nil {
		return err
	}

	if status.ErrorCode != "" {
		return errors.New(string(status.ErrorCode))
	}
	if status.StatusCode != "" && status.StatusCode != "SUCCESS" {
		return fmt.Errorf("stream request failed: %s", status.StatusCode)
	}

	return nil
}
