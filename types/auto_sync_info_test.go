package types

import (
	"testing"
)

func TestAutoSyncInfoRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "minimum_frequency_days": 123,
  "auto_remove": true,
  "consec_failures": 123,
  "next_refresh_by": 123
}`
	assertRoundTrip(t, wire, new(AutoSyncInfo))
}

func TestAutoSyncInfoRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "minimum_frequency_days": 123,
  "auto_remove": true,
  "consec_failures": 123,
  "next_refresh_by": 123
}`
	assertRoundTrip(t, wire, new(AutoSyncInfo))
}
