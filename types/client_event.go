package types

import (
	"fmt"
)

type ClientEvent string

const (
	ClientEventConversationInitiationMetadata ClientEvent = "conversation_initiation_metadata"
	ClientEventAsrInitiationMetadata          ClientEvent = "asr_initiation_metadata"
	ClientEventPing                           ClientEvent = "ping"
	ClientEventAudio                          ClientEvent = "audio"
	ClientEventInterruption                   ClientEvent = "interruption"
	ClientEventUserTranscript                 ClientEvent = "user_transcript"
	ClientEventTentativeUserTranscript        ClientEvent = "tentative_user_transcript"
	ClientEventAgentResponse                  ClientEvent = "agent_response"
	ClientEventAgentResponseCorrection        ClientEvent = "agent_response_correction"
	ClientEventClientToolCall                 ClientEvent = "client_tool_call"
	ClientEventMcpToolCall                    ClientEvent = "mcp_tool_call"
	ClientEventMcpConnectionStatus            ClientEvent = "mcp_connection_status"
	ClientEventAgentToolRequest               ClientEvent = "agent_tool_request"
	ClientEventAgentToolResponse              ClientEvent = "agent_tool_response"
	ClientEventAgentToolResponseFullPayload   ClientEvent = "agent_tool_response_full_payload"
	ClientEventAgentResponseMetadata          ClientEvent = "agent_response_metadata"
	ClientEventVadScore                       ClientEvent = "vad_score"
	ClientEventAgentChatResponsePart          ClientEvent = "agent_chat_response_part"
	ClientEventClientError                    ClientEvent = "client_error"
	ClientEventGuardrailTriggered             ClientEvent = "guardrail_triggered"
	ClientEventDtmfRequest                    ClientEvent = "dtmf_request"
	ClientEventAgentResponseComplete          ClientEvent = "agent_response_complete"
	ClientEventContextUsage                   ClientEvent = "context_usage"
	ClientEventInternalTurnProbability        ClientEvent = "internal_turn_probability"
	ClientEventInternalTentativeAgentResponse ClientEvent = "internal_tentative_agent_response"
)

func NewClientEventFromString(s string) (ClientEvent, error) {
	switch s {
	case "conversation_initiation_metadata":
		return ClientEventConversationInitiationMetadata, nil
	case "asr_initiation_metadata":
		return ClientEventAsrInitiationMetadata, nil
	case "ping":
		return ClientEventPing, nil
	case "audio":
		return ClientEventAudio, nil
	case "interruption":
		return ClientEventInterruption, nil
	case "user_transcript":
		return ClientEventUserTranscript, nil
	case "tentative_user_transcript":
		return ClientEventTentativeUserTranscript, nil
	case "agent_response":
		return ClientEventAgentResponse, nil
	case "agent_response_correction":
		return ClientEventAgentResponseCorrection, nil
	case "client_tool_call":
		return ClientEventClientToolCall, nil
	case "mcp_tool_call":
		return ClientEventMcpToolCall, nil
	case "mcp_connection_status":
		return ClientEventMcpConnectionStatus, nil
	case "agent_tool_request":
		return ClientEventAgentToolRequest, nil
	case "agent_tool_response":
		return ClientEventAgentToolResponse, nil
	case "agent_tool_response_full_payload":
		return ClientEventAgentToolResponseFullPayload, nil
	case "agent_response_metadata":
		return ClientEventAgentResponseMetadata, nil
	case "vad_score":
		return ClientEventVadScore, nil
	case "agent_chat_response_part":
		return ClientEventAgentChatResponsePart, nil
	case "client_error":
		return ClientEventClientError, nil
	case "guardrail_triggered":
		return ClientEventGuardrailTriggered, nil
	case "dtmf_request":
		return ClientEventDtmfRequest, nil
	case "agent_response_complete":
		return ClientEventAgentResponseComplete, nil
	case "context_usage":
		return ClientEventContextUsage, nil
	case "internal_turn_probability":
		return ClientEventInternalTurnProbability, nil
	case "internal_tentative_agent_response":
		return ClientEventInternalTentativeAgentResponse, nil
	}
	var t ClientEvent
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e ClientEvent) Ptr() *ClientEvent {
	return &e
}
