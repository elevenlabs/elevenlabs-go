package types

import (
	"testing"
)

func TestListAuthConnectionsResponseAuthConnectionsItemRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "auth_type": "oauth2_client_credentials",
  "provider": "value",
  "client_id": "value",
  "token_url": "value",
  "scopes": [
    "value"
  ],
  "extra_params": {
    "key": "value"
  },
  "basic_auth_in_header": true,
  "id": "value",
  "used_by": {
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
  },
  "status": "active",
  "status_detail": "value",
  "status_updated_at": "value",
  "custom_headers": {
    "key": "value"
  }
}`
	assertRoundTrip(t, wire, new(ListAuthConnectionsResponseAuthConnectionsItem))
}
