package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBytedanceSeedance2RequestAspectRatioVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BytedanceSeedance2RequestAspectRatio
		wire    string
	}
	cases := []enumCase{
		{variant: BytedanceSeedance2RequestAspectRatioAuto, wire: "\"auto\""},
		{variant: BytedanceSeedance2RequestAspectRatioTwo19, wire: "\"21:9\""},
		{variant: BytedanceSeedance2RequestAspectRatioOne69, wire: "\"16:9\""},
		{variant: BytedanceSeedance2RequestAspectRatioFour3, wire: "\"4:3\""},
		{variant: BytedanceSeedance2RequestAspectRatioOne1, wire: "\"1:1\""},
		{variant: BytedanceSeedance2RequestAspectRatioThree4, wire: "\"3:4\""},
		{variant: BytedanceSeedance2RequestAspectRatioNine16, wire: "\"9:16\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BytedanceSeedance2RequestAspectRatio
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
