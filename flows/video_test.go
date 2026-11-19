package flows

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestListVideoRequestStatusVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ListVideoRequestStatus
		wire    string
	}
	cases := []enumCase{
		{variant: ListVideoRequestStatusPending, wire: "\"pending\""},
		{variant: ListVideoRequestStatusGenerating, wire: "\"generating\""},
		{variant: ListVideoRequestStatusCompleted, wire: "\"completed\""},
		{variant: ListVideoRequestStatusFailed, wire: "\"failed\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ListVideoRequestStatus
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
