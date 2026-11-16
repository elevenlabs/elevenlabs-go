package types

import (
	"testing"
)

func TestToolMockConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "default_return_value": "value",
  "default_is_error": true
}`
	assertRoundTrip(t, wire, new(ToolMockConfig))
}
