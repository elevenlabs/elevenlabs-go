package types

import (
	"testing"
)

func TestAstNullNodeOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "null_literal"
}`
	assertRoundTrip(t, wire, new(AstNullNodeOutput))
}
