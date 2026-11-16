package types

import (
	"testing"
)

func TestPrivateKeyJwtResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "auth_type": "private_key_jwt",
  "provider": "value",
  "algorithm": "HS256",
  "key_id": "value",
  "issuer": "value",
  "audience": "value",
  "subject": "value",
  "expiration_seconds": 123,
  "extra_params": {
    "key": "value"
  },
  "extra_headers": {
    "key": "value"
  },
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
  "status_updated_at": "value"
}`
	assertRoundTrip(t, wire, new(PrivateKeyJwtResponse))
}
