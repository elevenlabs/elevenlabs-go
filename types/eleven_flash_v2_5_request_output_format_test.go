package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestElevenFlashV25RequestOutputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ElevenFlashV25RequestOutputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: ElevenFlashV25RequestOutputFormatMp32205032, wire: "\"mp3_22050_32\""},
		{variant: ElevenFlashV25RequestOutputFormatMp32400048, wire: "\"mp3_24000_48\""},
		{variant: ElevenFlashV25RequestOutputFormatMp34410032, wire: "\"mp3_44100_32\""},
		{variant: ElevenFlashV25RequestOutputFormatMp34410064, wire: "\"mp3_44100_64\""},
		{variant: ElevenFlashV25RequestOutputFormatMp34410096, wire: "\"mp3_44100_96\""},
		{variant: ElevenFlashV25RequestOutputFormatMp344100128, wire: "\"mp3_44100_128\""},
		{variant: ElevenFlashV25RequestOutputFormatMp344100192, wire: "\"mp3_44100_192\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ElevenFlashV25RequestOutputFormat
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
