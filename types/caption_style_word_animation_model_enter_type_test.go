package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCaptionStyleWordAnimationModelEnterTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant CaptionStyleWordAnimationModelEnterType
		wire    string
	}
	cases := []enumCase{
		{variant: CaptionStyleWordAnimationModelEnterTypeNone, wire: "\"none\""},
		{variant: CaptionStyleWordAnimationModelEnterTypeFade, wire: "\"fade\""},
		{variant: CaptionStyleWordAnimationModelEnterTypeScale, wire: "\"scale\""},
		{variant: CaptionStyleWordAnimationModelEnterTypePop, wire: "\"pop\""},
		{variant: CaptionStyleWordAnimationModelEnterTypeSlideUp, wire: "\"slide_up\""},
		{variant: CaptionStyleWordAnimationModelEnterTypeSlideDown, wire: "\"slide_down\""},
		{variant: CaptionStyleWordAnimationModelEnterTypeSlam, wire: "\"slam\""},
		{variant: CaptionStyleWordAnimationModelEnterTypeScaleDown, wire: "\"scale_down\""},
		{variant: CaptionStyleWordAnimationModelEnterTypeSlideIn, wire: "\"slide_in\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded CaptionStyleWordAnimationModelEnterType
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
