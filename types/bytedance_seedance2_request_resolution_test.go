package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBytedanceSeedance2RequestResolutionVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BytedanceSeedance2RequestResolution
		wire    string
	}
	cases := []enumCase{
		{variant: BytedanceSeedance2RequestResolutionFour80P, wire: "\"480p\""},
		{variant: BytedanceSeedance2RequestResolutionSeven20P, wire: "\"720p\""},
		{variant: BytedanceSeedance2RequestResolutionOne080P, wire: "\"1080p\""},
		{variant: BytedanceSeedance2RequestResolutionFourk, wire: "\"4k\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BytedanceSeedance2RequestResolution
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
