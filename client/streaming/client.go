package streaming

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io"
	"math/rand/v2"
	"net/http"

	"github.com/LonecastSystems/betfair-go/client/common"
)

const betfairUrl = "stream-api.betfair.com:443"

type (
	StreamingClient struct {
		Client     *common.JsonClient
		Connection *tls.Conn
		ReadChunks int
	}
)

func CreateClient(sessionToken string, app_key string, readBytes int) *StreamingClient {
	return &StreamingClient{Client: common.CreateClient(sessionToken, app_key), ReadChunks: readBytes}
}

func (client *StreamingClient) Do(req *http.Request) (*http.Response, error) {
	return client.Client.Do(req)
}

func (client *StreamingClient) Login(config *tls.Config) (err error) {
	client.Connection, err = tls.Dial("tcp", betfairUrl, config)
	if err != nil {
		return err
	}

	cm := &ConnectionMessage{}
	err = client.Read(cm)
	if err != nil {
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

func (client *StreamingClient) Read(response any) (err error) {
	buffer := bytes.NewBuffer(nil)
	for {
		chunk := make([]byte, client.ReadChunks)
		read, err := client.Connection.Read(chunk)
		if err != nil && err != io.EOF {
			return err
		}

		buffer.Write(chunk[:read])

		if err == io.EOF || read < client.ReadChunks {
			break
		}
	}

	return json.Unmarshal(buffer.Bytes(), &response)
}

func (client *StreamingClient) Write(request any, isRequest bool) (err error) {
	bytes, err := json.Marshal(request)
	if err != nil {
		return err
	}

	bytes = append(bytes, []byte("\r\n")...)
	if _, err = client.Connection.Write(bytes); err != nil {
		return err
	}

	if isRequest {
		status := &StatusMessage{}
		client.Read(status)

		if status.ErrorCode != "" {
			return errors.New(status.ErrorCode)
		}
	}

	return nil
}

func (client *StreamingClient) ReadStream(reads chan<- any) (err error) {
	dec := json.NewDecoder(client.Connection)

	for dec.More() {
		var x any
		err := dec.Decode(&x)
		if err != nil && err != io.EOF {
			return err
		}

		reads <- x
	}

	return nil
}

func (client *StreamingClient) WriteStream(reads chan any) (err error) {
	dec := json.NewEncoder(client.Connection)

	for x := range reads {
		err := dec.Encode(x)
		if err != nil {
			return err
		}

		reads <- x
	}

	return nil
}
