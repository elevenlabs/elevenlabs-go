package types

import (
	"encoding/json"
	"fmt"
)

type SubscribeEvent struct {
	ConversationInitiationMetadata *ConversationInitiationMetadata
	AgentResponseComplete          *AgentResponseComplete
	ContextUsage                   *ContextUsage
	McpToolCall                    *McpToolCall
	ClientError                    *ClientError
	GuardrailTriggered             *GuardrailTriggered
	ConversationUserTranscript     *ConversationUserTranscript
	ConversationAgentResponse      *ConversationAgentResponse
	QueueStatus                    *QueueStatus
	AgentResponseCorrection        *AgentResponseCorrection
	AgentResponseMetadata          *AgentResponseMetadata
	AudioResponse                  *AudioResponse
	Interruption                   *Interruption
	VadScore                       *VadScore
	AgentChatResponsePart          *AgentChatResponsePart
	ClientToolCall                 *ClientToolCall
	AgentToolResponse              *AgentToolResponse
	AgentToolResponseFullPayload   *AgentToolResponseFullPayload
	AgentToolRequest               *AgentToolRequest
	ConversationPing               *ConversationPing
	McpConnectionStatus            *McpConnectionStatus
	typ                            string
}

func (s *SubscribeEvent) GetConversationInitiationMetadata() *ConversationInitiationMetadata {
	if s == nil {
		return nil
	}
	if s.ConversationInitiationMetadata == nil {
		return nil
	}
	return s.ConversationInitiationMetadata
}

func (s *SubscribeEvent) GetAgentResponseComplete() *AgentResponseComplete {
	if s == nil {
		return nil
	}
	if s.AgentResponseComplete == nil {
		return nil
	}
	return s.AgentResponseComplete
}

func (s *SubscribeEvent) GetContextUsage() *ContextUsage {
	if s == nil {
		return nil
	}
	if s.ContextUsage == nil {
		return nil
	}
	return s.ContextUsage
}

func (s *SubscribeEvent) GetMcpToolCall() *McpToolCall {
	if s == nil {
		return nil
	}
	if s.McpToolCall == nil {
		return nil
	}
	return s.McpToolCall
}

func (s *SubscribeEvent) GetClientError() *ClientError {
	if s == nil {
		return nil
	}
	if s.ClientError == nil {
		return nil
	}
	return s.ClientError
}

func (s *SubscribeEvent) GetGuardrailTriggered() *GuardrailTriggered {
	if s == nil {
		return nil
	}
	if s.GuardrailTriggered == nil {
		return nil
	}
	return s.GuardrailTriggered
}

func (s *SubscribeEvent) GetConversationUserTranscript() *ConversationUserTranscript {
	if s == nil {
		return nil
	}
	if s.ConversationUserTranscript == nil {
		return nil
	}
	return s.ConversationUserTranscript
}

func (s *SubscribeEvent) GetConversationAgentResponse() *ConversationAgentResponse {
	if s == nil {
		return nil
	}
	if s.ConversationAgentResponse == nil {
		return nil
	}
	return s.ConversationAgentResponse
}

func (s *SubscribeEvent) GetQueueStatus() *QueueStatus {
	if s == nil {
		return nil
	}
	if s.QueueStatus == nil {
		return nil
	}
	return s.QueueStatus
}

func (s *SubscribeEvent) GetAgentResponseCorrection() *AgentResponseCorrection {
	if s == nil {
		return nil
	}
	if s.AgentResponseCorrection == nil {
		return nil
	}
	return s.AgentResponseCorrection
}

func (s *SubscribeEvent) GetAgentResponseMetadata() *AgentResponseMetadata {
	if s == nil {
		return nil
	}
	if s.AgentResponseMetadata == nil {
		return nil
	}
	return s.AgentResponseMetadata
}

func (s *SubscribeEvent) GetAudioResponse() *AudioResponse {
	if s == nil {
		return nil
	}
	if s.AudioResponse == nil {
		return nil
	}
	return s.AudioResponse
}

func (s *SubscribeEvent) GetInterruption() *Interruption {
	if s == nil {
		return nil
	}
	if s.Interruption == nil {
		return nil
	}
	return s.Interruption
}

func (s *SubscribeEvent) GetVadScore() *VadScore {
	if s == nil {
		return nil
	}
	if s.VadScore == nil {
		return nil
	}
	return s.VadScore
}

func (s *SubscribeEvent) GetAgentChatResponsePart() *AgentChatResponsePart {
	if s == nil {
		return nil
	}
	if s.AgentChatResponsePart == nil {
		return nil
	}
	return s.AgentChatResponsePart
}

func (s *SubscribeEvent) GetClientToolCall() *ClientToolCall {
	if s == nil {
		return nil
	}
	if s.ClientToolCall == nil {
		return nil
	}
	return s.ClientToolCall
}

func (s *SubscribeEvent) GetAgentToolResponse() *AgentToolResponse {
	if s == nil {
		return nil
	}
	if s.AgentToolResponse == nil {
		return nil
	}
	return s.AgentToolResponse
}

func (s *SubscribeEvent) GetAgentToolResponseFullPayload() *AgentToolResponseFullPayload {
	if s == nil {
		return nil
	}
	if s.AgentToolResponseFullPayload == nil {
		return nil
	}
	return s.AgentToolResponseFullPayload
}

func (s *SubscribeEvent) GetAgentToolRequest() *AgentToolRequest {
	if s == nil {
		return nil
	}
	if s.AgentToolRequest == nil {
		return nil
	}
	return s.AgentToolRequest
}

func (s *SubscribeEvent) GetConversationPing() *ConversationPing {
	if s == nil {
		return nil
	}
	if s.ConversationPing == nil {
		return nil
	}
	return s.ConversationPing
}

func (s *SubscribeEvent) GetMcpConnectionStatus() *McpConnectionStatus {
	if s == nil {
		return nil
	}
	if s.McpConnectionStatus == nil {
		return nil
	}
	return s.McpConnectionStatus
}

func (s *SubscribeEvent) UnmarshalJSON(data []byte) error {
	var unionKeys map[string]json.RawMessage
	_ = json.Unmarshal(data, &unionKeys)
	hasLiteral := func(key string, want interface{}) bool {
		raw, ok := unionKeys[key]
		if !ok {
			return false
		}
		var got interface{}
		return json.Unmarshal(raw, &got) == nil && got == want
	}
	matchedLiteral := false
	if hasLiteral("type", "conversation_initiation_metadata") {
		matchedLiteral = true
		var valueConversationInitiationMetadataLiteral *ConversationInitiationMetadata
		if err := json.Unmarshal(data, &valueConversationInitiationMetadataLiteral); err == nil {
			s.typ = "ConversationInitiationMetadata"
			s.ConversationInitiationMetadata = valueConversationInitiationMetadataLiteral
			return nil
		}
	}
	if hasLiteral("type", "agent_response_complete") {
		matchedLiteral = true
		var valueAgentResponseCompleteLiteral *AgentResponseComplete
		if err := json.Unmarshal(data, &valueAgentResponseCompleteLiteral); err == nil {
			s.typ = "AgentResponseComplete"
			s.AgentResponseComplete = valueAgentResponseCompleteLiteral
			return nil
		}
	}
	if hasLiteral("type", "context_usage") {
		matchedLiteral = true
		var valueContextUsageLiteral *ContextUsage
		if err := json.Unmarshal(data, &valueContextUsageLiteral); err == nil {
			s.typ = "ContextUsage"
			s.ContextUsage = valueContextUsageLiteral
			return nil
		}
	}
	if hasLiteral("type", "mcp_tool_call") {
		matchedLiteral = true
		var valueMcpToolCallLiteral *McpToolCall
		if err := json.Unmarshal(data, &valueMcpToolCallLiteral); err == nil {
			s.typ = "McpToolCall"
			s.McpToolCall = valueMcpToolCallLiteral
			return nil
		}
	}
	if hasLiteral("type", "client_error") {
		matchedLiteral = true
		var valueClientErrorLiteral *ClientError
		if err := json.Unmarshal(data, &valueClientErrorLiteral); err == nil {
			s.typ = "ClientError"
			s.ClientError = valueClientErrorLiteral
			return nil
		}
	}
	if hasLiteral("type", "guardrail_triggered") {
		matchedLiteral = true
		var valueGuardrailTriggeredLiteral *GuardrailTriggered
		if err := json.Unmarshal(data, &valueGuardrailTriggeredLiteral); err == nil {
			s.typ = "GuardrailTriggered"
			s.GuardrailTriggered = valueGuardrailTriggeredLiteral
			return nil
		}
	}
	if hasLiteral("type", "user_transcript") {
		matchedLiteral = true
		var valueConversationUserTranscriptLiteral *ConversationUserTranscript
		if err := json.Unmarshal(data, &valueConversationUserTranscriptLiteral); err == nil {
			s.typ = "ConversationUserTranscript"
			s.ConversationUserTranscript = valueConversationUserTranscriptLiteral
			return nil
		}
	}
	if hasLiteral("type", "agent_response") {
		matchedLiteral = true
		var valueConversationAgentResponseLiteral *ConversationAgentResponse
		if err := json.Unmarshal(data, &valueConversationAgentResponseLiteral); err == nil {
			s.typ = "ConversationAgentResponse"
			s.ConversationAgentResponse = valueConversationAgentResponseLiteral
			return nil
		}
	}
	if hasLiteral("type", "queue_status") {
		matchedLiteral = true
		var valueQueueStatusLiteral *QueueStatus
		if err := json.Unmarshal(data, &valueQueueStatusLiteral); err == nil {
			s.typ = "QueueStatus"
			s.QueueStatus = valueQueueStatusLiteral
			return nil
		}
	}
	if hasLiteral("type", "agent_response_correction") {
		matchedLiteral = true
		var valueAgentResponseCorrectionLiteral *AgentResponseCorrection
		if err := json.Unmarshal(data, &valueAgentResponseCorrectionLiteral); err == nil {
			s.typ = "AgentResponseCorrection"
			s.AgentResponseCorrection = valueAgentResponseCorrectionLiteral
			return nil
		}
	}
	if hasLiteral("type", "agent_response_metadata") {
		matchedLiteral = true
		var valueAgentResponseMetadataLiteral *AgentResponseMetadata
		if err := json.Unmarshal(data, &valueAgentResponseMetadataLiteral); err == nil {
			s.typ = "AgentResponseMetadata"
			s.AgentResponseMetadata = valueAgentResponseMetadataLiteral
			return nil
		}
	}
	if hasLiteral("type", "audio") {
		matchedLiteral = true
		var valueAudioResponseLiteral *AudioResponse
		if err := json.Unmarshal(data, &valueAudioResponseLiteral); err == nil {
			s.typ = "AudioResponse"
			s.AudioResponse = valueAudioResponseLiteral
			return nil
		}
	}
	if hasLiteral("type", "interruption") {
		matchedLiteral = true
		var valueInterruptionLiteral *Interruption
		if err := json.Unmarshal(data, &valueInterruptionLiteral); err == nil {
			s.typ = "Interruption"
			s.Interruption = valueInterruptionLiteral
			return nil
		}
	}
	if hasLiteral("type", "vad_score") {
		matchedLiteral = true
		var valueVadScoreLiteral *VadScore
		if err := json.Unmarshal(data, &valueVadScoreLiteral); err == nil {
			s.typ = "VadScore"
			s.VadScore = valueVadScoreLiteral
			return nil
		}
	}
	if hasLiteral("type", "agent_chat_response_part") {
		matchedLiteral = true
		var valueAgentChatResponsePartLiteral *AgentChatResponsePart
		if err := json.Unmarshal(data, &valueAgentChatResponsePartLiteral); err == nil {
			s.typ = "AgentChatResponsePart"
			s.AgentChatResponsePart = valueAgentChatResponsePartLiteral
			return nil
		}
	}
	if hasLiteral("type", "client_tool_call") {
		matchedLiteral = true
		var valueClientToolCallLiteral *ClientToolCall
		if err := json.Unmarshal(data, &valueClientToolCallLiteral); err == nil {
			s.typ = "ClientToolCall"
			s.ClientToolCall = valueClientToolCallLiteral
			return nil
		}
	}
	if hasLiteral("type", "agent_tool_response") {
		matchedLiteral = true
		var valueAgentToolResponseLiteral *AgentToolResponse
		if err := json.Unmarshal(data, &valueAgentToolResponseLiteral); err == nil {
			s.typ = "AgentToolResponse"
			s.AgentToolResponse = valueAgentToolResponseLiteral
			return nil
		}
	}
	if hasLiteral("type", "agent_tool_response_full_payload") {
		matchedLiteral = true
		var valueAgentToolResponseFullPayloadLiteral *AgentToolResponseFullPayload
		if err := json.Unmarshal(data, &valueAgentToolResponseFullPayloadLiteral); err == nil {
			s.typ = "AgentToolResponseFullPayload"
			s.AgentToolResponseFullPayload = valueAgentToolResponseFullPayloadLiteral
			return nil
		}
	}
	if hasLiteral("type", "agent_tool_request") {
		matchedLiteral = true
		var valueAgentToolRequestLiteral *AgentToolRequest
		if err := json.Unmarshal(data, &valueAgentToolRequestLiteral); err == nil {
			s.typ = "AgentToolRequest"
			s.AgentToolRequest = valueAgentToolRequestLiteral
			return nil
		}
	}
	if hasLiteral("type", "ping") {
		matchedLiteral = true
		var valueConversationPingLiteral *ConversationPing
		if err := json.Unmarshal(data, &valueConversationPingLiteral); err == nil {
			s.typ = "ConversationPing"
			s.ConversationPing = valueConversationPingLiteral
			return nil
		}
	}
	if hasLiteral("type", "mcp_connection_status") {
		matchedLiteral = true
		var valueMcpConnectionStatusLiteral *McpConnectionStatus
		if err := json.Unmarshal(data, &valueMcpConnectionStatusLiteral); err == nil {
			s.typ = "McpConnectionStatus"
			s.McpConnectionStatus = valueMcpConnectionStatusLiteral
			return nil
		}
	}
	if matchedLiteral {
		return fmt.Errorf("%s cannot be deserialized as a %T", data, s)
	}
	var valueConversationInitiationMetadata *ConversationInitiationMetadata
	if err := json.Unmarshal(data, &valueConversationInitiationMetadata); err == nil {
		s.typ = "ConversationInitiationMetadata"
		s.ConversationInitiationMetadata = valueConversationInitiationMetadata
		return nil
	}
	var valueAgentResponseComplete *AgentResponseComplete
	if err := json.Unmarshal(data, &valueAgentResponseComplete); err == nil {
		s.typ = "AgentResponseComplete"
		s.AgentResponseComplete = valueAgentResponseComplete
		return nil
	}
	var valueContextUsage *ContextUsage
	if err := json.Unmarshal(data, &valueContextUsage); err == nil {
		s.typ = "ContextUsage"
		s.ContextUsage = valueContextUsage
		return nil
	}
	var valueMcpToolCall *McpToolCall
	if err := json.Unmarshal(data, &valueMcpToolCall); err == nil {
		s.typ = "McpToolCall"
		s.McpToolCall = valueMcpToolCall
		return nil
	}
	var valueClientError *ClientError
	if err := json.Unmarshal(data, &valueClientError); err == nil {
		s.typ = "ClientError"
		s.ClientError = valueClientError
		return nil
	}
	var valueGuardrailTriggered *GuardrailTriggered
	if err := json.Unmarshal(data, &valueGuardrailTriggered); err == nil {
		s.typ = "GuardrailTriggered"
		s.GuardrailTriggered = valueGuardrailTriggered
		return nil
	}
	var valueConversationUserTranscript *ConversationUserTranscript
	if err := json.Unmarshal(data, &valueConversationUserTranscript); err == nil {
		s.typ = "ConversationUserTranscript"
		s.ConversationUserTranscript = valueConversationUserTranscript
		return nil
	}
	var valueConversationAgentResponse *ConversationAgentResponse
	if err := json.Unmarshal(data, &valueConversationAgentResponse); err == nil {
		s.typ = "ConversationAgentResponse"
		s.ConversationAgentResponse = valueConversationAgentResponse
		return nil
	}
	var valueQueueStatus *QueueStatus
	if err := json.Unmarshal(data, &valueQueueStatus); err == nil {
		s.typ = "QueueStatus"
		s.QueueStatus = valueQueueStatus
		return nil
	}
	var valueAgentResponseCorrection *AgentResponseCorrection
	if err := json.Unmarshal(data, &valueAgentResponseCorrection); err == nil {
		s.typ = "AgentResponseCorrection"
		s.AgentResponseCorrection = valueAgentResponseCorrection
		return nil
	}
	var valueAgentResponseMetadata *AgentResponseMetadata
	if err := json.Unmarshal(data, &valueAgentResponseMetadata); err == nil {
		s.typ = "AgentResponseMetadata"
		s.AgentResponseMetadata = valueAgentResponseMetadata
		return nil
	}
	var valueAudioResponse *AudioResponse
	if err := json.Unmarshal(data, &valueAudioResponse); err == nil {
		s.typ = "AudioResponse"
		s.AudioResponse = valueAudioResponse
		return nil
	}
	var valueInterruption *Interruption
	if err := json.Unmarshal(data, &valueInterruption); err == nil {
		s.typ = "Interruption"
		s.Interruption = valueInterruption
		return nil
	}
	var valueVadScore *VadScore
	if err := json.Unmarshal(data, &valueVadScore); err == nil {
		s.typ = "VadScore"
		s.VadScore = valueVadScore
		return nil
	}
	var valueAgentChatResponsePart *AgentChatResponsePart
	if err := json.Unmarshal(data, &valueAgentChatResponsePart); err == nil {
		s.typ = "AgentChatResponsePart"
		s.AgentChatResponsePart = valueAgentChatResponsePart
		return nil
	}
	var valueClientToolCall *ClientToolCall
	if err := json.Unmarshal(data, &valueClientToolCall); err == nil {
		s.typ = "ClientToolCall"
		s.ClientToolCall = valueClientToolCall
		return nil
	}
	var valueAgentToolResponse *AgentToolResponse
	if err := json.Unmarshal(data, &valueAgentToolResponse); err == nil {
		s.typ = "AgentToolResponse"
		s.AgentToolResponse = valueAgentToolResponse
		return nil
	}
	var valueAgentToolResponseFullPayload *AgentToolResponseFullPayload
	if err := json.Unmarshal(data, &valueAgentToolResponseFullPayload); err == nil {
		s.typ = "AgentToolResponseFullPayload"
		s.AgentToolResponseFullPayload = valueAgentToolResponseFullPayload
		return nil
	}
	var valueAgentToolRequest *AgentToolRequest
	if err := json.Unmarshal(data, &valueAgentToolRequest); err == nil {
		s.typ = "AgentToolRequest"
		s.AgentToolRequest = valueAgentToolRequest
		return nil
	}
	var valueConversationPing *ConversationPing
	if err := json.Unmarshal(data, &valueConversationPing); err == nil {
		s.typ = "ConversationPing"
		s.ConversationPing = valueConversationPing
		return nil
	}
	var valueMcpConnectionStatus *McpConnectionStatus
	if err := json.Unmarshal(data, &valueMcpConnectionStatus); err == nil {
		s.typ = "McpConnectionStatus"
		s.McpConnectionStatus = valueMcpConnectionStatus
		return nil
	}
	return fmt.Errorf("%s cannot be deserialized as a %T", data, s)
}

func (s SubscribeEvent) MarshalJSON() ([]byte, error) {
	if s.typ == "ConversationInitiationMetadata" || s.ConversationInitiationMetadata != nil {
		return json.Marshal(s.ConversationInitiationMetadata)
	}
	if s.typ == "AgentResponseComplete" || s.AgentResponseComplete != nil {
		return json.Marshal(s.AgentResponseComplete)
	}
	if s.typ == "ContextUsage" || s.ContextUsage != nil {
		return json.Marshal(s.ContextUsage)
	}
	if s.typ == "McpToolCall" || s.McpToolCall != nil {
		return json.Marshal(s.McpToolCall)
	}
	if s.typ == "ClientError" || s.ClientError != nil {
		return json.Marshal(s.ClientError)
	}
	if s.typ == "GuardrailTriggered" || s.GuardrailTriggered != nil {
		return json.Marshal(s.GuardrailTriggered)
	}
	if s.typ == "ConversationUserTranscript" || s.ConversationUserTranscript != nil {
		return json.Marshal(s.ConversationUserTranscript)
	}
	if s.typ == "ConversationAgentResponse" || s.ConversationAgentResponse != nil {
		return json.Marshal(s.ConversationAgentResponse)
	}
	if s.typ == "QueueStatus" || s.QueueStatus != nil {
		return json.Marshal(s.QueueStatus)
	}
	if s.typ == "AgentResponseCorrection" || s.AgentResponseCorrection != nil {
		return json.Marshal(s.AgentResponseCorrection)
	}
	if s.typ == "AgentResponseMetadata" || s.AgentResponseMetadata != nil {
		return json.Marshal(s.AgentResponseMetadata)
	}
	if s.typ == "AudioResponse" || s.AudioResponse != nil {
		return json.Marshal(s.AudioResponse)
	}
	if s.typ == "Interruption" || s.Interruption != nil {
		return json.Marshal(s.Interruption)
	}
	if s.typ == "VadScore" || s.VadScore != nil {
		return json.Marshal(s.VadScore)
	}
	if s.typ == "AgentChatResponsePart" || s.AgentChatResponsePart != nil {
		return json.Marshal(s.AgentChatResponsePart)
	}
	if s.typ == "ClientToolCall" || s.ClientToolCall != nil {
		return json.Marshal(s.ClientToolCall)
	}
	if s.typ == "AgentToolResponse" || s.AgentToolResponse != nil {
		return json.Marshal(s.AgentToolResponse)
	}
	if s.typ == "AgentToolResponseFullPayload" || s.AgentToolResponseFullPayload != nil {
		return json.Marshal(s.AgentToolResponseFullPayload)
	}
	if s.typ == "AgentToolRequest" || s.AgentToolRequest != nil {
		return json.Marshal(s.AgentToolRequest)
	}
	if s.typ == "ConversationPing" || s.ConversationPing != nil {
		return json.Marshal(s.ConversationPing)
	}
	if s.typ == "McpConnectionStatus" || s.McpConnectionStatus != nil {
		return json.Marshal(s.McpConnectionStatus)
	}
	return nil, fmt.Errorf("type %T does not include a non-empty union type", s)
}

type SubscribeEventVisitor interface {
	VisitConversationInitiationMetadata(*ConversationInitiationMetadata) error
	VisitAgentResponseComplete(*AgentResponseComplete) error
	VisitContextUsage(*ContextUsage) error
	VisitMcpToolCall(*McpToolCall) error
	VisitClientError(*ClientError) error
	VisitGuardrailTriggered(*GuardrailTriggered) error
	VisitConversationUserTranscript(*ConversationUserTranscript) error
	VisitConversationAgentResponse(*ConversationAgentResponse) error
	VisitQueueStatus(*QueueStatus) error
	VisitAgentResponseCorrection(*AgentResponseCorrection) error
	VisitAgentResponseMetadata(*AgentResponseMetadata) error
	VisitAudioResponse(*AudioResponse) error
	VisitInterruption(*Interruption) error
	VisitVadScore(*VadScore) error
	VisitAgentChatResponsePart(*AgentChatResponsePart) error
	VisitClientToolCall(*ClientToolCall) error
	VisitAgentToolResponse(*AgentToolResponse) error
	VisitAgentToolResponseFullPayload(*AgentToolResponseFullPayload) error
	VisitAgentToolRequest(*AgentToolRequest) error
	VisitConversationPing(*ConversationPing) error
	VisitMcpConnectionStatus(*McpConnectionStatus) error
}

func (s *SubscribeEvent) Accept(visitor SubscribeEventVisitor) error {
	if s.typ == "ConversationInitiationMetadata" || s.ConversationInitiationMetadata != nil {
		return visitor.VisitConversationInitiationMetadata(s.ConversationInitiationMetadata)
	}
	if s.typ == "AgentResponseComplete" || s.AgentResponseComplete != nil {
		return visitor.VisitAgentResponseComplete(s.AgentResponseComplete)
	}
	if s.typ == "ContextUsage" || s.ContextUsage != nil {
		return visitor.VisitContextUsage(s.ContextUsage)
	}
	if s.typ == "McpToolCall" || s.McpToolCall != nil {
		return visitor.VisitMcpToolCall(s.McpToolCall)
	}
	if s.typ == "ClientError" || s.ClientError != nil {
		return visitor.VisitClientError(s.ClientError)
	}
	if s.typ == "GuardrailTriggered" || s.GuardrailTriggered != nil {
		return visitor.VisitGuardrailTriggered(s.GuardrailTriggered)
	}
	if s.typ == "ConversationUserTranscript" || s.ConversationUserTranscript != nil {
		return visitor.VisitConversationUserTranscript(s.ConversationUserTranscript)
	}
	if s.typ == "ConversationAgentResponse" || s.ConversationAgentResponse != nil {
		return visitor.VisitConversationAgentResponse(s.ConversationAgentResponse)
	}
	if s.typ == "QueueStatus" || s.QueueStatus != nil {
		return visitor.VisitQueueStatus(s.QueueStatus)
	}
	if s.typ == "AgentResponseCorrection" || s.AgentResponseCorrection != nil {
		return visitor.VisitAgentResponseCorrection(s.AgentResponseCorrection)
	}
	if s.typ == "AgentResponseMetadata" || s.AgentResponseMetadata != nil {
		return visitor.VisitAgentResponseMetadata(s.AgentResponseMetadata)
	}
	if s.typ == "AudioResponse" || s.AudioResponse != nil {
		return visitor.VisitAudioResponse(s.AudioResponse)
	}
	if s.typ == "Interruption" || s.Interruption != nil {
		return visitor.VisitInterruption(s.Interruption)
	}
	if s.typ == "VadScore" || s.VadScore != nil {
		return visitor.VisitVadScore(s.VadScore)
	}
	if s.typ == "AgentChatResponsePart" || s.AgentChatResponsePart != nil {
		return visitor.VisitAgentChatResponsePart(s.AgentChatResponsePart)
	}
	if s.typ == "ClientToolCall" || s.ClientToolCall != nil {
		return visitor.VisitClientToolCall(s.ClientToolCall)
	}
	if s.typ == "AgentToolResponse" || s.AgentToolResponse != nil {
		return visitor.VisitAgentToolResponse(s.AgentToolResponse)
	}
	if s.typ == "AgentToolResponseFullPayload" || s.AgentToolResponseFullPayload != nil {
		return visitor.VisitAgentToolResponseFullPayload(s.AgentToolResponseFullPayload)
	}
	if s.typ == "AgentToolRequest" || s.AgentToolRequest != nil {
		return visitor.VisitAgentToolRequest(s.AgentToolRequest)
	}
	if s.typ == "ConversationPing" || s.ConversationPing != nil {
		return visitor.VisitConversationPing(s.ConversationPing)
	}
	if s.typ == "McpConnectionStatus" || s.McpConnectionStatus != nil {
		return visitor.VisitMcpConnectionStatus(s.McpConnectionStatus)
	}
	return fmt.Errorf("type %T does not include a non-empty union type", s)
}
