package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type BodyCreateAgentV1ConvaiAgentsCreatePost struct {
	// Deprecated: all agents are versioned. This parameter is ignored.
	EnableVersioning   *bool                 `json:"-" url:"enable_versioning,omitempty" explode:"true" serializationStyle:"form" queryParam:"enable_versioning"`
	ConversationConfig *ConversationalConfig `json:"conversation_config" xml:"conversation_config" required:"true"`
	// Platform settings for the agent are all settings that aren't related to the conversation orchestration and content.
	PlatformSettings *AgentPlatformSettingsRequestModel `json:"platform_settings,omitempty" xml:"platform_settings,omitempty"`
	Workflow         *AgentWorkflowRequestModel         `json:"workflow,omitempty" xml:"workflow,omitempty"`
	// A name to make the agent easier to find
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Tags to help classify and filter the agent
	Tags []string `json:"tags,omitempty" xml:"tags,omitempty"`
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) GetEnableVersioning() (value *bool) {
	if b == nil {
		return
	}
	return b.EnableVersioning
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) GetConversationConfig() (value *ConversationalConfig) {
	if b == nil {
		return
	}
	return b.ConversationConfig
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) GetPlatformSettings() (value *AgentPlatformSettingsRequestModel) {
	if b == nil {
		return
	}
	return b.PlatformSettings
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) GetWorkflow() (value *AgentWorkflowRequestModel) {
	if b == nil {
		return
	}
	return b.Workflow
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) GetName() (value *string) {
	if b == nil {
		return
	}
	return b.Name
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) GetTags() (value []string) {
	if b == nil {
		return
	}
	return b.Tags
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) SetEnableVersioning(enableVersioning *bool) {
	b.EnableVersioning = enableVersioning
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) SetConversationConfig(conversationConfig *ConversationalConfig) {
	b.ConversationConfig = conversationConfig
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) SetPlatformSettings(platformSettings *AgentPlatformSettingsRequestModel) {
	b.PlatformSettings = platformSettings
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) SetWorkflow(workflow *AgentWorkflowRequestModel) {
	b.Workflow = workflow
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) SetName(name *string) {
	b.Name = name
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) SetTags(tags []string) {
	b.Tags = tags
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) GetExtraProperties() map[string]interface{} {
	return nil
}

func (b BodyCreateAgentV1ConvaiAgentsCreatePost) String() string {
	jsonData, err := json.MarshalIndent(b, "", "  ")
	if err != nil {
		return "error converting struct: BodyCreateAgentV1ConvaiAgentsCreatePost to string"
	}
	return string(jsonData)
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, b)
}

func (b *BodyCreateAgentV1ConvaiAgentsCreatePost) MarshalJSON() ([]byte, error) {
	type embed BodyCreateAgentV1ConvaiAgentsCreatePost
	var marshaler = struct {
		embed
	}{
		embed: embed(*b),
	}
	return json.Marshal(marshaler)
}
