package helpers

import (
	"crypto/tls"
	"crypto/x509"
	"os"
)

func GetTLSConfig(pemFilePath string, certFilePath string, keyFilePath string) (*tls.Config, error) {
	caCert, err := os.ReadFile(pemFilePath)
	if err != nil {
		return nil, err
	}

	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	clientTLSCert, err := tls.LoadX509KeyPair(certFilePath, keyFilePath)
	if err != nil {
		return nil, err
	}

	certPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, err
	}

	return &tls.Config{
		RootCAs:      certPool,
		Certificates: []tls.Certificate{clientTLSCert},
	}, nil
}
