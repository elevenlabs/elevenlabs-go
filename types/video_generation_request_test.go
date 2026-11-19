package types

import (
	"testing"
)

func TestVideoGenerationRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhook": {
    "type": "all"
  },
  "model_id": "creatify-aurora",
  "image": {
    "type": "generation",
    "generation_id": "value"
  },
  "audio": {
    "type": "generation",
    "generation_id": "value"
  },
  "resolution": "480p",
  "guidance_scale": 4.5,
  "audio_guidance_scale": 4.5
}`
	assertRoundTrip(t, wire, new(VideoGenerationRequest))
}
