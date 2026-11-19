package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestResponseConversationErrorTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ResponseConversationErrorType
		wire    string
	}
	cases := []enumCase{
		{variant: ResponseConversationErrorTypeSystemError, wire: "\"system_error\""},
		{variant: ResponseConversationErrorTypeCallInitializationError, wire: "\"call_initialization_error\""},
		{variant: ResponseConversationErrorTypeLineBusy, wire: "\"line_busy\""},
		{variant: ResponseConversationErrorTypeNoAnswer, wire: "\"no_answer\""},
		{variant: ResponseConversationErrorTypeCallRejected, wire: "\"call_rejected\""},
		{variant: ResponseConversationErrorTypeBlockedByUser, wire: "\"blocked_by_user\""},
		{variant: ResponseConversationErrorTypeAgentConfigurationError, wire: "\"agent_configuration_error\""},
		{variant: ResponseConversationErrorTypeInvalidClientRequest, wire: "\"invalid_client_request\""},
		{variant: ResponseConversationErrorTypePermissionError, wire: "\"permission_error\""},
		{variant: ResponseConversationErrorTypeEntitlementExceeded, wire: "\"entitlement_exceeded\""},
		{variant: ResponseConversationErrorTypeClientDisconnected, wire: "\"client_disconnected\""},
		{variant: ResponseConversationErrorTypeLlmError, wire: "\"llm_error\""},
		{variant: ResponseConversationErrorTypeSpeechError, wire: "\"speech_error\""},
		{variant: ResponseConversationErrorTypeToolError, wire: "\"tool_error\""},
		{variant: ResponseConversationErrorTypeIntegrationError, wire: "\"integration_error\""},
		{variant: ResponseConversationErrorTypeGuardrailTriggered, wire: "\"guardrail_triggered\""},
		{variant: ResponseConversationErrorTypeSafetyViolation, wire: "\"safety_violation\""},
		{variant: ResponseConversationErrorTypeMaxDurationExceeded, wire: "\"max_duration_exceeded\""},
		{variant: ResponseConversationErrorTypePostProcessingError, wire: "\"post_processing_error\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ResponseConversationErrorType
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
