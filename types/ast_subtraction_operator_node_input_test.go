package types

import (
	"testing"
)

func TestAstSubtractionOperatorNodeInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "sub_operator",
  "left": {
    "type": "string_literal",
    "value": "value"
  },
  "right": {
    "type": "string_literal",
    "value": "value"
  }
}`
	assertRoundTrip(t, wire, new(AstSubtractionOperatorNodeInput))
}
