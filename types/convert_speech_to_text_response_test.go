package types

import (
	"testing"
)

func TestConvertSpeechToTextResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "language_code": "value",
  "language_probability": 9.99,
  "text": "value",
  "words": [
    {
      "text": "value",
      "start": 9.99,
      "end": 9.99,
      "type": "word",
      "speaker_id": "value",
      "logprob": 9.99,
      "characters": [
        {
          "text": "value",
          "start": 9.99,
          "end": 9.99
        }
      ],
      "channel_index": 123
    }
  ],
  "channel_index": 123,
  "additional_formats": [
    {
      "requested_format": "value",
      "file_extension": "value",
      "content_type": "value",
      "is_base64_encoded": true,
      "content": "value"
    }
  ],
  "transcription_id": "value",
  "entities": [
    {
      "text": "value",
      "entity_type": "value",
      "start_char": 123,
      "end_char": 123
    }
  ],
  "audio_duration_secs": 9.99,
  "edited_transcript": {
    "kind": "transcript",
    "text": "value",
    "edited_text": "value"
  }
}`
	assertRoundTrip(t, wire, new(ConvertSpeechToTextResponse))
}
