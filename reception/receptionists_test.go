package reception

import (
	"testing"
)

func TestReceptionUpdateCustomerFacingAgentRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "config": {
    "additional_instructions": "value",
    "tone_traits": [
      "value"
    ],
    "details_to_collect": [
      "value"
    ],
    "llm": "gpt-4o-mini",
    "reasoning_effort": "none",
    "disable_rag": true,
    "background_voice_detection": true,
    "turn_eagerness": "patient",
    "voice_id": "value",
    "language": "value",
    "additional_languages": [
      "value"
    ],
    "first_message": "value",
    "first_message_mode": "interruptible",
    "enable_end_call_tool": true,
    "enable_skip_turn_tool": true,
    "enable_play_keypad_touch_tone_tool": true,
    "hide_caller_number_on_transfers": true,
    "record_after_transfers": true,
    "enable_secure_mode": true,
    "enable_spam_protection": true
  },
  "name": "value"
}`
	assertRoundTrip(t, wire, new(ReceptionUpdateCustomerFacingAgentRequest))
}
