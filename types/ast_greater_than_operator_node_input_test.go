package types

import (
	"testing"
)

func TestAstGreaterThanOperatorNodeInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "gt_operator",
  "left": {
    "type": "string_literal",
    "value": "value"
  },
  "right": {
    "type": "string_literal",
    "value": "value"
  }
}`
	assertRoundTrip(t, wire, new(AstGreaterThanOperatorNodeInput))
}
