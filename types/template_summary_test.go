package types

import (
	"testing"
)

func TestTemplateSummaryRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "name": "value",
  "description": "value",
  "versions": [
    {
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
    }
  ],
  "has_more_versions": true
}`
	assertRoundTrip(t, wire, new(TemplateSummary))
}
