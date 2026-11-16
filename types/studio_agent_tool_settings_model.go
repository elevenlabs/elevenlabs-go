package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type StudioAgentToolSettingsModel struct {
	SkipConfirmation *bool `json:"skip_confirmation,omitempty" xml:"skip_confirmation,omitempty"`
	extraProperties  map[string]interface{}
}

func (s *StudioAgentToolSettingsModel) GetSkipConfirmation() (value *bool) {
	if s == nil {
		return
	}
	return s.SkipConfirmation
}

func (s *StudioAgentToolSettingsModel) SetSkipConfirmation(skipConfirmation *bool) {
	s.SkipConfirmation = skipConfirmation
}

func (s *StudioAgentToolSettingsModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StudioAgentToolSettingsModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StudioAgentToolSettingsModel to string"
	}
	return string(jsonData)
}

func (s *StudioAgentToolSettingsModel) UnmarshalJSON(data []byte) error {
	type unmarshaler StudioAgentToolSettingsModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = StudioAgentToolSettingsModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StudioAgentToolSettingsModel) MarshalJSON() ([]byte, error) {
	type embed StudioAgentToolSettingsModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
