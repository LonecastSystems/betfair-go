package streaming

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"math/rand/v2"
	"net/http"

	"github.com/LonecastSystems/betfair-go/common"
)

type StreamingClient struct {
	Client              *common.JsonClient
	Connection          *tls.Conn
	SegmentationEnabled bool
	ConflateMs          int
	HeartbeatMs         int
	InitialClk          string
	Clk                 string
}

func NewStreamingClient(sessionToken string, app_key string) *StreamingClient {
	return &StreamingClient{Client: common.NewJsonClient(sessionToken, app_key)}
}

func (client *StreamingClient) Do(req *http.Request) (*http.Response, error) {
	return client.Client.Do(req)
}

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

	StatusMessage struct {
		ID                   int    `json:"id"`
		StatusCode           string `json:"statusCode"`
		ConnectionClosed     bool   `json:"connectionClosed"`
		ErrorCode            string `json:"errorCode"`
		ErrorMessage         string `json:"errorMessage"`
		ConnectionsAvailable int    `json:"connectionsAvailable"`
	}
)

func (client *StreamingClient) Login(config *tls.Config) (err error) {
	client.Connection, err = tls.Dial("tcp", "stream-api.betfair.com:443", config)
	if err != nil {
		return err
	}

	cm := &ConnectionMessage{}
	if err = client.Read(cm); err != nil {
		return err
	}

	am := AuthenticationMessage{
		ID:      int(rand.Uint64()),
		Op:      "authentication",
		AppKey:  client.Client.ApplicationKey,
		Session: client.Client.SessionToken,
	}

	return client.Write(am, true)
}

func (client *StreamingClient) Read(response any) error {
	dec := json.NewDecoder(client.Connection)

	if err := dec.Decode(&response); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (client *StreamingClient) Write(request any, isRequest bool) error {
	enc := json.NewEncoder(client.Connection)

	if err := enc.Encode(request); err != nil {
		return err
	}

	if isRequest {
		status := &StatusMessage{}
		if err := client.Read(status); err != nil {
			return err
		}

		if status.ErrorCode != "" {
			return errors.New(status.ErrorCode)
		}
	}

	return nil
}

func ReadStream[T any](connection *tls.Conn, reads chan<- T) (err error) {
	dec := json.NewDecoder(connection)

	for dec.More() {
		var x T

		if err := dec.Decode(&x); err != nil && err != io.EOF {
			break
		}

		reads <- x
	}

	close(reads)
	return nil
}
