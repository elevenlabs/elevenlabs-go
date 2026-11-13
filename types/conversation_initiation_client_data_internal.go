package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationInitiationClientDataInternal struct {
	ConversationConfigOverride *ConversationConfigClientOverrideOutput `json:"conversation_config_override,omitempty" xml:"conversation_config_override,omitempty"`
	CustomLlmExtraBody         map[string]any                          `json:"custom_llm_extra_body,omitempty"`
	// ID of the end user participating in this conversation (for agent owner's user identification)
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// Information about the source of conversation initiation
	SourceInfo *ConversationInitiationSourceInfo `json:"source_info,omitempty" xml:"source_info,omitempty"`
	// ID of the agent branch to use for this conversation
	BranchId *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	// Environment to use for resolving environment variables
	Environment *string `json:"environment,omitempty" xml:"environment,omitempty"`
	// If set, start the workflow at this node id instead of the default entry
	StartingWorkflowNodeId *string `json:"starting_workflow_node_id,omitempty" xml:"starting_workflow_node_id,omitempty"`
	// If set, only these procedures are available to the starting agent. Each ID must be attached to that agent; unknown IDs fail conversation start. An empty list disables all of that agent's procedures. Not applied after an agent transfer. Requires enable_procedure_ids_from_client.
	ProcedureIds []string `json:"procedure_ids,omitempty" xml:"procedure_ids,omitempty"`
	// If set, these free-form procedures start when the conversation starts. Accepts at most five IDs. Each ID must be available after procedure_ids is applied. Unknown IDs and structured procedures fail conversation start. Requires enable_procedure_ids_from_client.
	StartProcedureIds []string       `json:"start_procedure_ids,omitempty" xml:"start_procedure_ids,omitempty"`
	DynamicVariables  map[string]any `json:"dynamic_variables,omitempty"`
	// Orchestrator-side config: tools are identified by resolved names.
	ToolMockConfig *OrchestratorToolMockBehaviorConfig `json:"tool_mock_config,omitempty" xml:"tool_mock_config,omitempty"`
	// Per-tool response mock overrides keyed by resolved tool name, applied ahead of the tool's shared mocks. Used for test-specific mocks.
	ToolMockOverrides map[string][]*ToolResponseMockConfigOutput `json:"tool_mock_overrides,omitempty"`
	extraProperties   map[string]interface{}
}

func (c *ConversationInitiationClientDataInternal) GetConversationConfigOverride() (value *ConversationConfigClientOverrideOutput) {
	if c == nil {
		return
	}
	return c.ConversationConfigOverride
}

func (c *ConversationInitiationClientDataInternal) GetCustomLlmExtraBody() (value map[string]any) {
	if c == nil {
		return
	}
	return c.CustomLlmExtraBody
}

func (c *ConversationInitiationClientDataInternal) GetUserId() (value *string) {
	if c == nil {
		return
	}
	return c.UserId
}

func (c *ConversationInitiationClientDataInternal) GetSourceInfo() (value *ConversationInitiationSourceInfo) {
	if c == nil {
		return
	}
	return c.SourceInfo
}

func (c *ConversationInitiationClientDataInternal) GetBranchId() (value *string) {
	if c == nil {
		return
	}
	return c.BranchId
}

func (c *ConversationInitiationClientDataInternal) GetEnvironment() (value *string) {
	if c == nil {
		return
	}
	return c.Environment
}

func (c *ConversationInitiationClientDataInternal) GetStartingWorkflowNodeId() (value *string) {
	if c == nil {
		return
	}
	return c.StartingWorkflowNodeId
}

func (c *ConversationInitiationClientDataInternal) GetProcedureIds() (value []string) {
	if c == nil {
		return
	}
	return c.ProcedureIds
}

func (c *ConversationInitiationClientDataInternal) GetStartProcedureIds() (value []string) {
	if c == nil {
		return
	}
	return c.StartProcedureIds
}

func (c *ConversationInitiationClientDataInternal) GetDynamicVariables() (value map[string]any) {
	if c == nil {
		return
	}
	return c.DynamicVariables
}

func (c *ConversationInitiationClientDataInternal) GetToolMockConfig() (value *OrchestratorToolMockBehaviorConfig) {
	if c == nil {
		return
	}
	return c.ToolMockConfig
}

func (c *ConversationInitiationClientDataInternal) GetToolMockOverrides() (value map[string][]*ToolResponseMockConfigOutput) {
	if c == nil {
		return
	}
	return c.ToolMockOverrides
}

func (c *ConversationInitiationClientDataInternal) SetConversationConfigOverride(conversationConfigOverride *ConversationConfigClientOverrideOutput) {
	c.ConversationConfigOverride = conversationConfigOverride
}

func (c *ConversationInitiationClientDataInternal) SetCustomLlmExtraBody(customLlmExtraBody map[string]any) {
	c.CustomLlmExtraBody = customLlmExtraBody
}

func (c *ConversationInitiationClientDataInternal) SetUserId(userId *string) {
	c.UserId = userId
}

func (c *ConversationInitiationClientDataInternal) SetSourceInfo(sourceInfo *ConversationInitiationSourceInfo) {
	c.SourceInfo = sourceInfo
}

func (c *ConversationInitiationClientDataInternal) SetBranchId(branchId *string) {
	c.BranchId = branchId
}

func (c *ConversationInitiationClientDataInternal) SetEnvironment(environment *string) {
	c.Environment = environment
}

func (c *ConversationInitiationClientDataInternal) SetStartingWorkflowNodeId(startingWorkflowNodeId *string) {
	c.StartingWorkflowNodeId = startingWorkflowNodeId
}

func (c *ConversationInitiationClientDataInternal) SetProcedureIds(procedureIds []string) {
	c.ProcedureIds = procedureIds
}

func (c *ConversationInitiationClientDataInternal) SetStartProcedureIds(startProcedureIds []string) {
	c.StartProcedureIds = startProcedureIds
}

func (c *ConversationInitiationClientDataInternal) SetDynamicVariables(dynamicVariables map[string]any) {
	c.DynamicVariables = dynamicVariables
}

func (c *ConversationInitiationClientDataInternal) SetToolMockConfig(toolMockConfig *OrchestratorToolMockBehaviorConfig) {
	c.ToolMockConfig = toolMockConfig
}

func (c *ConversationInitiationClientDataInternal) SetToolMockOverrides(toolMockOverrides map[string][]*ToolResponseMockConfigOutput) {
	c.ToolMockOverrides = toolMockOverrides
}

func (c *ConversationInitiationClientDataInternal) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataInternal) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataInternal to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataInternal) UnmarshalJSON(data []byte) error {
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

func (c *ConversationInitiationClientDataInternal) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataInternal
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
