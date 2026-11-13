package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestClipAnimationEnterEffectVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ClipAnimationEnterEffect
		wire    string
	}
	cases := []enumCase{
		{variant: ClipAnimationEnterEffectNone, wire: "\"none\""},
		{variant: ClipAnimationEnterEffectFade, wire: "\"fade\""},
		{variant: ClipAnimationEnterEffectFloat, wire: "\"float\""},
		{variant: ClipAnimationEnterEffectGentleFloat, wire: "\"gentle_float\""},
		{variant: ClipAnimationEnterEffectZoomIn, wire: "\"zoom_in\""},
		{variant: ClipAnimationEnterEffectDrop, wire: "\"drop\""},
		{variant: ClipAnimationEnterEffectSlideLeft, wire: "\"slide_left\""},
		{variant: ClipAnimationEnterEffectSlideRight, wire: "\"slide_right\""},
		{variant: ClipAnimationEnterEffectSlideUp, wire: "\"slide_up\""},
		{variant: ClipAnimationEnterEffectSlideDown, wire: "\"slide_down\""},
		{variant: ClipAnimationEnterEffectPop, wire: "\"pop\""},
		{variant: ClipAnimationEnterEffectBounce, wire: "\"bounce\""},
		{variant: ClipAnimationEnterEffectSpin, wire: "\"spin\""},
		{variant: ClipAnimationEnterEffectSlideBounce, wire: "\"slide_bounce\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ClipAnimationEnterEffect
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
