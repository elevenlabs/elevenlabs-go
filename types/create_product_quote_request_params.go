package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type CreateProductQuoteRequestParams struct {
	extraProperties map[string]interface{}
}

func (c *CreateProductQuoteRequestParams) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CreateProductQuoteRequestParams) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CreateProductQuoteRequestParams to string"
	}
	return string(jsonData)
}

func (c *CreateProductQuoteRequestParams) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateProductQuoteRequestParams
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateProductQuoteRequestParams(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "smb_tool_type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CreateProductQuoteRequestParams) MarshalJSON() ([]byte, error) {
	type embed CreateProductQuoteRequestParams
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
