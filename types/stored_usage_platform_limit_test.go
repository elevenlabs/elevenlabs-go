package types

import (
	"testing"
)

func TestStoredUsagePlatformLimitRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "limit": 123,
  "usage": 123
}`
	assertRoundTrip(t, wire, new(StoredUsagePlatformLimit))
}
