package types

import (
	"testing"
)

func TestToolExecutionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "taskSupport": "forbidden"
}`
	assertRoundTrip(t, wire, new(ToolExecution))
}
