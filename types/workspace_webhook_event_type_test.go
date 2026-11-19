package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestWorkspaceWebhookEventTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant WorkspaceWebhookEventType
		wire    string
	}
	cases := []enumCase{
		{variant: WorkspaceWebhookEventTypeVoiceLibraryRemovalNotice, wire: "\"voice_library_removal_notice\""},
		{variant: WorkspaceWebhookEventTypeSpeechToText, wire: "\"speech_to_text\""},
		{variant: WorkspaceWebhookEventTypeAgentQa, wire: "\"agent_qa\""},
		{variant: WorkspaceWebhookEventTypeFlows, wire: "\"flows\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded WorkspaceWebhookEventType
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
