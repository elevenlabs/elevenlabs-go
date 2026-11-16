package types

import (
	"testing"
)

func TestUpdateDocumentChunkingOverrideResponseModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "id": "value",
  "chunking_override": {
    "max_chunk_size": 123,
    "overlap_percentage": 123,
    "custom_separators": [
      "value"
    ]
  },
  "reindex_queued": true
}`
	assertRoundTrip(t, wire, new(UpdateDocumentChunkingOverrideResponseModel))
}
