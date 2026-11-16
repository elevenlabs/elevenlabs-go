package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CreateAgentRuleParams struct {
	extraProperties map[string]interface{}
}

func (c *CreateAgentRuleParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateAgentRuleParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateAgentRuleParams to string"
	}
	return string(jsonData)
}

func (c *CreateAgentRuleParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateAgentRuleParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateAgentRuleParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateAgentRuleParams) MarshalJSON() ([]byte, error) {
	type embed CreateAgentRuleParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
