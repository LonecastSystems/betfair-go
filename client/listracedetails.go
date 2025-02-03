package client

import "time"

type RaceStatus string

const (
	RS_DORMANT          = "DORMANT"
	RS_DELAYED          = "DELAYED"
	RS_PARADING         = "PARADING"
	RS_GOINGDOWN        = "GOINGDOWN"
	RS_GOINGBEHIND      = "GOINGBEHIND"
	RS_APPROACHING      = "APPROACHING"
	RS_GOINGINTRAPS     = "GOINGINTRAPS"
	RS_HARERUNNING      = "HARERUNNING"
	RS_ATTHEPOST        = "ATTHEPOST"
	RS_OFF              = "OFF"
	RS_FINISHED         = "FINISHED"
	RS_FINALRESULT      = "FINALRESULT"
	RS_FALSESTART       = "FALSESTART"
	RS_PHOTOGRAPH       = "PHOTOGRAPH"
	RS_RESULT           = "RESULT"
	RS_WEIGHEDIN        = "WEIGHEDIN"
	RS_RACEVOID         = "RACEVOID"
	RS_NORACE           = "NORACE"
	RS_MEETINGABANDONED = "MEETINGABANDONED"
	RS_RERUN            = "RERUN"
	RS_ABANDONED        = "ABANDONED"
)

type ResponseCode string

const (
	RC_OK                                = "OK"
	RC_NO_NEW_UPDATES                    = "NO_NEW_UPDATES"
	RC_NO_LIVE_DATA_AVAILABLE            = "NO_LIVE_DATA_AVAILABLE"
	RC_SERVICE_UNAVAILABLE               = "SERVICE_UNAVAILABLE"
	RC_UNEXPECTED_ERROR                  = "UNEXPECTED_ERROR"
	RC_LIVE_DATA_TEMPORARILY_UNAVAILABLE = "LIVE_DATA_TEMPORARILY_UNAVAILABLE"
)

type (
	RaceDetailsParams struct {
		MeetingIDs []string `json:"meetingIds,omitempty"`
		RaceIDs    []string `json:"raceIds,omitempty"`
	}

	RaceDetails struct {
		MeetingID    string       `json:"meetingId,omitempty"`
		RaceID       string       `json:"raceId,omitempty"`
		RaceStatus   RaceStatus   `json:"raceStatus,omitempty"`
		LastUpdated  time.Time    `json:"lastUpdated,omitempty"`
		Sequence     int64        `json:"sequence,omitempty"`
		ResponseCode ResponseCode `json:"responseCode,omitempty"`
	}
)

func (client *BetfairClient) ListRaceDetails(params RaceDetailsParams) ([]RaceDetails, error) {
	json := []RaceDetails{}

	if err := client.GetScores("listRaceDetails", params, &json); err != nil {
		return []RaceDetails{}, err
	}

	return json, nil
}
