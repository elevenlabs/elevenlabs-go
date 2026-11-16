package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDecisionDispositionVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant DecisionDisposition
		wire    string
	}
	cases := []enumCase{
		{variant: DecisionDispositionCompleted, wire: "\"completed\""},
		{variant: DecisionDispositionRetryQueued, wire: "\"retry_queued\""},
		{variant: DecisionDispositionCallbackQueued, wire: "\"callback_queued\""},
		{variant: DecisionDispositionRedirectQueued, wire: "\"redirect_queued\""},
		{variant: DecisionDispositionAttention, wire: "\"attention\""},
		{variant: DecisionDispositionCancelled, wire: "\"cancelled\""},
		{variant: DecisionDispositionDoNotCall, wire: "\"do_not_call\""},
		{variant: DecisionDispositionSuppressed, wire: "\"suppressed\""},
		{variant: DecisionDispositionWindowSuspended, wire: "\"window_suspended\""},
		{variant: DecisionDispositionAttemptCeilingReached, wire: "\"attempt_ceiling_reached\""},
		{variant: DecisionDispositionCallbackCeilingReached, wire: "\"callback_ceiling_reached\""},
		{variant: DecisionDispositionRedirectCeilingReached, wire: "\"redirect_ceiling_reached\""},
		{variant: DecisionDispositionCallbackTimeInvalid, wire: "\"callback_time_invalid\""},
		{variant: DecisionDispositionRedirectNumberInvalid, wire: "\"redirect_number_invalid\""},
		{variant: DecisionDispositionCouldntProcess, wire: "\"couldnt_process\""},
		{variant: DecisionDispositionSuppressedDoNotCall, wire: "\"suppressed_do_not_call\""},
		{variant: DecisionDispositionSuppressedAfterStop, wire: "\"suppressed_after_stop\""},
		{variant: DecisionDispositionPermissionNotGranted, wire: "\"permission_not_granted\""},
		{variant: DecisionDispositionPermissionDeclined, wire: "\"permission_declined\""},
		{variant: DecisionDispositionPermissionRequestBlocked, wire: "\"permission_request_blocked\""},
		{variant: DecisionDispositionPermissionExpired, wire: "\"permission_expired\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded DecisionDisposition
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
