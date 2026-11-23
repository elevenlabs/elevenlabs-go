package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AdhocAgentConfigOverrideForTestRequestModel struct {
	ConversationConfig *ConversationalConfig              `json:"conversation_config,omitempty" xml:"conversation_config,omitempty"`
	PlatformSettings   *AgentPlatformSettingsRequestModel `json:"platform_settings,omitempty" xml:"platform_settings,omitempty"`
	// A workflow to run instead of the agent's. It replaces the agent's workflow whole.
	Workflow        *AgentWorkflowRequestModel `json:"workflow,omitempty" xml:"workflow,omitempty"`
	extraProperties map[string]interface{}
}

func (a *AdhocAgentConfigOverrideForTestRequestModel) GetConversationConfig() (value *ConversationalConfig) {
	if a == nil {
		return
	}
	return a.ConversationConfig
}

func (a *AdhocAgentConfigOverrideForTestRequestModel) GetPlatformSettings() (value *AgentPlatformSettingsRequestModel) {
	if a == nil {
		return
	}
	return a.PlatformSettings
}

func (a *AdhocAgentConfigOverrideForTestRequestModel) GetWorkflow() (value *AgentWorkflowRequestModel) {
	if a == nil {
		return
	}
	return a.Workflow
}

func (a *AdhocAgentConfigOverrideForTestRequestModel) SetConversationConfig(conversationConfig *ConversationalConfig) {
	a.ConversationConfig = conversationConfig
}

func (a *AdhocAgentConfigOverrideForTestRequestModel) SetPlatformSettings(platformSettings *AgentPlatformSettingsRequestModel) {
	a.PlatformSettings = platformSettings
}

func (a *AdhocAgentConfigOverrideForTestRequestModel) SetWorkflow(workflow *AgentWorkflowRequestModel) {
	a.Workflow = workflow
}

func (a *AdhocAgentConfigOverrideForTestRequestModel) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AdhocAgentConfigOverrideForTestRequestModel) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AdhocAgentConfigOverrideForTestRequestModel to string"
	}
	return string(jsonData)
}

func (a *AdhocAgentConfigOverrideForTestRequestModel) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, a); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *a)
	if err != nil {
		return err
	}
	a.extraProperties = extraProperties
	return nil
}

func (a *AdhocAgentConfigOverrideForTestRequestModel) MarshalJSON() ([]byte, error) {
	type embed AdhocAgentConfigOverrideForTestRequestModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
