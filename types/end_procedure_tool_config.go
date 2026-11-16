package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type EndProcedureToolConfig struct {
	Procedures      map[string]*EndProcedureToolConfigProceduresValue `json:"procedures,omitempty"`
	extraProperties map[string]interface{}
}

func (e *EndProcedureToolConfig) GetProcedures() (value map[string]*EndProcedureToolConfigProceduresValue) {
	if e == nil {
		return
	}
	return e.Procedures
}

func (e *EndProcedureToolConfig) SetProcedures(procedures map[string]*EndProcedureToolConfigProceduresValue) {
	e.Procedures = procedures
}

func (e *EndProcedureToolConfig) GetExtraProperties() map[string]interface{} {
	if e == nil {
		return nil
	}
	return e.extraProperties
}

func (e EndProcedureToolConfig) String() string {
	jsonData, err := json.MarshalIndent(e, "", "  ")
	if err != nil {
		return "error converting struct: EndProcedureToolConfig to string"
	}
	return string(jsonData)
}

func (e *EndProcedureToolConfig) UnmarshalJSON(data []byte) error {
	type unmarshaler EndProcedureToolConfig
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*e = EndProcedureToolConfig(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *e, "system_tool_type")
	if err != nil {
		return err
	}
	e.extraProperties = extraProperties
	return nil
}

func (e *EndProcedureToolConfig) MarshalJSON() ([]byte, error) {
	type embed EndProcedureToolConfig
	var marshaler = struct {
		embed
	}{
		embed: embed(*e),
	}
	return json.Marshal(marshaler)
}
