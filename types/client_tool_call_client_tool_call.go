package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ClientToolCallClientToolCall struct {
	ToolName   string         `json:"tool_name" xml:"tool_name" required:"true"`
	ToolCallId string         `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	Parameters map[string]any `json:"parameters" xml:"parameters" required:"true"`
	EventId    int            `json:"event_id" xml:"event_id" required:"true"`
	// Whether the server expects a ClientToolResult response.
	ExpectsResponse bool `json:"expects_response" xml:"expects_response" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ClientToolCallClientToolCall) GetToolName() (value string) {
	if c == nil {
		return
	}
	return c.ToolName
}

func (c *ClientToolCallClientToolCall) GetToolCallId() (value string) {
	if c == nil {
		return
	}
	return c.ToolCallId
}

func (c *ClientToolCallClientToolCall) GetParameters() (value map[string]any) {
	if c == nil {
		return
	}
	return c.Parameters
}

func (c *ClientToolCallClientToolCall) GetEventId() (value int) {
	if c == nil {
		return
	}
	return c.EventId
}

func (c *ClientToolCallClientToolCall) GetExpectsResponse() (value bool) {
	if c == nil {
		return
	}
	return c.ExpectsResponse
}

func (c *ClientToolCallClientToolCall) SetToolName(toolName string) {
	c.ToolName = toolName
}

func (c *ClientToolCallClientToolCall) SetToolCallId(toolCallId string) {
	c.ToolCallId = toolCallId
}

func (c *ClientToolCallClientToolCall) SetParameters(parameters map[string]any) {
	c.Parameters = parameters
}

func (c *ClientToolCallClientToolCall) SetEventId(eventId int) {
	c.EventId = eventId
}

func (c *ClientToolCallClientToolCall) SetExpectsResponse(expectsResponse bool) {
	c.ExpectsResponse = expectsResponse
}

func (c *ClientToolCallClientToolCall) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ClientToolCallClientToolCall) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ClientToolCallClientToolCall to string"
	}
	return string(jsonData)
}

func (c *ClientToolCallClientToolCall) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ClientToolCallClientToolCall
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ClientToolCallClientToolCall(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ClientToolCallClientToolCall) MarshalJSON() ([]byte, error) {
	type embed ClientToolCallClientToolCall
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
