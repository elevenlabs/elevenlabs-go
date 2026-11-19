package types

import (
	"testing"
)

func TestAlertingSettingsRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "monitor_configs": {
    "key": {
      "enabled": true,
      "threshold": 0.5,
      "relative_increase_threshold": 9.99,
      "min_failure_count": 123,
      "min_history_bucket_count": 123,
      "min_sample_count": 123,
      "suspect_trigger_threshold": 10,
      "auto_resolve_after_inactive_minutes": 123
    }
  },
  "auto_resolve_after_inactive_minutes": 123,
  "notifiers": [
    {
      "type": "webhook",
      "webhook_id": "value"
    }
  ]
}`
	assertRoundTrip(t, wire, new(AlertingSettings))
}
