package types

import (
	"testing"
)

func TestDubbingTranscriptRevisionResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "revision": 123
}`
	assertRoundTrip(t, wire, new(DubbingTranscriptRevisionResponse))
}
