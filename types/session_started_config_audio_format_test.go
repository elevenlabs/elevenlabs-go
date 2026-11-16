package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSessionStartedConfigAudioFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant SessionStartedConfigAudioFormat
		wire    string
	}
	cases := []enumCase{
		{variant: SessionStartedConfigAudioFormatPcm8000, wire: "\"pcm_8000\""},
		{variant: SessionStartedConfigAudioFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: SessionStartedConfigAudioFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: SessionStartedConfigAudioFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: SessionStartedConfigAudioFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: SessionStartedConfigAudioFormatPcm48000, wire: "\"pcm_48000\""},
		{variant: SessionStartedConfigAudioFormatUlaw8000, wire: "\"ulaw_8000\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded SessionStartedConfigAudioFormat
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
