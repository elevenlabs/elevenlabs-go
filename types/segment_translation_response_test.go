package types

import (
	"testing"
)

func TestSegmentTranslationResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version": 123
}`
	assertRoundTrip(t, wire, new(SegmentTranslationResponse))
}
