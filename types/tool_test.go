package types

import (
	"testing"
)

func TestToolRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "title": "value",
  "description": "value",
  "inputSchema": {
    "key": {}
  },
  "execution": {
    "taskSupport": "forbidden"
  },
  "outputSchema": {
    "key": {}
  },
  "icons": [
    {
      "src": "value",
      "mimeType": "value",
      "sizes": [
        "value"
      ],
      "theme": "light"
    }
  ],
  "annotations": {
    "title": "value",
    "readOnlyHint": true,
    "destructiveHint": true,
    "idempotentHint": true,
    "openWorldHint": true
  },
  "_meta": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(Tool))
}
