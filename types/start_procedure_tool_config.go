package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type StartProcedureToolConfig struct {
	Procedures      map[string]*StartProcedureToolConfigProceduresValue `json:"procedures,omitempty"`
	extraProperties map[string]interface{}
}

func (s *StartProcedureToolConfig) GetProcedures() (value map[string]*StartProcedureToolConfigProceduresValue) {
	if s == nil {
		return
	}
	return s.Procedures
}

func (s *StartProcedureToolConfig) SetProcedures(procedures map[string]*StartProcedureToolConfigProceduresValue) {
	s.Procedures = procedures
}

func (s *StartProcedureToolConfig) GetExtraProperties() map[string]interface{} {
	if s == nil {
		return nil
	}
	return s.extraProperties
}

func (s StartProcedureToolConfig) String() string {
	jsonData, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return "error converting struct: StartProcedureToolConfig to string"
	}
	return string(jsonData)
}

func (s *StartProcedureToolConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler StartProcedureToolConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = StartProcedureToolConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *s, "system_tool_type")
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties
	return nil
}

func (s *StartProcedureToolConfig) MarshalJSON() ([]byte, error) {
	type embed StartProcedureToolConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*s),
	}
	return json.Marshal(marshaler)
}
