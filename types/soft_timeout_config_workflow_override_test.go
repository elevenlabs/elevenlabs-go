package types

import (
	"testing"
)

func TestSoftTimeoutConfigWorkflowOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "timeout_seconds": 9.99,
  "message": "value",
  "additional_soft_timeout_messages": [
    "value"
  ],
  "use_llm_generated_message": true,
  "randomize_fillers": true,
  "max_soft_timeouts_per_generation": 123,
  "llm_generated_message_prompt_override": "value",
  "disable_until_first_user_message": true
}`
	assertRoundTrip(t, wire, new(SoftTimeoutConfigWorkflowOverride))
}
