package types

import (
	"testing"
)

func TestTranscriptionOrderItemRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "kind": "transcription",
  "media_ids": [
    "prodmedia_aaaaaaaaaaaaaaaaaaaaaaaaaa"
  ],
  "source_language": "value",
  "verbatim": true,
  "instructions": "value"
}`
	assertRoundTrip(t, wire, new(TranscriptionOrderItemRequest))
}

func TestTranscriptionOrderItemRequestRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "kind": "transcription",
  "media_ids": [
    "prodmedia_aaaaaaaaaaaaaaaaaaaaaaaaaa"
  ],
  "source_language": "value",
  "verbatim": true,
  "instructions": "value"
}`
	assertRoundTrip(t, wire, new(TranscriptionOrderItemRequest))
}
