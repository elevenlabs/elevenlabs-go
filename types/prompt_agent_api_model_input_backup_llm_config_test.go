package types

import (
	"testing"
)

func TestPromptAgentApiModelInputBackupLlmConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "preference": "default"
}`
	assertRoundTrip(t, wire, new(PromptAgentApiModelInputBackupLlmConfig))
}
