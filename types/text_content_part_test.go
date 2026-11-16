package types

import (
	"testing"
)

func TestTextContentPartRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "text",
  "text": "value"
}`
	assertRoundTrip(t, wire, new(TextContentPart))
}
