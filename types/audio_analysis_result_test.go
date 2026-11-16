package types

import (
	"testing"
)

func TestAudioAnalysisResultRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "content_type": "value",
  "overall_pacing": "value",
  "segments": [
    {
      "start_ms": 123,
      "end_ms": 123,
      "description": "value",
      "segment_type": "value",
      "has_speech": true,
      "has_music": true,
      "pacing": "value"
    }
  ],
  "key_moments": [
    {
      "timestamp_ms": 123,
      "type": "value",
      "description": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(AudioAnalysisResult))
}
