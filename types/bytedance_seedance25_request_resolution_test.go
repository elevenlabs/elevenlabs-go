package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBytedanceSeedance25RequestResolutionVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BytedanceSeedance25RequestResolution
		wire    string
	}
	cases := []enumCase{
		{variant: BytedanceSeedance25RequestResolutionFour80P, wire: "\"480p\""},
		{variant: BytedanceSeedance25RequestResolutionSeven20P, wire: "\"720p\""},
		{variant: BytedanceSeedance25RequestResolutionOne080P, wire: "\"1080p\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BytedanceSeedance25RequestResolution
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
