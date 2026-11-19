package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

// Response model for testing tools available on an MCP server.
type ListMcpToolsResponseModel struct {
	// Indicates if the operation was successful.
	Success bool `json:"success" xml:"success" required:"true"`
	// A list of tools available on the MCP server.
	Tools []*Tool `json:"tools" xml:"tools" required:"true"`
	// Derived approval states for currently discovered tools. Populated only for persisted MCP servers using per-tool approval; otherwise empty.
	ToolApprovalStatuses []*McpToolApprovalStatus `json:"tool_approval_statuses,omitempty" xml:"tool_approval_statuses,omitempty"`
	// Error message if the operation was not successful.
	ErrorMessage    *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	extraProperties map[string]interface{}
}

func (l *ListMcpToolsResponseModel) GetSuccess() (value bool) {
	if l == nil {
		return
	}
	return l.Success
}

func (l *ListMcpToolsResponseModel) GetTools() (value []*Tool) {
	if l == nil {
		return
	}
	return l.Tools
}

func (l *ListMcpToolsResponseModel) GetToolApprovalStatuses() (value []*McpToolApprovalStatus) {
	if l == nil {
		return
	}
	return l.ToolApprovalStatuses
}

func (l *ListMcpToolsResponseModel) GetErrorMessage() (value *string) {
	if l == nil {
		return
	}
	return l.ErrorMessage
}

func (l *ListMcpToolsResponseModel) SetSuccess(success bool) {
	l.Success = success
}

func (l *ListMcpToolsResponseModel) SetTools(tools []*Tool) {
	l.Tools = tools
}

func (l *ListMcpToolsResponseModel) SetToolApprovalStatuses(toolApprovalStatuses []*McpToolApprovalStatus) {
	l.ToolApprovalStatuses = toolApprovalStatuses
}

func (l *ListMcpToolsResponseModel) SetErrorMessage(errorMessage *string) {
	l.ErrorMessage = errorMessage
}

func (l *ListMcpToolsResponseModel) GetExtraProperties() map[string]interface{} {
	if l == nil {
		return nil
	}
	return l.extraProperties
}

func (l ListMcpToolsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(l, "", "  ")
	if err != nil {
		return "error converting struct: ListMcpToolsResponseModel to string"
	}
	return string(jsonData)
}

func (l *ListMcpToolsResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, l); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties
	return nil
}

func (l *ListMcpToolsResponseModel) MarshalJSON() ([]byte, error) {
	type embed ListMcpToolsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*l),
	}
	return json.Marshal(marshaler)
}
