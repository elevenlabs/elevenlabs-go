package types

import (
	"testing"
)

func TestApiIntegrationOAuth2AuthCodeResponseRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "auth_type": "api_integration_oauth2_auth_code",
  "provider": "value",
  "token_url": "value",
  "scopes": [
    "value"
  ],
  "scope_separator": " ",
  "expires_at": "value",
  "integration_id": "value",
  "credential_id": "value",
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
	assertRoundTrip(t, wire, new(ApiIntegrationOAuth2AuthCodeResponse))
}
