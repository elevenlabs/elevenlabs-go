package types

import (
	"testing"
)

func TestConvAiWorkspaceStoredSecretConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "stored",
  "secret_id": "value",
  "name": "value",
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
    "tools_has_more": true,
    "agents": [
      {
        "referenced_resource_ids": [
          "value"
        ],
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "agents_has_more": true,
    "phone_numbers": [
      {
        "phone_number_id": "value",
        "phone_number": "value",
        "label": "value",
        "provider": "twilio"
      }
    ],
    "phone_numbers_has_more": true,
    "mcp_servers": [
      {
        "id": "value",
        "name": "value",
        "type": "available",
        "created_at_unix_secs": 123,
        "access_level": "admin"
      }
    ],
    "others": [
      "conversation_initiation_webhook"
    ]
  }
}`
	assertRoundTrip(t, wire, new(ConvAiWorkspaceStoredSecretConfig))
}
