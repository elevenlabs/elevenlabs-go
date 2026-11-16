package types

import (
	"testing"
)

func TestAstConditionalOperatorNodeOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "conditional_operator",
  "condition": {
    "type": "string_literal",
    "value": "value"
  },
  "trueExpression": {
    "type": "string_literal",
    "value": "value"
  },
  "falseExpression": {
    "type": "string_literal",
    "value": "value"
  }
}`
	assertRoundTrip(t, wire, new(AstConditionalOperatorNodeOutput))
}
