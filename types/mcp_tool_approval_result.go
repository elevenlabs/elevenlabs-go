package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type McpToolApprovalResult struct {
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The ID of the tool call this approval responds to.
	ToolCallId string `json:"tool_call_id" xml:"tool_call_id" required:"true"`
	// Whether the user approved the tool call.
	IsApproved      bool `json:"is_approved" xml:"is_approved" required:"true"`
	extraProperties map[string]interface{}
}

func (m *McpToolApprovalResult) GetType() (value *string) {
	if m == nil {
		return
	}
	return m.Type
}

func (m *McpToolApprovalResult) GetToolCallId() (value string) {
	if m == nil {
		return
	}
	return m.ToolCallId
}

func (m *McpToolApprovalResult) GetIsApproved() (value bool) {
	if m == nil {
		return
	}
	return m.IsApproved
}

func (m *McpToolApprovalResult) SetType(type_ *string) {
	m.Type = type_
}

func (m *McpToolApprovalResult) SetToolCallId(toolCallId string) {
	m.ToolCallId = toolCallId
}

func (m *McpToolApprovalResult) SetIsApproved(isApproved bool) {
	m.IsApproved = isApproved
}

func (m *McpToolApprovalResult) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m McpToolApprovalResult) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: McpToolApprovalResult to string"
	}
	return string(jsonData)
}

func (m *McpToolApprovalResult) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, m); err != nil {
		return err
	}
	type alias McpToolApprovalResult
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*m = McpToolApprovalResult(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *McpToolApprovalResult) MarshalJSON() ([]byte, error) {
	type embed McpToolApprovalResult
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
