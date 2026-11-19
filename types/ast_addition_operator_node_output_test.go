package types

import (
	"testing"
)

func TestAstAdditionOperatorNodeOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "add_operator",
  "left": {
    "type": "string_literal",
    "value": "value"
  },
  "right": {
    "type": "string_literal",
    "value": "value"
  }
}`
	assertRoundTrip(t, wire, new(AstAdditionOperatorNodeOutput))
}
