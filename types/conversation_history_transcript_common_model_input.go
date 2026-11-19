package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptCommonModelInput struct {
	Role                    ConversationHistoryTranscriptCommonModelInputRole               `json:"role" xml:"role" required:"true"`
	AgentMetadata           *AgentMetadata                                                  `json:"agent_metadata,omitempty" xml:"agent_metadata,omitempty"`
	Message                 *string                                                         `json:"message,omitempty" xml:"message,omitempty"`
	MultivoiceMessage       *ConversationHistoryMultivoiceMessageModel                      `json:"multivoice_message,omitempty" xml:"multivoice_message,omitempty"`
	ToolCalls               []*ConversationHistoryTranscriptToolCallCommonModelInput        `json:"tool_calls,omitempty" xml:"tool_calls,omitempty"`
	ToolResults             []*ConversationHistoryTranscriptCommonModelInputToolResultsItem `json:"tool_results,omitempty" xml:"tool_results,omitempty"`
	Feedback                *UserFeedback                                                   `json:"feedback,omitempty" xml:"feedback,omitempty"`
	LlmOverride             *string                                                         `json:"llm_override,omitempty" xml:"llm_override,omitempty"`
	ProducingLlm            *string                                                         `json:"producing_llm,omitempty" xml:"producing_llm,omitempty"`
	TimeInCallSecs          int                                                             `json:"time_in_call_secs" xml:"time_in_call_secs" required:"true"`
	ConversationTurnMetrics *ConversationTurnMetrics                                        `json:"conversation_turn_metrics,omitempty" xml:"conversation_turn_metrics,omitempty"`
	RagRetrievalInfo        *RagRetrievalInfo                                               `json:"rag_retrieval_info,omitempty" xml:"rag_retrieval_info,omitempty"`
	LlmUsage                *LlmUsageInput                                                  `json:"llm_usage,omitempty" xml:"llm_usage,omitempty"`
	Interrupted             *bool                                                           `json:"interrupted,omitempty" xml:"interrupted,omitempty"`
	IgnoredAsBackchannel    *bool                                                           `json:"ignored_as_backchannel,omitempty" xml:"ignored_as_backchannel,omitempty"`
	PlatformEvent           *TranscriptPlatformEvent                                        `json:"platform_event,omitempty" xml:"platform_event,omitempty"`
	OriginalMessage         *string                                                         `json:"original_message,omitempty" xml:"original_message,omitempty"`
	Reasoning               []*ConversationReasoningModel                                   `json:"reasoning,omitempty" xml:"reasoning,omitempty"`
	SourceMedium            *ChatSourceMedium                                               `json:"source_medium,omitempty" xml:"source_medium,omitempty"`
	SourceEventId           *int                                                            `json:"source_event_id,omitempty" xml:"source_event_id,omitempty"`
	UsedStaticKbDocumentIds []string                                                        `json:"used_static_kb_document_ids,omitempty" xml:"used_static_kb_document_ids,omitempty"`
	UserIdentifier          *string                                                         `json:"user_identifier,omitempty" xml:"user_identifier,omitempty"`
	Id                      *string                                                         `json:"id,omitempty" xml:"id,omitempty"`
	CreatedAtUnixMs         *int                                                            `json:"created_at_unix_ms,omitempty" xml:"created_at_unix_ms,omitempty"`
	TriggeredGuardrails     []*TriggeredGuardrailCommonModel                                `json:"triggered_guardrails,omitempty" xml:"triggered_guardrails,omitempty"`
	extraProperties         map[string]interface{}
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetRole() (value ConversationHistoryTranscriptCommonModelInputRole) {
	if c == nil {
		return
	}
	return c.Role
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetAgentMetadata() (value *AgentMetadata) {
	if c == nil {
		return
	}
	return c.AgentMetadata
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetMessage() (value *string) {
	if c == nil {
		return
	}
	return c.Message
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetMultivoiceMessage() (value *ConversationHistoryMultivoiceMessageModel) {
	if c == nil {
		return
	}
	return c.MultivoiceMessage
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetToolCalls() (value []*ConversationHistoryTranscriptToolCallCommonModelInput) {
	if c == nil {
		return
	}
	return c.ToolCalls
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetToolResults() (value []*ConversationHistoryTranscriptCommonModelInputToolResultsItem) {
	if c == nil {
		return
	}
	return c.ToolResults
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetFeedback() (value *UserFeedback) {
	if c == nil {
		return
	}
	return c.Feedback
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetLlmOverride() (value *string) {
	if c == nil {
		return
	}
	return c.LlmOverride
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetProducingLlm() (value *string) {
	if c == nil {
		return
	}
	return c.ProducingLlm
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetTimeInCallSecs() (value int) {
	if c == nil {
		return
	}
	return c.TimeInCallSecs
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetConversationTurnMetrics() (value *ConversationTurnMetrics) {
	if c == nil {
		return
	}
	return c.ConversationTurnMetrics
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetRagRetrievalInfo() (value *RagRetrievalInfo) {
	if c == nil {
		return
	}
	return c.RagRetrievalInfo
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetLlmUsage() (value *LlmUsageInput) {
	if c == nil {
		return
	}
	return c.LlmUsage
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetInterrupted() (value *bool) {
	if c == nil {
		return
	}
	return c.Interrupted
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetIgnoredAsBackchannel() (value *bool) {
	if c == nil {
		return
	}
	return c.IgnoredAsBackchannel
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetPlatformEvent() (value *TranscriptPlatformEvent) {
	if c == nil {
		return
	}
	return c.PlatformEvent
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetOriginalMessage() (value *string) {
	if c == nil {
		return
	}
	return c.OriginalMessage
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetReasoning() (value []*ConversationReasoningModel) {
	if c == nil {
		return
	}
	return c.Reasoning
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetSourceMedium() (value *ChatSourceMedium) {
	if c == nil {
		return
	}
	return c.SourceMedium
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetSourceEventId() (value *int) {
	if c == nil {
		return
	}
	return c.SourceEventId
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetUsedStaticKbDocumentIds() (value []string) {
	if c == nil {
		return
	}
	return c.UsedStaticKbDocumentIds
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetUserIdentifier() (value *string) {
	if c == nil {
		return
	}
	return c.UserIdentifier
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetId() (value *string) {
	if c == nil {
		return
	}
	return c.Id
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetCreatedAtUnixMs() (value *int) {
	if c == nil {
		return
	}
	return c.CreatedAtUnixMs
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetTriggeredGuardrails() (value []*TriggeredGuardrailCommonModel) {
	if c == nil {
		return
	}
	return c.TriggeredGuardrails
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetRole(role ConversationHistoryTranscriptCommonModelInputRole) {
	c.Role = role
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetAgentMetadata(agentMetadata *AgentMetadata) {
	c.AgentMetadata = agentMetadata
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetMessage(message *string) {
	c.Message = message
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetMultivoiceMessage(multivoiceMessage *ConversationHistoryMultivoiceMessageModel) {
	c.MultivoiceMessage = multivoiceMessage
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetToolCalls(toolCalls []*ConversationHistoryTranscriptToolCallCommonModelInput) {
	c.ToolCalls = toolCalls
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetToolResults(toolResults []*ConversationHistoryTranscriptCommonModelInputToolResultsItem) {
	c.ToolResults = toolResults
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetFeedback(feedback *UserFeedback) {
	c.Feedback = feedback
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetLlmOverride(llmOverride *string) {
	c.LlmOverride = llmOverride
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetProducingLlm(producingLlm *string) {
	c.ProducingLlm = producingLlm
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetTimeInCallSecs(timeInCallSecs int) {
	c.TimeInCallSecs = timeInCallSecs
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetConversationTurnMetrics(conversationTurnMetrics *ConversationTurnMetrics) {
	c.ConversationTurnMetrics = conversationTurnMetrics
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetRagRetrievalInfo(ragRetrievalInfo *RagRetrievalInfo) {
	c.RagRetrievalInfo = ragRetrievalInfo
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetLlmUsage(llmUsage *LlmUsageInput) {
	c.LlmUsage = llmUsage
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetInterrupted(interrupted *bool) {
	c.Interrupted = interrupted
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetIgnoredAsBackchannel(ignoredAsBackchannel *bool) {
	c.IgnoredAsBackchannel = ignoredAsBackchannel
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetPlatformEvent(platformEvent *TranscriptPlatformEvent) {
	c.PlatformEvent = platformEvent
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetOriginalMessage(originalMessage *string) {
	c.OriginalMessage = originalMessage
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetReasoning(reasoning []*ConversationReasoningModel) {
	c.Reasoning = reasoning
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetSourceMedium(sourceMedium *ChatSourceMedium) {
	c.SourceMedium = sourceMedium
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetSourceEventId(sourceEventId *int) {
	c.SourceEventId = sourceEventId
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetUsedStaticKbDocumentIds(usedStaticKbDocumentIds []string) {
	c.UsedStaticKbDocumentIds = usedStaticKbDocumentIds
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetUserIdentifier(userIdentifier *string) {
	c.UserIdentifier = userIdentifier
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetId(id *string) {
	c.Id = id
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetCreatedAtUnixMs(createdAtUnixMs *int) {
	c.CreatedAtUnixMs = createdAtUnixMs
}

func (c *ConversationHistoryTranscriptCommonModelInput) SetTriggeredGuardrails(triggeredGuardrails []*TriggeredGuardrailCommonModel) {
	c.TriggeredGuardrails = triggeredGuardrails
}

func (c *ConversationHistoryTranscriptCommonModelInput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptCommonModelInput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptCommonModelInput to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptCommonModelInput) UnmarshalJSON(data []byte) error {
	if err := unmarshal.UnmarshalNullable(data, c); err != nil {
		return err
	}
	extraProperties, err := internal.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties
	return nil
}

func (c *ConversationHistoryTranscriptCommonModelInput) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptCommonModelInput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
