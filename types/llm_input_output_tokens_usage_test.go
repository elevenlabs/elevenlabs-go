package types

import (
	"testing"
)

func TestLlmInputOutputTokensUsageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
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
  }
}`
	assertRoundTrip(t, wire, new(LlmInputOutputTokensUsage))
}

func TestLlmInputOutputTokensUsageRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
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
  }
}`
	assertRoundTrip(t, wire, new(LlmInputOutputTokensUsage))
}
