package types

import (
	"testing"
)

func TestAttachedUserEvaluationRefRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "user",
  "analysis_item_id": "value",
  "version_id": "value",
  "additional_version_ids": [
    "value"
  ],
  "scope": "conversation",
  "weight": 9.99
}`
	assertRoundTrip(t, wire, new(AttachedUserEvaluationRef))
}
