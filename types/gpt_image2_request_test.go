package types

import (
	"testing"
)

func TestGptImage2RequestRoundTrip(t *testing.T) {
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
  "model_id": "gpt-image-2",
  "aspect_ratio": "auto",
  "resolution": "1K"
}`
	assertRoundTrip(t, wire, new(GptImage2Request))
}
