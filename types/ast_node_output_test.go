package types

import (
	"testing"
)

func TestAstNodeOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "string_literal",
  "value": "value"
}`
	assertRoundTrip(t, wire, new(AstNodeOutput))
}
