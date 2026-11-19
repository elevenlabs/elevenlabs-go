package types

import (
	"testing"
)

func TestDynamicVariablesConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "dynamic_variable_placeholders": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(DynamicVariablesConfig))
}
