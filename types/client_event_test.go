package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestClientEventVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ClientEvent
		wire    string
	}
	cases := []enumCase{
		{variant: ClientEventConversationInitiationMetadata, wire: "\"conversation_initiation_metadata\""},
		{variant: ClientEventAsrInitiationMetadata, wire: "\"asr_initiation_metadata\""},
		{variant: ClientEventPing, wire: "\"ping\""},
		{variant: ClientEventAudio, wire: "\"audio\""},
		{variant: ClientEventInterruption, wire: "\"interruption\""},
		{variant: ClientEventUserTranscript, wire: "\"user_transcript\""},
		{variant: ClientEventTentativeUserTranscript, wire: "\"tentative_user_transcript\""},
		{variant: ClientEventAgentResponse, wire: "\"agent_response\""},
		{variant: ClientEventAgentResponseCorrection, wire: "\"agent_response_correction\""},
		{variant: ClientEventClientToolCall, wire: "\"client_tool_call\""},
		{variant: ClientEventMcpToolCall, wire: "\"mcp_tool_call\""},
		{variant: ClientEventMcpConnectionStatus, wire: "\"mcp_connection_status\""},
		{variant: ClientEventAgentToolRequest, wire: "\"agent_tool_request\""},
		{variant: ClientEventAgentToolResponse, wire: "\"agent_tool_response\""},
		{variant: ClientEventAgentToolResponseFullPayload, wire: "\"agent_tool_response_full_payload\""},
		{variant: ClientEventAgentResponseMetadata, wire: "\"agent_response_metadata\""},
		{variant: ClientEventVadScore, wire: "\"vad_score\""},
		{variant: ClientEventAgentChatResponsePart, wire: "\"agent_chat_response_part\""},
		{variant: ClientEventClientError, wire: "\"client_error\""},
		{variant: ClientEventGuardrailTriggered, wire: "\"guardrail_triggered\""},
		{variant: ClientEventDtmfRequest, wire: "\"dtmf_request\""},
		{variant: ClientEventAgentResponseComplete, wire: "\"agent_response_complete\""},
		{variant: ClientEventContextUsage, wire: "\"context_usage\""},
		{variant: ClientEventInternalTurnProbability, wire: "\"internal_turn_probability\""},
		{variant: ClientEventInternalTentativeAgentResponse, wire: "\"internal_tentative_agent_response\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ClientEvent
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}

func TestClientEventVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ClientEvent
		wire    string
	}
	cases := []enumCase{
		{variant: ClientEventConversationInitiationMetadata, wire: "\"conversation_initiation_metadata\""},
		{variant: ClientEventAsrInitiationMetadata, wire: "\"asr_initiation_metadata\""},
		{variant: ClientEventPing, wire: "\"ping\""},
		{variant: ClientEventAudio, wire: "\"audio\""},
		{variant: ClientEventInterruption, wire: "\"interruption\""},
		{variant: ClientEventUserTranscript, wire: "\"user_transcript\""},
		{variant: ClientEventTentativeUserTranscript, wire: "\"tentative_user_transcript\""},
		{variant: ClientEventAgentResponse, wire: "\"agent_response\""},
		{variant: ClientEventAgentResponseCorrection, wire: "\"agent_response_correction\""},
		{variant: ClientEventClientToolCall, wire: "\"client_tool_call\""},
		{variant: ClientEventMcpToolCall, wire: "\"mcp_tool_call\""},
		{variant: ClientEventMcpConnectionStatus, wire: "\"mcp_connection_status\""},
		{variant: ClientEventAgentToolRequest, wire: "\"agent_tool_request\""},
		{variant: ClientEventAgentToolResponse, wire: "\"agent_tool_response\""},
		{variant: ClientEventAgentToolResponseFullPayload, wire: "\"agent_tool_response_full_payload\""},
		{variant: ClientEventAgentResponseMetadata, wire: "\"agent_response_metadata\""},
		{variant: ClientEventVadScore, wire: "\"vad_score\""},
		{variant: ClientEventAgentChatResponsePart, wire: "\"agent_chat_response_part\""},
		{variant: ClientEventClientError, wire: "\"client_error\""},
		{variant: ClientEventGuardrailTriggered, wire: "\"guardrail_triggered\""},
		{variant: ClientEventDtmfRequest, wire: "\"dtmf_request\""},
		{variant: ClientEventAgentResponseComplete, wire: "\"agent_response_complete\""},
		{variant: ClientEventContextUsage, wire: "\"context_usage\""},
		{variant: ClientEventInternalTurnProbability, wire: "\"internal_turn_probability\""},
		{variant: ClientEventInternalTentativeAgentResponse, wire: "\"internal_tentative_agent_response\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ClientEvent
			if err := json.Unmarshal([]byte(tc.wire), &decoded); err != nil {
				t.Fatalf("unmarshal: %v", err)
			}
			if decoded != tc.variant {
				t.Fatalf("decoded mismatch: got %#v want %#v", decoded, tc.variant)
			}
			var want any
			var got any
			if err := json.Unmarshal([]byte(tc.wire), &want); err != nil {
				t.Fatalf("unmarshal want: %v", err)
			}
			if err := json.Unmarshal(encoded, &got); err != nil {
				t.Fatalf("unmarshal got: %v", err)
			}
			if !reflect.DeepEqual(got, want) {
				t.Fatalf("round-trip mismatch\n got: %#v\nwant: %#v", got, want)
			}
		})
	}
}
