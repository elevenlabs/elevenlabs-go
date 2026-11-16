package types

import (
	"testing"
)

func TestSystemToolConfigInputParamsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "system_tool_type": "end_call"
}`
	assertRoundTrip(t, wire, new(SystemToolConfigInputParams))
}
