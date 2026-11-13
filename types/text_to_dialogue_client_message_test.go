package types

import (
	"testing"
)

func TestTextToDialogueClientMessageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "inputs": [
    {
      "text": "value",
      "voice_id": "value",
      "new_turn": true
    }
  ],
  "flush": true,
  "close_socket": true,
  "keep_alive": true,
  "xi_api_key": "value",
  "authorization": "value",
  "single_use_token": "value",
  "voices": [
    "value"
  ],
  "voice_settings": {
    "stability": 0.5
  },
  "pronunciation_dictionary_locators": [
    {
      "pronunciation_dictionary_id": "value",
      "version_id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(TextToDialogueClientMessage))
}
