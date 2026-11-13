package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestStreamMusicRequestOutputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant StreamMusicRequestOutputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: StreamMusicRequestOutputFormatAuto, wire: "\"auto\""},
		{variant: StreamMusicRequestOutputFormatMp348000128, wire: "\"mp3_48000_128\""},
		{variant: StreamMusicRequestOutputFormatMp348000192, wire: "\"mp3_48000_192\""},
		{variant: StreamMusicRequestOutputFormatMp348000240, wire: "\"mp3_48000_240\""},
		{variant: StreamMusicRequestOutputFormatMp348000320, wire: "\"mp3_48000_320\""},
		{variant: StreamMusicRequestOutputFormatMp32205032, wire: "\"mp3_22050_32\""},
		{variant: StreamMusicRequestOutputFormatMp32400048, wire: "\"mp3_24000_48\""},
		{variant: StreamMusicRequestOutputFormatMp34410032, wire: "\"mp3_44100_32\""},
		{variant: StreamMusicRequestOutputFormatMp34410064, wire: "\"mp3_44100_64\""},
		{variant: StreamMusicRequestOutputFormatMp34410096, wire: "\"mp3_44100_96\""},
		{variant: StreamMusicRequestOutputFormatMp344100128, wire: "\"mp3_44100_128\""},
		{variant: StreamMusicRequestOutputFormatMp344100192, wire: "\"mp3_44100_192\""},
		{variant: StreamMusicRequestOutputFormatPcm8000, wire: "\"pcm_8000\""},
		{variant: StreamMusicRequestOutputFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: StreamMusicRequestOutputFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: StreamMusicRequestOutputFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: StreamMusicRequestOutputFormatPcm32000, wire: "\"pcm_32000\""},
		{variant: StreamMusicRequestOutputFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: StreamMusicRequestOutputFormatPcm48000, wire: "\"pcm_48000\""},
		{variant: StreamMusicRequestOutputFormatUlaw8000, wire: "\"ulaw_8000\""},
		{variant: StreamMusicRequestOutputFormatAlaw8000, wire: "\"alaw_8000\""},
		{variant: StreamMusicRequestOutputFormatOpus4800032, wire: "\"opus_48000_32\""},
		{variant: StreamMusicRequestOutputFormatOpus4800064, wire: "\"opus_48000_64\""},
		{variant: StreamMusicRequestOutputFormatOpus4800096, wire: "\"opus_48000_96\""},
		{variant: StreamMusicRequestOutputFormatOpus48000128, wire: "\"opus_48000_128\""},
		{variant: StreamMusicRequestOutputFormatOpus48000192, wire: "\"opus_48000_192\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded StreamMusicRequestOutputFormat
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
