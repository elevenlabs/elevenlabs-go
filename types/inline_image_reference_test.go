package types

import (
	"testing"
)

func TestInlineImageReferenceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "inline_base64",
  "content_base64": "value",
  "mime_type": "image/jpeg"
}`
	assertRoundTrip(t, wire, new(InlineImageReference))
}
