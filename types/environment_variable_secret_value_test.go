package types

import (
	"testing"
)

func TestEnvironmentVariableSecretValueRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "secret_id": "value"
}`
	assertRoundTrip(t, wire, new(EnvironmentVariableSecretValue))
}
