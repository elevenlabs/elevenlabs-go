package types

import (
	"testing"
)

func TestSafetyCommonModelOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "ivc": {
    "is_unsafe": true,
    "llm_reason": "value",
    "safety_prompt_version": 123,
    "matched_rule_id": [
      "sexual_minors"
    ]
  },
  "non_ivc": {
    "is_unsafe": true,
    "llm_reason": "value",
    "safety_prompt_version": 123,
    "matched_rule_id": [
      "sexual_minors"
    ]
  }
}`
	assertRoundTrip(t, wire, new(SafetyCommonModelOutput))
}
