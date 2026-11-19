package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGptImage25FlareRequestAspectRatioVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant GptImage25FlareRequestAspectRatio
		wire    string
	}
	cases := []enumCase{
		{variant: GptImage25FlareRequestAspectRatioAuto, wire: "\"auto\""},
		{variant: GptImage25FlareRequestAspectRatioOne1, wire: "\"1:1\""},
		{variant: GptImage25FlareRequestAspectRatioFour5, wire: "\"4:5\""},
		{variant: GptImage25FlareRequestAspectRatioFive4, wire: "\"5:4\""},
		{variant: GptImage25FlareRequestAspectRatioThree4, wire: "\"3:4\""},
		{variant: GptImage25FlareRequestAspectRatioFour3, wire: "\"4:3\""},
		{variant: GptImage25FlareRequestAspectRatioTwo3, wire: "\"2:3\""},
		{variant: GptImage25FlareRequestAspectRatioThree2, wire: "\"3:2\""},
		{variant: GptImage25FlareRequestAspectRatioOne2, wire: "\"1:2\""},
		{variant: GptImage25FlareRequestAspectRatioTwo1, wire: "\"2:1\""},
		{variant: GptImage25FlareRequestAspectRatioNine16, wire: "\"9:16\""},
		{variant: GptImage25FlareRequestAspectRatioOne69, wire: "\"16:9\""},
		{variant: GptImage25FlareRequestAspectRatioTwo19, wire: "\"21:9\""},
		{variant: GptImage25FlareRequestAspectRatioOne3, wire: "\"1:3\""},
		{variant: GptImage25FlareRequestAspectRatioThree1, wire: "\"3:1\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded GptImage25FlareRequestAspectRatio
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
