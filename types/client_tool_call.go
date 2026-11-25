package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ClientToolCall struct {
	ClientToolCall  *ClientToolCallClientToolCall `json:"client_tool_call" xml:"client_tool_call" required:"true"`
	Type            *string                       `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ClientToolCall) GetClientToolCall() (value *ClientToolCallClientToolCall) {
	if c == nil {
		return
	}
	return c.ClientToolCall
}

func (c *ClientToolCall) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ClientToolCall) SetClientToolCall(clientToolCall *ClientToolCallClientToolCall) {
	c.ClientToolCall = clientToolCall
}

func (c *ClientToolCall) SetType(type_ *string) {
	c.Type = type_
}

func (c *ClientToolCall) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ClientToolCall) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ClientToolCall to string"
	}
	return string(jsonData)
}

func (c *ClientToolCall) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ClientToolCall
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ClientToolCall(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ClientToolCall) MarshalJSON() ([]byte, error) {
	type embed ClientToolCall
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
