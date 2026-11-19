package types

import (
	"testing"
)

func TestTemplateRunResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "template_id": "value",
  "version_id": "value",
  "status": "pending",
  "outputs": {
    "key": {
      "id": "value",
      "status": "pending",
      "failure_reason": "timeout",
      "error_message": "value",
      "content_url": "value",
      "content_mime_type": "value",
      "type": "image"
    }
  }
}`
	assertRoundTrip(t, wire, new(TemplateRunResponse))
}
