package types

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCustomGuardrailConfigModelVariantsRoundTrip(t *testing.T) {
	t.Parallel()
	type enumCase struct {
		variant CustomGuardrailConfigModel
		wire    string
	}
	cases := []enumCase{
		{variant: CustomGuardrailConfigModelGemini25FlashLite, wire: "\"gemini-2.5-flash-lite\""},
		{variant: CustomGuardrailConfigModelGemini25Flash, wire: "\"gemini-2.5-flash\""},
		{variant: CustomGuardrailConfigModelGemini31FlashLite, wire: "\"gemini-3.1-flash-lite\""},
		{variant: CustomGuardrailConfigModelGemini35Flash, wire: "\"gemini-3.5-flash\""},
		{variant: CustomGuardrailConfigModelClaudeHaiku45, wire: "\"claude-haiku-4-5\""},
		{variant: CustomGuardrailConfigModelClaudeSonnet46, wire: "\"claude-sonnet-4-6\""},
		{variant: CustomGuardrailConfigModelGpt54Nano, wire: "\"gpt-5.4-nano\""},
		{variant: CustomGuardrailConfigModelGpt54Mini, wire: "\"gpt-5.4-mini\""},
	}
	for _, tc := range cases {
		tc := tc
		t.Run(tc.wire, func(t *testing.T) {
			t.Parallel()
			encoded, err := json.Marshal(tc.variant)
			if err != nil {
				t.Fatalf("marshal: %v", err)
			}
			var decoded CustomGuardrailConfigModel
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
