package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestConvertTextToSpeechRequestOutputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ConvertTextToSpeechRequestOutputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: ConvertTextToSpeechRequestOutputFormatAlaw8000, wire: "\"alaw_8000\""},
		{variant: ConvertTextToSpeechRequestOutputFormatMp32205032, wire: "\"mp3_22050_32\""},
		{variant: ConvertTextToSpeechRequestOutputFormatMp32400048, wire: "\"mp3_24000_48\""},
		{variant: ConvertTextToSpeechRequestOutputFormatMp344100128, wire: "\"mp3_44100_128\""},
		{variant: ConvertTextToSpeechRequestOutputFormatMp344100192, wire: "\"mp3_44100_192\""},
		{variant: ConvertTextToSpeechRequestOutputFormatMp34410032, wire: "\"mp3_44100_32\""},
		{variant: ConvertTextToSpeechRequestOutputFormatMp34410064, wire: "\"mp3_44100_64\""},
		{variant: ConvertTextToSpeechRequestOutputFormatMp34410096, wire: "\"mp3_44100_96\""},
		{variant: ConvertTextToSpeechRequestOutputFormatOpus48000128, wire: "\"opus_48000_128\""},
		{variant: ConvertTextToSpeechRequestOutputFormatOpus48000192, wire: "\"opus_48000_192\""},
		{variant: ConvertTextToSpeechRequestOutputFormatOpus4800032, wire: "\"opus_48000_32\""},
		{variant: ConvertTextToSpeechRequestOutputFormatOpus4800064, wire: "\"opus_48000_64\""},
		{variant: ConvertTextToSpeechRequestOutputFormatOpus4800096, wire: "\"opus_48000_96\""},
		{variant: ConvertTextToSpeechRequestOutputFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: ConvertTextToSpeechRequestOutputFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: ConvertTextToSpeechRequestOutputFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: ConvertTextToSpeechRequestOutputFormatPcm32000, wire: "\"pcm_32000\""},
		{variant: ConvertTextToSpeechRequestOutputFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: ConvertTextToSpeechRequestOutputFormatPcm48000, wire: "\"pcm_48000\""},
		{variant: ConvertTextToSpeechRequestOutputFormatPcm8000, wire: "\"pcm_8000\""},
		{variant: ConvertTextToSpeechRequestOutputFormatUlaw8000, wire: "\"ulaw_8000\""},
		{variant: ConvertTextToSpeechRequestOutputFormatWav16000, wire: "\"wav_16000\""},
		{variant: ConvertTextToSpeechRequestOutputFormatWav22050, wire: "\"wav_22050\""},
		{variant: ConvertTextToSpeechRequestOutputFormatWav24000, wire: "\"wav_24000\""},
		{variant: ConvertTextToSpeechRequestOutputFormatWav32000, wire: "\"wav_32000\""},
		{variant: ConvertTextToSpeechRequestOutputFormatWav44100, wire: "\"wav_44100\""},
		{variant: ConvertTextToSpeechRequestOutputFormatWav48000, wire: "\"wav_48000\""},
		{variant: ConvertTextToSpeechRequestOutputFormatWav8000, wire: "\"wav_8000\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ConvertTextToSpeechRequestOutputFormat
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
