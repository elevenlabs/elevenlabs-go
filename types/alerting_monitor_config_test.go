package types

import (
	"testing"
)

func TestAlertingMonitorConfigRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "threshold": 0.5,
  "relative_increase_threshold": 9.99,
  "min_failure_count": 123,
  "min_history_bucket_count": 123,
  "min_sample_count": 123,
  "suspect_trigger_threshold": 10,
  "auto_resolve_after_inactive_minutes": 123
}`
	assertRoundTrip(t, wire, new(AlertingMonitorConfig))
}
