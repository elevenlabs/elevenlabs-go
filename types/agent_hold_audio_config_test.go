package types

import (
	"testing"
)

func TestAgentHoldAudioConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "audio_path": "value",
  "audio_url": "value",
  "original_filename": "value",
  "duration_secs": 9.99,
  "size_bytes": 123
}`
	assertRoundTrip(t, wire, new(AgentHoldAudioConfig))
}

func TestAgentHoldAudioConfigRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "audio_path": "value",
  "audio_url": "value",
  "original_filename": "value",
  "duration_secs": 9.99,
  "size_bytes": 123
}`
	assertRoundTrip(t, wire, new(AgentHoldAudioConfig))
}
