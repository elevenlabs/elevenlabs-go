package types

import (
	"testing"
)

func TestVideoReferenceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "generation",
  "generation_id": "value"
}`
	assertRoundTrip(t, wire, new(VideoReference))
}
