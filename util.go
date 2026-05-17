package betfair

import (
	"crypto/tls"
	"encoding/json"
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

func ReadStream[T any](connection *tls.Conn) chan T {
	reads := make(chan T)

	dec := json.NewDecoder(connection)

	go func() {
		defer close(reads)

		for dec.More() {
			var x T

			if err := dec.Decode(&x); err != nil && err != io.EOF {
				break
			}

			reads <- x
		}
	}()

	return reads
}

func GetTLSConfig(certFilePath string, keyFilePath string) (*tls.Config, error) {
	clientTLSCert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{clientTLSCert},
	}, nil
}

func GetTLSConfigFromBytes(certBytes, keyBytes []byte) (*tls.Config, error) {
	clientTLSCert, err := tls.X509KeyPair(certBytes, keyBytes)
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		Certificates: []tls.Certificate{clientTLSCert},
	}, nil
}
