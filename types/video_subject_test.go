package types

import (
	"testing"
)

func TestVideoSubjectRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "description": "value"
}`
	assertRoundTrip(t, wire, new(VideoSubject))
}
