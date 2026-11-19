package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type StartSpeakerSeparationResponseModel struct {
	// The status of the start speaker seperation request. If the request was successful, the status will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (s *StartSpeakerSeparationResponseModel) GetStatus() (value string) {
	if s == nil {
		return
	}
	return s.Status
}

func (s *StartSpeakerSeparationResponseModel) SetStatus(status string) {
	s.Status = status
}

func (s *StartSpeakerSeparationResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StartSpeakerSeparationResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StartSpeakerSeparationResponseModel to string"
	}
	return string(jsonData)
}

func (s *StartSpeakerSeparationResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias StartSpeakerSeparationResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = StartSpeakerSeparationResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StartSpeakerSeparationResponseModel) MarshalJSON() ([]byte, error) {
	type embed StartSpeakerSeparationResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
