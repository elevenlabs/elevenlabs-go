package types

import (
	"encoding/json"
	"github.com/elevenlabs/elevenlabs-go/internal"
	"github.com/elevenlabs/elevenlabs-go/internal/unmarshal"
)

type ConversationHistoryTranscriptResponseModel struct {
	Role                    ConversationHistoryTranscriptResponseModelRole               `json:"role" xml:"role" required:"true"`
	AgentMetadata           *AgentMetadata                                               `json:"agent_metadata,omitempty" xml:"agent_metadata,omitempty"`
	Message                 *string                                                      `json:"message,omitempty" xml:"message,omitempty"`
	MultivoiceMessage       *ConversationHistoryMultivoiceMessageModel                   `json:"multivoice_message,omitempty" xml:"multivoice_message,omitempty"`
	ToolCalls               []*ConversationHistoryTranscriptToolCallCommonModelOutput    `json:"tool_calls,omitempty" xml:"tool_calls,omitempty"`
	ToolResults             []*ConversationHistoryTranscriptResponseModelToolResultsItem `json:"tool_results,omitempty" xml:"tool_results,omitempty"`
	Feedback                *UserFeedback                                                `json:"feedback,omitempty" xml:"feedback,omitempty"`
	LlmOverride             *string                                                      `json:"llm_override,omitempty" xml:"llm_override,omitempty"`
	ProducingLlm            *string                                                      `json:"producing_llm,omitempty" xml:"producing_llm,omitempty"`
	TimeInCallSecs          int                                                          `json:"time_in_call_secs" xml:"time_in_call_secs" required:"true"`
	ConversationTurnMetrics *ConversationTurnMetrics                                     `json:"conversation_turn_metrics,omitempty" xml:"conversation_turn_metrics,omitempty"`
	RagRetrievalInfo        *RagRetrievalInfo                                            `json:"rag_retrieval_info,omitempty" xml:"rag_retrieval_info,omitempty"`
	LlmUsage                *LlmUsageOutput                                              `json:"llm_usage,omitempty" xml:"llm_usage,omitempty"`
	Interrupted             *bool                                                        `json:"interrupted,omitempty" xml:"interrupted,omitempty"`
	IgnoredAsBackchannel    *bool                                                        `json:"ignored_as_backchannel,omitempty" xml:"ignored_as_backchannel,omitempty"`
	PlatformEvent           *TranscriptPlatformEvent                                     `json:"platform_event,omitempty" xml:"platform_event,omitempty"`
	OriginalMessage         *string                                                      `json:"original_message,omitempty" xml:"original_message,omitempty"`
	Reasoning               []*ConversationReasoningModel                                `json:"reasoning,omitempty" xml:"reasoning,omitempty"`
	SourceMedium            *ChatSourceMedium                                            `json:"source_medium,omitempty" xml:"source_medium,omitempty"`
	SourceEventId           *int                                                         `json:"source_event_id,omitempty" xml:"source_event_id,omitempty"`
	UsedStaticKbDocumentIds []string                                                     `json:"used_static_kb_document_ids,omitempty" xml:"used_static_kb_document_ids,omitempty"`
	UserIdentifier          *string                                                      `json:"user_identifier,omitempty" xml:"user_identifier,omitempty"`
	Id                      *string                                                      `json:"id,omitempty" xml:"id,omitempty"`
	CreatedAtUnixMs         *int                                                         `json:"created_at_unix_ms,omitempty" xml:"created_at_unix_ms,omitempty"`
	TriggeredGuardrails     []*TriggeredGuardrailCommonModel                             `json:"triggered_guardrails,omitempty" xml:"triggered_guardrails,omitempty"`
	// Deprecated: the first attachment on this turn. Use `file_inputs` to see every attachment.
	FileInput *ConversationHistoryTranscriptFileInputResponseModel `json:"file_input,omitempty" xml:"file_input,omitempty"`
	// All files attached to this turn, in the order the user attached them.
	FileInputs           []*ConversationHistoryTranscriptFileInputResponseModel `json:"file_inputs,omitempty" xml:"file_inputs,omitempty"`
	ContextualUpdateInfo *ContextualUpdateInfo                                  `json:"contextual_update_info,omitempty" xml:"contextual_update_info,omitempty"`
	Reasoned             *bool                                                  `json:"reasoned,omitempty" xml:"reasoned,omitempty"`
	extraProperties      map[string]interface{}
}

func (c *ConversationHistoryTranscriptResponseModel) GetRole() (value ConversationHistoryTranscriptResponseModelRole) {
	if c == nil {
		return
	}
	return c.Role
}

func (c *ConversationHistoryTranscriptResponseModel) GetAgentMetadata() (value *AgentMetadata) {
	if c == nil {
		return
	}
	return c.AgentMetadata
}

func (c *ConversationHistoryTranscriptResponseModel) GetMessage() (value *string) {
	if c == nil {
		return
	}
	return c.Message
}

func (c *ConversationHistoryTranscriptResponseModel) GetMultivoiceMessage() (value *ConversationHistoryMultivoiceMessageModel) {
	if c == nil {
		return
	}
	return c.MultivoiceMessage
}

func (c *ConversationHistoryTranscriptResponseModel) GetToolCalls() (value []*ConversationHistoryTranscriptToolCallCommonModelOutput) {
	if c == nil {
		return
	}
	return c.ToolCalls
}

func (c *ConversationHistoryTranscriptResponseModel) GetToolResults() (value []*ConversationHistoryTranscriptResponseModelToolResultsItem) {
	if c == nil {
		return
	}
	return c.ToolResults
}

func (c *ConversationHistoryTranscriptResponseModel) GetFeedback() (value *UserFeedback) {
	if c == nil {
		return
	}
	return c.Feedback
}

func (c *ConversationHistoryTranscriptResponseModel) GetLlmOverride() (value *string) {
	if c == nil {
		return
	}
	return c.LlmOverride
}

func (c *ConversationHistoryTranscriptResponseModel) GetProducingLlm() (value *string) {
	if c == nil {
		return
	}
	return c.ProducingLlm
}

func (c *ConversationHistoryTranscriptResponseModel) GetTimeInCallSecs() (value int) {
	if c == nil {
		return
	}
	return c.TimeInCallSecs
}

func (c *ConversationHistoryTranscriptResponseModel) GetConversationTurnMetrics() (value *ConversationTurnMetrics) {
	if c == nil {
		return
	}
	return c.ConversationTurnMetrics
}

func (c *ConversationHistoryTranscriptResponseModel) GetRagRetrievalInfo() (value *RagRetrievalInfo) {
	if c == nil {
		return
	}
	return c.RagRetrievalInfo
}

func (c *ConversationHistoryTranscriptResponseModel) GetLlmUsage() (value *LlmUsageOutput) {
	if c == nil {
		return
	}
	return c.LlmUsage
}

func (c *ConversationHistoryTranscriptResponseModel) GetInterrupted() (value *bool) {
	if c == nil {
		return
	}
	return c.Interrupted
}

func (c *ConversationHistoryTranscriptResponseModel) GetIgnoredAsBackchannel() (value *bool) {
	if c == nil {
		return
	}
	return c.IgnoredAsBackchannel
}

func (c *ConversationHistoryTranscriptResponseModel) GetPlatformEvent() (value *TranscriptPlatformEvent) {
	if c == nil {
		return
	}
	return c.PlatformEvent
}

func (c *ConversationHistoryTranscriptResponseModel) GetOriginalMessage() (value *string) {
	if c == nil {
		return
	}
	return c.OriginalMessage
}

func (c *ConversationHistoryTranscriptResponseModel) GetReasoning() (value []*ConversationReasoningModel) {
	if c == nil {
		return
	}
	return c.Reasoning
}

func (c *ConversationHistoryTranscriptResponseModel) GetSourceMedium() (value *ChatSourceMedium) {
	if c == nil {
		return
	}
	return c.SourceMedium
}

func (c *ConversationHistoryTranscriptResponseModel) GetSourceEventId() (value *int) {
	if c == nil {
		return
	}
	return c.SourceEventId
}

func (c *ConversationHistoryTranscriptResponseModel) GetUsedStaticKbDocumentIds() (value []string) {
	if c == nil {
		return
	}
	return c.UsedStaticKbDocumentIds
}

func (c *ConversationHistoryTranscriptResponseModel) GetUserIdentifier() (value *string) {
	if c == nil {
		return
	}
	return c.UserIdentifier
}

func (c *ConversationHistoryTranscriptResponseModel) GetId() (value *string) {
	if c == nil {
		return
	}
	return c.Id
}

func (c *ConversationHistoryTranscriptResponseModel) GetCreatedAtUnixMs() (value *int) {
	if c == nil {
		return
	}
	return c.CreatedAtUnixMs
}

func (c *ConversationHistoryTranscriptResponseModel) GetTriggeredGuardrails() (value []*TriggeredGuardrailCommonModel) {
	if c == nil {
		return
	}
	return c.TriggeredGuardrails
}

func (c *ConversationHistoryTranscriptResponseModel) GetFileInput() (value *ConversationHistoryTranscriptFileInputResponseModel) {
	if c == nil {
		return
	}
	return c.FileInput
}

func (c *ConversationHistoryTranscriptResponseModel) GetFileInputs() (value []*ConversationHistoryTranscriptFileInputResponseModel) {
	if c == nil {
		return
	}
	return c.FileInputs
}

func (c *ConversationHistoryTranscriptResponseModel) GetContextualUpdateInfo() (value *ContextualUpdateInfo) {
	if c == nil {
		return
	}
	return c.ContextualUpdateInfo
}

func (c *ConversationHistoryTranscriptResponseModel) GetReasoned() (value *bool) {
	if c == nil {
		return
	}
	return c.Reasoned
}

func (c *ConversationHistoryTranscriptResponseModel) SetRole(role ConversationHistoryTranscriptResponseModelRole) {
	c.Role = role
}

func (c *ConversationHistoryTranscriptResponseModel) SetAgentMetadata(agentMetadata *AgentMetadata) {
	c.AgentMetadata = agentMetadata
}

func (c *ConversationHistoryTranscriptResponseModel) SetMessage(message *string) {
	c.Message = message
}

func (c *ConversationHistoryTranscriptResponseModel) SetMultivoiceMessage(multivoiceMessage *ConversationHistoryMultivoiceMessageModel) {
	c.MultivoiceMessage = multivoiceMessage
}

func (c *ConversationHistoryTranscriptResponseModel) SetToolCalls(toolCalls []*ConversationHistoryTranscriptToolCallCommonModelOutput) {
	c.ToolCalls = toolCalls
}

func (c *ConversationHistoryTranscriptResponseModel) SetToolResults(toolResults []*ConversationHistoryTranscriptResponseModelToolResultsItem) {
	c.ToolResults = toolResults
}

func (c *ConversationHistoryTranscriptResponseModel) SetFeedback(feedback *UserFeedback) {
	c.Feedback = feedback
}

func (c *ConversationHistoryTranscriptResponseModel) SetLlmOverride(llmOverride *string) {
	c.LlmOverride = llmOverride
}

func (c *ConversationHistoryTranscriptResponseModel) SetProducingLlm(producingLlm *string) {
	c.ProducingLlm = producingLlm
}

func (c *ConversationHistoryTranscriptResponseModel) SetTimeInCallSecs(timeInCallSecs int) {
	c.TimeInCallSecs = timeInCallSecs
}

func (c *ConversationHistoryTranscriptResponseModel) SetConversationTurnMetrics(conversationTurnMetrics *ConversationTurnMetrics) {
	c.ConversationTurnMetrics = conversationTurnMetrics
}

func (c *ConversationHistoryTranscriptResponseModel) SetRagRetrievalInfo(ragRetrievalInfo *RagRetrievalInfo) {
	c.RagRetrievalInfo = ragRetrievalInfo
}

func (c *ConversationHistoryTranscriptResponseModel) SetLlmUsage(llmUsage *LlmUsageOutput) {
	c.LlmUsage = llmUsage
}

func (c *ConversationHistoryTranscriptResponseModel) SetInterrupted(interrupted *bool) {
	c.Interrupted = interrupted
}

func (c *ConversationHistoryTranscriptResponseModel) SetIgnoredAsBackchannel(ignoredAsBackchannel *bool) {
	c.IgnoredAsBackchannel = ignoredAsBackchannel
}

func (c *ConversationHistoryTranscriptResponseModel) SetPlatformEvent(platformEvent *TranscriptPlatformEvent) {
	c.PlatformEvent = platformEvent
}

func (c *ConversationHistoryTranscriptResponseModel) SetOriginalMessage(originalMessage *string) {
	c.OriginalMessage = originalMessage
}

func (c *ConversationHistoryTranscriptResponseModel) SetReasoning(reasoning []*ConversationReasoningModel) {
	c.Reasoning = reasoning
}

func (c *ConversationHistoryTranscriptResponseModel) SetSourceMedium(sourceMedium *ChatSourceMedium) {
	c.SourceMedium = sourceMedium
}

func (c *ConversationHistoryTranscriptResponseModel) SetSourceEventId(sourceEventId *int) {
	c.SourceEventId = sourceEventId
}

func (c *ConversationHistoryTranscriptResponseModel) SetUsedStaticKbDocumentIds(usedStaticKbDocumentIds []string) {
	c.UsedStaticKbDocumentIds = usedStaticKbDocumentIds
}

func (c *ConversationHistoryTranscriptResponseModel) SetUserIdentifier(userIdentifier *string) {
	c.UserIdentifier = userIdentifier
}

func (c *ConversationHistoryTranscriptResponseModel) SetId(id *string) {
	c.Id = id
}

func (c *ConversationHistoryTranscriptResponseModel) SetCreatedAtUnixMs(createdAtUnixMs *int) {
	c.CreatedAtUnixMs = createdAtUnixMs
}

func (c *ConversationHistoryTranscriptResponseModel) SetTriggeredGuardrails(triggeredGuardrails []*TriggeredGuardrailCommonModel) {
	c.TriggeredGuardrails = triggeredGuardrails
}

func (c *ConversationHistoryTranscriptResponseModel) SetFileInput(fileInput *ConversationHistoryTranscriptFileInputResponseModel) {
	c.FileInput = fileInput
}

func (c *ConversationHistoryTranscriptResponseModel) SetFileInputs(fileInputs []*ConversationHistoryTranscriptFileInputResponseModel) {
	c.FileInputs = fileInputs
}

func (c *ConversationHistoryTranscriptResponseModel) SetContextualUpdateInfo(contextualUpdateInfo *ContextualUpdateInfo) {
	c.ContextualUpdateInfo = contextualUpdateInfo
}

func (c *ConversationHistoryTranscriptResponseModel) SetReasoned(reasoned *bool) {
	c.Reasoned = reasoned
}

func (c *ConversationHistoryTranscriptResponseModel) GetExtraProperties() map[string]interface{} {
	if c == nil {
		return nil
	}
	return c.extraProperties
}

func (c ConversationHistoryTranscriptResponseModel) String() string {
	jsonData, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return "error converting struct: ConversationHistoryTranscriptResponseModel to string"
	}
	return string(jsonData)
}

func (c *ConversationHistoryTranscriptResponseModel) UnmarshalJSON(data []byte) error {
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

func (c *ConversationHistoryTranscriptResponseModel) MarshalJSON() ([]byte, error) {
	type embed ConversationHistoryTranscriptResponseModel
	var marshaler = struct {
		embed
	}{
		embed: embed(*c),
	}
	return json.Marshal(marshaler)
}
