package stateless

type (
	HeartbeatParams struct {
		PreferredTimeoutSeconds int `json:"preferredTimeoutSeconds"`
	}

	ActionPerformed string

	HeartbeatReport struct {
		ActionPerformed      ActionPerformed `json:"actionPerformed"`
		ActualTimeoutSeconds int             `json:"actualTimeoutSeconds"`
	}
)

func (client *StatelessClient) Heartbeat(params HeartbeatParams) (HeartbeatReport, error) {
	json := HeartbeatReport{}

	if err := GetHeartbeats(client, "heartbeat", params, &json); err != nil {
		return HeartbeatReport{}, err
	}

	return json, nil
}
