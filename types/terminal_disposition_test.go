package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestTerminalDispositionVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant TerminalDisposition
		wire    string
	}
	cases := []enumCase{
		{variant: TerminalDispositionCampaignEnded, wire: "\"campaign_ended\""},
		{variant: TerminalDispositionRemoved, wire: "\"removed\""},
		{variant: TerminalDispositionDoNotCall, wire: "\"do_not_call\""},
		{variant: TerminalDispositionCompleted, wire: "\"completed\""},
		{variant: TerminalDispositionCancelled, wire: "\"cancelled\""},
		{variant: TerminalDispositionSuppressed, wire: "\"suppressed\""},
		{variant: TerminalDispositionAttemptCeilingReached, wire: "\"attempt_ceiling_reached\""},
		{variant: TerminalDispositionCallbackCeilingReached, wire: "\"callback_ceiling_reached\""},
		{variant: TerminalDispositionRedirectCeilingReached, wire: "\"redirect_ceiling_reached\""},
		{variant: TerminalDispositionPermissionNotGranted, wire: "\"permission_not_granted\""},
		{variant: TerminalDispositionPermissionDeclined, wire: "\"permission_declined\""},
		{variant: TerminalDispositionPermissionRequestBlocked, wire: "\"permission_request_blocked\""},
		{variant: TerminalDispositionPermissionExpired, wire: "\"permission_expired\""},
		{variant: TerminalDispositionCallbackTimeInvalid, wire: "\"callback_time_invalid\""},
		{variant: TerminalDispositionRedirectNumberInvalid, wire: "\"redirect_number_invalid\""},
		{variant: TerminalDispositionCouldntProcess, wire: "\"couldnt_process\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded TerminalDisposition
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
