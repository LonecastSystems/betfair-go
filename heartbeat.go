package betfair

type ActionPerformed string

const (
	AP_NONE                           = "NONE"
	AP_CANCELLATION_REQUEST_SUBMITTED = "CANCELLATION_REQUEST_SUBMITTED"
	AP_ALL_BETS_CANCELLED             = "ALL_BETS_CANCELLED"
	AP_SOME_BETS_NOT_CANCELLED        = "SOME_BETS_NOT_CANCELLED"
	AP_CANCELLATION_REQUEST_ERROR     = "CANCELLATION_REQUEST_ERROR"
	AP_CANCELLATION_STATUS_UNKNOWN    = "CANCELLATION_STATUS_UNKNOWN"
)

type (
	HeartbeatParams struct {
		PreferredTimeoutSeconds int `json:"preferredTimeoutSeconds"`
	}

	HeartbeatReport struct {
		ActionPerformed      ActionPerformed `json:"actionPerformed"`
		ActualTimeoutSeconds int             `json:"actualTimeoutSeconds"`
	}
)

func (client *Client) Heartbeat(params HeartbeatParams) (HeartbeatReport, error) {
	json := HeartbeatReport{}

	if err := client.GetHeartbeats("heartbeat", params, &json); err != nil {
		return HeartbeatReport{}, err
	}

	return json, nil
}
