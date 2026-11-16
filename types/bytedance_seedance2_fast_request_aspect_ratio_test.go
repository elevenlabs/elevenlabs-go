package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBytedanceSeedance2FastRequestAspectRatioVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BytedanceSeedance2FastRequestAspectRatio
		wire    string
	}
	cases := []enumCase{
		{variant: BytedanceSeedance2FastRequestAspectRatioAuto, wire: "\"auto\""},
		{variant: BytedanceSeedance2FastRequestAspectRatioTwo19, wire: "\"21:9\""},
		{variant: BytedanceSeedance2FastRequestAspectRatioOne69, wire: "\"16:9\""},
		{variant: BytedanceSeedance2FastRequestAspectRatioFour3, wire: "\"4:3\""},
		{variant: BytedanceSeedance2FastRequestAspectRatioOne1, wire: "\"1:1\""},
		{variant: BytedanceSeedance2FastRequestAspectRatioThree4, wire: "\"3:4\""},
		{variant: BytedanceSeedance2FastRequestAspectRatioNine16, wire: "\"9:16\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BytedanceSeedance2FastRequestAspectRatio
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
