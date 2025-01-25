package stateless

import "time"

type (
	RaceDetailsParams struct {
		MeetingIDs []string `json:"meetingIds,omitempty"`
		RaceIDs    []string `json:"raceIds,omitempty"`
	}

	RaceStatus   string
	ResponseCode string

	RaceDetails struct {
		MeetingID    string       `json:"meetingId,omitempty"`
		RaceID       string       `json:"raceId,omitempty"`
		RaceStatus   RaceStatus   `json:"raceStatus,omitempty"`
		LastUpdated  time.Time    `json:"lastUpdated,omitempty"`
		Sequence     int64        `json:"sequence,omitempty"`
		ResponseCode ResponseCode `json:"responseCode,omitempty"`
	}
)

func (client *StatelessClient) ListRaceDetails(params RaceDetailsParams) ([]RaceDetails, error) {
	json := []RaceDetails{}

	if err := GetScores(client, "listRaceDetails", params, &json); err != nil {
		return []RaceDetails{}, err
	}

	return json, nil
}
