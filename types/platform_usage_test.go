package types

import (
	"testing"
)

func TestPlatformUsageRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "category_usage": {
    "key": {
      "credits": 123,
      "price": 9.99,
      "quantity": 9.99
    }
  }
}`
	assertRoundTrip(t, wire, new(PlatformUsage))
}
