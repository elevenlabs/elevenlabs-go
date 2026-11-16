package types

import (
	"testing"
)

func TestLlmDeprecationConfigModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "warning_start_days": 123,
  "fallback_start_days": 123,
  "fallback_complete_days": 123,
  "fallback_start_percentage": 123,
  "fallback_complete_percentage": 123
}`
	assertRoundTrip(t, wire, new(LlmDeprecationConfigModel))
}
