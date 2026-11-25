package types

import (
	"testing"
)

func TestSubtitlingCommittedSubtitleRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "message_type": "committed_subtitle",
  "cue_index": 123,
  "lines": [
    "value"
  ],
  "start_ms": 123,
  "end_ms": 123
}`
	assertRoundTrip(t, wire, new(SubtitlingCommittedSubtitle))
}
