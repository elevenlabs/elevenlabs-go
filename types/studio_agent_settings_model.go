package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type StudioAgentSettingsModel struct {
	ToolSettings    map[string]*StudioAgentToolSettingsModel `json:"tool_settings,omitempty"`
	extraProperties map[string]interface{}
}

func (s *StudioAgentSettingsModel) GetToolSettings() (value map[string]*StudioAgentToolSettingsModel) {
	if s == nil {
		return
	}
	return s.ToolSettings
}

func (s *StudioAgentSettingsModel) SetToolSettings(toolSettings map[string]*StudioAgentToolSettingsModel) {
	s.ToolSettings = toolSettings
}

func (s *StudioAgentSettingsModel) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StudioAgentSettingsModel) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StudioAgentSettingsModel to string"
	}
	return string(jsonData)
}

func (s *StudioAgentSettingsModel) UnmarshalJSON(data []byte) error {
	type unmarshaler StudioAgentSettingsModel
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = StudioAgentSettingsModel(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StudioAgentSettingsModel) MarshalJSON() ([]byte, error) {
	type embed StudioAgentSettingsModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
