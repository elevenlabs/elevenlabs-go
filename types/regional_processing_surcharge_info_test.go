package types

import (
	"testing"
)

func TestRegionalProcessingSurchargeInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "multiplier": 9.99
}`
	assertRoundTrip(t, wire, new(RegionalProcessingSurchargeInfo))
}
