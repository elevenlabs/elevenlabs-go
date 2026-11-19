package types

import (
	"testing"
)

func TestAstNumberNodeOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "number_literal",
  "value": 9.99
}`
	assertRoundTrip(t, wire, new(AstNumberNodeOutput))
}
