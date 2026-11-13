package types

import (
	"testing"
)

func TestSampleConfigDbModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "is_sample": true,
  "parent_id": "value",
  "parent_type": "read",
  "chapter_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(SampleConfigDbModel))
}
