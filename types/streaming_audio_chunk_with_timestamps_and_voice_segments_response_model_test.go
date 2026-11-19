package types

import (
	"testing"
)

func TestStreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "audio_base64": "value",
  "alignment": {
    "characters": [
      "value"
    ],
    "character_start_times_seconds": [
      9.99
    ],
    "character_end_times_seconds": [
      9.99
    ]
  },
  "normalized_alignment": {
    "characters": [
      "value"
    ],
    "character_start_times_seconds": [
      9.99
    ],
    "character_end_times_seconds": [
      9.99
    ]
  },
  "voice_segments": [
    {
      "voice_id": "value",
      "start_time_seconds": 9.99,
      "end_time_seconds": 9.99,
      "character_start_index": 123,
      "character_end_index": 123,
      "dialogue_input_index": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(StreamingAudioChunkWithTimestampsAndVoiceSegmentsResponseModel))
}
