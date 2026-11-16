package types

import (
	"testing"
)

func TestImageSchemaRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "description": "value",
  "type": "image"
}`
	assertRoundTrip(t, wire, new(ImageSchema))
}
