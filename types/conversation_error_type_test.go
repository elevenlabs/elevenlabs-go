package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestConversationErrorTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ConversationErrorType
		wire    string
	}
	cases := []enumCase{
		{variant: ConversationErrorTypeUnknown, wire: "\"unknown\""},
		{variant: ConversationErrorTypeCallInitializationError, wire: "\"call_initialization_error\""},
		{variant: ConversationErrorTypeLineBusy, wire: "\"line_busy\""},
		{variant: ConversationErrorTypeNoAnswer, wire: "\"no_answer\""},
		{variant: ConversationErrorTypeCallRejected, wire: "\"call_rejected\""},
		{variant: ConversationErrorTypeBlockedByUser, wire: "\"blocked_by_user\""},
		{variant: ConversationErrorTypeAgentConfigurationError, wire: "\"agent_configuration_error\""},
		{variant: ConversationErrorTypeInvalidClientRequest, wire: "\"invalid_client_request\""},
		{variant: ConversationErrorTypePermissionError, wire: "\"permission_error\""},
		{variant: ConversationErrorTypeEntitlementExceeded, wire: "\"entitlement_exceeded\""},
		{variant: ConversationErrorTypeClientDisconnected, wire: "\"client_disconnected\""},
		{variant: ConversationErrorTypeLlmError, wire: "\"llm_error\""},
		{variant: ConversationErrorTypeSpeechError, wire: "\"speech_error\""},
		{variant: ConversationErrorTypeToolError, wire: "\"tool_error\""},
		{variant: ConversationErrorTypeIntegrationError, wire: "\"integration_error\""},
		{variant: ConversationErrorTypeDependencyError, wire: "\"dependency_error\""},
		{variant: ConversationErrorTypeGuardrailTriggered, wire: "\"guardrail_triggered\""},
		{variant: ConversationErrorTypeSafetyViolation, wire: "\"safety_violation\""},
		{variant: ConversationErrorTypeMaxDurationExceeded, wire: "\"max_duration_exceeded\""},
		{variant: ConversationErrorTypeWorkflowEndFailure, wire: "\"workflow_end_failure\""},
		{variant: ConversationErrorTypeInternalError, wire: "\"internal_error\""},
		{variant: ConversationErrorTypePostProcessingError, wire: "\"post_processing_error\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ConversationErrorType
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
