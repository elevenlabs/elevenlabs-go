package types

import (
	"testing"
)

func TestTemplateRunListResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "runs": [
    {
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
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(TemplateRunListResponse))
}
