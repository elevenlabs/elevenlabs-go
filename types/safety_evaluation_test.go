package types

import (
	"testing"
)

func TestSafetyEvaluationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_unsafe": true,
  "llm_reason": "value",
  "safety_prompt_version": 123,
  "matched_rule_id": [
    "sexual_minors"
  ]
}`
	assertRoundTrip(t, wire, new(SafetyEvaluation))
}
