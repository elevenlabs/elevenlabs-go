package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestComposeDetailedMusicRequestOutputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ComposeDetailedMusicRequestOutputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: ComposeDetailedMusicRequestOutputFormatAuto, wire: "\"auto\""},
		{variant: ComposeDetailedMusicRequestOutputFormatMp348000128, wire: "\"mp3_48000_128\""},
		{variant: ComposeDetailedMusicRequestOutputFormatMp348000192, wire: "\"mp3_48000_192\""},
		{variant: ComposeDetailedMusicRequestOutputFormatMp348000240, wire: "\"mp3_48000_240\""},
		{variant: ComposeDetailedMusicRequestOutputFormatMp348000320, wire: "\"mp3_48000_320\""},
		{variant: ComposeDetailedMusicRequestOutputFormatMp32205032, wire: "\"mp3_22050_32\""},
		{variant: ComposeDetailedMusicRequestOutputFormatMp32400048, wire: "\"mp3_24000_48\""},
		{variant: ComposeDetailedMusicRequestOutputFormatMp34410032, wire: "\"mp3_44100_32\""},
		{variant: ComposeDetailedMusicRequestOutputFormatMp34410064, wire: "\"mp3_44100_64\""},
		{variant: ComposeDetailedMusicRequestOutputFormatMp34410096, wire: "\"mp3_44100_96\""},
		{variant: ComposeDetailedMusicRequestOutputFormatMp344100128, wire: "\"mp3_44100_128\""},
		{variant: ComposeDetailedMusicRequestOutputFormatMp344100192, wire: "\"mp3_44100_192\""},
		{variant: ComposeDetailedMusicRequestOutputFormatPcm8000, wire: "\"pcm_8000\""},
		{variant: ComposeDetailedMusicRequestOutputFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: ComposeDetailedMusicRequestOutputFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: ComposeDetailedMusicRequestOutputFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: ComposeDetailedMusicRequestOutputFormatPcm32000, wire: "\"pcm_32000\""},
		{variant: ComposeDetailedMusicRequestOutputFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: ComposeDetailedMusicRequestOutputFormatPcm48000, wire: "\"pcm_48000\""},
		{variant: ComposeDetailedMusicRequestOutputFormatUlaw8000, wire: "\"ulaw_8000\""},
		{variant: ComposeDetailedMusicRequestOutputFormatAlaw8000, wire: "\"alaw_8000\""},
		{variant: ComposeDetailedMusicRequestOutputFormatOpus4800032, wire: "\"opus_48000_32\""},
		{variant: ComposeDetailedMusicRequestOutputFormatOpus4800064, wire: "\"opus_48000_64\""},
		{variant: ComposeDetailedMusicRequestOutputFormatOpus4800096, wire: "\"opus_48000_96\""},
		{variant: ComposeDetailedMusicRequestOutputFormatOpus48000128, wire: "\"opus_48000_128\""},
		{variant: ComposeDetailedMusicRequestOutputFormatOpus48000192, wire: "\"opus_48000_192\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ComposeDetailedMusicRequestOutputFormat
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
