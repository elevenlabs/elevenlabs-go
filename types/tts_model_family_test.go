package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestTtsModelFamilyVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant TtsModelFamily
		wire    string
	}
	cases := []enumCase{
		{variant: TtsModelFamilyTurbo, wire: "\"turbo\""},
		{variant: TtsModelFamilyFlash, wire: "\"flash\""},
		{variant: TtsModelFamilyMultilingual, wire: "\"multilingual\""},
		{variant: TtsModelFamilyV3Conversational, wire: "\"v3_conversational\""},
		{variant: TtsModelFamilyV4, wire: "\"v4\""},
		{variant: TtsModelFamilyV4Turbo, wire: "\"v4_turbo\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded TtsModelFamily
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
