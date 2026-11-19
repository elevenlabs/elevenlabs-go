package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestConvertWithTimestampsTextToSpeechRequestOutputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ConvertWithTimestampsTextToSpeechRequestOutputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatAlaw8000, wire: "\"alaw_8000\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatMp32205032, wire: "\"mp3_22050_32\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatMp32400048, wire: "\"mp3_24000_48\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatMp344100128, wire: "\"mp3_44100_128\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatMp344100192, wire: "\"mp3_44100_192\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatMp34410032, wire: "\"mp3_44100_32\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatMp34410064, wire: "\"mp3_44100_64\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatMp34410096, wire: "\"mp3_44100_96\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus48000128, wire: "\"opus_48000_128\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus48000192, wire: "\"opus_48000_192\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus4800032, wire: "\"opus_48000_32\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus4800064, wire: "\"opus_48000_64\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatOpus4800096, wire: "\"opus_48000_96\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm32000, wire: "\"pcm_32000\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm48000, wire: "\"pcm_48000\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatPcm8000, wire: "\"pcm_8000\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatUlaw8000, wire: "\"ulaw_8000\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatWav16000, wire: "\"wav_16000\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatWav22050, wire: "\"wav_22050\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatWav24000, wire: "\"wav_24000\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatWav32000, wire: "\"wav_32000\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatWav44100, wire: "\"wav_44100\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatWav48000, wire: "\"wav_48000\""},
		{variant: ConvertWithTimestampsTextToSpeechRequestOutputFormatWav8000, wire: "\"wav_8000\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ConvertWithTimestampsTextToSpeechRequestOutputFormat
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
