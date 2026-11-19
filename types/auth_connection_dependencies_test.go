package types

import (
	"testing"
)

func TestAuthConnectionDependenciesRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tools": [
    {
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "mcp_servers": [
    {
      "id": "value",
      "name": "value",
      "type": "available",
      "created_at_unix_secs": 123,
      "access_level": "admin"
    }
  ],
  "integration_connections": [
    {
      "id": "value",
      "name": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(AuthConnectionDependencies))
}
