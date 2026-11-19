package types

import (
	"testing"
)

func TestAstBooleanNodeOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "boolean_literal",
  "value": true
}`
	assertRoundTrip(t, wire, new(AstBooleanNodeOutput))
}
