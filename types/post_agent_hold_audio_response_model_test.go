package types

import (
	"testing"
)

func TestPostAgentHoldAudioResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "agent_id": "value",
  "hold_audio": {
    "audio_path": "value",
    "audio_url": "value",
    "original_filename": "value",
    "duration_secs": 9.99,
    "size_bytes": 123
  }
}`
	assertRoundTrip(t, wire, new(PostAgentHoldAudioResponseModel))
}
