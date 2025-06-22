package betfair

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"math/rand/v2"
)

type (
	StreamingClientConfig struct {
		SegmentationEnabled bool
		ConflateMs          int
		HeartbeatMs         int
		InitialClk          string
		Clk                 string
		PreProduction       bool
	}

	StreamingClient struct {
		Connection *tls.Conn
		Config     *StreamingClientConfig
	}

	StatusMessage struct {
		ID                   int                    `json:"id"`
		StatusCode           string                 `json:"statusCode"`
		ConnectionClosed     bool                   `json:"connectionClosed"`
		ErrorCode            StatusMessageErrorCode `json:"errorCode"`
		ErrorMessage         string                 `json:"errorMessage"`
		ConnectionsAvailable int                    `json:"connectionsAvailable"`
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
		ID:      int(rand.Uint32()),
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

	return nil
}
