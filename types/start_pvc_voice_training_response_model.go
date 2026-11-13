package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type StartPvcVoiceTrainingResponseModel struct {
	// The status of the start PVC voice training request. If the request was successful, the status will be 'ok'. Otherwise an error message with status 500 will be returned.
	Status          string `json:"status" xml:"status" required:"true"`
	extraProperties map[string]interface{}
}

func (s *StartPvcVoiceTrainingResponseModel) GetStatus() (value string) {
	if s == nil {
		return
	}
	return s.Status
}

func (s *StartPvcVoiceTrainingResponseModel) SetStatus(status string) {
	s.Status = status
}

func (s *StartPvcVoiceTrainingResponseModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StartPvcVoiceTrainingResponseModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StartPvcVoiceTrainingResponseModel to string"
	}
	return string(jsonData)
}

func (s *StartPvcVoiceTrainingResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias StartPvcVoiceTrainingResponseModel
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = StartPvcVoiceTrainingResponseModel(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StartPvcVoiceTrainingResponseModel) MarshalJSON() ([]byte, error) {
	type embed StartPvcVoiceTrainingResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
