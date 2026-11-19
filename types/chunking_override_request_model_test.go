package types

import (
	"testing"
)

func TestChunkingOverrideRequestModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "max_chunk_size": 200,
  "overlap_percentage": 25,
  "custom_separators": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ChunkingOverrideRequestModel))
}
