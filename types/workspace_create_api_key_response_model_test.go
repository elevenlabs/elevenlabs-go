package types

import (
	"testing"
)

func TestWorkspaceCreateApiKeyResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "xi-api-key": "value",
  "key_id": "value"
}`
	assertRoundTrip(t, wire, new(WorkspaceCreateApiKeyResponseModel))
}
