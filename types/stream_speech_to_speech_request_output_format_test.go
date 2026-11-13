package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestStreamSpeechToSpeechRequestOutputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant StreamSpeechToSpeechRequestOutputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: StreamSpeechToSpeechRequestOutputFormatMp32205032, wire: "\"mp3_22050_32\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatMp32400048, wire: "\"mp3_24000_48\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatMp34410032, wire: "\"mp3_44100_32\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatMp34410064, wire: "\"mp3_44100_64\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatMp34410096, wire: "\"mp3_44100_96\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatMp344100128, wire: "\"mp3_44100_128\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatMp344100192, wire: "\"mp3_44100_192\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatPcm8000, wire: "\"pcm_8000\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatPcm32000, wire: "\"pcm_32000\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatPcm48000, wire: "\"pcm_48000\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatUlaw8000, wire: "\"ulaw_8000\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatAlaw8000, wire: "\"alaw_8000\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatOpus4800032, wire: "\"opus_48000_32\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatOpus4800064, wire: "\"opus_48000_64\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatOpus4800096, wire: "\"opus_48000_96\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatOpus48000128, wire: "\"opus_48000_128\""},
		{variant: StreamSpeechToSpeechRequestOutputFormatOpus48000192, wire: "\"opus_48000_192\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded StreamSpeechToSpeechRequestOutputFormat
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
