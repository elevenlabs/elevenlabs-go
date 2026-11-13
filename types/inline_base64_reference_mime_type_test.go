package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestInlineBase64ReferenceMimeTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant InlineBase64ReferenceMimeType
		wire    string
	}
	cases := []enumCase{
		{variant: InlineBase64ReferenceMimeTypeImageJpeg, wire: "\"image/jpeg\""},
		{variant: InlineBase64ReferenceMimeTypeImagePng, wire: "\"image/png\""},
		{variant: InlineBase64ReferenceMimeTypeImageWebp, wire: "\"image/webp\""},
		{variant: InlineBase64ReferenceMimeTypeImageHeic, wire: "\"image/heic\""},
		{variant: InlineBase64ReferenceMimeTypeImageHeif, wire: "\"image/heif\""},
		{variant: InlineBase64ReferenceMimeTypeAudioMpeg, wire: "\"audio/mpeg\""},
		{variant: InlineBase64ReferenceMimeTypeAudioWav, wire: "\"audio/wav\""},
		{variant: InlineBase64ReferenceMimeTypeVideoMp4, wire: "\"video/mp4\""},
		{variant: InlineBase64ReferenceMimeTypeVideoQuicktime, wire: "\"video/quicktime\""},
		{variant: InlineBase64ReferenceMimeTypeVideoWebm, wire: "\"video/webm\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded InlineBase64ReferenceMimeType
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
