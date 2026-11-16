package types

import (
	"testing"
)

func TestBaseTurnConfigRoundTrip(t *testing.T) {
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
  "transcribe_on_disabled_interruptions": true
}`
	assertRoundTrip(t, wire, new(BaseTurnConfig))
}
