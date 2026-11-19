package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type GetConvAiSettingsResponseModel struct {
	ConversationInitiationClientDataWebhook *ConversationInitiationClientDataWebhook `json:"conversation_initiation_client_data_webhook,omitempty" xml:"conversation_initiation_client_data_webhook,omitempty"`
	Webhooks                                *ConvAiWebhooks                          `json:"webhooks,omitempty" xml:"webhooks,omitempty"`
	// Whether the workspace can use MCP servers
	CanUseMcpServers       *bool `json:"can_use_mcp_servers,omitempty" xml:"can_use_mcp_servers,omitempty"`
	RagRetentionPeriodDays *int  `json:"rag_retention_period_days,omitempty" xml:"rag_retention_period_days,omitempty" max:"30"`
	// Days to retain conversation embeddings. None means use the system default (30 days).
	ConversationEmbeddingRetentionDays *int              `json:"conversation_embedding_retention_days,omitempty" xml:"conversation_embedding_retention_days,omitempty" max:"365"`
	DefaultLivekitStack                *LivekitStackType `json:"default_livekit_stack,omitempty" xml:"default_livekit_stack,omitempty"`
	extraProperties                    map[string]interface{}
}

func (g *GetConvAiSettingsResponseModel) GetConversationInitiationClientDataWebhook() (value *ConversationInitiationClientDataWebhook) {
	if g == nil {
		return
	}
	return g.ConversationInitiationClientDataWebhook
}

func (g *GetConvAiSettingsResponseModel) GetWebhooks() (value *ConvAiWebhooks) {
	if g == nil {
		return
	}
	return g.Webhooks
}

func (g *GetConvAiSettingsResponseModel) GetCanUseMcpServers() (value *bool) {
	if g == nil {
		return
	}
	return g.CanUseMcpServers
}

func (g *GetConvAiSettingsResponseModel) GetRagRetentionPeriodDays() (value *int) {
	if g == nil {
		return
	}
	return g.RagRetentionPeriodDays
}

func (g *GetConvAiSettingsResponseModel) GetConversationEmbeddingRetentionDays() (value *int) {
	if g == nil {
		return
	}
	return g.ConversationEmbeddingRetentionDays
}

func (g *GetConvAiSettingsResponseModel) GetDefaultLivekitStack() (value *LivekitStackType) {
	if g == nil {
		return
	}
	return g.DefaultLivekitStack
}

func (g *GetConvAiSettingsResponseModel) SetConversationInitiationClientDataWebhook(conversationInitiationClientDataWebhook *ConversationInitiationClientDataWebhook) {
	g.ConversationInitiationClientDataWebhook = conversationInitiationClientDataWebhook
}

func (g *GetConvAiSettingsResponseModel) SetWebhooks(webhooks *ConvAiWebhooks) {
	g.Webhooks = webhooks
}

func (g *GetConvAiSettingsResponseModel) SetCanUseMcpServers(canUseMcpServers *bool) {
	g.CanUseMcpServers = canUseMcpServers
}

func (g *GetConvAiSettingsResponseModel) SetRagRetentionPeriodDays(ragRetentionPeriodDays *int) {
	g.RagRetentionPeriodDays = ragRetentionPeriodDays
}

func (g *GetConvAiSettingsResponseModel) SetConversationEmbeddingRetentionDays(conversationEmbeddingRetentionDays *int) {
	g.ConversationEmbeddingRetentionDays = conversationEmbeddingRetentionDays
}

func (g *GetConvAiSettingsResponseModel) SetDefaultLivekitStack(defaultLivekitStack *LivekitStackType) {
	g.DefaultLivekitStack = defaultLivekitStack
}

func (g *GetConvAiSettingsResponseModel) GetExtraProperties() map[string]interface{} {
	if g == nil {
		return nil
	}
	return g.extraProperties
}

func (g GetConvAiSettingsResponseModel) String() string {
	jsonData, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		return "error converting struct: GetConvAiSettingsResponseModel to string"
	}
	return string(jsonData)
}

func (g *GetConvAiSettingsResponseModel) UnmarshalJSON(data []byte) error {
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

func (g *GetConvAiSettingsResponseModel) MarshalJSON() ([]byte, error) {
	type embed GetConvAiSettingsResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*g),
	}
	return json.Marshal(marshaler)
}
