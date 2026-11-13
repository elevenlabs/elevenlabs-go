package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGuardrailTypeVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant GuardrailType
		wire    string
	}
	cases := []enumCase{
		{variant: GuardrailTypeCustom, wire: "\"custom\""},
		{variant: GuardrailTypePromptInjection, wire: "\"prompt_injection\""},
		{variant: GuardrailTypeSelfHarmIntent, wire: "\"self_harm_intent\""},
		{variant: GuardrailTypeViolenceGraphic, wire: "\"violence_graphic\""},
		{variant: GuardrailTypeSexual, wire: "\"sexual\""},
		{variant: GuardrailTypeViolence, wire: "\"violence\""},
		{variant: GuardrailTypeHarassment, wire: "\"harassment\""},
		{variant: GuardrailTypeSexualMinors, wire: "\"sexual_minors\""},
		{variant: GuardrailTypeSelfHarm, wire: "\"self_harm\""},
		{variant: GuardrailTypeSelfHarmInstructions, wire: "\"self_harm_instructions\""},
		{variant: GuardrailTypeHarassmentThreatening, wire: "\"harassment_threatening\""},
		{variant: GuardrailTypeHate, wire: "\"hate\""},
		{variant: GuardrailTypeHateThreatening, wire: "\"hate_threatening\""},
		{variant: GuardrailTypeProfanity, wire: "\"profanity\""},
		{variant: GuardrailTypeReligionOrPolitics, wire: "\"religion_or_politics\""},
		{variant: GuardrailTypeMedicalAndLegal, wire: "\"medical_and_legal\""},
		{variant: GuardrailTypeGuardrail, wire: "\"guardrail\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded GuardrailType
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
