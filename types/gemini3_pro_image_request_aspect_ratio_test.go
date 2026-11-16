package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGemini3ProImageRequestAspectRatioVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Gemini3ProImageRequestAspectRatio
		wire    string
	}
	cases := []enumCase{
		{variant: Gemini3ProImageRequestAspectRatioAuto, wire: "\"auto\""},
		{variant: Gemini3ProImageRequestAspectRatioOne1, wire: "\"1:1\""},
		{variant: Gemini3ProImageRequestAspectRatioTwo3, wire: "\"2:3\""},
		{variant: Gemini3ProImageRequestAspectRatioThree2, wire: "\"3:2\""},
		{variant: Gemini3ProImageRequestAspectRatioThree4, wire: "\"3:4\""},
		{variant: Gemini3ProImageRequestAspectRatioFour3, wire: "\"4:3\""},
		{variant: Gemini3ProImageRequestAspectRatioFour5, wire: "\"4:5\""},
		{variant: Gemini3ProImageRequestAspectRatioFive4, wire: "\"5:4\""},
		{variant: Gemini3ProImageRequestAspectRatioNine16, wire: "\"9:16\""},
		{variant: Gemini3ProImageRequestAspectRatioOne69, wire: "\"16:9\""},
		{variant: Gemini3ProImageRequestAspectRatioTwo19, wire: "\"21:9\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Gemini3ProImageRequestAspectRatio
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
