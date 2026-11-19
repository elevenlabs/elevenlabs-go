package types

import (
	"testing"
)

func TestSpeakerSeparationResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(SpeakerSeparationResponseModel))
}

func TestSpeakerSeparationResponseModelRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(SpeakerSeparationResponseModel))
}
