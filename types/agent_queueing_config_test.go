package types

import (
	"testing"
)

func TestAgentQueueingConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "wait_timeout_seconds": 123,
  "hold_audio": {
    "audio_path": "value",
    "audio_url": "value",
    "original_filename": "value",
    "duration_secs": 9.99,
    "size_bytes": 123
  }
}`
	assertRoundTrip(t, wire, new(AgentQueueingConfig))
}
