package types

import (
	"encoding/json"
	"fmt"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// A `client_tool_result` whose `type` is required, so that it can share `input` with other items.
type ClientToolResultItem struct {
	type_ string
	// ID of the tool call this result corresponds to.
	ToolCallId string `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	Result     string `json:"result" xml:"result" required:"true"`
	IsError    bool   `json:"is_error" xml:"is_error" required:"true"`
	// Optional category for the failure reason. When set, `is_error` is treated as true. `user_rejected` is special-cased to mark the tool as not having been called.
	ErrorType       *ClientToolResultItemErrorType `json:"error_type,omitempty" xml:"error_type,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ClientToolResultItem) GetToolCallId() (value string) {
	if c == nil {
		return
	}
	return c.ToolCallId
}

func (c *ClientToolResultItem) GetResult() (value string) {
	if c == nil {
		return
	}
	return c.Result
}

func (c *ClientToolResultItem) GetIsError() (value bool) {
	if c == nil {
		return
	}
	return c.IsError
}

func (c *ClientToolResultItem) GetErrorType() (value *ClientToolResultItemErrorType) {
	if c == nil {
		return
	}
	return c.ErrorType
}

func (c *ClientToolResultItem) SetToolCallId(toolCallId string) {
	c.ToolCallId = toolCallId
}

func (c *ClientToolResultItem) SetResult(result string) {
	c.Result = result
}

func (c *ClientToolResultItem) SetIsError(isError bool) {
	c.IsError = isError
}

func (c *ClientToolResultItem) SetErrorType(errorType *ClientToolResultItemErrorType) {
	c.ErrorType = errorType
}

func (c *ClientToolResultItem) Type() string {
	return c.type_
}

func (c *ClientToolResultItem) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ClientToolResultItem) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ClientToolResultItem to string"
	}
	return string(jsonData)
}

func (c *ClientToolResultItem) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	var literals struct {
		Type *string `json:"type"`
	}
	if err := json.Unmarshal(data, &literals); err != nil {
		return err
	}
	if literals.Type != nil && *literals.Type != "client_tool_result" {
		return fmt.Errorf("unexpected value for literal on type %T; expected %v got %v", c, "client_tool_result", *literals.Type)
	}
	c.type_ = "client_tool_result"
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ClientToolResultItem) MarshalJSON() ([]byte, error) {
	type embed ClientToolResultItem
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*c),
		Type:  "client_tool_result",
	}
	return json.Marshal(marshaler)
}
