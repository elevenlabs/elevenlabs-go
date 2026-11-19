package types

import (
	"testing"
)

func TestBodySpeechToTextV1SpeechToTextPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "model_id": "value",
  "language_code": "value",
  "secondary_language_codes": [
    "value"
  ],
  "numbers_format": "digits",
  "transcript_edit": "value",
  "tag_audio_events": true,
  "num_speakers": 32,
  "timestamps_granularity": "none",
  "diarize": true,
  "diarization_threshold": 0.4,
  "additional_formats": [
    {
      "include_speakers": true,
      "include_timestamps": true,
      "format": "segmented_json",
      "segment_on_silence_longer_than_s": 9.99,
      "max_segment_duration_s": 9.99,
      "max_segment_chars": 123
    }
  ],
  "file_format": "pcm_s16le_16",
  "source_url": "value",
  "webhook": true,
  "webhook_id": "value",
  "temperature": 1.5,
  "seed": 123,
  "use_multi_channel": true,
  "multichannel_output_style": "separate",
  "webhook_metadata": "value",
  "entity_detection": "value",
  "no_verbatim": true,
  "use_speaker_library": true,
  "detect_speaker_roles": true,
  "entity_redaction": "value",
  "entity_redaction_mode": "value",
  "keyterms": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(BodySpeechToTextV1SpeechToTextPost))
}
