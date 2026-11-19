package agents

import (
	"testing"
)

func TestLlmUsageCalculatorRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "prompt_length": 123,
  "number_of_pages": 123,
  "rag_enabled": true,
  "service_tier": "standard"
}`
	assertRoundTrip(t, wire, new(LlmUsageCalculatorRequestModel))
}

func TestLlmUsageCalculatorPublicRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "prompt_length": 123,
  "number_of_pages": 123,
  "rag_enabled": true,
  "service_tier": "standard"
}`
	assertRoundTrip(t, wire, new(LlmUsageCalculatorPublicRequestModel))
}
