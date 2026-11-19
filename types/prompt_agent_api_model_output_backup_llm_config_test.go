package types

import (
	"testing"
)

func TestPromptAgentApiModelOutputBackupLlmConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "preference": "default"
}`
	assertRoundTrip(t, wire, new(PromptAgentApiModelOutputBackupLlmConfig))
}
