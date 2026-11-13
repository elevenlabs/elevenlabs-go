package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGptImage25SunburstRequestAspectRatioVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant GptImage25SunburstRequestAspectRatio
		wire    string
	}
	cases := []enumCase{
		{variant: GptImage25SunburstRequestAspectRatioAuto, wire: "\"auto\""},
		{variant: GptImage25SunburstRequestAspectRatioOne1, wire: "\"1:1\""},
		{variant: GptImage25SunburstRequestAspectRatioFour5, wire: "\"4:5\""},
		{variant: GptImage25SunburstRequestAspectRatioFive4, wire: "\"5:4\""},
		{variant: GptImage25SunburstRequestAspectRatioThree4, wire: "\"3:4\""},
		{variant: GptImage25SunburstRequestAspectRatioFour3, wire: "\"4:3\""},
		{variant: GptImage25SunburstRequestAspectRatioTwo3, wire: "\"2:3\""},
		{variant: GptImage25SunburstRequestAspectRatioThree2, wire: "\"3:2\""},
		{variant: GptImage25SunburstRequestAspectRatioOne2, wire: "\"1:2\""},
		{variant: GptImage25SunburstRequestAspectRatioTwo1, wire: "\"2:1\""},
		{variant: GptImage25SunburstRequestAspectRatioNine16, wire: "\"9:16\""},
		{variant: GptImage25SunburstRequestAspectRatioOne69, wire: "\"16:9\""},
		{variant: GptImage25SunburstRequestAspectRatioTwo19, wire: "\"21:9\""},
		{variant: GptImage25SunburstRequestAspectRatioOne3, wire: "\"1:3\""},
		{variant: GptImage25SunburstRequestAspectRatioThree1, wire: "\"3:1\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded GptImage25SunburstRequestAspectRatio
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
