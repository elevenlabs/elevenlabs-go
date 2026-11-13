package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGemini31FlashImageRequestAspectRatioVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Gemini31FlashImageRequestAspectRatio
		wire    string
	}
	cases := []enumCase{
		{variant: Gemini31FlashImageRequestAspectRatioAuto, wire: "\"auto\""},
		{variant: Gemini31FlashImageRequestAspectRatioOne1, wire: "\"1:1\""},
		{variant: Gemini31FlashImageRequestAspectRatioTwo3, wire: "\"2:3\""},
		{variant: Gemini31FlashImageRequestAspectRatioThree2, wire: "\"3:2\""},
		{variant: Gemini31FlashImageRequestAspectRatioThree4, wire: "\"3:4\""},
		{variant: Gemini31FlashImageRequestAspectRatioFour3, wire: "\"4:3\""},
		{variant: Gemini31FlashImageRequestAspectRatioFour5, wire: "\"4:5\""},
		{variant: Gemini31FlashImageRequestAspectRatioFive4, wire: "\"5:4\""},
		{variant: Gemini31FlashImageRequestAspectRatioNine16, wire: "\"9:16\""},
		{variant: Gemini31FlashImageRequestAspectRatioOne69, wire: "\"16:9\""},
		{variant: Gemini31FlashImageRequestAspectRatioTwo19, wire: "\"21:9\""},
		{variant: Gemini31FlashImageRequestAspectRatioOne4, wire: "\"1:4\""},
		{variant: Gemini31FlashImageRequestAspectRatioFour1, wire: "\"4:1\""},
		{variant: Gemini31FlashImageRequestAspectRatioOne8, wire: "\"1:8\""},
		{variant: Gemini31FlashImageRequestAspectRatioEight1, wire: "\"8:1\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Gemini31FlashImageRequestAspectRatio
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
