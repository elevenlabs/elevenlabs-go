package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type SpeakerUpdatedResponse struct {
	Version         int `json:"version" xml:"version" required:"true"`
	extraProperties map[string]interface{}
}

func (s *SpeakerUpdatedResponse) GetVersion() (value int) {
	if s == nil {
		return
	}
	return s.Version
}

func (s *SpeakerUpdatedResponse) SetVersion(version int) {
	s.Version = version
}

func (s *SpeakerUpdatedResponse) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s SpeakerUpdatedResponse) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: SpeakerUpdatedResponse to string"
	}
	return string(jsonData)
}

func (s *SpeakerUpdatedResponse) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, s); err != nil {
		return err
	}
	type alias SpeakerUpdatedResponse
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*s = SpeakerUpdatedResponse(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *SpeakerUpdatedResponse) MarshalJSON() ([]byte, error) {
	type embed SpeakerUpdatedResponse
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
