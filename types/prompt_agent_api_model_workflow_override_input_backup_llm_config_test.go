package types

import (
	"testing"
)

func TestPromptAgentApiModelWorkflowOverrideInputBackupLlmConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "preference": "default"
}`
	assertRoundTrip(t, wire, new(PromptAgentApiModelWorkflowOverrideInputBackupLlmConfig))
}
