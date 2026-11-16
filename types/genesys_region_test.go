package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGenesysRegionVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant GenesysRegion
		wire    string
	}
	cases := []enumCase{
		{variant: GenesysRegionUsEast1, wire: "\"us_east_1\""},
		{variant: GenesysRegionEuWest1, wire: "\"eu_west_1\""},
		{variant: GenesysRegionApSoutheast2, wire: "\"ap_southeast_2\""},
		{variant: GenesysRegionApNortheast1, wire: "\"ap_northeast_1\""},
		{variant: GenesysRegionEuCentral1, wire: "\"eu_central_1\""},
		{variant: GenesysRegionUsWest2, wire: "\"us_west_2\""},
		{variant: GenesysRegionCaCentral1, wire: "\"ca_central_1\""},
		{variant: GenesysRegionApNortheast2, wire: "\"ap_northeast_2\""},
		{variant: GenesysRegionEuWest2, wire: "\"eu_west_2\""},
		{variant: GenesysRegionApSouth1, wire: "\"ap_south_1\""},
		{variant: GenesysRegionUsEast2, wire: "\"us_east_2\""},
		{variant: GenesysRegionSaEast1, wire: "\"sa_east_1\""},
		{variant: GenesysRegionMeCentral1, wire: "\"me_central_1\""},
		{variant: GenesysRegionApNortheast3, wire: "\"ap_northeast_3\""},
		{variant: GenesysRegionEuCentral2, wire: "\"eu_central_2\""},
		{variant: GenesysRegionMxCentral1, wire: "\"mx_central_1\""},
		{variant: GenesysRegionApSoutheast1, wire: "\"ap_southeast_1\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded GenesysRegion
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
