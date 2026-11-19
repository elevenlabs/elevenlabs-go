package types

import (
	"testing"
)

func TestComputedUsagePlatformLimitRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "limit": 123
}`
	assertRoundTrip(t, wire, new(ComputedUsagePlatformLimit))
}
