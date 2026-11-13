package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestPendingSubscriptionSwitchResponseModelNextTierVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant PendingSubscriptionSwitchResponseModelNextTier
		wire    string
	}
	cases := []enumCase{
		{variant: PendingSubscriptionSwitchResponseModelNextTierFree, wire: "\"free\""},
		{variant: PendingSubscriptionSwitchResponseModelNextTierStarter, wire: "\"starter\""},
		{variant: PendingSubscriptionSwitchResponseModelNextTierGo, wire: "\"go\""},
		{variant: PendingSubscriptionSwitchResponseModelNextTierCreator, wire: "\"creator\""},
		{variant: PendingSubscriptionSwitchResponseModelNextTierPro, wire: "\"pro\""},
		{variant: PendingSubscriptionSwitchResponseModelNextTierGrowingBusiness, wire: "\"growing_business\""},
		{variant: PendingSubscriptionSwitchResponseModelNextTierScale20240810, wire: "\"scale_2024_08_10\""},
		{variant: PendingSubscriptionSwitchResponseModelNextTierGrantTier120250723, wire: "\"grant_tier_1_2025_07_23\""},
		{variant: PendingSubscriptionSwitchResponseModelNextTierGrantTier220250723, wire: "\"grant_tier_2_2025_07_23\""},
		{variant: PendingSubscriptionSwitchResponseModelNextTierGrant, wire: "\"grant\""},
		{variant: PendingSubscriptionSwitchResponseModelNextTierTrial, wire: "\"trial\""},
		{variant: PendingSubscriptionSwitchResponseModelNextTierEnterprise, wire: "\"enterprise\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded PendingSubscriptionSwitchResponseModelNextTier
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
