package types

import (
	"testing"
)

func TestSubtitlingPartialSubtitleRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "partial_subtitle",
  "cue_index": 123,
  "lines": [
    "value"
  ],
  "start_ms": 123
}`
	assertRoundTrip(t, wire, new(SubtitlingPartialSubtitle))
}
