package types

import (
	"testing"
)

func TestPlatformCategoryUsageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "credits": 123,
  "price": 9.99,
  "quantity": 9.99
}`
	assertRoundTrip(t, wire, new(PlatformCategoryUsage))
}
