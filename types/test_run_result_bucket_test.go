package types

import (
	"testing"
)

func TestTestRunResultBucketRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "test_run_ids": [
    "value"
  ],
  "title": "value",
  "reason": "value",
  "status": "pending"
}`
	assertRoundTrip(t, wire, new(TestRunResultBucket))
}
