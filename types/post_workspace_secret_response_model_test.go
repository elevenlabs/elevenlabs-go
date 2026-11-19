package types

import (
	"testing"
)

func TestPostWorkspaceSecretResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "stored",
  "secret_id": "value",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(PostWorkspaceSecretResponseModel))
}
