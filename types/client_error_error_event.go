package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ClientErrorErrorEvent struct {
	Code            int     `json:"code" xml:"code" required:"true"`
	ErrorName       string  `json:"error_name" xml:"error_name" required:"true"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ClientErrorErrorEvent) GetCode() (value int) {
	if c == nil {
		return
	}
	return c.Code
}

func (c *ClientErrorErrorEvent) GetErrorName() (value string) {
	if c == nil {
		return
	}
	return c.ErrorName
}

func (c *ClientErrorErrorEvent) GetMessage() (value *string) {
	if c == nil {
		return
	}
	return c.Message
}

func (c *ClientErrorErrorEvent) SetCode(code int) {
	c.Code = code
}

func (c *ClientErrorErrorEvent) SetErrorName(errorName string) {
	c.ErrorName = errorName
}

func (c *ClientErrorErrorEvent) SetMessage(message *string) {
	c.Message = message
}

func (c *ClientErrorErrorEvent) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ClientErrorErrorEvent) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ClientErrorErrorEvent to string"
	}
	return string(jsonData)
}

func (c *ClientErrorErrorEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ClientErrorErrorEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ClientErrorErrorEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ClientErrorErrorEvent) MarshalJSON() ([]byte, error) {
	type embed ClientErrorErrorEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
