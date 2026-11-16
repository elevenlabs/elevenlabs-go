package types

import (
	"testing"
)

func TestDeleteWorkspaceWebhookResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(DeleteWorkspaceWebhookResponseModel))
}
