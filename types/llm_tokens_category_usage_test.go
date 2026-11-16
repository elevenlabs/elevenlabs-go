package types

import (
	"testing"
)

func TestLlmTokensCategoryUsageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "tokens": 123,
  "price": 9.99
}`
	assertRoundTrip(t, wire, new(LlmTokensCategoryUsage))
}
