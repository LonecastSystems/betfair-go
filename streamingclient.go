package betfairgo

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"math/rand/v2"
)

type (
	StreamingClient struct {
		Connection          *tls.Conn
		SegmentationEnabled bool
		ConflateMs          int
		HeartbeatMs         int
		InitialClk          string
		Clk                 string
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

func NewStreamingClient() (client *StreamingClient) {
	return &StreamingClient{}
}

func (client *StreamingClient) Authenticate(config *tls.Config, applicationKey, sessionToken string) (err error) {
	if client.Connection, err = tls.Dial("tcp", "stream-api.betfair.com:443", config); err != nil {
		return err
	}

	cm := &ConnectionMessage{}
	if err = client.Read(cm); err != nil {
		client.Connection = nil
		return err
	}

	am := AuthenticationMessage{
		ID:      int(rand.Uint64()),
		Op:      "authentication",
		AppKey:  applicationKey,
		Session: sessionToken,
	}

	if err = client.Write(am, true); err != nil {
		client.Connection = nil
	}

	return err
}

func (client *StreamingClient) Read(response any) error {
	if client.Connection == nil {
		return errors.New("connection not established: please authenticate")
	}

	dec := json.NewDecoder(client.Connection)

	if err := dec.Decode(&response); err != nil && err != io.EOF {
		return err
	}

	return nil
}

func (client *StreamingClient) Write(request any, isRequest bool) error {
	if client.Connection == nil {
		return errors.New("connection not established: please authenticate")
	}

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
	defer close(reads)

	if connection == nil {
		return errors.New("connection not established: please authenticate")
	}

	dec := json.NewDecoder(connection)

	for dec.More() {
		var x T

		if err = dec.Decode(&x); err != nil && err != io.EOF {
			break
		}

		reads <- x
	}

	return err
}
