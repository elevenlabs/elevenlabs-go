package types

import (
	"testing"
)

func TestLlmDeprecationInfoModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "llm": "gpt-4o-mini",
  "is_deprecated": true,
  "is_in_warning_period": true,
  "is_in_fallback_period": true,
  "fallback_percentage": 123,
  "provider_deprecation_date": "2024-01-01T00:00:00Z",
  "replacement_model": "gpt-4o-mini",
  "deprecation_config": {
    "warning_start_days": 123,
    "fallback_start_days": 123,
    "fallback_complete_days": 123,
    "fallback_start_percentage": 123,
    "fallback_complete_percentage": 123
  }
}`
	assertRoundTrip(t, wire, new(LlmDeprecationInfoModel))
}
