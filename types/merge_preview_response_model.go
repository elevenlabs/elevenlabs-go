package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type MergePreviewResponseModel struct {
	// The ID of the agent
	AgentId string `json:"agent_id" xml:"agent_id" required:"true"`
	// The name of the agent
	Name               string                              `json:"name" xml:"name" required:"true"`
	ConversationConfig *ConversationalConfig               `json:"conversation_config" xml:"conversation_config" required:"true"`
	Metadata           *AgentMetadataResponseModel         `json:"metadata" xml:"metadata" required:"true"`
	PlatformSettings   *AgentPlatformSettingsResponseModel `json:"platform_settings,omitempty" xml:"platform_settings,omitempty"`
	// The phone numbers of the agent
	PhoneNumbers []*MergePreviewResponseModelPhoneNumbersItem `json:"phone_numbers,omitempty" xml:"phone_numbers,omitempty"`
	// WhatsApp accounts assigned to the agent
	WhatsappAccounts []*GetWhatsAppAccountResponse `json:"whatsapp_accounts,omitempty" xml:"whatsapp_accounts,omitempty"`
	Workflow         *AgentWorkflowResponseModel   `json:"workflow,omitempty" xml:"workflow,omitempty"`
	// The access information of the agent for the user
	AccessInfo *ResourceAccessInfo `json:"access_info,omitempty" xml:"access_info,omitempty"`
	// Agent tags used to categorize the agent
	Tags []string `json:"tags,omitempty" xml:"tags,omitempty"`
	// The ID of the version the agent is on
	VersionId *string `json:"version_id,omitempty" xml:"version_id,omitempty"`
	// The ID of the branch the agent is on
	BranchId *string `json:"branch_id,omitempty" xml:"branch_id,omitempty"`
	// The ID of the main branch for this agent
	MainBranchId *string `json:"main_branch_id,omitempty" xml:"main_branch_id,omitempty"`
	// Procedures keyed by procedure_id.
	Procedures map[string]*ProcedureRefResponseModel `json:"procedures,omitempty"`
	// URL of the default hold tone played to queued callers when no custom hold audio is uploaded, so the dashboard can preview it.
	DefaultHoldAudioUrl *string `json:"default_hold_audio_url,omitempty" xml:"default_hold_audio_url,omitempty"`
	// Dot-paths of config fields where both branches modified the same field relative to their common ancestor (conflicts). Present regardless of which side wins the conflict.
	OverriddenFields []string `json:"overridden_fields,omitempty" xml:"overridden_fields,omitempty"`
	// Structured view of the same conflicts as overridden_fields, each carrying the value on the base (common ancestor), source branch, and target branch so the divergence can be presented and resolved field-by-field. A merge or rebase is rejected while any of them has no resolution in `conflict_resolutions`, unless the merge is forced.
	Conflicts []*FieldConflict `json:"conflicts,omitempty" xml:"conflicts,omitempty"`
	// True when the merge/rebase would be a no-op, i.e. the merged result is identical to the source branch tip. The rebase endpoint rejects in this case.
	SourceIdenticalToTarget *bool `json:"source_identical_to_target,omitempty" xml:"source_identical_to_target,omitempty"`
	// Tip version of the source branch (the rebased branch for a rebase) the preview was computed from.
	SourceTipVersionId *string `json:"source_tip_version_id,omitempty" xml:"source_tip_version_id,omitempty"`
	// Tip version of the target branch (main for a rebase) the preview was computed from.
	TargetTipVersionId *string `json:"target_tip_version_id,omitempty" xml:"target_tip_version_id,omitempty"`
	// Why the merge doesn't combine into a valid configuration while conflicts are left to resolve. The preview then shows the branch the merge writes unchanged.
	MergeError      *string `json:"merge_error,omitempty" xml:"merge_error,omitempty"`
	extraProperties map[string]interface{}
}

func (m *MergePreviewResponseModel) GetAgentId() (value string) {
	if m == nil {
		return
	}
	return m.AgentId
}

func (m *MergePreviewResponseModel) GetName() (value string) {
	if m == nil {
		return
	}
	return m.Name
}

func (m *MergePreviewResponseModel) GetConversationConfig() (value *ConversationalConfig) {
	if m == nil {
		return
	}
	return m.ConversationConfig
}

func (m *MergePreviewResponseModel) GetMetadata() (value *AgentMetadataResponseModel) {
	if m == nil {
		return
	}
	return m.Metadata
}

func (m *MergePreviewResponseModel) GetPlatformSettings() (value *AgentPlatformSettingsResponseModel) {
	if m == nil {
		return
	}
	return m.PlatformSettings
}

func (m *MergePreviewResponseModel) GetPhoneNumbers() (value []*MergePreviewResponseModelPhoneNumbersItem) {
	if m == nil {
		return
	}
	return m.PhoneNumbers
}

func (m *MergePreviewResponseModel) GetWhatsappAccounts() (value []*GetWhatsAppAccountResponse) {
	if m == nil {
		return
	}
	return m.WhatsappAccounts
}

func (m *MergePreviewResponseModel) GetWorkflow() (value *AgentWorkflowResponseModel) {
	if m == nil {
		return
	}
	return m.Workflow
}

func (m *MergePreviewResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if m == nil {
		return
	}
	return m.AccessInfo
}

func (m *MergePreviewResponseModel) GetTags() (value []string) {
	if m == nil {
		return
	}
	return m.Tags
}

func (m *MergePreviewResponseModel) GetVersionId() (value *string) {
	if m == nil {
		return
	}
	return m.VersionId
}

func (m *MergePreviewResponseModel) GetBranchId() (value *string) {
	if m == nil {
		return
	}
	return m.BranchId
}

func (m *MergePreviewResponseModel) GetMainBranchId() (value *string) {
	if m == nil {
		return
	}
	return m.MainBranchId
}

func (m *MergePreviewResponseModel) GetProcedures() (value map[string]*ProcedureRefResponseModel) {
	if m == nil {
		return
	}
	return m.Procedures
}

func (m *MergePreviewResponseModel) GetDefaultHoldAudioUrl() (value *string) {
	if m == nil {
		return
	}
	return m.DefaultHoldAudioUrl
}

func (m *MergePreviewResponseModel) GetOverriddenFields() (value []string) {
	if m == nil {
		return
	}
	return m.OverriddenFields
}

func (m *MergePreviewResponseModel) GetConflicts() (value []*FieldConflict) {
	if m == nil {
		return
	}
	return m.Conflicts
}

func (m *MergePreviewResponseModel) GetSourceIdenticalToTarget() (value *bool) {
	if m == nil {
		return
	}
	return m.SourceIdenticalToTarget
}

func (m *MergePreviewResponseModel) GetSourceTipVersionId() (value *string) {
	if m == nil {
		return
	}
	return m.SourceTipVersionId
}

func (m *MergePreviewResponseModel) GetTargetTipVersionId() (value *string) {
	if m == nil {
		return
	}
	return m.TargetTipVersionId
}

func (m *MergePreviewResponseModel) GetMergeError() (value *string) {
	if m == nil {
		return
	}
	return m.MergeError
}

func (m *MergePreviewResponseModel) SetAgentId(agentId string) {
	m.AgentId = agentId
}

func (m *MergePreviewResponseModel) SetName(name string) {
	m.Name = name
}

func (m *MergePreviewResponseModel) SetConversationConfig(conversationConfig *ConversationalConfig) {
	m.ConversationConfig = conversationConfig
}

func (m *MergePreviewResponseModel) SetMetadata(metadata *AgentMetadataResponseModel) {
	m.Metadata = metadata
}

func (m *MergePreviewResponseModel) SetPlatformSettings(platformSettings *AgentPlatformSettingsResponseModel) {
	m.PlatformSettings = platformSettings
}

func (m *MergePreviewResponseModel) SetPhoneNumbers(phoneNumbers []*MergePreviewResponseModelPhoneNumbersItem) {
	m.PhoneNumbers = phoneNumbers
}

func (m *MergePreviewResponseModel) SetWhatsappAccounts(whatsappAccounts []*GetWhatsAppAccountResponse) {
	m.WhatsappAccounts = whatsappAccounts
}

func (m *MergePreviewResponseModel) SetWorkflow(workflow *AgentWorkflowResponseModel) {
	m.Workflow = workflow
}

func (m *MergePreviewResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	m.AccessInfo = accessInfo
}

func (m *MergePreviewResponseModel) SetTags(tags []string) {
	m.Tags = tags
}

func (m *MergePreviewResponseModel) SetVersionId(versionId *string) {
	m.VersionId = versionId
}

func (m *MergePreviewResponseModel) SetBranchId(branchId *string) {
	m.BranchId = branchId
}

func (m *MergePreviewResponseModel) SetMainBranchId(mainBranchId *string) {
	m.MainBranchId = mainBranchId
}

func (m *MergePreviewResponseModel) SetProcedures(procedures map[string]*ProcedureRefResponseModel) {
	m.Procedures = procedures
}

func (m *MergePreviewResponseModel) SetDefaultHoldAudioUrl(defaultHoldAudioUrl *string) {
	m.DefaultHoldAudioUrl = defaultHoldAudioUrl
}

func (m *MergePreviewResponseModel) SetOverriddenFields(overriddenFields []string) {
	m.OverriddenFields = overriddenFields
}

func (m *MergePreviewResponseModel) SetConflicts(conflicts []*FieldConflict) {
	m.Conflicts = conflicts
}

func (m *MergePreviewResponseModel) SetSourceIdenticalToTarget(sourceIdenticalToTarget *bool) {
	m.SourceIdenticalToTarget = sourceIdenticalToTarget
}

func (m *MergePreviewResponseModel) SetSourceTipVersionId(sourceTipVersionId *string) {
	m.SourceTipVersionId = sourceTipVersionId
}

func (m *MergePreviewResponseModel) SetTargetTipVersionId(targetTipVersionId *string) {
	m.TargetTipVersionId = targetTipVersionId
}

func (m *MergePreviewResponseModel) SetMergeError(mergeError *string) {
	m.MergeError = mergeError
}

func (m *MergePreviewResponseModel) GetExtraProperties() map[string]interface{} {
	if m == nil {
		return nil
	}
	return m.extraProperties
}

func (m MergePreviewResponseModel) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MergePreviewResponseModel to string"
	}
	return string(jsonData)
}

func (m *MergePreviewResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, m); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *m)
	if err != nil {
		return err
	}
	m.extraProperties = extraProperties
	return nil
}

func (m *MergePreviewResponseModel) MarshalJSON() ([]byte, error) {
	type embed MergePreviewResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}
