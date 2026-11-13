package types

import (
	"testing"
)

func TestListMcpToolsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "success": true,
  "tools": [
    {
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
    }
  ],
  "tool_approval_statuses": [
    {
      "tool_id": "value",
      "state": "up_to_date",
      "approval_policy": "auto_approved",
      "approved_definition": {
        "description": "value",
        "input_schema": {
          "key": {}
        }
      }
    }
  ],
  "error_message": "value"
}`
	assertRoundTrip(t, wire, new(ListMcpToolsResponseModel))
}
