package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestReceptionVerticalVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ReceptionVertical
		wire    string
	}
	cases := []enumCase{
		{variant: ReceptionVerticalSalonBarbershop, wire: "\"salon_barbershop\""},
		{variant: ReceptionVerticalNailBeauty, wire: "\"nail_beauty\""},
		{variant: ReceptionVerticalSpa, wire: "\"spa\""},
		{variant: ReceptionVerticalMedical, wire: "\"medical\""},
		{variant: ReceptionVerticalVeterinary, wire: "\"veterinary\""},
		{variant: ReceptionVerticalFitness, wire: "\"fitness\""},
		{variant: ReceptionVerticalHomeServices, wire: "\"home_services\""},
		{variant: ReceptionVerticalPhotography, wire: "\"photography\""},
		{variant: ReceptionVerticalConsulting, wire: "\"consulting\""},
		{variant: ReceptionVerticalRestaurant, wire: "\"restaurant\""},
		{variant: ReceptionVerticalDental, wire: "\"dental\""},
		{variant: ReceptionVerticalAutoService, wire: "\"auto_service\""},
		{variant: ReceptionVerticalTutoring, wire: "\"tutoring\""},
		{variant: ReceptionVerticalLegal, wire: "\"legal\""},
		{variant: ReceptionVerticalRealEstate, wire: "\"real_estate\""},
		{variant: ReceptionVerticalOther, wire: "\"other\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ReceptionVertical
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

func TestReceptionVerticalVariantsRoundTrip2(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant ReceptionVertical
		wire    string
	}
	cases := []enumCase{
		{variant: ReceptionVerticalSalonBarbershop, wire: "\"salon_barbershop\""},
		{variant: ReceptionVerticalNailBeauty, wire: "\"nail_beauty\""},
		{variant: ReceptionVerticalSpa, wire: "\"spa\""},
		{variant: ReceptionVerticalMedical, wire: "\"medical\""},
		{variant: ReceptionVerticalVeterinary, wire: "\"veterinary\""},
		{variant: ReceptionVerticalFitness, wire: "\"fitness\""},
		{variant: ReceptionVerticalHomeServices, wire: "\"home_services\""},
		{variant: ReceptionVerticalPhotography, wire: "\"photography\""},
		{variant: ReceptionVerticalConsulting, wire: "\"consulting\""},
		{variant: ReceptionVerticalRestaurant, wire: "\"restaurant\""},
		{variant: ReceptionVerticalDental, wire: "\"dental\""},
		{variant: ReceptionVerticalAutoService, wire: "\"auto_service\""},
		{variant: ReceptionVerticalTutoring, wire: "\"tutoring\""},
		{variant: ReceptionVerticalLegal, wire: "\"legal\""},
		{variant: ReceptionVerticalRealEstate, wire: "\"real_estate\""},
		{variant: ReceptionVerticalOther, wire: "\"other\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded ReceptionVertical
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
