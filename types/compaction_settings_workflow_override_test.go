package types

import (
	"testing"
)

func TestCompactionSettingsWorkflowOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "enabled": true,
  "soft_trigger_fraction": 9.99,
  "tail_size": 123,
  "min_reclaimable_tokens": 123
}`
	assertRoundTrip(t, wire, new(CompactionSettingsWorkflowOverride))
}
