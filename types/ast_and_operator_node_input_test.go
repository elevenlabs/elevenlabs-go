package types

import (
	"testing"
)

func TestAstAndOperatorNodeInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "and_operator",
  "children": [
    {
      "type": "string_literal",
      "value": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(AstAndOperatorNodeInput))
}
