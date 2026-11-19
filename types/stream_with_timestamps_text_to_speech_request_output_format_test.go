package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestStreamWithTimestampsTextToSpeechRequestOutputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant StreamWithTimestampsTextToSpeechRequestOutputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatMp32205032, wire: "\"mp3_22050_32\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatMp32400048, wire: "\"mp3_24000_48\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatMp34410032, wire: "\"mp3_44100_32\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatMp34410064, wire: "\"mp3_44100_64\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatMp34410096, wire: "\"mp3_44100_96\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatMp344100128, wire: "\"mp3_44100_128\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatMp344100192, wire: "\"mp3_44100_192\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatPcm8000, wire: "\"pcm_8000\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatPcm32000, wire: "\"pcm_32000\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatPcm48000, wire: "\"pcm_48000\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatUlaw8000, wire: "\"ulaw_8000\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatAlaw8000, wire: "\"alaw_8000\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatOpus4800032, wire: "\"opus_48000_32\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatOpus4800064, wire: "\"opus_48000_64\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatOpus4800096, wire: "\"opus_48000_96\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatOpus48000128, wire: "\"opus_48000_128\""},
		{variant: StreamWithTimestampsTextToSpeechRequestOutputFormatOpus48000192, wire: "\"opus_48000_192\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded StreamWithTimestampsTextToSpeechRequestOutputFormat
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
