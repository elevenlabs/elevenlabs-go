package types

import (
	"testing"
)

func TestTemplateBooleanOutputRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "status": "pending",
  "failure_reason": "timeout",
  "error_message": "value",
  "type": "boolean",
  "content": true
}`
	assertRoundTrip(t, wire, new(TemplateBooleanOutput))
}
