package types

import (
	"testing"
)

func TestEnvironmentVariableAuthConnectionValueRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "auth_connection_id": "value"
}`
	assertRoundTrip(t, wire, new(EnvironmentVariableAuthConnectionValue))
}
