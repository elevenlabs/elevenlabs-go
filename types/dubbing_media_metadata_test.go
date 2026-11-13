package types

import (
	"testing"
)

func TestDubbingMediaMetadataRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "content_type": "value",
  "duration": 9.99
}`
	assertRoundTrip(t, wire, new(DubbingMediaMetadata))
}
