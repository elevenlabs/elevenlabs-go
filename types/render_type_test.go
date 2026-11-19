package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestRenderTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant RenderType
		wire    string
	}
	cases := []enumCase{
		{variant: RenderTypeMp4, wire: "\"mp4\""},
		{variant: RenderTypeAac, wire: "\"aac\""},
		{variant: RenderTypeMp3, wire: "\"mp3\""},
		{variant: RenderTypeWav, wire: "\"wav\""},
		{variant: RenderTypeAaf, wire: "\"aaf\""},
		{variant: RenderTypeTracksZip, wire: "\"tracks_zip\""},
		{variant: RenderTypeClipsZip, wire: "\"clips_zip\""},
		{variant: RenderTypeZip, wire: "\"zip\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded RenderType
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
