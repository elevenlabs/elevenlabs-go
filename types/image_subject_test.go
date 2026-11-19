package types

import (
	"testing"
)

func TestImageSubjectRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(ImageSubject))
}
