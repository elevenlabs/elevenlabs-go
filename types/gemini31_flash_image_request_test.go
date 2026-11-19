package types

import (
	"testing"
)

func TestGemini31FlashImageRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhook": {
    "type": "all"
  },
  "prompt": "value",
  "model_id": "gemini-3.1-flash-image",
  "images": [
    {
      "type": "generation",
      "generation_id": "value"
    }
  ],
  "aspect_ratio": "auto",
  "resolution": "512"
}`
	assertRoundTrip(t, wire, new(Gemini31FlashImageRequest))
}
