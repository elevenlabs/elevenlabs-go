package types

import (
	"testing"
)

func TestLanguageDetectionToolConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "system_tool_type": "language_detection",
  "only_at_conversation_start": true
}`
	assertRoundTrip(t, wire, new(LanguageDetectionToolConfig))
}
