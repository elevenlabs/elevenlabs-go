package types

import (
	"testing"
)

func TestWorkspaceApiKeyListResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
  ]
}`
	assertRoundTrip(t, wire, new(WorkspaceApiKeyListResponseModel))
}
