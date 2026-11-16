package types

import (
	"testing"
)

func TestSessionStartedConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "sample_rate": 123,
  "audio_format": "pcm_8000",
  "language_code": "value",
  "secondary_languages": [
    "value"
  ],
  "timestamps_granularity": "none",
  "vad_commit_strategy": true,
  "vad_silence_threshold_secs": 9.99,
  "vad_threshold": 9.99,
  "min_speech_duration_ms": 123,
  "min_silence_duration_ms": 123,
  "max_tokens_to_recompute": 123,
  "model_id": "scribe_v2_realtime",
  "disable_logging": true,
  "include_timestamps": true,
  "include_language_detection": true,
  "filter_background_audio": true,
  "keyterms": [
    "value"
  ],
  "no_verbatim": true,
  "entity_detection": [
    "value"
  ],
  "entity_redaction": [
    "value"
  ],
  "entity_redaction_mode": "value",
  "keepalive_interval_ms": 123
}`
	assertRoundTrip(t, wire, new(SessionStartedConfig))
}
