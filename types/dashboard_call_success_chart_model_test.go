package types

import (
	"testing"
)

func TestDashboardCallSuccessChartModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "name": "value",
  "type": "call_success"
}`
	assertRoundTrip(t, wire, new(DashboardCallSuccessChartModel))
}
