package types

import (
	"testing"
)

func TestGemini25FlashImageRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhook": {
    "type": "all"
  },
  "prompt": "value",
  "model_id": "gemini-2.5-flash-image",
  "images": [
    {
      "type": "generation",
      "generation_id": "value"
    }
  ],
  "aspect_ratio": "auto"
}`
	assertRoundTrip(t, wire, new(Gemini25FlashImageRequest))
}
