package types

import (
	"testing"
)

func TestTemplateListResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "templates": [
    {
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
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(TemplateListResponse))
}
