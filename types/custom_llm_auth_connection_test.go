package types

import (
	"testing"
)

func TestCustomLlmAuthConnectionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "auth_connection_id": "value"
}`
	assertRoundTrip(t, wire, new(CustomLlmAuthConnection))
}

func TestCustomLlmAuthConnectionRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "auth_connection_id": "value"
}`
	assertRoundTrip(t, wire, new(CustomLlmAuthConnection))
}
