package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestComposeDetailedStreamMusicRequestOutputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ComposeDetailedStreamMusicRequestOutputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: ComposeDetailedStreamMusicRequestOutputFormatAuto, wire: "\"auto\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatMp348000128, wire: "\"mp3_48000_128\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatMp348000192, wire: "\"mp3_48000_192\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatMp348000240, wire: "\"mp3_48000_240\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatMp348000320, wire: "\"mp3_48000_320\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatMp32205032, wire: "\"mp3_22050_32\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatMp32400048, wire: "\"mp3_24000_48\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatMp34410032, wire: "\"mp3_44100_32\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatMp34410064, wire: "\"mp3_44100_64\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatMp34410096, wire: "\"mp3_44100_96\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatMp344100128, wire: "\"mp3_44100_128\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatMp344100192, wire: "\"mp3_44100_192\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatPcm8000, wire: "\"pcm_8000\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatPcm32000, wire: "\"pcm_32000\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatPcm48000, wire: "\"pcm_48000\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatUlaw8000, wire: "\"ulaw_8000\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatAlaw8000, wire: "\"alaw_8000\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatOpus4800032, wire: "\"opus_48000_32\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatOpus4800064, wire: "\"opus_48000_64\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatOpus4800096, wire: "\"opus_48000_96\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatOpus48000128, wire: "\"opus_48000_128\""},
		{variant: ComposeDetailedStreamMusicRequestOutputFormatOpus48000192, wire: "\"opus_48000_192\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ComposeDetailedStreamMusicRequestOutputFormat
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
