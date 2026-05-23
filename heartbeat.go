package betfair

import "context"

// ActionPerformed is the action performed since your last heartbeat request.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687861/Heartbeat+API#ActionPerformed
type ActionPerformed string

const (
	AP_NONE                           ActionPerformed = "NONE"                           // AP_NONE indicates no action was performed since last heartbeat, or this is the first heartbeat.
	AP_CANCELLATION_REQUEST_SUBMITTED ActionPerformed = "CANCELLATION_REQUEST_SUBMITTED" // AP_CANCELLATION_REQUEST_SUBMITTED indicates a request to cancel all unmatched bets was submitted since last heartbeat.
	AP_ALL_BETS_CANCELLED             ActionPerformed = "ALL_BETS_CANCELLED"             // AP_ALL_BETS_CANCELLED indicates all unmatched bets were cancelled since last heartbeat.
	AP_SOME_BETS_NOT_CANCELLED        ActionPerformed = "SOME_BETS_NOT_CANCELLED"        // AP_SOME_BETS_NOT_CANCELLED indicates not all unmatched bets were cancelled since last heartbeat.
	AP_CANCELLATION_REQUEST_ERROR     ActionPerformed = "CANCELLATION_REQUEST_ERROR"     // AP_CANCELLATION_REQUEST_ERROR indicates there was an error requesting cancellation, no bets have been cancelled.
	AP_CANCELLATION_STATUS_UNKNOWN    ActionPerformed = "CANCELLATION_STATUS_UNKNOWN"    // AP_CANCELLATION_STATUS_UNKNOWN indicates there was no response from requesting cancellation, cancellation status unknown.
)

type (
	// HeartbeatParams are the request parameters for heartbeat.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687861/Heartbeat+API
	HeartbeatParams struct {
		PreferredTimeoutSeconds int `json:"preferredTimeoutSeconds"` // PreferredTimeoutSeconds is the timeout requested.
	}

	// HeartbeatReport is the response from heartbeat.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687861/Heartbeat+API
	HeartbeatReport struct {
		ActionPerformed      ActionPerformed `json:"actionPerformed"`      // ActionPerformed is the action performed since your last heartbeat request.
		ActualTimeoutSeconds int             `json:"actualTimeoutSeconds"` // ActualTimeoutSeconds is the actual timeout applied.
	}
)

// Heartbeat helps customers have their positions managed automatically in the event of their API clients losing connectivity with the Betfair API. If a heartbeat request is not received within a prescribed time period, then Betfair will attempt to cancel all 'LIMIT' type bets for the given customer on the given exchange.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687861/Heartbeat+API
func (client *Client) Heartbeat(ctx context.Context, params HeartbeatParams) (HeartbeatReport, error) {
	json := HeartbeatReport{}

	return json, client.GetHeartbeats(ctx, "heartbeat", params, &json)
}
