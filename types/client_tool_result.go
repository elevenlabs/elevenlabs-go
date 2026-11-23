package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ClientToolResult struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// ID of the tool call this result corresponds to.
	ToolCallId string `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	Result     string `json:"result" xml:"result" required:"true"`
	IsError    bool   `json:"is_error" xml:"is_error" required:"true"`
	// Optional category for the failure reason. When set, `is_error` is treated as true. `user_rejected` is special-cased to mark the tool as not having been called.
	ErrorType       *ClientToolResultErrorType `json:"error_type,omitempty" xml:"error_type,omitempty"`
	extraProperties map[string]interface{}
}

func (c *ClientToolResult) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ClientToolResult) GetToolCallId() (value string) {
	if c == nil {
		return
	}
	return c.ToolCallId
}

func (c *ClientToolResult) GetResult() (value string) {
	if c == nil {
		return
	}
	return c.Result
}

func (c *ClientToolResult) GetIsError() (value bool) {
	if c == nil {
		return
	}
	return c.IsError
}

func (c *ClientToolResult) GetErrorType() (value *ClientToolResultErrorType) {
	if c == nil {
		return
	}
	return c.ErrorType
}

func (c *ClientToolResult) SetType(type_ *string) {
	c.Type = type_
}

func (c *ClientToolResult) SetToolCallId(toolCallId string) {
	c.ToolCallId = toolCallId
}

func (c *ClientToolResult) SetResult(result string) {
	c.Result = result
}

func (c *ClientToolResult) SetIsError(isError bool) {
	c.IsError = isError
}

func (c *ClientToolResult) SetErrorType(errorType *ClientToolResultErrorType) {
	c.ErrorType = errorType
}

func (c *ClientToolResult) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ClientToolResult) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ClientToolResult to string"
	}
	return string(jsonData)
}

func (c *ClientToolResult) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ClientToolResult
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ClientToolResult(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ClientToolResult) MarshalJSON() ([]byte, error) {
	type embed ClientToolResult
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
