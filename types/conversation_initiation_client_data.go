package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
)

type ConversationInitiationClientData struct {
	ConversationConfigOverride *ConversationInitiationClientDataConversationConfigOverride `json:"conversation_config_override,omitempty" xml:"conversation_config_override,omitempty"`
	CustomLlmExtraBody         map[string]any                                              `json:"custom_llm_extra_body,omitempty" xml:"custom_llm_extra_body,omitempty"`
	// ID of the end user participating in this conversation (for agent owner's user identification)
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// Information about the source of conversation initiation
	SourceInfo *ConversationInitiationClientDataSourceInfo `json:"source_info,omitempty" xml:"source_info,omitempty"`
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
	DynamicVariables  map[string]any `json:"dynamic_variables,omitempty" xml:"dynamic_variables,omitempty"`
	Type              *string        `json:"type,omitempty" xml:"type,omitempty"`
	extraProperties   map[string]interface{}
}

func (c *ConversationInitiationClientData) GetConversationConfigOverride() (value *ConversationInitiationClientDataConversationConfigOverride) {
	if c == nil {
		return
	}
	return c.ConversationConfigOverride
}

func (c *ConversationInitiationClientData) GetCustomLlmExtraBody() (value map[string]any) {
	if c == nil {
		return
	}
	return c.CustomLlmExtraBody
}

func (c *ConversationInitiationClientData) GetUserId() (value *string) {
	if c == nil {
		return
	}
	return c.UserId
}

func (c *ConversationInitiationClientData) GetSourceInfo() (value *ConversationInitiationClientDataSourceInfo) {
	if c == nil {
		return
	}
	return c.SourceInfo
}

func (c *ConversationInitiationClientData) GetBranchId() (value *string) {
	if c == nil {
		return
	}
	return c.BranchId
}

func (c *ConversationInitiationClientData) GetEnvironment() (value *string) {
	if c == nil {
		return
	}
	return c.Environment
}

func (c *ConversationInitiationClientData) GetStartingWorkflowNodeId() (value *string) {
	if c == nil {
		return
	}
	return c.StartingWorkflowNodeId
}

func (c *ConversationInitiationClientData) GetProcedureIds() (value []string) {
	if c == nil {
		return
	}
	return c.ProcedureIds
}

func (c *ConversationInitiationClientData) GetStartProcedureIds() (value []string) {
	if c == nil {
		return
	}
	return c.StartProcedureIds
}

func (c *ConversationInitiationClientData) GetDynamicVariables() (value map[string]any) {
	if c == nil {
		return
	}
	return c.DynamicVariables
}

func (c *ConversationInitiationClientData) GetType() (value *string) {
	if c == nil {
		return
	}
	return c.Type
}

func (c *ConversationInitiationClientData) SetConversationConfigOverride(conversationConfigOverride *ConversationInitiationClientDataConversationConfigOverride) {
	c.ConversationConfigOverride = conversationConfigOverride
}

func (c *ConversationInitiationClientData) SetCustomLlmExtraBody(customLlmExtraBody map[string]any) {
	c.CustomLlmExtraBody = customLlmExtraBody
}

func (c *ConversationInitiationClientData) SetUserId(userId *string) {
	c.UserId = userId
}

func (c *ConversationInitiationClientData) SetSourceInfo(sourceInfo *ConversationInitiationClientDataSourceInfo) {
	c.SourceInfo = sourceInfo
}

func (c *ConversationInitiationClientData) SetBranchId(branchId *string) {
	c.BranchId = branchId
}

func (c *ConversationInitiationClientData) SetEnvironment(environment *string) {
	c.Environment = environment
}

func (c *ConversationInitiationClientData) SetStartingWorkflowNodeId(startingWorkflowNodeId *string) {
	c.StartingWorkflowNodeId = startingWorkflowNodeId
}

func (c *ConversationInitiationClientData) SetProcedureIds(procedureIds []string) {
	c.ProcedureIds = procedureIds
}

func (c *ConversationInitiationClientData) SetStartProcedureIds(startProcedureIds []string) {
	c.StartProcedureIds = startProcedureIds
}

func (c *ConversationInitiationClientData) SetDynamicVariables(dynamicVariables map[string]any) {
	c.DynamicVariables = dynamicVariables
}

func (c *ConversationInitiationClientData) SetType(type_ *string) {
	c.Type = type_
}

func (c *ConversationInitiationClientData) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationInitiationClientData) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationInitiationClientData to string"
	}
	return string(jsonData)
}

func (c *ConversationInitiationClientData) UnmarshalJSON(data []byte) error {
	type unmarshaler ConversationInitiationClientData
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = ConversationInitiationClientData(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationInitiationClientData) MarshalJSON() ([]byte, error) {
	type embed ConversationInitiationClientData
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
