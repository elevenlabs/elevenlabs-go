package types

import (
	"testing"
)

func TestAgentTransferOpReplaceRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "type": "replace"
}`
	assertRoundTrip(t, wire, new(AgentTransferOpReplace))
}
