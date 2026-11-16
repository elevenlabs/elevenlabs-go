package types

import (
	"testing"
)

func TestDashboardDataCollectionChartModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "type": "data_collection",
  "data_collection_id": "value"
}`
	assertRoundTrip(t, wire, new(DashboardDataCollectionChartModel))
}
