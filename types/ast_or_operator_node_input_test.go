package types

import (
	"testing"
)

func TestAstOrOperatorNodeInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "or_operator",
  "children": [
    {
      "type": "string_literal",
      "value": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(AstOrOperatorNodeInput))
}
