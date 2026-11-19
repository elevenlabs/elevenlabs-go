package types

import (
	"testing"
)

func TestRegionConfigRequestRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "region_id": "us1",
  "token": "value",
  "edge_location": "ashburn"
}`
	assertRoundTrip(t, wire, new(RegionConfigRequest))
}
