package types

import (
	"testing"
)

func TestSearchHighlightSegmentRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "value": "value",
  "is_hit": true
}`
	assertRoundTrip(t, wire, new(SearchHighlightSegment))
}

func TestSearchHighlightSegmentRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "value": "value",
  "is_hit": true
}`
	assertRoundTrip(t, wire, new(SearchHighlightSegment))
}
