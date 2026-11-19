package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Custom SIP header for phone transfers with a static (validated) value.
type CustomSipHeader struct {
	// The SIP header name (e.g., 'X-Customer-ID')
	Key string `json:"key" xml:"key" required:"true" maxLength:"64" minLength:"1" pattern:"^[A-Za-z0-9]+(-[A-Za-z0-9]+)*$"`
	// The header value
	Value           string `json:"value" xml:"value" required:"true" maxLength:"512" minLength:"1" pattern:"^[\\x09\\x20-\\x7E]+$"`
	extraProperties map[string]interface{}
}

func (c *CustomSipHeader) GetKey() (value string) {
	if c == nil {
		return
	}
	return c.Key
}

func (c *CustomSipHeader) GetValue() (value string) {
	if c == nil {
		return
	}
	return c.Value
}

func (c *CustomSipHeader) SetKey(key string) {
	c.Key = key
}

func (c *CustomSipHeader) SetValue(value string) {
	c.Value = value
}

func (c *CustomSipHeader) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c CustomSipHeader) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: CustomSipHeader to string"
	}
	return string(jsonData)
}

func (c *CustomSipHeader) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias CustomSipHeader
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = CustomSipHeader(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *CustomSipHeader) MarshalJSON() ([]byte, error) {
	type embed CustomSipHeader
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
