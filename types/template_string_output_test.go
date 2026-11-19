package types

import (
	"testing"
)

func TestTemplateStringOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "pending",
  "failure_reason": "timeout",
  "error_message": "value",
  "type": "string",
  "content": "value"
}`
	assertRoundTrip(t, wire, new(TemplateStringOutput))
}
