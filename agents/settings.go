package agents

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
	"github.com/elevenlabs/elevenlabs-go/types"
)

type PatchConvAiSettingsRequest struct {
	ConversationInitiationClientDataWebhook *types.ConversationInitiationClientDataWebhook `json:"conversation_initiation_client_data_webhook,omitempty" xml:"conversation_initiation_client_data_webhook,omitempty"`
	Webhooks                                *types.ConvAiWebhooks                          `json:"webhooks,omitempty" xml:"webhooks,omitempty"`
	// Whether the workspace can use MCP servers
	CanUseMcpServers       *bool `json:"can_use_mcp_servers,omitempty" xml:"can_use_mcp_servers,omitempty"`
	RagRetentionPeriodDays *int  `json:"rag_retention_period_days,omitempty" xml:"rag_retention_period_days,omitempty" max:"30"`
	// Days to retain conversation embeddings. None means use the system default (30 days).
	ConversationEmbeddingRetentionDays *int                    `json:"conversation_embedding_retention_days,omitempty" xml:"conversation_embedding_retention_days,omitempty" max:"365"`
	DefaultLivekitStack                *types.LivekitStackType `json:"default_livekit_stack,omitempty" xml:"default_livekit_stack,omitempty"`
}

func (p *PatchConvAiSettingsRequest) GetConversationInitiationClientDataWebhook() (value *types.ConversationInitiationClientDataWebhook) {
	if p == nil {
		return
	}
	return p.ConversationInitiationClientDataWebhook
}

func (p *PatchConvAiSettingsRequest) GetWebhooks() (value *types.ConvAiWebhooks) {
	if p == nil {
		return
	}
	return p.Webhooks
}

func (p *PatchConvAiSettingsRequest) GetCanUseMcpServers() (value *bool) {
	if p == nil {
		return
	}
	return p.CanUseMcpServers
}

func (p *PatchConvAiSettingsRequest) GetRagRetentionPeriodDays() (value *int) {
	if p == nil {
		return
	}
	return p.RagRetentionPeriodDays
}

func (p *PatchConvAiSettingsRequest) GetConversationEmbeddingRetentionDays() (value *int) {
	if p == nil {
		return
	}
	return p.ConversationEmbeddingRetentionDays
}

func (p *PatchConvAiSettingsRequest) GetDefaultLivekitStack() (value *types.LivekitStackType) {
	if p == nil {
		return
	}
	return p.DefaultLivekitStack
}

func (p *PatchConvAiSettingsRequest) SetConversationInitiationClientDataWebhook(conversationInitiationClientDataWebhook *types.ConversationInitiationClientDataWebhook) {
	p.ConversationInitiationClientDataWebhook = conversationInitiationClientDataWebhook
}

func (p *PatchConvAiSettingsRequest) SetWebhooks(webhooks *types.ConvAiWebhooks) {
	p.Webhooks = webhooks
}

func (p *PatchConvAiSettingsRequest) SetCanUseMcpServers(canUseMcpServers *bool) {
	p.CanUseMcpServers = canUseMcpServers
}

func (p *PatchConvAiSettingsRequest) SetRagRetentionPeriodDays(ragRetentionPeriodDays *int) {
	p.RagRetentionPeriodDays = ragRetentionPeriodDays
}

func (p *PatchConvAiSettingsRequest) SetConversationEmbeddingRetentionDays(conversationEmbeddingRetentionDays *int) {
	p.ConversationEmbeddingRetentionDays = conversationEmbeddingRetentionDays
}

func (p *PatchConvAiSettingsRequest) SetDefaultLivekitStack(defaultLivekitStack *types.LivekitStackType) {
	p.DefaultLivekitStack = defaultLivekitStack
}

func (p *PatchConvAiSettingsRequest) GetExtraProperties() map[string]interface{} {
	return nil
}

func (p PatchConvAiSettingsRequest) String() string {
	jsonData, err := json.MarshalIndent(p, "", "  ")
	if err != nil {
		return "error converting struct: PatchConvAiSettingsRequest to string"
	}
	return string(jsonData)
}

func (p *PatchConvAiSettingsRequest) UnmarshalJSON(data []byte) error {
	return unmarshal.UnmarshalNullable(data, p)
}

func (p *PatchConvAiSettingsRequest) MarshalJSON() ([]byte, error) {
	type embed PatchConvAiSettingsRequest
	var marshaler = struct {
		embed
	}{
		embed: embed(*p),
	}
	return json.Marshal(marshaler)
}
