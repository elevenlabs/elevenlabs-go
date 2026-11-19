package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UtteranceResponseModel struct {
	// The start time of the utterance in seconds.
	Start float64 `json:"start" xml:"start" required:"true"`
	// The end time of the utterance in seconds.
	End             float64 `json:"end" xml:"end" required:"true"`
	extraProperties map[string]interface{}
}

func (u *UtteranceResponseModel) GetStart() (value float64) {
	if u == nil {
		return
	}
	return u.Start
}

func (u *UtteranceResponseModel) GetEnd() (value float64) {
	if u == nil {
		return
	}
	return u.End
}

func (u *UtteranceResponseModel) SetStart(start float64) {
	u.Start = start
}

func (u *UtteranceResponseModel) SetEnd(end float64) {
	u.End = end
}

func (u *UtteranceResponseModel) GetExtraProperties() map[string]interface{} {
	if u == nil {
		return nil
	}
	return u.extraProperties
}

func (u UtteranceResponseModel) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UtteranceResponseModel to string"
	}
	return string(jsonData)
}

func (u *UtteranceResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, u); err != nil {
		return err
	}
	type alias UtteranceResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*u = UtteranceResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties
	return nil
}

func (u *UtteranceResponseModel) MarshalJSON() ([]byte, error) {
	type embed UtteranceResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
