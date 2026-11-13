package types

import (
	"testing"
)

func TestPromptAgentDbModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tools": {}
}`
	assertRoundTrip(t, wire, new(PromptAgentDbModel))
}
