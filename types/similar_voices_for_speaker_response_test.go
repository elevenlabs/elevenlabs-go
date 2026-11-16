package types

import (
	"testing"
)

func TestSimilarVoicesForSpeakerResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "voices": [
    {
      "voice_id": "value",
      "name": "value",
      "category": "premade",
      "description": "value",
      "preview_url": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(SimilarVoicesForSpeakerResponse))
}
