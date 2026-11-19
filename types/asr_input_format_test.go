package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestAsrInputFormatVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant AsrInputFormat
		wire    string
	}
	cases := []enumCase{
		{variant: AsrInputFormatPcm8000, wire: "\"pcm_8000\""},
		{variant: AsrInputFormatPcm16000, wire: "\"pcm_16000\""},
		{variant: AsrInputFormatPcm22050, wire: "\"pcm_22050\""},
		{variant: AsrInputFormatPcm24000, wire: "\"pcm_24000\""},
		{variant: AsrInputFormatPcm44100, wire: "\"pcm_44100\""},
		{variant: AsrInputFormatPcm48000, wire: "\"pcm_48000\""},
		{variant: AsrInputFormatUlaw8000, wire: "\"ulaw_8000\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded AsrInputFormat
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
