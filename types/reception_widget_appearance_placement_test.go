package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestReceptionWidgetAppearancePlacementVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ReceptionWidgetAppearancePlacement
		wire    string
	}
	cases := []enumCase{
		{variant: ReceptionWidgetAppearancePlacementTopLeft, wire: "\"top-left\""},
		{variant: ReceptionWidgetAppearancePlacementTop, wire: "\"top\""},
		{variant: ReceptionWidgetAppearancePlacementTopRight, wire: "\"top-right\""},
		{variant: ReceptionWidgetAppearancePlacementBottomLeft, wire: "\"bottom-left\""},
		{variant: ReceptionWidgetAppearancePlacementBottom, wire: "\"bottom\""},
		{variant: ReceptionWidgetAppearancePlacementBottomRight, wire: "\"bottom-right\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ReceptionWidgetAppearancePlacement
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
