package types

import (
	"testing"
)

func TestWorkspaceWebhookUsageResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "usage_type": "ConvAI Agent Settings"
}`
	assertRoundTrip(t, wire, new(WorkspaceWebhookUsageResponseModel))
}
