package types

import (
	"testing"
)

func TestHtmlExportOptionsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "include_speakers": true,
  "include_timestamps": true,
  "format": "html",
  "segment_on_silence_longer_than_s": 9.99,
  "max_segment_duration_s": 9.99,
  "max_segment_chars": 123
}`
	assertRoundTrip(t, wire, new(HtmlExportOptions))
}
