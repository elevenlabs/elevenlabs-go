package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestNonStreamingOutputFormatsVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant NonStreamingOutputFormats
		wire    string
	}
	cases := []enumCase{
		{variant: NonStreamingOutputFormatsWav8000, wire: "\"wav_8000\""},
		{variant: NonStreamingOutputFormatsWav16000, wire: "\"wav_16000\""},
		{variant: NonStreamingOutputFormatsWav22050, wire: "\"wav_22050\""},
		{variant: NonStreamingOutputFormatsWav24000, wire: "\"wav_24000\""},
		{variant: NonStreamingOutputFormatsWav32000, wire: "\"wav_32000\""},
		{variant: NonStreamingOutputFormatsWav44100, wire: "\"wav_44100\""},
		{variant: NonStreamingOutputFormatsWav48000, wire: "\"wav_48000\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded NonStreamingOutputFormats
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
