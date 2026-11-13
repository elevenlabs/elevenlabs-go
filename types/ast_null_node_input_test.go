package types

import (
	"testing"
)

func TestAstNullNodeInputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "null_literal"
}`
	assertRoundTrip(t, wire, new(AstNullNodeInput))
}
