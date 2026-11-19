package types

import (
	"testing"
)

func TestDubbingTranslationRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "translation",
  "text": "value",
  "segment_id": "value",
  "start_s": 9.99,
  "end_s": 9.99,
  "speaker_id": "value"
}`
	assertRoundTrip(t, wire, new(DubbingTranslation))
}
