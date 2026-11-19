package types

import (
	"testing"
)

func TestInlineAudioReferenceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "inline_base64",
  "content_base64": "value",
  "mime_type": "audio/mpeg"
}`
	assertRoundTrip(t, wire, new(InlineAudioReference))
}
