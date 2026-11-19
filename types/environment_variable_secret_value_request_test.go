package types

import (
	"testing"
)

func TestEnvironmentVariableSecretValueRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "secret_id": "value"
}`
	assertRoundTrip(t, wire, new(EnvironmentVariableSecretValueRequest))
}
