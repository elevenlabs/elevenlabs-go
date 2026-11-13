package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptCommonModelOutput struct {
	Role                    ConversationHistoryTranscriptCommonModelOutputRole               `json:"role" xml:"role" required:"true"`
	AgentMetadata           *AgentMetadata                                                   `json:"agent_metadata,omitempty" xml:"agent_metadata,omitempty"`
	Message                 *string                                                          `json:"message,omitempty" xml:"message,omitempty"`
	MultivoiceMessage       *ConversationHistoryMultivoiceMessageModel                       `json:"multivoice_message,omitempty" xml:"multivoice_message,omitempty"`
	ToolCalls               []*ConversationHistoryTranscriptToolCallCommonModelOutput        `json:"tool_calls,omitempty" xml:"tool_calls,omitempty"`
	ToolResults             []*ConversationHistoryTranscriptCommonModelOutputToolResultsItem `json:"tool_results,omitempty" xml:"tool_results,omitempty"`
	Feedback                *UserFeedback                                                    `json:"feedback,omitempty" xml:"feedback,omitempty"`
	LlmOverride             *string                                                          `json:"llm_override,omitempty" xml:"llm_override,omitempty"`
	ProducingLlm            *string                                                          `json:"producing_llm,omitempty" xml:"producing_llm,omitempty"`
	TimeInCallSecs          int                                                              `json:"time_in_call_secs" xml:"time_in_call_secs" required:"true"`
	ConversationTurnMetrics *ConversationTurnMetrics                                         `json:"conversation_turn_metrics,omitempty" xml:"conversation_turn_metrics,omitempty"`
	RagRetrievalInfo        *RagRetrievalInfo                                                `json:"rag_retrieval_info,omitempty" xml:"rag_retrieval_info,omitempty"`
	LlmUsage                *LlmUsageOutput                                                  `json:"llm_usage,omitempty" xml:"llm_usage,omitempty"`
	Interrupted             *bool                                                            `json:"interrupted,omitempty" xml:"interrupted,omitempty"`
	IgnoredAsBackchannel    *bool                                                            `json:"ignored_as_backchannel,omitempty" xml:"ignored_as_backchannel,omitempty"`
	PlatformEvent           *TranscriptPlatformEvent                                         `json:"platform_event,omitempty" xml:"platform_event,omitempty"`
	OriginalMessage         *string                                                          `json:"original_message,omitempty" xml:"original_message,omitempty"`
	Reasoning               []*ConversationReasoningModel                                    `json:"reasoning,omitempty" xml:"reasoning,omitempty"`
	SourceMedium            *ChatSourceMedium                                                `json:"source_medium,omitempty" xml:"source_medium,omitempty"`
	SourceEventId           *int                                                             `json:"source_event_id,omitempty" xml:"source_event_id,omitempty"`
	UsedStaticKbDocumentIds []string                                                         `json:"used_static_kb_document_ids,omitempty" xml:"used_static_kb_document_ids,omitempty"`
	UserIdentifier          *string                                                          `json:"user_identifier,omitempty" xml:"user_identifier,omitempty"`
	Id                      *string                                                          `json:"id,omitempty" xml:"id,omitempty"`
	CreatedAtUnixMs         *int                                                             `json:"created_at_unix_ms,omitempty" xml:"created_at_unix_ms,omitempty"`
	TriggeredGuardrails     []*TriggeredGuardrailCommonModel                                 `json:"triggered_guardrails,omitempty" xml:"triggered_guardrails,omitempty"`
	extraProperties         map[string]interface{}
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetRole() (value ConversationHistoryTranscriptCommonModelOutputRole) {
	if c == nil {
		return
	}
	return c.Role
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetAgentMetadata() (value *AgentMetadata) {
	if c == nil {
		return
	}
	return c.AgentMetadata
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetMessage() (value *string) {
	if c == nil {
		return
	}
	return c.Message
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetMultivoiceMessage() (value *ConversationHistoryMultivoiceMessageModel) {
	if c == nil {
		return
	}
	return c.MultivoiceMessage
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetToolCalls() (value []*ConversationHistoryTranscriptToolCallCommonModelOutput) {
	if c == nil {
		return
	}
	return c.ToolCalls
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetToolResults() (value []*ConversationHistoryTranscriptCommonModelOutputToolResultsItem) {
	if c == nil {
		return
	}
	return c.ToolResults
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetFeedback() (value *UserFeedback) {
	if c == nil {
		return
	}
	return c.Feedback
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetLlmOverride() (value *string) {
	if c == nil {
		return
	}
	return c.LlmOverride
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetProducingLlm() (value *string) {
	if c == nil {
		return
	}
	return c.ProducingLlm
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetTimeInCallSecs() (value int) {
	if c == nil {
		return
	}
	return c.TimeInCallSecs
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetConversationTurnMetrics() (value *ConversationTurnMetrics) {
	if c == nil {
		return
	}
	return c.ConversationTurnMetrics
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetRagRetrievalInfo() (value *RagRetrievalInfo) {
	if c == nil {
		return
	}
	return c.RagRetrievalInfo
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetLlmUsage() (value *LlmUsageOutput) {
	if c == nil {
		return
	}
	return c.LlmUsage
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetInterrupted() (value *bool) {
	if c == nil {
		return
	}
	return c.Interrupted
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetIgnoredAsBackchannel() (value *bool) {
	if c == nil {
		return
	}
	return c.IgnoredAsBackchannel
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetPlatformEvent() (value *TranscriptPlatformEvent) {
	if c == nil {
		return
	}
	return c.PlatformEvent
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetOriginalMessage() (value *string) {
	if c == nil {
		return
	}
	return c.OriginalMessage
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetReasoning() (value []*ConversationReasoningModel) {
	if c == nil {
		return
	}
	return c.Reasoning
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetSourceMedium() (value *ChatSourceMedium) {
	if c == nil {
		return
	}
	return c.SourceMedium
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetSourceEventId() (value *int) {
	if c == nil {
		return
	}
	return c.SourceEventId
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetUsedStaticKbDocumentIds() (value []string) {
	if c == nil {
		return
	}
	return c.UsedStaticKbDocumentIds
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetUserIdentifier() (value *string) {
	if c == nil {
		return
	}
	return c.UserIdentifier
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetId() (value *string) {
	if c == nil {
		return
	}
	return c.Id
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetCreatedAtUnixMs() (value *int) {
	if c == nil {
		return
	}
	return c.CreatedAtUnixMs
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetTriggeredGuardrails() (value []*TriggeredGuardrailCommonModel) {
	if c == nil {
		return
	}
	return c.TriggeredGuardrails
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetRole(role ConversationHistoryTranscriptCommonModelOutputRole) {
	c.Role = role
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetAgentMetadata(agentMetadata *AgentMetadata) {
	c.AgentMetadata = agentMetadata
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetMessage(message *string) {
	c.Message = message
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetMultivoiceMessage(multivoiceMessage *ConversationHistoryMultivoiceMessageModel) {
	c.MultivoiceMessage = multivoiceMessage
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetToolCalls(toolCalls []*ConversationHistoryTranscriptToolCallCommonModelOutput) {
	c.ToolCalls = toolCalls
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetToolResults(toolResults []*ConversationHistoryTranscriptCommonModelOutputToolResultsItem) {
	c.ToolResults = toolResults
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetFeedback(feedback *UserFeedback) {
	c.Feedback = feedback
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetLlmOverride(llmOverride *string) {
	c.LlmOverride = llmOverride
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetProducingLlm(producingLlm *string) {
	c.ProducingLlm = producingLlm
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetTimeInCallSecs(timeInCallSecs int) {
	c.TimeInCallSecs = timeInCallSecs
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetConversationTurnMetrics(conversationTurnMetrics *ConversationTurnMetrics) {
	c.ConversationTurnMetrics = conversationTurnMetrics
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetRagRetrievalInfo(ragRetrievalInfo *RagRetrievalInfo) {
	c.RagRetrievalInfo = ragRetrievalInfo
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetLlmUsage(llmUsage *LlmUsageOutput) {
	c.LlmUsage = llmUsage
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetInterrupted(interrupted *bool) {
	c.Interrupted = interrupted
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetIgnoredAsBackchannel(ignoredAsBackchannel *bool) {
	c.IgnoredAsBackchannel = ignoredAsBackchannel
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetPlatformEvent(platformEvent *TranscriptPlatformEvent) {
	c.PlatformEvent = platformEvent
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetOriginalMessage(originalMessage *string) {
	c.OriginalMessage = originalMessage
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetReasoning(reasoning []*ConversationReasoningModel) {
	c.Reasoning = reasoning
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetSourceMedium(sourceMedium *ChatSourceMedium) {
	c.SourceMedium = sourceMedium
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetSourceEventId(sourceEventId *int) {
	c.SourceEventId = sourceEventId
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetUsedStaticKbDocumentIds(usedStaticKbDocumentIds []string) {
	c.UsedStaticKbDocumentIds = usedStaticKbDocumentIds
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetUserIdentifier(userIdentifier *string) {
	c.UserIdentifier = userIdentifier
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetId(id *string) {
	c.Id = id
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetCreatedAtUnixMs(createdAtUnixMs *int) {
	c.CreatedAtUnixMs = createdAtUnixMs
}

func (c *ConversationHistoryTranscriptCommonModelOutput) SetTriggeredGuardrails(triggeredGuardrails []*TriggeredGuardrailCommonModel) {
	c.TriggeredGuardrails = triggeredGuardrails
}

func (c *ConversationHistoryTranscriptCommonModelOutput) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptCommonModelOutput) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptCommonModelOutput to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptCommonModelOutput) UnmarshalJSON(data []byte) error {
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

func (c *ConversationHistoryTranscriptCommonModelOutput) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptCommonModelOutput
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
