package types

import (
	"testing"
)

func TestLlmUsageCalculatorResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "llm_prices": [
    {
      "llm": "gpt-4o-mini",
      "price_per_minute": 9.99,
      "price_per_message": 9.99
    }
  ]
}`
	assertRoundTrip(t, wire, new(LlmUsageCalculatorResponseModel))
}
