package types

import (
	"testing"
)

func TestEnvironmentVariableResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
}`
	assertRoundTrip(t, wire, new(EnvironmentVariableResponse))
}
