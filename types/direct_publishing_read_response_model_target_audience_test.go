package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestDirectPublishingReadResponseModelTargetAudienceVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant DirectPublishingReadResponseModelTargetAudience
		wire    string
	}
	cases := []enumCase{
		{variant: DirectPublishingReadResponseModelTargetAudienceChildren, wire: "\"children\""},
		{variant: DirectPublishingReadResponseModelTargetAudienceYoungAdult, wire: "\"young adult\""},
		{variant: DirectPublishingReadResponseModelTargetAudienceAdult, wire: "\"adult\""},
		{variant: DirectPublishingReadResponseModelTargetAudienceAllAges, wire: "\"all ages\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded DirectPublishingReadResponseModelTargetAudience
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
