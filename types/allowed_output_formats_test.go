package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAllowedOutputFormatsVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant AllowedOutputFormats
		wire    string
	}
	cases := []enumCase{
		{variant: AllowedOutputFormatsMp32205032, wire: "\"mp3_22050_32\""},
		{variant: AllowedOutputFormatsMp32400048, wire: "\"mp3_24000_48\""},
		{variant: AllowedOutputFormatsMp34410032, wire: "\"mp3_44100_32\""},
		{variant: AllowedOutputFormatsMp34410064, wire: "\"mp3_44100_64\""},
		{variant: AllowedOutputFormatsMp34410096, wire: "\"mp3_44100_96\""},
		{variant: AllowedOutputFormatsMp344100128, wire: "\"mp3_44100_128\""},
		{variant: AllowedOutputFormatsMp344100192, wire: "\"mp3_44100_192\""},
		{variant: AllowedOutputFormatsPcm8000, wire: "\"pcm_8000\""},
		{variant: AllowedOutputFormatsPcm16000, wire: "\"pcm_16000\""},
		{variant: AllowedOutputFormatsPcm22050, wire: "\"pcm_22050\""},
		{variant: AllowedOutputFormatsPcm24000, wire: "\"pcm_24000\""},
		{variant: AllowedOutputFormatsPcm32000, wire: "\"pcm_32000\""},
		{variant: AllowedOutputFormatsPcm44100, wire: "\"pcm_44100\""},
		{variant: AllowedOutputFormatsPcm48000, wire: "\"pcm_48000\""},
		{variant: AllowedOutputFormatsUlaw8000, wire: "\"ulaw_8000\""},
		{variant: AllowedOutputFormatsAlaw8000, wire: "\"alaw_8000\""},
		{variant: AllowedOutputFormatsOpus4800032, wire: "\"opus_48000_32\""},
		{variant: AllowedOutputFormatsOpus4800064, wire: "\"opus_48000_64\""},
		{variant: AllowedOutputFormatsOpus4800096, wire: "\"opus_48000_96\""},
		{variant: AllowedOutputFormatsOpus48000128, wire: "\"opus_48000_128\""},
		{variant: AllowedOutputFormatsOpus48000192, wire: "\"opus_48000_192\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded AllowedOutputFormats
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

func TestAllowedOutputFormatsVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant AllowedOutputFormats
		wire    string
	}
	cases := []enumCase{
		{variant: AllowedOutputFormatsMp32205032, wire: "\"mp3_22050_32\""},
		{variant: AllowedOutputFormatsMp32400048, wire: "\"mp3_24000_48\""},
		{variant: AllowedOutputFormatsMp34410032, wire: "\"mp3_44100_32\""},
		{variant: AllowedOutputFormatsMp34410064, wire: "\"mp3_44100_64\""},
		{variant: AllowedOutputFormatsMp34410096, wire: "\"mp3_44100_96\""},
		{variant: AllowedOutputFormatsMp344100128, wire: "\"mp3_44100_128\""},
		{variant: AllowedOutputFormatsMp344100192, wire: "\"mp3_44100_192\""},
		{variant: AllowedOutputFormatsPcm8000, wire: "\"pcm_8000\""},
		{variant: AllowedOutputFormatsPcm16000, wire: "\"pcm_16000\""},
		{variant: AllowedOutputFormatsPcm22050, wire: "\"pcm_22050\""},
		{variant: AllowedOutputFormatsPcm24000, wire: "\"pcm_24000\""},
		{variant: AllowedOutputFormatsPcm32000, wire: "\"pcm_32000\""},
		{variant: AllowedOutputFormatsPcm44100, wire: "\"pcm_44100\""},
		{variant: AllowedOutputFormatsPcm48000, wire: "\"pcm_48000\""},
		{variant: AllowedOutputFormatsUlaw8000, wire: "\"ulaw_8000\""},
		{variant: AllowedOutputFormatsAlaw8000, wire: "\"alaw_8000\""},
		{variant: AllowedOutputFormatsOpus4800032, wire: "\"opus_48000_32\""},
		{variant: AllowedOutputFormatsOpus4800064, wire: "\"opus_48000_64\""},
		{variant: AllowedOutputFormatsOpus4800096, wire: "\"opus_48000_96\""},
		{variant: AllowedOutputFormatsOpus48000128, wire: "\"opus_48000_128\""},
		{variant: AllowedOutputFormatsOpus48000192, wire: "\"opus_48000_192\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded AllowedOutputFormats
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

func TestAllowedOutputFormatsVariantsRoundTrip3(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant AllowedOutputFormats
		wire    string
	}
	cases := []enumCase{
		{variant: AllowedOutputFormatsMp32205032, wire: "\"mp3_22050_32\""},
		{variant: AllowedOutputFormatsMp32400048, wire: "\"mp3_24000_48\""},
		{variant: AllowedOutputFormatsMp34410032, wire: "\"mp3_44100_32\""},
		{variant: AllowedOutputFormatsMp34410064, wire: "\"mp3_44100_64\""},
		{variant: AllowedOutputFormatsMp34410096, wire: "\"mp3_44100_96\""},
		{variant: AllowedOutputFormatsMp344100128, wire: "\"mp3_44100_128\""},
		{variant: AllowedOutputFormatsMp344100192, wire: "\"mp3_44100_192\""},
		{variant: AllowedOutputFormatsPcm8000, wire: "\"pcm_8000\""},
		{variant: AllowedOutputFormatsPcm16000, wire: "\"pcm_16000\""},
		{variant: AllowedOutputFormatsPcm22050, wire: "\"pcm_22050\""},
		{variant: AllowedOutputFormatsPcm24000, wire: "\"pcm_24000\""},
		{variant: AllowedOutputFormatsPcm32000, wire: "\"pcm_32000\""},
		{variant: AllowedOutputFormatsPcm44100, wire: "\"pcm_44100\""},
		{variant: AllowedOutputFormatsPcm48000, wire: "\"pcm_48000\""},
		{variant: AllowedOutputFormatsUlaw8000, wire: "\"ulaw_8000\""},
		{variant: AllowedOutputFormatsAlaw8000, wire: "\"alaw_8000\""},
		{variant: AllowedOutputFormatsOpus4800032, wire: "\"opus_48000_32\""},
		{variant: AllowedOutputFormatsOpus4800064, wire: "\"opus_48000_64\""},
		{variant: AllowedOutputFormatsOpus4800096, wire: "\"opus_48000_96\""},
		{variant: AllowedOutputFormatsOpus48000128, wire: "\"opus_48000_128\""},
		{variant: AllowedOutputFormatsOpus48000192, wire: "\"opus_48000_192\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded AllowedOutputFormats
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

func TestAllowedOutputFormatsVariantsRoundTrip4(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant AllowedOutputFormats
		wire    string
	}
	cases := []enumCase{
		{variant: AllowedOutputFormatsMp32205032, wire: "\"mp3_22050_32\""},
		{variant: AllowedOutputFormatsMp32400048, wire: "\"mp3_24000_48\""},
		{variant: AllowedOutputFormatsMp34410032, wire: "\"mp3_44100_32\""},
		{variant: AllowedOutputFormatsMp34410064, wire: "\"mp3_44100_64\""},
		{variant: AllowedOutputFormatsMp34410096, wire: "\"mp3_44100_96\""},
		{variant: AllowedOutputFormatsMp344100128, wire: "\"mp3_44100_128\""},
		{variant: AllowedOutputFormatsMp344100192, wire: "\"mp3_44100_192\""},
		{variant: AllowedOutputFormatsPcm8000, wire: "\"pcm_8000\""},
		{variant: AllowedOutputFormatsPcm16000, wire: "\"pcm_16000\""},
		{variant: AllowedOutputFormatsPcm22050, wire: "\"pcm_22050\""},
		{variant: AllowedOutputFormatsPcm24000, wire: "\"pcm_24000\""},
		{variant: AllowedOutputFormatsPcm32000, wire: "\"pcm_32000\""},
		{variant: AllowedOutputFormatsPcm44100, wire: "\"pcm_44100\""},
		{variant: AllowedOutputFormatsPcm48000, wire: "\"pcm_48000\""},
		{variant: AllowedOutputFormatsUlaw8000, wire: "\"ulaw_8000\""},
		{variant: AllowedOutputFormatsAlaw8000, wire: "\"alaw_8000\""},
		{variant: AllowedOutputFormatsOpus4800032, wire: "\"opus_48000_32\""},
		{variant: AllowedOutputFormatsOpus4800064, wire: "\"opus_48000_64\""},
		{variant: AllowedOutputFormatsOpus4800096, wire: "\"opus_48000_96\""},
		{variant: AllowedOutputFormatsOpus48000128, wire: "\"opus_48000_128\""},
		{variant: AllowedOutputFormatsOpus48000192, wire: "\"opus_48000_192\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded AllowedOutputFormats
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
