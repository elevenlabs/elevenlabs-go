package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentWorkspaceOverridesOutput struct {
	// The webhook to send conversation initiation client data to
	ConversationInitiationClientDataWebhook *ConversationInitiationClientDataWebhook `json:"conversation_initiation_client_data_webhook,omitempty" xml:"conversation_initiation_client_data_webhook,omitempty"`
	Webhooks                                *ConvAiWebhooks                          `json:"webhooks,omitempty" xml:"webhooks,omitempty"`
	extraProperties                         map[string]interface{}
}

func (a *AgentWorkspaceOverridesOutput) GetConversationInitiationClientDataWebhook() (value *ConversationInitiationClientDataWebhook) {
	if a == nil {
		return
	}
	return a.ConversationInitiationClientDataWebhook
}

func (a *AgentWorkspaceOverridesOutput) GetWebhooks() (value *ConvAiWebhooks) {
	if a == nil {
		return
	}
	return a.Webhooks
}

func (a *AgentWorkspaceOverridesOutput) SetConversationInitiationClientDataWebhook(conversationInitiationClientDataWebhook *ConversationInitiationClientDataWebhook) {
	a.ConversationInitiationClientDataWebhook = conversationInitiationClientDataWebhook
}

func (a *AgentWorkspaceOverridesOutput) SetWebhooks(webhooks *ConvAiWebhooks) {
	a.Webhooks = webhooks
}

func (a *AgentWorkspaceOverridesOutput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentWorkspaceOverridesOutput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentWorkspaceOverridesOutput to string"
	}
	return string(jsonData)
}

func (a *AgentWorkspaceOverridesOutput) UnmarshalJSON(data []byte) error {
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

func (a *AgentWorkspaceOverridesOutput) MarshalJSON() ([]byte, error) {
	type embed AgentWorkspaceOverridesOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
