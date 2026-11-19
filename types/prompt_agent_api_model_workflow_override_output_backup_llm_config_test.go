package types

import (
	"testing"
)

func TestPromptAgentApiModelWorkflowOverrideOutputBackupLlmConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "preference": "default"
}`
	assertRoundTrip(t, wire, new(PromptAgentApiModelWorkflowOverrideOutputBackupLlmConfig))
}
