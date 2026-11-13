package types

import (
	"testing"
)

func TestCharacterMetadataResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "description": "value",
  "sample_message": "value",
  "voice_creation_prompt_suggestion": "value",
  "gender": "male",
  "age": "young",
  "accent": "value",
  "role": "narrator"
}`
	assertRoundTrip(t, wire, new(CharacterMetadataResponseModel))
}
