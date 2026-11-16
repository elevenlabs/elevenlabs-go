package types

import (
	"testing"
)

func TestDashboardCriteriaChartModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "type": "criteria",
  "criteria_id": "value"
}`
	assertRoundTrip(t, wire, new(DashboardCriteriaChartModel))
}
