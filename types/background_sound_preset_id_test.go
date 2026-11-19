package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestBackgroundSoundPresetIdVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant BackgroundSoundPresetId
		wire    string
	}
	cases := []enumCase{
		{variant: BackgroundSoundPresetIdOffice2, wire: "\"office2\""},
		{variant: BackgroundSoundPresetIdOffice1, wire: "\"office1\""},
		{variant: BackgroundSoundPresetIdRestaurant, wire: "\"restaurant\""},
		{variant: BackgroundSoundPresetIdCity, wire: "\"city\""},
		{variant: BackgroundSoundPresetIdTyping, wire: "\"typing\""},
		{variant: BackgroundSoundPresetIdElevator1, wire: "\"elevator1\""},
		{variant: BackgroundSoundPresetIdElevator2, wire: "\"elevator2\""},
		{variant: BackgroundSoundPresetIdElevator3, wire: "\"elevator3\""},
		{variant: BackgroundSoundPresetIdElevator4, wire: "\"elevator4\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded BackgroundSoundPresetId
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
