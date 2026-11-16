package types

import (
	"testing"
)

func TestBytedanceSeedream5ProRequestRoundTrip(t *testing.T) {
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
  "model_id": "bytedance-seedream-5-pro",
  "resolution": "1K"
}`
	assertRoundTrip(t, wire, new(BytedanceSeedream5ProRequest))
}
