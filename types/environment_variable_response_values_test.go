package types

import (
	"testing"
)

func TestEnvironmentVariableResponseValuesRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "key": "value"
}`
	assertRoundTrip(t, wire, new(EnvironmentVariableResponseValues))
}
