package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// `ClientToolCallClientEvent` as a stream event, because its custom `__init__` can't parse its own JSON.
type ClientToolCallEvent struct {
	ClientToolCall  *ClientToolCallClientEventData `json:"client_tool_call" xml:"client_tool_call" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ClientToolCallEvent) GetClientToolCall() (value *ClientToolCallClientEventData) {
	if c == nil {
		return
	}
	return c.ClientToolCall
}

func (c *ClientToolCallEvent) SetClientToolCall(clientToolCall *ClientToolCallClientEventData) {
	c.ClientToolCall = clientToolCall
}

func (c *ClientToolCallEvent) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ClientToolCallEvent) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ClientToolCallEvent to string"
	}
	return string(jsonData)
}

func (c *ClientToolCallEvent) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ClientToolCallEvent
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ClientToolCallEvent(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ClientToolCallEvent) MarshalJSON() ([]byte, error) {
	type embed ClientToolCallEvent
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
