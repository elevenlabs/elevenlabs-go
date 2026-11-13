package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestClipAnimationExitEffectVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ClipAnimationExitEffect
		wire    string
	}
	cases := []enumCase{
		{variant: ClipAnimationExitEffectNone, wire: "\"none\""},
		{variant: ClipAnimationExitEffectFade, wire: "\"fade\""},
		{variant: ClipAnimationExitEffectFloat, wire: "\"float\""},
		{variant: ClipAnimationExitEffectGentleFloat, wire: "\"gentle_float\""},
		{variant: ClipAnimationExitEffectZoomIn, wire: "\"zoom_in\""},
		{variant: ClipAnimationExitEffectDrop, wire: "\"drop\""},
		{variant: ClipAnimationExitEffectSlideLeft, wire: "\"slide_left\""},
		{variant: ClipAnimationExitEffectSlideRight, wire: "\"slide_right\""},
		{variant: ClipAnimationExitEffectSlideUp, wire: "\"slide_up\""},
		{variant: ClipAnimationExitEffectSlideDown, wire: "\"slide_down\""},
		{variant: ClipAnimationExitEffectPop, wire: "\"pop\""},
		{variant: ClipAnimationExitEffectBounce, wire: "\"bounce\""},
		{variant: ClipAnimationExitEffectSpin, wire: "\"spin\""},
		{variant: ClipAnimationExitEffectSlideBounce, wire: "\"slide_bounce\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ClipAnimationExitEffect
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
