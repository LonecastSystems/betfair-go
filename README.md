# betfair-go

Go wrapper library for interacting with the Betfair exchange, including both the standard and the streaming API. 

Currently only supports non-interactive login.

## Installation

To install the betfair-go library in your project, use the following command:

```sh
go get github.com/LonecastSystems/betfair-go
```

## Usage

### Standard Client

To use the standard Betfair client, you need to create a new client instance and authenticate with your Betfair credentials.

Example

```go
package main

import (
    "fmt"
    "github.com/LonecastSystems/betfair-go"
)

func main() {
	const certificate_crt_path = "C:\\Users\\User\\Documents\\Betfair\\client-2048.crt"
	const certificate_key_path = "C:\\Users\\User\\Documents\\Betfair\\client-2048.key"

	tlsConfig, err := betfair.GetTLSConfig(certificate_crt_path, certificate_key_path)
	if err != nil {
		panic(err)
	}

	config := &betfair.ClientConfig{Tls: tlsConfig, ApplicationKey: "AppKey", ApplicationName: "AppName", Rest: false}
	client := betfair.NewClient(config)

    //Login
	resp, err := client.Login("email@domain.com", "password")
	if err != nil {
		panic(err)
	}
}
```

### Streaming Client

To use the streaming client, you need an active session and use it to authenticate with the Betfair streaming API.

#### Without the standard client

Example
```go
package main

import (
    "fmt"
    "github.com/LonecastSystems/betfair-go"
)

func main() {
    streamClient := betfair.NewStreamingClient(&betfair.StreamingClientConfig{})

	const certificate_crt_path = "C:\\Users\\User\\Documents\\Betfair\\client-2048.crt"
	const certificate_key_path = "C:\\Users\\User\\Documents\\Betfair\\client-2048.key"

	tlsConfig, err := betfair.GetTLSConfig(certificate_crt_path, certificate_key_path)
	if err != nil {
		panic(err)
	}

    // Authenticate with the streaming API
    err = streamClient.Authenticate(tlsConfig, "AppKey", "SessionToken")
    if err != nil {
		panic(err)
    }
}
```

#### From the standard client

Example

```go
package main

import (
    "fmt"
    "github.com/LonecastSystems/betfair-go"
)

func main() {
	const certificate_crt_path = "C:\\Users\\User\\Documents\\Betfair\\client-2048.crt"
	const certificate_key_path = "C:\\Users\\User\\Documents\\Betfair\\client-2048.key"

	tlsConfig, err := betfair.GetTLSConfig(certificate_crt_path, certificate_key_path)
	if err != nil {
		panic(err)
	}

	config := &betfair.ClientConfig{Tls: tlsConfig, ApplicationKey: "AppKey", Rest: false, ApplicationName: "AppName"}
	client := betfair.NewClient(config)

    //Login
	resp, err := client.Login("email@domain.com", "password")
	if err != nil {
		panic(err)
	}

    // Authenticate with the streaming API
    streamClient, err := client.GetStream(&betfair.StreamingClientConfig{})
    if err != nil {
		panic(err)
    }
}
```