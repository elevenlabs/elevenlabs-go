package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptToolCallMcpDetails struct {
	McpServerId        string            `json:"mcp_server_id" xml:"mcp_server_id" required:"true"`
	McpServerName      string            `json:"mcp_server_name" xml:"mcp_server_name" required:"true"`
	IntegrationType    string            `json:"integration_type" xml:"integration_type" required:"true"`
	Parameters         map[string]string `json:"parameters,omitempty"`
	ApprovalPolicy     string            `json:"approval_policy" xml:"approval_policy" required:"true"`
	RequiresApproval   *bool             `json:"requires_approval,omitempty" xml:"requires_approval,omitempty"`
	McpToolName        *string           `json:"mcp_tool_name,omitempty" xml:"mcp_tool_name,omitempty"`
	McpToolDescription *string           `json:"mcp_tool_description,omitempty" xml:"mcp_tool_description,omitempty"`
	extraProperties    map[string]interface{}
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) GetMcpServerId() (value string) {
	if c == nil {
		return
	}
	return c.McpServerId
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) GetMcpServerName() (value string) {
	if c == nil {
		return
	}
	return c.McpServerName
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) GetIntegrationType() (value string) {
	if c == nil {
		return
	}
	return c.IntegrationType
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) GetParameters() (value map[string]string) {
	if c == nil {
		return
	}
	return c.Parameters
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) GetApprovalPolicy() (value string) {
	if c == nil {
		return
	}
	return c.ApprovalPolicy
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) GetRequiresApproval() (value *bool) {
	if c == nil {
		return
	}
	return c.RequiresApproval
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) GetMcpToolName() (value *string) {
	if c == nil {
		return
	}
	return c.McpToolName
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) GetMcpToolDescription() (value *string) {
	if c == nil {
		return
	}
	return c.McpToolDescription
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) SetMcpServerId(mcpServerId string) {
	c.McpServerId = mcpServerId
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) SetMcpServerName(mcpServerName string) {
	c.McpServerName = mcpServerName
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) SetIntegrationType(integrationType string) {
	c.IntegrationType = integrationType
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) SetParameters(parameters map[string]string) {
	c.Parameters = parameters
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) SetApprovalPolicy(approvalPolicy string) {
	c.ApprovalPolicy = approvalPolicy
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) SetRequiresApproval(requiresApproval *bool) {
	c.RequiresApproval = requiresApproval
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) SetMcpToolName(mcpToolName *string) {
	c.McpToolName = mcpToolName
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) SetMcpToolDescription(mcpToolDescription *string) {
	c.McpToolDescription = mcpToolDescription
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptToolCallMcpDetails) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptToolCallMcpDetails to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) UnmarshalJSON(data []byte) error {
	if err := unmarshal.ValidateRequiredJSONKeys(data, c); err != nil {
		return err
	}
	type alias ConversationHistoryTranscriptToolCallMcpDetails
	var tmp alias
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	*c = ConversationHistoryTranscriptToolCallMcpDetails(tmp)
	extraProperties, err := internal.ExtractExtraProperties(data, *c, "type")
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTranscriptToolCallMcpDetails) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptToolCallMcpDetails
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
