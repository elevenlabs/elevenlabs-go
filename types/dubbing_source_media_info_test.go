package types

import (
	"testing"
)

func TestDubbingSourceMediaInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "filename": "value",
  "duration_s": 9.99,
  "has_video": true,
  "mime_type": "value"
}`
	assertRoundTrip(t, wire, new(DubbingSourceMediaInfo))
}
