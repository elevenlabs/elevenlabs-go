package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptToolCallCommonModelOutput struct {
	Type      *ToolType `json:"type,omitempty" xml:"type,omitempty"`
	RequestId string    `json:"request_id" xml:"request_id" required:"true"`
	ToolName  string    `json:"tool_name" xml:"tool_name" required:"true"`
	// ID of the called tool. Not set for tools without an ID, such as system tools.
	ToolId            *string                                                            `json:"tool_id,omitempty" xml:"tool_id,omitempty"`
	ParamsAsJson      string                                                             `json:"params_as_json" xml:"params_as_json" required:"true"`
	ToolHasBeenCalled bool                                                               `json:"tool_has_been_called" xml:"tool_has_been_called" required:"true"`
	ToolDetails       *ConversationHistoryTranscriptToolCallCommonModelOutputToolDetails `json:"tool_details,omitempty" xml:"tool_details,omitempty"`
	extraProperties   map[string]interface{}
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) GetType() (value *ToolType) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) GetRequestId() (value string) {
	if c == nil {
		return
	}
	return c.RequestId
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) GetToolName() (value string) {
	if c == nil {
		return
	}
	return c.ToolName
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) GetToolId() (value *string) {
	if c == nil {
		return
	}
	return c.ToolId
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) GetParamsAsJson() (value string) {
	if c == nil {
		return
	}
	return c.ParamsAsJson
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) GetToolHasBeenCalled() (value bool) {
	if c == nil {
		return
	}
	return c.ToolHasBeenCalled
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) GetToolDetails() (value *ConversationHistoryTranscriptToolCallCommonModelOutputToolDetails) {
	if c == nil {
		return
	}
	return c.ToolDetails
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) SetType(type_ *ToolType) {
	c.Type = type_
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) SetRequestId(requestId string) {
	c.RequestId = requestId
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) SetToolName(toolName string) {
	c.ToolName = toolName
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) SetToolId(toolId *string) {
	c.ToolId = toolId
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) SetParamsAsJson(paramsAsJson string) {
	c.ParamsAsJson = paramsAsJson
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) SetToolHasBeenCalled(toolHasBeenCalled bool) {
	c.ToolHasBeenCalled = toolHasBeenCalled
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) SetToolDetails(toolDetails *ConversationHistoryTranscriptToolCallCommonModelOutputToolDetails) {
	c.ToolDetails = toolDetails
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptToolCallCommonModelOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptToolCallCommonModelOutput to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) UnmarshalJSON(data []byte) error {
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

func (c *ConversationHistoryTranscriptToolCallCommonModelOutput) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptToolCallCommonModelOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
