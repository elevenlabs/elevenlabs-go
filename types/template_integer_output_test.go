package types

import (
	"testing"
)

func TestTemplateIntegerOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "pending",
  "failure_reason": "timeout",
  "error_message": "value",
  "type": "integer",
  "content": 123
}`
	assertRoundTrip(t, wire, new(TemplateIntegerOutput))
}
