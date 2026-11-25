package types

import (
	"testing"
)

func TestLlmCategoryUsageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "irreversible_generation": {
    "model_usage": {
      "key": {
        "input": {
          "tokens": 123,
          "price": 9.99
        },
        "input_cache_read": {
          "tokens": 123,
          "price": 9.99
        },
        "input_cache_write": {
          "tokens": 123,
          "price": 9.99
        },
        "output_total": {
          "tokens": 123,
          "price": 9.99
        },
        "output_total_reasoning_portion": {
          "tokens": 123,
          "price": 9.99
        }
      }
    }
  },
  "initiated_generation": {
    "model_usage": {
      "key": {
        "input": {
          "tokens": 123,
          "price": 9.99
        },
        "input_cache_read": {
          "tokens": 123,
          "price": 9.99
        },
        "input_cache_write": {
          "tokens": 123,
          "price": 9.99
        },
        "output_total": {
          "tokens": 123,
          "price": 9.99
        },
        "output_total_reasoning_portion": {
          "tokens": 123,
          "price": 9.99
        }
      }
    }
  }
}`
	assertRoundTrip(t, wire, new(LlmCategoryUsage))
}
