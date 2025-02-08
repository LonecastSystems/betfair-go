package betfair

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func ReadJson(res *http.Response, response any) error {
	defer res.Body.Close()

	if body, err := io.ReadAll(res.Body); err == nil {
		return json.Unmarshal(body, &response)
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

func GetTLSConfig(certFilePath string, keyFilePath string) (*tls.Config, error) {
	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	clientTLSCert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientTLSCert},
	}, nil
}
