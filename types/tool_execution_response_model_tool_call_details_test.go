package types

import (
	"testing"
)

func TestToolExecutionResponseModelToolCallDetailsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "webhook",
  "method": "value",
  "url": "value",
  "headers": {
    "key": "value"
  },
  "path_params": {
    "key": "value"
  },
  "query_params": {
    "key": "value"
  },
  "body": "value"
}`
	assertRoundTrip(t, wire, new(ToolExecutionResponseModelToolCallDetails))
}
