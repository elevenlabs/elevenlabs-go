package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ClientError struct {
	ErrorEvent      *ClientErrorErrorEvent `json:"error_event" xml:"error_event" required:"true"`
	Type            *string                `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ClientError) GetErrorEvent() (value *ClientErrorErrorEvent) {
	if c == nil {
		return
	}
	return c.ErrorEvent
}

func (c *ClientError) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ClientError) SetErrorEvent(errorEvent *ClientErrorErrorEvent) {
	c.ErrorEvent = errorEvent
}

func (c *ClientError) SetType(type_ *string) {
	c.Type = type_
}

func (c *ClientError) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ClientError) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ClientError to string"
	}
	return string(jsonData)
}

func (c *ClientError) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ClientError
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ClientError(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ClientError) MarshalJSON() ([]byte, error) {
	type embed ClientError
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
