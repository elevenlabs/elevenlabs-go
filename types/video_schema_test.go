package types

import (
	"testing"
)

func TestVideoSchemaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "type": "video"
}`
	assertRoundTrip(t, wire, new(VideoSchema))
}
