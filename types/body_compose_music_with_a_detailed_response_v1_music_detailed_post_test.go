package types

import (
	"testing"
)

func TestBodyComposeMusicWithADetailedResponseV1MusicDetailedPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "prompt": "value",
  "composition_plan": {
    "positive_global_styles": [
      "value"
    ],
    "negative_global_styles": [
      "value"
    ],
    "sections": [
      {
        "section_name": "value",
        "positive_local_styles": [
          "value"
        ],
        "negative_local_styles": [
          "value"
        ],
        "duration_ms": 3000,
        "lines": [
          "value"
        ],
        "source_from": {
          "song_id": "value",
          "range": {
            "start_ms": 123,
            "end_ms": 123
          },
          "negative_ranges": [
            {
              "start_ms": 123,
              "end_ms": 123
            }
          ]
        }
      }
    ]
  },
  "music_length_ms": 3000,
  "model_id": "music_v1",
  "seed": 123,
  "force_instrumental": true,
  "finetune_id": "value",
  "respect_sections_durations": true,
  "store_for_inpainting": true,
  "with_timestamps": true,
  "with_waveform_visual": true,
  "sign_with_c2pa": true
}`
	assertRoundTrip(t, wire, new(BodyComposeMusicWithADetailedResponseV1MusicDetailedPost))
}
