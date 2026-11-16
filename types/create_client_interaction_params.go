package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CreateClientInteractionParams struct {
	extraProperties map[string]interface{}
}

func (c *CreateClientInteractionParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateClientInteractionParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateClientInteractionParams to string"
	}
	return string(jsonData)
}

func (c *CreateClientInteractionParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateClientInteractionParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateClientInteractionParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateClientInteractionParams) MarshalJSON() ([]byte, error) {
	type embed CreateClientInteractionParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
