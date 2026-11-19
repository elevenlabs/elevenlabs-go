package types

import (
	"testing"
)

func TestCharacterResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "character_id": "value",
  "name": "value",
  "metadata": {
    "description": "value",
    "sample_message": "value",
    "voice_creation_prompt_suggestion": "value",
    "gender": "male",
    "age": "young",
    "accent": "value",
    "role": "narrator"
  },
  "recommended_voice_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(CharacterResponseModel))
}
