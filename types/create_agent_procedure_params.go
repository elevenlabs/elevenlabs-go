package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CreateAgentProcedureParams struct {
	extraProperties map[string]interface{}
}

func (c *CreateAgentProcedureParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateAgentProcedureParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateAgentProcedureParams to string"
	}
	return string(jsonData)
}

func (c *CreateAgentProcedureParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateAgentProcedureParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateAgentProcedureParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateAgentProcedureParams) MarshalJSON() ([]byte, error) {
	type embed CreateAgentProcedureParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
