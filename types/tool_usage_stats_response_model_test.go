package types

import (
	"testing"
)

func TestToolUsageStatsResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "total_calls": 123,
  "avg_latency_secs": 9.99
}`
	assertRoundTrip(t, wire, new(ToolUsageStatsResponseModel))
}
