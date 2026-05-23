package betfair

import (
	"context"
	"time"
)

// RaceStatus is the current status of the race.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687662/Race+Status+API#RaceStatus
type RaceStatus string

const (
	RS_DORMANT          RaceStatus = "DORMANT"          // RS_DORMANT indicates there is no data available for this race.
	RS_DELAYED          RaceStatus = "DELAYED"          // RS_DELAYED indicates the start of the race has been delayed.
	RS_PARADING         RaceStatus = "PARADING"         // RS_PARADING indicates the horses/greyhounds are in the parade ring.
	RS_GOINGDOWN        RaceStatus = "GOINGDOWN"        // RS_GOINGDOWN indicates the horses are going down to the starting post.
	RS_GOINGBEHIND      RaceStatus = "GOINGBEHIND"      // RS_GOINGBEHIND indicates the horses are going behind the stalls.
	RS_APPROACHING      RaceStatus = "APPROACHING"      // RS_APPROACHING indicates the greyhounds are approaching the traps.
	RS_GOINGINTRAPS     RaceStatus = "GOINGINTRAPS"     // RS_GOINGINTRAPS indicates the greyhounds are being put in the traps.
	RS_HARERUNNING      RaceStatus = "HARERUNNING"      // RS_HARERUNNING indicates the hare has been started.
	RS_ATTHEPOST        RaceStatus = "ATTHEPOST"        // RS_ATTHEPOST indicates the horses are at the post.
	RS_OFF              RaceStatus = "OFF"              // RS_OFF indicates the greyhound/horse race has started.
	RS_FINISHED         RaceStatus = "FINISHED"         // RS_FINISHED indicates the race has finished.
	RS_FINALRESULT      RaceStatus = "FINALRESULT"      // RS_FINALRESULT indicates the result has been declared (Greyhounds only).
	RS_FALSESTART       RaceStatus = "FALSESTART"       // RS_FALSESTART indicates there has been a false start.
	RS_PHOTOGRAPH       RaceStatus = "PHOTOGRAPH"       // RS_PHOTOGRAPH indicates the result of the race is subject to a photo finish.
	RS_RESULT           RaceStatus = "RESULT"           // RS_RESULT indicates the result of the race has been announced.
	RS_WEIGHEDIN        RaceStatus = "WEIGHEDIN"        // RS_WEIGHEDIN indicates the jockeys have weighed in.
	RS_RACEVOID         RaceStatus = "RACEVOID"         // RS_RACEVOID indicates the race has been declared void.
	RS_NORACE           RaceStatus = "NORACE"           // RS_NORACE indicates the race has been declared a no race.
	RS_MEETINGABANDONED RaceStatus = "MEETINGABANDONED" // RS_MEETINGABANDONED indicates the meeting has been abandoned.
	RS_RERUN            RaceStatus = "RERUN"            // RS_RERUN indicates the race will be rerun.
	RS_ABANDONED        RaceStatus = "ABANDONED"        // RS_ABANDONED indicates the race has been abandoned.
)

// ResponseCode is the response code for a race status update.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687662/Race+Status+API#ResponseCode
type ResponseCode string

const (
	RC_OK                                ResponseCode = "OK"                                // RC_OK indicates data returned successfully.
	RC_NO_NEW_UPDATES                    ResponseCode = "NO_NEW_UPDATES"                    // RC_NO_NEW_UPDATES indicates no updates since the passes UpdateSequence.
	RC_NO_LIVE_DATA_AVAILABLE            ResponseCode = "NO_LIVE_DATA_AVAILABLE"            // RC_NO_LIVE_DATA_AVAILABLE indicates event scores are no longer available or are not on the schedule.
	RC_SERVICE_UNAVAILABLE               ResponseCode = "SERVICE_UNAVAILABLE"               // RC_SERVICE_UNAVAILABLE indicates the data feed for the event type (tennis/football etc) is currently unavailable.
	RC_UNEXPECTED_ERROR                  ResponseCode = "UNEXPECTED_ERROR"                  // RC_UNEXPECTED_ERROR indicates an unexpected error occurred retrieving score data.
	RC_LIVE_DATA_TEMPORARILY_UNAVAILABLE ResponseCode = "LIVE_DATA_TEMPORARILY_UNAVAILABLE" // RC_LIVE_DATA_TEMPORARILY_UNAVAILABLE indicates live data is temporarily unavailable (stale).
)

type (
	// ListRaceDetailsParams are the request parameters for listRaceDetails.
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687662/Race+Status+API
	ListRaceDetailsParams struct {
		MeetingIDs []string `json:"meetingIds,omitempty"` // MeetingIDs optionally restricts the results to the specified meeting IDs.
		RaceIDs    []string `json:"raceIds,omitempty"`    // RaceIDs optionally restricts the results to the specified race IDs in the format meetingid.raceTime (hhmm).
	}

	// ListRaceDetailsResult holds race details (RaceDetails).
	//
	// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687662/Race+Status+API#RaceDetails
	ListRaceDetailsResult struct {
		MeetingID    string       `json:"meetingId,omitempty"`    // MeetingID is the unique Id for the meeting equivalent to the eventId for that specific race as returned by listEvents.
		RaceID       string       `json:"raceId,omitempty"`       // RaceID is the unique Id for the race in the format meetingid.raceTime (hhmm).
		RaceStatus   RaceStatus   `json:"raceStatus,omitempty"`   // RaceStatus is the current status of the race.
		LastUpdated  time.Time    `json:"lastUpdated,omitempty"`  // LastUpdated is the time the data was last updated.
		Sequence     int64        `json:"sequence,omitempty"`     // Sequence is the unique identifier associated to each update of the data.
		ResponseCode ResponseCode `json:"responseCode,omitempty"` // ResponseCode is the response code for this race status update.
	}
)

// ListRaceDetails searches for races to get their details. This information is available for UK, Ireland and South African races only and is provided for information purposes only.
//
// Doc: https://betfair-developer-docs.atlassian.net/wiki/spaces/1smk3cen4v3lu3yomq5qye0ni/pages/2687662/Race+Status+API
func (client *Client) ListRaceDetails(ctx context.Context, params ListRaceDetailsParams) ([]ListRaceDetailsResult, error) {
	json := []ListRaceDetailsResult{}

	return json, client.GetScores(ctx, "listRaceDetails", params, &json)
}
