package types

import (
	"testing"
)

func TestTemplateVersionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "version_id": "value",
  "published_at_unix": 123,
  "is_latest": true,
  "inputs": [
    {
      "id": "value",
      "content_schema": {
        "title": "value",
        "description": "value",
        "type": "string",
        "enum": [
          "value"
        ]
      }
    }
  ],
  "outputs": [
    {
      "id": "value",
      "content_schema": {
        "title": "value",
        "description": "value",
        "type": "string",
        "enum": [
          "value"
        ]
      }
    }
  ]
}`
	assertRoundTrip(t, wire, new(TemplateVersion))
}
