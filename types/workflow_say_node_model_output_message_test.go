package types

import (
	"testing"
)

func TestWorkflowSayNodeModelOutputMessageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "literal",
  "text": "value",
  "text_translations": {
    "key": {
      "value": "value"
    }
  }
}`
	assertRoundTrip(t, wire, new(WorkflowSayNodeModelOutputMessage))
}
