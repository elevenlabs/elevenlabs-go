package types

import (
	"testing"
)

func TestPatchWorkspaceWebhookResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "value"
}`
	assertRoundTrip(t, wire, new(PatchWorkspaceWebhookResponseModel))
}
