package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestMetricTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant MetricType
		wire    string
	}
	cases := []enumCase{
		{variant: MetricTypeCredits, wire: "\"credits\""},
		{variant: MetricTypeTtsCharacters, wire: "\"tts_characters\""},
		{variant: MetricTypeMinutesUsed, wire: "\"minutes_used\""},
		{variant: MetricTypeRequestCount, wire: "\"request_count\""},
		{variant: MetricTypeTtfbAvg, wire: "\"ttfb_avg\""},
		{variant: MetricTypeTtfbP95, wire: "\"ttfb_p95\""},
		{variant: MetricTypeFiatUnitsSpent, wire: "\"fiat_units_spent\""},
		{variant: MetricTypeConcurrency, wire: "\"concurrency\""},
		{variant: MetricTypeConcurrencyAverage, wire: "\"concurrency_average\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded MetricType
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
