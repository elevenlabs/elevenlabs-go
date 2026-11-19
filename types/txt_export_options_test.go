package types

import (
	"testing"
)

func TestTxtExportOptionsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "max_characters_per_line": 123,
  "include_speakers": true,
  "include_timestamps": true,
  "format": "txt",
  "segment_on_silence_longer_than_s": 9.99,
  "max_segment_duration_s": 9.99,
  "max_segment_chars": 123
}`
	assertRoundTrip(t, wire, new(TxtExportOptions))
}
