package types

import (
	"testing"
)

func TestAstDivisionOperatorNodeOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "div_operator",
  "left": {
    "type": "string_literal",
    "value": "value"
  },
  "right": {
    "type": "string_literal",
    "value": "value"
  }
}`
	assertRoundTrip(t, wire, new(AstDivisionOperatorNodeOutput))
}
