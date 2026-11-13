package types

import (
	"testing"
)

func TestBodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatchRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_disabled": true,
  "name": "value",
  "retry_enabled": true,
  "request_headers": {
    "key": "value"
  },
  "events": [
    "voice_library_removal_notice"
  ]
}`
	assertRoundTrip(t, wire, new(BodyUpdateWorkspaceWebhookV1WorkspaceWebhooksWebhookIdPatch))
}
