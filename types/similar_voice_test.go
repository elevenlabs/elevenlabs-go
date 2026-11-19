package types

import (
	"testing"
)

func TestSimilarVoiceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voice_id": "value",
  "name": "value",
  "category": "premade",
  "description": "value",
  "preview_url": "value"
}`
	assertRoundTrip(t, wire, new(SimilarVoice))
}
