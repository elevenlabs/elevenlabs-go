package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestOutputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant OutputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: OutputFormatMp32205032, wire: "\"mp3_22050_32\""},
		{variant: OutputFormatMp34410032, wire: "\"mp3_44100_32\""},
		{variant: OutputFormatMp34410064, wire: "\"mp3_44100_64\""},
		{variant: OutputFormatMp34410096, wire: "\"mp3_44100_96\""},
		{variant: OutputFormatMp344100128, wire: "\"mp3_44100_128\""},
		{variant: OutputFormatMp344100192, wire: "\"mp3_44100_192\""},
		{variant: OutputFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: OutputFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: OutputFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: OutputFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: OutputFormatUlaw8000, wire: "\"ulaw_8000\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded OutputFormat
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
