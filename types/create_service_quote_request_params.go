package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CreateServiceQuoteRequestParams struct {
	extraProperties map[string]interface{}
}

func (c *CreateServiceQuoteRequestParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateServiceQuoteRequestParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateServiceQuoteRequestParams to string"
	}
	return string(jsonData)
}

func (c *CreateServiceQuoteRequestParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateServiceQuoteRequestParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateServiceQuoteRequestParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateServiceQuoteRequestParams) MarshalJSON() ([]byte, error) {
	type embed CreateServiceQuoteRequestParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
