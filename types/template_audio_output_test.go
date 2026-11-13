package types

import (
	"testing"
)

func TestTemplateAudioOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "pending",
  "failure_reason": "timeout",
  "error_message": "value",
  "content_url": "value",
  "content_mime_type": "value",
  "type": "audio"
}`
	assertRoundTrip(t, wire, new(TemplateAudioOutput))
}
