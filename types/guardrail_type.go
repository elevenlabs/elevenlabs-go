package types

import (
	"fmt"
)

type GuardrailType string

const (
	GuardrailTypeCustom                GuardrailType = "custom"
	GuardrailTypePromptInjection       GuardrailType = "prompt_injection"
	GuardrailTypeSelfHarmIntent        GuardrailType = "self_harm_intent"
	GuardrailTypeViolenceGraphic       GuardrailType = "violence_graphic"
	GuardrailTypeSexual                GuardrailType = "sexual"
	GuardrailTypeViolence              GuardrailType = "violence"
	GuardrailTypeHarassment            GuardrailType = "harassment"
	GuardrailTypeSexualMinors          GuardrailType = "sexual_minors"
	GuardrailTypeSelfHarm              GuardrailType = "self_harm"
	GuardrailTypeSelfHarmInstructions  GuardrailType = "self_harm_instructions"
	GuardrailTypeHarassmentThreatening GuardrailType = "harassment_threatening"
	GuardrailTypeHate                  GuardrailType = "hate"
	GuardrailTypeHateThreatening       GuardrailType = "hate_threatening"
	GuardrailTypeProfanity             GuardrailType = "profanity"
	GuardrailTypeReligionOrPolitics    GuardrailType = "religion_or_politics"
	GuardrailTypeMedicalAndLegal       GuardrailType = "medical_and_legal"
	GuardrailTypeGuardrail             GuardrailType = "guardrail"
)

func NewGuardrailTypeFromString(s string) (GuardrailType, error) {
	switch s {
	case "custom":
		return GuardrailTypeCustom, nil
	case "prompt_injection":
		return GuardrailTypePromptInjection, nil
	case "self_harm_intent":
		return GuardrailTypeSelfHarmIntent, nil
	case "violence_graphic":
		return GuardrailTypeViolenceGraphic, nil
	case "sexual":
		return GuardrailTypeSexual, nil
	case "violence":
		return GuardrailTypeViolence, nil
	case "harassment":
		return GuardrailTypeHarassment, nil
	case "sexual_minors":
		return GuardrailTypeSexualMinors, nil
	case "self_harm":
		return GuardrailTypeSelfHarm, nil
	case "self_harm_instructions":
		return GuardrailTypeSelfHarmInstructions, nil
	case "harassment_threatening":
		return GuardrailTypeHarassmentThreatening, nil
	case "hate":
		return GuardrailTypeHate, nil
	case "hate_threatening":
		return GuardrailTypeHateThreatening, nil
	case "profanity":
		return GuardrailTypeProfanity, nil
	case "religion_or_politics":
		return GuardrailTypeReligionOrPolitics, nil
	case "medical_and_legal":
		return GuardrailTypeMedicalAndLegal, nil
	case "guardrail":
		return GuardrailTypeGuardrail, nil
	}
	var t GuardrailType
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (e GuardrailType) Ptr() *GuardrailType {
	return &e
}
