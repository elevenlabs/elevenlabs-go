package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	types "github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost struct {
	// The id of an agent. This is returned on agent creation.
	AgentId string `json:"-" url:"-"`
	// ID of the version to branch from
	ParentVersionId string `json:"parent_version_id" xml:"parent_version_id" required:"true"`
	// Name of the branch. It is unique within the agent.
	Name string `json:"name" xml:"name" required:"true" maxLength:"140"`
	// Description for the branch
	Description string `json:"description" xml:"description" required:"true" maxLength:"4096"`
	// Changes to apply to conversation config
	ConversationConfig map[string]any `json:"conversation_config,omitempty"`
	// Changes to apply to platform settings
	PlatformSettings map[string]any `json:"platform_settings,omitempty"`
	// Updated workflow definition
	Workflow *types.AgentWorkflowRequestModel `json:"workflow,omitempty" xml:"workflow,omitempty"`
	// When true, the new branch uses the caller's draft procedure set instead of the branch tip. Requires parent_version_id to be the branch tip.
	IncludeDraft *bool `json:"include_draft,omitempty" xml:"include_draft,omitempty"`
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) GetParentVersionId() (value string) {
	if b == nil {
		return
	}
	return b.ParentVersionId
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) GetDescription() (value string) {
	if b == nil {
		return
	}
	return b.Description
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) GetConversationConfig() (value map[string]any) {
	if b == nil {
		return
	}
	return b.ConversationConfig
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) GetPlatformSettings() (value map[string]any) {
	if b == nil {
		return
	}
	return b.PlatformSettings
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) GetWorkflow() (value *types.AgentWorkflowRequestModel) {
	if b == nil {
		return
	}
	return b.Workflow
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) GetIncludeDraft() (value *bool) {
	if b == nil {
		return
	}
	return b.IncludeDraft
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) SetParentVersionId(parentVersionId string) {
	b.ParentVersionId = parentVersionId
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) SetName(name string) {
	b.Name = name
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) SetDescription(description string) {
	b.Description = description
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) SetConversationConfig(conversationConfig map[string]any) {
	b.ConversationConfig = conversationConfig
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) SetPlatformSettings(platformSettings map[string]any) {
	b.PlatformSettings = platformSettings
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) SetWorkflow(workflow *types.AgentWorkflowRequestModel) {
	b.Workflow = workflow
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) SetIncludeDraft(includeDraft *bool) {
	b.IncludeDraft = includeDraft
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateANewBranchV1ConvaiAgentsAgentIdBranchesPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch struct {
	// Unique identifier for the branch.
	BranchId string `json:"-" url:"-"`
	// The id of an agent. This is returned on agent creation.
	AgentId string `json:"-" url:"-"`
	// New name for the branch. Must be unique within the agent.
	Name *string `json:"name,omitempty" xml:"name,omitempty" maxLength:"140" minLength:"1"`
	// New description for the branch. Pass an empty string to clear it.
	Description *string `json:"description,omitempty" xml:"description,omitempty" maxLength:"4096"`
	// Whether the branch should be archived
	IsArchived *bool `json:"is_archived,omitempty" xml:"is_archived,omitempty"`
	// The protection level for the branch
	ProtectionStatus *types.BranchProtectionStatus `json:"protection_status,omitempty" xml:"protection_status,omitempty"`
	// When true, rebase this branch onto main automatically after each new main version. Cannot be enabled on the main branch.
	AutoRebase *bool `json:"auto_rebase,omitempty" xml:"auto_rebase,omitempty"`
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) GetBranchId() (value string) {
	if b == nil {
		return
	}
	return b.BranchId
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) GetName() (value *string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) GetDescription() (value *string) {
	if b == nil {
		return
	}
	return b.Description
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) GetIsArchived() (value *bool) {
	if b == nil {
		return
	}
	return b.IsArchived
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) GetProtectionStatus() (value *types.BranchProtectionStatus) {
	if b == nil {
		return
	}
	return b.ProtectionStatus
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) GetAutoRebase() (value *bool) {
	if b == nil {
		return
	}
	return b.AutoRebase
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) SetBranchId(branchId string) {
	b.BranchId = branchId
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) SetName(name *string) {
	b.Name = name
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) SetDescription(description *string) {
	b.Description = description
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) SetIsArchived(isArchived *bool) {
	b.IsArchived = isArchived
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) SetProtectionStatus(protectionStatus *types.BranchProtectionStatus) {
	b.ProtectionStatus = protectionStatus
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) SetAutoRebase(autoRebase *bool) {
	b.AutoRebase = autoRebase
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch to string"
	}
	return string(jsonData)
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch) MarshalJSON() ([]byte, error) {
	type embed BodyUpdateAgentBranchV1ConvaiAgentsAgentIdBranchesBranchIdPatch
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

type MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost struct {
	// Whether to archive the source branch after merging
	ArchiveSourceBranch *bool `json:"archive_source_branch,omitempty" xml:"archive_source_branch,omitempty"`
	// Keep the source branch's value for every conflict instead of sending `conflict_resolutions`
	Force *bool `json:"force,omitempty" xml:"force,omitempty"`
	// Side to keep, or custom text where the conflict accepts it, for each conflict the preview reported. A merge or rebase with conflicts is rejected unless there is one for every conflict, sent with both tip ids. On a merge it can't be combined with `force`, which lets the source win every conflict instead.
	ConflictResolutions []*types.ConflictResolution `json:"conflict_resolutions,omitempty" xml:"conflict_resolutions,omitempty"`
	// Source branch tip version the conflicts were resolved against (`source_tip_version_id` of the preview). The request fails with 409 if the branch moved since.
	SourceTipId *string `json:"source_tip_id,omitempty" xml:"source_tip_id,omitempty"`
	// Target branch tip version the conflicts were resolved against (`target_tip_version_id` of the preview). The request fails with 409 if the branch moved since.
	TargetTipId *string `json:"target_tip_id,omitempty" xml:"target_tip_id,omitempty"`
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) GetArchiveSourceBranch() (value *bool) {
	if m == nil {
		return
	}
	return m.ArchiveSourceBranch
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) GetForce() (value *bool) {
	if m == nil {
		return
	}
	return m.Force
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) GetConflictResolutions() (value []*types.ConflictResolution) {
	if m == nil {
		return
	}
	return m.ConflictResolutions
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) GetSourceTipId() (value *string) {
	if m == nil {
		return
	}
	return m.SourceTipId
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) GetTargetTipId() (value *string) {
	if m == nil {
		return
	}
	return m.TargetTipId
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) SetArchiveSourceBranch(archiveSourceBranch *bool) {
	m.ArchiveSourceBranch = archiveSourceBranch
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) SetForce(force *bool) {
	m.Force = force
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) SetConflictResolutions(conflictResolutions []*types.ConflictResolution) {
	m.ConflictResolutions = conflictResolutions
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) SetSourceTipId(sourceTipId *string) {
	m.SourceTipId = sourceTipId
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) SetTargetTipId(targetTipId *string) {
	m.TargetTipId = targetTipId
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (m MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) String() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		return "error converting struct: MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost to string"
	}
	return string(jsonData)
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, m)
}

func (m *MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) MarshalJSON() ([]byte, error) {
	type embed MergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*m),
	}
	return json.Marshal(marshaler)
}

type BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost struct {
	// The ID of the branch to pull changes from. Defaults to the agent's main branch.
	SourceBranchId *string `json:"-" url:"source_branch_id,omitempty" explode:"true" serializationStyle:"form" queryParam:"source_branch_id"`
	// Unique identifier for the source branch to merge from.
	BranchId string `json:"-" url:"-"`
	// The id of an agent. This is returned on agent creation.
	AgentId string `json:"-" url:"-"`
	// Side to keep, or custom text where the conflict accepts it, for each conflict the preview reported. A merge or rebase with conflicts is rejected unless there is one for every conflict, sent with both tip ids. On a merge it can't be combined with `force`, which lets the source win every conflict instead.
	ConflictResolutions []*types.ConflictResolution `json:"conflict_resolutions,omitempty" xml:"conflict_resolutions,omitempty"`
	// Tip version of the branch being rebased (`source_tip_version_id` of the rebase preview) the conflicts were resolved against. The request fails with 409 if it moved since.
	BranchTipId *string `json:"branch_tip_id,omitempty" xml:"branch_tip_id,omitempty"`
	// Tip version of the branch the changes are pulled from (`source_branch_id`, the main branch by default; `target_tip_version_id` of the rebase preview) the conflicts were resolved against. The request fails with 409 if it moved since.
	SourceTipId *string `json:"source_tip_id,omitempty" xml:"source_tip_id,omitempty"`
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) GetSourceBranchId() (value *string) {
	if b == nil {
		return
	}
	return b.SourceBranchId
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) GetBranchId() (value string) {
	if b == nil {
		return
	}
	return b.BranchId
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) GetConflictResolutions() (value []*types.ConflictResolution) {
	if b == nil {
		return
	}
	return b.ConflictResolutions
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) GetBranchTipId() (value *string) {
	if b == nil {
		return
	}
	return b.BranchTipId
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) GetSourceTipId() (value *string) {
	if b == nil {
		return
	}
	return b.SourceTipId
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) SetSourceBranchId(sourceBranchId *string) {
	b.SourceBranchId = sourceBranchId
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) SetBranchId(branchId string) {
	b.BranchId = branchId
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) SetConflictResolutions(conflictResolutions []*types.ConflictResolution) {
	b.ConflictResolutions = conflictResolutions
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) SetBranchTipId(branchTipId *string) {
	b.BranchTipId = branchTipId
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) SetSourceTipId(sourceTipId *string) {
	b.SourceTipId = sourceTipId
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost to string"
	}
	return string(jsonData)
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost) MarshalJSON() ([]byte, error) {
	type embed BodyRebaseABranchOntoMainV1ConvaiAgentsAgentIdBranchesBranchIdRebasePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	listBranchesRequestFieldIncludeArchived     = big.NewInt(1 << 0)
	listBranchesRequestFieldLimit               = big.NewInt(1 << 1)
	listBranchesRequestFieldIncludeCommitStatus = big.NewInt(1 << 2)
)

type ListBranchesRequest struct {
	IncludeArchived     *bool `json:"-" url:"include_archived,omitempty"`
	Limit               *int  `json:"-" url:"limit,omitempty"`
	IncludeCommitStatus *bool `json:"-" url:"include_commit_status,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (l *ListBranchesRequest) require(field *big.Int) {
	if l.explicitFields == nil {
		l.explicitFields = big.NewInt(0)
	}
	l.explicitFields.Or(l.explicitFields, field)
}

// SetIncludeArchived sets the IncludeArchived field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListBranchesRequest) SetIncludeArchived(includeArchived *bool) {
	l.IncludeArchived = includeArchived
	l.require(listBranchesRequestFieldIncludeArchived)
}

// SetLimit sets the Limit field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListBranchesRequest) SetLimit(limit *int) {
	l.Limit = limit
	l.require(listBranchesRequestFieldLimit)
}

// SetIncludeCommitStatus sets the IncludeCommitStatus field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (l *ListBranchesRequest) SetIncludeCommitStatus(includeCommitStatus *bool) {
	l.IncludeCommitStatus = includeCommitStatus
	l.require(listBranchesRequestFieldIncludeCommitStatus)
}

var (
	previewMergeBranchesRequestFieldTargetBranchId                            = big.NewInt(1 << 0)
	previewMergeBranchesRequestFieldForce                                     = big.NewInt(1 << 1)
	previewMergeBranchesRequestFieldIncludeProcedureGeneratedWorkflowElements = big.NewInt(1 << 2)
)

type PreviewMergeBranchesRequest struct {
	TargetBranchId                            string `json:"-" url:"target_branch_id,omitempty"`
	Force                                     *bool  `json:"-" url:"force,omitempty"`
	IncludeProcedureGeneratedWorkflowElements *bool  `json:"-" url:"include_procedure_generated_workflow_elements,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (p *PreviewMergeBranchesRequest) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetTargetBranchId sets the TargetBranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PreviewMergeBranchesRequest) SetTargetBranchId(targetBranchId string) {
	p.TargetBranchId = targetBranchId
	p.require(previewMergeBranchesRequestFieldTargetBranchId)
}

// SetForce sets the Force field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PreviewMergeBranchesRequest) SetForce(force *bool) {
	p.Force = force
	p.require(previewMergeBranchesRequestFieldForce)
}

// SetIncludeProcedureGeneratedWorkflowElements sets the IncludeProcedureGeneratedWorkflowElements field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PreviewMergeBranchesRequest) SetIncludeProcedureGeneratedWorkflowElements(includeProcedureGeneratedWorkflowElements *bool) {
	p.IncludeProcedureGeneratedWorkflowElements = includeProcedureGeneratedWorkflowElements
	p.require(previewMergeBranchesRequestFieldIncludeProcedureGeneratedWorkflowElements)
}

var (
	bodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostFieldArchiveSourceBranch = big.NewInt(1 << 0)
	bodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostFieldForce               = big.NewInt(1 << 1)
	bodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostFieldConflictResolutions = big.NewInt(1 << 2)
	bodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostFieldSourceTipId         = big.NewInt(1 << 3)
	bodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostFieldTargetTipId         = big.NewInt(1 << 4)
	bodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostFieldTargetBranchId      = big.NewInt(1 << 5)
)

type BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost struct {
	ArchiveSourceBranch *bool                       `json:"archive_source_branch,omitempty" url:"-"`
	Force               *bool                       `json:"force,omitempty" url:"-"`
	ConflictResolutions []*types.ConflictResolution `json:"conflict_resolutions,omitempty" url:"-"`
	SourceTipId         *string                     `json:"source_tip_id,omitempty" url:"-"`
	TargetTipId         *string                     `json:"target_tip_id,omitempty" url:"-"`
	TargetBranchId      string                      `json:"-" url:"target_branch_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (b *BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) require(field *big.Int) {
	if b.explicitFields == nil {
		b.explicitFields = big.NewInt(0)
	}
	b.explicitFields.Or(b.explicitFields, field)
}

// SetArchiveSourceBranch sets the ArchiveSourceBranch field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) SetArchiveSourceBranch(archiveSourceBranch *bool) {
	b.ArchiveSourceBranch = archiveSourceBranch
	b.require(bodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostFieldArchiveSourceBranch)
}

// SetForce sets the Force field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) SetForce(force *bool) {
	b.Force = force
	b.require(bodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostFieldForce)
}

// SetConflictResolutions sets the ConflictResolutions field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) SetConflictResolutions(conflictResolutions []*types.ConflictResolution) {
	b.ConflictResolutions = conflictResolutions
	b.require(bodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostFieldConflictResolutions)
}

// SetSourceTipId sets the SourceTipId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) SetSourceTipId(sourceTipId *string) {
	b.SourceTipId = sourceTipId
	b.require(bodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostFieldSourceTipId)
}

// SetTargetTipId sets the TargetTipId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) SetTargetTipId(targetTipId *string) {
	b.TargetTipId = targetTipId
	b.require(bodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostFieldTargetTipId)
}

// SetTargetBranchId sets the TargetBranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (b *BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) SetTargetBranchId(targetBranchId string) {
	b.TargetBranchId = targetBranchId
	b.require(bodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePostFieldTargetBranchId)
}

func (b *BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) UnmarshalJSON(data []byte) error {
	type unmarshaler BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost
	var body unmarshaler
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	*b = BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost(body)
	return nil
}

func (b *BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost) MarshalJSON() ([]byte, error) {
	type embed BodyMergeABranchIntoATargetBranchV1ConvaiAgentsAgentIdBranchesSourceBranchIdMergePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	previewRebaseBranchesRequestFieldSourceBranchId                            = big.NewInt(1 << 0)
	previewRebaseBranchesRequestFieldIncludeProcedureGeneratedWorkflowElements = big.NewInt(1 << 1)
)

type PreviewRebaseBranchesRequest struct {
	SourceBranchId                            *string `json:"-" url:"source_branch_id,omitempty"`
	IncludeProcedureGeneratedWorkflowElements *bool   `json:"-" url:"include_procedure_generated_workflow_elements,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (p *PreviewRebaseBranchesRequest) require(field *big.Int) {
	if p.explicitFields == nil {
		p.explicitFields = big.NewInt(0)
	}
	p.explicitFields.Or(p.explicitFields, field)
}

// SetSourceBranchId sets the SourceBranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PreviewRebaseBranchesRequest) SetSourceBranchId(sourceBranchId *string) {
	p.SourceBranchId = sourceBranchId
	p.require(previewRebaseBranchesRequestFieldSourceBranchId)
}

// SetIncludeProcedureGeneratedWorkflowElements sets the IncludeProcedureGeneratedWorkflowElements field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (p *PreviewRebaseBranchesRequest) SetIncludeProcedureGeneratedWorkflowElements(includeProcedureGeneratedWorkflowElements *bool) {
	p.IncludeProcedureGeneratedWorkflowElements = includeProcedureGeneratedWorkflowElements
	p.require(previewRebaseBranchesRequestFieldIncludeProcedureGeneratedWorkflowElements)
}
