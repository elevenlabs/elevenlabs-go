package types

import (
	"testing"
)

func TestToolAnnotationsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "title": "value",
  "readOnlyHint": true,
  "destructiveHint": true,
  "idempotentHint": true,
  "openWorldHint": true
}`
	assertRoundTrip(t, wire, new(ToolAnnotations))
}
