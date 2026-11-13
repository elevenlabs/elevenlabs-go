package types

import (
	"testing"
)

func TestTurnConfigWorkflowOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "turn_timeout": 9.99,
  "initial_wait_time": 9.99,
  "silence_end_call_timeout": 9.99,
  "turn_eagerness": "patient",
  "spelling_patience": "auto",
  "speculative_turn": true,
  "retranscribe_on_turn_timeout": true,
  "turn_model": "turn_v2",
  "interruption_ignore_terms": [
    "value"
  ],
  "interruption_ignore_term_languages": [
    "value"
  ],
  "merge_with_default_ignore_terms": true,
  "transcribe_on_disabled_interruptions": true,
  "soft_timeout_config": {
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
  },
  "noisy_environment_message": "value"
}`
	assertRoundTrip(t, wire, new(TurnConfigWorkflowOverride))
}
