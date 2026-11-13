package types

import (
	"testing"
)

func TestAttachedSystemDataCollectionRefRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "source": "system",
  "analysis_item_id": "__system_data_collection_topic",
  "scope": "conversation"
}`
	assertRoundTrip(t, wire, new(AttachedSystemDataCollectionRef))
}
