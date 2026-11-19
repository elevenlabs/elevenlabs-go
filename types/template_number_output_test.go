package types

import (
	"testing"
)

func TestTemplateNumberOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "pending",
  "failure_reason": "timeout",
  "error_message": "value",
  "type": "number",
  "content": 9.99
}`
	assertRoundTrip(t, wire, new(TemplateNumberOutput))
}
