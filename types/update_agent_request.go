package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type UpdateAgentRequest struct {
	// Whether to include the workflow nodes and edges generated from structured procedures. When false, the workflow contains only authored elements.
	IncludeProcedureGeneratedWorkflowElements *bool `json:"-" url:"include_procedure_generated_workflow_elements,omitempty" explode:"true" serializationStyle:"form" queryParam:"include_procedure_generated_workflow_elements"`
	// The ID of the branch to use
	BranchId *string `json:"-" url:"branch_id,omitempty" explode:"true" serializationStyle:"form" queryParam:"branch_id"`
	// Deprecated: all agents are versioned. This parameter is ignored.
	EnableVersioningIfNotEnabled *bool `json:"-" url:"enable_versioning_if_not_enabled,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_versioning_if_not_enabled"`
	// The id of an agent. This is returned on agent creation.
	AgentId            string                             `json:"-" url:"-"`
	ConversationConfig *ConversationalConfig              `json:"conversation_config,omitempty" xml:"conversation_config,omitempty"`
	PlatformSettings   *AgentPlatformSettingsRequestModel `json:"platform_settings,omitempty" xml:"platform_settings,omitempty"`
	Workflow           *AgentWorkflowRequestModel         `json:"workflow,omitempty" xml:"workflow,omitempty"`
	// A name to make the agent easier to find
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Tags to help classify and filter the agent
	Tags []string `json:"tags,omitempty" xml:"tags,omitempty"`
	// Description for this version when publishing changes (only applicable for versioned agents)
	VersionDescription *string `json:"version_description,omitempty" xml:"version_description,omitempty"`
	// Procedure versions to publish, keyed by procedure_id. When provided, this map replaces the procedures from the current draft or branch tip. When omitted or null, unpublished procedure edits are used if present; otherwise, the branch tip's procedures are retained. Pass an empty object to remove all procedures.
	Procedures map[string]*ProcedureVersionRef `json:"procedures,omitempty"`
}

func (u *UpdateAgentRequest) GetIncludeProcedureGeneratedWorkflowElements() (value *bool) {
	if u == nil {
		return
	}
	return u.IncludeProcedureGeneratedWorkflowElements
}

func (u *UpdateAgentRequest) GetBranchId() (value *string) {
	if u == nil {
		return
	}
	return u.BranchId
}

func (u *UpdateAgentRequest) GetEnableVersioningIfNotEnabled() (value *bool) {
	if u == nil {
		return
	}
	return u.EnableVersioningIfNotEnabled
}

func (u *UpdateAgentRequest) GetAgentId() (value string) {
	if u == nil {
		return
	}
	return u.AgentId
}

func (u *UpdateAgentRequest) GetConversationConfig() (value *ConversationalConfig) {
	if u == nil {
		return
	}
	return u.ConversationConfig
}

func (u *UpdateAgentRequest) GetPlatformSettings() (value *AgentPlatformSettingsRequestModel) {
	if u == nil {
		return
	}
	return u.PlatformSettings
}

func (u *UpdateAgentRequest) GetWorkflow() (value *AgentWorkflowRequestModel) {
	if u == nil {
		return
	}
	return u.Workflow
}

func (u *UpdateAgentRequest) GetName() (value *string) {
	if u == nil {
		return
	}
	return u.Name
}

func (u *UpdateAgentRequest) GetTags() (value []string) {
	if u == nil {
		return
	}
	return u.Tags
}

func (u *UpdateAgentRequest) GetVersionDescription() (value *string) {
	if u == nil {
		return
	}
	return u.VersionDescription
}

func (u *UpdateAgentRequest) GetProcedures() (value map[string]*ProcedureVersionRef) {
	if u == nil {
		return
	}
	return u.Procedures
}

func (u *UpdateAgentRequest) SetIncludeProcedureGeneratedWorkflowElements(includeProcedureGeneratedWorkflowElements *bool) {
	u.IncludeProcedureGeneratedWorkflowElements = includeProcedureGeneratedWorkflowElements
}

func (u *UpdateAgentRequest) SetBranchId(branchId *string) {
	u.BranchId = branchId
}

func (u *UpdateAgentRequest) SetEnableVersioningIfNotEnabled(enableVersioningIfNotEnabled *bool) {
	u.EnableVersioningIfNotEnabled = enableVersioningIfNotEnabled
}

func (u *UpdateAgentRequest) SetAgentId(agentId string) {
	u.AgentId = agentId
}

func (u *UpdateAgentRequest) SetConversationConfig(conversationConfig *ConversationalConfig) {
	u.ConversationConfig = conversationConfig
}

func (u *UpdateAgentRequest) SetPlatformSettings(platformSettings *AgentPlatformSettingsRequestModel) {
	u.PlatformSettings = platformSettings
}

func (u *UpdateAgentRequest) SetWorkflow(workflow *AgentWorkflowRequestModel) {
	u.Workflow = workflow
}

func (u *UpdateAgentRequest) SetName(name *string) {
	u.Name = name
}

func (u *UpdateAgentRequest) SetTags(tags []string) {
	u.Tags = tags
}

func (u *UpdateAgentRequest) SetVersionDescription(versionDescription *string) {
	u.VersionDescription = versionDescription
}

func (u *UpdateAgentRequest) SetProcedures(procedures map[string]*ProcedureVersionRef) {
	u.Procedures = procedures
}

func (u *UpdateAgentRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (u UpdateAgentRequest) String() string {
	jsonData, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		return "error converting struct: UpdateAgentRequest to string"
	}
	return string(jsonData)
}

func (u *UpdateAgentRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, u)
}

func (u *UpdateAgentRequest) MarshalJSON() ([]byte, error) {
	type embed UpdateAgentRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*u),
	}
	return json.Marshal(marshaler)
}
