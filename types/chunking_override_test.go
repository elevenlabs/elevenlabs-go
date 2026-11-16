package types

import (
	"testing"
)

func TestChunkingOverrideRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "max_chunk_size": 123,
  "overlap_percentage": 123,
  "custom_separators": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ChunkingOverride))
}

func TestChunkingOverrideRoundTrip2(t *testing.T) {
	t.Parallel()
	const wire = `{
  "max_chunk_size": 123,
  "overlap_percentage": 123,
  "custom_separators": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(ChunkingOverride))
}
