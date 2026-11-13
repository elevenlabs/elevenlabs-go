package types

import (
	"testing"
)

func TestLlmSchemaOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "llm",
  "prompt": "value"
}`
	assertRoundTrip(t, wire, new(LlmSchemaOverride))
}

func TestLlmSchemaOverrideRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "llm",
  "prompt": "value"
}`
	assertRoundTrip(t, wire, new(LlmSchemaOverride))
}

func TestLlmSchemaOverrideRoundTrip3(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "llm",
  "prompt": "value"
}`
	assertRoundTrip(t, wire, new(LlmSchemaOverride))
}
