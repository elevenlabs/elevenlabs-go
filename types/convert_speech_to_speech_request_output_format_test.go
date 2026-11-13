package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestConvertSpeechToSpeechRequestOutputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ConvertSpeechToSpeechRequestOutputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: ConvertSpeechToSpeechRequestOutputFormatAlaw8000, wire: "\"alaw_8000\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatMp32205032, wire: "\"mp3_22050_32\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatMp32400048, wire: "\"mp3_24000_48\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatMp344100128, wire: "\"mp3_44100_128\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatMp344100192, wire: "\"mp3_44100_192\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatMp34410032, wire: "\"mp3_44100_32\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatMp34410064, wire: "\"mp3_44100_64\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatMp34410096, wire: "\"mp3_44100_96\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatOpus48000128, wire: "\"opus_48000_128\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatOpus48000192, wire: "\"opus_48000_192\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatOpus4800032, wire: "\"opus_48000_32\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatOpus4800064, wire: "\"opus_48000_64\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatOpus4800096, wire: "\"opus_48000_96\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatPcm32000, wire: "\"pcm_32000\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatPcm48000, wire: "\"pcm_48000\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatPcm8000, wire: "\"pcm_8000\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatUlaw8000, wire: "\"ulaw_8000\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatWav16000, wire: "\"wav_16000\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatWav22050, wire: "\"wav_22050\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatWav24000, wire: "\"wav_24000\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatWav32000, wire: "\"wav_32000\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatWav44100, wire: "\"wav_44100\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatWav48000, wire: "\"wav_48000\""},
		{variant: ConvertSpeechToSpeechRequestOutputFormatWav8000, wire: "\"wav_8000\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ConvertSpeechToSpeechRequestOutputFormat
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
