package types

import (
	"testing"
)

func TestBytedanceSeedance2MiniRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhook": {
    "type": "all"
  },
  "prompt": "value",
  "aspect_ratio": "auto",
  "duration_secs": 15,
  "seed": 123,
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
      "type": "generation",
      "generation_id": "value"
    }
  ],
  "videos": [
    {
      "type": "generation",
      "generation_id": "value"
    }
  ],
  "audios": [
    {
      "type": "generation",
      "generation_id": "value"
    }
  ],
  "model_id": "bytedance-seedance-v2-mini",
  "resolution": "480p"
}`
	assertRoundTrip(t, wire, new(BytedanceSeedance2MiniRequest))
}
