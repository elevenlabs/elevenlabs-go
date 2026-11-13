package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestColumnUnitVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ColumnUnit
		wire    string
	}
	cases := []enumCase{
		{variant: ColumnUnitMs, wire: "\"ms\""},
		{variant: ColumnUnitS, wire: "\"s\""},
		{variant: ColumnUnitMin, wire: "\"min\""},
		{variant: ColumnUnitDuration, wire: "\"duration\""},
		{variant: ColumnUnitCredits, wire: "\"credits\""},
		{variant: ColumnUnitUsd, wire: "\"usd\""},
		{variant: ColumnUnitEur, wire: "\"eur\""},
		{variant: ColumnUnitInr, wire: "\"inr\""},
		{variant: ColumnUnitPln, wire: "\"pln\""},
		{variant: ColumnUnitGbp, wire: "\"gbp\""},
		{variant: ColumnUnitRatio, wire: "\"ratio\""},
		{variant: ColumnUnitRating, wire: "\"rating\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ColumnUnit
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
