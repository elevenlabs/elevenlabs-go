package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestSafetyRuleVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant SafetyRule
		wire    string
	}
	cases := []enumCase{
		{variant: SafetyRuleSexualMinors, wire: "\"sexual_minors\""},
		{variant: SafetyRuleForgetModeration, wire: "\"forget_moderation\""},
		{variant: SafetyRuleExtremism, wire: "\"extremism\""},
		{variant: SafetyRuleScamFraud, wire: "\"scam_fraud\""},
		{variant: SafetyRulePolitical, wire: "\"political\""},
		{variant: SafetyRuleSelfHarm, wire: "\"self_harm\""},
		{variant: SafetyRuleIllegalDistributionMedical, wire: "\"illegal_distribution_medical\""},
		{variant: SafetyRuleSexualAdults, wire: "\"sexual_adults\""},
		{variant: SafetyRuleUnknown, wire: "\"unknown\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded SafetyRule
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
