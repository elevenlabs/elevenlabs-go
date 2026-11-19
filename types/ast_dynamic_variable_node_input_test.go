package types

import (
	"testing"
)

func TestAstDynamicVariableNodeInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "dynamic_variable",
  "name": "value"
}`
	assertRoundTrip(t, wire, new(AstDynamicVariableNodeInput))
}
