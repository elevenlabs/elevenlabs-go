package types

import (
	"testing"
)

func TestFeatureStatusCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "used": true
}`
	assertRoundTrip(t, wire, new(FeatureStatusCommonModel))
}
