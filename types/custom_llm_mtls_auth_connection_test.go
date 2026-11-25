package types

import (
	"testing"
)

func TestCustomLlmMtlsAuthConnectionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "auth_connection_id": "value"
}`
	assertRoundTrip(t, wire, new(CustomLlmMtlsAuthConnection))
}
