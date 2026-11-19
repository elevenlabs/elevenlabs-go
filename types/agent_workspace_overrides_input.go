package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type AgentWorkspaceOverridesInput struct {
	// The webhook to send conversation initiation client data to
	ConversationInitiationClientDataWebhook *ConversationInitiationClientDataWebhook `json:"conversation_initiation_client_data_webhook,omitempty" xml:"conversation_initiation_client_data_webhook,omitempty"`
	Webhooks                                *ConvAiWebhooks                          `json:"webhooks,omitempty" xml:"webhooks,omitempty"`
	extraProperties                         map[string]interface{}
}

func (a *AgentWorkspaceOverridesInput) GetConversationInitiationClientDataWebhook() (value *ConversationInitiationClientDataWebhook) {
	if a == nil {
		return
	}
	return a.ConversationInitiationClientDataWebhook
}

func (a *AgentWorkspaceOverridesInput) GetWebhooks() (value *ConvAiWebhooks) {
	if a == nil {
		return
	}
	return a.Webhooks
}

func (a *AgentWorkspaceOverridesInput) SetConversationInitiationClientDataWebhook(conversationInitiationClientDataWebhook *ConversationInitiationClientDataWebhook) {
	a.ConversationInitiationClientDataWebhook = conversationInitiationClientDataWebhook
}

func (a *AgentWorkspaceOverridesInput) SetWebhooks(webhooks *ConvAiWebhooks) {
	a.Webhooks = webhooks
}

func (a *AgentWorkspaceOverridesInput) GetExtraProperties() map[string]interface{} {
	if a == nil {
		return nil
	}
	return a.extraProperties
}

func (a AgentWorkspaceOverridesInput) String() string {
	jsonData, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return "error converting struct: AgentWorkspaceOverridesInput to string"
	}
	return string(jsonData)
}

func (a *AgentWorkspaceOverridesInput) UnmarshalJSON(data []byte) error {
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

func (a *AgentWorkspaceOverridesInput) MarshalJSON() ([]byte, error) {
	type embed AgentWorkspaceOverridesInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*a),
	}
	return json.Marshal(marshaler)
}
