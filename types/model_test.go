package types

import (
	"testing"
)

func TestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "model_id": "value",
  "name": "value",
  "can_be_finetuned": true,
  "can_do_text_to_speech": true,
  "can_do_voice_conversion": true,
  "can_use_style": true,
  "can_use_speaker_boost": true,
  "serves_pro_voices": true,
  "token_cost_factor": 9.99,
  "description": "value",
  "requires_alpha_access": true,
  "max_characters_request_free_user": 123,
  "max_characters_request_subscribed_user": 123,
  "maximum_text_length_per_request": 123,
  "languages": [
    {
      "language_id": "value",
      "name": "value"
    }
  ],
  "model_rates": {
    "character_cost_multiplier": 9.99,
    "cost_discount_multiplier": 9.99
  },
  "concurrency_group": "value"
}`
	assertRoundTrip(t, wire, new(Model))
}
