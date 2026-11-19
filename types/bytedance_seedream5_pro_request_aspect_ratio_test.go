package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBytedanceSeedream5ProRequestAspectRatioVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BytedanceSeedream5ProRequestAspectRatio
		wire    string
	}
	cases := []enumCase{
		{variant: BytedanceSeedream5ProRequestAspectRatioAuto, wire: "\"auto\""},
		{variant: BytedanceSeedream5ProRequestAspectRatioOne1, wire: "\"1:1\""},
		{variant: BytedanceSeedream5ProRequestAspectRatioThree4, wire: "\"3:4\""},
		{variant: BytedanceSeedream5ProRequestAspectRatioOne69, wire: "\"16:9\""},
		{variant: BytedanceSeedream5ProRequestAspectRatioFour3, wire: "\"4:3\""},
		{variant: BytedanceSeedream5ProRequestAspectRatioNine16, wire: "\"9:16\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BytedanceSeedream5ProRequestAspectRatio
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
