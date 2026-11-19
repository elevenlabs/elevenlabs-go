package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
	big "math/big"
)

type BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost struct {
	// The ID of the agent branch to use
	BranchId string `json:"-" url:"branch_id" explode:"true" serializationStyle:"form" queryParam:"branch_id" required:"true"`
	// The id of an agent. This is returned on agent creation.
	AgentId string `json:"-" url:"-"`
	// Conversation config for the draft
	ConversationConfig map[string]any `json:"conversation_config" required:"true"`
	// Platform settings for the draft
	PlatformSettings map[string]any                   `json:"platform_settings" required:"true"`
	Workflow         *types.AgentWorkflowRequestModel `json:"workflow" xml:"workflow" required:"true"`
	// Name for the draft
	Name string `json:"name" xml:"name" required:"true"`
	// Tags to help classify and filter the agent
	Tags []string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) GetBranchId() (value string) {
	if b == nil {
		return
	}
	return b.BranchId
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) GetAgentId() (value string) {
	if b == nil {
		return
	}
	return b.AgentId
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) GetConversationConfig() (value map[string]any) {
	if b == nil {
		return
	}
	return b.ConversationConfig
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) GetPlatformSettings() (value map[string]any) {
	if b == nil {
		return
	}
	return b.PlatformSettings
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) GetWorkflow() (value *types.AgentWorkflowRequestModel) {
	if b == nil {
		return
	}
	return b.Workflow
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) GetName() (value string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) GetTags() (value []string) {
	if b == nil {
		return
	}
	return b.Tags
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) SetBranchId(branchId string) {
	b.BranchId = branchId
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) SetAgentId(agentId string) {
	b.AgentId = agentId
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) SetConversationConfig(conversationConfig map[string]any) {
	b.ConversationConfig = conversationConfig
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) SetPlatformSettings(platformSettings map[string]any) {
	b.PlatformSettings = platformSettings
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) SetWorkflow(workflow *types.AgentWorkflowRequestModel) {
	b.Workflow = workflow
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) SetName(name string) {
	b.Name = name
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) SetTags(tags []string) {
	b.Tags = tags
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateAgentDraftV1ConvaiAgentsAgentIdDraftsPost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}

var (
	deleteDraftsRequestFieldBranchId = big.NewInt(1 << 0)
)

type DeleteDraftsRequest struct {
	BranchId string `json:"-" url:"branch_id,omitempty"`

	// Private bitmask of fields set to an explicit value and therefore not to be omitted
	explicitFields *big.Int `json:"-" url:"-"`
}

func (d *DeleteDraftsRequest) require(field *big.Int) {
	if d.explicitFields == nil {
		d.explicitFields = big.NewInt(0)
	}
	d.explicitFields.Or(d.explicitFields, field)
}

// SetBranchId sets the BranchId field and marks it as non-optional;
// this prevents an empty or null value for this field from being omitted during serialization.
func (d *DeleteDraftsRequest) SetBranchId(branchId string) {
	d.BranchId = branchId
	d.require(deleteDraftsRequestFieldBranchId)
}
