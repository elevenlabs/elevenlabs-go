package types

import (
	"testing"
)

func TestTestsFeatureUsageCommonModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "tests_ran_after_last_modification": true,
  "tests_ran_in_last_7_days": true
}`
	assertRoundTrip(t, wire, new(TestsFeatureUsageCommonModel))
}
