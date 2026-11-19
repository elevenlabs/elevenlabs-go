package types

import (
	"testing"
)

func TestAstEqualsOperatorNodeOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "eq_operator",
  "left": {
    "type": "string_literal",
    "value": "value"
  },
  "right": {
    "type": "string_literal",
    "value": "value"
  }
}`
	assertRoundTrip(t, wire, new(AstEqualsOperatorNodeOutput))
}
