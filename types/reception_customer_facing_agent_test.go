package types

import (
	"testing"
)

func TestReceptionCustomerFacingAgentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
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
    "disable_first_message_interruptions": true,
    "enable_end_call_tool": true,
    "enable_skip_turn_tool": true,
    "enable_play_keypad_touch_tone_tool": true,
    "agent_rules": [
      {
        "title": "value",
        "description": "value",
        "priority": "strict",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "transfer_rules": [
      {
        "label": "value",
        "phone_number": "value",
        "condition": "value",
        "post_dial_digits": "0",
        "id": "value",
        "created_at_unix": 123
      }
    ],
    "hide_caller_number_on_transfers": true,
    "record_after_transfers": true,
    "procedures": [
      {
        "name": "value",
        "trigger": "value",
        "body": "value",
        "id": "value",
        "created_at_unix": 123,
        "source_key": "value"
      }
    ],
    "enable_secure_mode": true,
    "enable_spam_protection": true
  },
  "deactivated_at_unix": 123,
  "deactivation_reason": "tier_downgrade"
}`
	assertRoundTrip(t, wire, new(ReceptionCustomerFacingAgent))
}
