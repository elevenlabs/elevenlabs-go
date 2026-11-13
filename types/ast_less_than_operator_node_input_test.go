package types

import (
	"testing"
)

func TestAstLessThanOperatorNodeInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "lt_operator",
  "left": {
    "type": "string_literal",
    "value": "value"
  },
  "right": {
    "type": "string_literal",
    "value": "value"
  }
}`
	assertRoundTrip(t, wire, new(AstLessThanOperatorNodeInput))
}
