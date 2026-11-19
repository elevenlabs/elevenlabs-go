package types

import (
	"testing"
)

func TestReadMetadataChapterDbModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "chapter_name": "value",
  "word_count": 123,
  "char_count": 123,
  "starting_char_offset": 123,
  "has_parsed_html": true,
  "has_summary": true,
  "duration_seconds": 9.99,
  "file_number": "value",
  "is_fallback_name": true,
  "chapter_id": "value"
}`
	assertRoundTrip(t, wire, new(ReadMetadataChapterDbModel))
}
