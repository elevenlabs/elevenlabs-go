package types

import (
	"testing"
)

func TestAstGreaterThanOrEqualsOperatorNodeOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "gte_operator",
  "left": {
    "type": "string_literal",
    "value": "value"
  },
  "right": {
    "type": "string_literal",
    "value": "value"
  }
}`
	assertRoundTrip(t, wire, new(AstGreaterThanOrEqualsOperatorNodeOutput))
}
