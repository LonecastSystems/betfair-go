package helpers

import (
	"crypto/tls"
	"crypto/x509"
)

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
