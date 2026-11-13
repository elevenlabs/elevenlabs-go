package music

import (
	"testing"
)

func TestCreateCompositionPlanResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(CreateCompositionPlanResponse))
}

func TestBodyGenerateCompositionPlanV1MusicPlanPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "prompt": "value",
  "music_length_ms": 3000,
  "source_composition_plan": {
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
  "model_id": "music_v1"
}`
	assertRoundTrip(t, wire, new(BodyGenerateCompositionPlanV1MusicPlanPost))
}

func TestBodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlanRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(BodyGenerateCompositionPlanV1MusicPlanPostSourceCompositionPlan))
}
