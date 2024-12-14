package streaming

type (
	StatusMessage struct {
		ID                   int    `json:"id"`
		StatusCode           string `json:"statusCode"`
		ConnectionClosed     bool   `json:"connectionClosed"`
		ErrorCode            string `json:"errorCode"`
		ErrorMessage         string `json:"errorMessage"`
		ConnectionsAvailable int    `json:"connectionsAvailable"`
	}
)

// Connection/Authentication
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
