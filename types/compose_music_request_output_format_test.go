package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestComposeMusicRequestOutputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ComposeMusicRequestOutputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: ComposeMusicRequestOutputFormatAuto, wire: "\"auto\""},
		{variant: ComposeMusicRequestOutputFormatMp348000128, wire: "\"mp3_48000_128\""},
		{variant: ComposeMusicRequestOutputFormatMp348000192, wire: "\"mp3_48000_192\""},
		{variant: ComposeMusicRequestOutputFormatMp348000240, wire: "\"mp3_48000_240\""},
		{variant: ComposeMusicRequestOutputFormatMp348000320, wire: "\"mp3_48000_320\""},
		{variant: ComposeMusicRequestOutputFormatMp32205032, wire: "\"mp3_22050_32\""},
		{variant: ComposeMusicRequestOutputFormatMp32400048, wire: "\"mp3_24000_48\""},
		{variant: ComposeMusicRequestOutputFormatMp34410032, wire: "\"mp3_44100_32\""},
		{variant: ComposeMusicRequestOutputFormatMp34410064, wire: "\"mp3_44100_64\""},
		{variant: ComposeMusicRequestOutputFormatMp34410096, wire: "\"mp3_44100_96\""},
		{variant: ComposeMusicRequestOutputFormatMp344100128, wire: "\"mp3_44100_128\""},
		{variant: ComposeMusicRequestOutputFormatMp344100192, wire: "\"mp3_44100_192\""},
		{variant: ComposeMusicRequestOutputFormatPcm8000, wire: "\"pcm_8000\""},
		{variant: ComposeMusicRequestOutputFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: ComposeMusicRequestOutputFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: ComposeMusicRequestOutputFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: ComposeMusicRequestOutputFormatPcm32000, wire: "\"pcm_32000\""},
		{variant: ComposeMusicRequestOutputFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: ComposeMusicRequestOutputFormatPcm48000, wire: "\"pcm_48000\""},
		{variant: ComposeMusicRequestOutputFormatUlaw8000, wire: "\"ulaw_8000\""},
		{variant: ComposeMusicRequestOutputFormatAlaw8000, wire: "\"alaw_8000\""},
		{variant: ComposeMusicRequestOutputFormatOpus4800032, wire: "\"opus_48000_32\""},
		{variant: ComposeMusicRequestOutputFormatOpus4800064, wire: "\"opus_48000_64\""},
		{variant: ComposeMusicRequestOutputFormatOpus4800096, wire: "\"opus_48000_96\""},
		{variant: ComposeMusicRequestOutputFormatOpus48000128, wire: "\"opus_48000_128\""},
		{variant: ComposeMusicRequestOutputFormatOpus48000192, wire: "\"opus_48000_192\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ComposeMusicRequestOutputFormat
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
