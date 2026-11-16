package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetAgentResponseModel struct {
	// The ID of the agent
	AgentId string `json:"agent_id" xml:"agent_id" required:"true"`
	// The name of the agent
	Name               string                              `json:"name" xml:"name" required:"true"`
	ConversationConfig *ConversationalConfig               `json:"conversation_config" xml:"conversation_config" required:"true"`
	Metadata           *AgentMetadataResponseModel         `json:"metadata" xml:"metadata" required:"true"`
	PlatformSettings   *AgentPlatformSettingsResponseModel `json:"platform_settings,omitempty" xml:"platform_settings,omitempty"`
	// The phone numbers of the agent
	PhoneNumbers []*GetAgentResponseModelPhoneNumbersItem `json:"phone_numbers,omitempty" xml:"phone_numbers,omitempty"`
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
	extraProperties     map[string]interface{}
}

func (g *GetAgentResponseModel) GetAgentId() (value string) {
	if g == nil {
		return
	}
	return g.AgentId
}

func (g *GetAgentResponseModel) GetName() (value string) {
	if g == nil {
		return
	}
	return g.Name
}

func (g *GetAgentResponseModel) GetConversationConfig() (value *ConversationalConfig) {
	if g == nil {
		return
	}
	return g.ConversationConfig
}

func (g *GetAgentResponseModel) GetMetadata() (value *AgentMetadataResponseModel) {
	if g == nil {
		return
	}
	return g.Metadata
}

func (g *GetAgentResponseModel) GetPlatformSettings() (value *AgentPlatformSettingsResponseModel) {
	if g == nil {
		return
	}
	return g.PlatformSettings
}

func (g *GetAgentResponseModel) GetPhoneNumbers() (value []*GetAgentResponseModelPhoneNumbersItem) {
	if g == nil {
		return
	}
	return g.PhoneNumbers
}

func (g *GetAgentResponseModel) GetWhatsappAccounts() (value []*GetWhatsAppAccountResponse) {
	if g == nil {
		return
	}
	return g.WhatsappAccounts
}

func (g *GetAgentResponseModel) GetWorkflow() (value *AgentWorkflowResponseModel) {
	if g == nil {
		return
	}
	return g.Workflow
}

func (g *GetAgentResponseModel) GetAccessInfo() (value *ResourceAccessInfo) {
	if g == nil {
		return
	}
	return g.AccessInfo
}

func (g *GetAgentResponseModel) GetTags() (value []string) {
	if g == nil {
		return
	}
	return g.Tags
}

func (g *GetAgentResponseModel) GetVersionId() (value *string) {
	if g == nil {
		return
	}
	return g.VersionId
}

func (g *GetAgentResponseModel) GetBranchId() (value *string) {
	if g == nil {
		return
	}
	return g.BranchId
}

func (g *GetAgentResponseModel) GetMainBranchId() (value *string) {
	if g == nil {
		return
	}
	return g.MainBranchId
}

func (g *GetAgentResponseModel) GetProcedures() (value map[string]*ProcedureRefResponseModel) {
	if g == nil {
		return
	}
	return g.Procedures
}

func (g *GetAgentResponseModel) GetDefaultHoldAudioUrl() (value *string) {
	if g == nil {
		return
	}
	return g.DefaultHoldAudioUrl
}

func (g *GetAgentResponseModel) SetAgentId(agentId string) {
	g.AgentId = agentId
}

func (g *GetAgentResponseModel) SetName(name string) {
	g.Name = name
}

func (g *GetAgentResponseModel) SetConversationConfig(conversationConfig *ConversationalConfig) {
	g.ConversationConfig = conversationConfig
}

func (g *GetAgentResponseModel) SetMetadata(metadata *AgentMetadataResponseModel) {
	g.Metadata = metadata
}

func (g *GetAgentResponseModel) SetPlatformSettings(platformSettings *AgentPlatformSettingsResponseModel) {
	g.PlatformSettings = platformSettings
}

func (g *GetAgentResponseModel) SetPhoneNumbers(phoneNumbers []*GetAgentResponseModelPhoneNumbersItem) {
	g.PhoneNumbers = phoneNumbers
}

func (g *GetAgentResponseModel) SetWhatsappAccounts(whatsappAccounts []*GetWhatsAppAccountResponse) {
	g.WhatsappAccounts = whatsappAccounts
}

func (g *GetAgentResponseModel) SetWorkflow(workflow *AgentWorkflowResponseModel) {
	g.Workflow = workflow
}

func (g *GetAgentResponseModel) SetAccessInfo(accessInfo *ResourceAccessInfo) {
	g.AccessInfo = accessInfo
}

func (g *GetAgentResponseModel) SetTags(tags []string) {
	g.Tags = tags
}

func (g *GetAgentResponseModel) SetVersionId(versionId *string) {
	g.VersionId = versionId
}

func (g *GetAgentResponseModel) SetBranchId(branchId *string) {
	g.BranchId = branchId
}

func (g *GetAgentResponseModel) SetMainBranchId(mainBranchId *string) {
	g.MainBranchId = mainBranchId
}

func (g *GetAgentResponseModel) SetProcedures(procedures map[string]*ProcedureRefResponseModel) {
	g.Procedures = procedures
}

func (g *GetAgentResponseModel) SetDefaultHoldAudioUrl(defaultHoldAudioUrl *string) {
	g.DefaultHoldAudioUrl = defaultHoldAudioUrl
}

func (g *GetAgentResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetAgentResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetAgentResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetAgentResponseModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, g); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties
	return nil
}

func (g *GetAgentResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetAgentResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
