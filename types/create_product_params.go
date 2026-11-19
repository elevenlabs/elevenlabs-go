package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CreateProductParams struct {
	extraProperties map[string]interface{}
}

func (c *CreateProductParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateProductParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateProductParams to string"
	}
	return string(jsonData)
}

func (c *CreateProductParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateProductParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateProductParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateProductParams) MarshalJSON() ([]byte, error) {
	type embed CreateProductParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
