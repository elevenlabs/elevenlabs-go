package types

import (
	"testing"
)

func TestPendingBlocksMetadataModelRoundTrip(t *testing.T) {
	t.Parallel()
	const wire = `{
  "target_global_offset_ms": 123,
  "block_ids": [
    "value"
  ]
}`
	assertRoundTrip(t, wire, new(PendingBlocksMetadataModel))
}
