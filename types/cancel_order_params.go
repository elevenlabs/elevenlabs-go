package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CancelOrderParams struct {
	extraProperties map[string]interface{}
}

func (c *CancelOrderParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CancelOrderParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CancelOrderParams to string"
	}
	return string(jsonData)
}

func (c *CancelOrderParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CancelOrderParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CancelOrderParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CancelOrderParams) MarshalJSON() ([]byte, error) {
	type embed CancelOrderParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
