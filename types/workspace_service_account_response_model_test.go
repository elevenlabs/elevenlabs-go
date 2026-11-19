package types

import (
	"testing"
)

func TestWorkspaceServiceAccountResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "service_account_user_id": "value",
  "name": "value",
  "created_at_unix": 123,
  "api-keys": [
    {
      "name": "value",
      "hint": "value",
      "key_id": "value",
      "service_account_user_id": "value",
      "created_at_unix": 123,
      "is_disabled": true,
      "permissions": [
        "text_to_speech"
      ],
      "disable_reason": "trial_ended",
      "character_limit": 123,
      "character_count": 123,
      "hashed_xi_api_key": "value",
      "allowed_ips": [
        "value"
      ],
      "third_party_disable_allowed": true,
      "platform_limits": {
        "credits": {
          "limit": 123,
          "usage": 123
        },
        "pvc": {
          "limit": 123
        },
        "concurrency": {
          "limit": 123
        },
        "dubbing_concurrency": {
          "limit": 123
        },
        "music_concurrency": {
          "limit": 123
        }
      },
      "expires_at_unix": 123
    }
  ],
  "default_sharing_groups": [
    {
      "group": {
        "name": "value",
        "id": "value",
        "members": [
          "value"
        ],
        "permissions": [
          "text_to_speech"
        ],
        "group_usage_limit": 123,
        "group_pvc_limit": 123,
        "group_concurrency_limit": 123,
        "group_dubbing_concurrency_limit": 123,
        "group_music_concurrency_limit": 123,
        "character_count": 123,
        "is_scim_synced": true,
        "scim_group": {
          "scim_external_id": "value",
          "display_name": "value",
          "created_at_unix": 123,
          "updated_at_unix": 123,
          "seat_type": "workspace_admin"
        },
        "scim_frozen": true
      },
      "permission_level": "admin"
    }
  ]
}`
	assertRoundTrip(t, wire, new(WorkspaceServiceAccountResponseModel))
}
