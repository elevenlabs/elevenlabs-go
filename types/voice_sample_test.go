package types

import (
	"testing"
)

func TestVoiceSampleRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "sample_id": "value",
  "file_name": "value",
  "mime_type": "value",
  "size_bytes": 123,
  "hash": "value",
  "duration_secs": 9.99,
  "remove_background_noise": true,
  "has_isolated_audio": true,
  "has_isolated_audio_preview": true,
  "speaker_separation": {
    "voice_id": "value",
    "sample_id": "value",
    "status": "not_started",
    "speakers": {
      "key": {
        "speaker_id": "value",
        "duration_secs": 9.99,
        "utterances": [
          {
            "start": 9.99,
            "end": 9.99
          }
        ]
      }
    },
    "selected_speaker_ids": [
      "value"
    ]
  },
  "trim_start": 123,
  "trim_end": 123
}`
	assertRoundTrip(t, wire, new(VoiceSample))
}

func TestVoiceSampleRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "sample_id": "value",
  "file_name": "value",
  "mime_type": "value",
  "size_bytes": 123,
  "hash": "value",
  "duration_secs": 9.99,
  "remove_background_noise": true,
  "has_isolated_audio": true,
  "has_isolated_audio_preview": true,
  "speaker_separation": {
    "voice_id": "value",
    "sample_id": "value",
    "status": "not_started",
    "speakers": {
      "key": {
        "speaker_id": "value",
        "duration_secs": 9.99,
        "utterances": [
          {
            "start": 9.99,
            "end": 9.99
          }
        ]
      }
    },
    "selected_speaker_ids": [
      "value"
    ]
  },
  "trim_start": 123,
  "trim_end": 123
}`
	assertRoundTrip(t, wire, new(VoiceSample))
}
