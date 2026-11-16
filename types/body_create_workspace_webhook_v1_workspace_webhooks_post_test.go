package types

import (
	"testing"
)

func TestBodyCreateWorkspaceWebhookV1WorkspaceWebhooksPostRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "settings": {
    "auth_type": "hmac",
    "name": "value",
    "webhook_url": "value",
    "request_headers": {
      "key": "value"
    },
    "log_redacted_fields": [
      "value"
    ]
  }
}`
	assertRoundTrip(t, wire, new(BodyCreateWorkspaceWebhookV1WorkspaceWebhooksPost))
}
