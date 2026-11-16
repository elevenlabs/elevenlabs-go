package types

import (
	"testing"
)

func TestLlmListResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "llms": [
    {
      "llm": "gpt-4o-mini",
      "is_checkpoint": true,
      "max_tokens_limit": 123,
      "max_context_limit": 123,
      "supports_image_input": true,
      "supports_document_input": true,
      "supports_parallel_tool_calls": true,
      "available_reasoning_efforts": [
        "none"
      ],
      "available_service_tiers": [
        "standard"
      ],
      "deprecation_info": {
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
      },
      "regional_processing_surcharge": {
        "multiplier": 9.99
      }
    }
  ],
  "default_deprecation_config": {
    "warning_start_days": 123,
    "fallback_start_days": 123,
    "fallback_complete_days": 123,
    "fallback_start_percentage": 123,
    "fallback_complete_percentage": 123
  }
}`
	assertRoundTrip(t, wire, new(LlmListResponseModel))
}
