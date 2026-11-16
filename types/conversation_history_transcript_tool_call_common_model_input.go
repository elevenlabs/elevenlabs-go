package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptToolCallCommonModelInput struct {
	Type      *ToolType `json:"type,omitempty" xml:"type,omitempty"`
	RequestId string    `json:"request_id" xml:"request_id" required:"true"`
	ToolName  string    `json:"tool_name" xml:"tool_name" required:"true"`
	// ID of the called tool. Not set for tools without an ID, such as system tools.
	ToolId            *string                                                           `json:"tool_id,omitempty" xml:"tool_id,omitempty"`
	ParamsAsJson      string                                                            `json:"params_as_json" xml:"params_as_json" required:"true"`
	ToolHasBeenCalled bool                                                              `json:"tool_has_been_called" xml:"tool_has_been_called" required:"true"`
	ToolDetails       *ConversationHistoryTranscriptToolCallCommonModelInputToolDetails `json:"tool_details,omitempty" xml:"tool_details,omitempty"`
	extraProperties   map[string]interface{}
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) GetType() (value *ToolType) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) GetRequestId() (value string) {
	if c == nil {
		return
	}
	return c.RequestId
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) GetToolName() (value string) {
	if c == nil {
		return
	}
	return c.ToolName
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) GetToolId() (value *string) {
	if c == nil {
		return
	}
	return c.ToolId
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) GetParamsAsJson() (value string) {
	if c == nil {
		return
	}
	return c.ParamsAsJson
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) GetToolHasBeenCalled() (value bool) {
	if c == nil {
		return
	}
	return c.ToolHasBeenCalled
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) GetToolDetails() (value *ConversationHistoryTranscriptToolCallCommonModelInputToolDetails) {
	if c == nil {
		return
	}
	return c.ToolDetails
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) SetType(type_ *ToolType) {
	c.Type = type_
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) SetRequestId(requestId string) {
	c.RequestId = requestId
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) SetToolName(toolName string) {
	c.ToolName = toolName
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) SetToolId(toolId *string) {
	c.ToolId = toolId
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) SetParamsAsJson(paramsAsJson string) {
	c.ParamsAsJson = paramsAsJson
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) SetToolHasBeenCalled(toolHasBeenCalled bool) {
	c.ToolHasBeenCalled = toolHasBeenCalled
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) SetToolDetails(toolDetails *ConversationHistoryTranscriptToolCallCommonModelInputToolDetails) {
	c.ToolDetails = toolDetails
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptToolCallCommonModelInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptToolCallCommonModelInput to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTranscriptToolCallCommonModelInput) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptToolCallCommonModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
