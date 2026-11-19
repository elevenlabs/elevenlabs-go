package agents

import (
	"testing"
)

func TestPostWorkspaceSecretRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "new",
  "name": "value",
  "value": "value"
}`
	assertRoundTrip(t, wire, new(PostWorkspaceSecretRequest))
}

func TestPatchWorkspaceSecretRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "update",
  "name": "value",
  "value": "value"
}`
	assertRoundTrip(t, wire, new(PatchWorkspaceSecretRequest))
}
