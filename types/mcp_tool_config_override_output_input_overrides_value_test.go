package types

import (
	"testing"
)

func TestMcpToolConfigOverrideOutputInputOverridesValueRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "constant",
  "constant_value": "value"
}`
	assertRoundTrip(t, wire, new(McpToolConfigOverrideOutputInputOverridesValue))
}
