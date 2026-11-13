package types

import (
	"testing"
)

func TestWorkspaceWebhookListResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhooks": [
    {
      "name": "value",
      "webhook_id": "value",
      "webhook_url": "value",
      "is_disabled": true,
      "is_auto_disabled": true,
      "created_at_unix": 123,
      "auth_type": "hmac",
      "usage": [
        {
          "usage_type": "ConvAI Agent Settings"
        }
      ],
      "events": [
        "voice_library_removal_notice"
      ],
      "most_recent_failure_error_code": 123,
      "most_recent_failure_timestamp": 123
    }
  ]
}`
	assertRoundTrip(t, wire, new(WorkspaceWebhookListResponseModel))
}
