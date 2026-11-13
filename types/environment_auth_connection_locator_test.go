package types

import (
	"testing"
)

func TestEnvironmentAuthConnectionLocatorRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "env_var_label": "value"
}`
	assertRoundTrip(t, wire, new(EnvironmentAuthConnectionLocator))
}
