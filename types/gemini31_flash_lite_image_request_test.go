package types

import (
	"testing"
)

func TestGemini31FlashLiteImageRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhook": {
    "type": "all"
  },
  "prompt": "value",
  "model_id": "gemini-3.1-flash-lite-image",
  "images": [
    {
      "type": "generation",
      "generation_id": "value"
    }
  ],
  "aspect_ratio": "auto",
  "resolution": "1K"
}`
	assertRoundTrip(t, wire, new(Gemini31FlashLiteImageRequest))
}
