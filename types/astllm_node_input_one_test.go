package types

import (
	"testing"
)

func TestAstllmNodeInputOneRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "llm",
  "prompt": "value"
}`
	assertRoundTrip(t, wire, new(AstllmNodeInputOne))
}
