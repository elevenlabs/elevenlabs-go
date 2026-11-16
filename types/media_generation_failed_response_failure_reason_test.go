package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMediaGenerationFailedResponseFailureReasonVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant MediaGenerationFailedResponseFailureReason
		wire    string
	}
	cases := []enumCase{
		{variant: MediaGenerationFailedResponseFailureReasonTimeout, wire: "\"timeout\""},
		{variant: MediaGenerationFailedResponseFailureReasonModelError, wire: "\"model_error\""},
		{variant: MediaGenerationFailedResponseFailureReasonModerated, wire: "\"moderated\""},
		{variant: MediaGenerationFailedResponseFailureReasonInvalidParameters, wire: "\"invalid_parameters\""},
		{variant: MediaGenerationFailedResponseFailureReasonDependencyFailed, wire: "\"dependency_failed\""},
		{variant: MediaGenerationFailedResponseFailureReasonChargingFailed, wire: "\"charging_failed\""},
		{variant: MediaGenerationFailedResponseFailureReasonInternalError, wire: "\"internal_error\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded MediaGenerationFailedResponseFailureReason
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
