package types

import (
	"testing"
)

func TestSkipTurnToolConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "system_tool_type": "skip_turn",
  "wait_timeout_secs": 9.99
}`
	assertRoundTrip(t, wire, new(SkipTurnToolConfig))
}
