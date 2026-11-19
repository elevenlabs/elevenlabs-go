package types

import (
	"testing"
)

func TestBytedanceSeedance25RequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhook": {
    "type": "all"
  },
  "model_id": "bytedance-seedance-v2.5",
  "prompt": "value",
  "aspect_ratio": "auto",
  "resolution": "480p",
  "duration_secs": 30,
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
  ]
}`
	assertRoundTrip(t, wire, new(BytedanceSeedance25Request))
}
