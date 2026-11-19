package types

import (
	"fmt"
)

type SafetyRule string

const (
	SafetyRuleSexualMinors               SafetyRule = "sexual_minors"
	SafetyRuleForgetModeration           SafetyRule = "forget_moderation"
	SafetyRuleExtremism                  SafetyRule = "extremism"
	SafetyRuleScamFraud                  SafetyRule = "scam_fraud"
	SafetyRulePolitical                  SafetyRule = "political"
	SafetyRuleSelfHarm                   SafetyRule = "self_harm"
	SafetyRuleIllegalDistributionMedical SafetyRule = "illegal_distribution_medical"
	SafetyRuleSexualAdults               SafetyRule = "sexual_adults"
	SafetyRuleUnknown                    SafetyRule = "unknown"
)

func NewSafetyRuleFromString(s string) (SafetyRule, error) {
	switch s {
	case "sexual_minors":
		return SafetyRuleSexualMinors, nil
	case "forget_moderation":
		return SafetyRuleForgetModeration, nil
	case "extremism":
		return SafetyRuleExtremism, nil
	case "scam_fraud":
		return SafetyRuleScamFraud, nil
	case "political":
		return SafetyRulePolitical, nil
	case "self_harm":
		return SafetyRuleSelfHarm, nil
	case "illegal_distribution_medical":
		return SafetyRuleIllegalDistributionMedical, nil
	case "sexual_adults":
		return SafetyRuleSexualAdults, nil
	case "unknown":
		return SafetyRuleUnknown, nil
	}
	var t SafetyRule
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e SafetyRule) Ptr() *SafetyRule {
	return &e
}
