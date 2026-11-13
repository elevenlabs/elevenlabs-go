package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGemini3ProImageRequestResolutionVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant Gemini3ProImageRequestResolution
		wire    string
	}
	cases := []enumCase{
		{variant: Gemini3ProImageRequestResolutionOneK, wire: "\"1K\""},
		{variant: Gemini3ProImageRequestResolutionTwoK, wire: "\"2K\""},
		{variant: Gemini3ProImageRequestResolutionFourK, wire: "\"4K\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded Gemini3ProImageRequestResolution
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
