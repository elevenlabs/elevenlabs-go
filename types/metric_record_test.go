package types

import (
	"testing"
)

func TestMetricRecordRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "elapsed_time": 9.99
}`
	assertRoundTrip(t, wire, new(MetricRecord))
}
