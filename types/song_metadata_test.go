package types

import (
	"testing"
)

func TestSongMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "genres": [
    "value"
  ],
  "languages": [
    "value"
  ],
  "is_explicit": true
}`
	assertRoundTrip(t, wire, new(SongMetadata))
}
