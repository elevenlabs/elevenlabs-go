package types

import (
	"testing"
)

func TestTemplateArrayOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "pending",
  "failure_reason": "timeout",
  "error_message": "value",
  "type": "array",
  "content": [
    {
      "id": "value",
      "status": "pending",
      "failure_reason": "timeout",
      "error_message": "value",
      "content_url": "value",
      "content_mime_type": "value",
      "type": "image"
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(TemplateArrayOutput))
}
