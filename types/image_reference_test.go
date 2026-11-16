package types

import (
	"testing"
)

func TestImageReferenceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "generation",
  "generation_id": "value"
}`
	assertRoundTrip(t, wire, new(ImageReference))
}

func TestImageReferenceRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "generation",
  "generation_id": "value"
}`
	assertRoundTrip(t, wire, new(ImageReference))
}
