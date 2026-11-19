package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ClientToolCallRequiredAction struct {
	Type            *string        `json:"type,omitempty" xml:"type,omitempty"`
	ToolCallId      string         `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	ToolName        string         `json:"tool_name" xml:"tool_name" required:"true"`
	Parameters      map[string]any `json:"parameters" required:"true"`
	extraProperties map[string]interface{}
}

func (c *ClientToolCallRequiredAction) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ClientToolCallRequiredAction) GetToolCallId() (value string) {
	if c == nil {
		return
	}
	return c.ToolCallId
}

func (c *ClientToolCallRequiredAction) GetToolName() (value string) {
	if c == nil {
		return
	}
	return c.ToolName
}

func (c *ClientToolCallRequiredAction) GetParameters() (value map[string]any) {
	if c == nil {
		return
	}
	return c.Parameters
}

func (c *ClientToolCallRequiredAction) SetType(type_ *string) {
	c.Type = type_
}

func (c *ClientToolCallRequiredAction) SetToolCallId(toolCallId string) {
	c.ToolCallId = toolCallId
}

func (c *ClientToolCallRequiredAction) SetToolName(toolName string) {
	c.ToolName = toolName
}

func (c *ClientToolCallRequiredAction) SetParameters(parameters map[string]any) {
	c.Parameters = parameters
}

func (c *ClientToolCallRequiredAction) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ClientToolCallRequiredAction) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ClientToolCallRequiredAction to string"
	}
	return string(jsonData)
}

func (c *ClientToolCallRequiredAction) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ClientToolCallRequiredAction
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ClientToolCallRequiredAction(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ClientToolCallRequiredAction) MarshalJSON() ([]byte, error) {
	type embed ClientToolCallRequiredAction
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
