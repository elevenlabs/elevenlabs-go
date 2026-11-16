package types

import (
	"testing"
)

func TestFileRefreshStatusRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "queued",
  "enqueued_at": 123,
  "started_at": 123,
  "completed_at": 123,
  "last_synced_at": 123,
  "error_message": "value"
}`
	assertRoundTrip(t, wire, new(FileRefreshStatus))
}

func TestFileRefreshStatusRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "status": "queued",
  "enqueued_at": 123,
  "started_at": 123,
  "completed_at": 123,
  "last_synced_at": 123,
  "error_message": "value"
}`
	assertRoundTrip(t, wire, new(FileRefreshStatus))
}
