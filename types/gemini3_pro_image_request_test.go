package types

import (
	"testing"
)

func TestGemini3ProImageRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhook": {
    "type": "all"
  },
  "prompt": "value",
  "model_id": "gemini-3-pro-image",
  "images": [
    {
      "type": "generation",
      "generation_id": "value"
    }
  ],
  "aspect_ratio": "auto",
  "resolution": "1K"
}`
	assertRoundTrip(t, wire, new(Gemini3ProImageRequest))
}
