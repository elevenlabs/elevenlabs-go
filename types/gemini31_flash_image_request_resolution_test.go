package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGemini31FlashImageRequestResolutionVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Gemini31FlashImageRequestResolution
		wire    string
	}
	cases := []enumCase{
		{variant: Gemini31FlashImageRequestResolutionFive12, wire: "\"512\""},
		{variant: Gemini31FlashImageRequestResolutionOneK, wire: "\"1K\""},
		{variant: Gemini31FlashImageRequestResolutionTwoK, wire: "\"2K\""},
		{variant: Gemini31FlashImageRequestResolutionFourK, wire: "\"4K\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Gemini31FlashImageRequestResolution
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
