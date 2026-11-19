package types

import (
	"testing"
)

func TestBodyCreatePodcastV1StudioPodcastsPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "model_id": "value",
  "mode": {
    "type": "conversation",
    "conversation": {
      "host_voice_id": "value",
      "guest_voice_id": "value"
    }
  },
  "source": {
    "type": "text",
    "text": "value"
  },
  "quality_preset": "standard",
  "duration_scale": "short",
  "language": "va",
  "intro": "value",
  "outro": "value",
  "instructions_prompt": "value",
  "highlights": [
    "valuexxxxx"
  ],
  "callback_url": "value",
  "apply_text_normalization": "auto"
}`
	assertRoundTrip(t, wire, new(BodyCreatePodcastV1StudioPodcastsPost))
}
