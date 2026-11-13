package types

import (
	"testing"
)

func TestWorkspaceCreateWebhookResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "webhook_id": "value",
  "webhook_secret": "value"
}`
	assertRoundTrip(t, wire, new(WorkspaceCreateWebhookResponseModel))
}
