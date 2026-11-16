package types

import (
	"testing"
)

func TestBytedanceSeedream5LiteRequestRoundTrip(t *testing.T) {
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
  "aspect_ratio": "auto",
  "seed": 123,
  "model_id": "bytedance-seedream-5-lite",
  "resolution": "2K"
}`
	assertRoundTrip(t, wire, new(BytedanceSeedream5LiteRequest))
}
