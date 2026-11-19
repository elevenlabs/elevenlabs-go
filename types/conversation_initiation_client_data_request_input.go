package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationInitiationClientDataRequestInput struct {
	ConversationConfigOverride *ConversationConfigClientOverrideInput `json:"conversation_config_override,omitempty" xml:"conversation_config_override,omitempty"`
	CustomLlmExtraBody         map[string]any                         `json:"custom_llm_extra_body,omitempty"`
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
	extraProperties   map[string]interface{}
}

func (c *ConversationInitiationClientDataRequestInput) GetConversationConfigOverride() (value *ConversationConfigClientOverrideInput) {
	if c == nil {
		return
	}
	return c.ConversationConfigOverride
}

func (c *ConversationInitiationClientDataRequestInput) GetCustomLlmExtraBody() (value map[string]any) {
	if c == nil {
		return
	}
	return c.CustomLlmExtraBody
}

func (c *ConversationInitiationClientDataRequestInput) GetUserId() (value *string) {
	if c == nil {
		return
	}
	return c.UserId
}

func (c *ConversationInitiationClientDataRequestInput) GetSourceInfo() (value *ConversationInitiationSourceInfo) {
	if c == nil {
		return
	}
	return c.SourceInfo
}

func (c *ConversationInitiationClientDataRequestInput) GetBranchId() (value *string) {
	if c == nil {
		return
	}
	return c.BranchId
}

func (c *ConversationInitiationClientDataRequestInput) GetEnvironment() (value *string) {
	if c == nil {
		return
	}
	return c.Environment
}

func (c *ConversationInitiationClientDataRequestInput) GetStartingWorkflowNodeId() (value *string) {
	if c == nil {
		return
	}
	return c.StartingWorkflowNodeId
}

func (c *ConversationInitiationClientDataRequestInput) GetProcedureIds() (value []string) {
	if c == nil {
		return
	}
	return c.ProcedureIds
}

func (c *ConversationInitiationClientDataRequestInput) GetStartProcedureIds() (value []string) {
	if c == nil {
		return
	}
	return c.StartProcedureIds
}

func (c *ConversationInitiationClientDataRequestInput) GetDynamicVariables() (value map[string]any) {
	if c == nil {
		return
	}
	return c.DynamicVariables
}

func (c *ConversationInitiationClientDataRequestInput) SetConversationConfigOverride(conversationConfigOverride *ConversationConfigClientOverrideInput) {
	c.ConversationConfigOverride = conversationConfigOverride
}

func (c *ConversationInitiationClientDataRequestInput) SetCustomLlmExtraBody(customLlmExtraBody map[string]any) {
	c.CustomLlmExtraBody = customLlmExtraBody
}

func (c *ConversationInitiationClientDataRequestInput) SetUserId(userId *string) {
	c.UserId = userId
}

func (c *ConversationInitiationClientDataRequestInput) SetSourceInfo(sourceInfo *ConversationInitiationSourceInfo) {
	c.SourceInfo = sourceInfo
}

func (c *ConversationInitiationClientDataRequestInput) SetBranchId(branchId *string) {
	c.BranchId = branchId
}

func (c *ConversationInitiationClientDataRequestInput) SetEnvironment(environment *string) {
	c.Environment = environment
}

func (c *ConversationInitiationClientDataRequestInput) SetStartingWorkflowNodeId(startingWorkflowNodeId *string) {
	c.StartingWorkflowNodeId = startingWorkflowNodeId
}

func (c *ConversationInitiationClientDataRequestInput) SetProcedureIds(procedureIds []string) {
	c.ProcedureIds = procedureIds
}

func (c *ConversationInitiationClientDataRequestInput) SetStartProcedureIds(startProcedureIds []string) {
	c.StartProcedureIds = startProcedureIds
}

func (c *ConversationInitiationClientDataRequestInput) SetDynamicVariables(dynamicVariables map[string]any) {
	c.DynamicVariables = dynamicVariables
}

func (c *ConversationInitiationClientDataRequestInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientDataRequestInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientDataRequestInput to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientDataRequestInput) UnmarshalJSON(data []byte) error {
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

func (c *ConversationInitiationClientDataRequestInput) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientDataRequestInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
