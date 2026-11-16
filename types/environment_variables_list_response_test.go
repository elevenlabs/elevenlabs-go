package types

import (
	"testing"
)

func TestEnvironmentVariablesListResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "environment_variables": [
    {
      "label": "value",
      "created_at_unix_secs": 123,
      "updated_at_unix_secs": 123,
      "created_by_user_id": "value",
      "type": "string",
      "id": "value",
      "workspace_id": "value",
      "values": {
        "key": "value"
      }
    }
  ],
  "next_cursor": "value",
  "has_more": true
}`
	assertRoundTrip(t, wire, new(EnvironmentVariablesListResponse))
}
