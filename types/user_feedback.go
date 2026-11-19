package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UserFeedback struct {
	Score           UserFeedbackScore `json:"score" xml:"score" required:"true"`
	TimeInCallSecs  int               `json:"time_in_call_secs" xml:"time_in_call_secs" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UserFeedback) GetScore() (value UserFeedbackScore) {
	if u == nil {
		return
	}
	return u.Score
}

func (u *UserFeedback) GetTimeInCallSecs() (value int) {
	if u == nil {
		return
	}
	return u.TimeInCallSecs
}

func (u *UserFeedback) SetScore(score UserFeedbackScore) {
	u.Score = score
}

func (u *UserFeedback) SetTimeInCallSecs(timeInCallSecs int) {
	u.TimeInCallSecs = timeInCallSecs
}

func (u *UserFeedback) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UserFeedback) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UserFeedback to string"
	}
	return string(jsonData)
}

func (u *UserFeedback) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UserFeedback
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UserFeedback(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UserFeedback) MarshalJSON() ([]byte, error) {
	type embed UserFeedback
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
