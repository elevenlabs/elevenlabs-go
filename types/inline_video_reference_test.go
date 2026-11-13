package types

import (
	"testing"
)

func TestInlineVideoReferenceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "inline_base64",
  "content_base64": "value",
  "mime_type": "video/mp4"
}`
	assertRoundTrip(t, wire, new(InlineVideoReference))
}
