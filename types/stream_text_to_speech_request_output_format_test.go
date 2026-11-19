package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestStreamTextToSpeechRequestOutputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant StreamTextToSpeechRequestOutputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: StreamTextToSpeechRequestOutputFormatMp32205032, wire: "\"mp3_22050_32\""},
		{variant: StreamTextToSpeechRequestOutputFormatMp32400048, wire: "\"mp3_24000_48\""},
		{variant: StreamTextToSpeechRequestOutputFormatMp34410032, wire: "\"mp3_44100_32\""},
		{variant: StreamTextToSpeechRequestOutputFormatMp34410064, wire: "\"mp3_44100_64\""},
		{variant: StreamTextToSpeechRequestOutputFormatMp34410096, wire: "\"mp3_44100_96\""},
		{variant: StreamTextToSpeechRequestOutputFormatMp344100128, wire: "\"mp3_44100_128\""},
		{variant: StreamTextToSpeechRequestOutputFormatMp344100192, wire: "\"mp3_44100_192\""},
		{variant: StreamTextToSpeechRequestOutputFormatPcm8000, wire: "\"pcm_8000\""},
		{variant: StreamTextToSpeechRequestOutputFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: StreamTextToSpeechRequestOutputFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: StreamTextToSpeechRequestOutputFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: StreamTextToSpeechRequestOutputFormatPcm32000, wire: "\"pcm_32000\""},
		{variant: StreamTextToSpeechRequestOutputFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: StreamTextToSpeechRequestOutputFormatPcm48000, wire: "\"pcm_48000\""},
		{variant: StreamTextToSpeechRequestOutputFormatUlaw8000, wire: "\"ulaw_8000\""},
		{variant: StreamTextToSpeechRequestOutputFormatAlaw8000, wire: "\"alaw_8000\""},
		{variant: StreamTextToSpeechRequestOutputFormatOpus4800032, wire: "\"opus_48000_32\""},
		{variant: StreamTextToSpeechRequestOutputFormatOpus4800064, wire: "\"opus_48000_64\""},
		{variant: StreamTextToSpeechRequestOutputFormatOpus4800096, wire: "\"opus_48000_96\""},
		{variant: StreamTextToSpeechRequestOutputFormatOpus48000128, wire: "\"opus_48000_128\""},
		{variant: StreamTextToSpeechRequestOutputFormatOpus48000192, wire: "\"opus_48000_192\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded StreamTextToSpeechRequestOutputFormat
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
