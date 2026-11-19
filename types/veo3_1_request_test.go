package types

import (
	"testing"
)

func TestVeo31RequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhook": {
    "type": "all"
  },
  "prompt": "value",
  "negative_prompt": "value",
  "seed": 123,
  "enhance_prompt": true,
  "duration_secs": 123,
  "aspect_ratio": "16:9",
  "resolution": "720p",
  "generate_audio": true,
  "start_frame": {
    "type": "generation",
    "generation_id": "value"
  },
  "end_frame": {
    "type": "generation",
    "generation_id": "value"
  },
  "images": [
    {
      "image": {
        "type": "generation",
        "generation_id": "value"
      },
      "role": "subject"
    }
  ],
  "model_id": "veo-3.1-generate-001"
}`
	assertRoundTrip(t, wire, new(Veo31Request))
}
