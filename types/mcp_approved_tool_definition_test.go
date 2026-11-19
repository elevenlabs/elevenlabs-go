package types

import (
	"testing"
)

func TestMcpApprovedToolDefinitionRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "description": "value",
  "input_schema": {
    "key": {}
  }
}`
	assertRoundTrip(t, wire, new(McpApprovedToolDefinition))
}
