package types

import (
	"testing"
)

func TestTemplatePortRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "content_schema": {
    "title": "value",
    "description": "value",
    "type": "string",
    "enum": [
      "value"
    ]
  }
}`
	assertRoundTrip(t, wire, new(TemplatePort))
}
