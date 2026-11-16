package types

import (
	"testing"
)

func TestMcpServerConfigOutputSecretTokenRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "secret_id": "value"
}`
	assertRoundTrip(t, wire, new(McpServerConfigOutputSecretToken))
}
