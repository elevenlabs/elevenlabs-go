package types

import (
	"testing"
)

func TestImageGenerationRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhook": {
    "type": "all"
  },
  "prompt": "value",
  "images": [
    {
      "type": "generation",
      "generation_id": "value"
    }
  ],
  "mask": {
    "type": "generation",
    "generation_id": "value"
  },
  "quality": "low",
  "background": "transparent",
  "aspect_ratio": "1:1",
  "model_id": "gpt-image-1"
}`
	assertRoundTrip(t, wire, new(ImageGenerationRequest))
}
